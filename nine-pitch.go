// https://gist.github.com/tmathews/42f663e85c333791d720d7911347eb77
package frame

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type NinePatch struct {
	tiles [9]*ebiten.Image
}

func NewNinePatchOfSize(src *ebiten.Image, size, x, y, w, h int) *NinePatch {
	return NewNinePatch(src, image.Rectangle{
		Min: image.Point{X: x, Y: y},
		Max: image.Point{X: x + w, Y: y + h},
	}, image.Rectangle{
		Min: image.Point{X: x + size, Y: y + size},
		Max: image.Point{X: x + size*2, Y: y + size*2},
	})
}

func NewNinePatch(src *ebiten.Image, out, in image.Rectangle) *NinePatch {
	var n NinePatch
	// Top Left
	n.tiles[0] = NewSubImageAt(src,
		out.Min.X, out.Min.Y,
		in.Min.X, in.Min.Y)
	// Top Center
	n.tiles[1] = NewSubImageAt(src,
		in.Min.X, out.Min.Y,
		in.Max.X, in.Min.Y)
	// Top Right
	n.tiles[2] = NewSubImageAt(src,
		in.Max.X, out.Min.Y,
		out.Max.X, in.Min.Y)
	// Center Left
	n.tiles[3] = NewSubImageAt(src,
		out.Min.X, in.Min.Y,
		in.Min.X, in.Max.Y)
	// Center
	n.tiles[4] = NewSubImageAt(src,
		in.Min.X, in.Min.Y,
		in.Max.X, in.Max.Y)
	// Center Right
	n.tiles[5] = NewSubImageAt(src,
		in.Max.X, in.Min.Y,
		out.Max.X, in.Max.Y)
	// Bottom Left
	n.tiles[6] = NewSubImageAt(src,
		out.Min.X, in.Max.Y,
		in.Min.X, out.Max.Y)
	// Bottom Center
	n.tiles[7] = NewSubImageAt(src,
		in.Min.X, in.Max.Y,
		in.Max.X, out.Max.Y)
	// Bottom Right
	n.tiles[8] = NewSubImageAt(src,
		in.Max.X, in.Max.Y,
		out.Max.X, out.Max.Y)
	return &n
}

func (n *NinePatch) Draw(dst *ebiten.Image, x, y, w, h int) {
	var img *ebiten.Image
	var opts *ebiten.DrawImageOptions
	var iw, ih, cx1, cy1, cx2, cy2 int

	// Top Left
	img = n.tiles[0]
	iw, ih = img.Size()
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	dst.DrawImage(img, opts)
	cx1, cy1 = x+iw, y+ih
	// Top Right
	img = n.tiles[2]
	iw, ih = img.Size()
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x+w-iw), float64(y))
	dst.DrawImage(img, opts)
	// Bottom Left
	img = n.tiles[6]
	iw, ih = img.Size()
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(y+h-ih))
	dst.DrawImage(img, opts)
	// Bottom Right
	img = n.tiles[8]
	iw, ih = img.Size()
	cx2, cy2 = x+w-iw, y+h-ih
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(cx2), float64(cy2))
	dst.DrawImage(img, opts)
	// Center
	img = n.tiles[4]
	iw, ih = img.Size()
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(cx1), float64(cy1))
	for x := cx1; x < cx2; x += iw {
		var count int
		for y := cy1; y < cy2; y += ih {
			dst.DrawImage(img, opts)
			opts.GeoM.Translate(0, float64(ih))
			count++
		}
		opts.GeoM.Translate(float64(iw), float64(-count*ih))
	}
	// North
	img = n.tiles[1]
	iw, ih = img.Size()
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(cx1), float64(y))
	for x := cx1; x < cx2; x += iw {
		dst.DrawImage(img, opts)
		opts.GeoM.Translate(float64(iw), 0)
	}
	// South
	img = n.tiles[7]
	iw, ih = img.Size()
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(cx1), float64(cy2))
	for x := cx1; x < cx2; x += iw {
		dst.DrawImage(img, opts)
		opts.GeoM.Translate(float64(iw), 0)
	}
	// East
	img = n.tiles[3]
	iw, ih = img.Size()
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x), float64(cy1))
	for y := cy1; y < cy2; y += ih {
		dst.DrawImage(img, opts)
		opts.GeoM.Translate(0, float64(ih))
	}
	// West
	img = n.tiles[5]
	iw, ih = img.Size()
	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(cx2), float64(cy1))
	for y := cy1; y < cy2; y += ih {
		dst.DrawImage(img, opts)
		opts.GeoM.Translate(0, float64(ih))
	}
}

func NewSubImageAt(src *ebiten.Image, x1, y1, x2, y2 int) *ebiten.Image {
	return src.SubImage(image.Rectangle{
		Min: image.Point{X: x1, Y: y1},
		Max: image.Point{X: x2, Y: y2},
	}).(*ebiten.Image)
}

