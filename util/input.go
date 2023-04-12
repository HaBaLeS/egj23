package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var gamepadIDs []ebiten.GamepadID

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
		if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton0) {
			evl[BTN_A] = 1
		} else if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton1) {
			evl[BTN_B] = 1
		} else if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton2) {
			evl[BTN_X] = 1
		} else if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton3) {
			evl[BTN_Y] = 1
		} else if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton4) {
			evl[SHOULDER_LEFT] = 1
		} else if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton5) {
			evl[SHOULDER_RIGHT] = 1
		} else if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton14) {
			evl[LEFT] = 1
		} else if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton15) {
			evl[UP] = 1
		} else if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton12) {
			evl[RIGHT] = 1
		} else if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton13) {
			evl[DOWN] = 1
		} else if ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton11) {
			evl[UP] = 1
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		evl[BTN_A] = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		evl[BTN_A] = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		evl[DOWN] = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		evl[UP] = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		evl[RIGHT] = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		evl[LEFT] = 1
	}

	//fmt.Printf("got %v\n", evl)
	return evl
}
