package dispatcher

import (
	"net"
	"strconv"
)

type Instance struct {
	host string
	port int
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

func (p *Instance) GetHost() string {
	return p.host
}

func (p *Instance) GetPort() int {
	return p.port
}

func (p *Instance) String() string {
	return p.host + ":" + strconv.Itoa(p.port)
}

func toInstance(server string) *Instance {
	var inst Instance
	ip, err := net.ResolveTCPAddr("", server)
	if err != nil {
		return nil
	}
	inst.host = ip.IP.String()
	inst.port = ip.Port
	return &inst
}

func NewInstances(servers ...string) []*Instance {
	instances := make([]*Instance, 0)
	for _, server := range servers {
		inst := toInstance(server)
		if inst != nil {
			instances = append(instances, inst)
		}
	}
	return instances
}
