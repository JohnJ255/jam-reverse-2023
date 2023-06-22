package framework

type IComponent interface {
	GetName() string
	Start(f *Framework)
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

func (g *GameEntity) Start(f *Framework) {
	for _, c := range g.Components {
		if c.IsEnabled() {
			c.Start(f)
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

func (g *GameEntity) GetComponents() []IComponent {
	return g.Components
}

func (g *GameEntity) GetComponent(name string) IComponent {
	for _, c := range g.Components {
		if c.GetName() == name {
			return c
		}
	}
	return nil
}
