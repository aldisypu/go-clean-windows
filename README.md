# go-clean-windows
`go-clean-windows` is a CLI tool designed for Windows to perform various cleaning tasks.

## Features

- Clean Prefetch folder
- Clean Temp folder
- Clean %temp% folder
- Run Disk cleanup
- Run `chkdsk /f`

## Download

To download the latest version:

1. Go to the [Releases](https://github.com/aldisypu/go-clean-windows/releases) page.
2. [Download go-clean-windows.exe](https://github.com/aldisypu/go-clean-windows/releases/download/v1.0.0/go-clean-windows.exe) file from the latest release.

Once downloaded, **run go-clean-windows.exe <mark>Run as administrator</mark>**

Alternatively, you can clone the repository and build the binary manually:

```bash
git clone https://github.com/aldisypu/go-clean-windows.git
cd go-clean-windows
go build -o go-clean-windows.exe main.go
```

## Manual Clean Windows
1. Delete files in Prefetch
2. Delete files in Temp
3. Delete files in %temp%
4. Empty Recycle Bin
5. Disk cleanup
6. Open command prompt <mark>Run as administrator</mark> => `chkdsk /f` => `y`

## Tech Stack

- Golang : https://github.com/golang/go

## Framework & Library

- Bubble Tea (TUI framework) : https://github.com/charmbracelet/bubbletea
- Lip Gloss (Style, format and layout tools) : https://github.com/charmbracelet/lipgloss