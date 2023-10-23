# Ode
Ode is a Discord bot. This project houses all the code responsible for interacting with Discord. \
Ode can do a lot of things!
- Pomodoro
- Animal pictures (dogs and cats)
- and more! 

[Add to Server](https://discord.com/oauth2/authorize?client_id=1140718447018377417&scope=bot&permissions=40138682993408)

## Contributing

1. Install Go 1.11 or higher
2. `go get -u github.com/rmwiesenberg/ode`

## Running

The easiest way to run Ode is to use Docker Compose.

1. `docker-compose build`
2. `docker-compose up`

## Long-term deployment
1. install docker, add run user to the docker group, restart
2. `sudo cp ode.service /etc/systemd/system`
3. `sudo systemctl daemon-reload`
4. `sudo systemctl edit ode` to add the environment variables
```
[Service]
Environment="DISCORD_BOT_TOKEN=..."
Environment="CATS_API_TOKEN=..."
Environment="DICTIONARY_API_TOKEN=..."
```
5. `sudo systemctl start ode`
6. `sudo systemctl enable ode`
