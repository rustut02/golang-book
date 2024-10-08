# 3.5.3

* `UTF-8` - это кодировка переменной длины символов `Unicode` в виде байтов, изобретенная Кеном Томпсоном и Робом
  Пайком, является стандартом `Unicode`;
* Руны в `UTF-8` могут использовать `от 1 до 4 байтов`, но символы `ASCII` занимают только `1 байт`, большинство
  распространенных рун используют `2` или `3 байта`;
* `UTF-8` обладает свойством самосинхронизации (можно найти начало символа, просмотрев не более 3 байт) и совместимости с `ASCII`, что упрощает работу с текстом;
* Для работы с отдельными рунами в Go можно использовать пакеты `unicode` и `unicode/utf8`, которые предоставляют
  функции для кодирования, декодирования и обработки рун;
* Управляющие последовательности `Unicode` в строковых литералах Go позволяют указывать символы с помощью их числового кода: `'\uhhhh'` для 16-разрядных значений и `'\uhhhhhhhh'` для 32-разрядных;
* `UTF-8` позволяет проводить некоторые строковые операции без декодирования, например, проверить является ли одна
  строка префиксом, суффиксом или содержит ли подстроку другой строки;
* Если необходимо работать с отдельными символами `Unicode`, нужно использовать другие механизмы, такие как функция
  `utf8.RuneCountInString()` для определения количества рун в строке;
* Для работы с символами `Unicode` в Go необходимо декодирование `UTF-8`, для этого можно использовать
  пакет `unicode/utf8`;
* При декодировании с помощью функции `utf8.DecodeRuneInString()` возвращается руна и количество байтов, занятых её
  кодом `UTF-8`;
* Цикл по диапазону Go, применённый к строке, выполняет декодирование `UTF-8` неявно (каждый элемент строки,
  рассматривается как `rune`);
* Чтобы посчитать количество рун в строке, можно использовать простой цикл `range` или
  функцию `utf8.RuneCountInString(s)`;
* Тексты строк в Go интерпретируются как последовательности символов `Unicode` в кодировке `UTF-8`, это необходимо для правильного использования циклов по диапазону;
* Если в строке содержатся ошибки кодировки, то при декодировании генерируется специальный замещающий символ Unicode `'\uFFFD'`;
* Преобразование `[]rune` примененное к строке с кодировкой `UTF-8` возвращает последовательность символов `Unicode`, закодированных в этой строке;
* Преобразование целочисленного значения в строку рассматривает это число как значение руны и даёт её представление в кодировке UTF-8;
* Если руна некорректна, вместо неё подставляется замещающий символ `�`.