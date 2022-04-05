package msgserver

import (
    "fmt"
    "log"
    "context"
    "strconv"
    "modbusd/rtu"
    "github.com/gin-gonic/gin"
)

type MsgServer struct{
    Port int
    chData chan rtu.RTU
    Rtus rtu.RTUS
}

func NewServer(ctx context.Context, port int, ch chan rtu.RTU) *MsgServer {

    s := &MsgServer{Port: port,
                    chData: ch}
    return s    
}

func(s *MsgServer) Run() {

 
    router := gin.Default()
    router.RedirectFixedPath = true
    router.POST("/rtu", s.rtuHandler)
    if s.Port > 0 {
        router.Run(fmt.Sprintf(":%d", s.Port))
    }
}

func(s *MsgServer) rtuHandler(c *gin.Context) {
    cmd := c.GetHeader("cmd")
    sys := c.GetHeader("system")
    a := c.GetHeader("address")
    l := c.GetHeader("loop")
    v := c.GetHeader("value")
    co := c.GetHeader("code")
    r := rtu.RTU {
        Cmd: cmd,
        System: strconv.Atoi(sys),
        Address: a,
        Loop: l,
        Value: v, 
        Code: co,
    }
    
    log.Println(cmd, system, rtuAddress, data)
}

