# Go HTTP Server

![Go Version](https://img.shields.io/github/go-mod/go-version/golang/go)
![License](https://img.shields.io/github/license/golang/go)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen)

## Описание

Этот проект — простой HTTP сервер, написанный на языке Go. Он обрабатывает различные маршруты и возвращает простые текстовые ответы. Это учебный проект, который будет расширяться ежедневно, добавляя новый функционал и возможности.

## Текущий функционал

- **Маршрут "/"**: Возвращает строку "Main Handler".
- **Маршрут "/about"**: Возвращает строку "About Handler".
- Чтение переменных окружения из файла `.env` (порт сервера).

## Используемые технологии

- [Go](https://golang.org/) — основной язык программирования.
- [godotenv](https://github.com/joho/godotenv) — для работы с переменными окружения из файла `.env`.

## Установка и запуск проекта

### 1. Клонирование репозитория

Сначала клонируйте репозиторий на ваш локальный компьютер:

```bash
git clone https://github.com/yourusername/your-repo-name.git
cd your-repo-name
```
