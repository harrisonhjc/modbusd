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
 
 	"os"
    "os/signal"
    "syscall"
 	
)
 
var SaveValue map[int]int
 
 
func main() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	chData := make(chan rtu.RTU)
	defer close(chData)
	//go getMessage(ctx)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//rtu.GetRTU()

	msgserver.NewServer(ctx, 8081, chData)
	//for _, v := range rtu.RTUs{
	//	log.Println(v)
	//}

	serv := modbusx.NewModbusServer(ctx)	
	defer serv.Close()

	time.Sleep(4 * time.Second)
	data := []byte{0,3,0,4,0,5,0,6}
	modbusx.WriteRegisters(ctx, 0, 4, data)
	
	<-quit
	
	log.Println("Shutdown Server ...")
}

