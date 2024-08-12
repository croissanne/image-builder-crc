package unleash

import (
	"encoding/json"
	"net/http"

	"github.com/stretchr/testify/mock"
)

const (
	mockerServer   = "http://foo.com"
	mockAppName    = "unleash-client-go-tests"
	mockInstanceId = "1234"
)

type MockedListener struct {
	mock.Mock
}

func (l *MockedListener) OnError(err error) {
	l.Called(err)
}

func (l *MockedListener) OnWarning(warning error) {
	l.Called(warning)
}

func (l *MockedListener) OnReady() {
	l.Called()
}

func (l *MockedListener) OnCount(name string, enabled bool) {
	l.Called(name, enabled)
}

func (l *MockedListener) OnSent(payload MetricsData) {
	l.Called(payload)
}

func (l *MockedListener) OnRegistered(payload ClientData) {
	l.Called(payload)
}

func writeJSON(rw http.ResponseWriter, x interface{}) {
	enc := json.NewEncoder(rw)
	if err := enc.Encode(x); err != nil {
		panic(err)
	}
}
