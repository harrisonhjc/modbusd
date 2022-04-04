package modbusx

import (
	//"flag"
	//"fmt"
	"log"
	"context"
	//"time"
 	//"modbusd/rtu"

	//"github.com/tbrandon/mbserver"
	//"github.com/goburrow/serial"
	"github.com/goburrow/modbus"
)

func WriteRegisters(ctx context.Context,address int, count int, value []byte) {
	
	handler := modbus.NewTCPClientHandler("localhost:3502")
	// Connect manually so that multiple requests are handled in one session
	err := handler.Connect()
	defer handler.Close()
	client := modbus.NewClient(handler)

	_, err = client.WriteMultipleRegisters(address, count, value)
	if err != nil {
		log.Printf("%v\n", err)
	}

	results, err := client.ReadHoldingRegisters(address, count)
	if err != nil {
		log.Printf("%v\n", err)
	}
	log.Printf("results %v\n", results)
	return
}