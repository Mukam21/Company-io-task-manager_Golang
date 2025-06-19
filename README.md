# IO-Bound Task Manager API

**I/O-bound Task Manager** — это HTTP API на Go, которое позволяет создавать, отслеживать и удалять задачи с длительной обработкой (от 3 до 5 минут). 

Все данные хранятся **в оперативной памяти** (in-memory), без сторонних баз данных, очередей и хранилищ.

##  Возможности

- Создание задач

- Получение всех задач

- Получение задачи по ID

- Удаление задачи

- Статусы: `pending (в ожидании)`, `processing (в процессе)`, `completed (завершено)`, `failed (ошибка)`

- Вся логика реализована in-memory

##  Требования

- **Go 1.24.4+**  
  > Важно: проект использует функции из Go 1.21+ и протестирован на Go 1.24.4. Старые версии Go могут не поддерживать часть кода.

- [Docker (опционально)](https://docs.docker.com/get-docker/) — для запуска без установки Go

##  Установка и запуск вручную (без Docker)

```bash
# Клонируем проект

  git clone https://github.com/Mukam21/Company-io-task-manager_Golang.git

  cd Company-io-task-manager_Golang

# Устанавливаем зависимости
  go mod tidy

# Запускаем сервер
  go run ./cmd/server/main.go

После запуска сервер будет доступен на:

  http://localhost:8080

Запуск с помощью Docker:

  docker build -t io-bound-task-api .

  docker run -p 8080:8080 io-bound-task-api

Примеры HTTP-запросов:

Если вы используете Postman, вы можете выполнять эти запросы вручную или импортировать коллекцию.

Ниже приведены примеры curl для справки.

1. Создание задачи:

  curl -X POST http://localhost:8080/tasks

  Postman:

  Метод: POST

  URL: http://localhost:8080/tasks

2. Получение всех задач:

  curl http://localhost:8080/tasks

  Postman:

  Метод: GET

  URL: http://localhost:8080/tasks

3. Получение задачи по ID:

  curl http://localhost:8080/tasks/{task_id}

  Postman:

  Метод: GET

  URL: http://localhost:8080/tasks/<вставьте ID задачи>

4. Удаление задачи:

  curl -X DELETE http://localhost:8080/tasks/{task_id}

  Postman:

  Метод: DELETE

  URL: http://localhost:8080/tasks/<вставьте ID задачи>

Тестирование:

Проект включает модульные тесты:

  go test ./... -v

Вывод будет содержать:

  === RUN   TestTaskService_ProcessTask
  --- PASS: TestTaskService_ProcessTask (0.20s)
  PASS
  ok      github.com/Mukam21/Company-io-task-manager_Golang/pkg/service

Структура проекта:

├── cmd/
│   └── server/         # Точка входа (main.go)
├── pkg/
│   ├── entity/         # Модель Task
│   ├── handlers/       # In-memory хранилище
│   ├── router/         # HTTP маршруты
│   ├── service/        # Логика обработки задач и воркеры
│   └── transport/
│       └── http_trans/ # HTTP хендлеры
├── Dockerfile          # Контейнеризация
├── go.mod              # Go зависимости
├── README.md           # Инструкция


Автор:

Golang-разработчик — Мукам Усманов
