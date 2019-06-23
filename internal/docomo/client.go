package docomo

type Client interface {
	Login(id string, password string) (string, error)
}

type ScrappingClient struct{}
