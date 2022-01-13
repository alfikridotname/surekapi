package kategoripenerima

type Tabler interface {
	TableName() string
}

type KategoriPenerima struct {
	ID         int    `json:"id"`
	Kategori   string `json:"kategori"`
	Keterangan string `json:"keterangan"`
	IsActive   bool   `json:"is_active"`
}

func (KategoriPenerima) TableName() string {
	return "ref_kategori_penerima"
}
