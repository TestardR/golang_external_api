package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"example.com/clients/restclient"
	"example.com/domain/repositories"
	"github.com/stretchr/testify/assert"
)

var (
	urlCreateRepo = "https://api.github.com/user/repos"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidIputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}
func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockUp(restclient.Mock{
		Url:        urlCreateRepo,
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{
				"message": "Requires authentication",
				"documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"
			}`)),
		},
	})
	request := repositories.CreateRepoRequest{Name: "testing"}
	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockUp(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{
				"id": 123,
				"name": "Romain",
				"full_name": "Golang-Tut"
			}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}
	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "Romain", result.Name)
	assert.EqualValues(t, "", result.Owner)
}
