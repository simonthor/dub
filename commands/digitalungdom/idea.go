package digitalungdom

import (
	"context"
	"strings"

	"github.com/digitalungdom-se/dub/pkg"
	"go.mongodb.org/mongo-driver/bson"
)

var Idea = pkg.Command{
	Name:        "idea",
	Description: "Föreslå något till Digital Ungdom",
	Aliases:     []string{"förslag"},
	Group:       "digitalungdom",
	Usage:       "idea <idea>",
	Example:     "idea skaffa programerings tutorials",
	ServerOnly:  false,
	AdminOnly:   false,

	Execute: func(ctx *pkg.Context) error {
		if len(ctx.Args) < 3 {
			ctx.Reply("Du måste ge ett förslag")
			return nil
		}

		idea := bson.M{
			"type":    "idea",
			"where":   "discord",
			"message": strings.Join(ctx.Args, " "),
			"author":  ctx.Message.Author.ID}

		ctx.Server.Database.Collection("notifications").InsertOne(context.TODO(), idea)

		ctx.Reply("Du har nu skickat in ditt förslag. Tack för din medverkan!")

		return nil
	},
}