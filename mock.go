package goutils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockResponse(code int, body string) (*httptest.Server, *http.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	client := &http.Client{Transport: transport}

	return server, client
}

func TestSuccess(t *testing.T) {
	server, client := MockResponse(200, `{"upstream_url":"http:\/\/mock.url.for.path.mapper","id":"92e6a0fc-a962-4497-828a-95a9463e983","name":"wenchma.stage1.example.net","created_at":1464845908000,"request_host":"wenchma.stage1.example.net"}`)
	defer server.Close()
	req, _ = http.NewRequest("GET", "http://example.net"+path, nil)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	expectedResp := struct {
		Upstream_url, Id, Name, Request_host string
		Created_at                           int64
	}{"http://mock.url.for.path.mapper", "92e6a0fc-a962-4497-828a-95a9463e983", "wenchma.stage1.example.net", "wenchma.stage1.example.net", 1464845908000}

	resultBody := struct {
		Upstream_url, Id, Name, Request_host string
		Created_at                           int64
	}{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &resultBody)

	expect(t, resp.StatusCode, 200)
	expect(t, reflect.DeepEqual(expectedResp, resultBody), true)
}
