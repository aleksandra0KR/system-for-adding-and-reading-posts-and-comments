# System-for-adding-and-reading-posts-and-comments
Cистема для добавления и чтения постов и комментариев с использованием GraphQL

### Характеристики системы постов:
•	Можно просмотреть список постов.

•	Можно просмотреть пост и комментарии под ним.

•	Пользователь, написавший пост, может запретить оставление комментариев к своему посту.

### Характеристики системы комментариев к постам:
•	Комментарии организованы иерархически, позволяя вложенность без ограничений.

•	Длина текста комментария ограничена до, например, 2000 символов.

•	Система пагинации для получения списка комментариев.

•	Комментарии к постам доставляются асинхронно, т.е. клиенты, подписанные на определенный пост, должны получать уведомления о новых комментариях без необходимости повторного запроса.

# Запуск в Docker
Склонировать проект с гита

```
git clone https://github.com/aleksandra0KR/system-for-adding-and-reading-posts-and-comments
```
Перейти в директорию проекта
```
cd system-for-adding-and-reading-posts-and-comments
```
Забилдить
```
docker compose build
```
Запустить:
```
docker compose up
```
---
# Поменять postgres на in-memory или наоборот

  ### В файле .env поставить нужное хранилище : postgres или in-memory

 ```
STORAGE=postgres
```

![](https://github.com/aleksandra0KR/system-for-adding-and-reading-posts-and-comments/exampleImg/img1.png)

![](https://github.com/aleksandra0KR/system-for-adding-and-reading-posts-and-comments/exampleImg/img2.png)

![](https://github.com/aleksandra0KR/system-for-adding-and-reading-posts-and-comments/exampleImg/img3.png)

![](https://github.com/aleksandra0KR/system-for-adding-and-reading-posts-and-comments/exampleImg/img4.png)
