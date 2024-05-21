# File Server

This is a simple file server written in Go. It serves files from the current (or specific) directory.

## Usage

Download the binary from 'Release' from github

To start the server, run the following command:

`./file_server_linux_amd64`

By default the current directory is served, and a random port number is selected.


To specify a direct, use `-d`
e.g. 
`./file_server_linux_amd64 -d /tmp`

To specify a port number, use `-p`
e.g.
`./file_server_linux_amd64 -p 8080`

To access the files, open a web browser and navigate to the printed address.

