# Usage

## Built-in methods

```go
import "github.com/yandzee/config/configurator"

type Config struct {
	configurator.Configurator

	Port           uint16
	APIPrefix      string
	FilePaths      []string
	SignPrivateKey *ecdsa.PrivateKey
	DatabaseURL    string
	IsInMemoryMode bool
}

func (c *Config) Init() {
	c.Port = c.Env("PORT").Uint16Or(8080)
	c.APIPrefix = c.Env("API_PREFIX").StringOr("/api")
	c.FilePaths = c.Env("FILE_PATHS").StringsOr([]string{}, ";", ",")
	c.SignPrivateKey = c.Env("SIGNATURE_KEY").ECPrivateKey()

	c.IsInMemoryMode = c.Env("IN_MEMORY_MODE").BoolOr(false)
	c.DatabaseURL = c.Env("DATABASE_URL").StringOrFn(func() (string, error) {
		if !c.IsInMemoryMode {
			return "", fmt.Errorf("Database url must be specified unless in memory mode is on")
		}

		return "", nil
	})
}
```

After this code executed, you can inspect the result either directly by checking
`c.Configurator.ValueResults` or by getting `c.LogRecords()`:

```go
func main() {
	cfg := Config{}
	cfg.Init()

	log := slog.Default()

	hasFatal := false
	for _, logRecord := range cfg.LogRecords() {
		if err := log.Handler().Handle(context.Background(), logRecord); err != nil {
			panic(err.Error())
		}

		hasFatal = hasFatal || logRecord.Level == slog.LevelError
	}

	if hasFatal {
		os.Exit(1)
	}
}
```

Output:
```
➜  config-usage-ex go run main.go
2025/02/27 18:55:48 WARN Value set name=PORT kind=env is-defaulted=true value=8080
2025/02/27 18:55:48 WARN Value set name=API_PREFIX kind=env is-defaulted=true value=/api
2025/02/27 18:55:48 WARN Value set name=FILE_PATHS kind=env is-defaulted=true value=[]
2025/02/27 18:55:48 ERROR Not set name=SIGNATURE_KEY kind=env is-required=true value=<nil>
2025/02/27 18:55:48 WARN Value set name=IN_MEMORY_MODE kind=env is-defaulted=true value=false
2025/02/27 18:55:48 ERROR Database url must be specified unless in memory mode is on name=DATABASE_URL kind=env is-custom-error=true value=""
2025/02/27 18:55:48 ERROR Not set name=B64_TWO_STRINGS kind=env is-required=true value=""
exit status 1
```

## Custom parsing logic

```go
import (
	"github.com/yandzee/config/transform"
	"github.com/yandzee/config/transformers"
)

func (c *Config) Init() {
	...

	c.Env("B64_TWO_STRINGS").String(
		transformers.Unbase64,
		transformers.Split(","),
		transform.Map(func(splited []string) (string, error) {
			if len(splited) < 2 {
				return "", fmt.Errorf("Value must have at least two string comma separated")
			}

			return splited[1], nil
		}),
	)
}
```

By calling ```cfg.LogRecords(configurator.LogWithValue)``` you can see the actual value
assigned after initialization:

```
➜  config-usage-ex B64_TWO_STRINGS=b25lLHR3bw== go run main.go
2025/02/27 18:56:20 WARN Value set name=PORT kind=env is-defaulted=true value=8080
2025/02/27 18:56:20 WARN Value set name=API_PREFIX kind=env is-defaulted=true value=/api
2025/02/27 18:56:20 WARN Value set name=FILE_PATHS kind=env is-defaulted=true value=[]
2025/02/27 18:56:20 ERROR Not set name=SIGNATURE_KEY kind=env is-required=true value=<nil>
2025/02/27 18:56:20 WARN Value set name=IN_MEMORY_MODE kind=env is-defaulted=true value=false
2025/02/27 18:56:20 ERROR Database url must be specified unless in memory mode is on name=DATABASE_URL kind=env is-custom-error=true value=""
2025/02/27 18:56:20 INFO Value set name=B64_TWO_STRINGS kind=env is-required=true is-presented=true value=two
exit status 1
```
