package custom_errors

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

type BusinessError struct {
	Message string
}

func (e BusinessError) Error() string {
	return e.Message
}

type ValidationError struct {
	Errors string
}

func (e ValidationError) Error() string {
	return e.Errors
}

type NotAuthorizedError struct {
}

func (e NotAuthorizedError) Error() string {
	return "You are not authorized."
}

type ForbiddenError struct {
}

func (e ForbiddenError) Error() string {
	return "Forbidden."
}
