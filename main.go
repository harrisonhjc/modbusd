package main
 
import (
	//"flag"
	//"fmt"
	"time"
	"log"
	"context"
 	"modbusd/rtu"
 	"modbusd/modbusx"
 	"modbusd/msgserver"
 	
)
 
var SaveValue map[int]int
 
 
func main() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	chData := make(chan RTU)
	defer close(chData)
	//go getMessage(ctx)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rtu.GetRTU()

	msgserver.NewServer(8081, chData)
	//for _, v := range rtu.RTUs{
	//	log.Println(v)
	//}

	serv := modbusx.NewModbusServer(ctx)	
	defer serv.Close()

	time.Sleep(4 * time.Second)
	modbusx.WriteRegisters(ctx)
	
	<-quit
	
	log.Println("Shutdown Server ...")
}

