package events

import (
	"context"
	"encoding/json"
	"fmt"
	"misha/extensions"
	"misha/models"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func GuildCreate(c *extensions.Ex, s *discordgo.Session, e *discordgo.Event) {

	res := &discordgo.GuildCreate{}
	buf, _ := e.RawData.MarshalJSON()
	json.Unmarshal(buf, res)
	var result models.GuildDoc
	err := c.DB.Colls["guilds"].FindOne(context.TODO(), bson.M{"GuildID": res.Guild.ID}).Decode(&result)
	if err != nil {
		c.DB.Colls["guilds"].InsertOne(context.TODO(), bson.M{"GuildID": res.Guild.ID})
	}

}
func GuildDelete(s *discordgo.Session, e *discordgo.Event) {
	res := &discordgo.GuildDelete{}
	buf, _ := e.RawData.MarshalJSON()
	json.Unmarshal(buf, res)
	fmt.Println(res.Guild.Name)
}
