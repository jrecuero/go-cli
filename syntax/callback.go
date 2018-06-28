package syntax

// EnterFunc represents generic callback for enter functions.
type EnterFunc = func(ctx *Context, arguments interface{}) error

// ExitFunc represents generic callback for exit functions.
type ExitFunc = func(ctx *Context) error

// Callback represents the type for any command callback.
type Callback struct {
	Enter     EnterFunc
	PostEnter EnterFunc
	Exit      ExitFunc
	PostExit  ExitFunc
}

// SetEnter sets new value for Enter field.
func (cb *Callback) SetEnter(enter EnterFunc) *Callback {
	cb.Enter = enter
	return cb
}

// SetExit sets new value for Exit field.
func (cb *Callback) SetExit(exit ExitFunc) *Callback {
	cb.Exit = exit
	return cb
}

// defaultEnter is the default callback, it executes with the running context.
func defaultEnter(ctx *Context, arguments interface{}) error {
	return nil
}

// defaultPostEnter belongs only to mode commands and it executes with the new
// namespace context.
func defaultPostEnter(ctx *Context, arguments interface{}) error {
	return nil
}

// defaultExit belongs only to mode commands and it executes with the running
// namespace context.
func defaultExit(ctx *Context) error {
	return nil
}

// defaultPostExit belongs only to mode commands and it execute withe the parent
// namespace context, which will become the actual context.
func defaultPostExit(ctx *Context) error {
	return nil
}

// NewCallback creates a new Callback instance.
func NewCallback(enterCb EnterFunc, postEnterCb EnterFunc, exitCb ExitFunc, postExitCb ExitFunc) *Callback {
	callback := &Callback{
		Enter:     defaultEnter,
		PostEnter: defaultPostEnter,
		Exit:      defaultExit,
		PostExit:  defaultPostExit,
	}
	if enterCb != nil {
		callback.Enter = enterCb
	}
	if postEnterCb != nil {
		callback.PostEnter = postEnterCb
	}
	if exitCb != nil {
		callback.Exit = exitCb
	}
	if postExitCb != nil {
		callback.PostExit = postExitCb
	}
	return callback
}

// NewDefaultCallback creates a new Callback instance with default methods.
func NewDefaultCallback() *Callback {
	return NewCallback(nil, nil, nil, nil)
}
