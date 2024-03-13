# gha-test
This project exists to manually test github actions OpenIdProvider. We plan to automate this long term.

1. Run `go get github.com/<github user name>/openpubkey@<commit>` where github user name and branch name are your github user name and latest commit of branch you want to test.
2. Take returned go path and put it in go.mod file. For instance: `replace github.com/openpubkey/openpubkey v0.2.2 => github.com/ethanheilman/openpubkey v0.0.0-20240313000340-1b3454599ebe`
3. Run go mod tidy
4. run go test ./... to ensure the code is correct
5. Push to origin to run the test against github

