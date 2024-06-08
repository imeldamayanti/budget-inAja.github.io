package main

import (
	"fmt"
	"time"
)

// Struktur untuk menyimpan data makanan
type Food struct {
	name     string
	price    float64
	rating   float64
	distance float64
	location string
}

// Fungsi untuk data makanan
func foodsData() []Food {
	return []Food{
		{"Ayam Gepuk Pak Gembus", 24, 4.4, 1.2, "https://maps.app.goo.gl/5Ss9Ry4gCaM1GV7z6"},
		{"Mie Ayam Baso Budi", 14, 4.8, 1, "https://maps.app.goo.gl/BtxLJo47FzSzrCmL9"},
		{"Batagor MSU", 10, 5, 1.3, "https://maps.app.goo.gl/Sx9945WJkQXs24UP7"},
		{"Sate Suramadu", 25, 4.3, 1.3, "https://maps.app.goo.gl/skXXHUmvV1WVZVXx8"},
		{"Ketoprak Mas No", 15, 4.9, 0.8, "https://maps.app.goo.gl/GMP4A8dHAc9pS9CL6"},
		{"Nasi Pecel Lele Lumintu 768", 15, 4.5, 1, "https://maps.app.goo.gl/DXgV4WZjcoSELyaAA"},
		{"Geprek Crispy Ayam Crisbar", 14, 4.4, 1.1, "https://maps.app.goo.gl/FcjSGXRb5xTKavMf6"},
		{"Mie Baso Mas Japra Solo", 15, 4.7, 1.2, "https://maps.app.goo.gl/nHSgDhB6JcXZHTVT8"},
		{"Mie Ayam Jabrig", 10, 4.5, 1.2, "https://maps.app.goo.gl/t9giWhD6E3MdT7o78"},
		{"Yakiniku Rice Bowl Spicy Yakiniku", 20, 4.8, 0.95, "https://maps.app.goo.gl/EgWp5aQSPKZvhNPS7"},
		{"Classic Beef Kebab Merhaba Kebab", 18, 5, 1.4, "https://maps.app.goo.gl/fpLQPxgiSeJDwdjx8"},
		{"Paket Jumbo Chicken Chicken William", 15, 4.3, 1.7, "https://maps.app.goo.gl/drA7Kznw118A6uiG6"},
		{"Paket Nila Merah Nasi Ikan Dan Ayam Bakar Pesona Bali", 26, 5, 3.1, "https://maps.app.goo.gl/BHxZu27AUfJAmxum7"},
		{"Seblak Azka", 15, 4, 1, "https://maps.app.goo.gl/3ithL5ePS8kyPU4NA"},
		{"Nasi Goreng Telor Kedai Aas", 12, 4.6, 1.4, "https://maps.app.goo.gl/2Lh7thn32hUgzbR69"},
		{"Kwetiau Goreng Seafood Kedai Aas", 14, 4.6, 1.4, "https://maps.app.goo.gl/2Lh7thn32hUgzbR69"},
		{"Nasi Goreng Moro Tresno", 13, 2.3, 0.9, "https://maps.app.goo.gl/4gkAkb4WL5AN5WAt6"},
		{"Mie Tek-Tek Warkop Djoeang", 10, 3, 1.1, "https://maps.app.goo.gl/RfnkhfRrsA6qjum49"},
		{"Dori Fish Mr. Mangkok", 18, 4.5, 1.8, "https://maps.app.goo.gl/hPWTHf8jFhcSNnK66"},
		{"Capcay Ayam Kedai Aas", 13, 4.6, 1.4, "https://maps.app.goo.gl/2Lh7thn32hUgzbR69"},
		{"Nasi Bakar Ayam Nasi Bakar Dapoer Koering", 16, 5, 1.1, "https://maps.app.goo.gl/xqBhbfdYghr3Ukzt5"},
		{"Nasi Bakar Cumi Nasi Bakar Dapoer Koering", 17, 5, 1.1, "https://maps.app.goo.gl/xqBhbfdYghr3Ukzt5"},
		{"Soto Betawi Daging Soto Betawi Pak Odang", 20, 4.7, 1.2, "https://maps.app.goo.gl/4YVTsCYyfR9gwXy67"},
		{"Soto Betawi Ayam Soto Betawi Pak Odang", 15, 4.7, 1.2, "https://maps.app.goo.gl/4YVTsCYyfR9gwXy67"},
		{"Mie Baek Asin / Manis", 12, 4.8, 1.3, "https://maps.app.goo.gl/yff8xUZwVR9QBsJs9"},
		{"Nasi Chicken Teriyaki WS Hotplate", 16, 4.5, 1.1, "https://maps.app.goo.gl/jHK47AawhdgoshrUA"},
		{"Mie Bakso Cincang Sedang Bakso Joko Tingkir", 14, 3.8, 1.3, "https://maps.app.goo.gl/tNhZectsrekuQzL79"},
		{"Mie Bakso Cincang Jumbo Bakso Joko Tingkir", 18, 3.8, 1.3, "https://maps.app.goo.gl/tNhZectsrekuQzL79"},
		{"Nasi Ayam Bistik Waroeng Steak 77", 20, 5, 0.95, "https://maps.app.goo.gl/hhYPM9cWkt5vJUKL9"},
		{"Steak Ayam Waroeng Steak 77", 25, 5, 0.95, "https://maps.app.goo.gl/hhYPM9cWkt5vJUKL9"},
	}
}

