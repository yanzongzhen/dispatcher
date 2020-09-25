package dispatcher

import (
	"errors"
)

func init() {
	RegisterDispatcher("roundrobin", &RoundRobinDispatch{})
}

type RoundRobinDispatch struct {
	curIndex int
}

func (p *RoundRobinDispatch) DoDispatch(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance ")
		return
	}
	lens := len(insts)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens
	return
}
