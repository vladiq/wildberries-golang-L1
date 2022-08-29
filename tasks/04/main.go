package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// функция для воркера, которая заканчивается выполнение при получение значения из канала ctx.Done()
func worker(ctx context.Context, workerId int, numbersToProcess <-chan int) {
	for {
		select {
		case value := <-numbersToProcess:
			fmt.Printf("worker %d got %d\n", workerId, value)
		case <-ctx.Done():
			fmt.Printf("worker %d finished working\n", workerId)
			return
		}
	}
}

func main() {
	// для graceful shutdown будем использовать контекст
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numbersToProcess := make(chan int)
	var wg sync.WaitGroup

	// считывание количество воркеров при запуске программы
	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	// воркеры считывают информацию из канала
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			worker(ctx, workerId, numbersToProcess)
		}(w)
	}

	// записываем информацию в канал
	go func() {
		for i := 0; true; i++ {
			numbersToProcess <- i
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// ловим сигнал в канал
	gracefulShutdown := make(chan os.Signal)
	signal.Notify(gracefulShutdown, syscall.SIGINT)

	<-gracefulShutdown
	fmt.Println("-------------------------------SIGINT CAUGHT------------------------------------")
	// говорим контексту, что нужно закругляться
	cancel()
	// ждём завершения работы всех воркеров
	wg.Wait()
	fmt.Println("Graceful shutdown performed!")
}
