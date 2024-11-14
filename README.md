<p align="center">
  <img src="https://github.com/k5sha/lifeEasier/blob/master/media/logo.png" alt="Logo" width="256"/>
</p>

<div align="center">
  
  ![Go Badge](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
  ![Telegram Badge](https://img.shields.io/badge/Telegram-bot-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)
  ![Docker Badge](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
</div >

# Life_Easier

A Telegram bot built in Golang to help people remember to do their useful things

#### You can try it [here](https://t.me/Life_Easier_bot)

### How it works:
If you find something interesting on the internet or want to do something in the future, simply send the link or a message to the bot. It will store your request and send you the link or reminder at a later time, when you need it most.
<p align="center">
  <img src="https://github.com/k5sha/lifeEasier/blob/master/media/how.jpg" alt="How work" width="726"/>
</p>

### Database schema
<br>
<p align="center">
  <img src="https://github.com/k5sha/lifeEasier/blob/master/media/db.svg" alt="db" width="256"/>
</p>

### Config

- *config.yaml*
```yaml
telegram_bot_token: 'YOUR_TOKEN'
database_dsn: 'postgres://postgres:postgres@localhost:5432/your_db?sslmode=disable"'
send_interval: '24h'
min_hours_random: 24
max_hours_random: 72
```
- *env*
```env
TELEGRAM_BOT_TOKEN=<YOUR_TOKEN>
DATABASE_DSN=postgres://postgres:postgres@localhost:5432/your_db?sslmode=disable"
SEND_INTERVAL=24h
MIN_HOURS_RANDOM=24
MIN_HOURS_RANDOM=72
```
### Startup

- via Docker
```bash
docker-compose -f docker-compose.dev.yml up -d
```

### Nice to have features (backlog)
- [ ]  Use webhook for better performance
- [ ]  Implement  reminders with custom intervals (e.g., daily, weekly)
- [ ]  Ability to add notes or tags to each reminder 
- [x]  Add beautiful message with a markdown 
- [ ]  Ability to send multimedia reminders (e.g., images, videos, audio)
- [ ]  Another type of SQL or NOSQL db (mongo, mysql)
- [ ]  Command handler
- [ ]  Add settings where you can change interval when the bot can send you message
- [ ]  Add tests
- [ ]  Video guide 
### Author:
**Yurii (k5sha) Yevtushenko**
