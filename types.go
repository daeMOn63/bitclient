package bitclient

// ##############################################################################
// Type
// ##############################################################################

type BitClientType interface {
	object()
}
type BitClientTypeImpl struct{}

func (b BitClientTypeImpl) object() {}

type Link map[string]string
type Links map[string][]Link

const (
	REPO_READ  = "REPO_READ"
	REPO_WRITE = "REPO_WRITE"
	REPO_ADMIN = "REPO_ADMIN"
)

type Error struct {
	BitClientTypeImpl
	Context       string
	Message       string
	ExceptionName string
}

type Project struct {
	BitClientTypeImpl
	Key         string
	Id          uint
	Name        string
	Description string
	Public      bool
	Type        string
	Links       Links
}

type Repository struct {
	BitClientTypeImpl
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
	BitClientTypeImpl
	Name         string
	EmailAddress string
	Id           uint
	DisplayName  string
	Active       bool
	Slug         string
	Type         string
}
