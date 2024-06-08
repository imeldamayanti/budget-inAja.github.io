package main

import (
	"container/heap"
	"fmt"
	"sort"
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

// DataFoods mengisi data makanan
func DataFoods() []Food {
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

// Struktur untuk menyimpan node dalam pohon ruang status
type Node struct {
	level    int
	rating   float64
	price    float64
	bound    float64
	included []bool
	index    int // index diperlukan oleh heap interface
}

// PriorityQueue mengimplementasikan heap.Interface dan menyimpan Nodes
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].bound > pq[j].bound // Menggunakan bound sebagai prioritas
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

// Fungsi untuk menghitung bound dari node
func getBound(node *Node, n int, budget float64, foods []Food) float64 {
	if node.price >= budget {
		return 0
	}

	ratingBound := node.rating
	j := node.level + 1
	totalPrice := node.price

	for j < n && totalPrice+foods[j].price <= budget {
		totalPrice += foods[j].price
		ratingBound += foods[j].rating
		j++
	}

	if j < n {
		ratingBound += (budget - totalPrice) * foods[j].rating / foods[j].price
	}

	return ratingBound
}

// Fungsi untuk algoritma Branch and Bound
func branchAndBound(foods []Food, budget float64, mealsPerWeek int) ([]int, int) {
	// Mengurutkan foods berdasarkan densitas (rating/price)
	sort.Slice(foods, func(i, j int) bool {
		return foods[i].rating/foods[i].price > foods[j].rating/foods[j].price
	})

	n := len(foods)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	u := &Node{level: -1, rating: 0, price: 0, included: make([]bool, n)}
	maxRating := 0.0
	var bestFoods []int
	nodeCount := 0

	u.bound = getBound(u, n, budget, foods)
	heap.Push(&pq, u)
	nodeCount++

	for pq.Len() > 0 {
		u = heap.Pop(&pq).(*Node)
		if u.bound > maxRating {
			if u.level+1 < n { // Ensure we don't go out of bounds
				// Menambahkan item ke dalam node
				v := &Node{level: u.level + 1, rating: u.rating + foods[u.level+1].rating, price: u.price + foods[u.level+1].price, included: make([]bool, n)}
				copy(v.included, u.included)
				v.included[u.level+1] = true
				nodeCount++

				if v.price <= budget && v.rating > maxRating && countTrue(v.included) == mealsPerWeek {
					maxRating = v.rating
					bestFoods = getIncludedItems(v.included)
				}

				v.bound = getBound(v, n, budget, foods)
				if v.bound > maxRating {
					heap.Push(&pq, v)
				}

				// Tidak menambahkan item ke dalam node
				v2 := &Node{level: u.level + 1, rating: u.rating, price: u.price, included: make([]bool, n)}
				copy(v2.included, u.included)
				nodeCount++
				v2.bound = getBound(v2, n, budget, foods)
				if v2.bound > maxRating {
					heap.Push(&pq, v2)
				}
			}
		}
	}

	return bestFoods, nodeCount
}

func countTrue(arr []bool) int {
	count := 0
	for _, v := range arr {
		if v {
			count++
		}
	}
	return count
}

func getIncludedItems(included []bool) []int {
	var items []int
	for i, v := range included {
		if v {
			items = append(items, i)
		}
	}
	return items
}

// Fungsi untuk menemukan kombinasi makanan optimal
func findOptimalFoodCombination(foods []Food, budget float64, mealsPerDay int) ([]int, int) {
	mealsPerWeek := mealsPerDay * 7
	optimalFoods, nodeCount := branchAndBound(foods, budget, mealsPerWeek)
	if len(optimalFoods) < mealsPerWeek {
		return nil, nodeCount
	}
	return optimalFoods, nodeCount
}

// Fungsi untuk memetakan makanan per hari
func mapFoodsPerDay(foods []Food, foodIndices []int, mealsPerDay int) map[string][]Food {
	weeklyMenu := make(map[string][]Food)
	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	for i := 0; i < 7; i++ {
		for j := 0; j < mealsPerDay; j++ {
			weeklyMenu[days[i]] = append(weeklyMenu[days[i]], foods[foodIndices[i*mealsPerDay+j]])
		}
	}

	return weeklyMenu
}

func main() {
	// Data makanan yang tersedia
	foods := DataFoods()

	// Input dari user
	var budget float64
	var mealsPerDay int

	fmt.Println("Masukkan budget makan per minggu: ")
	fmt.Scan(&budget)
	fmt.Println("Masukkan jumlah kali makan per hari (2-4): ")
	fmt.Scan(&mealsPerDay)

	// Validasi input jumlah kali makan per hari
	if mealsPerDay < 2 || mealsPerDay > 4 {
		fmt.Println("Jumlah kali makan per hari tidak valid.")
		return
	}

	startTime := time.Now()

	// Menemukan kombinasi makanan optimal
	optimalCombination, nodeCount := findOptimalFoodCombination(foods, budget, mealsPerDay)

	duration := time.Since(startTime)

	// Menampilkan rekomendasi makanan
	if optimalCombination == nil {
		fmt.Println("Tidak ada rekomendasi makanan yang sesuai dengan budget Anda.")
	} else {
		weeklyMenu := mapFoodsPerDay(foods, optimalCombination, mealsPerDay)
		totalPrice := 0.0
		totalRating := 0.0

		fmt.Println("Rekomendasi makanan per hari:")
		for day, meals := range weeklyMenu {
			fmt.Printf("%s:\n", day)
			for _, food := range meals {
				fmt.Printf("  - %s (Harga: %.2f, Rating: %.2f, Jarak: %.2f, Lokasi: %s)\n", food.name, food.price, food.rating, food.distance, food.location)
				totalPrice += food.price
				totalRating += food.rating
			}
		}
		fmt.Printf("\nTotal harga: %.2f\n", totalPrice)
		fmt.Printf("Total rating: %.2f\n", totalRating)
		fmt.Printf("Durasi running: %s\n", duration*100)
		fmt.Printf("Jumlah node yang terbentuk: %d\n", nodeCount)
	}
}
