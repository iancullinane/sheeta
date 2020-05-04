package cloud

type cloud struct {
	r Resources
}

// Resources are API's needed to execute a task
type Resources struct {
	S3     S3Client
	CF     CFClient
	Logger Logger
}

// NewCloud returns a new cloud client
func NewCloud(r Resources) *cloud {
	return &cloud{
		r: r,
	}
}
