package bitclient

import (
	"fmt"
)

type SetRepositoryUserPermissionRequest struct {
	Username   string `url:"name"`
	Permission string `url:"permissions"`
}

func (bc *BitClient) SetRepositoryUserPermission(projectKey string, repositorySlug string, params SetRepositoryUserPermissionRequest) error {

	_, err := bc.DoPut(
		fmt.Sprintf("/projects/%s/repos/%s/permissions", projectKey, repositorySlug),
		params,
		nil,
	)

	return err
}
