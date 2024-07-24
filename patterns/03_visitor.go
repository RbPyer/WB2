package patterns

import "math"

type Shape interface {
}

type Circle struct {
	Radius float64
}

func (c *Circle) accept(visitor ShapeVisitor) {
	visitor.visitCircle(c)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) accept(visitor ShapeVisitor) {
	visitor.visitRectangle(r)
}

type ShapeVisitor interface {
	visitCircle(*Circle)
	visitRectangle(*Rectangle)
}

type AreaVisitor struct {
	area float64
}

func (a *AreaVisitor) visitCircle(circle *Circle) {
	a.area = circle.Radius * circle.Radius * math.Pi
}

func (a *AreaVisitor) visitRectangle(rectangle *Rectangle) {
	a.area = rectangle.Width * rectangle.Height
}

type PerimeterVisitor struct {
	perimeter float64
}

func (p *PerimeterVisitor) visitCircle(circle *Circle) {
	p.perimeter = 2 * math.Pi * circle.Radius
}

func (p *PerimeterVisitor) visitRectangle(rectangle *Rectangle) {
	p.perimeter = 2 * (rectangle.Height + rectangle.Width)
}
