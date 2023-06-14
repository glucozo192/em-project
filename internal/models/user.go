package models

type AddRoleRequest struct {
	Name string `json:"name" binding:"required"`
}
