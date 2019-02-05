package dgohelper

import (
    "strings"

    "github.com/bwmarrin/discordgo"
)

func GetPermissionBasedOnString(arg string) int {
    switch strings.ToLower(arg) {
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
    case "view_audit_log":
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
    //case "priority_speaker":
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
    }
    return 0
}