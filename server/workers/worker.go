package workers

type Worker interface {
	Work() error
}
