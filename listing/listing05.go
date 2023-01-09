package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}
func test() *customError {
	{
		// do something
	}
	return nil
}
func main() {
	var err error
	err = test()

	if err != nil {
		println("error")
		return
	}
	println("ok")
}

/*
Вывод:
error

Функция test возвращает интерфейс, в котором хранится значение типа (customError) и указатель на сам объект.
return nil создает пустой интерфейс нужного типа, но не выделяет память под структуру.

Структура customError имеет метод Error() string, таким образом реализуя базовый интерфейс error, что позволяет
переменной типа error присвоить значение типа *customError. В результате err имеет тип *customError.

Интерфейс будет равен nil только в случае, если он является пустым (и без типа, и значения). Поэтому программа печатает error и завершает работу
*/
