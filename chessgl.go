package main

import (
	"bytes"
	"image"
	"image/draw"
	_ "image/png"
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var (
	cWIDTH       = 720
	cHEIGHT      = 720
	xSCALE       = float32(cWIDTH) / 8
	ySCALE       = float32(cHEIGHT) / 8
	pieces       uint32
	tile         uint32
	selectedTile uint32
	pubWin       *glfw.Window
)

func handleResize(w *glfw.Window, cWidth int, cHeight int) {
	cWIDTH = cWidth
	cHEIGHT = cHeight
	xSCALE = float32(cWIDTH) / 8
	ySCALE = float32(cHEIGHT) / 8
	setupScene()
	drawScene()
}

func handleClick(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	xpos, ypos := w.GetCursorPos()

	if button == glfw.MouseButton2 {
		movePiece(uint(xpos/float64(xSCALE)), uint(ypos/float64(ySCALE)))
	}

	if button == glfw.MouseButton1 {
		grabPiece(uint(xpos/float64(xSCALE)), uint(ypos/float64(ySCALE)))
	}
}

func setup() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(cWIDTH, cHEIGHT, "Chess", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	pieces = loadPiecesTexture()
	defer gl.DeleteTextures(1, &pieces)

	tile = loadTileTexture()
	defer gl.DeleteTextures(1, &tile)

	selectedTile = loadSelectedTileTexture()
	defer gl.DeleteTextures(1, &selectedTile)

	setupScene()

	window.SetMouseButtonCallback(handleClick)
	window.SetSizeCallback(handleResize)

	pubWin = window

	for !window.ShouldClose() {
		drawScene()
		glfw.WaitEvents()
	}
}

func loadPiecesTexture() uint32 {
	imageBytes, _ := piecesPngBytes()

	img, _, _ := image.Decode(bytes.NewReader(imageBytes))

	rgba := image.NewRGBA(img.Bounds())

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture
}

func loadTileTexture() uint32 {
	imageBytes, _ := tilePngBytes()

	img, _, _ := image.Decode(bytes.NewReader(imageBytes))

	rgba := image.NewRGBA(img.Bounds())

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture
}

func loadSelectedTileTexture() uint32 {
	imageBytes, _ := selecttilePngBytes()

	img, _, _ := image.Decode(bytes.NewReader(imageBytes))

	rgba := image.NewRGBA(img.Bounds())

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture
}

func setupScene() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(cWIDTH), float64(cHEIGHT), 0, -1, 1)
	gl.Viewport(0, 0, int32(cWIDTH), int32(cHEIGHT))
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.DepthRange(0, 1)
}

func destroyScene() {
}

func drawPiece(x float32, y float32, w float32, h float32, piece *piece) {
	var z float32 = 0.8

	x1 := x
	y1 := y
	x2 := x + w
	y2 := y + h

	gl.BindTexture(gl.TEXTURE_2D, pieces)

	gl.Color4f(1, 1, 1, 1)

	gl.Begin(gl.QUADS)

	txmin := float32(piece.class) / 6.0
	txmax := float32(piece.class+1) / 6.0
	tymin := float32(piece.faction) * 0.5
	tymax := tymin + 0.5

	gl.TexCoord2f(txmin, tymin)
	gl.Vertex3f(x1, y1, z)
	gl.TexCoord2f(txmax, tymin)
	gl.Vertex3f(x2, y1, z)
	gl.TexCoord2f(txmax, tymax)
	gl.Vertex3f(x2, y2, z)
	gl.TexCoord2f(txmin, tymax)
	gl.Vertex3f(x1, y2, z)

	gl.End()
}

func drawTile(x float32, y float32, w float32, h float32, s bool) {
	var z float32 = 0.5

	x1 := x
	y1 := y
	x2 := x + w
	y2 := y + h

	if s {
		gl.BindTexture(gl.TEXTURE_2D, selectedTile)
	} else {
		gl.BindTexture(gl.TEXTURE_2D, tile)
	}

	gl.Color4f(1, 1, 1, 1)

	gl.Begin(gl.QUADS)

	gl.TexCoord2f(0, 0)
	gl.Vertex3f(x1, y1, z)

	gl.TexCoord2f(1, 0)
	gl.Vertex3f(x2, y1, z)

	gl.TexCoord2f(1, 1)
	gl.Vertex3f(x2, y2, z)

	gl.TexCoord2f(0, 1)
	gl.Vertex3f(x1, y2, z)

	gl.End()
}

func drawScene() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)

	gl.ClearDepth(1)
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for x := uint(0); x < GRIDLENGTH; x++ {
		for y := uint(0); y < GRIDLENGTH; y++ {
			selectTile := selected != nil && selected.x == x && selected.y == y
			drawTile(float32(x)*xSCALE, float32(y)*ySCALE, xSCALE, ySCALE, selectTile)
		}
	}

	for x := uint(0); x < GRIDLENGTH; x++ {
		for y := uint(0); y < GRIDLENGTH; y++ {
			if board[x][y] != nil {
				drawPiece(float32(x)*xSCALE, float32(y)*ySCALE, xSCALE, ySCALE, board[x][y])
			}
		}
	}

	pubWin.SwapBuffers()
}
