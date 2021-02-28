package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Keys for environment with app prefix.
var (
	envDBFileName = fmt.Sprintf("%s_DB.FILENAME", prefixEnvironmet)
	envDBHost     = fmt.Sprintf("%s_DB.HOST", prefixEnvironmet)
	envDBPort     = fmt.Sprintf("%s_DB.PORT", prefixEnvironmet)
	envWebSrvHost = fmt.Sprintf("%s_WEBSRV.HOST", prefixEnvironmet)
	envWebSrvPort = fmt.Sprintf("%s_WEBSRV.PORT", prefixEnvironmet)
)

func TestConfig(t *testing.T) {
	type environment struct {
		DBFileName string
		DBHost     string
		DBPort     int32
		WebSrvHost string
		WebSrvPort int32
	}

	setEnvironvent := func(env environment) {
		// key - value
		os.Setenv(envDBFileName, env.DBFileName)
		os.Setenv(envDBHost, env.DBHost)
		os.Setenv(envDBPort, fmt.Sprint(env.DBPort))
		os.Setenv(envWebSrvHost, env.WebSrvHost)
		os.Setenv(envWebSrvPort, fmt.Sprint(env.WebSrvPort))
	}

	unsetEnvironment := func() {
		os.Unsetenv(envDBFileName)
		os.Unsetenv(envDBHost)
		os.Unsetenv(envDBPort)
		os.Unsetenv(envWebSrvHost)
		os.Unsetenv(envWebSrvPort)
	}

	testCases := []struct {
		name       string
		wantErr    bool
		useEnv     bool
		fileConfig string
		env        environment
		expect     *Config
	}{
		{
			name:       "From environment",
			wantErr:    false,
			useEnv:     true,
			fileConfig: "fixtures/test",
			env: environment{
				DBFileName: "db/other.db",
				DBHost:     "",
				DBPort:     3606,
				WebSrvHost: "localhost",
				WebSrvPort: 8080,
			},
			expect: &Config{
				DB: DBConfig{
					FileName: "db/other.db",
					Host:     "localhost",
					Port:     3606,
				},
				WebSrv: WebServerConfig{
					Host: "localhost",
					Port: 8080,
				},
			},
		},
		{
			name:       "From config file",
			wantErr:    false,
			useEnv:     false,
			fileConfig: "fixtures/test",
			env: environment{
				DBFileName: "db/other.db",
				DBHost:     "",
				DBPort:     3606,
				WebSrvHost: "localhost",
				WebSrvPort: 8080,
			},
			expect: &Config{
				DB: DBConfig{
					FileName: "db/development.db",
					Host:     "localhost",
					Port:     3606,
				},
				WebSrv: WebServerConfig{
					Host: "localhost",
					Port: 8080,
				},
			},
		},
	}

	for _, tc := range testCases {
		//t.Log(tc)

		// Setup environment from test case.
		if tc.useEnv {
			setEnvironvent(tc.env)
		} else {
			unsetEnvironment()
		}

		// Init config with config file.
		result, err := Init(tc.fileConfig)
		if tc.wantErr {
			assert.NotEmpty(t, err)
		} else {
			assert.Nil(t, err)
		}

		// Check.
		assert.Equal(t, tc.expect, result)
	}
}

func TestWebServerConfig_GetServerAddress(t *testing.T) {
	testCases := []struct {
		name            string
		webServerConfig WebServerConfig
		expect          string
	}{
		{
			name: "Have all: hostname and port",
			webServerConfig: WebServerConfig{
				Host: "localhost",
				Port: 3001,
			},
			expect: "localhost:3001",
		},
		{
			name: "Have only port",
			webServerConfig: WebServerConfig{
				Host: "",
				Port: 3001,
			},
			expect: ":3001",
		},
	}

	for _, tc := range testCases {
		result := tc.webServerConfig.GetServerAddress()
		assert.Equal(t, tc.expect, result)

	}
}
