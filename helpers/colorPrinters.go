package helpers

import "github.com/fatih/color"

// PositivePrintln prints bold and green colored text with line ending
var PositivePrintln = color.New(color.FgGreen, color.Bold).PrintlnFunc()

// NegativePrintln prints bold and red colored text with line ending
var NegativePrintln = color.New(color.FgRed, color.Bold).PrintlnFunc()
