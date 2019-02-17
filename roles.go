package dgohelper

import (
    "strings"
    
    "github.com/bwmarrin/discordgo"
)

func GetRoleIDByName(Roles []*discordgo.Role, RoleName string) (id string) {
    for ri := range Roles {
        if strings.ToLower(Roles[ri].Name) == strings.ToLower(RoleName) {
            return Roles[ri].ID
        }
    }
    return ""
}

func GetRoleNameByID(Roles []*discordgo.Role, RoleID string) (id string) {
    for ri := range Roles {
        if Roles[ri].ID == RoleID {
            return Roles[ri].Name
        }
    }
    return ""
}