package bitclient

import (
	"fmt"
)

type GetRepositoryBranchRestrictionRequest struct {
	PagedRequest
	Type        string `url:"type,omitempty"`
	MatcherType string `url:"matcherType,omitempty"`
	MatcherId   string `url:"matcherId,omitempty"`
	Effective   bool   `url:"effective,omitempty"`
}

type GetRepositoryBranchRestrictionResponse struct {
	PagedResponse
	Values []BranchRestriction
}

func (bc *BitClient) GetRepositoryBranchRestrictions(projectKey, repositorySlug string, params GetRepositoryBranchRestrictionRequest) ([]BranchRestriction, error) {
	rError := new(ErrorResponse)

	response := GetRepositoryBranchRestrictionResponse{}

	url := fmt.Sprintf("/rest/branch-permissions/2.0/projects/%s/repos/%s/restrictions", projectKey, repositorySlug)
	resp, err := bc.sling.New().Get(url).QueryStruct(params).Receive(&response, rError)

	resp, err = bc.checkReponse(resp, err)

	return response.Values, err
}

type SetRepositoryBranchRestrictionsRequest struct {
	Id      int      `json:"id,omitempty"`
	Type    string   `json:"type,omitempty"`
	Matcher Matcher  `json:"matcher,omitempty"`
	Users   []string `json:"users,omitempty"`
	Groups  []string `json:"groups,omitempty"`
}

func (bc *BitClient) SetRepositoryBranchRestrictions(projectKey, repositorySlug string, params SetRepositoryBranchRestrictionsRequest) error {
	rError := new(ErrorResponse)

	url := fmt.Sprintf("/rest/branch-permissions/2.0/projects/%s/repos/%s/restrictions", projectKey, repositorySlug)
	resp, err := bc.sling.New().Post(url).BodyJSON(params).Receive(nil, rError)

	resp, err = bc.checkReponse(resp, err)

	return err
}

func (bc *BitClient) CloneRepositoryMasterBranchRestrictions(sourceProjectKey, sourceRepositorySlug, targetProjectKey, targetRepositorySlug string) error {

	getParams := GetRepositoryBranchRestrictionRequest{
		MatcherType: "BRANCH",
		MatcherId:   "refs/heads/master",
	}

	response, err := bc.GetRepositoryBranchRestrictions(sourceProjectKey, sourceRepositorySlug, getParams)
	if err != nil {
		return err
	}

	for _, branchRestriction := range response {

		var users []string
		for _, user := range branchRestriction.Users {
			if user.Active == true {
				users = append(users, user.Name)
			}
		}

		postParams := SetRepositoryBranchRestrictionsRequest{
			Id:      branchRestriction.Id,
			Type:    branchRestriction.Type,
			Matcher: branchRestriction.Matcher,
			Groups:  branchRestriction.Groups,
			Users:   users,
		}

		err := bc.SetRepositoryBranchRestrictions(targetProjectKey, targetRepositorySlug, postParams)
		if err != nil {
			return err
		}
	}

	return nil
}
