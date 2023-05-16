# Серверная часть MusicDay

## Курсового проекта "Мобильное приложение-аудиотека для обмена музыкаль- ными предпочтениями «MusicDay»"
## Выполнил: Киселев Иван БПИ217

### Установка и запуск программы
Данный репозиторий представляет из себя серверную часть приложения,
предполгается его запускать на VPS (виртуальном сервере),
в нашем же случае уже запущен, и к нему можно обратиться по адресу
http://134.0.119.220:8000/

В случае, если требуется установить данный програмный продукт.
Зайдя в данную директорию нужно прописать
```docker run --name=musicday_db -e POSTGRES_PASSWORD=password  -p 5432:5432 -d postgres```,
чтобы запустить в docker контейнер postgres, затем нужно запустить саму программу командой ```./build.exe```

Для запуска программы требуется наличие docker и postgreSQL (для помещения postgreSQL в docker надо прописть```docker pull postgres```).
### Описание репозитория
Данный репозиторий представляет из себя серверную часть приложения MusicDay


## Строение репозитория:

### Точка входа в программу и инициализация конфигураций
**[main.go](https://github.com/musicday/server/tree/main/cmd/app)**
> В данном файле запускается вся прогамма:
> 
> Инициализируется logrus (библеотека логирования)
> 
> Подключаеся БД (инициализируется СУБД sqlx вместе с драйвером)
> 
> Инициализируются все три слоя чистой арихтектуры ссылаясь 
> друг на друга
```go
repos := repository.New(db)
services := service.NewService(repos)
handlers := transport.NewHandler(services)
```
> Запускается сервер на 8000 порту
