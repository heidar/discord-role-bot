discord-role-bot:
	go build ./cmd/discord-role-bot

clean:
	rm discord-role-bot

install:
	install -m 755 discord-role-bot /usr/local/bin

uninstall:
	rm /usr/local/bin/discord-role-bot

