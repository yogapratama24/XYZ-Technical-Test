package model

import "time"

type Transaction struct {
	Id            int       `json:"id" gorm:"primaryKey"`
	NomorKontrak  string    `json:"nomor_kontrak"`
	Otr           string    `json:"otr"`
	AdminFee      int64     `json:"admin_fee"`
	JumlahCicilan int64     `json:"jumlah_cicilan"`
	JumlahBunga   float64   `json:"jumlah_bunga"`
	NamaAsset     string    `json:"nama_asset"`
	UserId        int       `json:"user_id"`
	User          *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type CreateTransactionRequest struct {
	NomorKontrak  string  `json:"nomor_kontrak" validate:"required"`
	Otr           string  `json:"otr" validate:"required"`
	AdminFee      int64   `json:"admin_fee" validate:"required"`
	JumlahCicilan int64   `json:"jumlah_cicilan" validate:"required"`
	JumlahBunga   float64 `json:"jumlah_bunga" validate:"required"`
	NamaAsset     string  `json:"nama_asset" validate:"required"`
	UserId        int     `json:"user_id" validate:"required"`
}

type ReadTransactionResponse struct {
	Id            int             `json:"id" gorm:"primaryKey"`
	NomorKontrak  string          `json:"nomor_kontrak"`
	Otr           string          `json:"otr"`
	AdminFee      int64           `json:"admin_fee"`
	JumlahCicilan int64           `json:"jumlah_cicilan"`
	JumlahBunga   float64         `json:"jumlah_bunga"`
	NamaAsset     string          `json:"nama_asset"`
	UserId        int             `json:"user_id"`
	User          UserTransaction `json:"user"`
	CreatedAt     time.Time       `json:"created_at,omitempty"`
	UpdatedAt     time.Time       `json:"updated_at,omitempty"`
}

type UserTransaction struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
}
