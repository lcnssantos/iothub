package rmq

type RMQClient struct {
	host     string
	port     string
	user     string
	password string
}

func NewRMQClient(host string, port string, user string, password string) *RMQClient {
	return &RMQClient{host: host, port: port, user: user, password: password}
}

func (this RMQClient) CreateAccount(login string, password string) error {
	return nil
}
