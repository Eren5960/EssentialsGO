# Essentials GO
Essential commands for [Dragonfly](https://github.com/df-mc/dragonfly).

## Ready for:
- [x] /help - Show server commands and descriptions.
- [x] /gamemode - Changes the player to a specific game mode.
- [x] /teleport - Teleport everywhere.
- [x] /defaultgamemode - Set the default gamemode.
- [x] /setworldspawn - Sets a worlds's spawn point.
- [x] /xyz - Show/hide coordinates.
- [x] /op - Give op permissions to player.
- [x] /deop - Take op permissions from player.
- [x] /stop - Stop the server from in game.
- [x] /time - Changes or queries the world's game time.

## Usage
### Get the package
`go get -u github.com/eren5960/essentialsgo`
### Import package
```go 
import "github.com/eren5960/essentialsgo"
```
### Load commands
```go
essentialsgo.LoadCommands(server, nil) // the server is *dragonfly.Server{}, nil for load all commands
```
#### Load commands without some commands
```go
essentialsgo.LoadCommands(server, []string{"stop", "defaultgamemode"}) // All will be loaded except "stop" and "defaultgamemode" command
```
#### Finis
If you want to support the project, you can give a star, report errors and bugs as issues.

# VIP
## Simple Console Command Sender
```go
essentialsgo.LoadConsole()
```
![](https://user-images.githubusercontent.com/35738714/92475213-b255e580-f1e5-11ea-9e15-c5cbfc71e98e.gif)
