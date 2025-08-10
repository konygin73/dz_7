package model

type TransType byte

const (
	AnyType TransType = iota
	CarType
	AirplaneType
	BoatType
)

type TransInterface interface {
	GetName() string
	GetType() TransType
	StepLeft()
	StepRight()
}

type Trans struct {
	Name string
	Type TransType
	Curs int
}

func (th Trans) GetName() string {
	return th.Name
}

func (th Trans) GetType() TransType {
	return th.Type
}

type Airplane struct {
	Trans
	Height int
}

func (th Airplane) StepLeft() {
	th.Curs -= 5
}

func (th Airplane) StepRight() {
	th.Curs += 5
}

type Car struct {
	Trans
	Weight int
}

func (th Car) StepLeft() {
	th.Curs -= 3
}

func (th Car) StepRight() {
	th.Curs += 3
}

type Boat struct {
	Trans
	Depth int
}

func (th Boat) StepLeft() {
	th.Curs -= 3
}

func (th Boat) StepRight() {
	th.Curs += 3
}
