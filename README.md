# version
Go module to maintain an applicaiton version

## Usage
During build Add the following linker arguments to inject the version:
```
-ldflags="-X github.com/v3io/version/version.gitCommit=$(GIT_COMMIT) \
-X github.com/v3io/version/version.label=$(GIT_TAG) \
-X github.com/v3io/version/version.os=$(GOOS) \
-X github.com/v3io/version/version.arch=$(GOARCH)"
```

In your application, you can then:
```go
import "github.com/v3io/version-go"

// print the version
version.Get().String()

// use the version components
versionLabel := version.Get().Label
```
