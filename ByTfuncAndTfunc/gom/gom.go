package gom

type Point struct {
	x float64
}

func (p Point) X() float64 {
	return p.x
}

func (p *Point) SetX(x float64) {
	p.x = x
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}
