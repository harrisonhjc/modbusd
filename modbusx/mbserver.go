package modbusx

import (
	//"flag"
	//"fmt"
	"log"
	"context"
	"time"
 	//"modbusd/rtu"

	"github.com/tbrandon/mbserver"
	"github.com/goburrow/serial"
)

func NewModbusServer(ctx context.Context) (serv *mbserver.Server) {
	
	serv = mbserver.NewServer()
	err := serv.ListenTCP("0.0.0.0:3502")
	if err != nil {
		log.Printf("%v\n", err)
	}
	
	err = serv.ListenRTU(&serial.Config{
		Address:  "/dev/ttys000",
		BaudRate: 115200,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
		Timeout:  10 * time.Second})
	if err != nil {
		log.Fatalf("failed to listen, got %v\n", err)
	}
	
	/*
	err = serv.ListenRTU(&serial.Config{
		Address:  "/dev/ttyACM0",
		BaudRate: 9600,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
		Timeout:  10 * time.Second,
		RS485: serial.RS485Config{
			Enabled: true,
			DelayRtsBeforeSend: 2 * time.Millisecond,
			DelayRtsAfterSend: 3 * time.Millisecond,
			RtsHighDuringSend: false,
			RtsHighAfterSend: false,
			RxDuringTx: false,
			},
		})
	if err != nil {
		log.Fatalf("failed to listen, got %v\n", err)
	}
	*/
	//defer serv.Close()
	return
}