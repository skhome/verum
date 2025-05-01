# verum

An assertion library for Go. Verum matchers are self-describing and extendable.
You can extend the library with your own custom matchers.

## Usage

### Import

```go
import (
    "github.com/skhome/verum/assert"
    . "gitgub.com/skhome/verum/matchers"
)
```

### Example

```go
func TestRingsOfPower(t *testing.T) {
    assert.That(t, "nenya").
        DescribedBy("ring of power").
        Matches(IsEqualTo("vilya"))
}
```

output:

```plain
[ring of power] value equal to <vilya>, but got <nenya>
```
