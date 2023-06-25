package hello

import "context"

//encore:api auth path=/hello/:name
func World(_ context.Context, name string) (*Response, error) {
	msg := "Hello, " + name + "!"
	return &Response{Message: msg}, nil
}

type Response struct {
	Message string
}
