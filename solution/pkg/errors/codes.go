package errors

type Code uint32

const (
	OK Code = iota
	BadRequest
	NotFound
	Internal
	Conflict
	Unauthorized
	Forbidden
)
