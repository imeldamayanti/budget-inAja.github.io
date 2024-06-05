package makananmodel

import (
	"tubes_sa/config"
	"tubes_sa/entities"
)


func GetAll() []entities.Makanan {
	rows, err := config.DB.Query(`SELECT * FROM makanan`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var dt_makanan []entities.Makanan

	for rows.Next() {
		var makanan entities.Makanan
		if err := rows.Scan(&makanan.ID, &makanan.Nama, &makanan.Harga, &makanan.Rating, &makanan.Jarak, &makanan.Lokasi); err != nil {
			panic(err)
		}

		// fmt.Printf("ID: %d, Nama: %s, Harga: %.2f, Rating: %.2f, Jarak: %.2f, Lokasi: %s\n",
		// makanan.ID, makanan.Nama, makanan.Harga, makanan.Rating, makanan.Jarak, makanan.Lokasi)
		// fmt.Print(dt_makanan)
		dt_makanan = append(dt_makanan, makanan)
	}

	return dt_makanan
}
