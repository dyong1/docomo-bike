package login

type ClientMock struct {
	LoginFunc func(id string, password string) (string, error)
}

func (m *ClientMock) Login(id string, password string) (string, error) {
	return m.LoginFunc(id, password)
}
