package bitclient

// ##############################################################################
// Response
// ##############################################################################

type ErrorResponse struct {
	Errors []Error
}

type PagedResponse struct {
	Size       uint
	Limit      uint
	IsLastPage bool
	Start      uint
}

type GetProjectsResponse struct {
	PagedResponse
	Values []Project
}

type GetRepositoriesResponse struct {
	PagedResponse
	Values []Repository
}

type GetUsersResponse struct {
	PagedResponse
	Values []User
}
