package api

// Context is the api Context
type Context struct {
	ApiURL      string
	Token       string
	Account     string
	Application string
}

// CreateAPI Creates an api instance from the Context
func (context *Context) CreateAPI() API {
	return API{
		context: *context,
	}
}

// Create a streaming api instance from the Context
