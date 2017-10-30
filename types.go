package bitclient

const (
	REPO_READ  = "REPO_READ"
	REPO_WRITE = "REPO_WRITE"
	REPO_ADMIN = "REPO_ADMIN"
)

type Link map[string]string
type Links map[string][]Link

type Error struct {
	Context       string
	Message       string
	ExceptionName string
}

type Project struct {
	Key         string
	Id          uint
	Name        string
	Description string
	Public      bool
	Type        string
	Links       Links
}

type Repository struct {
	Slug          string
	Id            uint
	Name          string
	ScmId         string
	State         string
	StatusMessage string
	Forkable      bool
	Project       Project
	Public        bool
	Links         Links
}

type User struct {
	Name         string
	EmailAddress string
	Id           uint
	DisplayName  string
	Active       bool
	Slug         string
	Type         string
}

type DetailedUser struct {
	User                        User
	DirectoryName               string
	Deletable                   bool
	LastAuthenticationTimestamp uint
	MutableDetails              bool
	MutableGroups               bool
}

type Group struct {
	Name      string
	Deletable bool
}

type Tag struct {
	Id              string
	DisplayId       string
	Type            string
	LatestCommit    string
	LatestChangeset string
	Hash            string
}

type RefChange struct {
	Ref      Ref
	RefId    string
	FromHash string
	ToHash   string
	Type     string
}

type Ref struct {
	Id        string
	DisplayId string
	Type      string
}

type PullRequestSuggestion struct {
	ChangeTime uint
	RefChange  RefChange
	Repository Repository
	FromRef    Ref
	ToRef      Ref
}

type Author struct {
	User     User
	Role     string
	Approved bool
	Status   string
}

type Participant struct {
	User               User
	LastReviewedCommit string
	Role               string
	Approved           bool
	Status             string
}

type PullRequest struct {
	Id           uint
	Version      uint
	Title        string
	Description  string
	State        string
	Open         bool
	Closed       bool
	CreatedDate  uint
	UpdatedDate  uint
	FromRef      Ref
	ToRef        Ref
	Locked       bool
	Author       Author
	Reviewers    []Participant
	Participants []Participant
	Links        Links
}

type SonarSettings struct {
	Project       SonarSettingsProject       `json:"project,omitempty"`
	Issues        SonarSettingsIssues        `json:"issues,omitempty"`
	DuplicateCode SonarSettingsDuplicateCode `json:"duplicateCode,omitempty"`
	TestCoverage  SonarSettingsTestCoverage  `json:"testCoverage,omitempty"`
	Statistics    SonarSettingsStatistics    `json:"statistics,omitempty"`
	MergeChecks   SonarSettingsMergeChecks   `json:"mergeChecks,omitempty"`
	Provisioning  SonarSettingsProvisioning  `json:"provisioning,omitempty"`
}

type SonarSettingsProject struct {
	SonarEnabled                 bool   `json:"sonarEnabled,omitempty"`
	ServerConfigId               int    `json:"serverConfigId,omitempty"`
	MasterProjectKey             string `json:"masterProjectKey,omitempty"`
	ProjectBaseKey               string `json:"projectBaseKey,omitempty"`
	AnalysisMode                 string `json:"analysisMode,omitempty"`
	BuildType                    string `json:"buildType,omitempty"`
	IllegalBranchCharReplacement string `json:"illegalBranchCharReplacement,omitempty"`
	BranchPrefix                 string `json:"branchPrefix,omitempty"`
	PullRequestBranch            string `json:"pullRequestBranch,omitempty"`
	ShowOnlyNewOrChangedLines    bool   `json:"showOnlyNewOrChangedLines,omitempty"`
	ProjectCleanupEnabled        bool   `json:"projectCleanupEnabled,omitempty"`
	ForkCleanupEnabled           bool   `json:"forkCleanupEnabled,omitempty"`
	MatchingBranchesRegex        string `json:"matchingBranchesRegex,omitempty"`
	IncrementalModeForMatches    string `json:"incrementalModeForMatches,omitempty"`
}

type SonarSettingsIssues struct {
	ShowSonarIssues        bool   `json:"showOnlyNewIssues,omitempty"`
	ShowOnlyNewIssues      bool   `json:"showOnlyNewIssues,omitempty"`
	SonarIssuesMinSeverity string `json:"showOnlyNewIssues,omitempty"`
}

type SonarSettingsDuplicateCode struct {
	ShowDuplicateCode bool `json:"showDuplicateCode,omitempty"`
}

type SonarSettingsTestCoverage struct {
	ShowCoverage bool   `json:"showCoverage,omitempty"`
	CoverageType string `json:"coverageType,omitempty"`
}

type SonarSettingsStatistics struct {
	ShowStatistics           bool `json:"showStatistics,omitempty"`
	CollapsedByDefault       bool `json:"collapsedByDefault,omitempty"`
	ShowQualityGates         bool `json:"showQualityGates,omitempty"`
	ShowBehindCommitsWarning bool `json:"showBehindCommitsWarning,omitempty"`
}

type SonarSettingsMergeChecks struct {
	MergeChecksEnabled       bool   `json:"mergeChecksEnabled,omitempty"`
	QualityGatesEnabled      bool   `json:"qualityGatesEnabled,omitempty"`
	TechnicalDebtMaxIncrease int    `json:"technicalDebtMaxIncrease,omitempty"`
	SonarIssuesMaxNew        int    `json:"sonarIssuesMaxNew,omitempty"`
	SonarIssueTypeMaxNew     string `json:"sonarIssueTypeMaxNew,omitempty"`
	DuplicateCodeMaxIncrease int    `json:"duplicateCodeMaxIncrease,omitempty"`
	CoverageMinPercentage    int    `json:"coverageMinPercentage,omitempty"`
}

type SonarSettingsProvisioning struct {
	QualityProfileProvisioningEnabled bool `json:"qualityProfileProvisioningEnabled,omitempty"`
	PropertiesProvisioningEnabled     bool `json:"propertiesProvisioningEnabled,omitempty"`
}

type UserPermission struct {
	User       User
	Permission string `json:"permission,omitempty"`
}

type GroupPermission struct {
	Group      Group
	Permission string `json:"permission,omitempty"`
}
