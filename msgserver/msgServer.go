package msgserver

import (
    "fmt"
    "log"
    "context"

    "github.com/gin-gonic/gin"
)

type MsgServer struct{
    Port int
    chData chan RTU
}

func NewServer(port int, ch chan RTU) *Server {

    s := &MsgServer{Port: port,
                    chData: ch}
    return s    
}

func(s *MsgServer) Run() {

    router := gin.Default()
    router.RedirectFixedPath = true
    router.POST("/rtu", rtuHandler)
    if Port > 0 {
        router.Run(fmt.sprintf(":%d", Port))
    }
}

func(s *MsgServer) rtuHandler(c *gin.Context) {
    cmd := c.GetHeader("cmd")
    system := c.GetHeader("system")
    rtuAddress := c.GetHeader("address")
    data := c.GetHeader("data")
}