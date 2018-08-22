package queue

// GeneratorContent represents ...
type GeneratorContent struct {
	label   string
	workout int
}

// GetLabel is ...
func (gc *GeneratorContent) GetLabel() string {
	return gc.label
}

// Generate is ...
func (gc *GeneratorContent) Generate() (interface{}, bool) {
	job := NewJob(gc.label, gc.workout)
	return job, true
}

// NewGeneratorContent is ...
func NewGeneratorContent(label string, workout int) *GeneratorContent {
	return &GeneratorContent{
		label:   label,
		workout: workout,
	}
}

// GetGeneratorContent is ...
func GetGeneratorContent(gen *Generator) *GeneratorContent {
	return gen.Content.(*GeneratorContent)
}
