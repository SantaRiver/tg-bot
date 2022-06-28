### Запуск PostgresSQL в докере
    docker run --name workshop-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:14.2
### Подключение к БД через psql и наполнение ии тестовыми данными
    psql -h 127.0.0.1 -p 5432 -U postgres -d postgres
    psql -h 127.0.0.1 -p 5432 -U postgres -f ~/Desktop/demo_big.sql
### Генерация заглушки на БД
    minimock -i github.com/mpak/testgrpc/internal/app.Repository -o ./repository_mock.go
### Нагрузочное тестрирование с ghz
    ghz --insecure --proto ./api/api.proto --call api.SomeStuff.GetAircraft -d '{"code": "773"}' localhost:8080