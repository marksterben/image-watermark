package helper

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func AddWatermark(src image.Image, text string) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	// Buat canvas baru dengan ukuran yang sama dengan gambar asli
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(rgba, bounds, src, image.Point{}, draw.Src)

	// Tambahkan teks watermark
	col := color.RGBA{255, 255, 255, 255} // Putih
	point := fixed.Point26_6{
		X: fixed.I(10), // Posisi X kiri atas
		Y: fixed.I(20), // Posisi Y kiri atas
	}
	d := &font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)

	return rgba
}
