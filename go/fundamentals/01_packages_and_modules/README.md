### Exported vs Unexported

```bash
var fooBaz string // This is an unexported variable.
var FooBar string // This is an exported variable.

func fooBaz() {...} // This is an unexported function.
func FooBar() {...} // This is an exported function.

type fooBaz struct {...} // This is an unexported type.
type FooBar struct {...} // This is an exported type.
```

### Organizing Import Statement

```bash
import (
    {standard library packages}

    {packages from the current module}

    {third-party packages}

    {aliased packages}
)
```

### Go Module (Example)

```bash
go mod init github.com/naufalihsan/academy
go get github.com/fatih/color@latest # third-party packages
go mod tidy
```

### References
- https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules