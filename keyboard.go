package xzkey

import (
	"fmt"
	"time"

	"github.com/CxZMoE/NetEase-CMD/logger"
	"github.com/MarinX/keylogger"
)

// 映射
const (
	NUMLOCK   = "NUMLOCK"
	Q         = "Q"
	W         = "W"
	E         = "E"
	R         = "R"
	T         = "T"
	Y         = "Y"
	U         = "U"
	I         = "I"
	O         = "O"
	P         = "P"
	A         = "A"
	S         = "S"
	D         = "D"
	F         = "F"
	G         = "G"
	H         = "H"
	J         = "J"
	K         = "K"
	L         = "L"
	Z         = "Z"
	X         = "X"
	C         = "C"
	V         = "V"
	B         = "B"
	N         = "N"
	M         = "M"
	UP        = "UP"
	DOWN      = "DOWN"
	LEFT      = "LEFT"
	RIGHT     = "RIGHT"
	PgUp      = "PgUp"
	PgDn      = "PgDn"
	KEY0      = "KEY0"
	KEY1      = "KEY1"
	KEY2      = "KEY2"
	KEY3      = "KEY3"
	KEY4      = "KEY4"
	KEY5      = "KEY5"
	KEY6      = "KEY6"
	KEY7      = "KEY7"
	KEY8      = "KEY8"
	KEY9      = "KEY9"
	BACKSPACE = "BACKSPACE"
	DELETE    = "DELETE"
	INSERT    = "INSERT"
	SPACE     = "SPACE"
	HOME      = "HOME"
	F1        = "F1"
	F2        = "F2"
	F3        = "F3"
	F4        = "F4"
	F5        = "F5"
	F6        = "F6"
	F7        = "F7"
	F8        = "F8"
	F9        = "F9"
	F10       = "F10"
	F11       = "F11"
	F12       = "F12"
	ESC       = "ESC"
	CapsLock  = "CapsLock"
	LSHIFT    = "LSHIFT"
	RSHIFT    = "RSHIFT"
	RCTRL     = "RCTRL"
	LCTRL     = "LCTRL"
	LALT      = "LALT"
	RALT      = "RALT"
	ENTER     = "ENTER"
	TAB       = "TAB"
	KeyOther  = "KeyOther"
	Zkhz      = "["
	Zkhy      = "]"
)

// Key 按键
type Key struct {
	Pressed  bool // 按住
	Released bool // 松开
}

// Keyboard 键盘
type Keyboard struct {
	Keys   map[string]*Key
	Status KeyboardStatus
	Logger *keylogger.KeyLogger
	Event  *chan keylogger.InputEvent
}

// KeyboardStatus 键盘状态信息
type KeyboardStatus struct {
	Device         string // path to dev in /dev/input/eventX
	EventHandlers  map[string]int
	HandleHandlers map[string]func()
}

// KeyBindHandler 键盘事件处理
type KeyBindHandler interface {
	KeyHandler()
}

// NewKeyboard 创建一个新键盘
func NewKeyboard() *Keyboard {
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Printf("\n[ERR] Failed to get keyboard: %v", err)
			logger.WriteLog(fmt.Sprint(err))
		}
	}()
	keyboard := &Keyboard{
		Status: KeyboardStatus{},
		Logger: nil,
		Event:  nil,
	}
	keyboard.Status.EventHandlers = make(map[string]int, 100)
	keyboard.Keys = keyMap
	keyboard.GetKeyboardDevice()
	if keyboard.Init() == nil {
		return nil
	}
	keyboard.StartReadEvent()
	go keyboard.MainLoop()
	return keyboard
}

// GetKeyboardDevice 创建新的键盘设备
func (k *Keyboard) GetKeyboardDevice() string {
	device := keylogger.FindKeyboardDevice()
	if len(device) <= 0 {
		fmt.Printf("\n[INFO] Failed to find keyboard device.")
		return ""
	}
	k.Status.Device = device
	return device
}

