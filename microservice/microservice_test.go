package microservice

import "testing"

func TestMicroservice_HelloWorld(t *testing.T) {
	uut := microservice{}

	cases := []struct {
		in, want string
	}{
		{"", "Hello, World"},
	}
	for _, c := range cases {
		got, _ := uut.HelloWorld()
		if got != c.want {
			t.Errorf("HelloWorld() == %q, want %q", got, c.want)
		}
	}
}