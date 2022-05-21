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


## Configuration

**All configuration must be set as environment variables**  
In development (default), it will look for a `.env` file next to the binary/entrypoint `./cmd/parser.go`.

| **Configuration Key** | **Type**   | **Default**        | **Description**                                                                                                     |
|-----------------------|------------|--------------------|---------------------------------------------------------------------------------------------------------------------|
| ENV                   | `string`   | `development`      |                                                                                                                     |
| PORT                  | `int`      | `3000`             | Webhook Server port to listen on                                                                                    |
| TG_BOT_TOKEN          | `string`   |                    | (REQUIRED) Telegram Bot Token                                                                                       |
| TG_WEBHOOK_URL        | `string`   |                    | (REQUIRED) Webhook server URL for your bot to send updates to                                                       |
| ALLOWED_PRIVATE_USERS | `[]string` | `[]`               | List of users (comma-separated) to allow handling messages from private chat with bot. <br>Example: `user_a,user_b` |
| DB_HOST               | `string`   | `localhost`        |                                                                                                                     |
| DB_PORT               | `int`      | `5432`             |                                                                                                                     |
| DB_USERNAME           | `string`   | `postgres`         |                                                                                                                     |
| DB_PASSWORD           | `string`   | `postgrespassword` |                                                                                                                     |
| DB_DATABASE           | `string`   | `postgres`         |                                                                                                                     |
| DB_UNIX_SOCKET        | `string`   | ``                 | Used to connect to GCP Cloud SQL using Unix Socket                                                                  |

## Parsing Messages
### Expenses

Example message:  
```
200
shopping
cash
```

Required data:  
**ON SEPARATE LINES! In any order**
* Amount (float64)
* Category
* Payment Method