// Init 初始化键盘事件记录对象
func (k *Keyboard) Init() *keylogger.KeyLogger {
	kl, err := keylogger.New(k.Status.Device)
	if err != nil {
		//fmt.Printf("\n[INFO] Failed to create key logger.")
		//return nil
	}
	if !kl.IsRoot() {
		fmt.Printf("\n[INFO] 没有权限,已关闭全局热键.")
		return nil
	}
	k.Logger = kl
	return kl
}

// StartReadEvent 开始读取键盘事件
func (k *Keyboard) StartReadEvent() *chan keylogger.InputEvent {
	e := k.Logger.Read()
	k.Event = &e
	return &e
}

// StopReadEvent 停止读取键盘事件
func (k *Keyboard) StopReadEvent() error {
	if k.Logger.IsRoot() {
		err := k.Logger.Close()
		return err
	}
	return nil
}

// MainLoop 获取按键状态主循环
func (k *Keyboard) MainLoop() {
	for e := range *k.Event {
		switch e.Type {
		case keylogger.EvKey:
			//var justReleased = ""
			if e.KeyPress() {
				switch e.KeyString() {
				case "ESC":
					k.Keys["ESC"].Pressed = true
					k.Keys["ESC"].Released = false
					break
				case "1":
					k.Keys["KEY1"].Pressed = true
					k.Keys["KEY1"].Released = false
					break
				case "2":
					k.Keys["KEY2"].Pressed = true
					k.Keys["KEY2"].Released = false
					break
				case "3":
					k.Keys["KEY3"].Pressed = true
					k.Keys["KEY3"].Released = false
					break
				case "4":
					k.Keys["KEY4"].Pressed = true
					k.Keys["KEY4"].Released = false
					break
				case "5":
					k.Keys["KEY5"].Pressed = true
					k.Keys["KEY5"].Released = false
					break
				case "6":
					k.Keys["KEY6"].Pressed = true
					k.Keys["KEY6"].Released = false
					break
				case "7":
					k.Keys["KEY7"].Pressed = true
					k.Keys["KEY7"].Released = false
					break
				case "8":
					k.Keys["KEY8"].Pressed = true
					k.Keys["KEY8"].Released = false
					break
				case "9":
					k.Keys["KEY9"].Pressed = true
					k.Keys["KEY9"].Released = false
					break
				case "0":
					k.Keys["KEY0"].Pressed = true
					k.Keys["KEY0"].Released = false
					break
				case "-":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "=":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "BS":
					k.Keys["BACKSPACE"].Pressed = true
					k.Keys["BACKSPACE"].Released = false
					break
				case "TAB":
					k.Keys["TAB"].Pressed = true
					k.Keys["TAB"].Released = false
					break
				case "Q":
					k.Keys["Q"].Pressed = true
					k.Keys["Q"].Released = false
					break
				case "W":
					k.Keys["W"].Pressed = true
					k.Keys["W"].Released = false
					break
				case "E":
					k.Keys["E"].Pressed = true
					k.Keys["E"].Released = false
					break
				case "R":
					k.Keys["R"].Pressed = true
					k.Keys["R"].Released = false
					break
				case "T":
					k.Keys["T"].Pressed = true
					k.Keys["T"].Released = false
					break
				case "Y":
					k.Keys["Y"].Pressed = true
					k.Keys["Y"].Released = false
					break
				case "U":
					k.Keys["U"].Pressed = true
					k.Keys["U"].Released = false
					break
				case "I":
					k.Keys["I"].Pressed = true
					k.Keys["I"].Released = false
					break
				case "O":
					k.Keys["O"].Pressed = true
					k.Keys["O"].Released = false
					break
				case "P":
					k.Keys["P"].Pressed = true
					k.Keys["P"].Released = false
					break
				case "[":
					k.Keys["["].Pressed = true
					k.Keys["["].Released = false
					break
				case "]":
					k.Keys["]"].Pressed = true
					k.Keys["]"].Released = false
					break
				case "ENTER":
					k.Keys["ENTER"].Pressed = true
					k.Keys["ENTER"].Released = false
					break
				case "L_CTRL":
					k.Keys["LCTRL"].Pressed = true
					k.Keys["LCTRL"].Released = false
					break
				case "A":
					k.Keys["A"].Pressed = true
					k.Keys["A"].Released = false
					break
				case "S":
					k.Keys["S"].Pressed = true
					k.Keys["S"].Released = false
					break
				case "D":
					k.Keys["D"].Pressed = true
					k.Keys["D"].Released = false
					break
				case "F":
					k.Keys["F"].Pressed = true
					k.Keys["F"].Released = false
					break
				case "G":
					k.Keys["G"].Pressed = true
					k.Keys["G"].Released = false
					break
				case "H":
					k.Keys["H"].Pressed = true
					k.Keys["H"].Released = false
					break
				case "J":
					k.Keys["J"].Pressed = true
					k.Keys["J"].Released = false
					break
				case "K":
					k.Keys["K"].Pressed = true
					k.Keys["K"].Released = false
					break
				case "L":
					k.Keys["L"].Pressed = true
					k.Keys["L"].Released = false
					break
				case ";":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "'":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "`":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "L_SHIFT":
					k.Keys["LSHIFT"].Pressed = true
					k.Keys["LSHIFT"].Released = false
					break
				case "\\":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "Z":
					k.Keys["Z"].Pressed = true
					k.Keys["Z"].Released = false
					break
				case "X":
					k.Keys["X"].Pressed = true
					k.Keys["X"].Released = false
					break
				case "C":
					k.Keys["C"].Pressed = true
					k.Keys["C"].Released = false
					break
				case "V":
					k.Keys["V"].Pressed = true
					k.Keys["V"].Released = false
					break
				case "B":
					k.Keys["B"].Pressed = true
					k.Keys["B"].Released = false
					break
				case "N":
					k.Keys["N"].Pressed = true
					k.Keys["N"].Released = false
					break
				case "M":
					k.Keys["M"].Pressed = true
					k.Keys["M"].Released = false
					break
				case ",":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case ".":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "/":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "R_SHIFT":
					k.Keys["RSHIFT"].Pressed = true
					k.Keys["RSHIFT"].Released = false
					break
				case "*":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "L_ALT":
					k.Keys["LALT"].Pressed = true
					k.Keys["LALT"].Released = false
					break
				case "SPACE":
					k.Keys["SPACE"].Pressed = true
					k.Keys["SPACE"].Released = false
					break
				case "CAPS_LOCK":
					k.Keys["CapsLock"].Pressed = true
					k.Keys["CapsLock"].Released = false
					break
				case "F1":
					k.Keys["F1"].Pressed = true
					k.Keys["F1"].Released = false
					break
				case "F2":
					k.Keys["F2"].Pressed = true
					k.Keys["F2"].Released = false
					break
				case "F3":
					k.Keys["F3"].Pressed = true
					k.Keys["F3"].Released = false
					break
				case "F4":
					k.Keys["F4"].Pressed = true
					k.Keys["F4"].Released = false
					break
				case "F5":
					k.Keys["F5"].Pressed = true
					k.Keys["F5"].Released = false
					break
				case "F6":
					k.Keys["F6"].Pressed = true
					k.Keys["F6"].Released = false
					break
				case "F7":
					k.Keys["F7"].Pressed = true
					k.Keys["F7"].Released = false
					break
				case "F8":
					k.Keys["F8"].Pressed = true
					k.Keys["F8"].Released = false
					break
				case "F9":
					k.Keys["F9"].Pressed = true
					k.Keys["F9"].Released = false
					break
				case "F10":
					k.Keys["F10"].Pressed = true
					k.Keys["F10"].Released = false
					break
				case "NUM_LOCK":
					k.Keys["ESC"].Pressed = true
					k.Keys["ESC"].Released = false
					break
				case "SCROLL_LOCK":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "HOME":
					k.Keys["HOME"].Pressed = true
					k.Keys["HOME"].Released = false
					break
				case "UP_8":
					k.Keys["UP"].Pressed = true
					k.Keys["UP"].Released = false
					break
				case "PGUP_9":
					k.Keys["PgUp"].Pressed = true
					k.Keys["PgUp"].Released = false
					break
				case "RT_ARROW_6":
					k.Keys["RIGHT"].Pressed = true
					k.Keys["RIGHT"].Released = false
					break
				case "+":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "END_1":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "DOWN":
					k.Keys["DOWN"].Pressed = true
					k.Keys["DOWN"].Released = false
					break
				case "PGDN_3":
					k.Keys["PgDn"].Pressed = true
					k.Keys["PgDn"].Released = false
					break
				case "INS":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "DEL":
					k.Keys["DELETE"].Pressed = true
					k.Keys["DELETE"].Released = false
					break
				case "F11":
					k.Keys["F11"].Pressed = true
					k.Keys["F11"].Released = false
					break
				case "F12":
					k.Keys["F12"].Pressed = true
					k.Keys["F12"].Released = false
					break
				case "R_ENTER":
					k.Keys["ENTER"].Pressed = true
					k.Keys["ENTER"].Released = false
					break
				case "R_CTRL":
					k.Keys["RCTRL"].Pressed = true
					k.Keys["RCTRL"].Released = false
					break
				case "PRT_SCR":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "R_ALT":
					k.Keys["RALT"].Pressed = true
					k.Keys["RALT"].Released = false
					break
				case "Home":
					k.Keys["HOME"].Pressed = true
					k.Keys["HOME"].Released = false
					break
				case "Up":
					k.Keys["UP"].Pressed = true
					k.Keys["UP"].Released = false
					break
				case "PgUp":
					k.Keys["PgUp"].Pressed = true
					k.Keys["PgUp"].Released = false
					break
				case "Left":
					k.Keys["LEFT"].Pressed = true
					k.Keys["LEFT"].Released = false
					break
				case "Right":
					k.Keys["RIGHT"].Pressed = true
					k.Keys["RIGHT"].Released = false
					break
				case "End":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				case "Down":
					k.Keys["DOWN"].Pressed = true
					k.Keys["DOWN"].Released = false
					break
				case "PgDn":
					k.Keys["PgDn"].Pressed = true
					k.Keys["PgDn"].Released = false
					break
				case "Insert":
					k.Keys["INSERT"].Pressed = true
					k.Keys["INSERT"].Released = false
					break
				case "Del":
					k.Keys["DELETE"].Pressed = true
					k.Keys["DELETE"].Released = false
					break
				case "Pause":
					k.Keys["KeyOther"].Pressed = true
					k.Keys["KeyOther"].Released = false
					break
				default:
					//fmt.Printf("\n[INFO] Unknown key.")
					break
				}
			}
			if e.KeyRelease() {
				switch e.KeyString() {
				case "ESC":
					k.Keys["ESC"].Pressed = false
					k.Keys["ESC"].Released = true
					break
				case "1":
					k.Keys["KEY1"].Pressed = false
					k.Keys["KEY1"].Released = true
					break
				case "2":
					k.Keys["KEY2"].Pressed = false
					k.Keys["KEY2"].Released = true
					break
				case "3":
					k.Keys["KEY3"].Pressed = false
					k.Keys["KEY3"].Released = true
					break
				case "4":
					k.Keys["KEY4"].Pressed = false
					k.Keys["KEY4"].Released = true
					break
				case "5":
					k.Keys["KEY5"].Pressed = false
					k.Keys["KEY5"].Released = true
					break
				case "6":
					k.Keys["KEY6"].Pressed = false
					k.Keys["KEY6"].Released = true
					break
				case "7":
					k.Keys["KEY7"].Pressed = false
					k.Keys["KEY7"].Released = true
					break
				case "8":
					k.Keys["KEY8"].Pressed = false
					k.Keys["KEY8"].Released = true
					break
				case "9":
					k.Keys["KEY9"].Pressed = false
					k.Keys["KEY9"].Released = true
					break
				case "0":
					k.Keys["KEY0"].Pressed = false
					k.Keys["KEY0"].Released = true
					break
				case "-":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "=":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "BS":
					k.Keys["BACKSPACE"].Pressed = false
					k.Keys["BACKSPACE"].Released = true
					break
				case "TAB":
					k.Keys["TAB"].Pressed = false
					k.Keys["TAB"].Released = true
					break
				case "Q":
					k.Keys["Q"].Pressed = false
					k.Keys["Q"].Released = true
					break
				case "W":
					k.Keys["W"].Pressed = false
					k.Keys["W"].Released = true
					break
				case "E":
					k.Keys["E"].Pressed = false
					k.Keys["E"].Released = true
					break
				case "R":
					k.Keys["R"].Pressed = false
					k.Keys["R"].Released = true
					break
				case "T":
					k.Keys["T"].Pressed = false
					k.Keys["T"].Released = true
					break
				case "Y":
					k.Keys["Y"].Pressed = false
					k.Keys["Y"].Released = true
					break
				case "U":
					k.Keys["U"].Pressed = false
					k.Keys["U"].Released = true
					break
				case "I":
					k.Keys["I"].Pressed = false
					k.Keys["I"].Released = true
					break
				case "O":
					k.Keys["O"].Pressed = false
					k.Keys["O"].Released = true
					break
				case "P":
					k.Keys["P"].Pressed = false
					k.Keys["P"].Released = true
					break
				case "[":
					k.Keys["["].Pressed = false
					k.Keys["["].Released = true
					break
				case "]":
					k.Keys["]"].Pressed = false
					k.Keys["]"].Released = true
					break
				case "ENTER":
					k.Keys["ENTER"].Pressed = false
					k.Keys["ENTER"].Released = true
					break
				case "L_CTRL":
					k.Keys["LCTRL"].Pressed = false
					k.Keys["LCTRL"].Released = true
					break
				case "A":
					k.Keys["A"].Pressed = false
					k.Keys["A"].Released = true
					break
				case "S":
					k.Keys["S"].Pressed = false
					k.Keys["S"].Released = true
					break
				case "D":
					k.Keys["D"].Pressed = false
					k.Keys["D"].Released = true
					break
				case "F":
					k.Keys["F"].Pressed = false
					k.Keys["F"].Released = true
					break
				case "G":
					k.Keys["G"].Pressed = false
					k.Keys["G"].Released = true
					break
				case "H":
					k.Keys["H"].Pressed = false
					k.Keys["H"].Released = true
					break
				case "J":
					k.Keys["J"].Pressed = false
					k.Keys["J"].Released = true
					break
				case "K":
					k.Keys["K"].Pressed = false
					k.Keys["K"].Released = true
					break
				case "L":
					k.Keys["L"].Pressed = false
					k.Keys["L"].Released = true
					break
				case ";":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "'":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "`":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "L_SHIFT":
					k.Keys["LSHIFT"].Pressed = false
					k.Keys["LSHIFT"].Released = true
					break
				case "\\":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "Z":
					k.Keys["Z"].Pressed = false
					k.Keys["Z"].Released = true
					break
				case "X":
					k.Keys["X"].Pressed = false
					k.Keys["X"].Released = true
					break
				case "C":
					k.Keys["C"].Pressed = false
					k.Keys["C"].Released = true
					break
				case "V":
					k.Keys["V"].Pressed = false
					k.Keys["V"].Released = true
					break
				case "B":
					k.Keys["B"].Pressed = false
					k.Keys["B"].Released = true
					break
				case "N":
					k.Keys["N"].Pressed = false
					k.Keys["N"].Released = true
					break
				case "M":
					k.Keys["M"].Pressed = false
					k.Keys["M"].Released = true
					break
				case ",":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case ".":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "/":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "R_SHIFT":
					k.Keys["RSHIFT"].Pressed = false
					k.Keys["RSHIFT"].Released = true
					break
				case "*":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "L_ALT":
					k.Keys["LALT"].Pressed = false
					k.Keys["LALT"].Released = true
					break
				case "SPACE":
					k.Keys["SPACE"].Pressed = false
					k.Keys["SPACE"].Released = true
					break
				case "CAPS_LOCK":
					k.Keys["CapsLock"].Pressed = false
					k.Keys["CapsLock"].Released = true
					break
				case "F1":
					k.Keys["F1"].Pressed = false
					k.Keys["F1"].Released = true
					break
				case "F2":
					k.Keys["F2"].Pressed = false
					k.Keys["F2"].Released = true
					break
				case "F3":
					k.Keys["F3"].Pressed = false
					k.Keys["F3"].Released = true
					break
				case "F4":
					k.Keys["F4"].Pressed = false
					k.Keys["F4"].Released = true
					break
				case "F5":
					k.Keys["F5"].Pressed = false
					k.Keys["F5"].Released = true
					break
				case "F6":
					k.Keys["F6"].Pressed = false
					k.Keys["F6"].Released = true
					break
				case "F7":
					k.Keys["F7"].Pressed = false
					k.Keys["F7"].Released = true
					break
				case "F8":
					k.Keys["F8"].Pressed = false
					k.Keys["F8"].Released = true
					break
				case "F9":
					k.Keys["F9"].Pressed = false
					k.Keys["F9"].Released = true
					break
				case "F10":
					k.Keys["F10"].Pressed = false
					k.Keys["F10"].Released = true
					break
				case "NUM_LOCK":
					k.Keys["ESC"].Pressed = false
					k.Keys["ESC"].Released = true
					break
				case "SCROLL_LOCK":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "HOME":
					k.Keys["HOME"].Pressed = false
					k.Keys["HOME"].Released = true
					break
				case "UP_8":
					k.Keys["UP"].Pressed = false
					k.Keys["UP"].Released = true
					break
				case "PGUP_9":
					k.Keys["PgUp"].Pressed = false
					k.Keys["PgUp"].Released = true
					break
				case "RT_ARROW_6":
					k.Keys["RIGHT"].Pressed = false
					k.Keys["RIGHT"].Released = true
					break
				case "+":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "END_1":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "DOWN":
					k.Keys["DOWN"].Pressed = false
					k.Keys["DOWN"].Released = true
					break
				case "PGDN_3":
					k.Keys["PgDn"].Pressed = false
					k.Keys["PgDn"].Released = true
					break
				case "INS":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "DEL":
					k.Keys["DELETE"].Pressed = false
					k.Keys["DELETE"].Released = true
					break
				case "F11":
					k.Keys["F11"].Pressed = false
					k.Keys["F11"].Released = true
					break
				case "F12":
					k.Keys["F12"].Pressed = false
					k.Keys["F12"].Released = true
					break
				case "R_ENTER":
					k.Keys["ENTER"].Pressed = false
					k.Keys["ENTER"].Released = true
					break
				case "R_CTRL":
					k.Keys["RCTRL"].Pressed = false
					k.Keys["RCTRL"].Released = true
					break
				case "PRT_SCR":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "R_ALT":
					k.Keys["RALT"].Pressed = false
					k.Keys["RALT"].Released = true
					break
				case "Home":
					k.Keys["HOME"].Pressed = false
					k.Keys["HOME"].Released = true
					break
				case "Up":
					k.Keys["UP"].Pressed = false
					k.Keys["UP"].Released = true
					break
				case "PgUp":
					k.Keys["PgUp"].Pressed = false
					k.Keys["PgUp"].Released = true
					break
				case "Left":
					k.Keys["LEFT"].Pressed = false
					k.Keys["LEFT"].Released = true
					break
				case "Right":
					k.Keys["RIGHT"].Pressed = false
					k.Keys["RIGHT"].Released = true
					break
				case "End":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				case "Down":
					k.Keys["DOWN"].Pressed = false
					k.Keys["DOWN"].Released = true
					break
				case "PgDn":
					k.Keys["PgDn"].Pressed = false
					k.Keys["PgDn"].Released = true
					break
				case "Insert":
					k.Keys["INSERT"].Pressed = false
					k.Keys["INSERT"].Released = true
					break
				case "Del":
					k.Keys["DELETE"].Pressed = false
					k.Keys["DELETE"].Released = true
					break
				case "Pause":
					k.Keys["KeyOther"].Pressed = false
					k.Keys["KeyOther"].Released = true
					break
				default:
					//fmt.Printf("\n[INFO] Unknown key.")

					break
				}

			}
			break
		}
	}
}

