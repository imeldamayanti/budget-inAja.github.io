package services

import (
	"KnapSack/entities"
)

// Struktur untuk menyimpan node dalam pohon ruang status
type Node struct {
	level    int
	rating   float64
	bound    float64
	harga    float64
	included []bool
}

// Mengurutkan data berdasarkan density
func MergeSort(foods []entities.Makanan) []entities.Makanan {
	if len(foods) > 1 {
		mid := len(foods) / 2

		// Allocate memory for the left and right subfoodsays
		L := make([]entities.Makanan, mid)
		R := make([]entities.Makanan, len(foods)-mid)

		// Copy data to the left and right subfoodsays
		for i := 0; i < mid; i++ {
			L[i] = foods[i]
		}
		for i := mid; i < len(foods); i++ {
			R[i-mid] = foods[i]
		}

		MergeSort(L)
		MergeSort(R)

		i, j, k := 0, 0, 0

		// Merge the two halves
		for i < len(L) && j < len(R) {
			if (L[i].Rating / L[i].Harga) >= (R[j].Rating / R[i].Harga) {
				foods[k] = L[i]
				i++
			} else {
				foods[k] = R[j]
				j++
			}
			k++
		}

		// Copy any remaining elements of L[], if any
		for i < len(L) {
			foods[k] = L[i]
			i++
			k++
		}

		// Copy any remaining elements of R[], if any
		for j < len(R) {
			foods[k] = R[j]
			j++
			k++
		}
	}
	return foods
}

// Fungsi untuk menghitung bound dari node
func bound(node Node, foods []entities.Makanan, budget float64, n_item int) float64 {
	if node.harga >= budget {
		return 0
	}

	ratingBound := node.rating
	j := node.level + 1
	totalWeight := node.harga

	for j < n_item && totalWeight+foods[j].Harga <= budget {
		totalWeight += foods[j].Harga
		ratingBound += foods[j].Rating
		j++
	}

	if j < n_item {
		ratingBound += (budget - totalWeight) * foods[j].Rating / foods[j].Harga
	}

	return ratingBound
}

// Fungsi untuk algoritma Branch and Bound
func branchAndBound(foods []entities.Makanan, budget float64) ([]entities.Makanan, float64, float64) {
	foods = MergeSort(foods)

	n := len(foods)
	var queue []Node
	u := Node{-1, 0.0, 0.0, 0.0, make([]bool, n)} // root
	v := Node{0, 0.0, 0.0, 0.0, make([]bool, n)}

	maxRating := 0.0
	bestItems := make([]bool, n)

	u.bound = bound(u, foods, budget, n)
	queue = append(queue, u)

	for len(queue) > 0 {
		u = queue[0]
		queue = queue[1:]

		if u.level == -1 {
			v.level = 0
		} else if u.level == n-1 {
			continue
		} else {
			v.level = u.level + 1
		}

		v.harga = u.harga + foods[v.level].Harga
		v.rating = u.rating + foods[v.level].Rating
		v.included = make([]bool, n)
		copy(v.included, u.included)
		v.included[v.level] = true

		if v.harga <= budget && v.rating > maxRating {
			maxRating = v.rating
			bestItems = make([]bool, n)
			copy(bestItems, v.included)
		}

		v.bound = bound(v, foods, budget, n)
		if v.bound > maxRating {
			queue = append(queue, v)
		}

		v = Node{u.level + 1, u.rating, u.harga, 0.0, make([]bool, n)}
		copy(v.included, u.included)
		v.bound = bound(v, foods, budget, n)
		if v.bound > maxRating {
			queue = append(queue, v)
		}
	}

	var result []entities.Makanan
	totalRating := 0.0
	totalHarga := 0.0
	for i := 0; i < n; i++ {
		if bestItems[i] && totalHarga+foods[i].Harga <= budget {
			result = append(result, foods[i])
			totalRating += foods[i].Rating
			totalHarga += foods[i].Harga
		}
	}

	return result, totalRating, totalHarga
}