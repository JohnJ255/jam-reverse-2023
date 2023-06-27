package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/loader"
	"strconv"
)

type ILevelFillter interface {
	Fill(level *Level)
}

type Level struct {
	*framework.Sprite
	index     int
	name      string
	size      framework.Size
	player    *entities.CarEntity
	camera    framework.ICamera
	framework *framework.Framework
	entities  []framework.Entity
}

func NewLevel(index int, g *Game) *Level {
	bgSize := framework.Size{800, 800}
	level := &Level{
		Sprite:   framework.InitSprites(bgSize),
		size:     bgSize,
		index:    index,
		entities: make([]framework.Entity, 0),
	}
	level.name = level.makeName(index)
	level.camera = framework.NewFollowCamera(level.size.Sub(g.WindowSize.AsVec2()), level.Sprite)
	level.LoadResources(&loader.ResourceLoader{}, loader.LevelFileNames[index])

	return level
}

func (l *Level) Init(f *framework.Framework) {
	l.framework = f
	f.FlushCollisions()
	l.Fill()
}

func (l *Level) GetPlayer() framework.Entity {
	return l.player
}

func (l *Level) Update(dt float64) {
	l.camera.Control(l.player)
}

func (l *Level) GetSize() framework.Size {
	return l.size
}

func (l *Level) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	return l.Sprite.PivotTransform(scale, framework.VecUV{})
}

func (l *Level) Change(f *framework.Framework, index int) {
	for _, entity := range l.entities {
		f.RemoveEntity(entity)
	}
	f.FlushCollisions()

	l.Sprite.Imgs = make([]*ebiten.Image, 0)
	l.LoadResources(&loader.ResourceLoader{}, loader.LevelFileNames[index])
	l.index = index
	l.name = l.makeName(index)
	l.Fill()
}

func (l *Level) AddEntity(entity framework.Entity) {
	l.entities = append(l.entities, entity)
	l.framework.AddEntity(entity)
}

func (l *Level) makeName(index int) string {
	return "level " + strconv.Itoa(index)
}

func (l *Level) Fill() {
	var levelFiller ILevelFillter

	switch l.name {
	case "level 1":
		levelFiller = &Level1{}
	case "level 2":
		levelFiller = &Level2{}
	case "level 3":
		levelFiller = &Level3{}
	case "level 4":
		levelFiller = &Level4{}
	case "level 5":
		levelFiller = &Level5{}
	}

	levelFiller.Fill(l)
}

func (l *Level) GetName() string {
	return l.name
}