var keyMap = map[string]*Key{
	"NUMLOCK":   &Key{},
	"Q":         &Key{},
	"W":         &Key{},
	"E":         &Key{},
	"R":         &Key{},
	"T":         &Key{},
	"Y":         &Key{},
	"U":         &Key{},
	"I":         &Key{},
	"O":         &Key{},
	"P":         &Key{},
	"A":         &Key{},
	"S":         &Key{},
	"D":         &Key{},
	"F":         &Key{},
	"G":         &Key{},
	"H":         &Key{},
	"J":         &Key{},
	"K":         &Key{},
	"L":         &Key{},
	"Z":         &Key{},
	"X":         &Key{},
	"C":         &Key{},
	"V":         &Key{},
	"B":         &Key{},
	"N":         &Key{},
	"M":         &Key{},
	"UP":        &Key{},
	"DOWN":      &Key{},
	"LEFT":      &Key{},
	"RIGHT":     &Key{},
	"PgUp":      &Key{},
	"PgDn":      &Key{},
	"KEY0":      &Key{},
	"KEY1":      &Key{},
	"KEY2":      &Key{},
	"KEY3":      &Key{},
	"KEY4":      &Key{},
	"KEY5":      &Key{},
	"KEY6":      &Key{},
	"KEY7":      &Key{},
	"KEY8":      &Key{},
	"KEY9":      &Key{},
	"BACKSPACE": &Key{},
	"DELETE":    &Key{},
	"INSERT":    &Key{},
	"SPACE":     &Key{},
	"HOME":      &Key{},
	"F1":        &Key{},
	"F2":        &Key{},
	"F3":        &Key{},
	"F4":        &Key{},
	"F5":        &Key{},
	"F6":        &Key{},
	"F7":        &Key{},
	"F8":        &Key{},
	"F9":        &Key{},
	"F10":       &Key{},
	"F11":       &Key{},
	"F12":       &Key{},
	"ESC":       &Key{},
	"CapsLock":  &Key{},
	"LSHIFT":    &Key{},
	"RSHIFT":    &Key{},
	"RCTRL":     &Key{},
	"LCTRL":     &Key{},
	"LALT":      &Key{},
	"RALT":      &Key{},
	"ENTER":     &Key{},
	"TAB":       &Key{},
	"[":         &Key{},
	"]":         &Key{},
	"KeyOther":  &Key{},
}

