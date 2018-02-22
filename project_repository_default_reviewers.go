package bitclient

import (
	"fmt"
)

func (bc *BitClient) GetRepositoryDefaultReviewers(projectKey, repositorySlug string) ([]DefaultReviewers, error) {
	rError := new(ErrorResponse)

	var response []DefaultReviewers

	url := fmt.Sprintf("/rest/default-reviewers/1.0/projects/%s/repos/%s/conditions", projectKey, repositorySlug)
	resp, _ := bc.sling.New().Get(url).Receive(&response, rError)

	resp, err := bc.checkReponse(resp, rError)

	return response, err
}

func (bc *BitClient) CreateRepositoryDefaultReviewers(projectKey, repositorySlug string, params DefaultReviewers) (DefaultReviewers, error) {
	rError := new(ErrorResponse)

	request := RepositoryDefaultReviewersRequest{
		SourceMatcher:     params.FromRefMatcher,
		TargetMatcher:     params.ToRefMatcher,
		RequiredApprovals: params.RequiredApprovals,
		Reviewers:         params.Reviewers,
	}

	response := DefaultReviewers{}

	url := fmt.Sprintf("/rest/default-reviewers/latest/projects/%s/repos/%s/condition", projectKey, repositorySlug)
	resp, _ := bc.sling.New().Post(url).BodyJSON(request).Receive(&response, rError)

	resp, err := bc.checkReponse(resp, rError)

	return response, err
}

type RepositoryDefaultReviewersRequest struct {
	SourceMatcher     Matcher `json:"sourceMatcher,omitempty"`
	TargetMatcher     Matcher `json:"targetMatcher,omitempty"`
	RequiredApprovals int     `json:"requiredApprovals"`
	Reviewers         []User  `json:"reviewers,omitempty"`
}

func (bc *BitClient) UpdateRepositoryDefaultReviewers(projectKey, repositorySlug string, params DefaultReviewers) (DefaultReviewers, error) {
	rError := new(ErrorResponse)

	request := RepositoryDefaultReviewersRequest{
		SourceMatcher:     params.FromRefMatcher,
		TargetMatcher:     params.ToRefMatcher,
		RequiredApprovals: params.RequiredApprovals,
		Reviewers:         params.Reviewers,
	}

	response := DefaultReviewers{}

	url := fmt.Sprintf("/rest/default-reviewers/1.0/projects/%s/repos/%s/condition/%d", projectKey, repositorySlug, params.Id)
	resp, _ := bc.sling.New().Put(url).BodyJSON(request).Receive(&response, rError)

	resp, err := bc.checkReponse(resp, rError)

	return response, err
}
