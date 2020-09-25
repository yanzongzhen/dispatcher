package main

import (
	"fmt"
	"github.com/yanzongzhen/dispatcher/dispatcher"
	"math/rand"
	"os"
	"time"
)

func main() {
	var insts []*dispatcher.Instance
	var hosts []string
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d:8080", rand.Intn(255), rand.Intn(255))
		hosts = append(hosts, host)
	}
	insts = dispatcher.NewInstances(hosts...)
	var balanceName = "hash"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}
	for {
		inst, err := dispatcher.DoDispatch(balanceName, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			fmt.Fprintf(os.Stdout, "do balance err\n")
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}