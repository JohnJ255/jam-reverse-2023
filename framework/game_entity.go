package framework

type IComponent interface {
	Start()
	Update(dt float64)
	Disable()
	Enable()
	IsEnabled() bool
	GetOwner() Updating
	SetOwner(g Updating)
}

type GameEntity struct {
	Name       string
	Enabled    bool
	Components []IComponent
}

func InitGameEntity() *GameEntity {
	return &GameEntity{
		Components: make([]IComponent, 0),
	}
}

func (g *GameEntity) Start() {
	for _, c := range g.Components {
		if c.IsEnabled() {
			c.Start()
		}
	}
}

func (g *GameEntity) Update(dt float64) {
	for _, c := range g.Components {
		if c.IsEnabled() {
			c.Update(dt)
		}
	}
}

func (g *GameEntity) AddComponent(c IComponent) {
	g.Components = append(g.Components, c)
}
