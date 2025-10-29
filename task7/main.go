package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int { // Функция для записи набора чисел типа int в канал
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int { // Функция объединения двух каналов Fan-In like
	c := make(chan int)
	go func() {
		for {
			select { // В случае если имеется возможность получить сообщение из двух каналов, выберется случайный
			case v, ok := <-a: // если можно получить сообщение из канала a
				if ok { // если канал не закрыт
					c <- v
				} else { // если канал закрыт
					a = nil // обнуление переменной, дабы не ожидать с нее сообщений
				}
			case v, ok := <-b: // если можно получить сообщение из канала b
				if ok { // если канал не закрыт
					c <- v
				} else { // если канал закрыт
					b = nil // обнуление переменной, дабы не ожидать с нее сообщений
				}
			}
			if a == nil && b == nil {
				close(c)
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().Unix())
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Print(v)
	}
}
