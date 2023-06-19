package framework

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
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

func (b *SpriteEntity) PivotTransform(scale float64, size helper.Size, pivot helper.VecUV) *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Rotate(b.DrawAngle)

	spriteSize := b.GetSprite().Bounds().Size()
	scaleX := scale * size.Length / (float64(spriteSize.X)*math.Cos(b.DrawAngle) + float64(spriteSize.Y)*math.Sin(b.DrawAngle))
	scaleY := scale * size.Width / (float64(spriteSize.X)*math.Sin(b.DrawAngle) + float64(spriteSize.Y)*math.Cos(b.DrawAngle))
	op.GeoM.Scale(scaleX, scaleY)

	tx := size.Length * (1 - pivot.U)
	ty := -size.Width * pivot.V
	op.GeoM.Translate(tx, ty)

	return op
}
