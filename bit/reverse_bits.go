package bit

// import "fmt"

func reverseBits(num uint32) uint32 {
	var out uint32
	for i := 0; i < 32; i++ {
		num, out = num / 2, out * 2 + num % 2
	}
	return out
}
