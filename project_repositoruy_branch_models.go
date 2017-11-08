package bitclient

import (
	"fmt"
)

func (bc *BitClient) GetBranchingModel(projectKey, repositorySlug string) (BranchingModel, error) {
	rError := new(ErrorResponse)

	response := BranchingModel{}

	url := fmt.Sprintf("/rest/branch-utils/1.0/projects/%s/repos/%s/branchmodel/configuration", projectKey, repositorySlug)
	resp, err := bc.sling.New().Get(url).Receive(&response, rError)

	resp, err = bc.checkReponse(resp, err)

	return response, err
}

func (bc *BitClient) SetBranchingModel(projectKey, repositorySlug string, settings BranchingModel) error {
	rError := new(ErrorResponse)

	url := fmt.Sprintf("/rest/branch-utils/1.0/projects/%s/repos/%s/branchmodel/configuration", projectKey, repositorySlug)
	resp, err := bc.sling.New().Put(url).BodyJSON(settings).Receive(nil, rError)

	resp, err = bc.checkReponse(resp, err)

	return err
}
