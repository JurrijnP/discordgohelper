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

// MemberHasRole checks if a user has a role.
// userRoles   : Slice of role IDs.
// roleID      : ID of the role you want to check.
func MemberHasRole(userRoles []string, roleID string) bool {
    for uri := range userRoles {
        if userRoles[uri] == roleID {
            return true
        }
    }
    
    return false
}