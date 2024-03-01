```
rei@teaparty MINGW64 ~/Documents/GitHub/telebotspam (main)
$ ./telebotspam-windows-latest-amd64

    ╔╦╗┌─┐┬  ┌─┐╔╗ ┌─┐┌┬┐╔═╗╔═╗╔═╗╔╦╗
     ║ ├┤ │  ├┤ ╠╩╗│ │ │ ╚═╗╠═╝╠═╣║║║
     ╩ └─┘┴─┘└─┘╚═╝└─┘ ┴ ╚═╝╩  ╩ ╩╩ ╩
    By @elliottophellia        v1.0.0

The message to send? mati aja lo kontol
How many times to send the message? 0
[+] Success.
[+] Success.
[+] Success.
[+] Success.
[+] Success.

Received an interrupt, stopping...
```
# General usage:

### Download the binary

<a href="https://github.com/elliottophellia/telebotspam/releases/latest/download/telebotspam-windows-amd64"><img src="https://img.shields.io/badge/DOWNLOAD-FOR%20WINDOWS-white?style=for-the-badge&logo=github"/></a>
<a href="https://github.com/elliottophellia/telebotspam/releases/latest/download/telebotspam-linux-amd64"><img src="https://img.shields.io/badge/DOWNLOAD-FOR%20LINUX-white?style=for-the-badge&logo=github"/></a>
<a href="https://github.com/elliottophellia/telebotspam/releases/latest/download/telebotspam-darwin-amd64"><img src="https://img.shields.io/badge/DOWNLOAD-FOR%20DARWIN-white?style=for-the-badge&logo=github"/></a>

### Run it

```
./telebotspam-windows-amd64

    ╔╦╗┌─┐┬  ┌─┐╔╗ ┌─┐┌┬┐╔═╗╔═╗╔═╗╔╦╗
     ║ ├┤ │  ├┤ ╠╩╗│ │ │ ╚═╗╠═╝╠═╣║║║
     ╩ └─┘┴─┘└─┘╚═╝└─┘ ┴ ╚═╝╩  ╩ ╩╩ ╩
    By @elliottophellia        v1.0.0

The message to send?
```

#### Flags:
```
-loop value
        How many times to send the message
-msg string
        The message to send
```

#### Notes:
In unix-like operating systems, first you need to give permission to the binary to execute it.
```
chmod +x telebotspam-linux-amd64
```
Before running the binary, make sure you have `config` file on the same directory
```
botToken=TARGET_TOKEN
chatID=TARGET_CHATID
```

# Development usage:

### Clone the repository
```
git clone https://github.com/elliottophellia/telebotspam.git
```

### Go to directory
```
cd telebotspam
```

### Run the script
```
go run main.go
```

### Build the binary
```
GOARCH=amd64 go build -o telebotspam main.go
```



