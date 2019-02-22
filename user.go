package dgohelper

import (
    //"strings"
    
    "github.com/bwmarrin/discordgo"
)

// MemberHasPermission checks if a member has the permission(s) from their roles.
// guildRoles   : Slice of guild roles.
// member       : discordgo.Member struct of the member.
// perm         : Permission(s) you want to check.
func MemberHasPermission(guildRoles []*discordgo.Role, member *discordgo.Member, perm int) bool {
    var perms int
    for _, gr := range guildRoles {
        for _, r := range member.Roles {
            if gr.ID == r {
                perms |= gr.Permissions
            }
        }
    }
    
    if perms&perm == perm {
        return true
    }
    
    return false
}