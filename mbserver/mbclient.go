package mbserver

import (
	//"flag"
	//"fmt"
	"context"
	"log"

	//"time"
	//"modbusd/rtu"

	//"github.com/tbrandon/mbserver"
	//"github.com/goburrow/serial"
	"github.com/goburrow/modbus"
)

func WriteRegisters(ctx context.Context, address uint16, count uint16, value []byte) {

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
func WriteCoils(ctx context.Context, address uint16, count uint16, value []byte) {

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

func ReadRegisters(ctx context.Context, address uint16, count uint16) []byte {
	handler := modbus.NewTCPClientHandler("localhost:3502")
	// Connect manually so that multiple requests are handled in one session
	err := handler.Connect()
	defer handler.Close()
	client := modbus.NewClient(handler)

	results, err := client.ReadHoldingRegisters(address, count)
	if err != nil {
		log.Printf("%v\n", err)
	}
	log.Printf("results %v\n", results)
	return results
}
