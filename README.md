# SQaLice-scanner

<a href="https://ibb.co/4YJPnGn"><img src="https://i.ibb.co/4YJPnGn/SQa-Lice-logo.png" alt="SQa-Lice-logo" border="0"></a>

Пакет сканнера SQaLice предназначен для гибкого получения данных из __sql.Rows__. В функцию-сканнер *Scan* передается модель, на которую осуществляется сканирование,
строки rows и тег, по которому подбирается целевое поле для сканирования. Поля модели, не содержащие переданного тега, пропускаются. Функция поддерживает сканирование на вложенные структуры и массивы без дополнительных условий

__UPDATE 0.1.1__
Начиная с данной версии компилятор поддерживает считывание тегов полей из вложенных структур, что делает возможным сканирование строк на содержащие их модели

Пример ошибки компилятора:

```go
"[SQaLice] Dest must be pointer to struct; got string"
```

## TO-DO

| TO-DO             | Статус        |
| ----------------- | ------------- |
| Покрытие тестами  | Не выполнено  |