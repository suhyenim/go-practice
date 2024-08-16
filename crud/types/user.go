package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type UserResponse struct {
	*ApiResponse
	*User
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"result"`
}

type CreateUserResquest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

func (c *CreateUserResquest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type CreateUserResponse struct {
	*ApiResponse
}

type UpdateUserResquest struct {
	Name       string `json:"name" binding:"required"`
	UpdatedAge int64  `json:"updatedAge" binding:"required"`
}

type UpdateUserResponse struct {
	*ApiResponse
}

type DeleteUserResquest struct {
	Name string `json:"name" binding:"required"`
}

type DeleteUserResponse struct {
	*ApiResponse
}
