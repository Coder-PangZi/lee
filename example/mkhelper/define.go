package main

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

var mOPMap = OPMap{
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
	op, ok := mOPMap[name]
	return op, ok
}
