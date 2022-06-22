package respkey

type ResponseKey int

const (
	Success			ResponseKey = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
)

func (r ResponseKey) GetKey() string {
	return [...] string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}[r-1]
}

func (r ResponseKey) GetMessage() string {
	return [...] string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized"}[r-1]
}