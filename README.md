<h1 align=center><code>sedcord-go</code></h1>
<p align=center>Edit messages without editing messages</p>

## 💡About
A simple bot to run sed at Discord messages.

## 📥 Installation

### Heroku
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/RedsonBr140/sedcord-go)

### Local/VPS
```sh
git clone https://github.com/RedsonBr140/sedcord-go.git
cd sedcord-go
cp .env.sample .env
# Edit the .env file
go run main.go
```
### Add sedcord to your server
Don't want to self-host the bot? Ok, you can add it to you server by clicking [here](https://discord.com/api/oauth2/authorize?client_id=982312252793307196&permissions=8&scope=bot)

As the bot is hosted on Heroku, it is offline between `00:00 -0300` and `06:00 -0300`

## ⌨ Usage
Reply to a message with `!s <your regex>` for example: `!s s/game/half-life/`

## 💌 License
MIT License

---
> ❤️ Keep It Simple, Stupid
