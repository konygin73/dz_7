package service

import (
	//"errors"
	"DZ_7/internal/model"
	"math/rand"
)

const COUNT int = 7

var dictionary [COUNT]string = [COUNT]string{"BMW",
	"Skoda", "Volvo", "Toyota", "Opel", "Ford", "Lada"}

func Create() model.TransInterface {
	var ret model.TransInterface
	typeMod := rand.Intn(3) + 1
	nameNum := rand.Intn(COUNT)
	switch typeMod {
	case int(model.AirplaneType):
		airplane := model.Airplane{Trans: model.Trans{Name: dictionary[nameNum],
			Type: model.AirplaneType, Curs: 0}, Height: 5000}
		//fmt.Println("create air")
		ret = airplane
	case int(model.CarType):
		car := model.Car{Trans: model.Trans{Name: dictionary[nameNum],
			Type: model.CarType, Curs: 0}, Weight: 2000}
		//fmt.Println("create car")
		ret = car
	case int(model.BoatType):
		boat := model.Boat{Trans: model.Trans{Name: dictionary[nameNum],
			Type: model.BoatType, Curs: 0}, Depth: 100000}
		//fmt.Println("create boat")
		ret = boat
	}
	return ret
}
