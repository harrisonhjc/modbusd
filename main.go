package main
 
import (
	"flag"
	"fmt"
	"log"
 	"modbusd/rtu"
 	"modbusd/mbus_server"

	modbus "github.com/advancedclimatesystems/goldfish"
)
 
var SaveValue map[int]int
 
 
func main() {

	//go getMessage(ctx)
	rtu.GetRTU()
	for _, v := range rtu.RTUs{
		fmt.Println(v)
	}

	

}

