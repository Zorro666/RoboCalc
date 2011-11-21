package main

import "os"
import "log"
import "fmt"

import "image"
import "time"
import "image/draw"
import "exp/gui"
import "exp/gui/x11"

var (
	red = image.NewColorImage(image.RGBAColor{0xFF, 0, 0, 0xFF})
	green = image.NewColorImage(image.RGBAColor{0x00, 0xFF, 0, 0xFF})
	blue = image.NewColorImage(image.RGBAColor{0x00, 0, 0xFF, 0xFF})
	yellow = image.NewColorImage(image.RGBAColor{0xFF, 0xFF, 0x00, 0xFF})
	white = image.NewColorImage(image.RGBAColor{0xFF, 0xFF, 0xFF, 0xFF})
)

var g_debugLevel int = 0
var g_clearImage bool = true
var g_board Board

type Board struct {
	m_values[5][5] int
	m_columnScores[5] int
	m_rowScores[5] int
}

func (board Board) String() string {
	var ret string = ""
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			ret += fmt.Sprintf("%3d ", board.m_values[x][y])
		}
		ret += fmt.Sprintf("= %3d", board.m_rowScores[y])
		ret += "\n"
	}
	ret += " ||  ||  ||  ||  ||\n"
	for x := 0; x < 5; x++ {
		ret += fmt.Sprintf("%3d ", board.m_columnScores[x])
	}

	return ret
}

func ComputeScore(values[5] int) (score int) {
	var counts[6] int
	for i := 0; i < 6; i++ {
		counts[0] = 0
	}
	for i := 0; i < 5; i++ {
		var v = values[i]
		counts[v]++
	}
	score = 0
	for i := 0; i < 6; i++ {
		var count int = counts[i]
		switch count {
			case 1:
				score += i
			case 2:
				score += 10*i
			case 3, 4, 5:
				score += 100
		}
	}
	return
}

func (board *Board) GetRow(row int) (values[5] int) {
	for x := 0; x < 5; x++ {
		values[x] = board.m_values[x][row]
	}
	return
}

func (board *Board) GetColumn(column int) (values[5] int) {
	for y := 0; y < 5; y++ {
		values[y] = board.m_values[column][y]
	}
	return
}

func (board *Board) ComputeScores() {
	for y := 0; y < 5; y++ {
		board.m_rowScores[y] = ComputeScore( board.GetRow(y) )
	}
	for x := 0; x < 5; x++ {
		board.m_columnScores[x] = ComputeScore( board.GetColumn(x) )
	}
}

func render(window gui.Window) {

	for {
		if g_clearImage == true {
			var canvas draw.Image = window.Screen()
			renderFrame(canvas)
			window.FlushImage()
		}
		time.Sleep(1)
	}
}

func Log(format string, a ...interface{}) {
	if g_debugLevel > 0 {
		log.Printf(format, a...)
	}
	return
}

func renderFrame(canvas draw.Image) {

	if g_clearImage == true {
		//var width int = canvas.Bounds().Max.X - canvas.Bounds().Min.X
		//var height int = canvas.Bounds().Max.Y - canvas.Bounds().Min.Y

		draw.Draw(canvas, canvas.Bounds(), image.Transparent, image.ZP, draw.Src)
		log.Println("Clear Image", canvas.Bounds())
		g_board.m_values[0][0] = 0
		g_board.m_values[1][0] = 0
		g_board.m_values[2][0] = 0
		g_board.m_values[3][0] = 3
		g_board.m_values[4][0] = 4
		g_board.m_values[0][1] = 5
		g_board.m_values[1][1] = 0
		g_board.m_values[2][1] = 1
		g_board.m_values[3][1] = 1
		g_board.m_values[4][1] = 3
		g_board.m_values[0][2] = 4
		g_board.m_values[1][2] = 5
		g_board.m_values[2][2] = 0
		g_board.m_values[3][2] = 1
		g_board.m_values[4][2] = 2
		g_board.m_values[0][3] = 3
		g_board.m_values[1][3] = 3
		g_board.m_values[2][3] = 5
		g_board.m_values[3][3] = 0
		g_board.m_values[4][3] = 1
		g_board.m_values[0][4] = 2
		g_board.m_values[1][4] = 3
		g_board.m_values[2][4] = 4
		g_board.m_values[3][4] = 5
		g_board.m_values[4][4] = 5
		g_board.ComputeScores()

		fmt.Println(g_board)
		g_clearImage = false
	}
}

func handleMouseEvent(mouseEvent gui.MouseEvent) {
	Log("Mouse Event Buttons 0x%X Position:%d,%d\n", mouseEvent.Buttons, mouseEvent.Loc.X, mouseEvent.Loc.Y)

	if mouseEvent.Loc.X == 0 && mouseEvent.Loc.Y == 0 {
		return
	}

	// mouseEvent.Buttons == 0x1 {
	// mouseEvent.Buttons == 0x4 {
}

type MyKeyEvent struct {
	drawKeyEvent gui.KeyEvent
}

func (keyEvent MyKeyEvent) String() string {
	key := keyEvent.drawKeyEvent.Key
	isPressed := "Press"
	if key < 0 {
		isPressed = "Release"
		key = -key
	}

	keyString := "'UNKNOWN'"
	if ' ' <= key && key <= 'z' {
		keyString = fmt.Sprintf("'%c'", key)
	}

	return fmt.Sprintf("%s %s %v 0x%X", isPressed, keyString, key, key)
}

func clampInt(value int, min int, max int) (result int) {
	result = value
	if result < min {
		result = min
	}
	if result > max {
		result = max
	}
	return result
}

func handleKeyEvent(keyEvent gui.KeyEvent) {
	var myKeyEvent MyKeyEvent
	myKeyEvent.drawKeyEvent = keyEvent
	log.Println("Key Event", myKeyEvent)
	// Key Release (key -ve)
	if keyEvent.Key < 0 {
		var key uint = uint(-keyEvent.Key)
		var numberKey uint = uint(key) -0x31
		if numberKey < 5 {
		}

		if key == 'r' {
			g_clearImage = true
		}
		if key == 'd' {
			g_debugLevel += 1
		}
		if key == 'D' {
			g_debugLevel -= 1
		}
		g_debugLevel = clampInt(g_debugLevel, 0, 5)
	}
}

type Empty interface{}

func main() {
	var mainWindow gui.Window
	var error os.Error
	log.SetFlags(log.Ltime|log.Lmicroseconds)
	mainWindow, error = x11.NewWindow()

	if error != nil {
		log.Fatalf("%s", error.String())
	}

	go render(mainWindow)

loop:
	for {
		var windowEvent Empty = <-mainWindow.EventChan()
		switch event := windowEvent.(type) {
		case gui.MouseEvent:
			handleMouseEvent(event)
		case gui.KeyEvent:
			handleKeyEvent(event)
			if event.Key == 65307 { // ESC
				break loop
			}
		case gui.ConfigEvent:
			log.Printf("Config Event\n")
		case gui.ErrEvent:
			log.Printf("Error Event\n")
			break loop
		}

	}
	error = mainWindow.Close()
}
