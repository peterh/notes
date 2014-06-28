Notes
=====

This is a quick-and-dirty personal notes system. It reads markdown files
from `$GOPATH/src/notes`, and launches a web browser for read-only access.

Editing the notes is done with your favourite editor; this is a read-only system.

Because it is read only, it should be relatively safe. However, I have not
thought about it deeply and do not advise exposing `notes` as a public web
server.

Usage
-----
`mkdir $GOPATH/src/notes; echo "Hello, World!" > $GOPATH/src/notes/test; notes`

License
-------
The license of `notes` is MIT/X11. See `COPYING` for details.
