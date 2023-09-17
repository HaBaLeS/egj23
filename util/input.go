package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"math"
)

var gamepadIDs []ebiten.GamepadID
var deadZone = 0.3

const (
	BTN_X = iota
	BTN_A
	BTN_B
	BTN_Y
	LEFT
	RIGHT
	UP
	DOWN
	MIDDLE
	START
	MENUE
	SHOULDER_RIGHT
	SHOULDER_LEFT
	AXIS_LEFT_VERT
	AXIS_LEFT_HZN
	AXIS_RIGHT_VERT
	AXIS_RIGHT_HZN
)

func GetInputEvents() map[int]float64 {
	evl := make(map[int]float64, 0) //fixme reuse but clear

	gamepadIDs = ebiten.AppendGamepadIDs(nil)
	for _, id := range gamepadIDs {
		//fmt.Printf("%s -> %s\n", ebiten.GamepadSDLID(id), ebiten.GamepadName(id))
		//maxAxis := ebiten.GamepadAxisCount(id)

		//for a := 0; a < maxAxis; a++ {
		//	v := ebiten.GamepadAxisValue(id, a)
		//	evl[fmt.Sprintf("axis_%d", a)] = v
		//}

		// AXIS
		if v := ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickVertical); math.Abs(v) > deadZone {
			evl[AXIS_LEFT_VERT] = v
		}
		if v := ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickHorizontal); math.Abs(v) > deadZone {
			evl[AXIS_LEFT_HZN] = v
		}
		if v := ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisRightStickVertical); math.Abs(v) > deadZone {
			evl[AXIS_RIGHT_VERT] = v
		}
		if v := ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisRightStickHorizontal); math.Abs(v) > deadZone {
			evl[AXIS_RIGHT_HZN] = v
		}

		// BUTTONS
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton0) {
			evl[BTN_A] = 1
		}
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton1) {
			evl[BTN_B] = 1
		}
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton2) {
			evl[BTN_X] = 1
		}
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton3) {
			evl[BTN_Y] = 1
		}
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton4) {
			evl[SHOULDER_LEFT] = 1
		}
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton5) {
			evl[SHOULDER_RIGHT] = 1
		}
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton14) {
			evl[LEFT] = 1
		}
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton15) {
			evl[UP] = 1
		}
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton12) {
			evl[RIGHT] = 1
		}
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton13) {
			evl[DOWN] = 1
		}
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton11) {
			evl[UP] = 1
		}
	}

	//KEYBOARD
	if inpututil.KeyPressDuration(ebiten.KeySpace) > 0 {
		evl[BTN_A] = 1
	}
	if inpututil.KeyPressDuration(ebiten.KeySpace) > 0 {
		evl[BTN_A] = 1
	}
	if inpututil.KeyPressDuration(ebiten.KeyDown) > 0 {
		evl[DOWN] = 1
	}
	if inpututil.KeyPressDuration(ebiten.KeyUp) > 0 {
		evl[UP] = 1
	}
	if inpututil.KeyPressDuration(ebiten.KeyRight) > 0 {
		evl[RIGHT] = 1
	}
	if inpututil.KeyPressDuration(ebiten.KeyLeft) > 0 {
		evl[LEFT] = 1
	}

	return evl
}
