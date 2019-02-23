package dgohelper

import (
    "regexp"
    "strings"
    
    "github.com/bwmarrin/discordgo"
)

var (
    patternChannels        = regexp.MustCompile("<#([^>]*)>")
    patternMentions        = regexp.MustCompile("<(@|!@)[^>]*>")
    patternSnowflake       = regexp.MustCompile("[0-9]{16,18}")
    patternChannelMentions = regexp.MustCompile("<@&[^>]*>")
)

// ContentWithMentionsRemoved returns the content with all user mentions removed.
// m   : discordgo.Message struct.
func ContentWithMentionsRemoved(m *discordgo.Message) (content string) {
    content = m.Content
    
    for _, user := range m.Mentions {
        content = strings.Replace(content, "<@"+user.ID+">", "", -1)
        content = strings.Replace(content, "<@!"+user.ID+">", "", -1)
    }
    
    return strings.TrimSpace(content)
}

// GetMembersFromContent searches for all user snowflakes in the message content,
// returns slice of discordgo.Member.
// s         : discordgo.Session struct.
// guildid   : The guildID.
// msg       : content of the message.
func GetMembersFromContent(s *discordgo.Session, guildID, content string) (ms []*discordgo.Member) {
    uids := patternSnowflake.FindAllString(content, -1)
    for ui := range uids {
        if m, err := s.State.Member(guildID, uids[ui]); err == nil {
            ms = append(ms, m)
        }
    }
    return ms
}

// GetChannelMention searches for a specific channel mention in the message content,
// returns discordgo.Channel if channel is found.
// s           : discordgo.Session struct.
// channelID   : The ChannelID of the channel you want.
// msg         : content of the message.
func GetChannelMention(s *discordgo.Session, msg, channelID string) *discordgo.Channel {
    if len(GetChannelMentions(s, msg)) > 0 {
        for _, ci := range GetChannelMentions(s, msg) {
            if ci.ID == channelID {
                return ci
            }
        }
    }
    return nil
}

// GetChannelMentions searches for all channel mention in the message content,
// returns slice of discordgo.Channel for channels found.
// s           : discordgo.Session struct.
// msg         : content of the message.
func GetChannelMentions(s *discordgo.Session, msg string) (cm []*discordgo.Channel) {
    cms := patternChannels.FindAllString(msg, -1)
    
    for cmi := range cms {
        if c, err := s.State.Channel(patternChannels.ReplaceAllString(cms[cmi], "$1")); err == nil {
            cm = append(cm, c)
        }
    }
    return cm
}