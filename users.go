package bitclient

type GetUsersResponse struct {
	PagedResponse
	Values []User
}

func (bc *BitClient) GetUsers(params PagedRequest) (GetUsersResponse, error) {

	response := GetUsersResponse{}

	_, err := bc.DoGet(
		"/users",
		params,
		&response,
	)

	return response, err
}
