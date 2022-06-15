package domain

// 足りないカラムある
type RequestPayloadMeraki struct {
	Version          string `json:"version"`
	SharedSecret     string `json:"sharedSecret"`
	SentAt           string `json:"sentAt"`
	OrganizationId   string `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
	NetWorkId        string `json:"networkId"`
	NetWorkName      string `json:"networkName"`
	DeviceSerial     string `json:"deviceSerial"`
	AlertType        string `json:"alertType"`
	AlertTypeId      string `json:"alertId"`
	AlertLevel       string `json:"alertLevel"`
	OccuredAt        string `json:"occurredAt"`
	*AlertData       `json:"alertData"`
}

type AlertData struct {
	AlertConfigId   string         `json:"alertConfigId"`
	TriggerDatas    *[]TriggerData `json:"triggerData"`
	StartedAlerting bool           `json:"startedAlerting"`
}

type TriggerData struct {
	*Trigger `json:"trigger"`
}

type Trigger struct {
	Ts          int    `json:"ts"`
	Type        string `json:"type"`
	NodeId      int    `json:"nodeId"`
	SensorValue int    `json:"sensorValue"`
}
