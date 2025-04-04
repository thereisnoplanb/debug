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
// this function writes the messages to the standard output without appending a newline.
// If Enabled is false, no output will be produced.
//
// # Example
//
//  debug.Enabled = true
//  debug.Print("First part of the message", " ", "Second part.")
//  // Output: First part of the message Second part.
//
//  debug.Enabled = false
//  debug.Print("This will not be printed.")
func Print(messages ...any) {
	if Enabled {
		fmt.Print(messages...)
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
// to the standard output without appending a newline. If either Enabled is false
// or the condition evaluates to false, no output will be produced.
//
// # Example
//
//  debug.Enabled = true
//  debug.PrintIf(5 > 3, "Condition met", " ", "Message continues.")
//  // Output: Condition met Message continues.
//
//  debug.Enabled = true
//  debug.PrintIf(false, "This will not be printed.")
//
//  debug.Enabled = false
//  debug.PrintIf(true, "Debugging is disabled.")
func PrintIf(condition bool, messages ...any) {
	if Enabled && condition {
		fmt.Print(messages...)
	}
}
