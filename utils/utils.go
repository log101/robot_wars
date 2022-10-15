package utils

import "math/rand"

// 0'dan 3'e kadar rastgele bir sayı ver
// Robotların kullanacağı yeteneği rastgele bir şekilde seçmek için
func Random3() int {
	return rand.Intn(3)
}
