package aodata

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPrices(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/prices.json")
		if err != nil {
			t.Fatal(err)
		}
		w.Write(response)
	}))

	defer ts.Close()

	client := Client{
		BaseURL:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	prices, err := client.GetPrices("T4_2H_FIRESTAFF")
	require.NoError(t, err)
	assert.Len(t, prices, 3)
}
