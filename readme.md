# API Загрузки изображений, сжатия изображений

## Логика работы сервиса
- Изображение загружается, с указанием ссылки на него
- Изображение сжимается, под нужные разрешения: 1920px, 1000px, 500px. Также изображение переводится под современные WEB форматы (webP и тд.)
- Данные об изображении заносятся в БД
- Обработанные изображения сохраняются в каталоге public по шаблону: `<ссылка, переданная при загрузке>/<разрешение>.<расширение>`
- по интервалу времени изображения удаляются из файловой системы и БД
#

## Запуск в prod версии:
- `docker-compose up --build -d`

#

## Запуск в dev версии:

- Изменить конфигурацию .env для СУБД
- В файле docker-compose.yaml закоментировать сервис API и оставить конфиг для СУБД
- Выполнить команду - `docker-compose up --build`
- Выполнить компанды:
`go mod vendor`, `swag init --parseDependency --parseInternal -g cmd/main.go`, `air`
