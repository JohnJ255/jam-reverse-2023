package framework

type Component interface {
	Start()
	Update(dt float64)
	Disable()
	Enable()
	IsEnabled() bool
}

type GameEntity struct {
	Name       string
	Enabled    bool
	Components []Component
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

func (g *GameEntity) AddComponent(c Component) {
	g.Components = append(g.Components, c)
}
