package main

import (
	//"flag"
	//"fmt"
	"context"
	"log"
	"modbusd/mbserver"
	"modbusd/msgserver"
	"modbusd/rtu"

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

	serv := mbserver.NewModbusServer(ctx)
	defer serv.Close()
	log.Println("ModBusd Server Start...")

	c2 := context.WithValue(ctx, "modBusServ", serv)

	msgServer := msgserver.NewServer(c2, 8081, chData)
	go msgServer.Run()
	log.Println("HTTP Server Start...")
	//for _, v := range rtu.RTUs{
	//	log.Println(v)
	//}

	// time.Sleep(2 * time.Second)
	// data := []byte{0, 3, 0, 4, 0, 5, 0, 6}
	// mbserver.WriteRegisters(ctx, 0, 4, data)

	<-quit
	log.Println("Shutdown Server ...")
}
