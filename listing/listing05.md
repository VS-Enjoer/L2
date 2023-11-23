Что выведет программа? Объяснить вывод программы.

```
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

```

Ответ: `error`

Здесь в функции `test()`, хоть она и возвращает `nil`, она возвращает `nil` указатель типа `*customError`, значит 
она возвращает `nil` не как значение типа `customError`, а как нулевой указатель этого типа. Когда мы присваиваем 
результат `test()` переменной `err` типа `error`, этот `nil` указатель все равно будет интерфейсным значением, и условие 
`if err != nil` выполнится, поскольку `err` не является `nil` по интерфейсу.

