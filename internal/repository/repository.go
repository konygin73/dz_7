package repository

import (
	"DZ_7/internal/model"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var SlAirplane []model.Airplane
var AirMutex sync.Mutex

const slAirplaneName = "./slAirplaneName"

var SlCar []model.Car
var CarMutex sync.Mutex

const slCarName = "./slCarName"

var SlBoat []model.Boat
var BoatMutex sync.Mutex

const slBoatName = "./slBoatName"

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
		saveToJSON(&SlAirplane)
		AirMutex.Unlock()
		//fmt.Println("append air:", any.GetType())
	case model.Car:
		CarMutex.Lock()
		SlCar = append(SlCar, any.(model.Car))
		saveToJSON(&SlCar)
		CarMutex.Unlock()
		//fmt.Println("append car", any.GetType())
	case model.Boat:
		BoatMutex.Lock()
		SlBoat = append(SlBoat, any.(model.Boat))
		saveToJSON(&SlBoat)
		BoatMutex.Unlock()
		//fmt.Println("append boat", any.GetType())
	default:
		fmt.Println("Unknown type:", t)
	}
}

func saveToJSON(data interface{}) {
	fileName := GetFileNameModel(data)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Error json encode:", err)
	}
}

func GetFileNameModel(data interface{}) string {
	switch data.(type) {
	case *[]model.Airplane:
		return slAirplaneName
	case *[]model.Car:
		return slCarName
	case *[]model.Boat:
		return slBoatName
	}
	return ""
}

func GetFileNameType(any model.TransType) string {
	switch any {
	case model.AirplaneType:
		return slAirplaneName
	case model.CarType:
		return slCarName
	case model.BoatType:
		return slBoatName
	}
	return ""
}

func PrintModel(any model.TransType) {
	fileName := GetFileNameType(any)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error open file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data interface{}
	switch any {
	case model.AirplaneType:
		data = SlAirplane
	case model.CarType:
		data = SlCar
	case model.BoatType:
		data = SlBoat
	}
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error encode file:", err)
		return
	}
	fmt.Println("Slice:", data)
}

func Init() {
	InitModel(&SlAirplane)
	InitModel(&SlCar)
	InitModel(&SlBoat)
}

func InitModel(model interface{}) {
	fileName := GetFileNameModel(model)

	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Error open init file:", fileName, err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(model)
	if err != nil {
		fmt.Println("Error encode file:", err)
		return
	}
}
