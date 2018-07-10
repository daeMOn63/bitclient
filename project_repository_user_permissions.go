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

type UnsetRepositoryUserPermissionRequest struct {
	Username string `url:"name"`
}

func (bc *BitClient) UnsetRepositoryUserPermission(projectKey string, repositorySlug string, params UnsetRepositoryUserPermissionRequest) error {

	_, err := bc.DoDeleteUrl(
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

func (bc *BitClient) CloneRepositoryUserPermissions(sourceProjectKey, sourceRepositorySlug, targetProjectKey, targetRepositorySlug string) error {

	response, err := bc.GetRepositoryUserPermission(sourceProjectKey, sourceRepositorySlug, GetRepositoryUserPermissionRequest{})
	if err != nil {
		return err
	}

	for _, permission := range response.Values {

		if permission.User.Active == false {
			continue
		}

		params := SetRepositoryUserPermissionRequest{
			Username:   permission.User.Name,
			Permission: permission.Permission,
		}
		err := bc.SetRepositoryUserPermission(targetProjectKey, targetRepositorySlug, params)
		if err != nil {
			return fmt.Errorf("error adding %s permission to user %s on %s/%s: %s", permission.Permission, permission.User.Name, targetProjectKey, targetRepositorySlug, err)
		}
	}

	return nil
}
