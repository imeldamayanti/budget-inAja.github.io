package makananmodel

import (
	"KnapSack/config"
	"KnapSack/entities"
	"container/heap"
	"fmt"
	"sort"
	// "time"
	// "KnapSack/services"
	// "os"
)

type Node struct {
	level    int
	rating   float64
	price    float64
	bound    float64
	included []bool
	index    int // index diperlukan oleh heap interface
}

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
func getBound(node *Node, n int, budget float64, foods []entities.Makanan) float64 {
	if node.price >= budget {
		return 0
	}

	ratingBound := node.rating
	j := node.level + 1
	totalPrice := node.price

	for j < n && totalPrice+foods[j].Harga <= budget {
		totalPrice += foods[j].Harga
		ratingBound += foods[j].Rating
		j++
	}

	if j < n {
		ratingBound += (budget - totalPrice) * foods[j].Rating / foods[j].Harga
	}

	return ratingBound
}

// Fungsi untuk algoritma Branch and Bound
func branchAndBound(foods []entities.Makanan, budget float64, mealsPerWeek int) ([]int, int) {
	// Mengurutkan foods berdasarkan densitas (rating/price)
	sort.Slice(foods, func(i, j int) bool {
		return foods[i].Rating/foods[i].Harga > foods[j].Rating/foods[j].Harga
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
				v := &Node{level: u.level + 1, rating: u.rating + foods[u.level+1].Rating, price: u.price + foods[u.level+1].Harga, included: make([]bool, n)}
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
func findOptimalFoodCombination(foods []entities.Makanan, budget float64, mealsPerDay int) ([]int, int) {
	mealsPerWeek := mealsPerDay * 7
	optimalFoods, nodeCount := branchAndBound(foods, budget, mealsPerWeek)
	if len(optimalFoods) < mealsPerWeek {
		return nil, nodeCount
	}
	return optimalFoods, nodeCount
}

// Fungsi untuk memetakan makanan per hari
func mapFoodsPerDay(foods []entities.Makanan, foodIndices []int, mealsPerDay int) map[string][]entities.Makanan {
	weeklyMenu := make(map[string][]entities.Makanan)
	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	for i := 0; i < 7; i++ {
		for j := 0; j < mealsPerDay; j++ {
			weeklyMenu[days[i]] = append(weeklyMenu[days[i]], foods[foodIndices[i*mealsPerDay+j]])
		}
	}

	return weeklyMenu
}

func GetDayName(day int) string {
	var dayString string

	switch day {
	case 1:
		dayString = "SENIN"
	case 2:
		dayString = "SELASA"
	case 3:
		dayString = "RABU"
	case 4:
		dayString = "KAMIS"
	case 5:
		dayString = "JUMAT"
	case 6:
		dayString = "SABTU"
	case 7:
		dayString = "MINGGU"
	}

	return dayString
}

func GenerateData(mealsPerDay int, budget float64) []entities.Jadwal {
	rows, err := config.DB.Query(`SELECT * FROM makanan`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var foods []entities.Makanan
	for rows.Next() {
		var food entities.Makanan
		err := rows.Scan(&food.ID, &food.Nama, &food.Harga, &food.Rating, &food.Jarak, &food.Lokasi)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		foods = append(foods, food)
	}

	if mealsPerDay >= 5 {
		mealsPerDay = 5
	}

	if budget >= 1000 {
		budget = 1000
	}

	// Get meals recommendation
	optimalCombination, _ := findOptimalFoodCombination(foods, budget, mealsPerDay)

	// Filter combination per-day
	var totalSeluruhHarga float64
	mealPlans := make([]entities.Jadwal, 0)
	if optimalCombination != nil {
		weeklyMenu := mapFoodsPerDay(foods, optimalCombination, mealsPerDay)

		// Order of days
		days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

		for _, day := range days {
			meals := weeklyMenu[day]
			// Create empty slice jadwal
			jadwal := entities.Jadwal{
				Hari:       day,
				Menu:       make([]entities.Makanan, 0),
				TotalHarga: 0,
			}

			for _, food := range meals {
				jadwal.Menu = append(jadwal.Menu, food)
				jadwal.TotalHarga += food.Harga
				totalSeluruhHarga += food.Harga
			}

			mealPlans = append(mealPlans, jadwal)
		}
	}

	return mealPlans
}

