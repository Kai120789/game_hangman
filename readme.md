# Виселица

## Мотивация проекта

Проект является разминочным, и его основная цель — проверить себя на то, что от теории и решения задач уже можно переходить к реализации цельных проектов. Если приложения подобного уровня вы уже реализовывали без возникновения трудностей, проект можно пропустить.

## Функционал приложения и меню консольного интерфейса

- При старте приложение предлагает начать новую игру или выйти из приложения.
- При начале новой игры случайным образом загадывается слово, и игрок начинает процесс по его отгадыванию.
- После каждой введённой буквы выводим в консоль счётчик ошибок и текущее состояние виселицы (нарисованное ASCII символами).
- По завершении игры выводим результат (победа или поражение) и возвращаемся к состоянию №1 — предложение начать новую игру или выйти из приложения.

## План работы над приложением

1. Найти в интернете словарь существительных в именительном падеже, отбросить из него слишком короткие слова. Этот словарь будет источником для выбора случайного загаданного слова для каждого раунда игры.
2. Реализовать игровой цикл отгадывания букв и отображения текущего состояния виселицы.
3. Реализовать цикл по перезапуску игры после победы/поражения.

## Чеклист для самопроверки

### Функциональные проблемы:

- Отсутствие валидации вводимых символов (валидными могут считаться, например, только маленькие буквы кириллицы). Невалидный ввод не должен увеличивать счётчик ошибок пользователя в игровом раунде.
- Повторно вводимый символ, отсутствующий в секретном слове, не должен считаться за ошибку.
