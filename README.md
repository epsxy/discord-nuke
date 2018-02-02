# discord-nuke

A small program to mass-delete your own messages on Discord.

Not sure why I made this honestly.

Currently it deletes every single message you've sent in every server you're currently in. If you just want to delete some
messages, either wait until I add that capability or find something else.

## Usage

```bash
go get -u github.com/lamados/discord-nuke
go install github.com/lamados/discord-nuke
discord-nuke -e <discord email> -p <discord password>
```

## Todo

- add DM conversation nuking
- allow user to specify time periods to delete in
- saving messages to a file before deleting them
- allow user to select what servers and channels to delete messages from

## Disclaimer

I take no responsibility from what happens as a consequence of using this program. This program will, provided it doesn't get
interrupted, delete every message you've sent in every server you're currently in. Be careful when using this program and make
sure that you definitely want to delete __everything__ before you run it, as it doesn't prompt you for confirmation before
starting.
