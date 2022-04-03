package main
 
import (
	//"flag"
	//"fmt"
	"time"
	"log"
 	"modbusd/rtu"
 	"modbusd/modbusx"
 	"context"
)
 
var SaveValue map[int]int
 
 
func main() {

	chWait := make(chan bool)
	//go getMessage(ctx)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rtu.GetRTU()
	//for _, v := range rtu.RTUs{
	//	log.Println(v)
	//}

	serv := modbusx.NewModbusServer(ctx)	
	time.Sleep(8 * time.Second)
	modbusx.NewModbusClient(ctx)
	<-chWait
	serv.Close()
	log.Println("main -----")
}

