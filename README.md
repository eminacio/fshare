# fshare

fshare is very simple webserver that is useful for development, testing and file sharing on a local network.

## Install

Make sure you have go installed in your machine or follow the instructions on [Go website](https://go.dev/doc/install) to install go. Once installed, run the following command on your terminal to install `fshare` server.

`go install github.com/eminacio/fshare`

## Usage

`fshare [Options]`

### OPTIONS

    The options which apply to the fshare command are:

    -port, --port PORT_NUMBER
        Set the server port where to listen for requests.
        By default, when this option is not specified,
        fshare listens for requests on port 8000.

    -directory, --directory DIRECTORY
     Set a specific directory where to serve files from.
     The path to the directory can be either relative or absolute path.
     When this option is not set, the default behavior is to serve files
     from the current working directory.
