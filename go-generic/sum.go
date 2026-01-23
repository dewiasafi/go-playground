package main

func Sum[V int | float32 | float64](numbers []V) V {
	var total V
	for _, e := range numbers {
		total += e
	}
	return total
}
