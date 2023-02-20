package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func rumCommand() {
	var commandApp = os.Args[1]   //"ping"              //os.Args[1]
	var commandArgs = os.Args[2:] //"www.bilibili.com" //os.Args[1:]
	fmt.Printf("args line: %s  %s\n", commandApp, commandArgs)
	cmd := exec.Command(commandApp, commandArgs...)
	var outB, errB bytes.Buffer
	cmd.Stdout = &outB
	cmd.Stderr = &errB
	err := cmd.Run()
	if err != nil {
		fmt.Println("err:", errB.String())
		log.Fatal(err)
	}
	fmt.Println("out:", outB.String())

}

func main() {
	rumCommand()
}
