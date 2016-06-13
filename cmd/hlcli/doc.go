package main

// GENERATED BY gen.go

var methodParamMap = map[string][]string{
	"encodeXML":    []string{"v"},
	"decodeXML":    []string{"buf", "takeFirstEl"},
	"buildRequest": []string{"urlstr", "v"},
	"doReq":        []string{"path", "v", "takeFirstEl"},
	"doReqString":  []string{"path", "v", "elName"},
	"doReqCheckOK": []string{"path", "v"},
	"Do":           []string{"path", "v"},
	"NewSessionAndTokenID": []string{},
	"SetSessionAndTokenID": []string{"sessionID", "tokenID"},
	"GlobalConfig":         []string{},
	"NetworkTypes":         []string{},
	"PCAssistantConfig":    []string{},
	"DeviceConfig":         []string{},
	"WebUIConfig":          []string{},
	"SmsConfig":            []string{},
	"WlanBasicSettings":    []string{},
	"CradleStatusInfo":     []string{},
	"AutorunVersion":       []string{},
	"DeviceBasicInfo":      []string{},
	"PublicKey":            []string{},
	"Reboot":               []string{},
	"DeviceFeatures":       []string{},
	"DeviceInfo":           []string{},
	"SignalInfo":           []string{},
	"ConnectionInfo":       []string{},
	"GlobalFeatures":       []string{},
	"Language":             []string{},
	"NotificationInfo":     []string{},
	"SimInfo":              []string{},
	"StatusInfo":           []string{},
	"TrafficStatistics":    []string{},
	"NetworkInfo":          []string{},
	"WifiFeatures":         []string{},
	"ModeInfo":             []string{},
	"PinInfo":              []string{},
	"doReqPin":             []string{"pt", "cur", "new", "puk"},
	"PinEnter":             []string{"pin"},
	"PinActivate":          []string{"pin"},
	"PinDeactivate":        []string{"pin"},
	"PinChange":            []string{"pin", "new"},
	"PinEnterPuk":          []string{"puk", "new"},
	"Connect":              []string{},
	"Disconnect":           []string{},
	"ProfileInfo":          []string{},
	"SmsList":              []string{"boxType", "page", "count", "ascending", "unreadPreferred"},
	"SmsCount":             []string{},
	"SmsSend":              []string{"msg", "to"},
	"SmsSendStatus":        []string{},
	"SmsSetRead":           []string{"id"},
	"SmsDelete":            []string{"id"},
	"UssdStatus":           []string{},
	"UssdCode":             []string{"code"},
	"UssdContent":          []string{},
	"UssdRelease":          []string{},
}

var methodCommentMap = map[string]string{
	"encodeXML":    "encodeXML encodes a map to standard XML values.",
	"decodeXML":    "decodeXML decodes buf into its simple xml values.",
	"buildRequest": "buildRequest creates a request for use with the Client.",
	"doReq":        "doReq sends a request to the server with the provided path. If data is nil,then GET will be used as the HTTP method, otherwise POST will be used.",
	"doReqString":  "doReqString wraps a request operation, returning the data of the specifiedchild node named elName as a string.",
	"doReqCheckOK": "doReqCheckOK wraps a request operation (ie, connect, disconnect, etc),checking success via the presence of 'OK' in the XML <response/>.",
	"Do":           "Do sends a request to the server with the provided path. If data is nil,then GET will be used as the HTTP method, otherwise POST will be used.",
	"NewSessionAndTokenID": "NewSessionAndTokenID starts a session with the server, and returns thesession and token.",
	"SetSessionAndTokenID": "SetSessionID sets the sessionID for the Client.",
	"GlobalConfig":         "GlobalConfig retrieves the global Hilink configuration.",
	"NetworkTypes":         "NetworkTypes retrieves the available network types.",
	"PCAssistantConfig":    "PCAssistantConfig retrieves the PC Assistant configuration.",
	"DeviceConfig":         "DeviceConfig retrieves device configuration.",
	"WebUIConfig":          "WebUIConfig retrieves the WebUI configuration.",
	"SmsConfig":            "SmsConfig retrieves device SMS configuration.",
	"WlanBasicSettings":    "WlanBasicSettings returns the basic WLAN settings.",
	"CradleStatusInfo":     "CradleStatusInfo retrieves cradle status information.",
	"AutorunVersion":       "AutorunVersion retrieves autorun version.",
	"DeviceBasicInfo":      "DeviceBasicInfo retrieves basic device information.",
	"PublicKey":            "PublicKey retrieves the hilink public key.",
	"Reboot":               "Reboot restarts the Hilink device.",
	"DeviceFeatures":       "DeviceFeatures retrieves device feature information.",
	"DeviceInfo":           "DeviceInfo retrieves Hilink device information.",
	"SignalInfo":           "SignalInfo retrieves signal information.",
	"ConnectionInfo":       "ConnectionInfo retrieves connection (dialup) information.",
	"GlobalFeatures":       "GlobalFeatures retrieves global feature information.",
	"Language":             "Language retrieves the current language.",
	"NotificationInfo":     "NotificationInfo retrieves notification information.",
	"SimInfo":              "SimInfo retrieves SIM card information.",
	"StatusInfo":           "StatusInfo retrieves Hilink connection status information.",
	"TrafficStatistics":    "TrafficStatistics retrieves traffic statistics.",
	"NetworkInfo":          "NetworkInfo retrieves network provider information.",
	"WifiFeatures":         "WifiFeatures retrieves wifi feature information.",
	"ModeInfo":             "ModeInfo retrieves network mode information.",
	"PinInfo":              "PinInfo retrieves SIM PIN status information.",
	"doReqPin":             "doReqPin wraps a SIM PIN manipulation request.",
	"PinEnter":             "PinEnter enters a SIM PIN.",
	"PinActivate":          "PinActivate activates a SIM PIN.",
	"PinDeactivate":        "PinDeactivate deactivates a SIM PIN.",
	"PinChange":            "PinChange changes a SIM PIN.",
	"PinEnterPuk":          "PinEnterPuk enters a puk SIM PIN.",
	"Connect":              "Connect connects the Hilink device to the network provider.",
	"Disconnect":           "Disconnect disconnects the Hilink device from the network provider.",
	"ProfileInfo":          "ProfileInfo retrieves profile information (ie, APN).",
	"SmsList":              "SmsList retrieves a list of SMS from an inbox.",
	"SmsCount":             "SmsCount retrieves the count of SMS per inbox type.",
	"SmsSend":              "SmsSend sends an SMS.",
	"SmsSendStatus":        "SmsSendStatus retrieves the SMS send status information.",
	"SmsSetRead":           "SmsSetRead sets the read status of a SMS.",
	"SmsDelete":            "SmsDelete deletes a specified SMS.",
	"UssdStatus":           "UssdStatus determines if the Hilink device is currently engaged in a USSDsession.",
	"UssdCode":             "UssdCode sends a USSD code to the Hilink device.",
	"UssdContent":          "UssdContent retrieves the content buffer of the active USSD session.",
	"UssdRelease":          "UssdRelease releases the active USSD session.",
}