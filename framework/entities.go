package framework

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawing interface {
	GetSprite() *ebiten.Image
	GetTransforms(scale float64) *ebiten.DrawImageOptions
}

type Controlling interface {
	Drawing
	Control(params interface{})
}

type ResourceLoader interface {
	GetSprite(fileName string) *ebiten.Image
}

type SpriteEntity struct {
	Imgs            []*ebiten.Image
	DrawAngle       float64
	GetSpriteFunc   func() *ebiten.Image
	CurrentImgIndex int
}

func InitSprites(angle float64) *SpriteEntity {
	b := &SpriteEntity{
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
