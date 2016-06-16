package main

// Borrowed from:
// https://github.com/polygo/polygo/blob/master/keyboard.go

import "github.com/go-gl/glfw/v3.1/glfw"

type Key int
type ModifierKey int
type Keyboard [glfw.KeyLast]bool

//Keyboard structure manipulation
func (k Keyboard) glfwKeyCallback(
	window *glfw.Window,
	key glfw.Key,
	scancode int,
	action glfw.Action,
	mods glfw.ModifierKey,
) {
	if action == glfw.Press {
		k[key] = true
	} else if action == glfw.Release {
		k[key] = false
	}
}

//Test if a key is down
func (k Keyboard) IsDown(key Key) bool {
	return k[glfw.Key(int(key))]
}

//Input constants
const (
	ModShift   ModifierKey = ModifierKey(int(glfw.ModShift))
	ModControl ModifierKey = ModifierKey(int(glfw.ModControl))
	ModAlt     ModifierKey = ModifierKey(int(glfw.ModAlt))
	ModSuper   ModifierKey = ModifierKey(int(glfw.ModSuper))
)

const (
	KeyUnknown      Key = Key(int(glfw.KeyUnknown))
	KeySpace        Key = Key(int(glfw.KeySpace))
	KeyApostrophe   Key = Key(int(glfw.KeyApostrophe))
	KeyComma        Key = Key(int(glfw.KeyComma))
	KeyMinus        Key = Key(int(glfw.KeyMinus))
	KeyPeriod       Key = Key(int(glfw.KeyPeriod))
	KeySlash        Key = Key(int(glfw.KeySlash))
	Key0            Key = Key(int(glfw.Key0))
	Key1            Key = Key(int(glfw.Key1))
	Key2            Key = Key(int(glfw.Key2))
	Key3            Key = Key(int(glfw.Key3))
	Key4            Key = Key(int(glfw.Key4))
	Key5            Key = Key(int(glfw.Key5))
	Key6            Key = Key(int(glfw.Key6))
	Key7            Key = Key(int(glfw.Key7))
	Key8            Key = Key(int(glfw.Key8))
	Key9            Key = Key(int(glfw.Key9))
	KeySemicolon    Key = Key(int(glfw.KeySemicolon))
	KeyEqual        Key = Key(int(glfw.KeyEqual))
	KeyA            Key = Key(int(glfw.KeyA))
	KeyB            Key = Key(int(glfw.KeyB))
	KeyC            Key = Key(int(glfw.KeyC))
	KeyD            Key = Key(int(glfw.KeyD))
	KeyE            Key = Key(int(glfw.KeyE))
	KeyF            Key = Key(int(glfw.KeyF))
	KeyG            Key = Key(int(glfw.KeyG))
	KeyH            Key = Key(int(glfw.KeyH))
	KeyI            Key = Key(int(glfw.KeyI))
	KeyJ            Key = Key(int(glfw.KeyJ))
	KeyK            Key = Key(int(glfw.KeyK))
	KeyL            Key = Key(int(glfw.KeyL))
	KeyM            Key = Key(int(glfw.KeyM))
	KeyN            Key = Key(int(glfw.KeyN))
	KeyO            Key = Key(int(glfw.KeyO))
	KeyP            Key = Key(int(glfw.KeyP))
	KeyQ            Key = Key(int(glfw.KeyQ))
	KeyR            Key = Key(int(glfw.KeyR))
	KeyS            Key = Key(int(glfw.KeyS))
	KeyT            Key = Key(int(glfw.KeyT))
	KeyU            Key = Key(int(glfw.KeyU))
	KeyV            Key = Key(int(glfw.KeyV))
	KeyW            Key = Key(int(glfw.KeyW))
	KeyX            Key = Key(int(glfw.KeyX))
	KeyY            Key = Key(int(glfw.KeyY))
	KeyZ            Key = Key(int(glfw.KeyZ))
	KeyLeftBracket  Key = Key(int(glfw.KeyLeftBracket))
	KeyBackslash    Key = Key(int(glfw.KeyBackslash))
	KeyRightBracket Key = Key(int(glfw.KeyRightBracket))
	KeyGraveAccent  Key = Key(int(glfw.KeyGraveAccent))
	KeyWorld1       Key = Key(int(glfw.KeyWorld1))
	KeyWorld2       Key = Key(int(glfw.KeyWorld2))
	KeyEscape       Key = Key(int(glfw.KeyEscape))
	KeyEnter        Key = Key(int(glfw.KeyEnter))
	KeyTab          Key = Key(int(glfw.KeyTab))
	KeyBackspace    Key = Key(int(glfw.KeyBackspace))
	KeyInsert       Key = Key(int(glfw.KeyInsert))
	KeyDelete       Key = Key(int(glfw.KeyDelete))
	KeyRight        Key = Key(int(glfw.KeyRight))
	KeyLeft         Key = Key(int(glfw.KeyLeft))
	KeyDown         Key = Key(int(glfw.KeyDown))
	KeyUp           Key = Key(int(glfw.KeyUp))
	KeyPageUp       Key = Key(int(glfw.KeyPageUp))
	KeyPageDown     Key = Key(int(glfw.KeyPageDown))
	KeyHome         Key = Key(int(glfw.KeyHome))
	KeyEnd          Key = Key(int(glfw.KeyEnd))
	KeyCapsLock     Key = Key(int(glfw.KeyCapsLock))
	KeyScrollLock   Key = Key(int(glfw.KeyScrollLock))
	KeyNumLock      Key = Key(int(glfw.KeyNumLock))
	KeyPrintScreen  Key = Key(int(glfw.KeyPrintScreen))
	KeyPause        Key = Key(int(glfw.KeyPause))
	KeyF1           Key = Key(int(glfw.KeyF1))
	KeyF2           Key = Key(int(glfw.KeyF2))
	KeyF3           Key = Key(int(glfw.KeyF3))
	KeyF4           Key = Key(int(glfw.KeyF4))
	KeyF5           Key = Key(int(glfw.KeyF5))
	KeyF6           Key = Key(int(glfw.KeyF6))
	KeyF7           Key = Key(int(glfw.KeyF7))
	KeyF8           Key = Key(int(glfw.KeyF8))
	KeyF9           Key = Key(int(glfw.KeyF9))
	KeyF10          Key = Key(int(glfw.KeyF10))
	KeyF11          Key = Key(int(glfw.KeyF11))
	KeyF12          Key = Key(int(glfw.KeyF12))
	KeyF13          Key = Key(int(glfw.KeyF13))
	KeyF14          Key = Key(int(glfw.KeyF14))
	KeyF15          Key = Key(int(glfw.KeyF15))
	KeyF16          Key = Key(int(glfw.KeyF16))
	KeyF17          Key = Key(int(glfw.KeyF17))
	KeyF18          Key = Key(int(glfw.KeyF18))
	KeyF19          Key = Key(int(glfw.KeyF19))
	KeyF20          Key = Key(int(glfw.KeyF20))
	KeyF21          Key = Key(int(glfw.KeyF21))
	KeyF22          Key = Key(int(glfw.KeyF22))
	KeyF23          Key = Key(int(glfw.KeyF23))
	KeyF24          Key = Key(int(glfw.KeyF24))
	KeyF25          Key = Key(int(glfw.KeyF25))
	KeyKP0          Key = Key(int(glfw.KeyKP0))
	KeyKP1          Key = Key(int(glfw.KeyKP1))
	KeyKP2          Key = Key(int(glfw.KeyKP2))
	KeyKP3          Key = Key(int(glfw.KeyKP3))
	KeyKP4          Key = Key(int(glfw.KeyKP4))
	KeyKP5          Key = Key(int(glfw.KeyKP5))
	KeyKP6          Key = Key(int(glfw.KeyKP6))
	KeyKP7          Key = Key(int(glfw.KeyKP7))
	KeyKP8          Key = Key(int(glfw.KeyKP8))
	KeyKP9          Key = Key(int(glfw.KeyKP9))
	KeyKPDecimal    Key = Key(int(glfw.KeyKPDecimal))
	KeyKPDivide     Key = Key(int(glfw.KeyKPDivide))
	KeyKPMultiply   Key = Key(int(glfw.KeyKPMultiply))
	KeyKPSubtract   Key = Key(int(glfw.KeyKPSubtract))
	KeyKPAdd        Key = Key(int(glfw.KeyKPAdd))
	KeyKPEnter      Key = Key(int(glfw.KeyKPEnter))
	KeyKPEqual      Key = Key(int(glfw.KeyKPEqual))
	KeyLeftShift    Key = Key(int(glfw.KeyLeftShift))
	KeyLeftControl  Key = Key(int(glfw.KeyLeftControl))
	KeyLeftAlt      Key = Key(int(glfw.KeyLeftAlt))
	KeyLeftSuper    Key = Key(int(glfw.KeyLeftSuper))
	KeyRightShift   Key = Key(int(glfw.KeyRightShift))
	KeyRightControl Key = Key(int(glfw.KeyRightControl))
	KeyRightAlt     Key = Key(int(glfw.KeyRightAlt))
	KeyRightSuper   Key = Key(int(glfw.KeyRightSuper))
	KeyMenu         Key = Key(int(glfw.KeyMenu))
	KeyLast         Key = Key(int(glfw.KeyLast))
)
