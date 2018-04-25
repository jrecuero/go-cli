package syntax

// Execute represents the method that executes a command callback.
func Execute(cmd *Command, arguments interface{}, ctx *Context) error {
	return cmd.Enter(ctx, arguments)
}
