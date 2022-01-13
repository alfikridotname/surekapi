package naskahdinas

type Tabler interface {
	TableName() string
}

type NaskahDinas struct {
	ID       int    `json:"id"`
	Nama     string `json:"kategori"`
	Template string `json:"keterangan"`
	IsActive bool   `json:"is_active"`
}

func (NaskahDinas) TableName() string {
	return "ref_naskah_dinas"
}
