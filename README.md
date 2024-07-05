## Веб-сервис парсер


### Описание

Реализует точки доступа через `GET`- запросы, для кодирования и декодирования `base64` и `jwt`

```sh
http://localhost:8060
```
Endpoints, соответственно:
* `/encode?secret=` -- енкодер base64
* `/decode?secret=` -- декодер base64
* `/jwtdecode?jwt=` -- декодер jwt


### Сборка и использование


```sh
go build .
./webparser
```

Запросы из браузера:

```sh
http://localhost:8060/encode?secret=эту_строку_хочу_закодировать_в_base64
```

```sh
http://localhost:8060/decode?secret=эту_строку_хочу_декодировать_из_base64
```

```sh
http://localhost:8060/jwtdecode?jwt=этот_токен_хочу_декодировать
```

