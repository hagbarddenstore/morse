# Morse

A simple ASCII to Morse encoder and decoder.

## Installation

Run ```go get github.com/hagbarddenstore/morse```.

## Usage

```morse``` reads from stdin and prints to stdout.

### Encoding

```
echo 'SOS' | morse

# Output: - - -   --- --- ---   - - -
```

### Decoding:

```
echo '- - -   --- --- ---   - - -' | morse -decode

# Output: SOS
```
