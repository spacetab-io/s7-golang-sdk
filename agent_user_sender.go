package sdk

type AgentUserSender struct {
	OtherIDs    *OtherIDs
	PseudoCity  string
	AgentUserID string
	UserRole    string
}

type OtherIDs struct {
	OtherID []*OtherID
}

type OtherID struct {
	Description string `xml:"Description,attr"`
	Value       string `xml:",chardata"`
}

func MakeAgentUserSender(airportCode, erspUserID, isoCountry, requestorID, requestorType, pseudoCity, agentUserID, userRole string) *AgentUserSender {
	agentUserSender := new(AgentUserSender)

	others := make([]*OtherID, 0)
	others = append(others, &OtherID{
		Description: "airlineVendorID",
		Value:       "S7",
	})
	others = append(others, &OtherID{
		Description: "airportCode",
		Value:       airportCode,
	})
	others = append(others, &OtherID{
		Description: "erspUserID",
		Value:       erspUserID,
	})
	others = append(others, &OtherID{
		Description: "isoCountry",
		Value:       isoCountry,
	})
	others = append(others, &OtherID{
		Description: "requestorID",
		Value:       requestorID,
	})
	others = append(others, &OtherID{
		Description: "requestorType",
		Value:       requestorType,
	})

	agentUserSender.OtherIDs = &OtherIDs{others}
	agentUserSender.PseudoCity = pseudoCity
	agentUserSender.AgentUserID = agentUserID
	agentUserSender.UserRole = userRole

	return agentUserSender
}
