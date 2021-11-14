package main

import (
	"github.com/mfsyahrz/bareksaproject/internal/interface/ioc"
	"github.com/mfsyahrz/bareksaproject/internal/interface/server/rest"
)

func main() {
	rest.StartRestServer(ioc.Setup())
}
