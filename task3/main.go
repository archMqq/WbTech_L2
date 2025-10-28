package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil // конкретный тип *os.PathError, но присваивается значение nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        // вывод nil
	fmt.Println(err == nil) // false, поскольку err интерфейс err конкретного типа, а не nil
}

// Интерфейс хранит значение (тип, значение)
// nil хранит значение (nil, nil)
// В данном примере nil это (тип, nil)

// error - это интерфейс с одним единственным методом Error() string
// interface{} - это пустой интерфейс не имеющий методов, то есть его инициализирует любой тип. Тоже содержит (тип, значение)
