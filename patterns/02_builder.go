package patterns

type Director struct {
	builder HouseBuilder
}

func NewDirector(builder HouseBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) BuildHouse() House {
	return d.builder.BuildWalls("wood").
		BuildWindows("glass").
		BuildRoof("Bricks").
		BuildDoor("steel").
		GetResult()
}

type House struct {
	Roof    string
	Walls   string
	Door    string
	Windows string
}

type HouseBuilder interface {
	BuildRoof(material string) HouseBuilder
	BuildWalls(material string) HouseBuilder
	BuildDoor(material string) HouseBuilder
	BuildWindows(material string) HouseBuilder
	GetResult() House
}

type ConcreteHouseBuilder struct {
	house House
}

func NewConcreteHouseBuilder() *ConcreteHouseBuilder {
	return &ConcreteHouseBuilder{
		house: House{},
	}
}

func (b *ConcreteHouseBuilder) BuildRoof(material string) HouseBuilder {
	b.house.Roof = material
	return b
}

func (b *ConcreteHouseBuilder) BuildWalls(material string) HouseBuilder {
	b.house.Walls = material
	return b
}

func (b *ConcreteHouseBuilder) BuildDoor(material string) HouseBuilder {
	b.house.Door = material
	return b
}

func (b *ConcreteHouseBuilder) BuildWindows(material string) HouseBuilder {
	b.house.Windows = material
	return b
}

func (b *ConcreteHouseBuilder) GetResult() House {
	return b.house
}
