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
	//3:ReadHoldingRegisters
	s.RegisterFunctionHandler(3, handleRegisters)

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

func handleRegisters(s *Server, frame Framer) ([]byte, *Exception){

	register, numRegs, endRegister := frame.registerAddressAndNumber()
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