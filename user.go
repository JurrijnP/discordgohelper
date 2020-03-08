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
    for gri := range guildRoles {
        for mri := range member.Roles {
            if guildRoles[gri].ID == member.Roles[mri] {
                perms |= guildRoles[gri].Permissions
            }
        }
    }
    
    if perms & perm == perm {
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