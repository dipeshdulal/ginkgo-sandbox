package gomega_sandbox

type Books struct {
	Name   string
	Author string
	ISBN   string
}

func (b *Books) GetBookName() string {
	return b.Name
}
