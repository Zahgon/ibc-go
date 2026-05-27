package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"os"
)

const (
	testNamePrefix     = "Test"
	testFileNameSuffix = "_test.go"
	e2eTestDirectory   = "e2e"
	// testEntryPointEnv specifies a single test function to run if provided.
	testEntryPointEnv = "TEST_ENTRYPOINT"
	// testExclusionsEnv is a comma separated list of test function names that will not be included
	// in the results of this script.
	testExclusionsEnv = "TEST_EXCLUSIONS"
	// testNameEnv if provided returns a single test entry so that only one test is actually run.
	testNameEnv = "TEST_NAME"
)

// GithubActionTestMatrix represents
type GithubActionTestMatrix struct {
	Include []TestSuitePair `json:"include"`
}

type TestSuitePair struct {
	Test       string `json:"test"`
	EntryPoint string `json:"entrypoint"`
}

func main() {
	githubActionMatrix, err := getGithubActionMatrixForTests(e2eTestDirectory, getTestToRun(), getTestEntrypointToRun(), getExcludedTestFunctions())
	if err != nil {
		fmt.Printf("error generating github action json: %s", err)
		os.Exit(1)
	}

	ghBytes, err := json.Marshal(githubActionMatrix)
	if err != nil {
		fmt.Printf("error marshalling github action json: %s", err)
		os.Exit(1)
	}
	fmt.Println(string(ghBytes))
}

// getTestEntrypointToRun returns the specified test function to run if present, otherwise
// it returns an empty string which will result in running all test suites.
func getTestEntrypointToRun() string { _ = "STUB: not implemented"; return "" }

// getTestToRun returns the specified test function to run if present.
// If specified, only this test will be run.
func getTestToRun() string { _ = "STUB: not implemented"; return "" }

// getExcludedTestFunctions returns a list of test functions that we don't want to run.
func getExcludedTestFunctions() []string { _ = "STUB: not implemented"; return nil }

// getGithubActionMatrixForTests returns a json string representing the contents that should go in the matrix
// field in a github action workflow. This string can be used with `fromJSON(str)` to dynamically build
// the workflow matrix to include all E2E tests under the e2eRootDirectory directory.
func getGithubActionMatrixForTests(e2eRootDirectory, testName string, suite string, excludedItems []string) (GithubActionTestMatrix, error) {
	_ = "STUB: not implemented"
	return *new(GithubActionTestMatrix), nil
}

// only look at test files

// Sort the test cases by name so that the order is consistent.

// extractSuiteAndTestNames extracts the name of the test suite function as well
// as all tests associated with it in the same file.
func extractSuiteAndTestNames(file *ast.File) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// isTestSuiteMethod returns true if the function is a test suite function.
// e.g. func TestFeeMiddlewareTestSuite(t *testing.T) { ... }
func isTestSuiteMethod(f *ast.FuncDecl) bool { _ = "STUB: not implemented"; return false }

// isTestFunction returns true if the function name starts with "Test" and has no parameters.
// as test suite functions do not accept a *testing.T.
func isTestFunction(f *ast.FuncDecl) bool { _ = "STUB: not implemented"; return false }
