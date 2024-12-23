### Обзор проекта Calculator-API

**Название:** Calculator-API

**Язык программирования:** Go (Golang)

**Тип проекта:** RESTful API для вычислений математических выражений.

---

#### Описание проекта

Calculator-API — это веб-приложение, реализующее RESTful API, которое позволяет пользователям выполнять математические вычисления, отправляя выражения в виде JSON. API поддерживает основные арифметические операции, такие как сложение, вычитание, умножение и деление. Это удобный инструмент для разработчиков, которым необходимо выполнять вычисления динамически и программно.

#### Основные функции

- Обработка запросов на вычисление арифметических выражений.
- Поддержка базовых операций: `+`, `-`, `*`, `/`.
- Возможность обработки выражений в инфиксной нотации с преобразованием в постфиксную нотацию для удобного вычисления.
- Ответ в формате JSON с результатом вычисления.

#### Архитектура и компоненты

- **HTTP-сервер:** Использует библиотеку `net/http` для реализации веб-сервера и обработки HTTP-запросов.
- **Декодирование JSON:** Обрабатывает входящие данные в формате JSON через `encoding/json`.
- **Валидация входных данных:** Проверяет корректность переданных математических выражений.
- **Логика вычислений:** Включает функции для токенизации, преобразования инфиксных выражений в постфиксные и вычисления результата.

#### Как пользоваться проектом

1. **Запуск сервера:**
   - Чтобы запустить сервер, выполните команду в терминале, находясь в директории проекта:
     ```bash
     go run ./cmd/app
     ```
   - Сервер будет запущен и доступен на `http://localhost:8080`.

2. **Отправка запроса на вычисление:**
   - Для выполнения вычисления необходимо отправить POST-запрос на конечную точку `/api/v1/calculate` с телом запроса в формате JSON. Воспользуйтесь для этого терминалом Git Bash. Например, используя `curl`:
     ```bash
     curl -X POST http://localhost:8080/api/v1/calculate -d '{"expression": "3 + 5 * 2"}' -H "Content-Type: application/json"
     ```
   - В данном примере выражение `"3 + 5 * 2"` будет вычислено.

3. **Формат ответа:**
   - Если запрос успешно обработан, ответ будет возвращен в формате JSON, содержащем результат вычисления. Например:
     ```json
     {
       "result": "13"
     }
     ```
   - В случае ошибки (например, некорректного выражения) сервер вернет соответствующий код состояния HTTP и сообщение об ошибке.

#### Примеры запросов представлены в полной документации docs. Запросы необходимо отправлять через консоль Git Bash.
     
     https://github.com/butorovv/Calculator-API/blob/master/Calculate-service/docs/swagger.yaml

#### Замечания
- API не поддерживает сложные математические выражения (например, функции или тригонометрические операции) — только базовые арифметические операции.
---

Calculator-API представляет собой простой и удобный инструмент для вычислений, который может быть интегрирован в более крупные проекты или использован изолированно для выполнения математических операций.
