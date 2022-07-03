package msgserver

import (
	"context"
	"fmt"
	"log"
	"modbusd/mbserver"
	"modbusd/rtu"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type MsgServer struct {
	Port   int
	chData chan rtu.RTU
	Rtus   rtu.RTUS
	Ctx    context.Context
	Mutex  sync.Mutex
}

func NewServer(ctx context.Context, port int, ch chan rtu.RTU) *MsgServer {

	s := &MsgServer{Port: port,
		chData: ch}
	s.Ctx = ctx
	return s
}

func (s *MsgServer) Run() {

	router := gin.Default()
	router.RedirectFixedPath = true
	router.POST("/rtu", s.rtuHandler)
	if s.Port > 0 {
		router.Run(fmt.Sprintf(":%d", s.Port))
	}
}

func (s *MsgServer) rtuHandler(c *gin.Context) {

	cmd := c.GetHeader("cmd")
	sys, _ := strconv.Atoi(c.GetHeader("system"))
	a, _ := strconv.Atoi(c.GetHeader("address"))
	l, _ := strconv.Atoi(c.GetHeader("loop"))
	v, _ := strconv.Atoi(c.GetHeader("value"))
	co, _ := strconv.Atoi(c.GetHeader("code"))
	r := rtu.RTU{
		Cmd:     cmd,
		System:  sys,
		Address: a,
		Loop:    l,
		Value:   v,
		Code:    co,
	}

	log.Println(r)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	Cmd := r.Cmd
	system := r.System
	address := r.Address
	loop := r.Loop
	modBusAddress := 0
	value := strings.Trim(fmt.Sprintf("% 05b", r.Value), " ")
	code := r.Code
	remainder := (address) % 4
	switch Cmd {

	case "J0", "J1":
		log.Println("WriteRegisters ")
		modBusAddress = (address)/4 + 1 + (system-1)*64

		s.Mutex.Lock()
		var getRegistersData []byte = mbserver.ReadRegisters(ctx, uint16(modBusAddress)-1, 1)
		var s1, s2 string

		s1 = strings.Trim(fmt.Sprintf("% 09b", getRegistersData[0]), " ")
		s2 = strings.Trim(fmt.Sprintf("% 09b", getRegistersData[1]), " ")
		fmt.Println(s1 + s2)
		fmt.Println(value)

		switch remainder {
		case 0:
			s2 = s2[0:4] + value
		case 1:
			s2 = value + s2[4:8]
		case 2:
			s1 = s1[0:4] + value
		case 3:
			s1 = value + s1[4:8]
		}
		i1, err1 := strconv.ParseInt(s1, 2, 64)
		i2, err2 := strconv.ParseInt(s2, 2, 64)
		if err1 != nil {
			panic(err1)
		}
		if err2 != nil {
			panic(err2)
		}
		fmt.Println(i1, i2)

		writeRegistersData := []byte{uint8(i1), uint8(i2)}

		mbserver.WriteRegisters(ctx, uint16(modBusAddress)-1, 1, writeRegistersData)

		s.Mutex.Unlock()

	case "C0":
		modBusAddress = 1024 + address + 1 + (system-1)*256

		s.Mutex.Lock()
		var getRegistersData []byte = mbserver.ReadRegisters(ctx, uint16(modBusAddress)-1, 1)
		var s1, s2 string

		s1 = strings.Trim(fmt.Sprintf("% 09b", getRegistersData[0]), " ")
		s2 = strings.Trim(fmt.Sprintf("% 09b", getRegistersData[1]), " ")
		value := strings.Trim(fmt.Sprintf("% 02b", r.Value), " ")

		fmt.Println(s1 + s2)
		fmt.Println(value)

		switch loop {
		case 1:
			switch code {
			case 0:
				s2 = s2[0:7] + value
			case 1:
				s2 = s2[0:6] + value + s2[7:8]
			case 2:
				s2 = s2[0:5] + value + s2[6:8]
			case 3:
				s2 = s2[0:4] + value + s2[5:8]
			}

		case 2:
			switch code {
			case 0:
				s2 = s2[0:3] + value + s2[4:8]
			case 1:
				s2 = s2[0:2] + value + s2[3:8]
			case 2:
				s2 = s2[0:1] + value + s2[2:8]
			case 3:
				s2 = value + s2[1:8]
			}
		case 3:
			switch code {
			case 0:
				s1 = s1[0:7] + value
			case 1:
				s1 = s1[0:6] + value + s1[7:8]
			case 2:
				s1 = s1[0:5] + value + s1[6:8]
			case 3:
				s1 = s1[0:4] + value + s1[5:8]
			}
		case 4:
			switch code {
			case 0:
				s1 = s1[0:3] + value + s1[4:8]
			case 1:
				s1 = s1[0:2] + value + s1[3:8]
			case 2:
				s1 = s1[0:1] + value + s1[2:8]
			case 3:
				s1 = value + s1[1:8]
			}
		}
		i1, err1 := strconv.ParseInt(s1, 2, 64)
		i2, err2 := strconv.ParseInt(s2, 2, 64)
		if err1 != nil {
			panic(err1)
		}
		if err2 != nil {
			panic(err2)
		}
		fmt.Println(i1, i2)

		writeRegistersData := []byte{uint8(i1), uint8(i2)}

		mbserver.WriteRegisters(ctx, uint16(modBusAddress)-1, 1, writeRegistersData)
		s.Mutex.Unlock()

	case "F0":
		modBusAddress = address*4 + loop + (system-1)*1024

		serv := s.Ctx.Value("modBusServ").(*mbserver.Server)

		s.Mutex.Lock()
		serv.DiscreteInputs[modBusAddress-1] = byte(uint8(r.Value))
		s.Mutex.Unlock()

		mbserver.ReadtheDiscreteInputs(ctx, 0, 2000)

	case "T0":
		modBusAddress = 16384 + address*4 + loop + (system-1)*1024

		serv := s.Ctx.Value("modBusServ").(*mbserver.Server)

		s.Mutex.Lock()
		serv.DiscreteInputs[modBusAddress-1] = byte(uint8(r.Value))
		s.Mutex.Unlock()

		mbserver.ReadtheDiscreteInputs(ctx, 0, 2000)

	case "B0":
		modBusAddress = 32768 + code + (system)*54
		if system != 0 {
			modBusAddress = modBusAddress + (system)*2
		}
		serv := s.Ctx.Value("modBusServ").(*mbserver.Server)

		s.Mutex.Lock()
		serv.DiscreteInputs[modBusAddress-1] = byte(uint8(r.Value))
		s.Mutex.Unlock()

	case "F1":
		modBusAddress = 33728 + system

		serv := s.Ctx.Value("modBusServ").(*mbserver.Server)

		s.Mutex.Lock()
		serv.DiscreteInputs[modBusAddress-1] = byte(uint8(r.Value))
		s.Mutex.Unlock()

	default:
		log.Println("No Cmd Pares ")
		log.Println("ReadRegisters : ")
		mbserver.ReadRegisters(ctx, 0, 20)
		mbserver.ReadtheDiscreteInputs(ctx, 0, 2000)

		return
	}
}
