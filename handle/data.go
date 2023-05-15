package handle

import "math/rand"

func GenRandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < length; i++ {
		result = append(result, bytes[RandInt(0, len(bytes))])
	}
	return string(result)
}

func RandInt(min, max int) int {
	if min >= max {
		return min
	}
	return min + rand.Intn(max-min)
}

func RandFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func Round(data float64, precision int) float64 {
	p := 1
	for i := 0; i < precision; i++ {
		p *= 10
	}
	return float64(int(data*float64(p))) / float64(p)
}

func AdjustNum(data, unit float64) float64 {
	return 0
}
