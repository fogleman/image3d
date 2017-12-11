package image3d

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

type Image3D struct {
	W, H, D int
	Slices  []*image.NRGBA64
}

func NewImage3D(images []image.Image) *Image3D {
	w := images[0].Bounds().Size().X
	h := images[0].Bounds().Size().Y
	d := len(images)
	slices := make([]*image.NRGBA64, len(images))
	for i, src := range images {
		switch src := src.(type) {
		case *image.NRGBA64:
			slices[i] = src
		default:
			dst := image.NewNRGBA64(src.Bounds())
			draw.Draw(dst, dst.Rect, src, image.ZP, draw.Src)
			slices[i] = dst
		}
	}
	return &Image3D{w, h, d, slices}
}

func (im *Image3D) At(x, y, z float64) color.NRGBA64 {
	x0 := int(math.Floor(x))
	y0 := int(math.Floor(y))
	z0 := int(math.Floor(z))
	if z0 < 0 || z0 >= len(im.Slices) {
		return color.NRGBA64{}
	}
	im0 := im.Slices[z0]

	x -= float64(x0)
	y -= float64(y0)
	z -= float64(z0)
	if x == 0 && y == 0 && z == 0 {
		return im0.NRGBA64At(x0, y0)
	}
	X := 1 - x
	Y := 1 - y
	Z := 1 - z

	x1 := x0 + 1
	y1 := y0 + 1
	z1 := z0 + 1
	if z1 >= im.D {
		return color.NRGBA64{}
	}
	im1 := im.Slices[z1]

	c000 := im0.NRGBA64At(x0, y0)
	c001 := im1.NRGBA64At(x0, y0)
	c010 := im0.NRGBA64At(x0, y1)
	c011 := im1.NRGBA64At(x0, y1)
	c100 := im0.NRGBA64At(x1, y0)
	c101 := im1.NRGBA64At(x1, y0)
	c110 := im0.NRGBA64At(x1, y1)
	c111 := im1.NRGBA64At(x1, y1)
	r000, g000, b000, a000 := c000.RGBA()
	r001, g001, b001, a001 := c001.RGBA()
	r010, g010, b010, a010 := c010.RGBA()
	r011, g011, b011, a011 := c011.RGBA()
	r100, g100, b100, a100 := c100.RGBA()
	r101, g101, b101, a101 := c101.RGBA()
	r110, g110, b110, a110 := c110.RGBA()
	r111, g111, b111, a111 := c111.RGBA()
	r00 := float64(r000)*X + float64(r100)*x
	r01 := float64(r001)*X + float64(r101)*x
	r10 := float64(r010)*X + float64(r110)*x
	r11 := float64(r011)*X + float64(r111)*x
	r0 := r00*Y + r10*y
	r1 := r01*Y + r11*y
	r := uint16(r0*Z + r1*z)
	g00 := float64(g000)*X + float64(g100)*x
	g01 := float64(g001)*X + float64(g101)*x
	g10 := float64(g010)*X + float64(g110)*x
	g11 := float64(g011)*X + float64(g111)*x
	g0 := g00*Y + g10*y
	g1 := g01*Y + g11*y
	g := uint16(g0*Z + g1*z)
	b00 := float64(b000)*X + float64(b100)*x
	b01 := float64(b001)*X + float64(b101)*x
	b10 := float64(b010)*X + float64(b110)*x
	b11 := float64(b011)*X + float64(b111)*x
	b0 := b00*Y + b10*y
	b1 := b01*Y + b11*y
	b := uint16(b0*Z + b1*z)
	a00 := float64(a000)*X + float64(a100)*x
	a01 := float64(a001)*X + float64(a101)*x
	a10 := float64(a010)*X + float64(a110)*x
	a11 := float64(a011)*X + float64(a111)*x
	a0 := a00*Y + a10*y
	a1 := a01*Y + a11*y
	a := uint16(a0*Z + a1*z)
	return color.NRGBA64{r, g, b, a}
}
