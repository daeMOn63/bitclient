package bitclient

import (
	"fmt"
)

type SetRepositoryGroupPermissionRequest struct {
	Name       string `url:"name"`
	Permission string `url:"permission"`
}

func (bc *BitClient) SetRepositoryGroupPermission(projectKey string, repositorySlug string, params SetRepositoryGroupPermissionRequest) error {

	_, err := bc.DoPutUrl(
		fmt.Sprintf("/projects/%s/repos/%s/permissions/groups", projectKey, repositorySlug),
		params,
		nil,
	)

	return err
}

type GetRepositoryGroupPermissionRequest struct {
	PagedRequest
	Filter string
}

type GetRepositoryGroupPermissionResponse struct {
	PagedResponse
	Values []GroupPermission
}

func (bc *BitClient) GetRepositoryGroupPermission(projectKey string, repositorySlug string, params GetRepositoryGroupPermissionRequest) (GetRepositoryGroupPermissionResponse, error) {

	response := GetRepositoryGroupPermissionResponse{}

	_, err := bc.DoGet(
		fmt.Sprintf("/projects/%s/repos/%s/permissions/groups", projectKey, repositorySlug),
		params,
		&response,
	)

	return response, err
}

func (bc *BitClient) CloneRepositoryGroupPermissions(sourceProjectKey, sourceRepositorySlug, targetProjectKey, targetRepositorySlug string) error {

	response, err := bc.GetRepositoryGroupPermission(sourceProjectKey, sourceRepositorySlug, GetRepositoryGroupPermissionRequest{})
	if err != nil {
		return err
	}

	for _, permission := range response.Values {

		params := SetRepositoryGroupPermissionRequest{
			Name:       permission.Group.Name,
			Permission: permission.Permission,
		}
		err := bc.SetRepositoryGroupPermission(targetProjectKey, targetRepositorySlug, params)
		if err != nil {
			return fmt.Errorf("error adding %s permission to group %s on %s/%s: %s", permission.Permission, permission.Group.Name, targetProjectKey, targetRepositorySlug, err)
		}
	}

	return nil
}
