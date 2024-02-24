# Go Concat Strings

A little test to find the most efficient way to concatenate strings in Go.

Based off an [article](https://shantanubansal.medium.com/golang-how-to-efficiently-concatenate-strings-f2e51564f8d) by @shantanubansal

## Findings
* When the lenght of a string is known, the most efficient way of concatonating a string is using copy
* When the lenght if a string is unknown, the most effiecent way of concatonating a string is using StringBuilder with WriteString
* Using the strings.Builder String() method does not decrease efficiency