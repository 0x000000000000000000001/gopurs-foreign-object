package Gen

import "math"

func Float32ToInt32(n float64) int {
	return int(int32(math.Float32bits(float32(n))))
}
