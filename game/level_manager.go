package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/loader"
	"strconv"
)

const NewLevelScore = 1000

const (
	LevelIndex1 = iota + 1
	LevelIndex2
	LevelIndex3
	LevelIndex4
	LevelIndex5
	LevelIndex6

	LastLevelIndex
)

type ILevelFillter interface {
	Fill(level *LevelManager)
	GetSize() framework.Size
}

type LevelManager struct {
	*framework.Sprite
	index     int
	name      string
	size      framework.Size
	player    *entities.CarEntity
	camera    framework.ICamera
	framework *framework.Framework
	entities  []framework.Entity
	Score     int
}

func NewLevel(index int, g *Game) *LevelManager {
	bgSize := framework.Size{800, 800}
	level := &LevelManager{
		Sprite:   framework.InitSprites(bgSize),
		size:     bgSize,
		index:    index,
		entities: make([]framework.Entity, 0),
	}
	level.camera = framework.NewFollowCamera(level.size.Sub(g.WindowSize.AsVec2()), level.Sprite)

	return level
}

func (l *LevelManager) Init(f *framework.Framework) {
	l.framework = f
	l.Change(f, l.index)
}

func (l *LevelManager) GetPlayer() framework.Entity {
	return l.player
}

func (l *LevelManager) Update(dt float64) {
	if l.camera != nil && l.player != nil {
		l.camera.Control(l.player)
	}
}

func (l *LevelManager) GetSize() framework.Size {
	return l.size
}

func (l *LevelManager) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	l.Sprite.SetToSize(l.size)
	return l.Sprite.PivotTransform(scale, framework.VecUV{})
}

func (l *LevelManager) Change(f *framework.Framework, index int) {
	for _, entity := range l.entities {
		f.RemoveEntity(entity)
	}
	f.FlushCollisions()

	if index == 1 {
		l.Score = 0
	}
	if index > 0 && index != LastLevelIndex {
		l.Score += NewLevelScore

		if l.index < index {
			f.Audio.PlayOnce("win")
		}
	}
	l.Sprite.Imgs = make([]*ebiten.Image, 0)
	l.index = index
	l.name = l.makeName(index)

	if _, ok := loader.LevelFileNames[index]; ok {
		l.LoadResources(&loader.ResourceLoader{}, loader.LevelFileNames[index])
		l.Fill()
	}
}

func (l *LevelManager) AddEntity(entity framework.Entity) {
	l.entities = append(l.entities, entity)
	l.framework.AddEntity(entity)
}

func (l *LevelManager) makeName(index int) string {
	return "level " + strconv.Itoa(index)
}

func (l *LevelManager) Fill() {
	levels := []ILevelFillter{
		nil,
		&Level1{},
		&Level2{},
		&Level3{},
		&Level4{},
		&Level5{},
		&Level6{},
		&LevelAbout{},
	}

	levelFiller := levels[l.index]
	l.size = levelFiller.GetSize()

	levelFiller.Fill(l)
}

func (l *LevelManager) GetName() string {
	return l.name
}

func (l *LevelManager) AddScore(score int) {
	l.Score += score
}
