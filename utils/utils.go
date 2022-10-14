package utils

import "math/rand"

// 0'dan 3'e kadar rastgele bir sayı ver
// Robotların kullanacağı yeteneği rastgele bir şekilde seçmek için
func random3() int {
	return rand.Intn(4)
}
