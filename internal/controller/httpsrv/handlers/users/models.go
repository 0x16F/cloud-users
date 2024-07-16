package users

import "github.com/0x16F/cloud-users/internal/entity"

type UpdatePasswordReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UpdateUsernameReq struct {
	Username string `json:"username"`
}

type UpdateEmailReq struct {
	Email string `json:"email"`
}

type GetUsersReq struct {
	Limit    int    `query:"limit"`
	LastID   uint64 `query:"last_id"`
	Username string `query:"username"`
	Email    string `query:"email"`
}

type GetUsersResp struct {
	Users []entity.User `json:"users"`
}
