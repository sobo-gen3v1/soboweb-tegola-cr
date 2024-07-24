package server_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/sobo-gen3v1/soboweb-tegola-cr/server"
)

func TestHandleMasterLayers(t *testing.T) {
	type tcase struct {
		hostname string
		port     string
		uri      string
		method   string
		body     []byte
		expected any
	}

	fn := func(tc tcase) func(*testing.T) {
		return func(t *testing.T) {
			var err error

			server.HostName = tc.hostname
			server.Port = tc.port
			a := newTestMapWithLayers(testLayer1, testLayer2, testLayer3)

			w, _, err := doRequest(a, tc.method, tc.uri, bytes.NewBuffer(tc.body))
			if err != nil {
				t.Fatal(err)
			}

			if w.Code != http.StatusOK {
				t.Errorf("status code, expected %v got %v", http.StatusOK, w.Code)
				return
			}

			bytes, err := ioutil.ReadAll(w.Body)
			if err != nil {
				t.Errorf("error response body, expected nil got %v", err)
				return
			}

			var result map[string]any

			// read the respons body
			if err := json.Unmarshal(bytes, &result); err != nil {
				t.Errorf("error unmarshal JSON, expected nil got %v", err)
				return
			}

			t.Logf("result = %+v", result)

			// if !reflect.DeepEqual(tc.expected, layers) {
			// 	t.Errorf("response body, \n  expected %+v\n  got %+v", tc.expected, layers)
			// }

		}
	}

	t1 := map[string]string{
		"area_code":  "01204",
		"properties": "",
		"tags":       "",
	}
	t2 := map[string]string{
		"area_code":  "01462",
		"properties": "",
		"tags":       "",
	}
	var layers []map[string]string
	layers = append(layers, t1)
	layers = append(layers, t2)
	version := &map[string]string{
		"version":    "1721189825856",
		"start_time": "2024-07-17T12:00:00+08:00",
		"end_time":   "2024-07-17T12:17:00+08:00",
	}

	bodyJson, _ := json.Marshal(map[string]any{
		"layers":  layers,
		"version": version,
	})

	tests := map[string]tcase{
		"empty host port": {
			//  With empty hostname and no port specified in config, urls should have host:port matching request uri.
			uri:      "http://localhost:8181/api/m/layers",
			method:   "PUT",
			body:     bodyJson,
			expected: make([]interface{}, 0),
		},
	}

	for name, tc := range tests {
		t.Run(name, fn(tc))
	}
}
