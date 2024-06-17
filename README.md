# Golang to WASM Compilation Template
## Environment
- Ubuntu 20.04 LTS (WSL 2)
- Go 1.20
- Goland IDE
## Usage
1. Run `./build.sh`
2. Open `test/index.html` using a local static server
## Why Go Instead of C/C++/Rust
I prioritize achieving high development efficiency. I am unconcerned about the overhead of loading the Go runtime because the resulting file size is only 1.5 MB, which is acceptable to me.