package bitclient

import (
	"fmt"
)

type GetHooksRequest struct {
	PagedRequest
	Type string
}

type GetHooksResponse struct {
	PagedResponse
	Values []Hook
}

func (bc *BitClient) GetHooks(projectKey, repositorySlug string, params GetHooksRequest) (GetHooksResponse, error) {
	response := GetHooksResponse{}

	_, err := bc.DoGet(
		fmt.Sprintf("/projects/%s/repos/%s/settings/hooks", projectKey, repositorySlug),
		params,
		&response,
	)

	return response, err
}

func (bc *BitClient) EnableHook(projectKey, repositorySlug, hookKey string, hookSettings interface{}) error {

	_, err := bc.DoPut(
		fmt.Sprintf("/projects/%s/repos/%s/settings/hooks/%s/enabled", projectKey, repositorySlug, hookKey),
		hookSettings,
		nil,
	)

	return err
}

func (bc *BitClient) DisableHook(projectKey, repositorySlug, hookKey string) error {
	_, err := bc.DoDeleteUrl(
		fmt.Sprintf("/projects/%s/repos/%s/settings/hooks/%s/enabled", projectKey, repositorySlug, hookKey),
		nil,
		nil,
	)

	return err
}
