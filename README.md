## Cobra
Install Cobra CLI:
```shell
go install github.com/spf13/cobra-cli@latest
```

Init go project:
```shell
mkdir my-cli && cd my-cli
go mod init
cobra-cli init
```

Add command:
```shell
cobra-cli add get
```

## Task

Install [Task][TaskUrl]
```shell
brew install go-task
```

Create Taskfile.yaml:
```yaml
version: "3"

tasks:
  build:
    desc: Build the app
    cmds:
      - GOFLAGS=-mod=mod go build -o bin/my-cli main.go

  run:
    desc: Run the app
    cmds:
      - GOFLAGS=-mod=mod go run main.go

  clean:
    desc: Remove all retrieved *.png files
    cmds:
      - rm -rf img
```

Usage:
```shell
task build
./bin/my-cli get 5th-element
task clean
```

[TaskUrl]: https://taskfile.dev/installation/