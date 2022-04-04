package rtu
 
import (
	"bufio"
	//"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type RTU struct {
	cmd		string
	system  int 
	address string
	loop    int 
	value   int
}

type RTUS struct{
	RTUs map[int]RTU
}


func(rtus *RTUS) GetRTU() {

	rtus.RTUs = make(map[int]RTU)
	file, err := os.Open("./rtu.db")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()
 	
 	for _, eachline := range txtlines {
 		str := strings.Split(eachline, ":")
 		cmd := str[0]
 		system, _ := strconv.Atoi(str[1])
 		addr := str[2]
 		l, err := strconv.Atoi(str[3])
 		v, err1 := strconv.Atoi(str[4])
 		if len(addr) > 0 && err ==nil && err1 == nil{
 			rtu := RTU{
 				cmd: cmd,
 				system: system,
 				address: addr,
 				loop: l,
 				value: v}	
 			rtus.RTUs[9] = rtu
 			
 		} 
	}
	//for _, v := range RTUs{
	//	fmt.Println(v)
	//}
}