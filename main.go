package main
 
import (
	"flag"
	"fmt"
	"log"
 
	modbus "github.com/advancedclimatesystems/goldfish"
)
 
var SaveValue map[int]int
 
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
 
func handleRegisters(unitID, start, quantity int) ([]modbus.Value, error) {
	
	fmt.Printf("ID:%d, start:%d, quantity:%d\n",unitID, start, quantity)

	registers := make([]modbus.Value, quantity)
	for i := 0; i < quantity; i++ {
		registers[i], _ = modbus.NewValue(SaveValue[unitID*2 + start + i])
	}
 
	return registers, nil
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
 
func main() {
	addr := flag.String("addr", ":3502", "address to listen on.")
	flag.Parse()
 
	SaveValue = make(map[int]int)
	s, err := modbus.NewServer(*addr)
 
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to start Modbus server: %v", err))
	}

	for i := 0; i<200; i++{
		SaveValue[i] = 3 + i
	}
	//s.Handle(modbus.ReadCoils, modbus.NewReadHandler(handleReadCoils))
	s.Handle(modbus.ReadHoldingRegisters, modbus.NewReadHandler(handleRegisters))
	//s.Handle(modbus.WriteSingleCoil, modbus.NewWriteHandler(handleWriteCoils, modbus.Signed))
	//s.Handle(modbus.WriteSingleRegister, modbus.NewWriteHandler(handleWriteRegisters, modbus.Signed))
	//s.Handle(modbus.WriteMultipleRegisters, modbus.NewWriteHandler(handleWriteRegisters, modbus.Signed))
 
	s.Listen()

}