package bitclient

import (
	"fmt"
)

type SetRepositoryUserPermissionRequest struct {
	Username   string `url:"name"`
	Permission string `url:"permission"`
}

func (bc *BitClient) SetRepositoryUserPermission(projectKey string, repositorySlug string, params SetRepositoryUserPermissionRequest) error {

	_, err := bc.DoPutUrl(
		fmt.Sprintf("/projects/%s/repos/%s/permissions/users", projectKey, repositorySlug),
		params,
		nil,
	)

	return err
}

type GetRepositoryUserPermissionRequest struct {
	PagedRequest
	Filter string
}

type GetRepositoryUserPermissionResponse struct {
	PagedResponse
	Values []UserPermission
}

func (bc *BitClient) GetRepositoryUserPermission(projectKey string, repositorySlug string, params GetRepositoryUserPermissionRequest) (GetRepositoryUserPermissionResponse, error) {

	response := GetRepositoryUserPermissionResponse{}

	_, err := bc.DoGet(
		fmt.Sprintf("/projects/%s/repos/%s/permissions/users", projectKey, repositorySlug),
		params,
		&response,
	)

	return response, err
}
