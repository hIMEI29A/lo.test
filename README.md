# lo.test

Тестовое задание

## Build & run

go run main.go

## Task

```
type Task struct {
	Id          int // Обязательное поле. Устанавливается автоматически при помещении в хранилище.
	Status      Status // Обязательное поле. Cм. соотв. раздел
	Description string // Обязательное поле.
}
```

## Usage

```bash
# Создание задачи
curl -X POST -H "Content-Type: application/json" -d '{
    "status": "new",
    "description": "test task"
}' http://localhost:8080/tasks

# Получение всех задач
curl http://localhost:8080/tasks

# Фильтрация по статусу
curl http://localhost:8080/tasks?status=new

В случае, если параметр `status` не передан, вернется список со всеми имеющимися задачами.

# Получение задачи по ID
curl http://localhost:8080/tasks/1
```

### Available statuses
	"new"
	"complete"
	"pending"

При попытке создать задачу с любым другим статусом будет возвращена ошибка `"invalid status"`