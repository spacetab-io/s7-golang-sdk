package structsV052

// AgentUserSender struct describes AgentUserSender section of S7 request
type AgentUserSender struct {
	Name string `xml:",omitempty"`
	// Type string `xml:",omitempty"`
	// Contacts	*AgentUserContacts	`xml:",omitempty"`
	OtherIDs    *OtherIDs
	PseudoCity  string
	AgentUserID string
	UserRole    string
}

//type AgentUserContacts struct {
//	Contact		Contact
//}

// OtherIDs is an OtherIDs subsection in AgentUserSender section of S7 request
type OtherIDs struct {
	OtherID []*OtherID
}

// OtherID is an element of OtherIDs array
type OtherID struct {
	Description string `xml:"Description,attr"`
	Value       string `xml:",chardata"`
}

// MakeS7AgentUserSender creates AgentUserSender section of S7 request
func MakeS7AgentUserSender(
	agentName,
	posType,
	requestorType,
	password,
	pseudoCity,
	agentUserID,
	userRole string,
) *AgentUserSender {
	var others []*OtherID
	others = append(others, &OtherID{
		Description: "POS_Type",
		Value:       posType,
	})
	others = append(others, &OtherID{
		Description: "requestorType",
		Value:       requestorType,
	})
	others = append(others, &OtherID{
		Description: "Password",
		Value:       password,
	})

	return &AgentUserSender{
		Name: agentName,
		// Type: agentType,
		//Contacts:    contacts,
		OtherIDs:    &OtherIDs{OtherID: others},
		PseudoCity:  pseudoCity,
		AgentUserID: agentUserID,
		UserRole:    userRole,
	}
}
