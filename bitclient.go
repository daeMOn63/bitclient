package bitclient

import (
	"fmt"
	"github.com/dghubble/sling"
)

// ##############################################################################
// Client
// ##############################################################################

type BitClient struct {
	sling *sling.Sling
}

func NewBitClient(url string, username string, password string) *BitClient {
	return &BitClient{
		sling: sling.New().Base(url).Set("Accept", "application/json").SetBasicAuth(username, password),
	}
}

func (bc *BitClient) GetProjects() ([]Project, error) {

	rData := new(GetProjectsResponse)
	rError := new(ErrorResponse)
	_, err := bc.DoGetPaged("/rest/api/1.0/projects", rData, rError)

	return rData.Values, err
}

func (bc *BitClient) GetRepositories(projectKey string) ([]Repository, error) {
	rData := new(GetRepositoriesResponse)
	rError := new(ErrorResponse)

	_, err := bc.DoGetPaged(fmt.Sprintf("/rest/api/1.0/projects/%s/repos", projectKey), rData, rError)

	return rData.Values, err
}

func (bc *BitClient) GetUsers() ([]User, error) {
	rData := new(GetUsersResponse)
	rError := new(ErrorResponse)

	_, err := bc.DoGetPaged("/rest/api/1.0/users", rData, rError)

	return rData.Values, err
}

func (bc *BitClient) CreateRepository(projectKey string, data CreateRepositoryRequest) (*Repository, error) {
	rData := new(Repository)
	rError := new(ErrorResponse)

	_, err := bc.DoPost(fmt.Sprintf("/rest/api/1.0/projects/%s/repos", projectKey), data, rData, rError)

	return rData, err
}

func (bc *BitClient) SetRepositoryUserPermission(projectKey string, repositorySlug string, username string, permission string) error {

	rError := new(ErrorResponse)
	_, err := bc.DoPut(
		fmt.Sprintf(
			"/rest/api/1.0/projects/%s/repos/%s/permissions/users?name=%s&permission=%s",
			projectKey,
			repositorySlug,
			username,
			permission,
		),
		nil,
		nil,
		rError,
	)

	return err
}
