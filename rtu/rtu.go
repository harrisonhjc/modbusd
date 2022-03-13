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
	address string
	length	int
	value   int
}

var RTUs map[int]RTU

func GetRTU() {

	RTUs = make(map[int]RTU)
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
 		var addr string
 		addr = str[0]
 		id, err := strconv.Atoi(str[1])
 		length, err1:= strconv.Atoi(str[2])
 		if len(addr) > 0 && err ==nil && err1 == nil{
 			rtu := RTU{
 				address: addr,
 				length: length,
 				value: 0}	
 			RTUs[id] = rtu
 			
 		} 
	}
	//for _, v := range RTUs{
	//	fmt.Println(v)
	//}
}