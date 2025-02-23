# project-lite-go

Discord CLI client but better and written in Go (inspired by [project-lite](https://github.com/tjf1dev/project-lite))

> [!NOTE]
> The client doesn't send your token to anyone othen then to Discord servers.
## is it against TOS of Discord?
> This client uses the Discord API to perform user actions. Discord technically does not allow 3rd party clients, but the chance of getting banned is low, though use the client at your own risk.

## Progress (these are test fields, the progress isn't for now accurate)
- Messaging
  - [x] Sending text
  - [x] Recieving text
  - [x] Format mentions
  - [ ] Sending attachments
  - [ ] Recieving attachments
  - [ ] App commands
  - ~~[ ] Emojis~~
  - ~~[ ] Stickers~~
  - ~~[ ] Gifs~~
  - [ ] Polls
- Information
  - [x] List guilds
  - [x] List channels
  - [x] Select guilds
  - [x] Select channels
  - [x] Get username
  - [x] Get profile picture
- Voice
  - WebRTC libraries are available on Go and can be implemented here but currently not implemented.
- Other
  - [ ] Log in as bot