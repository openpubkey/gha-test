# gha-test

To manually test github actions

1. Change `replace github.com/openpubkey/openpubkey v0.2.2 => github.com/<fork>/openpubkey <branch name>` where fork and branch name is the name of your fork and branch.
2. Run go mod tidy
3. Push to origin to run the test

