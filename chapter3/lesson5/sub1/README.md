* Строковые литералы в Go - это последовательность байтов, заключенная в двойные кавычки, позволяют записывать      текстовые значения, включая символы `Unicode`;
* Управляющие последовательности, начинающиеся с обратной косой черты `'\'`, используются для вставки произвольных
  значений байтов в строке, обрабатывают управляющие коды ASCII и специальные символы;
* Шестнадцатеричные управляющие последовательности записываются как `\xhh` и представляют один байт с указанным
  значением, ровно с двумя шестнадцатеричными цифрами `h` (в верхнем или нижнем регистре);
* Восьмеричные управляющие последовательности записываются как `\ooo` и тоже представляют один байт с указанным
  значением, с ровно тремя восьмеричными цифрами (от 0 до 7), не превышающими значение `\377`;
* Неформатированные строковые литералы (`raw string literal`) используют обратные одинарные кавычки `` и позволяют
  записывать многострочный текст, в котором управляющие последовательности не обрабатываются;
* При использовании неформатированных строковых литералов символы возврата каретки `(\r)` удаляются, чтобы значение
  строки было одинаковым на всех платформах;
* Неформатированные строковые литералы удобно использовать для записи регулярных выражений, шаблонов HTML, литералов
  JSON, сообщений об использовании программы и других текстов, занимающих несколько строк.