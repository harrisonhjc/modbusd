package mbserver

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

func NewModbusServer(ctx context.Context){
	
	s := mbserver.NewServer()
	err := s.ListenRTU(&serial.Config{
		Address:  "/dev/serial1",
		BaudRate: 115200,
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

	defer s.Close()
}
/*
func handleRegisters(unitID, start, quantity int) ([]modbus.Value, error) {
	
	fmt.Printf("ID:%d, start:%d, quantity:%d\n",unitID, start, quantity)

	registers := make([]modbus.Value, quantity)
	for i := 0; i < quantity; i++ {
		if v, ok := rtu.RTUs[unitID]; ok {
    		registers[i], _ = modbus.NewValue(v.Value)
    	}
	}
	return registers, nil
}

func handleReadCoils(unitID, start, quantity int) ([]modbus.Value, error) {
	fmt.Println("ReadCoils")
	coils := make([]modbus.Value, quantity)
	for i := 0; i < quantity; i++ {
		v, err := modbus.NewValue((i + start) % 2)
		if err != nil {
			return coils, modbus.SlaveDeviceFailureError
		}
 
		coils[i] = v
	}
 
	return coils, nil
}
 

 
func handleWriteRegisters(unitID, start int, values []modbus.Value) error {
	fmt.Println("WriteRegisters")
	for i, value := range values {
		fmt.Printf("[%d]: %d\n", i+start, value.Get())
		SaveValue[i+start] = value.Get()
	}
 
	return nil
}
 
func handleWriteCoils(unitID, start int, values []modbus.Value) error {
	fmt.Println("WriteCoils")
	if start == 1 {
		return modbus.IllegalAddressError
	}
	return nil
}
*/