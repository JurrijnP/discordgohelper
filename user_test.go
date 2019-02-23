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

func TestMemberHasRole(t *testing.T) {
    userRoles := []string{"role1", "role2", "role3"}
    
    if result := MemberHasRole(userRoles, "role2"); !result {
        t.Error(result)
    }
    
    if result := MemberHasRole(userRoles, "role4"); result {
        t.Error(result)
    }
}