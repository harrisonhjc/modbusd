package msgserver

import (
	"context"
	"fmt"
	"log"
	"modbusd/mbserver"
	"modbusd/rtu"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MsgServer struct {
	Port   int
	chData chan rtu.RTU
	Rtus   rtu.RTUS
}

func NewServer(ctx context.Context, port int, ch chan rtu.RTU) *MsgServer {

	s := &MsgServer{Port: port,
		chData: ch}
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

	system := r.System
	address := r.Address
	modBusAddress := 0
	value := strings.Trim(fmt.Sprintf("% 05b", r.Value), " ")
	remainder := (address) % 4
	switch r.Cmd {

	case "J0", "J1":
		log.Println("WriteRegisters ")
		modBusAddress = (address)/4 + 1 + (system-1)*64

	case "C0":
		modBusAddress = 1024 + address + 1 + (system-1)*256
	default:
		log.Println("No Cmd Pares ")
		log.Println("ReadRegisters : ")
		mbserver.ReadRegisters(ctx, 0, 20)
		return
	}

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
}

// func Int64ToBytes(i int64) []byte {
// 	var buf = make([]byte, 8)
// 	binary.BigEndian.PutUint64(buf, uint64(i))
// 	return buf
// }
