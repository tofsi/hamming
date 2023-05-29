# Error Correcting (Hamming) Codes

## Purpose

Information is represented as a sequence of words consisting of bits (0 or 1). When transmitting information, a bit may be flipped by interference from some source.

It is of interest to 

1. Know that an error has occured

2. Recover the original information

Hamming codes solve both of these problems by adding redundancy.
Hamming codes detect two errors, and correct one error per code word.
Notation: the Hamming code H(n, k) has code words of length n (block length) and words of length k (word length).
For Hamming codes, n = 2<sup>m</sup> - 1 and k = 2<sup>m</sup> - m - 1, where m is a positive integer.

## Installation

`go get github.com/tofsi/hamming`

## Documentation

Documentation is found in the Documentation.md file

## Example

The following is an example of how to encode and decode data.

```go
data := []bool{true, false, true, false, true, true, false, true}
c := hamming.New(3)
c.Print() // To get an overview of the code
encodedData := c.EncodeBin(data)
encodedData[3] = !encodedData[3] // Oh no! A flipped bit
decodedData, bitFlips := c.DecodeBin(encodedData, 0) // If working with H(7, 4) and
                                                    // bytes, remainder is always 0
fmt.Println("Original Data:")
fmt.Println(data)
fmt.Println("Detected errors in bits:")
fmt.Println(bitFlips)
fmt.Println("Decoded Data:")
fmt.Println(decodedData)
```
