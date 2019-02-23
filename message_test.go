package dgohelper

import (
    "testing"
    
    "github.com/bwmarrin/discordgo"
)

func TestContentWithMentionsRemoved(t *testing.T) {
    user1 := &discordgo.User{
        ID: "user1",
    }
    user2 := &discordgo.User{
        ID: "user2",
    }
    user3 := &discordgo.User{
        ID: "user3",
    }
    
    
    message := &discordgo.Message{
        Content: "<@user1> a <@!user2> b <@user3>",
    }
    
    message.Mentions = append(message.Mentions, user1)
    message.Mentions = append(message.Mentions, user2)
    message.Mentions = append(message.Mentions, user3)
    
    if result := ContentWithMentionsRemoved(message); result != "a  b" {
        t.Error("\""+result+"\"")
    }
}

func TestGetMembersFromContent(t *testing.T) {
    s := &discordgo.Session{
        StateEnabled: true,
        State: discordgo.NewState(),
    }
    
	user1 := &discordgo.User{
		ID: "123456789012345678",
    }
    
    user2 := &discordgo.User{
		ID: "234567890123456789",
	}
    
    s.State.GuildAdd(&discordgo.Guild{ID: "guild"})
	s.State.MemberAdd(&discordgo.Member{
		User:    user1,
		GuildID: "guild",
    })
	s.State.MemberAdd(&discordgo.Member{
		User:    user2,
		GuildID: "guild",
	})
    
    if result := GetMembersFromContent(s, "guild", "text 123456789012345678 more text 234567890123456789"); result == nil || len(result) != 2 {
        t.Error(result)
    }
}

func TestGetChannelMention(t *testing.T) {
    s := &discordgo.Session{
        StateEnabled: true,
        State: discordgo.NewState(),
    }
    
    s.State.GuildAdd(&discordgo.Guild{
        ID: "guild",
    })
    
    s.State.ChannelAdd(&discordgo.Channel{
        ID: "channel1",
        GuildID: "guild",
    })
    
    s.State.ChannelAdd(&discordgo.Channel{
        ID: "channel3",
        GuildID: "guild",
    })
    
    if result := GetChannelMention(s, "<#channel1> text <#channel3>", "channel1"); result == nil || result.ID != "channel1" {
        t.Error(result)
    }
}

func TestGetChannelMentions(t *testing.T) {
    s := &discordgo.Session{
        StateEnabled: true,
        State: discordgo.NewState(),
    }
    
    s.State.GuildAdd(&discordgo.Guild{
        ID: "guild",
    })
    
    s.State.ChannelAdd(&discordgo.Channel{
        ID: "channel1",
        GuildID: "guild",
    })
    
    s.State.ChannelAdd(&discordgo.Channel{
        ID: "channel3",
        GuildID: "guild",
    })
    
    if result := GetChannelMentions(s, "<#channel1> text <#channel3>"); result == nil || len(result) != 2 {
        t.Error(result)
    }
}