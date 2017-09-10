package bitclient

type GetPullRequestSuggestionsRequest struct {
	PagedRequest
	ChangeSince uint `url:"changeSince,omitempty"`
	Limit       uint `url:"limit,omitempty"`
}

type GetPullRequestSuggestionsResponse struct {
	PagedResponse
	Values []PullRequestSuggestion
}

func (bc *BitClient) GetPullRequestSuggestions(params GetPullRequestSuggestionsRequest) (GetPullRequestSuggestionsResponse, error) {

	response := GetPullRequestSuggestionsResponse{}

	_, err := bc.DoGet(
		"/dashboard/pull-request-suggestions",
		params,
		&response,
	)

	return response, err
}

type GetPullRequestsRequest struct {
	PagedRequest
	State             string `url:"state,omitempty"`
	Role              string `url:"role,omitempty"`
	ParticipantStatus string `url:"participantStatus,omitempty"`
	Order             string `url:"order,omitempty"`
	ClosedSince       string `url:"closedSince,omitempty"`
}

type GetPullRequestsResponse struct {
	PagedResponse PagedResponse
	PullRequests  []PullRequest
}

func (bc *BitClient) GetPullRequests(params GetPullRequestsRequest) (GetPullRequestsResponse, error) {

	response := GetPullRequestsResponse{}

	_, err := bc.DoGet(
		"/dashboard/pull-requests",
		params,
		&response,
	)

	return response, err
}
