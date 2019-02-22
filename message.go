package dgohelper

import (
    "regexp"
    "strings"
    
    "github.com/bwmarrin/discordgo"
)

var (
    patternChannels        = regexp.MustCompile("<#[^>]*>")
    patternMentions        = regexp.MustCompile("<(@|!@)[^>]*>")
    patternSnowflake       = regexp.MustCompile("[0-9]{16,18}")
    patternChannelMentions = regexp.MustCompile("<@&[^>]*>")
)

func ContentWithMentionsRemoved(m *discordgo.Message) (content string) {
    content = m.Content
    
    for _, user := range m.Mentions {
        content = strings.Replace(content, "<@"+user.ID+">", "", -1)
        content = strings.Replace(content, "<@!"+user.ID+">", "", -1)
    }
    
    return strings.TrimSpace(content)
}

func GetMembersFromContent(s *discordgo.Session, guildid, msg string) (ms []*discordgo.Member) {
    uids := patternSnowflake.FindAllString(msg, -1)
    for ui := range uids {
        if m, err := s.State.Member(guildid, uids[ui]); err == nil {
            ms = append(ms, m)
        }
    }
    return ms
}

func GetChannelMention(s *discordgo.Session, msg, channelid string) *discordgo.Channel {
    if len(GetChannelMentions(s, msg)) > 0 {
        for _, ci := range GetChannelMentions(s, msg) {
            if ci.ID == channelid {
                return ci
            }
        }
    }
    return nil
}

func GetChannelMentions(s *discordgo.Session, msg string) (cm []*discordgo.Channel) {
    cms := patternChannels.FindAllString(msg, -1)
    
    for cmi := range cms {
        if c, err := s.State.Channel(cms[cmi]); err == nil {
            cm = append(cm, c)
        }
    }
    return cm
}