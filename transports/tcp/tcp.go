package tcp

import (
	"fmt"
	"net"

	"github.com/phedoreanu/logspout/adapters/raw"
	"github.com/phedoreanu/logspout/router"
)

func init() {
	router.AdapterTransports.Register(new(tcpTransport), "tcp")
	// convenience adapters around raw adapter
	router.AdapterFactories.Register(rawTCPAdapter, "tcp")
}

func rawTCPAdapter(route *router.Route) (router.LogAdapter, error) {
	route.Adapter = "raw+tcp"
	return raw.NewRawAdapter(route)
}

type tcpTransport int

func (_ *tcpTransport) Dial(addr string, options map[string]string) (net.Conn, error) {
	fmt.Printf("# addr: %s\n", addr)
	raddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Printf("# raddr: %s\n", raddr.String())
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return conn, nil
}
