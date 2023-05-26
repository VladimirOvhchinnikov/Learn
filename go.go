package main

var noLocal int = 3 //не локальная переменная
func localint() {

	var local int = 2 //локальная переменная
	println(local)

	var enterf interface{}         // создание переменной типо интерфейс без присвоения
	var enterf_2 = enterf          // создание переменной без указания типа типа интерфейс с присвоением интерфейса
	var enterf_3 interface{} = nil // идентичная запись
	println(enterf_2)
	println(enterf_3)

	/*nil - нужно указывать всегда для ссылочных типов данных
	таких как срезы, указатели, отображений, каналов, функций
	это гарантирует, что переменная имеет значения*/

	var s string
	println(s) //выведет пустую строку, а не ошибку

	var v, t, e = true, "te", 123
	println(v, t, e)
}

func main() {
	localint()
}
