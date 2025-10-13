package models

type RecipientType string

const (
	RecipientsTypeDirect  RecipientType = "direct"
	RecipientsTypeStream  RecipientType = "stream"
	RecipientsTypePrivate RecipientType = "private"
	RecipientsTypeChannel RecipientType = "channel"
)
