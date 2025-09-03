# go-ascii


I needed some custom text art for my nvim dashboard config and was learning go at the time. 
So, I made a quick and easy cli tool in go to convert images into ASCII, ANSI, or Braille text art right from your terminal!

---

## Installation

### Prerequisites
- [Go 1.22+](https://golang.org/dl/) installed on your system  
- `$GOPATH/bin` or `$HOME/go/bin` (or `$GOBIN`) on your `PATH`

### Install with Go

```bash
go install github.com/KhasarMunkh/go-ascii@latest
```

### Or clone the repo and run locally:
```sh
git clone https://github.com/KhasarMunkh/go-ascii.git
cd go-ascii
go run . --mode ascii --width 120 assets/zoro.jpeg
```
## Usage

```sh
go-ascii --mode ascii|ansi|braille --width <cols> [--out <file or ->] <image file or ->
```

- `--mode`   : Rendering mode (`ascii`, `ansi`, or `braille`). Default: `braille`
- `--width`  : Output width in character cells. Default: `200`
- `--out`    : Output file (default: stdout).
- `<image>`  : Path to image file 

### Example
<img width="1876" height="1107" alt="2025-09-03-171404_hyprshot" src="https://github.com/user-attachments/assets/aabfa63c-fc79-4b4f-9781-5ec75d24d0a7" />

## License
MIT
