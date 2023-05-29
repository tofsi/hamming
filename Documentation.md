# Documentation

## Overview

Package hamming provides encoding, decoding, and error detection/correction using Hamming codes.

Hamming codes are characterized by the word length `w` and the minimal distance `delta` between code words.

## Index
type Code

func New(w int, delta int) \*Code

func (c Code) EncodeBin(x []bool) []bool

func (c Code) DecodeBin(y []bool) ([]bool, "")

func (c Code) EncodeString(x string) []bool

func (c Code) DecodeString(y []bin) string 


## Types

### type Code

```go
type Code struct {
    // Contains filtered or unexported fields
}
```
Code stores the pertinent information of the error correcting code

### func New

```go
func New(w int, delta int) *Code
```
`New` creates a new code, for a word length `w` and a minimal distance `delta`.

### func EncodeBin

```go
func (c Code) EncodeBin(x []bool) []bool
```
`EncodeBin` encodes the sequence of words `x` to a sequence of code words of the code `c`.

### func DecodeBin
```go
func (c Code) DecodeBin(y []bool) ([]bool, "")
```
`DecodeBin` corrects and detects errors in the sequence of code words `y` of `c` and tries to recover the original sequence of words. If there are any errors, a string with the relevant information is returned as well.

### func EncodeString
```go
func (c Code) EncodeString(x string) []bool
```
`EncodeString` encodes the string `x` to a sequence of code words of the code `c`.

### func DecodeString
```go
func (c Code) DecodeString(y []bin) string 
```
`DecodeString` corrects and detects errors in the sequence of code words `y` of `c` and tries to recover the original string, (encoded using `EncodeString`). If there are any errors, a string with the relevant information is returned as well.
