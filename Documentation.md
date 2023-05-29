# Documentation

## Overview

Package hamming provides encoding, decoding, and error detection/correction using Hamming codes.

Hamming codes are characterized by the word length `w` and the minimal distance `delta` between code words.

## Index
type Code

func New(m int) \*Code

func (c Code) EncodeBin(x []bool) []bool

func (c Code) DecodeBin(y []bool, remainder int) ([]bool, []int)

func (c Code) EncodeString(x string) []bool

func (c Code) DecodeString(y []bin, rematinder int) (string, []int) 


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
func New(m int) *Code
```
`New` creates the Hamming code H(n, k) = H(2<sup>m</sup> - 1, 2<sup>m</sup> - m - 1).
The length of words is k and the length of blocks (data and parity bits) is n.
The code corrects one error and detects two.

### func EncodeBin

```go
func (c Code) EncodeBin(x []bool) []bool
```
`EncodeBin` encodes the sequence of words `x` to a sequence of code words of the code `c`.

### func DecodeBin
```go
func (c Code) DecodeBin(y []bool, remainder int) ([]bool, []int)
```
`DecodeBin` corrects and detects errors in the sequence of code words `y` of `c` and tries to recover the original sequence of words. `remainder` is equal to the length of the message, in bits, modulo the length of one block, `c.n`. A slice containing the positions of detected errors is returned as well.

### func EncodeString
```go
func (c Code) EncodeString(x string) []bool
```
`EncodeString` encodes the string `x` to a sequence of code words of the code `c`.

### func DecodeString
```go
func (c Code) DecodeString(y []bin) (string, []int)
```
`DecodeString` corrects and detects errors in the sequence of code words `y` of `c` and tries to recover the original string, (encoded using `EncodeString`). `remainder` is equal to the length of the message, in bits, modulo the length of one block, `c.n`. A slice containing the positions of characters in which errors where detected is also returned.
