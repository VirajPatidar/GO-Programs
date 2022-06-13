package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func learnGo(learn chan bool) {
	i := 10
	fmt.Print("Learning in...")

	for i > 0 {
		time.Sleep(time.Second)
		fmt.Println(i)
		i--
		clear()
	}
	fmt.Println("...GOing to learn")

	learn <- true
}

func main() {

	learn := make(chan bool, 1)
	go learnGo(learn)

	<-learn

}
