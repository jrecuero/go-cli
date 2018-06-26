package syntax

import "fmt"

// Callback represents the type for any command callback.
type Callback struct{}

// Enter is the default callback, it executes with the running context.
func (c *Callback) Enter(ctx *Context, arguments interface{}) error {
	fmt.Println(">>>>> Default Enter")
	return nil
}

// PostEnter belongs only to mode commands and it executes with the new
// namespace context.
func (c *Callback) PostEnter(ctx *Context, arguments interface{}) error {
	return nil
}

// Exit belongs only to mode commands and it executes with the running
// namespace context.
func (c *Callback) Exit(ctx *Context) error {
	fmt.Println(">>>>> Default Exit")
	return nil
}

// PostExit belongs only to mode commands and it execute withe the parent
// namespace context, which will become the actual context.
func (c *Callback) PostExit(ctx *Context) error {
	return nil
}

var _ ICallback = (*Callback)(nil)
