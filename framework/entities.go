package framework

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"reverse-jam-2023/helper"
)

type Drawing interface {
	GetSprite() *ebiten.Image
	GetTransforms(scale float64) *ebiten.DrawImageOptions
}

type ResourceLoader interface {
	GetSprite(fileName string) *ebiten.Image
}

type Entity interface {
	GetTransforms(scale float64) *ebiten.DrawImageOptions
	GetSprite() *ebiten.Image
	Update(dt float64)
}

type SpriteEntity struct {
	Visible         bool
	Imgs            []*ebiten.Image
	DrawAngle       float64
	GetSpriteFunc   func() *ebiten.Image
	CurrentImgIndex int
}

func InitSprites(angle float64) *SpriteEntity {
	b := &SpriteEntity{
		Visible:   true,
		DrawAngle: angle,
		Imgs:      make([]*ebiten.Image, 0),
	}
	b.GetSpriteFunc = b.getBaseSprite
	return b
}

func (b *SpriteEntity) LoadResources(res ResourceLoader, fileNames []string) {
	for _, fileName := range fileNames {
		img := res.GetSprite(fileName)
		b.Imgs = append(b.Imgs, img)
	}
}

func (b *SpriteEntity) GetSprite() *ebiten.Image {
	return b.GetSpriteFunc()
}

func (b *SpriteEntity) getBaseSprite() *ebiten.Image {
	return b.Imgs[b.CurrentImgIndex]
}

func (b *SpriteEntity) PivotTransform(scale float64, size helper.Size, pivot helper.PositionUV) *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	spriteSize := image.Point{b.GetSprite().Bounds().Size().Y, b.GetSprite().Bounds().Size().X}
	scaleX := scale * size.Length / float64(spriteSize.X)
	scaleY := scale * size.Width / float64(spriteSize.Y)
	op.GeoM.Rotate(-b.DrawAngle)
	tx := size.Length * (1 - pivot.U)
	ty := -size.Width * pivot.V
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(tx, ty)

	return op
}
