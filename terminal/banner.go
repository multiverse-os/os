package terminal

import (
	"fmt"
)

func PrintBanner(name, version string) {
	fmt.Println(Header(name) + "  v" + Light(version))
	fmt.Println(Repeat("=", 80))
}
