package constant

type HeadersName int
type HeadersPrefix int
type CacheName int

const (
	Authorization 		HeadersName = iota + 1
)

const (
	Bearer				HeadersPrefix = iota + 1
	Basic
)

const (
	UserSession				CacheName = iota + 1
	RefreshTokenSession
)

func (h HeadersName) GetHeadersName() string {
	return [...] string{"Authorization"}[h-1]
}

func (h HeadersPrefix) GetHeadersPrefix() string {
	return [...] string{"Bearer ", "Basic "}[h-1]
}

func (c CacheName) GetCacheName() string {
	return [...] string{"USER_SESSION", "REFRESH_TOKEN_SESSION"}[c-1]
}