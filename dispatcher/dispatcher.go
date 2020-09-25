package dispatcher

type Dispatcher interface {
	DoDispatch([]*Instance, ...string) (*Instance, error)
}