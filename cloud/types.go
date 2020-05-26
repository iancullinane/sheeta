package cloud

// Services are API's needed to execute module tasks
type Services struct {
	S3 S3Client
	CF CFClient
}

// StackConfig is used to generate the request
type StackConfig struct {
	Name        string                 `yaml:"name"`
	CloudConfig map[string]interface{} `yaml:"cloud-config"`
	Tags        map[string]string      `yaml:"tags"`
}
