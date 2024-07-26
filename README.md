```   
,adPPYba,   ,adPPYba,  8b,dPPYba,  8b       d8   ,adPPYba,   
I8[    ""  a8P_____88  88P'   "Y8  `8b     d8'  a8"     "8a  
 `"Y8ba,   8PP"""""""  88           `8b   d8'   8b       d8  
aa    ]8I  "8b,   ,aa  88            `8b,d8'    "8a,   ,a8"  
`"YbbdP"'   `"Ybbd8"'  88              "8"       `"YbbdP"'
-----------------------------------------------------------
the perfect tool for multi-device ux/ui testing over LAN
```

## compile and link
```
$ go build
```
create a symlink from the executable to your PATH for use system-wide
```
$ sudo ln -s absolute/path/to/servo /usr/local/bin/servo
```

## usage
### basic
serves current working directory on 8080
```
$ servo
```
### custom
on specific port
```
$ servo [PORT]
```
