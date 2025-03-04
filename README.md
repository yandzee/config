# Config [![Go Reference](https://pkg.go.dev/badge/github.com/yandzee/config.svg)](https://pkg.go.dev/github.com/yandzee/config)

Extensible utility for initializing fields / configs with values, taken from
[StringSources](https://github.com/yandzee/config/blob/c04fb38e63c0b62f4dabd87e47ee08bb993dd5fd/configurator/getter.go#L77),
like environment variables. Allows to easily get `[]slog.LogRecord` for logging purposes.
Convenient, frequently used
[value checkers](https://github.com/yandzee/config/blob/c04fb38e63c0b62f4dabd87e47ee08bb993dd5fd/checkers/exports.go) and
[value transformers](https://github.com/yandzee/config/blob/c04fb38e63c0b62f4dabd87e47ee08bb993dd5fd/transformers/transformers.go)
included.

## Basic usage

```go
import "github.com/yandzee/config"

func main() {
	cfg := AppConfig{}

	// Return value form
	cfg.Port := config.Uint16().Env("PORT", 8080)
	// Note that optional second argument is used to specify the default value

	// Generic setter form
	config.Set(&cfg.Port).Env("PORT", 8080)
}
```

## Value checks

Validators (checkers) can be engaged using appropriate
[methods](https://github.com/yandzee/config/blob/adcfb7550acdd417cfb2c37d65a4353d2e00d681/configurator/getter.go#L176-L184)
on value getters:

```go
import (
	"github.com/yandzee/config"
	"github.com/yandzee/config/checkers"
)

type Config struct {
	TLSCerts []string
}

func (c *Config) Init() {
	c.TLSCerts = config.
		Strings(",", ";"). // Multiple separators can be used for string split
		Checks(checkers.FilesExist).
		Env("TLS_CERTS", []string{}) // Default value as an optional second argument
}
```

#### Example: logic for connected values

Sometimes variables are connected logically and should be checked appropriately:

```go
type Config struct {
	DatabaseURL    string
	IsInMemoryMode bool
}

func (c *Config) Init() {
	c.IsInMemoryMode = config.Bool().Env("IN_MEMORY_MODE", false)

	c.DatabaseURL = config.
		String().
		Check(func(r *result.Result[string]) (bool, string) {
			// Custom check function for wiring up logically connected values
			if c.IsInMemoryMode || !r.Flags.IsDefaulted() {
				return true, ""
			}

			return false, "Must be specified unless IN_MEMORY_MODE is on"
		}).
		Env("DATABASE_URL", "")
}
```

## Value transformers

Some values may require additional preprocessing before being stored in config:

```go
import (
	"github.com/yandzee/config"
	"github.com/yandzee/config/checkers"
)

type Config struct {
	AESKey []byte
}

func (c *Config) Init() {
	c.AESKey = config.Bytes().
        Pre(transformers.Unhex).  // <--- hex.DecodeString()
        Env("AES_KEY")
}
```

#### Example: custom transformers

Custom getters with custom transformers can be easily written by hand and reused
as it was there built in.

```go
import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/yandzee/config"
	"github.com/yandzee/config/configurator"
	"github.com/yandzee/config/transform"
	"github.com/yandzee/config/transformers"
)

type Config struct {
	SignPrivateKey *ecdsa.PrivateKey
}

func (c *Config) Init() {
	c.SignPrivateKey = c.ECPrivateKey(). // <--- Custom getter
		Pre(transformers.Unbase64).
		Env("SIGNATURE_PRIVATE_KEY")
}

func (c *Config) ECPrivateKey() *configurator.Getter[*ecdsa.PrivateKey] {
	return config.
		Custom[*ecdsa.PrivateKey]().
		Post(
			transformers.ToBytes,
			transform.Map(func(v []byte) (*ecdsa.PrivateKey, error) {
				block, _ := pem.Decode([]byte(v))
				if block == nil {
					return nil, fmt.Errorf("PEM block is not found")
				}

				pk, err := x509.ParseECPrivateKey(block.Bytes)
				if err != nil {
					return nil, errors.Join(
						fmt.Errorf("Failed to x509.ParseECPrivateKey"),
						err,
					)
				}

				return pk, nil
			}),
		)
}
```

## Logging

Library doesn't force the program to log anything. Instead, `[]slog.Record` can
be obtained for further handling:

```go
func main() {
	cfg := Config{}
	cfg.Init()

	log := slog.Default()
	fctx := context.Background()

	// Logging is decoupled and managed by your code
	hasFatal := false
	for _, logRecord := range config.LogRecords() {
		_ = log.Handler().Handle(ctx, logRecord)

		hasFatal = hasFatal || logRecord.Level == slog.LevelError
	}

	if hasFatal {
		os.Exit(1)
	}
}
```

Note that by default, no value is included as log record attribute for security reasons.
It can be changed by appropriate
[option](https://github.com/yandzee/config/blob/1c3966d792e5834fe50a8d79b76fdc6bcf4c4cc3/configurator/configurator.go#L11)
for `config.LogRecords()` method.

Example output:
```
➜  config-usage-ex go run main.go
2025/02/27 18:55:48 WARN Value set name=PORT kind=env is-defaulted=true value=8080
2025/02/27 18:55:48 WARN Value set name=TLS_CERTS kind=env is-defaulted=true value=[]
2025/02/27 18:55:48 ERROR Not set name=SIGNATURE_KEY kind=env is-required=true value=<nil>
2025/02/27 18:55:48 WARN Value set name=IN_MEMORY_MODE kind=env is-defaulted=true value=false
2025/02/27 18:55:48 ERROR Database url must be specified unless in memory mode is on name=DATABASE_URL kind=env is-check-failed=true value=""
exit status 1
```
