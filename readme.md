## CI pipeline

A CI pipeline for Go applications that does the following:

1. Compiles the code to the binary
2. Runs unit tests
3. Formats the source code according to the standard go formatter.
4. Pushes commits to the remote git repository of the project.

Steps to run:

1. Clone the git repository.
2. cd to the directory and run "go build".

Usage:

`./go-ci -p <project directory>`

## NOTE:

In order to push to a remote repository on GitHub, you need to use SSH Url of your project as the remote URL instead of HTTPS and generate a public key pair as well.

How to generate SSH Key: https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent

How to add SSH Key to GithHub: https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account
