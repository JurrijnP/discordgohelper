package dgohelper

import (
    "regexp"
    "strings"
    
    "github.com/bwmarrin/discordgo"
)

var (
    patternChannels        = regexp.MustCompile("<#[^>]*>")
    patternMentions        = regexp.MustCompile("<(@|!@)[^>]*>")
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

func GetUserIDs(s *discordgo.Session, msg string) (uids []string) {
    return uids
}

func GetChannelMention(s *discordgo.Session, msg string) (cm string) {
    if len(GetAllChannelMention(s, msg)) > 0 {
        return GetAllChannelMention(s, msg)[0]
    }
    return msg
}

func GetAllChannelMention(s *discordgo.Session, msg string) (cm []string) {
    cm = patternChannels.FindAllString(msg, -1)
    
    for cmi := range cm {
        cm[cmi] = patternChannels.ReplaceAllStringFunc(cm[cmi], func(mention string) string {
            channel, err := s.State.Channel(mention[2 : len(mention)-1])
            if err != nil || channel.Type == discordgo.ChannelTypeGuildVoice {
                return mention
            }

            return channel.ID
        })
    }
    return cm
}