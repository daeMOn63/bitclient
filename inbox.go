package bitclient

type GetInboxPullRequestsRequest struct {
	PagedRequest
	Start uint   `url:"start,omitempty"`
	Limit uint   `url:"limit,omitempty"`
	Role  string `url:"role,omitempty"`
}

type GetInboxPullRequestsResponse struct {
	PagedResponse
	Values []PullRequest
}

func (bc *BitClient) GetInboxPullRequests(params GetInboxPullRequestsRequest) (GetInboxPullRequestsResponse, error) {

	response := GetInboxPullRequestsResponse{}

	_, err := bc.DoGet(
		"/inbox/pull-requests",
		params,
		&response,
	)

	return response, err
}

type GetInboxPullRequestsCountResponse struct {
	Count uint `json:"count"`
}

func (bc *BitClient) GetInboxPullRequestsCount() (GetInboxPullRequestsCountResponse, error) {

	response := GetInboxPullRequestsCountResponse{}

	_, err := bc.DoGet(
		"/inbox/pull-requests/count",
		nil,
		&response,
	)

	return response, err
}
