package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_handleRequest(t *testing.T) {

	t.Run("responde com hello world", func(t *testing.T) {
		recoder := httptest.NewRecorder()
		reqTest := httptest.NewRequest("GET", "/", nil)

		handleRequest(recoder, reqTest)

		if recoder.Result().StatusCode != 200 {
			t.Errorf("Waiting %d, receive %d", 200, recoder.Result().StatusCode)
		}

		bodyTest, err := ioutil.ReadAll(recoder.Result().Body)

		if err != nil {
			t.Errorf("response Body with error")
		}

		htmlForm := `
		<form action="/formResult" method="POST">
			<input name="first_name" />
			<input type="submit" value="OK" />
		</form>
	`

		if string(bodyTest) != htmlForm {
			t.Errorf("Waiting %s, receive %s", "test", string(bodyTest))
		}
	})
}

func Test_handleResponse(t *testing.T) {
	t.Run("response form", func(t *testing.T) {
		firstName := strings.NewReader("first_name=test-name")

		recoder := httptest.NewRecorder()
		reqTest := httptest.NewRequest("POST", "/formResult", firstName)
		reqTest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		handleResponse(recoder, reqTest)

		if recoder.Result().StatusCode != 200 {
			t.Errorf("Waiting %d, receive %d", 200, recoder.Result().StatusCode)
		}

		bodyTest, err := ioutil.ReadAll(recoder.Result().Body)

		if err != nil {
			t.Errorf("response Body with error")
		}

		if string(bodyTest) != "test-name" {
			t.Errorf("Waiting %s, receive %s", "test-name", string(bodyTest))
		}
	})
}

func Test_authBasic(t *testing.T) {
	t.Run("test authBasic", func(t *testing.T) {

		recoder := httptest.NewRecorder()
		reqTest := httptest.NewRequest("GET", "/auth", nil)
		reqTest.Header.Set("Authorization", "Basic YWxhZGRpbjpvcGVuc2VzYW1lCg==")

		authBasic(recoder, reqTest)

		if recoder.Result().StatusCode != 200 {
			t.Errorf("Waiting %d, receive %d", 200, recoder.Result().StatusCode)
		}

		bodyTest, err := ioutil.ReadAll(recoder.Result().Body)
		if err != nil {
			t.Errorf("error reading from record.Result() - %v", err)
		}

		if string(bodyTest) != "OK" {
			t.Errorf("Waiting %s, returned %s", "OK", bodyTest)
		}
	})

	t.Run("test not auth", func(t *testing.T) {

		recoder := httptest.NewRecorder()
		reqTest := httptest.NewRequest("GET", "/auth", nil)
		reqTest.Header.Set("Authorization", "Basic b2k6dGNoYXU=")

		authBasic(recoder, reqTest)

		if recoder.Result().StatusCode != 401 {
			t.Errorf("Waiting %d, receive %d", 401, recoder.Result().StatusCode)
		}
	})
}
