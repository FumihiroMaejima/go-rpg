package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"os"

	"go-rpg/util"
	"go-rpg/web"
)

func main() {
	checkEnvMessage := util.CheckEnvValue()
	if checkEnvMessage != "" {
		log.Output(0, "execution failed.\n")
		return
	}

	command := os.Args[1]
	log.Output(0, " ------------------\n execute applicataion: "+command)

	fmt.Println("command value!!", command)
	web.App()

	log.Output(0, " ------------------\n fishish application: "+command)
}
