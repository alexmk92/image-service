// +build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetImages(t *testing.T) {
    client := resty.New()
    resp, err := client.R().Get(BASE_URL + "/api/image")
    if err != nil {
        t.Fail()
    }

    assert.Equal(t, 200, resp.StatusCode())
}

func TestPostImage(t *testing.T) {
    client := resty.New()
    resp, err := client.R().
        SetBody(`{"AspectRatio":"1:1","StorageUrl":"/"}`).
        Post(BASE_URL + "/api/image")

    assert.NoError(t, err)

    assert.Equal(t, 200, resp.StatusCode())
}
