package bitclient

import (
	"fmt"
)

type GetTagsResponse struct {
	PagedResponse
	Values []Tag
}

func (bc *BitClient) GetTags(projectKey string, repositorySlug string, params PagedRequest) (GetTagsResponse, error) {

	response := GetTagsResponse{}

	_, err := bc.DoGet(
		fmt.Sprintf("/projects/%s/repos/%s/tags", projectKey, repositorySlug),
		params,
		&response,
	)

	return response, err
}

type CreateTagRequest struct {
	Name       string `json:"name"`
	StartPoint string `json:"startPoint"`
	Message    string `json:"message"`
}

func (bc *BitClient) CreateTag(projectKey string, repositorySlug string, params CreateTagRequest) (Tag, error) {

	response := Tag{}

	_, err := bc.DoPost(
		fmt.Sprintf("/projects/%s/repos/%s/tags", projectKey, repositorySlug),
		params,
		&response,
	)

	return response, err
}
