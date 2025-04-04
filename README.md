# Debug Package

The `debug` package provides a simple mechanism for conditional debugging output. It allows you to control whether debug messages are printed to the standard output using a global flag `Enabled`.

## Features

- Print messages conditionally based on the `Enabled` flag.
- Support for formatted output similar to `fmt.Printf`.
- Flexible conditional printing using additional conditions.

## Installation

```cmd
go get github.com/thereisnoplanb/debug
```

To use the `debug` package in your Go project, include it in your project and import it as follows:

```go
import "github.com/thereisnoplanb/debug"
```

# Usage

### Global Flag

The behavior of the package is controlled by the global `Enabled` flag:

```go
debug.Enabled = true // Enable debug output
debug.Enabled = false // Disable debug output
```

### Functions

`Println`

```go
func Println(messages ...any)
```

Outputs the provided messages to the standard output if debugging is enabled.

##### Parameters

- messages `...any`
A variable number of arguments representing the messages to be printed.

##### Remarks
When `Enabled` is set to `true`, this function writes the messages to the standard output without appending a newline.
If `Enabled` is `false`, no output will be produced.

##### Example

```go
debug.Enabled = true
debug.Println("Debugging message", 123, true)
// Output: Debugging message 123 true

debug.Enabled = false
debug.Println("This will not be printed.")
```

`PrintlnIf`

```go
func PrintlnIf(condition bool, messages ...any)
```

Outputs the provided messages to the standard output if debugging is enabled and the specified condition is `true`.

##### Parameters

- condition `bool`
A boolean value to determine if the messages should be printed.

- messages `...any`
A variable number of arguments representing the messages to be printed.

##### Remarks

No output if Enabled is false or condition is false.

##### Example

```go
debug.Enabled = true
debug.PrintlnIf(5 > 3, "Condition met", 123)
// Output: Condition met 123
```

`Printf`

```go
func Printf(format string, messages ...any)
```

Outputs a formatted string and provided messages to the standard output if debugging is enabled.

##### Parameters

- format `string`
The format string to define the output format.
- messages `...any`
Arguments to be formatted and printed.

##### Remarks

When Enabled is false, no output is produced.

##### Example:

```go
debug.Enabled = true
debug.Printf("Value: %d, Name: %s", 42, "Example")
// Output: Value: 42, Name: Example
```

`PrintfIf`

```go
func PrintfIf(condition bool, format string, messages ...any)
```

Outputs a formatted string and provided messages to the standard output if debugging is enabled and the condition is true.

##### Parameters

- condition `bool`
Determines whether the output should occur.
- format `string`
The format string to define the output format.
- messages `...any`
Arguments to be formatted and printed.

##### Remarks

No output if Enabled is false or condition is false.

##### Example

```go
debug.Enabled = true
debug.PrintfIf(5 > 3, "Value: %d, Status: %s", 42, "Active")
// Output: Value: 42, Status: Active
```

`Print`

```go
func Print(messages ...any)
```

Outputs the provided messages to the standard output if debugging is enabled.

##### Parameters

- messages `...any`
Messages to be printed without a newline.

##### Remarks

When Enabled is false, no output is produced.

##### Example

```go
debug.Enabled = true
debug.Print("First part", " ", "Second part.")
// Output: First part Second part.
```

`PrintIf`

```go
func PrintIf(condition bool, messages ...any)
```

Outputs the provided messages to the standard output if debugging is enabled and the condition is true.

##### Parameters

- condition `bool`
Determines whether the output should occur.
- messages `...any`
Messages to be printed without a newline.

##### Remarks

No output if Enabled is false or condition is false.

##### Example

```go
debug.Enabled = true
debug.PrintIf(5 > 3, "Condition met", " ", "Message continues.")
// Output: Condition met Message continues.
```

# License

MIT
This package is open-source. Feel free to use it in your projects under an appropriate license.