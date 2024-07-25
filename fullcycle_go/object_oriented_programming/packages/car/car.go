package car

type Car struct {
	Name  string // Exported
	color string // Unexported
}

func (c Car) Start() string {
	return c.Name + " has been started"
}
