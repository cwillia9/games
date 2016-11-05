package pente

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var (
	img_w, img_h, border_padding = 820, 820, 80
	BLUE                         = color.RGBA{0, 0, 255, 255}
	RED                          = color.RGBA{255, 0, 0, 255}
)

func drawRect() {

	img := image.NewRGBA(image.Rect(0, 0, img_w, img_h))

	black := color.RGBA{0, 0, 0, 255}

	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	drawGrid(img, boardSize, border_padding, img_h, img_w)

	x, y := associatePosToCoords(0, 0)
	r := gridDist(img_h, border_padding, boardSize+1) / 2
	drawPiece(img, x, y, r, BLUE)

	x, y = associatePosToCoords(0, 1)
	drawPiece(img, x, y, r, RED)
	x, y = associatePosToCoords(1, 0)
	drawPiece(img, x, y, r, RED)

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func associatePosToCoords(x, y int) (x_coord, y_coord int) {
	dist := gridDist(img_h, border_padding, boardSize+1)

	x_coord = border_padding + dist + (dist * x)
	y_coord = border_padding + dist + (dist * y)
	return
}

func gridDist(height, padding, lines int) int {
	activeHeight := height - (padding * 2)
	return activeHeight / lines
}

func drawGrid(src *image.RGBA, lines, padding, height, width int) {
	lineDist := gridDist(height, padding, lines+1)

	for i := 1; i <= lines; i++ {
		drawHorizontalLine(src, padding, padding+(i*lineDist))
		drawVerticalLine(src, padding, padding+(i*lineDist))
	}
}

func drawHorizontalLine(src *image.RGBA, padding, y int) {
	white := color.RGBA{255, 255, 255, 255}

	for x := padding; x < img_w-padding; x++ {
		src.Set(x, y, white)
	}

}

func drawVerticalLine(src *image.RGBA, padding, x int) {
	white := color.RGBA{255, 255, 255, 255}

	for y := padding; y < img_w-padding; y++ {
		src.Set(x, y, white)
	}
}

func drawPiece(src *image.RGBA, x, y, r int, clr color.RGBA) {
	r2 := r * r
	for w := -r; w < r; w++ {
		for h := -r; h < r; h++ {
			if ((w * w) + (h * h)) <= r2 {
				src.Set(x+w, y+h, clr)
			}
		}
	}

}
