package microservice

type HelloWorldService interface {
	HelloWorld() (string, error)
}

type helloWorldService struct{}

func (h helloWorldService) HelloWorld() (string, error) {
	return "Hello, World", nil
}


