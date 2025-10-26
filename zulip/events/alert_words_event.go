package events

type AlertWordsEvent struct {
	event

	// An array of strings, where each string is an alert word (or phrase) configured by the user.
	AlertWords []string `json:"alert_words,omitempty"`
}
