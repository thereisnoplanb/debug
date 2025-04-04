package debug

import "fmt"

// Outputs the provided messages to the standard output if debugging is enabled.
//
// # Parameters
//
//  messages ...any
//
// A variable number of arguments representing the messages to be printed.
// These messages are passed as <any> type and will be concatenated into a single output.
//
// # Remarks
//
// Debugging is controlled by the global Enabled flag. When Enabled is set to true,
// this function writes the messages to the standard output.
// If Enabled is false, no output will be produced.
//
// # Example
//
//	debug.Enabled = true
//	debug.Println("Debugging message", 123, true)
//	// Output: Debugging message 123 true
//
//	debug.Enabled = false
//	debug.Println("This will not be printed.")
func Println(messages ...any) {
	if Enabled {
		fmt.Println(messages...)
	}
}

// Outputs the provided messages to the standard output if debugging is enabled
// and the specified condition is true.
//
// # Parameters
//
//  condition bool
//
// A boolean value that determines whether the messages should be printed.
// The output occurs only if this condition evaluates to true.
//
//  messages ...any
//
// A variable number of arguments representing the messages to be printed.
// These messages are passed as <any> type and will be concatenated into a single output.
//
// # Remarks
//
// Debugging is controlled by the global Enabled flag. When Enabled is set to true
// and the specified condition evaluates to true, this function writes the messages
// to the standard output. If either Enabled is false or the condition evaluates to false,
// no output will be produced.
//
// # Example
//
//  debug.Enabled = true
//  debug.PrintlnIf(5 > 3, "Condition met", 123)
//  // Output: Condition met 123
//
//  debug.Enabled = true
//  debug.PrintlnIf(false, "This will not be printed.")
//
//  debug.Enabled = false
//  debug.PrintlnIf(true, "Debugging is disabled.")
func PrintlnIf(condition bool, messages ...any) {
	if Enabled && condition {
		fmt.Println(messages...)
	}
}
