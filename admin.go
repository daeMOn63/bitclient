package bitclient

type GetGroupUsersNonMemberRequest struct {
	PagedRequest
	Context string
	Filter  string
}

type GetGroupUsersNonMemberResponse struct {
	PagedResponse
	Values []DetailedUser
}

func (bc *BitClient) GetGroupUsersNonMember(params GetGroupUsersNonMemberRequest) (GetGroupUsersNonMemberResponse, error) {

	response := GetGroupUsersNonMemberResponse{}

	_, err := bc.DoGet(
		"/admin/groups/more-non-members",
		params,
		&response,
	)

	return response, err
}

type GetUserGroupsRequest struct {
	PagedRequest
	Context string
	Filter  string
}

type GetUserGroupsResponse struct {
	PagedResponse
	Values []Group
}

func (bc *BitClient) GetUserGroups(params GetUserGroupsRequest) (GetUserGroupsResponse, error) {

	response := GetUserGroupsResponse{}

	_, err := bc.DoGet(
		"/admin/users/more-members",
		params,
		&response,
	)

	return response, err
}

type GetUserGroupsNonMembersRequest struct {
	PagedRequest
	Context string
	Filter  string
}

type GetUserGroupsNonMembersResponse struct {
	PagedResponse
	Values []Group
}

func (bc *BitClient) GetUserGroupsNonMember(params GetUserGroupsNonMembersRequest) (GetUserGroupsNonMembersResponse, error) {

	response := GetUserGroupsNonMembersResponse{}

	_, err := bc.DoGet(
		"/admin/users/more-non-members",
		params,
		&response,
	)

	return response, err
}

type SearchUsersRequest struct {
	PagedRequest,
	Filter string
}

type SearchUsersResponse struct {
	PagedResponse
	Values []DetailedUser
}

func (bc *BitClient) SearchUsers(params SearchUsersRequest) (SearchUsersResponse, error) {

	response := SearchUsersResponse{}

	_, err := bc.DoGet(
		"/admin/users",
		params,
		&response,
	)

	return response, err
}

type CreateUserRequest struct {
	Name              string
	Password          string
	DisplayName       string
	EmailAddress      string
	AddToDefaultGroup bool
	Notify            bool
}

func (bc *BitClient) CreateUser(params CreateUserRequest) error {

	_, err := bc.DoPostUrl(
		"/admin/users",
		params,
		nil,
	)

	return err
}

type DeleteUserRequest struct {
	Name string
}

func (bc *BitClient) DeleteUser(params DeleteUserRequest) (DetailedUser, error) {

	response := DetailedUser{}

	_, err := bc.DoDeleteUrl(
		"/admin/users",
		params,
		&response,
	)

	return response, err
}

type UpdateUserRequest struct {
	Name        string `json:"name"`
	DisplayName string `jsin:"displayName"`
	Email       string `json:"email"`
}

func (bc *BitClient) UpdateUser(params UpdateUserRequest) (DetailedUser, error) {

	response := DetailedUser{}

	_, err := bc.DoPut(
		"/admin/users",
		params,
		&response,
	)

	return response, err
}

type RenameUserRequest struct {
	Name    string `json:"name"`
	NewName string `json:"newName"`
}

func (bc *BitClient) RenameUser(params RenameUserRequest) (DetailedUser, error) {

	response := DetailedUser{}

	_, err := bc.DoPost(
		"/admin/users/rename",
		params,
		&response,
	)

	return response, err
}

type UpdateUserPasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
	Name        string `json:"name"`
}

func (bc *BitClient) UpdateUserPassword(params UpdateUserPasswordRequest) error {

	_, err := bc.DoPut(
		"/admin/users/credentials",
		params,
		nil,
	)

	return err
}
