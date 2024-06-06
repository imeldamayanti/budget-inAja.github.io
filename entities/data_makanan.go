package entities

type Jadwal struct {
	Hari       string
	Menu       []Makanan
	TotalHarga float64
}

type Makanan struct {
	ID     int
	Nama   string
	Harga  float64
	Rating float64
	Jarak  float64
	Lokasi string
}
