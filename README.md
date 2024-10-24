# mdf - Markdown Live Formatter

mdf renders your written .md file in a web browser on a local server:
```
http://localhost:8080/view
```
![20241024_16h51m22s_grim](https://github.com/user-attachments/assets/1eacee9d-bf40-4c29-91fc-bbe5ed75a64d)



## What is this about?

**Markdown** is a lightweight markup language for creating formatted text using a plain-text editor. John Gruber created Markdown in 2004 as an easy-to-read markup language. Markdown is widely used for blogging and instant messaging, and also used elsewhere in online forums, collaborative software, documentation pages, and readme files.

[Reference - Wiki](https://en.wikipedia.org/wiki/Markdown)

**mdf** is a markdown formatter, written in Golang - that displays a given .md file on a local webserver. 
This is just a small tool, to help for quick and easy markdown creation. 

## Install

- Using the compiled binary or build from source.

### Compiled Binary 

Using the compiled bin in the build folder:

```
cd /build
```

```
./mdf --help
```


### Build Source

Clone this repository:
```
git clone https://github.com/nyangkosense/mdf
```

Using makefile:
```
cd mdf
make && make install
```

Or:

```
go build 
```
---

## Usage

Use `mdf --help` to display all available options.

```
Usage:
  mdf [options]

Options:
  --cheatsheet
        Display Markdown syntax cheatsheet
  --filepath string
        Path to the Markdown file
  --help
        Display help information
  --theme string
        Theme to use (github-dark, github-light, dracula, nord) (default "github-dark")
```

### Examples:

```
mdf --filepath /path/to/file.md
mdf --filepath /path/to/file.md --theme dracula
mdf --cheatsheet
mdf --help
```

#### media

-- example goes in here --



## Headline
