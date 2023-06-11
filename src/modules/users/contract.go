package users

type ErrorResponse struct {
	Message string
}

type UserResponse struct{
	Message string
}

type UserResponseWithSession struct{
	SID string
}