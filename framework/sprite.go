package framework

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type ResourceLoader interface {
	GetSprite(fileName string) *ebiten.Image
}

type ImageResource interface {
	GetFileNames() []string
	GetBaseAngle() Radian
}

type Sprite struct {
	Visible         bool
	Imgs            []*ebiten.Image
	DrawAngle       Radian
	GetSpriteFunc   func() *ebiten.Image
	CurrentImgIndex int
}

func InitSprites() *Sprite {
	b := &Sprite{
		Visible: true,
		Imgs:    make([]*ebiten.Image, 0),
	}
	b.GetSpriteFunc = b.getBaseSprite
	return b
}

func (b *Sprite) LoadResources(res ResourceLoader, imgs ImageResource) {
	b.DrawAngle = imgs.GetBaseAngle()
	for _, fileName := range imgs.GetFileNames() {
		img := res.GetSprite(fileName)
		b.Imgs = append(b.Imgs, img)
	}
}

func (b *Sprite) GetSprite() *ebiten.Image {
	return b.GetSpriteFunc()
}

func (b *Sprite) getBaseSprite() *ebiten.Image {
	return b.Imgs[b.CurrentImgIndex]
}

func (b *Sprite) PivotTransform(scale float64, size Size, pivot VecUV) *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Rotate(float64(b.DrawAngle))

	spriteSize := b.GetSprite().Bounds().Size()
	scaleX := scale * size.Length / (float64(spriteSize.X)*math.Cos(float64(b.DrawAngle)) + float64(spriteSize.Y)*math.Sin(float64(b.DrawAngle)))
	scaleY := scale * size.Height / (float64(spriteSize.X)*math.Sin(float64(b.DrawAngle)) + float64(spriteSize.Y)*math.Cos(float64(b.DrawAngle)))
	op.GeoM.Scale(scaleX, scaleY)

	tx := -size.Length * (pivot.U - math.Abs(math.Sin(float64(b.DrawAngle))))
	ty := -size.Height * pivot.V
	op.GeoM.Translate(tx, ty)

	return op
}
