package debug

import "fmt"

// Outputs the formatted string and provided messages to the standard output if debugging is enabled.
//
// # Parameters
//
//  format string
//
// A format string that specifies how the messages will be formatted.
// Follows the same formatting rules as `fmt.Printf`.
//
//  messages ...any
//
// A variable number of arguments representing the messages to be included in the formatted output.
// These messages are passed as <any> type and will be formatted according to the format string.
//
// # Remarks
//
// Debugging is controlled by the global Enabled flag. When Enabled is set to true,
// this function writes the formatted string and messages to the standard output.
// If Enabled is false, no output will be produced.
//
// # Example
//
//  debug.Enabled = true
//  debug.Printf("Value: %d, Name: %s", 42, "Example")
//  // Output: Value: 42, Name: Example
//
//  debug.Enabled = false
//  debug.Printf("This will not be printed.")
func Printf(format string, messages ...any) {
	if Enabled {
		fmt.Printf(format, messages...)
	}
}

// Outputs the formatted string and provided messages to the standard output if debugging is enabled
// and the specified condition is true.
//
// # Parameters
//
//  condition bool
//
// A boolean value that determines whether the messages should be printed.
// The output occurs only if this condition evaluates to true.
//
//  format string
//
// A format string that specifies how the messages will be formatted.
// Follows the same formatting rules as `fmt.Printf`.
//
//  messages ...any
//
// A variable number of arguments representing the messages to be included in the formatted output.
// These messages are passed as <any> type and will be formatted according to the format string.
//
// # Remarks
//
// Debugging is controlled by the global Enabled flag. When Enabled is set to true
// and the specified condition evaluates to true, this function writes the formatted string
// and messages to the standard output. If either Enabled is false or the condition evaluates to false,
// no output will be produced.
//
// # Example
//
//  debug.Enabled = true
//  debug.PrintfIf(5 > 3, "Value: %d, Status: %s", 42, "Active")
//  // Output: Value: 42, Status: Active
//
//  debug.Enabled = true
//  debug.PrintfIf(false, "This will not be printed.")
//
//  debug.Enabled = false
//  debug.PrintfIf(true, "Debugging is disabled.")
func PrintfIf(condition bool, format string, messages ...any) {
	if Enabled && condition {
		fmt.Printf(format, messages...)
	}
}
