package syntax

// Execute represents the method that executes a command callback.
func Execute(cmd *Command, arguments []interface{}, context interface{}) bool {
	return cmd.Cb(context, arguments)
}
