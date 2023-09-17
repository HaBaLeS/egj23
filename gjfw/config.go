package gjfw

var CFG Config = Config{
	ScreenHeight: 720,
	ScreenWidth:  1280,
	Volume:       0,
	Debug:        false,
}

type Config struct {
	ScreenHeight int
	ScreenWidth  int
	Volume       float64
	Debug        bool
}
