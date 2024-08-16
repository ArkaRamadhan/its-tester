package models

import (
	"time"

	"gorm.io/gorm"
)

// model for sag
type Sag struct {
	gorm.Model
	Tanggal time.Time `json:"tanggal"`
	NoMemo  string    `json:"no_memo"`
	Perihal string    `json:"perihal"`
	Pic     string    `json:"pic"`
}

// model for memo
type Memo struct {
	gorm.Model
	Tanggal time.Time `json:"tanggal"`
	NoMemo  string    `json:"no_memo"`
	Perihal string    `json:"perihal"`
	Pic     string    `json:"pic"`
}

// model for role
type Role struct {
	gorm.Model
	Name         string `json:"name"`
	description string `json:"description		"`
}

// model for iso
type Iso struct {
	gorm.Model
	Tanggal time.Time `json:"tanggal"`
	NoMemo  string    `json:"no_memo"`
	Perihal string    `json:"perihal"`
	Pic     string    `json:"pic"`
}

// model for surat
type Surat struct {
	gorm.Model
	Tanggal time.Time `json:"tanggal"`
	NoSurat  string    `json:"no_surat"`
	Perihal string    `json:"perihal"`
	Pic     string    `json:"pic"`
}

// model for berita acara
type BeritaAcara struct {
	gorm.Model
	Tanggal time.Time `json:"tanggal"`
	NoSurat string    `json:"no_surat"`
	Perihal string    `json:"perihal"`
	Pic     string    `json:"pic"`
}

// model for sk
type Sk struct {
	gorm.Model
	Tanggal time.Time `json:"tanggal"`
	NoSurat string    `json:"no_surat"`
	Perihal string    `json:"perihal"`
	Pic     string    `json:"pic"`
}

// model for project
type Project struct {
	gorm.Model
	KodeProject     string    `json:"kode_project"`
	JenisPengadaan  string    `json:"jenis_pengadaan"`
	NamaPengadaan   string    `json:"nama_pengadaan"`
	DivInisiasi     string    `json:"div_inisiasi"`
	Bulan           time.Time `json:"bulan"`
	SumberPendanaan string    `json:"sumber_pendanaan"`
	Anggaran        int64     `json:"anggaran"`
	NoIzin          string    `json:"no_izin"`
	TanggalIzin     time.Time `json:"tanggal_izin"`
	TanggalTor      time.Time `json:"tanggal_tor"`
	Pic             string    `json:"pic"`
}

// model for perdin
type Perdin struct {
	gorm.Model
	NoPerdin  string    `json:"no_perdin"`
	Tanggal   time.Time `json:"tanggal"`
	Hotel     string    `json:"hotel"`
	Transport string    `json:"transport"`
}

// model for suratMasuk
type SuratMasuk struct {
	gorm.Model
	NoSurat    string    `json:"no_surat"`
	Title      string    `json:"title"`
	RelatedDiv string    `json:"related_div"`
	DestinyDiv string    `json:"destiny_div"`
	Tanggal    time.Time `json:"tanggal"`
}

// model for suratKeluar
type SuratKeluar struct {
	gorm.Model
	NoSurat string    `json:"no_surat"`
	Title   string    `json:"title"`
	From    string    `json:"from"`
	Pic     string    `json:"pic"`
	Tanggal time.Time `json:"tanggal"`
}

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
}

type Admin struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
}