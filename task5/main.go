package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error   // создается переменная err, хранящая (тип, значение), а именно (nil, nil)
	err = test()    // переменной присваивается значение nil, а тип становится customError
	if err != nil { // результат сравнения равен true, поскольку (*customError, nil) != (nil, nil)
		println("error") // вывод строки "error"
		return           // выход из функции
	}
	println("ok")
}
