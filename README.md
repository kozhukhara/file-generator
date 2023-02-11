# File Generator

File Generator is a Go module that generates a file of a given size, with optional parameters for the file name and file extension.

## Usage

To use File Generator, run the following command:
```shell
go run file_generator.go -s <size value> -u <size unit> -n <file path> 
```
where:

- `-s <size value>` is the file size, with the format <value(int)>. If not specified, the default value is 1024.
- `-u <size unit>` is the unit for the file size. can be b (bytes), kb (kilobytes), mb (megabytes), gb (gigabytes), or tb (terabytes). If not specified, the default value is b.
- `-o <file path>` is the output file name or absolute path. If not specified, the default value is `output.tmp`.

## Example

To generate a file of size 1 kilobyte, with the name random.tmp, located in current directory, run the following command:
```shell
go run main.go -s 1024 -u b -o random.tmp
```

The output will be:
```
Created successfully
- File name: random.tmp
- Full path: /absolute/path/to/random.tmp
- File size: 1024 byte(s)
```
