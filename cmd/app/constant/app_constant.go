package constant

type HeadersName int
type HeadersPrefix int
type CacheName int
type ContextKey int
type Role int

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

const (
	JwtClaims				ContextKey = iota + 1
)

const (
	Admin Role = iota + 1
	Finance
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

func (c ContextKey) GetContextKey() string {
	return [...] string{"Jwt-Claims"}[c-1]
}

func (r Role) GetRole() string {
	return [...] string{"ROLE_ADMIN", "ROLE_FINANCE"}[r-1]
}