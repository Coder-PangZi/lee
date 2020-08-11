package main

import "fmt"

type Device int
type Operator int
type Target interface{}

const (
	_ Operator = iota
	mouse_beg
	MouseClick
	MouseDbClick
	MouseMove
	MouseSmooth
	MouseDown
	MouseUp
	MouseScroll
	mouse_end

	key_board_beg
	KeyClick
	KeyDown
	KeyUp
	key_board_end

	input_beg
	InputString
	input_end

	screen_beg
	ScreenColor
	screen_end
)

type OPMap map[string]Operator
type OPMap2 map[Operator]string

var mOPMapNames = OPMap{
	"Mouse.Click":   MouseClick,
	"Mouse.DbClick": MouseDbClick,
	"Mouse.Move":    MouseMove,
	"Mouse.Smooth":  MouseSmooth,
	"Mouse.Down":    MouseDown,
	"Mouse.Up":      MouseUp,
	"Mouse.Scroll":  MouseScroll,
	"Key.Click":     KeyClick,
	"Key.Down":      KeyDown,
	"Key.Up":        KeyUp,
	"Key.String":    InputString,
	"Screen.Color":  ScreenColor,
}
var mOPMapEvent = OPMap2{
	MouseClick:   "Mouse.Click",
	MouseDbClick: "Mouse.DbClick",
	MouseMove:    "Mouse.Move",
	MouseSmooth:  "Mouse.Smooth",
	MouseDown:    "Mouse.Down",
	MouseUp:      "Mouse.Up",
	MouseScroll:  "Mouse.Scroll",
	KeyClick:     "Key.Click",
	KeyDown:      "Key.Down",
	KeyUp:        "Key.Up",
	InputString:  "Key.String",
	ScreenColor:  "Screen.Color",
}

type operator struct {
	Op     Operator
	Target Target
	Args   []interface{}
}

func newOperator(op Operator, target Target, args []interface{}) *operator {
	if args == nil {
		args = []interface{}{}
	}
	return &operator{Op: op, Target: target, Args: args}
}

func GetOpByName(name string) (Operator, bool) {
	op, ok := mOPMapNames[name]
	return op, ok
}

func (o *operator) String() string {
	return fmt.Sprintf("[%s]\t%s:\t%v", mOPMapEvent[o.Op], o.Target, o.Args)
}
