# 2.3 Переменные
* Каждое объявление имеет общий вид `var name type = expression`
* Тип и начальное значение переменной могут быть опущены, но не одновременно. При опущенном типе он определяется из инициализирующего выражения, при опущенном значении - присваивается нулевое значение для данного типа (`var s := 1`; `var s int`, соответственно);
* Множество переменных может быть инициализировано с помощью вызова функции, возвращающей несколько значений. Например, функция `os.Open(name)` возвращает файл и ошибку (`var f, err = os.Open(name)`).
* В одном объявлении можно объявить и инициализировать несколько переменных разных типов (`s, v := 1, "hello"`);