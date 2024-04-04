package incominghandler

func init() {
	instance = &Impl{}
}

type Impl struct{}

func (i *Impl) Handle() string {
	return "hello from greet in main.go"
}

//export foo-namespace:pkg/foo@2.0.0#greet
func GreetExported() *string {
	x := instance.Greet()
	return &x
}

func HandleExported(*IncomingRequest, *ResponseOutparam) {

}
