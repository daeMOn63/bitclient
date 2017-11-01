package bitclient

import (
	"fmt"
)

func (bc *BitClient) GetPullRequestSettings(projectKey string, repositorySlug string) (PullRequestSettings, error) {
	response := PullRequestSettings{}

	_, err := bc.DoGet(
		fmt.Sprintf("/projects/%s/repos/%s/settings/pull-requests", projectKey, repositorySlug),
		nil,
		&response,
	)

	return response, err
}

func (bc *BitClient) SetPullRequestSettings(projectKey string, repositorySlug string, params PullRequestSettings) error {

	// Remove disable mergeConfig strategies or they'll get enabled :/
	var filteredStrategies []PullRequestStrategy
	for _, strategy := range params.MergeConfig.Strategies {
		if strategy.Enabled == true {
			filteredStrategies = append(filteredStrategies, strategy)
		}
	}
	params.MergeConfig.Strategies = filteredStrategies

	_, err := bc.DoPost(
		fmt.Sprintf("/projects/%s/repos/%s/settings/pull-requests", projectKey, repositorySlug),
		params,
		nil,
	)

	return err
}
