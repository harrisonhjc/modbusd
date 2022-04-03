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

func NewModbusClient(ctx context.Context) {
	
	handler := modbus.NewTCPClientHandler("localhost:3502")
	// Connect manually so that multiple requests are handled in one session
	err := handler.Connect()
	defer handler.Close()
	client := modbus.NewClient(handler)

	_, err = client.WriteMultipleRegisters(8, 3, []byte{0, 3, 0, 4, 0, 5})
	if err != nil {
		log.Printf("%v\n", err)
	}

	results, err := client.ReadHoldingRegisters(8, 3)
	if err != nil {
		log.Printf("%v\n", err)
	}
	log.Printf("results %v\n", results)
	return
}