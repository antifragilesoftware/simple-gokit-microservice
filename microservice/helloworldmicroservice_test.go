package microservice

import "testing"

func TestHelloWorld(t *testing.T) {
	service := helloWorldService{}

	cases := []struct {
		in, want string
	}{
		{"", "Hello, World"},
	}
	for _, c := range cases {
		got, _ := service.HelloWorld()
		if got != c.want {
			t.Errorf("HelloWorld() == %q, want %q", got, c.want)
		}
	}
}