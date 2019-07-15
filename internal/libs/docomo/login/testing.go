package login

type ClientMock struct {
	LoginFunc func(id string, password string) (string, error)
	TestFunc  func(userID string, sessionKey string) (bool, error)
}

func (m *ClientMock) Login(id string, password string) (string, error) {
	return m.LoginFunc(id, password)
}
func (m *ClientMock) Test(userID string, sessionKey string) (bool, error) {
	return m.TestFunc(userID, sessionKey)
}
