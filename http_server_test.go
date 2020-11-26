package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_handleRequest(t *testing.T) {

	t.Run("responde com hello world", func(t *testing.T) {
		recoder := httptest.NewRecorder()
		reqTest := httptest.NewRequest("GET", "/?key=test", nil)

		handleRequest(recoder, reqTest)

		if recoder.Result().StatusCode != 200 {
			t.Errorf("Waiting %d, receive %d", 200, recoder.Result().StatusCode)
		}

		bodyTest, err := ioutil.ReadAll(recoder.Result().Body)

		if err != nil {
			t.Errorf("response Body with error")
		}

		if string(bodyTest) != "test" {
			t.Errorf("Waiting %s, receive %s", "test", string(bodyTest))
		}
	})
}
