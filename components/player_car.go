package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/models"
)

type IScoreManager interface {
	AddScore(score int)
}

type PlayerCarControl struct {
	*framework.Component
	levelSize framework.Size
	scores    IScoreManager
	f         *framework.Framework
}

func (c *PlayerCarControl) GetName() string {
	return "PlayerCarControl"
}

func NewPlayerCarControl(levelSize framework.Size, scores IScoreManager) *PlayerCarControl {
	return &PlayerCarControl{
		levelSize: levelSize,
		Component: framework.InitComponent(),
		scores:    scores,
	}
}

func (c *PlayerCarControl) Start(f *framework.Framework) {
	cc := c.GetOwner().(*entities.CarEntity).GetComponent("CarCollision").(*CarCollision)
	defaultFunc := cc.BehaviourOnCollide
	f.Events.AddListener("TrailerCollision", func(event *framework.Event) {
		if c.isPlayersTrailerCollided(event.Data) {
			c.OnCollide(event.Data["collide"].(*framework.Collide))
		}
	})
	cc.BehaviourOnCollide = func(collide *framework.Collide) {
		defaultFunc(collide)
		c.OnCollide(collide)
	}
	c.f = f
	f.Audio.SetVolume("forward", 0.1)
	f.Audio.SetVolume("reverse", 0.1)
	f.Events.AddListener("Forward", func(event *framework.Event) {
		f.Audio.PlayMany("forward", 600)
	})
	f.Events.AddListener("Reverse", func(event *framework.Event) {
		f.Audio.PlayMany("forward", 500)
	})
}

func (c *PlayerCarControl) Update(dt float64) {
	entity := c.GetOwner().(*entities.CarEntity)
	accelerate := 0.0
	wheel := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		wheel = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		wheel = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		accelerate = 1.0
		c.f.Events.Dispatch(&framework.Event{
			Name: "Forward",
		})
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		accelerate = -0.3
		c.f.Events.Dispatch(&framework.Event{
			Name: "Reverse",
		})
	}
	//if ebiten.IsKeyPressed(ebiten.KeyR) {
	//	entity.Car.TowbarToggle()
	//}

	entity.Car.Control(accelerate, wheel)

	entity.Car.Position.X = framework.Limited(entity.Car.Position.X, 0, c.levelSize.Length)
	entity.Car.Position.Y = framework.Limited(entity.Car.Position.Y, 0, c.levelSize.Height)
}

func (c *PlayerCarControl) OnCollide(collide *framework.Collide) {
	if _, ok := collide.Collision.GetEntity().(framework.IPhysicsObject); ok {
		c.scores.AddScore(-1)
	}
}

func (c *PlayerCarControl) isPlayersTrailerCollided(data map[string]interface{}) bool {
	if mtr, ok := data["traktor"].(*models.Car); ok {
		pcar := c.GetOwner().(*entities.CarEntity)
		return pcar.Car == mtr
	}
	return false
}
