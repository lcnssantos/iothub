package rmq

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RMQClient struct {
	host       string
	user       string
	password   string
	httpClient *http.Client
}

func NewRMQClient(host string, user string, password string) *RMQClient {
	return &RMQClient{host: host, user: user, password: password, httpClient: &http.Client{Timeout: time.Duration(20) * time.Second}}
}

type CreateUserRequest struct {
	Password string `json:"password"`
	Tags     string `json:"tags"`
}

func (c RMQClient) CreateAccount(login string, password string) error {
	endpoint := fmt.Sprintf("/api/users/%s", login)
	url := fmt.Sprintf("%s%s", c.host, endpoint)
	userRequest := CreateUserRequest{Password: password, Tags: ""}

	jsonReq, err := json.Marshal(userRequest)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonReq))
	req.SetBasicAuth(c.user, c.password)
	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	statusCode := strconv.Itoa(res.StatusCode)

	if strings.Split(statusCode, "")[0] != "2" {
		return errors.New(fmt.Sprintf("RMQ Api Error | Status: %s", statusCode))
	}

	return nil
}
