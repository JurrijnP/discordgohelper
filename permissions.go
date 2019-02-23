package dgohelper

import (
    "strings"

    "github.com/bwmarrin/discordgo"
)

// GetPermissionBasedOnString returns an integer from a permission based on a string.
// arg   : The a permission name.
func GetPermissionBasedOnString(arg string) int {
    switch strings.ToLower(strings.Replace(arg, " ", "_", -1)) {
    case "create_instant_invite":
        return discordgo.PermissionCreateInstantInvite
    case "kick_members":
        return discordgo.PermissionKickMembers
    case "ban_members":
        return discordgo.PermissionBanMembers
    case "administrator":
        return discordgo.PermissionAdministrator
    case "manage_channels":
        return discordgo.PermissionManageChannels
    case "manage_guild":
        return discordgo.PermissionManageServer
    case "manage_server":
        return discordgo.PermissionManageServer
    case "add_reactions":
        return discordgo.PermissionAddReactions
    case "view_audit_logs":
        return discordgo.PermissionViewAuditLogs
    case "view_channel":
        return discordgo.PermissionReadMessages
    case "read_messages":
        return discordgo.PermissionReadMessages
    case "send_messages":
        return discordgo.PermissionSendMessages
    case "send_tts_messages":
        return discordgo.PermissionSendTTSMessages
    case "manage_messages":
        return discordgo.PermissionManageMessages
    case "embed_links":
        return discordgo.PermissionEmbedLinks
    case "attach_files":
        return discordgo.PermissionAttachFiles
    case "read_message_history":
        return discordgo.PermissionReadMessageHistory
    case "mention_everyone":
        return discordgo.PermissionMentionEveryone
    case "use_external_emojis":
        return discordgo.PermissionUseExternalEmojis
    case "connect":
        return discordgo.PermissionVoiceConnect
    case "speak":
        return discordgo.PermissionVoiceSpeak
    case "mute_members":
        return discordgo.PermissionVoiceMuteMembers
    case "deafen_members":
        return discordgo.PermissionVoiceDeafenMembers
    case "move_members":
        return discordgo.PermissionVoiceMoveMembers
    case "use_vad":
        return discordgo.PermissionVoiceUseVAD
    case "use_voice_activity":
        return discordgo.PermissionVoiceUseVAD
    case "priority_speaker":
        return 0x00000100
        //return discordgo.PermissionVoicePrioritySpeaker
    case "change_nickname":
        return discordgo.PermissionChangeNickname
    case "manage_nicknames":
        return discordgo.PermissionManageNicknames
    case "manage_roles":
        return discordgo.PermissionManageRoles
    case "manage_permissions":
        return discordgo.PermissionManageRoles
    case "manage_webhooks":
        return discordgo.PermissionManageWebhooks
    case "manage_emojis":
        return discordgo.PermissionManageEmojis
        
    case "text_all":
        return discordgo.PermissionAllText
    case "all_text":
        return discordgo.PermissionAllText
    case "voice_all":
        return (discordgo.PermissionAllVoice|0x00000100)
    case "all_voice":
        return (discordgo.PermissionAllVoice|0x00000100)
    case "channel_all":
        return discordgo.PermissionAllChannel
    case "all_channel":
        return discordgo.PermissionAllChannel
    case "all":
        return discordgo.PermissionAll
    }
    return 0
}

