package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Pengembalian struct {
	gorm.Model
	NIM             int       `json:"nim" form:"nim" validate:"required"`
	Judul           string    `json:"judul" form:"judul"`
	Tanggal_kembali time.Time `json:"tanggal_kembali" form:"tanggal_kembali"`
}

// For Response Peminjaman
type PengembalianResponse struct {
	NIM             int    `json:"nim" form:"nim"`
	Judul           string `json:"judul" form:"judul"`
	Tanggal_kembali string `json:"tanggal_kembali" form:"tanggal_kembali"`
}
