package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	command := os.Args[1]
	log.Output(0, " ------------------\n execute applicataion: "+command)

	fmt.Println("command value!!", command)

	log.Output(0, " ------------------\n fishish application: "+command)
}
