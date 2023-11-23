Что выведет программа? Объяснить вывод программы

```
package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func asChan(vs ...int) <-chan int {
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
 
func merge(a, b <-chan int) <-chan int {
   c := make(chan int)
   go func() {
       for {
           select {
               case v := <-a:
                   c <- v
              case v := <-b:
                   c <- v
           }
      }
   }()
 return c
}
 
func main() {
 
   a := asChan(1, 3, 5, 7)
   b := asChan(2, 4 ,6, 8)
   c := merge(a, b )
   for v := range c {
       fmt.Println(v)
   }
}

```

Ответ: `1` `2` `3` `4` `5` `6` `7` `8` (В рандомном порядке) и потом будут выводиться нули пока не завершим программу.

Тут проблемка в том что мы не проверяем статус канала, закрыт он или нет. Т.к. канал закрыт и мы с него читаем, мы читаем
дефолтное значение, для `int` это как раз `0`, поэтому так много нулей в выводе
По хорошему нужно добавить `ok` в `select`

```
    case v, ok := <-a:
        if !ok {
            return
        }
        c-<v
    case v, ok := <-b:
        if !ok {
            return
        }
        c-<v
```