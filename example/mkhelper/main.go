package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/go-vgo/robotgo"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	RuneSelf   = 0x80
	CharsSpace = " \r\n\t"
)

var (
	spaceCharAS asciiSet
)

type asciiSet [8]uint32

func init() {
	var ok bool
	spaceCharAS, ok = makeASCIISet(CharsSpace)
	if !ok {
		panic("error ascii set")
	}
}

func main() {
	fileName := flag.String("file", "", "please point to a code file")
	flag.Parse()
	fd, err := os.Open(*fileName)
	if err != nil {
		panic(err.Error())
	}
	parse(fd)

}

func parse(file io.Reader) ([]*operator, error) {

	var line = 0
	ops := make([]*operator, 0)
	for {
		line++
		// 读取一行
		buf, toolong, err := bufio.NewReader(file).ReadLine()
		if toolong {
			return nil, fmt.Errorf("code line [%d] is too long", line)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		// 去除空格字符
		buf = bytes.TrimLeft(buf, " \r\n\t")
		// 跳过注释
		if len(buf) > 1 && buf[0] == '/' && buf[1] == '/' {
			continue
		}
		ops = append(ops, parseLine(buf))
	}
	return ops, nil
}

func (as *asciiSet) contains(c uint8) bool {
	return (as[c>>5] & (1 << uint(c&31))) != 0
}

func parseLine(buf []byte) (*operator, error) {
	length := len(buf)
	var cache bytes.Buffer
	var args = make([][]byte, 0)
	for i := 0; i < length; i++ {
		if spaceCharAS.contains(buf[i]) {
			if cache.Len() > 0 {
				args = append(args, cache.Bytes())
				cache.Reset()
			}
			continue
		}
		cache.WriteByte(buf[i])
	}
	if cache.Len() > 0 {
		args = append(args, cache.Bytes())
	}
	cache.Reset()
	op := operator{}
	switch len(args) {
	case 3:
		split := strings.Split(string(args[2]), ",")
		arg := make([]interface{}, len(split))
		for i, s := range split {
			arg[i] = s
		}
		op.Args = arg
		fallthrough
	case 2:
		split := strings.Split(string(args[1]), ",")
		if len(split) > 1 {
			arg := make([]interface{}, len(split))
			for i, s := range split {
				arg[i] = s
			}

			op.Target = arg
		} else {
			op.Target = args[1]
		}
		fallthrough
	case 1:
		opt, ok := GetOpByName(string(buf[0]))
		if !ok {
			return nil, fmt.Errorf("unknow operator: %s", buf)
		}
		op.Op = opt
	default:
		return nil, fmt.Errorf("error line: %s", buf)
	}
	return &op, nil
}

func makeASCIISet(chars string) (as asciiSet, ok bool) {
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c >= RuneSelf {
			return as, false
		}
		as[c>>5] |= 1 << uint(c&31)
	}
	return as, true
}

func eval(op *operator) interface{} {
	switch op.Op {
	case MouseClick:
		robotgo.MouseClick(op.Target.(string), false)
	case MouseDbClick:
		robotgo.MouseClick(op.Target.(string), true)
	case MouseMove:
		target, ok := op.Target.([]int)
		if !ok {
			return nil
		}
		robotgo.Move(target[0], target[1])
	case MouseSmooth:
		target, ok := op.Target.([]int)
		if !ok {
			return nil
		}
		robotgo.MoveSmooth(target[0], target[1], op.Args...)
	case MouseDown:
		robotgo.MouseToggle("down", op.Target.(string))
	case MouseUp:
		robotgo.MouseToggle("up", op.Target.(string))
	case MouseScroll:
		if len(op.Args) == 1 {
			robotgo.ScrollMouse(op.Args[0].(int), op.Target.(string))
		}
	case KeyClick:
		robotgo.KeyTap(op.Target.(string), op.Args...)
	case KeyDown:
		robotgo.KeyToggle(op.Target.(string), "down")
	case KeyUp:
		robotgo.KeyToggle(op.Target.(string), "up")
	case InputString:
		robotgo.TypeStr(op.Target.(string))
	case ScreenColor:
		robotgo.GetPixelColor(robotgo.GetMousePos())
	}

	return nil
}

func runCmd(str string) {
	cmd := exec.Command(str)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	fmt.Println(cmd.Args)
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	fmt.Println(out.String())
}
