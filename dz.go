package main

import (
	"DZ_7/internal/model"
	"DZ_7/internal/repository"
	"DZ_7/internal/service"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan model.TransInterface)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				close(ch)
				fmt.Println("cancel1")
				return
			case <-time.After(80 * time.Millisecond):
				trans := service.Create()
				fmt.Println("+create: ", trans.GetName(), trans.GetType())
				ch <- trans
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
			//fmt.Println("insert: ", result.GetName())
		}
	}()
	wg.Add(1)

	go func() {
		countAir := 0
		countCar := 0
		countBoat := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel2")
				return
			case <-time.After(200 * time.Millisecond):
				//time.Sleep(200 * time.Millisecond)
				tmp := repository.GetCount(model.AirplaneType)
				for ; countAir < tmp; countAir++ {
					fmt.Println("-air: ", repository.GetAir(countAir).Name)
				}

				tmp = repository.GetCount(model.CarType)
				for ; countCar < tmp; countCar++ {
					fmt.Println("-car: ", repository.GetCar(countCar).Name)
				}

				tmp = repository.GetCount(model.BoatType)
				for ; countBoat < tmp; countBoat++ {
					fmt.Println("-boat: ", repository.GetBoat(countBoat).Name)
				}
			}
		}
	}()

	fmt.Println("wait...")

	sig := <-sigs
	if sig == syscall.SIGINT {
		fmt.Println("ctrl-c")
	}

	cancel()
	wg.Wait()

	fmt.Println("rez:")
	fmt.Println("count air: ", repository.GetCount(model.AirplaneType))
	fmt.Println("count car: ", repository.GetCount(model.CarType))
	fmt.Println("count boat: ", repository.GetCount(model.BoatType))
}