// GetPermissionBasedOnString returns a name from a permission based on a integer.
// perm   : Integer of the permission.
// mode   : Determines the string to return in some cases.
//          (e.g. Manage Roles <> Manage Permissions)
func GetPermissionAsString(perm, mode int) string {
    switch perm {
    case discordgo.PermissionCreateInstantInvite:
        return "Create Instant Invite"
    case discordgo.PermissionKickMembers:
        return "Kick Members"
    case discordgo.PermissionBanMembers:
        return "Ban Members"
    case discordgo.PermissionAdministrator:
        return "Administrator"
    case discordgo.PermissionManageChannels:
        return "Manage Channels"
    case discordgo.PermissionManageServer:
        return "Manage Server"
    case discordgo.PermissionAddReactions:
        return "Add Reactions"
    case discordgo.PermissionViewAuditLogs:
        return "View Audit Logs"
    case discordgo.PermissionReadMessages:
        return "Read Messages"
    case discordgo.PermissionSendMessages:
        return "Send Messages"
    case discordgo.PermissionSendTTSMessages:
        return "Send TTS Messages"
    case discordgo.PermissionManageMessages:
        return "Manage Messages"
    case discordgo.PermissionEmbedLinks:
        return "Embed Links"
    case discordgo.PermissionAttachFiles:
        return "Attach Files"
    case discordgo.PermissionReadMessageHistory:
        return "Read Message History"
    case discordgo.PermissionMentionEveryone:
        return "Mention Everyone"
    case discordgo.PermissionUseExternalEmojis:
        return "Use External Emojis"
    case discordgo.PermissionVoiceConnect:
        return "Connect (Voice)"
    case discordgo.PermissionVoiceSpeak:
        return "Speak (Voice)"
    case discordgo.PermissionVoiceMuteMembers:
        return "Mute Members (Voice)"
    case discordgo.PermissionVoiceDeafenMembers:
        return "Deafen Members (Voice)"
    case discordgo.PermissionVoiceMoveMembers:
        return "Move Members (Voice)"
    case discordgo.PermissionVoiceUseVAD:
        return "Use Voice Activity"
    case 0x00000100:
    //case discordgo.PermissionVoicePrioritySpeaker:
        return "Priority Speaker (Voice)"
    case discordgo.PermissionChangeNickname:
        return "Change Nickname"
    case discordgo.PermissionManageNicknames:
        return "Manage Nicknames"
    case discordgo.PermissionManageRoles:
        switch mode {
        case 0:
            return "Manage Roles"
        case 1:
            return "Manage Permissions"
        }
    case discordgo.PermissionManageWebhooks:
        return "Manage Webhooks"
    case discordgo.PermissionManageEmojis:
        return "Manage Emojis"
    }
    return ""
}

// GetPermissionsFromSlice returns an integer consisting of (multiple) permissions based on a slice of strings.
// perms   : Slice of strings of permission names.
func GetPermissionsFromSlice(perms []string) (apermissions int) {
    for pi := range perms {
        switch strings.ToLower(strings.Replace(perms[pi], " ", "_", -1)) {
        case "create_instant_invite":
            apermissions |= discordgo.PermissionCreateInstantInvite
        case "kick_members":
            apermissions |= discordgo.PermissionKickMembers
        case "ban_members":
            apermissions |= discordgo.PermissionBanMembers
        case "administrator":
            apermissions |= discordgo.PermissionAdministrator
        case "manage_channels":
            apermissions |= discordgo.PermissionManageChannels
        case "manage_guild":
            apermissions |= discordgo.PermissionManageServer
        case "manage_server":
            apermissions |= discordgo.PermissionManageServer
        case "add_reactions":
            apermissions |= discordgo.PermissionAddReactions
        case "view_audit_logs":
            apermissions |= discordgo.PermissionViewAuditLogs
        case "view_channel":
            apermissions |= discordgo.PermissionReadMessages
        case "read_messages":
            apermissions |= discordgo.PermissionReadMessages
        case "send_messages":
            apermissions |= discordgo.PermissionSendMessages
        case "send_tts_messages":
            apermissions |= discordgo.PermissionSendTTSMessages
        case "manage_messages":
            apermissions |= discordgo.PermissionManageMessages
        case "embed_links":
            apermissions |= discordgo.PermissionEmbedLinks
        case "attach_files":
            apermissions |= discordgo.PermissionAttachFiles
        case "read_message_history":
            apermissions |= discordgo.PermissionReadMessageHistory
        case "mention_everyone":
            apermissions |= discordgo.PermissionMentionEveryone
        case "use_external_emojis":
            apermissions |= discordgo.PermissionUseExternalEmojis
        case "connect":
            apermissions |= discordgo.PermissionVoiceConnect
        case "speak":
            apermissions |= discordgo.PermissionVoiceSpeak
        case "mute_members":
            apermissions |= discordgo.PermissionVoiceMuteMembers
        case "deafen_members":
            apermissions |= discordgo.PermissionVoiceDeafenMembers
        case "move_members":
            apermissions |= discordgo.PermissionVoiceMoveMembers
        case "use_vad":
            apermissions |= discordgo.PermissionVoiceUseVAD
        case "use_voice_activity":
            apermissions |= discordgo.PermissionVoiceUseVAD
        case "priority_speaker":
            apermissions |= 0x00000100
            //apermissions |= discordgo.PermissionVoicePrioritySpeaker
        case "change_nickname":
            apermissions |= discordgo.PermissionChangeNickname
        case "manage_nicknames":
            apermissions |= discordgo.PermissionManageNicknames
        case "manage_roles":
            apermissions |= discordgo.PermissionManageRoles
        case "manage_permissions":
            apermissions |= discordgo.PermissionManageRoles
        case "manage_webhooks":
            apermissions |= discordgo.PermissionManageWebhooks
        case "manage_emojis":
            apermissions |= discordgo.PermissionManageEmojis
            
        case "text_all":
            apermissions |= discordgo.PermissionAllText
        case "all_text":
            apermissions |= discordgo.PermissionAllText
        case "voice_all":
            apermissions |= (discordgo.PermissionAllVoice|0x00000100)
        case "all_voice":
            apermissions |= (discordgo.PermissionAllVoice|0x00000100)
        case "channel_all":
            apermissions |= discordgo.PermissionAllChannel
        case "all_channel":
            apermissions |= discordgo.PermissionAllChannel
        case "all":
            apermissions |= discordgo.PermissionAll
        }
    }
    return apermissions
}

