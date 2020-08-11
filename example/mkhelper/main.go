package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/go-vgo/robotgo"
	"io"
	"log"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
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

const sc = `
Mouse.Move	100,500
Mouse.Click  left
Mouse.Move	200,500
Mouse.Down left
Mouse.Move	500,500
Mouse.Up left
Mouse.DbClick left 600,500 
Mouse.Move	700,500
Mouse.Down left
Mouse.Smooth	1000,500
Mouse.Scroll	down 5
Mouse.Up left
//
Mouse.Move 3000,500
Mouse.Click  left
Key.Click a
Key.Down w
Key.Up w
Key.String hello
Screen.Color
`

func main() {
	//fileName := flag.String("f", "", "please point to a code file")
	//flag.Parse()
	//fd, err := os.Open(*fileName)
	//if err != nil {
	//	panic(err.Error())
	//}
	fd := strings.NewReader(sc)
	ops, err := parse(fd)
	if err != nil {
		panic(err)
	}
	for _, op := range ops {
		err := eval(op)
		if err != nil {
			panic(err)
		}
	}

}

func parse(file io.Reader) ([]*operator, error) {

	var line = 0
	ops := make([]*operator, 0)
	reader := bufio.NewReader(file)
	for {
		line++
		// 读取一行
		buf, toolong, err := reader.ReadLine()
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
		if len(buf) == 0 || (len(buf) > 1 && buf[0] == '/' && buf[1] == '/') {
			continue
		}
		op, err := parseLine(buf)
		if err != nil {
			return nil, err
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func (as *asciiSet) contains(c uint8) bool {
	return (as[c>>5] & (1 << uint(c&31))) != 0
}

func parseLine(buf []byte) (*operator, error) {
	length := len(buf)
	var cache bytes.Buffer
	var args = make([][]byte, 0, 3)
	for i := 0; i < length; i++ {
		if spaceCharAS.contains(buf[i]) {
			if cache.Len() > 0 {
				args = append(args, copyByteSlice(cache.Bytes()))
				cache.Reset()
			}
			continue
		}
		cache.WriteByte(buf[i])
	}
	if cache.Len() > 0 {
		args = append(args, copyByteSlice(cache.Bytes()))
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
		opt, ok := GetOpByName(string(args[0]))
		if !ok {
			return nil, fmt.Errorf("unknow operator: %s", buf)
		}
		op.Op = opt
	default:
		return nil, fmt.Errorf("error line: %s", buf)
	}
	return &op, nil
}

func copyByteSlice(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
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

func eval(op *operator) error {
	println(op.String())
	switch op.Op {
	case MouseClick:
		robotgo.MouseClick(byte2String(op.Target.([]byte)), false)
	case MouseDbClick:
		robotgo.MouseClick(byte2String(op.Target.([]byte)), true)
	case MouseMove:
		target, ok := op.Target.([]interface{})
		if !ok {
			return fmt.Errorf("wrong target of MouseMove event")
		}
		pos, err := interfaceSlice2IntSlice(target)
		if err != nil {
			return err
		}
		robotgo.Move(pos[0], pos[1])
	case MouseSmooth:
		target, ok := op.Target.([]interface{})
		if !ok {
			return fmt.Errorf("wrong target of MouseMove event")
		}
		pos, err := interfaceSlice2IntSlice(target)
		if err != nil {
			return err
		}
		robotgo.MoveSmooth(pos[0], pos[1], op.Args...)
	case MouseDown:
		robotgo.MouseToggle("down", byte2String(op.Target.([]byte)))
	case MouseUp:
		robotgo.MouseToggle("up", byte2String(op.Target.([]byte)))
	case MouseScroll:
		if len(op.Args) == 1 {
			d, err := strconv.Atoi(op.Args[0].(string))
			if err != nil {
				return err
			}
			robotgo.ScrollMouse(d, byte2String(op.Target.([]byte)))
		}
	case KeyClick:
		robotgo.KeyTap(byte2String(op.Target.([]byte)), op.Args...)
	case KeyDown:
		robotgo.KeyToggle(byte2String(op.Target.([]byte)), "down")
	case KeyUp:
		robotgo.KeyToggle(byte2String(op.Target.([]byte)), "up")
	case InputString:
		robotgo.TypeStr(byte2String(op.Target.([]byte)))
	case ScreenColor:
		robotgo.GetPixelColor(robotgo.GetMousePos())
	}

	return nil
}

func byte2String(src []byte) (dst string) {
	pBytes := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	pString := (*reflect.StringHeader)(unsafe.Pointer(&dst))
	pString.Data = pBytes.Data
	pString.Len = pBytes.Len
	return
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

func interfaceSlice2IntSlice(src []interface{}) ([]int, error) {
	var err error
	var dstInt = make([]int, len(src), len(src))
	for i, arg := range src {
		argStr, ok := arg.(string)
		if !ok {
			return nil, fmt.Errorf("wrong type:[%T] %v", arg, arg)
		}
		dstInt[i], err = strconv.Atoi(argStr)
		if err != nil {
			return nil, err
		}

	}
	return dstInt, nil
}

//func getPosFromStrSliceX(src []interface{}) (x, y int, err error) {
//	pString := (*reflect.StringHeader)(unsafe.Pointer(&src))
//	return interfaceSlice2IntSlice(pString)
//}
