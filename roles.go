package dgohelper

import (
    "strconv"
    "strings"
    "math/bits"
    
    "github.com/bwmarrin/discordgo"
)

type RoleProperty int

const (
    RolePropertyName RoleProperty = 1 << iota
    RolePropertyID
    RolePropertyPosition
)

func GetRoleByProperty(guildRoles []*discordgo.Role, property RoleProperty, value interface{}) (r *discordgo.Role) {
    var (
        sv string
        iv int64
    )
    
    switch v := value.(type) {
    case string:
        sv = v
    case int:
        iv = int64(v)
    default:
        return nil
    }

    if bits.OnesCount64(uint64(property)) > 1 {
        if property & RolePropertyName == RolePropertyName {
            if r = GetRoleByName(guildRoles, sv); r != nil {
                return
            }
        }

        if property & RolePropertyID == RolePropertyID {
            if r = GetRoleByID(guildRoles, sv); r != nil {
                return
            }
        }

        if property & RolePropertyPosition == RolePropertyPosition {
            if iv == 0 && sv != "" {
                if _, err := strconv.ParseInt(sv, 10, 64); err != nil {
                    return nil
                }
                iv, _ = strconv.ParseInt(sv, 10, 32)
            }

            if r = GetRoleByPosition(guildRoles, int(iv)); r != nil {
                return
            }
        }
    } else {
        switch property {
        case RolePropertyName:
            return GetRoleByName(guildRoles, sv)
        case RolePropertyID:
            return GetRoleByID(guildRoles, sv)
        case RolePropertyPosition:
            if iv == 0 && sv != "" {
                if _, err := strconv.ParseInt(sv, 10, 64); err != nil {
                    return nil
                }
                iv, _ = strconv.ParseInt(sv, 10, 32)
            }
            return GetRoleByPosition(guildRoles, int(iv))
        }
    }

    return nil
}

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

// GetRoleByID returns a discordgo.Role struct based on an id, returns nil if not found.
// guildroles   : Slice of the roles in a guild.
// roleID       : ID of the role you want.
func GetRoleByID(guildroles []*discordgo.Role, roleID string) (r *discordgo.Role) {
    for ri := range guildroles {
        if strings.ToLower(guildroles[ri].ID) == roleID {
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