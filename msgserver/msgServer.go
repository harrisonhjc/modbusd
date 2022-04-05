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
    sys, _:= strconv.Atoi(c.GetHeader("system"))
    a, _:= strconv.Atoi(c.GetHeader("address"))
    l, _ := strconv.Atoi(c.GetHeader("loop"))
    v, _ := strconv.Atoi(c.GetHeader("value"))
    co, _ := strconv.Atoi(c.GetHeader("code"))
    r := rtu.RTU {
        Cmd: cmd,
        System: sys,
        Address: a,
        Loop: l,
        Value: v, 
        Code: co,
    }
    
    log.Println(r)
}

