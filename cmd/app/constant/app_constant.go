package constant

type HeadersName int
type HeadersPrefix int

const (
	Authorization 		HeadersName = iota + 1
)

const (
	Bearer				HeadersPrefix = iota + 1
	Basic
)

func (h HeadersName) GetHeadersName() string {
	return [...] string{"Authorization"}[h-1]
}

func (h HeadersPrefix) GetHeadersPrefix() string {
	return [...] string{"Bearer ", "Basic "}[h-1]
}