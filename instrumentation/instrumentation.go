package instrumentation

type Instrumenter struct{
	projectId string
}

func NewInstrumenter(projectId string) *Instrumenter {
	return &Instrumenter{
		projectId: projectId,
	}
}