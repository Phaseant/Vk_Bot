package telegramEvents

const StartMessage = `
Привет! Я бот, который умеет рассказывать анекдоты, стихи, песни и отправлять картинки.
Я был создан в качестве тестового задания для компании VK.
Если тебе интересно, как я работаю, то можешь посмотреть мой код на GitHub 🌐: github.com/Phaseant/VK_Bot
Если ты хочешь узнать, что я умею, то введи /help
Сделано на Go с ❤️`

const HelpMessage = `
/start - начать работу с ботом
/help - получить справку о командах
/joke - получить анекдот
/poem - получить стихотворение
/song - получить песню
/picture - получить картинку
Также ты можешь пользоваться кнопками внизу экрана.`

const JokeMessage = `Два грузина тащат по лесу убитого медведя. Навстречу им третий.
-Вах, вах, какой большой! Гризли?
-Нэт, руками душылы.`

const PoemMessage = `Во глубине сибирских руд
Храните гордое терпенье,
Не пропадет ваш скорбный труд
И дум высокое стремленье.

Несчастью верная сестра,
Надежда в мрачном подземелье
Разбудит бодрость и веселье,
Придет желанная пора:

Любовь и дружество до вас
Дойдут сквозь мрачные затворы,
Как в ваши каторжные норы
Доходит мой свободный глас.

Оковы тяжкие падут,
Темницы рухнут — и свобода
Вас примет радостно у входа,
И братья меч вам отдадут.

А.С. Пушкин
1827`

// const SongMessage = выбор русской и английской песни

const RusSongMessage = `Когда закончатся полёты первых ласточек,
и ты усталая придёшь к себе домой.
Увидишь из окна слова из ярких лампочек,
я напишу тебе «не бойся я с тобой».
Мы можем быть только на расстоянии и невесомости,
хочешь упасть, я неволить не стану, хочешь лететь, лети.

Но я тысячу раз обрывал провода,
сам себе кричал, ухожу навсегда.
Непонятно как доживал до утра, салют Вера
Но я буду с тобой или буду один,
дальше не сбежать ближе не подойти.
Прежде чем навек поменять номера, Салют Вера.

Ты не сбываешься, хоть снишься в ночь на пятницу,
не отзываешься не на один пароль.
Не ошибаешься, и мне всё чаще кажется,
что ты посланница неведомых миров.
Мы можем быть только на расстоянии и невесомости,
хочешь упасть, я неволить не стану, хочешь лететь, лети.

Но я тысячу раз обрывал провода
сам себе кричал, ухожу навсегда
непонятно как доживал до утра, салют Вера
Но я буду с тобой или буду один
дальше не сбежать ближе не подойти
прежде чем навек поменять номера, Салют Вера!`

const EngSongMessage = `Underneath the bridge
Tarp has sprung a leak
And the animals I've trapped
Have all become my pets
And I'm living off of grass
And the drippings from my ceiling
It's okay to eat fish
Cause they don't have any feelings

Something in the way
Mmm-mmm
Something in the way, yeah
Mmm-mmm
Something in the way
Mmm-mmm
Something in the way, yeah
Mmm-mmm
Something in the way
Mmm-mmm
Something in the way, yeah
Mmm-mmm

Underneath the bridge
Tarp has sprung a leak
And the animals I've trapped
Have all become my pets
And I'm living off of grass
And the drippings from the ceiling
It's okay to eat fish
Cause they don't have any feelings

Something in the way
Mmm-mmm
Something in the way, yeah
Mmm-mmm
Something in the way
Mmm-mmm
Something in the way, yeah
Mmm-mmm
Something in the way
Mmm-mmm
Something in the way, yeah
Mmm-mmm
Something in the way
Mmm-mmm
Something in the way, yeah
Mmm-mmm`

const unknownMessage = `Я не знаю такой команды 😥
Введи /help, чтобы узнать, что я умею.`
