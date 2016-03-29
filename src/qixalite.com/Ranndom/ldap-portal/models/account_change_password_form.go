package models

type AccountChangePasswordForm struct {
	CurrentPassword string `validate:"required" name:"Current password"`
	NewPassword     string `validate:"required" name:"New password"`
	ConfirmPassword string `validate:"required" name:"Confirmation password"`
}

