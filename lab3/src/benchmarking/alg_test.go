package benchmarking

import (
	"math/rand"
	"testing"
	"time"
)

var (
	s1, s2 string
)

// Генерация случайной строки заданной длины
func randString(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Настройка тестовых данных
func init() {
	// Генерация данных один раз перед всеми тестами
	s1 = randString(2000)
	s2 = randString(2000)
}

// Бенчмарки для разных размеров входных данных
func BenchmarkLevenshtein100_100(b *testing.B)  { bench(b, 100,100) }
func BenchmarkLevenshtein400_100(b *testing.B)  { bench(b, 400, 100) }
func BenchmarkLevenshtein400_200(b *testing.B)  { bench(b, 400, 200) }
func BenchmarkLevenshtein800_400(b *testing.B)  { bench(b, 800, 400) }
func BenchmarkLevenshtein1600_200(b *testing.B)  { bench(b, 1600, 200) }

func bench(b *testing.B, sSize, tSize int) {
	s := []rune(s1[:sSize])
	t := []rune(s2[:tSize])
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		levenshteinDistance(s, t)
	}
}
