# word-counter
Golang program to count words in files

## Purpose
To count words in English or other languages in files.

A word is defined as a sequence of Unicode letters, possibly separated by one or more apostrophes.

Results are sorted in descending order by frequency.

## Status
Ready to use

## Installation
This program is written in Google Go language. Make sure that Go is installed and the GOPATH is set up as described in
[How to Write Go Code](https://golang.org/doc/code.html).

The install this program and its dependencies by running:

    go get github.com/BluntSporks/word-counter

## Usage
There are two modes of usage:
* Count all the words in a file.
* Look up the word counts of words in a file using an existing word count file.

Usage:

    word-counter FILENAME
    word-counter -lookup=LOOKUPFILE FILENAME

Options:

    -lookup=LOOKUPFILE  Existing word count to use for lookup
