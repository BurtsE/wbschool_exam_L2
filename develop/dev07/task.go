package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var or = func(channels ...<-chan interface{}) <-chan interface{} {
		// Объединенный канал
		res := make(chan interface{})
		defer close(res)

		// Как только все каналы для чтения будут закрыты, завершается выполнение функции и общий канал закрывается
		wg := new(sync.WaitGroup)

		// Создаем горутину для каждого канала. Когда канал закрывается, горутина завершает работу
		for _, channel := range channels {
			wg.Add(1)
			go func(ch <-chan interface{}, wg *sync.WaitGroup) {
				defer wg.Done()
				for {
					if val, ok := <-ch; ok {
						res <- val
					} else {
						return
					}
				}
			}(channel, wg)
		}
		wg.Wait()
		return res
	}
	//Пример использования функции:
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)
	fmt.Printf("fone after %v\n", time.Since(start))
}
