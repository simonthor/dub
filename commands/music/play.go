package music

import (
	"github.com/digitalungdom-se/dub/pkg"
)

var Play = pkg.Command{
	Name:        "play",
	Description: "Spelar en låt",
	Aliases:     []string{"spela", "pl"},
	Group:       "music",
	Usage:       "play play <youtube link>",
	Example:     "play https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	ServerOnly:  true,
	AdminOnly:   false,

	Execute: func(context *pkg.Context) error {
		context.Delete()
		if len(context.Args) == 0 {
			return nil
		}

		err := context.Server.Controller.AddToQueue(context)

		return err
	},
}
