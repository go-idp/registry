package server

// Config is the configuration for the server
type Config struct {
	// Port is the port the server listens on
	Port int `config:"port"`

	// Registries is the list of registries
	Registries Registries `config:"registries"`
}

// Registries is the list of registries
type Registries struct {
	// Npm is the npm registry
	Npm Registry `config:"npm"`

	// Docker is the docker registry
	Docker Registry `config:"docker"`

	// Maven is the maven registry
	Maven Registry `config:"maven"`

	// Go is the go registry
	Go Registry `config:"go"`

	// Python is the python registry
	Python Registry `config:"python"`

	// Ruby is the ruby registry
	Ruby Registry `config:"ruby"`

	// Apt is the apt registry
	Apt Registry `config:"apt"`

	// Apk is the apk registry
	Apk Registry `config:"apk"`

	// Yum is the yum registry
	Yum Registry `config:"yum"`
}

// Registry is the registry
type Registry struct {
	// Server is the server
	Server string `config:"server"`

	// Headers is the headers
	Headers map[string]string `config:"headers"`
}
