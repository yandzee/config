package main

//	type Transformy] struct {
//		Datum T
//	}

// type ValueTransformer[T any] struct {
// 	Value T
// }
//
// func (vt *ValueTransformer[T]) Split(splitter Transformer[T, []T]) []T {
// 	return splitter.Transform(vt.Value)
// }
//
// func (vt *ValueTransformer[T]) Step(t Transformer[T, T]) T {
// 	return t.Transform(vt.Value)
// }
//
// func (st *StringTransformer) Split(seps ...string) []string {
//
// }

// type StringSplitter struct {
// 	Separators []string
// }
//
// func Splitter(seps ...string) *StringSplitter {
// 	return &StringSplitter{
// 		Separators: seps,
// 	}
// }
//
// func (s *StringSplitter) Transform(v string) ([]string, error) {
// 	return []string{}, nil
// }
//
// func main() {
// 	str := "ydz;was,here:and there"
//
// 	pl := transform.Pipe2(
// 		Unhex(),
// 		Splitter(","),
// 	)
//
// 	transformed, err := pl.Transform(str)
// 	if err != nil {
// 		panic(err.Error())
// 	}
//
// 	fmt.Printf("transformed: %v\n", transformed)
// }

func main() {
}
