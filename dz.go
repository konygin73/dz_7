package main

import (
	"DZ_7/internal/model"
	"DZ_7/internal/repository"
	"DZ_7/internal/service"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan model.TransInterface)
	go func() {
		defer wg.Done()
		count := 0
		for {
			trans := service.Create()
			ch <- trans
			time.Sleep(80 * time.Millisecond)
			count++
			if count > 100 { //8 sec
				close(ch)
				break
			}
		}
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			result, ok := <-ch
			if !ok {
				fmt.Println("close channel")
				return
			}
			repository.Insert(result)
		}
	}()
	wg.Add(1)

	go func() {
		countAir := 0
		countCar := 0
		countBoat := 0
		for {
			time.Sleep(200 * time.Millisecond)
			tmp := repository.GetCount(model.AirplaneType)
			for ; countAir < tmp; countAir++ {
				fmt.Println("air: ", repository.GetAir(countAir).Name)
			}

			tmp = repository.GetCount(model.CarType)
			for ; countCar < tmp; countCar++ {
				fmt.Println("car: ", repository.GetCar(countCar).Name)
			}

			tmp = repository.GetCount(model.BoatType)
			for ; countBoat < tmp; countBoat++ {
				fmt.Println("boat: ", repository.GetBoat(countBoat).Name)
			}
		}
	}()

	fmt.Println("wait...")
	wg.Wait()

	fmt.Println("rez:")
	fmt.Println("count air: ", repository.GetCount(model.AirplaneType))
	fmt.Println("count car: ", repository.GetCount(model.CarType))
	fmt.Println("count boat: ", repository.GetCount(model.BoatType))
}
