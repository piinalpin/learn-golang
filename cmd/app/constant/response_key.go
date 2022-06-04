package constant

type ResponseKey int

const (
	Success			ResponseKey = iota + 1
	DataNotFound
	UnknownError
)

func (r ResponseKey) GetKey() string {
	return [...] string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR"}[r-1]
}

func (r ResponseKey) GetMessage() string {
	return [...] string{"Success", "Data Not Found", "Unknown Error"}[r-1]
}