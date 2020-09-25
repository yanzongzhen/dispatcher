package dispatcher

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterDispatcher("random", &RandomDispatch{})
}

type RandomDispatch struct {
}

func (p *RandomDispatch) DoDispatch(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance ")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}
