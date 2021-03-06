package scheduler

type App struct {
	AppId  string
	Tasks  []Task
	Labels map[string]string
	Env    map[string]string
}

type Task struct {
	Host         string
	Ports        []int
	ServicePorts []int
	StagedAt     string
	StartedAt    string
	Version      string
}

type BeethovenInstance struct {
	Host string
	Port int
}

type ShutdownChan chan bool

type Scheduler interface {
	// Watch for changes using streams and make callbacks to the specified
	// handler when apps have been added, removed or health changes.
	Watch(reload chan bool)
	// Shutdown the current stream watching
	Shutdown()
	// FetchApps will find all applications/services from the scheduler source
	FetchApps() (map[string]*App, error)

	// FetchBeethovenInstances will find all beethoven running instances
	FetchBeethovenInstances() ([]*BeethovenInstance, error)
}
