package model

import "time"

type User struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	Nik          string    `json:"nik" gorm:"unique"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir string    `json:"tanggal_lahir"`
	Gaji         int64     `json:"gaji"`
	FotoKtp      string    `json:"foto_ktp"`
	FotoSelfie   string    `json:"foto_selfie"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type UserCreateRequest struct {
	Nik          string `json:"nik" form:"nik" validate:"required"`
	FullName     string `json:"full_name" form:"full_name" validate:"required"`
	LegalName    string `json:"legal_name" form:"legal_name" validate:"required"`
	TempatLahir  string `json:"tempat_lahir" form:"tempat_lahir" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" form:"tanggal_lahir" validate:"required"`
	Gaji         int64  `json:"gaji" form:"gaji" validate:"required"`
	FotoKtp      string `json:"foto_ktp"`
	FotoSelfie   string `json:"foto_selfie"`
	Password     string `json:"password" form:"password" validate:"required"`
}

type UserLoginRequest struct {
	Nik      string `json:"nik" form:"nik" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
