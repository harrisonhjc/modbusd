package mbserver

import (
	//"flag"
	//"fmt"
	"context"
	"log"
	"time"

	"github.com/goburrow/serial"
	//"modbusd/rtu"
)

func NewModbusServer(ctx context.Context) (serv *Server) {

	serv = NewServer()
	err := serv.ListenTCP("0.0.0.0:3502")
	if err != nil {
		log.Printf("%v\n", err)
	}
	err = serv.ListenTCP("0.0.0.0:1502")
	if err != nil {
		log.Printf("%v\n", err)
	}

	/*
			serv.RegisterFunctionHandler(2,
		    func(s *Server, frame Framer) ([]byte, *Exception) {
		        register, numRegs, endRegister := registerAddressAndNumber(frame)
		        // Check the request is within the allocated memory
		        if endRegister > 65535 {
		            return []byte{}, &IllegalDataAddress
		        }
		        dataSize := numRegs / 8
		        if (numRegs % 8) != 0 {
		            dataSize++
		        }
		        data := make([]byte, 1+dataSize)
		        data[0] = byte(dataSize)
		        for i := range s.DiscreteInputs[register:endRegister] {
		            // Return all 1s, regardless of the value in the DiscreteInputs array.
		            shift := uint(i) % 8
		            data[1+i/8] |= byte(1 << shift)
		        }
		        return data, &Success
		    })
	*/
	//set discrete input value
	// var i byte
	// i = 1
	// serv.DiscreteInputs[i] = i

	err = serv.ListenRTU(&serial.Config{
		Address:  "/dev/ttyAMA0",
		BaudRate: 9600,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
		Timeout:  30 * time.Second})
	if err != nil {
		log.Fatalf("failed to listen, got %v\n", err)
	}

	/*err = serv.ListenRTU(&serial.Config{
		Address:  "/dev/ttyAMA0",
		BaudRate: 9600,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
		Timeout:  10 * time.Second,
		RS485: serial.RS485Config{
			Enabled:            true,
			DelayRtsBeforeSend: 2 * time.Millisecond,
			DelayRtsAfterSend:  3 * time.Millisecond,
			RtsHighDuringSend:  false,
			RtsHighAfterSend:   false,
			RxDuringTx:         false,
		},
	})
	if err != nil {
		log.Fatalf("failed to listen, got %v\n", err)
	}*/

	//defer serv.Close()
	return
}
