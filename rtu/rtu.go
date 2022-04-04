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


func(rtu *RTUS) GetRTU() {

	rtu.RTUs = make(map[int]RTU)
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
 		addr := str[1]
 		id, err := strconv.Atoi(str[2])
 		length, err1:= strconv.Atoi(str[3])
 		if len(addr) > 0 && err ==nil && err1 == nil{
 			rtu := RTU{
 				cmd: cmd,
 				address: addr,
 				length: length,
 				value: 0}	
 			rtu.RTUs[id] = rtu
 			
 		} 
	}
	//for _, v := range RTUs{
	//	fmt.Println(v)
	//}
}