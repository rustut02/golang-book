# 2.4.1 Присваивание кортежу
* Присваивание кортежу - это метод, который позволяет присваивать значения нескольким переменным одновременно, делая код более компактным и понятным;
* Все выражения справа вычисляются перед присваиванием значений переменным слева, что позволяет обменивать значения
  переменных, например, вычислять наибольший общий делитель или числа Фибоначчи:
	* ``` go
      func fib(n int) int {
        x, y, := 0, 1
        for i := 0; i < n; i++ {
            x, y = y, x+y
        }
        return x
      }
      ```
* Присваивание кортежу удобно для обмена значений между переменными, такими как `x, y = y, x` или `a[i], a[j] = a[j], a[i]`;
* Для функций с несколькими результатами присваивание кортежу используется для того, чтобы присвоить значения каждому из результатов, например: `f, err = os.Open("foo.txt")`;
* Дополнительные результаты функций часто используются для указания на ошибки, возвращая значение `error` или булево значение `ok`;
* В присваивании кортежу можно использовать пустой идентификатор `_`, чтобы игнорировать ненужные значения, например:  `_, err = io.Copy(dst, src)`;
* Присваивание кортежу повышает читаемость и понятность кода, но нужно избегать его использования при наличии сложных выражений, так как последовательность отдельных инструкций может быть легче для чтения.