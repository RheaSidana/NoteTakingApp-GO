package users

type ErrorResponse struct {
	Message string
}

type UserResponse struct {
	Message string
}

type UserRequestWithSession struct {
	SID string
}
