# Ode
Ode is a Discord bot. This project houses all the code responsible for interacting with Discord.

## Contributing

1. Install Go 1.11 or higher
2. `go get -u github.com/rmwiesenberg/ode`

## Running

The easiest way to run Ode is to use Docker Compose.

3. `docker-compose build`
4. `docker-compose up`

## Add to Discord Server
[Add to Server](https://discord.com/oauth2/authorize?client_id=1140718447018377417&scope=bot&permissions=40138682993408)

## Long-term deployment
1. install docker, add run user to the docker group, restart
2. edit `ode.service` and add .env file with environment variables
3. `sudo cp ode.service /etc/systemd/system`
4. `sudo systemctl daemon-reload`
5. `sudo systemctl start ode`
6. `sudo systemctl enable ode`
7. 