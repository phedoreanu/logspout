package main

import (
	_ "github.com/phedoreanu/logspout/adapters/raw"
	_ "github.com/phedoreanu/logspout/adapters/syslog"
	_ "github.com/phedoreanu/logspout/httpstream"
	_ "github.com/phedoreanu/logspout/routesapi"
	_ "github.com/phedoreanu/logspout/transports/tcp"
	_ "github.com/phedoreanu/logspout/transports/tls"
	_ "github.com/phedoreanu/logspout/transports/udp"
)
