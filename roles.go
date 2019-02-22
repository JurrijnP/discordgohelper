package dgohelper

import (
    "strings"
    
    "github.com/bwmarrin/discordgo"
)

func GetRoleByName(Roles []*discordgo.Role, RoleName string) (r *discordgo.Role) {
    for ri := range Roles {
        if strings.ToLower(Roles[ri].Name) == strings.ToLower(RoleName) {
            return Roles[ri]
        }
    }
    return nil
}

func GetHighestRole(Roles []*discordgo.Role, UserRoles []string) (hp int) {
    for ri := range Roles {
        for uri := range UserRoles {
            if Roles[ri].ID == UserRoles[uri] && Roles[ri].Position > hp {
                hp = Roles[ri].Position
            }
        }
    }
    return hp
}