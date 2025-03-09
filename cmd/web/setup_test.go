package main

import (
	"os"
	"simple-go-web-page/adapters"
	"simple-go-web-page/configuration"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testBackend := &adapters.TestBackend{}
	testAdapter := &adapters.RemoteService{Remote: testBackend}
	testApp = application{
		App: configuration.New(nil, testAdapter),
	}

	os.Exit(m.Run())
}
