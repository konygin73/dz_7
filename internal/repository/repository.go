package repository

import (
	"DZ_7/internal/model"
	"fmt"
	"sync"
)

var SlAirplane []model.Airplane
var AirMutex sync.Mutex
var SlCar []model.Car
var CarMutex sync.Mutex
var SlBoat []model.Boat
var BoatMutex sync.Mutex

func GetCount(any model.TransType) int {
	count := 0
	switch any {
	case model.AirplaneType:
		AirMutex.Lock()
		count = len(SlAirplane)
		AirMutex.Unlock()
	case model.CarType:
		CarMutex.Lock()
		count = len(SlCar)
		CarMutex.Unlock()
	case model.BoatType:
		BoatMutex.Lock()
		count = len(SlBoat)
		BoatMutex.Unlock()
	default:
		fmt.Println("Unknown type: ", any)
	}
	return count
}

func GetAir(item int) model.Airplane {
	AirMutex.Lock()
	ret := SlAirplane[item]
	AirMutex.Unlock()
	return ret
}

func GetCar(item int) model.Car {
	CarMutex.Lock()
	ret := SlCar[item]
	CarMutex.Unlock()
	return ret
}

func GetBoat(item int) model.Boat {
	BoatMutex.Lock()
	ret := SlBoat[item]
	BoatMutex.Unlock()
	return ret
}

func Insert(any model.TransInterface) {
	switch t := any.(type) {
	case model.Airplane:
		AirMutex.Lock()
		SlAirplane = append(SlAirplane, any.(model.Airplane))
		AirMutex.Unlock()
		//fmt.Println("append air:", any.GetType())
	case model.Car:
		CarMutex.Lock()
		SlCar = append(SlCar, any.(model.Car))
		CarMutex.Unlock()
		//fmt.Println("append car", any.GetType())
	case model.Boat:
		BoatMutex.Lock()
		SlBoat = append(SlBoat, any.(model.Boat))
		BoatMutex.Unlock()
		//fmt.Println("append boat", any.GetType())
	default:
		fmt.Println("Unknown type:", t)
	}
}
