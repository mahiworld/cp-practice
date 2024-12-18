package main

// func canPlaceFlowers(flowerbed []int, n int) bool {
//     if n <= 0{
//         return true
//     }

//     length := len(flowerbed) - 1

// 	if (length ==0 && flowerbed[0] == 0) || (flowerbed[0] == 0 && flowerbed[1] == 0) {
// 		flowerbed[0] = 1
// 		n--
// 	}

//     if n <= 0{
//         return true
//     }

// 	if flowerbed[length] == 0 && flowerbed[length-1] == 0 {
// 		flowerbed[length] = 1
// 		n--
// 	}

//     if n <= 0{
//         return true
//     }

//     for i:=1; i<length-1;i++{
//         if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
// 		    flowerbed[i] = 1
// 		    n--
// 	    }
//     }

//     if n <= 0{
//         return true
//     }

//     return false
// }

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)

	for i := 0; i < length && n > 0; i++ {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == length-1 || flowerbed[i+1] == 0) {

			flowerbed[i] = 1
			n--
			i++
		}
	}

	return n <= 0
}
