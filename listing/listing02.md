Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
defer - ключевое слово, которое позволяет осуществлять отложенное выполнение функций. Грубо говоря, эти функции выполняются в обратном порядке их объявления, когда окружающая функция завершает выполнение (добавляются в конец стекфрейма).

Конкретно в данном примере мы получаем вывод:

2
1

Это происходит за счёт разницы в работе между конструкциями return <value> и return.

В случае с test(), когда мы пишем return, мы возвращаем значение переменной, которая была объявлена как возвращаемая в самой сигнатуре функции.  Здесь в момент вызова return отложенные функции выполнятся ПЕРЕД возвратом, что позволяет изменить значение x (в данном случае, инкрементировав его).


В случае же с anotherTest(), когда мы пишем return x, мы явно указываем, какое значение должно вернуться из функции. То есть тут функция вернёт нам текущее значение переменной x, которое равно 1 на момент вызова return x. Отложенная функция, инкрементирующая x, вызовется уже ПОСЛЕ возврата, но никак не повлияет на возвращаемое значение.

```