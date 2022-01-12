package main

// structure containing the debug flag to check on
type debug struct {
	PrintDebug bool
}

// Flags
// flags structures
// separated in 2 structure because some function needs only the mandatory
type flags struct {
	mandatoryFlags MandatoryFlags
	optionalFlags  optionalFlags
}

type MandatoryFlags struct {
	orgID          string
	endpointAPI    string
	apiToken       string
	jiraProjectID  string
	jiraProjectKey string
}

type optionalFlags struct {
	projectID              string
	jiraTicketType         string
	severity               string
	issueType              string
	maturityFilterString   string
	assigneeID             string
	assigneeName           string
	labels                 string
	priorityIsSeverity     bool
	priorityScoreThreshold int
	debug                  bool
	dryRun                 bool
	ifUpgradeAvailableOnly bool
}