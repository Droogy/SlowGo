# SlowGo
 A slow and naive TCP portscanner written in Go

### Usage

- Create a file called ports.conf, put each port you want to scan on a newline within the file
- Compile the Go program and make it executable
  - `go build naive_TCP.go && chmod +x naive_TCP`

- Finally run the program and add your target as the first and only command line argument
  - `./naive_TCP scanme.nmap.org`

### To-Do

- Add concurrency to the scanning routine
- More robust scanning options
- Maybe get rid of need for configuration file?
- Scan top 100 ports by default
- Options for an output file - for now just pipe execution into tee or however you prefer