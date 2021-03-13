package tests

import "github.com/go-testfixtures/testfixtures/v3"

// PrepareTestDatabase ...
func PrepareTestDatabase(fixtures *testfixtures.Loader) {
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}
