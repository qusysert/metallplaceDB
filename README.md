# metallplaceDB

#### Запуск
В корне проекта создать файл app.env со следующим содержимым

    DB_HOST=[хост БД]
    DB_PORT=[БД]
    DB_USER=[имя пользователя для подключения к БД]
    DB_PASSWORD=[пароль к БД]
    DB_NAME=[имя БД]
    HTTP_PORT=[порт сервера]
    
Далее запустить файл cmd/metallplace/main.go

Готово, сервер запущен и слушает запросы!

#### Сервер

Сервер реализует следующие хендлеры:
- getMaterials - возвращает существующие уникальные связки Материал - Источник - Рынок - Валюта продажи в формате JSON
- getPrice - возвращает фид цены определенной связки за определенный промежуток. На вход принимает id связки, дату нижней и верхней границы (несторогое сравнение)
 
  POST localhost:8080/getPrice
  Content-Type: application/json

  {
    "material_source_id": 1,
    "start": "2017-01-04",
    "finish": "2017-04-05"
  }
