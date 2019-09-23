package file

type LinkType int

const (
	Hard LinkType = iota
	Link
)

type Link struct {
	From   string
	To     string
	Hard   bool
	Broken bool
	File
}

func SymbolicLink(from, to string) *Link { return nil }
func HardLink(from, to string) *Link     { return nil }
