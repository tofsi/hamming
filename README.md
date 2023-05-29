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

`go install -u github.com/tofsi/hamming`

## Documentation

Documentation is found in the Documentation.md file
