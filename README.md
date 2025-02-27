# Main-Journal

## Идеи на доработку

### Мелкие
- ВЫТАЩИТЬ ИЗ ФИКС-АКТА функции работы с базой данных и скинуть их в МОДЕЛЬ
- инициализировать базу данных в main
    - убрать инициализацию базы данных из fixact.NewApp()
- СОЗДАТЬ папку КОНТРОЛЛЕР (И ОБДУМЫВАТЬ ТО, как он будет взаимодействовать с моделью и вью)


### Глобальные
В общем и целом есть идея построить архитектуру приложения по типу MVC (model-view-controller). 
Есть конечно гигантская вероятность того, что я вообще не понимаю, что такое MVC, но считаю, что к данному
проекту я должен относиться как к тренировке (обучение программированию), поэтому в принципе даже если я сделаю
что-то ужасное и дерьмовое, то в конечном итоге у меня будет возможность классно поржать над самим собой в будущем.


Приложение нужно написать так, чтобы иметь возможность в дальнейшем отвалиться от библиотеки fyne и использовать другой gui
или вообще CLI (шутка).

В MVC model не общается напрямую с view? Значит вызов ф-ии создания ДБ напрямую из приложения fyne - НЕ ХОРОШО! ПЕРЕПИСАТЬ!

## Мои заметки о программировании (не смейтесь)

Суть написания хорошого приложения скорее всего можно позаимствовать с машиностроения. Автомобиль как конечный продукт можно сравнить
с приложением. Автомобиль состоит из узлов (трансмиссия, подвеска, кузов), которые в свою очередь состоят из более локальных (мелких)
деталей, которые возможно состоят из еще более мелких деталей. Как правило есть более-менее универсальные детали, например
колеса (которые подойдут как на Жигули, так и на мерседес) или детали специфические (которые подходят только определенной модели).

В программировании так-же. Мы разбиваем приложение на глобальные модули. Те в свою очередь состоят из более мелких, которые состоят из
еще более мелких и так-далее.

## Версии

### 0.0.10 (27.02.2025)
Мысль... В следующей версии необходимо пересмотреть структуру проекта и отказаться от модели MVC. Нужна кастомная архитектура. И вообще уж больно сильно сейчас всё завязано на отображении в библиотеке fyne. Поэтому пытаться реализовать какую-то супер-пупер модель, в которой всё можно выдернуть и заменить на всё - считаю не состоятельным.

Вообще стоит относиться к этому приложению как к учебному проекта (с персональной заинтересованностью)