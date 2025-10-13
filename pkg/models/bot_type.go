package models

type BotType int

const (
	BotTypeGeneric         BotType = 1
	BotTypeIncomingWebhook BotType = 2
	BotTypeOutgoingWebhook BotType = 3
	BotTypeEmbedded        BotType = 4
)
