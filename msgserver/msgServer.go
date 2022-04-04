package msgserver

import (
    "fmt"
    "log"
    "context"
    "modbusd/rtu"
    "github.com/gin-gonic/gin"
)

type MsgServer struct{
    Port int
    chData chan rtu.RTU
}

func NewServer(ctx context.Context, port int, ch chan rtu.RTU) *MsgServer {

    s := &MsgServer{Port: port,
                    chData: ch}
    return s    
}

func(s *MsgServer) Run() {

    router := gin.Default()
    router.RedirectFixedPath = true
    router.POST("/rtu", rtuHandler)
    if s.Port > 0 {
        router.Run(fmt.Sprintf(":%d", s.Port))
    }
}

func rtuHandler(c *gin.Context) {
    cmd := c.GetHeader("cmd")
    system := c.GetHeader("system")
    rtuAddress := c.GetHeader("address")
    data := c.GetHeader("data")
    log.Println(cmd, system, rtuAddress, data)
}