package dgohelper

import (
    "testing"
    
    "github.com/bwmarrin/discordgo"
)

func TestGetRoleByName(t *testing.T) {
    var guildRoles []*discordgo.Role
    
    guildRole := &discordgo.Role{
        Name: "role",
    }
    
    guildRoles = append(guildRoles, guildRole)
    
    if result := GetRoleByName(guildRoles, "role"); result == nil {
        t.Error(result)
    }
}

func TestGetRoleByPosition(t *testing.T) {
    var guildRoles []*discordgo.Role
    
    guildRole := &discordgo.Role{
        Position: 5,
    }
    
    guildRoles = append(guildRoles, guildRole)
    
    if result := GetRoleByPosition(guildRoles, 5); result == nil {
        t.Error(result)
    }
}

func TestGetHighestRole(t *testing.T) {
    var guildRoles []*discordgo.Role
    
    guildRole := &discordgo.Role{
        ID:       "role1",
        Position: 5,
    }
    guildRole2 := &discordgo.Role{
        ID:       "role2",
        Position: 8,
    }
    guildRole3 := &discordgo.Role{
        ID:       "role3",
        Position: 10,
    }
    
    guildRoles = append(guildRoles, guildRole)
    guildRoles = append(guildRoles, guildRole2)
    guildRoles = append(guildRoles, guildRole3)
    
    uroles := []string{"role1", "role2", "role3"}
    
    if result := GetHighestRole(guildRoles, uroles); result != 10 {
        t.Error(result)
    }
}
