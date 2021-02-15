// +build integration

package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	integration "github.com/lfourky/go-rest-service-template/pkg/integration_tests"
	"github.com/lfourky/go-rest-service-template/pkg/model/dto"
	"github.com/lfourky/go-rest-service-template/pkg/server"
	"github.com/lfourky/go-rest-service-template/pkg/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	sqlSeedRegisterUserUserExists = `
		INSERT INTO "user" (id,name,email) 
		VALUES (
			'4468918f-cb92-426d-a0dc-b452b94a151c', 
			'Aron Aronson', 
			'aronaronson@someinvalidemaildomain.com'
		);
	`
)

func TestRegisterUser(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	smtpServer := integration.NewSMTPServer(config.Mail)
	smtpServerCleanup, err := smtpServer.Run()
	defer smtpServerCleanup()
	require.NoError(err)

	tests := []struct {
		name        string
		seeds       []string
		request     *dto.RegisterUserRequest
		mailHandler func(remoteAddr net.Addr, from string, to []string, data []byte) error
		assertions  func(req *dto.RegisterUserRequest, responseBody []byte)
	}{
		{
			name: "successfully registered a new user",
			request: &dto.RegisterUserRequest{
				Name:  "Aron Aronson",
				Email: "aronaronson@someinvalidemaildomain.com",
			},
			mailHandler: func(remoteAddr net.Addr, from string, to []string, data []byte) error {
				assert.Equal("aronaronson@someinvalidemaildomain.com", to[0])
				assert.Equal(config.Mail.SenderAddress, from)
				assert.Contains(string(data), fmt.Sprintf("Thanks for registering, %s", "Aron Aronson"))

				return nil
			},
			assertions: func(req *dto.RegisterUserRequest, responseBody []byte) {
				res := &dto.RegisterUserResponse{}
				err = json.Unmarshal(responseBody, res)
				require.NoError(err)

				assert.NotEmpty(res.ID)
				assert.Equal(res.Email, req.Email)
				assert.Equal(res.Email, req.Email)
			},
		},
		{
			name: "user not registered since one already exists with that email",
			request: &dto.RegisterUserRequest{
				Name:  "Aron Aronson",
				Email: "aronaronson@someinvalidemaildomain.com",
			},
			seeds: []string{sqlSeedRegisterUserUserExists},
			assertions: func(req *dto.RegisterUserRequest, responseBody []byte) {
				var res server.HTTPError
				err := json.Unmarshal(responseBody, &res)
				require.NoError(err)

				assert.Equal(usecase.ErrDatabaseUserAlreadyExists.Message, res.Message)
				assert.Equal(usecase.ErrDatabaseUserAlreadyExists.Code, res.Code)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			integration.ClearDatabase(config.PostgresConfig, t)
			integration.SeedDatabase(config.PostgresConfig, t, test.seeds...)

			smtpServer.SetHandler(test.mailHandler)

			body, err := json.Marshal(test.request)
			require.NoError(err)

			res, err := http.Post(address+"/v1/user", "application/json", bytes.NewBuffer(body))
			require.NoError(err)
			defer res.Body.Close()

			respBody, err := ioutil.ReadAll(res.Body)
			require.NoError(err)

			test.assertions(test.request, respBody)
		})
	}
}
