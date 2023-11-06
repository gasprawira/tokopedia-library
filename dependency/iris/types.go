package iris

type (
	// Data describes event data with default parameters to be passed
	Data struct {
		UserAgent       string      `json:"user_agent,omitempty"`       // Browser user agent
		SessionID       string      `json:"session_id,omitempty"`       // Tokopedia Session ID
		SourceType      string      `json:"source_type,omitempty"`      // Please follow source type name mentioned in this guideline https://phab.tokopedia.com/w/api/jarvis/iris/guideline/
		IPAddress       string      `json:"ip_address,omitempty"`       // Ip Address from HTTP Header
		ClientTimestamp float64     `json:"client_timestamp,omitempty"` // Client unix timestamp
		EventData       interface{} `json:"data"`                       // Free form data based on your defined struct
	}
)
