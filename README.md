# fshare

fshare is very simple webserver that is useful for development, testing and file sharing on a local network.

## Install

Make sure you have go installed in your machine or follow the instructions on [Go website](https://go.dev/doc/install) to install go. Once installed, run the following command on your terminal to install `fshare` server.

```
go install github.com/eminacio/fshare/cmd/fshare@latest
```

## Usage

Navigate to the folder that contains the files you want to share or serve with `fshare` and run the following command:

```
fshare [Options]
```

### Example

```
fshare
```

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

## Example

```
fshare --directory /path/to/my/files
```

```
fshare --port 3000
```

After running the command, your server will be active. You can access the shared files from another machine on the same network by navigating to `http://IP_ADDRESS:PORT_NUMBER` replacing `IP_ADDRESS` AND `PORT_NUMBER` with the local IP of the machine running the server. To view the files on the same machine, you can use the localhost address: `http://localhost:8000`.
