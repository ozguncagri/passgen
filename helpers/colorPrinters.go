package helpers

import "github.com/fatih/color"

// PositivePrint prints bold and green colored text
var PositivePrint = color.New(color.FgGreen, color.Bold).PrintFunc()

// NegativePrint prints bold and red colored text
var NegativePrint = color.New(color.FgRed, color.Bold).PrintFunc()

// PositivePrintf prints bold and green colored text with formatter
var PositivePrintf = color.New(color.FgGreen, color.Bold).PrintfFunc()

// NegativePrintf prints bold and red colored text with formatter
var NegativePrintf = color.New(color.FgRed, color.Bold).PrintfFunc()

// PositivePrintln prints bold and green colored text with line ending
var PositivePrintln = color.New(color.FgGreen, color.Bold).PrintlnFunc()

// NegativePrintln prints bold and red colored text with line ending
var NegativePrintln = color.New(color.FgRed, color.Bold).PrintlnFunc()
