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
