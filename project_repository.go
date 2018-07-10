package bitclient

import (
	"fmt"
)

type UpdateRepositoryRequest struct {
	Name     string  `json:"name,omitempty"`
	Forkable bool    `json:"forkable,omitempty"`
	Project  Project `json:"project,omitempty"`
	Public   bool    `json:"public,omitempty"`
}

func (bc *BitClient) UpdateRepository(projectKey string, repositorySlug string, params UpdateRepositoryRequest) (Repository, error) {

	response := Repository{}

	url := fmt.Sprintf("/projects/%s/repos/%s", projectKey, repositorySlug)
	_, err := bc.DoPut(url, params, response)

	return response, err
}
