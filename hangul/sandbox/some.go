package sandbox

import "fmt"
import "github.com/myworld225/go/hangul"

func main() {
	fmt.Println("Hangul Test")
	var a = hangul.HasConsonantSuffix("테스트")
	var b = hangul.HasConsonantSuffix("입니깡")
	fmt.Println(a, b)
}
