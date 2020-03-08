package dgohelper

import (
    "testing"
    
    "github.com/bwmarrin/discordgo"
)

func TestGetPermissionInt(t *testing.T) {
    if result := GetPermissionInt("kick_members"); result != discordgo.PermissionKickMembers {
        t.Error(result)
    }
    
    if result := GetPermissionInt("view_audit_logs"); result != discordgo.PermissionViewAuditLogs {
        t.Error(result)
    }
    
    if result, result2 := GetPermissionInt("manage_roles"), GetPermissionInt("manage_permissions"); result != discordgo.PermissionManageRoles || result != result2 {
        t.Error(result)
    }
    
    if result := GetPermissionInt("text_all"); result != discordgo.PermissionAllText {
        t.Error(result)
    }
}

func TestGetPermissionString(t *testing.T) {
    if result := GetPermissionString(discordgo.PermissionKickMembers, 0); result != "Kick Members" {
        t.Error(result)
    }
    
    if result := GetPermissionString(discordgo.PermissionViewAuditLogs, 0); result != "View Audit Logs" {
        t.Error(result)
    }
    
    if result, result2 := GetPermissionString(discordgo.PermissionManageRoles, 0), GetPermissionString(discordgo.PermissionManageRoles, 1); result != "Manage Roles" || result2 != "Manage Permissions" {
        t.Error(result, result2)
    }
}

func TestGetPermissionIntAll(t *testing.T) {
    perms1 := []string{"kick_members", "view audit logs"}
    perms2 := []string{"administrator", "manage_roles", "manage_permissions"}
    ip1 := discordgo.PermissionKickMembers   | discordgo.PermissionViewAuditLogs
    ip2 := discordgo.PermissionAdministrator | discordgo.PermissionManageRoles
    
    if result1, result2 := GetPermissionIntAll(perms1), GetPermissionIntAll(perms2); result1 != ip1 || result2 != ip2 {
        t.Error(result1, result2)
    }
}

func TestGetPermissionSlice(t *testing.T) {
    perms1 := []string{"Kick Members", "View Audit Logs"}
    perms2 := []string{"Administrator", "Manage Roles"}
    perms3 := []string{"Administrator", "Manage Permissions"}
    ip1 := discordgo.PermissionKickMembers   | discordgo.PermissionViewAuditLogs
    ip2 := discordgo.PermissionAdministrator | discordgo.PermissionManageRoles
    
    if result1, result2, result3 := GetPermissionSlice(ip1, 0), GetPermissionSlice(ip2, 0), GetPermissionSlice(ip2, 1); (result1[0] != perms1[0] || result1[1] != perms1[1]) || result2[0] != perms2[0] || result2[1] != perms2[1] || (result3[0] != perms3[0] || result3[1] != perms3[1]) {
        t.Error(result1, result2, result3)
    }
}