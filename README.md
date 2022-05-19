# Telebot

## First Setup
### Creating a new telegram bot

1. Open a session with [BotFather](https://telegram.me/BotFather)
2. Type `/newbot`
3. Enter display name of the bot
4. Enter username for the bot (must end with `bot`)
5. Don't forget to save the token!

### (Optional) Disable bot from joining groups

1. Open a session with [BotFather](https://telegram.me/BotFather)
2. Type `/setjoingroups`
3. Enter the name of the bot. `@example_bot`
4. Enter `Disable`

### (Optional) Disable privacy to receive all messages from group (and not only those with `/` prefix)

1. Open a session with [BotFather](https://telegram.me/BotFather)
2. Type `/setprivacy`
3. Enter the name of the bot. `@example_bot`
4. Enter `Disable`


## Required Configurations

### `.env`
* `PORT` - Webhook server port [Default: `3000`]
* `TG_BOT_TOKEN` - Telegram bot token
* `TG_WEBHOOK_URL` - Telegram bot webhook server url (this application)

## Parsing Messages
### Expenses

Example message:  
```
200
shopping
cash
```

Required data:  
* Amount (float64)
* Category
* Payment Method
