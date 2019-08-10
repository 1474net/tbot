# tbot
golang test telegram bot

go version : **1.12.7**

## Установка 

- Сделать клон репозитория 

```bash
git clone https://github.com/1474net/tbot.git 
```

- Собрать проект со всеми зависимостями 

```
~tbot# go build
```

- Внести настройки в `config.yalm`

```yaml
#Для использование https прокси 
proxy_addr: ""
#Webhook - адрес вашего https сервера 
webhookurl: ""
#token telegram bota
token: ""
```

*В качнстве доступа к локальному компьютеру можно использовать*  ngrok 

- Запускаем исполняемый файл
- Вы восхитительны 

