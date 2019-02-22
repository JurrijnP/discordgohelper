package dgohelper

import (
    "strings"
    
    "github.com/bwmarrin/discordgo"
)

// GetRoleByName returns a discordgo.Role struct based on a name, returns nil if not found.
// guildroles   : Slice of the roles in a guild.
// rolename     : Name of the role you want.
func GetRoleByName(guildroles []*discordgo.Role, rolename string) (r *discordgo.Role) {
    for ri := range guildroles {
        if strings.ToLower(guildroles[ri].Name) == strings.ToLower(rolename) {
            return guildroles[ri]
        }
    }
    return nil
}

// GetRoleByName returns a discordgo.Role struct based on a position, returns nil if not found.
// guildroles   : Slice of the roles in a guild.
// position     : Position of the role you want to be returned.
func GetRoleByPosition(guildroles []*discordgo.Role, position int) (r *discordgo.Role) {
    for ri := range guildroles {
        if guildroles[ri].Position == position {
            return guildroles[ri]
        }
    }
    return nil
}

// GetHighestRole returns an integer based on the highest role of a user.
// guildroles   : Slice of the roles in a guild.
// uroles       : Slice of IDs of the roles from a user.
func GetHighestRole(guildroles []*discordgo.Role, uroles []string) (hp int) {
    for ri := range guildroles {
        for uri := range uroles {
            if guildroles[ri].ID == uroles[uri] && guildroles[ri].Position > hp {
                hp = guildroles[ri].Position
            }
        }
    }
    return hp
}