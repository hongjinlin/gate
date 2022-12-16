package main

import (
	"gate/enums"
	"gate/service"
	"gate/sys"
	"log"
)

func main() {
	mode, b := enums.ParseMode(sys.MODE)
	if !b {
		log.Fatalf("mode error")
	}
	switch mode {
	case enums.Check:
		c := new(service.CheckService)
		c.Check()
	default:
		s := new(service.GateService)
		s.Gates()
	}

}
