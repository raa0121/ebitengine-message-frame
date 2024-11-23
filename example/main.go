package main

import (
	"bytes"
	"embed"
	"image/color"
	"image/png"
	"path"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/raa0121/ebitengine-message-frame"
)

//go:embed assets
var assets embed.FS

var (
	dotGothic16Source *text.GoTextFaceSource
	dotGothic16Face   *text.GoTextFace
	frameImage        *ebiten.Image
)

func init() {
	font, err := assets.ReadFile(path.Join("assets", "DotGothic16-Regular.ttf"))
	if err != nil {
		panic(err)
	}
	s, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	dotGothic16Source = s
	dotGothic16Face = &text.GoTextFace{
		Source: dotGothic16Source,
		Size: 32,
	}

	frameFile, err := assets.Open(path.Join("assets", "frame.png"))
	if err != nil {
		panic(err)
	}
	img, err := png.Decode(frameFile)
	frameImage = ebiten.NewImageFromImage(img) 
}

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x0, 0x0, 0xff, 0xff})
	frame.DrawFrame(screen, frameImage, []string{"abcd", "あああいいいうううえええ", "漢字テスト"}, dotGothic16Face, 10, 10)
	frame.DrawFrame(screen, frameImage, []string{"アメンボ赤いなあいうえお", "赤巻紙青巻紙黄巻紙", "隣の客はよく柿食う客だ"}, dotGothic16Face, 110, 310)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	game := &Game{}
	ebiten.SetWindowTitle("フレームテスト")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
	if runtime.GOARCH == "wasm" {
		done := make(chan struct{}, 0)
		<-done
	}
}
