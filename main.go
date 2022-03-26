package main
 
import (
	//"flag"
	"fmt"
	//"log"
 	"modbusd/rtu"
 	"modbusd/mbserver"
 	"context"
)
 
var SaveValue map[int]int
 
 
func main() {

	//go getMessage(ctx)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rtu.GetRTU()
	for _, v := range rtu.RTUs{
		fmt.Println(v)
	}

	mbserver.NewModbusServer(ctx)	

}

