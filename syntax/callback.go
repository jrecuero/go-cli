package syntax

// ICallback represents the interface for any command callback.
type ICallback interface {
	Enter(ctx *Context, arguments interface{}) bool
	PostEnter(ctx *Context, arguments interface{}) bool
	Exit(ctx *Context) bool
	PostExit(ctx *Context) bool
}

// Callback represents the type for any command callback.
type Callback struct{}

// Enter is the default callback, it executes with the running context.
func (c *Callback) Enter(ctx *Context, arguments interface{}) bool {
	return true
}

// PostEnter belongs only to mode commands and it executes with the new
// namespace context.
func (c *Callback) PostEnter(ctx *Context, arguments interface{}) bool {
	return true
}

// Exit belongs only to mode commands and it executes with the running
// namespace context.
func (c *Callback) Exit(ctx *Context) bool {
	return true
}

// PostExit belongs only to mode commands and it execute withe the parent
// namespace context, which will become the actual context.
func (c *Callback) PostExit(ctx *Context) bool {
	return true
}

var _ ICallback = (*Callback)(nil)
