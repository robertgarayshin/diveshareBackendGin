# Driveshare Backend (КЗ 5 ПВП2023)

Серверная часть приложения Driveshare разработана с учетом современных требований к быстродействию информационных систем.
Технологии, применяемые при разработке доказывают свою эффективность в крупнейших российских и зарубежных компаниях:
- Go - высокопроизводительный язык, рассчитанный на HighLoad системы
- PostgreSQL - СУБД, способная обрабатывать несколько тысяч запросов в секунду
- Фреймворк Gin-Gonic/Gin, отличающийся высокой производительностью и простотой разработки
- SQLX - надстройка для встроенной библиотеки для обработки SQL-запросов
- Docker и Docker Compose для удобной развертки системы
- GitHub Actions для развертки CI (со встроенными go vet и golint)  
Проект разработан по технологии чистой архитектуры и требует минимальных усилий для начала работы:  

```
cd <директория проекта>  
docker-compose build
docker compose up -d
```

