```   
,adPPYba,   ,adPPYba,  8b,dPPYba,  8b       d8   ,adPPYba,   
I8[    ""  a8P_____88  88P'   "Y8  `8b     d8'  a8"     "8a  
 `"Y8ba,   8PP"""""""  88           `8b   d8'   8b       d8  
aa    ]8I  "8b,   ,aa  88            `8b,d8'    "8a,   ,a8"  
`"YbbdP"'   `"Ybbd8"'  88              "8"       `"YbbdP"'
-----------------------------------------------------------
the perfect tool for multi-device ux/ui testing over LAN
```

## Pre-Requirements
- [Go 1.22.5+](https://go.dev/doc/install)

## Compile and Link
```
$ go build
```
Create a symlink from the executable to your PATH for use system-wide
```
$ sudo ln -s absolute/path/to/servo /usr/local/bin/servo
```

## Usage
### Basic
Serves current working directory on 8080
```
$ servo
```
### Custom
On specific port
```
$ servo [PORT]
```
