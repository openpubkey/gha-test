# gha-test
This project exists to manually test github actions OpenIdProvider. We plan to automate this long term.

1. Run `go mod edit -replace github.com/openpubkey/openpubkey@v0.2.2=github.com/<github user name>/openpubkey@<commit>` where github user name and branch name are your github user name and latest commit of branch you want to test. Alternatively you can use `go mod edit -replace github.com/openpubkey/openpubkey@v0.2.2=github.com/<github user name>/openpubkey@<fork>` where fork is the the fork you want use.
For instance: `go mod edit -replace github.com/openpubkey/openpubkey@v0.2.2=github.com/ethanheilman/openpubkey@43d9a377b2336bb8e3cabcd9e74918e6debd9e4f` or `go mod edit -replace github.com/openpubkey/openpubkey@v0.2.2=github.com/ethanheilman/openpubkey@memguardhide`

2. Run `go mod tidy`

3. run go test ./... to ensure the code is correct

4. Push to origin to run the test in a github-action

