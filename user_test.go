package dgohelper

import (
    "testing"
    
    "github.com/bwmarrin/discordgo"
)

func TestMemberHasPermission(t *testing.T) {
    var guildRoles []*discordgo.Role
    
    guildRole := &discordgo.Role{
        ID: "role",
        Permissions: 2359304,
    }
    
    guildRoles = append(guildRoles, guildRole)
    
    member := &discordgo.Member{
        Roles: []string{"role"},
    }
    
    perm := 8
    
    if result := MemberHasPermission(guildRoles, member, perm); !result {
        t.Error(result)
    }
}