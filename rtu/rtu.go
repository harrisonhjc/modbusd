package rtu
 
import (
	//"fmt"
	//"log"
	//"os"
	//"strconv"
	//"strings"
)

type RTU struct {
	Cmd		string
	System  int 
	Address int
	Loop    int 
	Value   int
	Code    int
	Modbus_addr int 
}

type RTUS struct{
	RTUs map[string]RTU
}

func(rtus *RTUS) NewRTUs() *RTUS {
	rs := make(map[string]RTU)
	r:= RTUS{
		RTUs : rs,
	}
	return &r
}

func(rtus *RTUS) SetModbusAddress(data RTU) {

	
	switch data.Cmd{
	case "J0":

	case "J1":
	case "C0":
	case "F0":
	case "F1":
	case "T0":
	case "B0":
	default:
		return
	}
	
	
}