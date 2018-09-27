package battle

// IStage represents ...
type IStage interface{}

// Battle represents ...
type Battle struct {
	Actors []IActor
	stage  IStage
}
