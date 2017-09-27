package bitclient

import (
	"fmt"
)

func (bc *BitClient) GetSonarSettings(projectKey string, repositorySlug string) (SonarSettings, error) {

	params := SonarSettings{}

	_, err := bc.DoGet(
		fmt.Sprintf("/rest/sonar4stash/1.0/projects/%s/repos/%s/settings", projectKey, repositorySlug),
		nil,
		&params,
	)

	return params, err
}

func (bc *BitClient) SetSonarSettings(projectKey string, repositorySlug string, settings SonarSettings) error {

	_, err := bc.DoPost(
		fmt.Sprintf("/rest/sonar4stash/1.0/projects/%s/repos/%s/settings", projectKey, repositorySlug),
		settings,
		nil,
	)

	return err
}
