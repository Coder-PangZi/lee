package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-vgo/robotgo"
	"testing"
)

func TestRobotgo(t *testing.T) {
	//cmd := exec.Command("notepad")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Start()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("Waiting for command to finish...")
	//fmt.Println(cmd.Args)
	//go func() {
	//	err = cmd.Wait()
	//	if err != nil {
	//		log.Printf("Command finished with error: %v", err)
	//	}
	//}()
	x, y := robotgo.GetMousePos()
	fmt.Println("mouse pos: ", x, y)

	clo := robotgo.GetPixelColor(x, y)
	fmt.Println("color: #", clo)

	// clipboard
	spew.Dump(robotgo.WriteAll("#" + clo))

}
