package bitclient

type GetMyPullRequestSuggestionsRequest struct {
	PagedRequest
	ChangeSince uint `url:"changeSince,omitempty"`
	Limit       uint `url:"limit,omitempty"`
}

type GetMyPullRequestSuggestionsResponse struct {
	PagedResponse
	Values []PullRequestSuggestion
}

func (bc *BitClient) GetMyPullRequestSuggestions(params GetMyPullRequestSuggestionsRequest) (GetMyPullRequestSuggestionsResponse, error) {

	response := GetMyPullRequestSuggestionsResponse{}

	_, err := bc.DoGet(
		"/dashboard/pull-request-suggestions",
		params,
		&response,
	)

	return response, err
}

type GetMyPullRequestsRequest struct {
	PagedRequest
	State             string `url:"state,omitempty"`
	Role              string `url:"role,omitempty"`
	ParticipantStatus string `url:"participantStatus,omitempty"`
	Order             string `url:"order,omitempty"`
	ClosedSince       string `url:"closedSince,omitempty"`
}

type GetMyPullRequestsResponse struct {
	PagedResponse PagedResponse
	PullRequests  []PullRequest
}

func (bc *BitClient) GetMyPullRequests(params GetMyPullRequestsRequest) (GetMyPullRequestsResponse, error) {

	response := GetMyPullRequestsResponse{}

	_, err := bc.DoGet(
		"/dashboard/pull-requests",
		params,
		&response,
	)

	return response, err
}
