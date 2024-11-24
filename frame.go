package frame

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func DrawFrame(screen, frameImage *ebiten.Image, strs []string, face *text.GoTextFace, x, y float64) {
	var ws, hs []float64
	for _, str := range strs {
		w, h := text.Measure(str, face, face.Size)
		ws = append(ws, w)
		hs = append(hs, h)
	}
	wsMax := int(slices.Max(ws))
	hsMax := int(slices.Min(hs))
	pitch := NewNinePatchOfSize(frameImage, 10, 0, 0, wsMax + 10, hsMax * len(strs))
	pitch.Draw(screen, int(x), int(y), wsMax + 20, hsMax * len(strs) + 20)
	for i, str := range strs {
		_, h := text.Measure(str, face, face.Size)
		op := &text.DrawOptions{}
		op.GeoM.Translate(x + 10 , y + 10 + h * float64(i))
		text.Draw(screen, str, face, op)
	}
}

func DrawChoiceFrame(screen, frameImage *ebiten.Image, strs []string, choice int, face *text.GoTextFace, x, y float64) {
	if choice < 0 {
		choice = choice % len(strs) + len(strs)
	}
	if len(strs) - 1 < choice {
		choice = choice % len(strs)
	}
	for i, str := range strs {
		if i == choice {
			strs[choice] = "→" + str
		} else {
			strs[i] = "　" + str
		}
	}
	DrawFrame(screen, frameImage, strs, face, x, y)
}