// Fungsi brute force untuk mencari kombinasi makanan terbaik
func bruteForce(foods []Food, budget float64, mealsPerWeek int) ([]Food, float64, int) {
	n := len(foods)
	maxRating := 0.0
	bestCombination := []Food{}
	totalCombinations := 0 // Menyimpan jumlah kombinasi yang telah diperiksa

	// Fungsi rekursif untuk memeriksa semua kombinasi
	var findCombination func(currentCombination []Food, currentBudget, currentRating float64, start int)
	findCombination = func(currentCombination []Food, currentBudget, currentRating float64, start int) {
		// Jika jumlah makanan dalam kombinasi sesuai kebutuhan
		if len(currentCombination) == mealsPerWeek {
			totalCombinations++
			if currentRating > maxRating {
				maxRating = currentRating
				bestCombination = make([]Food, mealsPerWeek)
				copy(bestCombination, currentCombination)
			}
			return
		}

		// Memeriksa semua kombinasi
		for i := start; i < n; i++ {
			totalCombinations++
			food := foods[i]
			if currentBudget+food.price <= budget {
				findCombination(append(currentCombination, food), currentBudget+food.price, currentRating+food.rating, i+1)
			}
		}
	}

	// Memulai pencarian dari kombinasi kosong
	findCombination([]Food{}, 0, 0, 0)

	return bestCombination, maxRating, totalCombinations
}

func main() {
	// Inisialisasi data makanan
	foods := foodsData()

	// Input dari user
	var budget float64
	var mealsPerDay int
	fmt.Println("Masukkan budget makan per minggu: ")
	fmt.Scan(&budget)
	fmt.Println("Masukkan berapa kali makan per hari (2-4 kali): ")
	fmt.Scan(&mealsPerDay)

	if mealsPerDay < 2 || mealsPerDay > 4 {
		fmt.Println("Jumlah kali makan per hari tidak valid.")
		return
	}

	// Hitung jumlah makanan yang dibutuhkan
	neededMeals := 7 * mealsPerDay

	// Mulai timer
	start := time.Now()

	// Mencari kombinasi makanan terbaik
	bestCombination, maxRating, totalCombinations := bruteForce(foods, budget, neededMeals)

	// Hitung waktu eksekusi
	elapsed := time.Since(start)

	// Periksa apakah ada rekomendasi makanan yang sesuai
	if len(bestCombination) < neededMeals {
		fmt.Println("Tidak ada rekomendasi makanan yang sesuai dengan budget Anda")
	} else {
		// Tampilkan rekomendasi makanan per hari
		fmt.Println("Rekomendasi makanan per hari:")
		for day := 0; day < 7; day++ {
			fmt.Printf("Hari %d:\n", day+1)
			for meal := 0; meal < mealsPerDay; meal++ {
				food := bestCombination[day*mealsPerDay+meal]
				fmt.Printf("- %s (Harga: %.2f, Rating: %.2f, Lokasi: %s)\n", food.name, food.price, food.rating, food.location)
			}
		}

		// Tampilkan total harga, total rating, dan jumlah kombinasi yang diperiksa
		totalPrice := 0.0
		for i := 0; i < neededMeals; i++ {
			totalPrice += bestCombination[i].price
		}
		fmt.Printf("Total harga: %.2f\n", totalPrice)
		fmt.Printf("Total rating: %.2f\n", maxRating)
		fmt.Printf("Jumlah kombinasi yang diperiksa: %d\n", totalCombinations)
	}

	// Tampilkan lama waktu running program
	fmt.Printf("Lama waktu running program: %s\n", elapsed)
}
