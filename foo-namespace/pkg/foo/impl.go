package foo

type Impl struct{}

//export foo-namespace:pkg/foo@2.0.0#greet
func (i *Impl) Greet() string {
	return "hello from greet"
}
