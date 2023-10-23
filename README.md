# Ode
Ode is a Discord bot. This project houses all the code responsible for interacting with Discord. \
Ode can do a lot of things!
- Pomodoro
- Animal pictures (dogs and cats)
- and more! 

[Add to Server](https://discord.com/oauth2/authorize?client_id=1140718447018377417&scope=bot&permissions=40138682993408)

## Long-term deployment
1. install docker, add run user to the docker group, restart
2. Clone repo (ideally into dedicated user "ode", otherwise need to edit the service file)
3. `sudo cp ode.service /etc/systemd/system`
4. `sudo systemctl daemon-reload`
5. `sudo systemctl edit ode` to add the environment variables
```
[Service]
Environment="DISCORD_BOT_TOKEN=..."
Environment="CATS_API_TOKEN=..."
Environment="DICTIONARY_API_TOKEN=..."
```
6. `sudo systemctl start ode`
7. `sudo systemctl enable ode`
