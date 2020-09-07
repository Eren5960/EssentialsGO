package commands
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
)

type DefaultGameMode struct {
	GameMode mode // @see Gamemode.go for "mode"
	// Target   target will be soon
}

func (t DefaultGameMode) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)
	mode := StringToGameMode(string(t.GameMode))

	p.World().SetDefaultGameMode(mode)
	output.Printf("The world's default game mode is now %s.", GameModeToName(mode))
}