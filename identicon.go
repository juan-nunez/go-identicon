package identicon

import (
	"crypto/md5"
	"image/color"
	"image"
	"os"
	"fmt"
	"image/png"
)

type identicon struct {
	key string
	hash [16]byte
	dimension int
	pixelDimension int
}

func NewIdenticon(key string) *identicon {
	icon := &identicon{key: key, dimension: 4, pixelDimension: 256}
	icon.hash = icon.generate([]byte(key))
	return icon
}

func (icon identicon) generate(key []byte) [16]byte{
	return md5.Sum(key)
}

func (icon identicon) getColorFromHash() color.NRGBA {
	rgb := icon.hash[13:]
	return color.NRGBA{
		rgb[0],
		rgb[1],
		rgb[2],
		255,
	}
}

func (icon identicon) ToImage() bool {
	hashColor := icon.getColorFromHash()
	white := color.NRGBA{255,255,255,255}
	img := image.NewNRGBA(image.Rect(0,0,icon.pixelDimension,icon.pixelDimension))
	bounds := icon.getBounds()
	for i, b := range icon.hash {
		if (b & 1) == 1 {
			icon.drawRect(img, bounds[i], hashColor)
		} else {
			icon.drawRect(img, bounds[i], white)
		}
	}

	f, err := os.Create(icon.key + ".png")
	if err != nil {
		fmt.Println("Something went wrong creating file")
		return false
	}
	err = png.Encode(f, img)
	if err != nil {
		fmt.Println("Something went wrong encoding to png")
		return false
	}
	return true
}

type bound struct {
	x0 int
	y0 int
	x1 int
	y1 int
}

func (icon identicon) getBounds() []bound{
	bounds := make([]bound, 16)
	chunk := 64
	idx := 0
	for i := 0; i < icon.dimension; i++ {
		for j := 0; j < icon.dimension; j++ {
			bound := bound{chunk * i, chunk * j, (chunk * i) + chunk, (chunk*j) + chunk}
			bounds[idx] = bound
			idx++
		}
	}
	return bounds
}

func (icon identicon) drawRect(img *image.NRGBA, bound bound, color color.NRGBA) {
	x0 := bound.x0
	y0 := bound.y0
	x1 := bound.x1
	y1 := bound.y1
	for i := x0; i < x1; i++ {
		for j := y0; j < y1; j++ {
			img.SetNRGBA(i, j, color)
		}
	}

}







