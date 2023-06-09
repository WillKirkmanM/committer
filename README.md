# Committer
⭐Push Your Git Commit Messages to the Moon

[![asciicast](https://asciinema.org/a/GeoloSaR4JFMcGHWZHveJmObB.svg)](https://asciinema.org/a/GeoloSaR4JFMcGHWZHveJmObB)

## Developing

To add a development build to the path. Make sure you have

`export PATH=$PATH:~/.local/bin/`

in your shell rc file.
Once that is done, you can run the command to add the development binary to your path.:

```
$ make path
```

To remove the binary from your path, run this command:

```
$ make cleanPath
```

## Building

To build normally, run the command:

```
$ make build
```

To build with the lowest binary size run:

```
$ make production
```
