# Natonize

_Natonize is unstable and mostly a toy._

Natonize provides a small library to convert strings to their NATO phonetic alphabet form.

## Usage

Simply pass whatever you want translated to phonetic alphabet as command line arguments to Natonize.
```
$ natonize Illinois
capital-india lima lima india november oscar india sierra
```

Add `--reverse` to reverse the process.
```
$ natonize --reverse capital-india lima lima india november oscar india sierra
Illinois
```

## Installation
```
go install github.com/sakjur/natonize/cmd/natonize
```