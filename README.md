# go-clean-windows
Go-Lang Clean Windows

Open go-clean-windows.exe <mark>Run as administrator</mark>

## Manual Clean Windows
1. Delete files in Prefetch
2. Delete files in Temp
3. Delete files in %temp%
4. Empty Recycle Bin
5. Disk cleanup
6. Open command prompt <mark>Run as administrator</mark> => `chkdsk /f` => `y`

## Run Go Build
```bash
go build -o go-clean-windows.exe main.go
```

## Tech Stack

- Golang : https://github.com/golang/go

## Framework & Library

- Bubble Tea (TUI framework) : https://github.com/charmbracelet/bubbletea
- Lip Gloss (Style, format and layout tools) : https://github.com/charmbracelet/lipgloss