package instrumentation

type Instrumenter struct {
	projectID string
}

func NewInstrumenter(projectID string) *Instrumenter {
	return &Instrumenter{
		projectID: projectID,
	}
}
