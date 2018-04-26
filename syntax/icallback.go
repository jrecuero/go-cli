package syntax

// ICallback represents the interface for any command callback.
type ICallback interface {
	Enter(ctx *Context, arguments interface{}) error
	PostEnter(ctx *Context, arguments interface{}) error
	Exit(ctx *Context) error
	PostExit(ctx *Context) error
}