// GetPermissionsFromSlice returns a slice of strings based on an integer consisting of (multiple) permissions.
// perms   : Slice of strings of permission names.
// mode    : Determines the string to return in some cases.
//           (e.g. Manage Roles <> Manage Permissions)
func GetAllPermissionsFromBit(perms, mode int) (apermissions []string) {
    for i := 0; 1<<uint(i) < perms; i++ {
        if perms&(1<<uint(i)) != 0 {
            switch 1<<uint(i) {
            case discordgo.PermissionCreateInstantInvite:
                apermissions = append(apermissions, "Create Instant Invite")
            case discordgo.PermissionKickMembers:
                apermissions = append(apermissions, "Kick Members")
            case discordgo.PermissionBanMembers:
                apermissions = append(apermissions, "Ban Members")
            case discordgo.PermissionAdministrator:
                apermissions = append(apermissions, "Administrator")
            case discordgo.PermissionManageChannels:
                apermissions = append(apermissions, "Manage Channels")
            case discordgo.PermissionManageServer:
                apermissions = append(apermissions, "Manage Server")
            case discordgo.PermissionAddReactions:
                apermissions = append(apermissions, "Add Reactions")
            case discordgo.PermissionViewAuditLogs:
                apermissions = append(apermissions, "View Audit Logs")
            case discordgo.PermissionReadMessages:
                apermissions = append(apermissions, "Read Messages")
            case discordgo.PermissionSendMessages:
                apermissions = append(apermissions, "Send Messages")
            case discordgo.PermissionSendTTSMessages:
                apermissions = append(apermissions, "Send TTS Messages")
            case discordgo.PermissionManageMessages:
                apermissions = append(apermissions, "Manage Messages")
            case discordgo.PermissionEmbedLinks:
                apermissions = append(apermissions, "Embed Links")
            case discordgo.PermissionAttachFiles:
                apermissions = append(apermissions, "Attach Files")
            case discordgo.PermissionReadMessageHistory:
                apermissions = append(apermissions, "Read Message History")
            case discordgo.PermissionMentionEveryone:
                apermissions = append(apermissions, "Mention Everyone")
            case discordgo.PermissionUseExternalEmojis:
                apermissions = append(apermissions, "Use External Emojis")
            case discordgo.PermissionVoiceConnect:
                apermissions = append(apermissions, "Connect (Voice)")
            case discordgo.PermissionVoiceSpeak:
                apermissions = append(apermissions, "Speak (Voice)")
            case discordgo.PermissionVoiceMuteMembers:
                apermissions = append(apermissions, "Mute Members (Voice)")
            case discordgo.PermissionVoiceDeafenMembers:
                apermissions = append(apermissions, "Deafen Members (Voice)")
            case discordgo.PermissionVoiceMoveMembers:
                apermissions = append(apermissions, "Move Members (Voice)")
            case discordgo.PermissionVoiceUseVAD:
                apermissions = append(apermissions, "Use Voice Activity")
            case 0x00000100:
            //case discordgo.PermissionVoicePrioritySpeaker:
                apermissions = append(apermissions, "Priority Speaker (Voice)")
            case discordgo.PermissionChangeNickname:
                apermissions = append(apermissions, "Change Nickname")
            case discordgo.PermissionManageNicknames:
                apermissions = append(apermissions, "Manage Nicknames")
            case discordgo.PermissionManageRoles:
                switch mode {
                case 0:
                    apermissions = append(apermissions, "Manage Roles")
                case 1:
                    apermissions = append(apermissions, "Manage Permissions")
                }
            case discordgo.PermissionManageWebhooks:
                apermissions = append(apermissions, "Manage Webhooks")
            case discordgo.PermissionManageEmojis:
                apermissions = append(apermissions, "Manage Emojis")
            }
        }
    }
    return apermissions
}