// BindKeyEvent 绑定按键事件
func (k *Keyboard) BindKeyEvent(eventName string, handler func(), keys ...*Key) {
	go func() {
		defer func() {
			//捕获test抛出的panic
			if err := recover(); err != nil {
				fmt.Printf("\n[ERR] Failed to get keyboard: %v", err)
				logger.WriteLog(fmt.Sprint(err))
			}
		}()
		k.Status.EventHandlers[eventName] = 1

		for {
			if k.Status.EventHandlers[eventName] != 1 {
				break
			}

			var allPressed = false
			for {
				allPressed = true
				for _, v := range keys {
					if v.Pressed == false {
						allPressed = false
					}
				}
				if allPressed == true {
					break
				}
				time.Sleep(time.Millisecond * 50)
			}

			var allReleased = false
			for {
				allReleased = true
				for _, v := range keys {
					if v.Released == false {
						allReleased = false
					}
				}
				if allReleased == true {
					break
				}
				time.Sleep(time.Millisecond * 50)
			}
			//if allPressed == true {
			if allReleased {
				handler()
				for _, v := range keys {
					v.Released = false
				}

			} else {
				for _, v := range keys {
					v.Released = false
				}
			}
			continue
			//time.Sleep(time.Millisecond * 50)
			//}
		}
	}()
}

// UnbindKeyEvent 解绑按键事件
func (k *Keyboard) UnbindKeyEvent(eventName string) {
	k.Status.EventHandlers[eventName] = 0
}
