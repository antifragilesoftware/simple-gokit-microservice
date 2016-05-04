package microservice

type Microservice interface {
	HelloWorld() (string, error)
}

type microservice struct{}

func (h microservice) HelloWorld() (string, error) {
	return "Hello, World", nil
}


