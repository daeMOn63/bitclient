package bitclient

import (
	"fmt"
)

func (bc *BitClient) GetSonarSettings(projectKey string, repositorySlug string) (SonarSettings, error) {
	rError := new(ErrorResponse)
	settings := SonarSettings{}

	url := fmt.Sprintf("/rest/sonar4stash/1.0/projects/%s/repos/%s/settings", projectKey, repositorySlug)
	resp, _ := bc.sling.New().Get(url).Receive(&settings, rError)

	resp, err := bc.checkReponse(resp, rError)

	return settings, err
}

func (bc *BitClient) SetSonarSettings(projectKey string, repositorySlug string, settings SonarSettings) error {
	rError := new(ErrorResponse)

	url := fmt.Sprintf("/rest/sonar4stash/1.0/projects/%s/repos/%s/settings", projectKey, repositorySlug)
	resp, _ := bc.sling.New().Post(url).BodyJSON(settings).Receive(nil, rError)

	resp, err := bc.checkReponse(resp, rError)

	return err
}
