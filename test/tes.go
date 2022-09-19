package main

import "fmt"

func deretBilanganCase1() []int {
	// Contoh kasus pertama:
	// Inputan deret pertama = 2
	// Inputan deret kedua = 4
	// Value x = 5
	// Maka akan menghasilkan deret angka 2,4,6,8,10

	// contoh kodingan soal nomor 3 case pertama
	var bil1 int = 2
	var bil2 int = 4
	var x int = 5

	selisih := bil2 - bil1
	i := 0
	output := make([]int, x)

	for i < x {
		output[i] = bil1 + (i * selisih)
		i += 1
	}
	return output

}

func deretBilanganCase2() []int {
	// Contoh kasus kedua:
	// Inputan deret pertama = 5
	// Inputan deret kedua = 8
	// Value x = 7
	// Maka akan menghasilkan deret angka 5,8,11,14,17,20,23
	// contoh kodingan soal nomor 3 case pertama
	var bil1 int = 5
	var bil2 int = 8
	var x int = 7

	selisih := bil2 - bil1
	i := 0
	output := make([]int, x)

	for i < x {
		output[i] = bil1 + (i * selisih)
		i += 1
	}
	return output

}

func BubbleSortascending() []float32 {
	array := []float32{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				// array[j], array[j+1] = array[j+1], array[j]

				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
			}
		}
	}
	return array
}

func BubbleSortdescending() []float32 {
	array := []float32{4, -7, -5, 3, 3.3, 9, 0, 10, 0.2}
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] < array[j+1] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
			}
		}
	}
	return array
}

func main() {

	fmt.Println(deretBilanganCase1())
	fmt.Println(deretBilanganCase2())
	fmt.Println(BubbleSortascending())
	fmt.Println(BubbleSortdescending())

}
