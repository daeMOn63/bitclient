package bitclient

import (
	"fmt"
)

// Those branchmodel/configuration calls are not officially documented. Thanks to this comment for detailing usage:
// https://jira.atlassian.com/browse/BSERV-5411?focusedCommentId=1517096&page=com.atlassian.jira.plugin.system.issuetabpanels%3Acomment-tabpanel#comment-1517096
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
