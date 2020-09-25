package dispatcher

import "fmt"

type DispatchMgr struct {
	allDispatcher map[string]Dispatcher
}

var mgr = DispatchMgr{
	allDispatcher: make(map[string]Dispatcher),
}

func (p *DispatchMgr) registerDispatcher(name string, b Dispatcher) {
	p.allDispatcher[name] = b
}

func RegisterDispatcher(name string, b Dispatcher) {
	mgr.registerDispatcher(name, b)
}

func DoDispatch(name string, insts []*Instance) (inst *Instance, err error) {
	Dispatcher, ok := mgr.allDispatcher[name]
	if !ok {
		err = fmt.Errorf("Not found %s Dispatcher ", name)
		return
	}
	fmt.Printf("use %s Dispatcher\n", name)
	inst, err = Dispatcher.DoDispatch(insts)
	return
}
