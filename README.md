# Version

Go module to maintain an application version

## Install

```sh
go get -u github.com/v3io/version-go
```

## Usage

During build, add the following linker arguments to inject the version:

```sh
-ldflags="-X github.com/v3io/version-go/version.gitCommit=$(GIT_COMMIT) \
-X github.com/v3io/version-go/version.label=$(GIT_TAG)
```

In your application, you can then:

```go
import "github.com/v3io/version-go"

// print the version
version.Get().String()

// use the version components
versionLabel := version.Get().Label
```
