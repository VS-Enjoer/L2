Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```
package main
 
import (
    "fmt"
    "os"
)
 
func Foo() error {
    var err *os.PathError = nil
    return err
}
 
func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}
```
Ответ:
`<nil>`
`true`
Присваиваем переменной `err` значение из функции `Foo()` значение из функции `Foo()`
будет равно типу `error` со значением `nil`. Соответственно вывод `nil` в 1 случае, во 2 `true`.

По поводу интерфейсов, Интерфейсы имеют 2 значения `value` и `type`. Интерфейсы нужны для реализации каких-либо методов,
соответственно не пустой интерфейс будет подразумевать что его кто-то реализует.
Пустой интерфейс нужен когда мы еще не понимаем какой тип данных может прийти.