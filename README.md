# discord-role-bot
A really simple bot to manage roles on Discord based on reactions to messages.

## Dependencies
The only dependency is Go. This has only been tested on Go 1.18.

## Building
Clone the repository and build the project.

```
make
```

This should give you an executable in the top level folder called
`discord-role-bot`.

## Installing
The bot is by default installed to `/usr/local/bin`.

```
make install
```

## Configuration
Copy `configs/config.json.example` to `configs/config.json`. If you plan on
running this on a server as a service, it will look for configs in
`/etc/discord-role-bot/config.json` as well, if it fails to find it in the
`configs` folder.

```
cp configs/config.json.example configs/config.json
```

Go into Discord settings -> Advanced and turn on Developer Mode.

Then create a discord bot in the [Discord Developer Portal](https://discord.com/developers/applications).
If you are unsure of how to do this, there are plenty of tutorials on how to create a bot with Discord.

You need to invite the bot to your server. It needs two permissions:

* Manage Roles
* Read Messages/View Channels

The permission integer should be `268436480`.

Copy the token for your bot and paste it into `discordToken` in `config.json`.

```
{
  "discordToken": "your-token",
  ...
}
```

Then right click on your server in Discord and copy its ID, paste it into
`guildID` in `config.json`.

```
{
  ...
  "guildID": "your-guild-id",
  ...
}
```

Next up write a message in a channel on your Discord server to use for
reactions, then copy its ID and paste it into `config.json`.

```
{
  ...
  "roleConfig": {
    "your-message-id": {
  ...
}
```

Create a role on your Discord server that you want to be added when a user
reacts to your message. Once you have created the role, copy its ID and paste
it into `config.json`.

```
{
  ...
  "roleConfig": {
    "your-message-id": {
      "your-emoji": "your-role-id"
  ...
}
```

Finally, pick an emoji and paste it into `config.json`. This should be the
actual emoji you want to use, for example 📢.

## Running
Make sure you've built the bot and configured it, then run the executable.

```
./discord-role-bot
```

The bot doesn't fork into the background, it just runs and logs in the
foreground and needs to be backgrounded if desired.

## Logging
Logs get written to stdout. If you have issues, check the logs first.

## Uninstalling
Use make to uninstall the bot.

```
make uninstall
```

## FAQ
The bot has all the permissions but still can't add my role, how come?
Make sure the bot's role is higher up in the role list than the users it
will be modifying.

## Reporting issues
I don't have a public issue tracker. My email is in the commit history, feel
free to drop me an email if you have problems, I'm happy to help.

## Thanks
Loosely based on [Distortions81/M45-ReactBot](https://github.com/Distortions81/M45-ReactBot).
