# Параллельная выборка URL
* `Горутина` - функция, которая может выполнятся параллельно;
* `Канал` - механизм связи, который позволяет одной горутине передавать значения определенного типа другой горутине. Функция main выполняется в горутине а инструкция go создает дополнительные горутины;
* Функция `io.Copy()` считывает тело ответа  и игнориурет его, записывая в выходой поток `io.Discard`.
`Copy` возвращаетколичество байтов и информацию о произошедших ошибках;
* Когда одна горутина пытается отправить или получить информацию по каналу, она блокируется, пока другая горутина пытается выполнить те же действия. После передачи информации обе горутины продолжают работу.
* Весь вывод осуществляется функцией main, что гарантирует, что вывод каждой горутины будет обработан как единое целое без чередования вывода при завершени двух горутин в один и тот же момент времени. 