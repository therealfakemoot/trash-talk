package trash

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

// Mock accepts a string and returns alternating case string.
func Mock(s string) string {
	b := strings.Builder{}
	reg, _ := regexp.Compile("[^\\p{L}\\p{Z}]+")

	n := reg.ReplaceAllString(s, "")
	b.Grow(len(n))
	for i, c := range n {
		if i%2 == 0 {
			b.WriteString(strings.ToUpper(string(c)))
		} else {
			b.WriteString(strings.ToLower(string(c)))
		}
	}
	return b.String()
}

func main() {
	var s string
	flag.StringVar(&s, "in", "", "Any string can be passed in.")

	flag.Parse()

	// fmt.Println("Input string: `" + s + "`")

	fmt.Println(Mock(s))

}
