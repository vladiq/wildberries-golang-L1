package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. Создаём канал и от спустя заданное время отправляем через него сигнал о завершении
	stop := make(chan struct{})

	go func() {
		select {
		case <-stop:
			fmt.Println("received stop signal, stopping function 1")
			return
		}
	}()

	stop <- struct{}{}
	time.Sleep(time.Millisecond * 300)

	// 2. То же самое, но через контекст
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("received ctx.Done signal, stopping function 2")
			return
		}
	}(ctx)

	cancel()
	time.Sleep(time.Millisecond * 300)

	// 3. Закрытие канала
	dataChan := make(chan int)

	go func() {
		for {
			if _, ok := <-dataChan; !ok {
				fmt.Println("channel closed, stopping function 3")
				return
			}
		}
	}()

	dataChan <- 1234
	close(dataChan)
	time.Sleep(time.Millisecond * 300)

	// 4. Используем time.After
	dataChan = make(chan int)

	go func() {
		for {
			select {
			case <-dataChan:
				fmt.Println("got a value from dataChan")
			case <-time.After(time.Second):
				fmt.Println("got a value from time.After, stopping function 4")
				return
			}
		}
	}()

	dataChan <- 1234
	time.Sleep(time.Second * 2)
}
