package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/knieriem/markdown"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var port = flag.String("port", ":0", "Port to listen on (don't forget :)")
var rootDir = flag.String("root", findNotes(), "Root directory full of .md files")

type rootHandler struct{}

var markdownMode = &markdown.Extensions{
	Smart:      true,
	FilterHTML: true,
}

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const header1 = "<html><head><title>"
	const header2 = "</title></head>\n<body>\n"
	const footer = "</body></html>"
	if strings.Contains(r.URL.Path, "/.") {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	subdir := path.Clean(r.URL.Path)
	f, err := os.Open(filepath.Join(*rootDir, subdir))
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	defer f.Close()
	s, err := f.Stat()
	if s.IsDir() {
		list, err := f.Readdir(0)
		if err != nil {
			http.Error(w, "Unexpected error: "+err.Error(), http.StatusSeeOther)
			return
		}
		io.WriteString(w, header1)
		io.WriteString(w, "Notes: ")
		io.WriteString(w, subdir)
		io.WriteString(w, header2)
		if subdir != "/" {
			parent, _ := path.Split(subdir)
			io.WriteString(w, `<a href="`)
			io.WriteString(w, parent)
			io.WriteString(w, "\">Up to Parent</a>\n")
		}
		io.WriteString(w, "<ul>\n")
		for _, v := range list {
			if strings.HasPrefix(v.Name(), ".") {
				continue
			}
			fmt.Fprintf(w, `<li><a href="%s">%s</a>`,
				path.Join(subdir, v.Name()), v.Name())
			if v.IsDir() {
				io.WriteString(w, "/")
			}
			fmt.Fprintln(w, "</li>")
		}
		io.WriteString(w, "</ul>\n")
		io.WriteString(w, footer)
	} else {
		bufout := bufio.NewWriter(w)

		io.WriteString(bufout, header1)
		io.WriteString(bufout, "Notes: ")
		io.WriteString(bufout, subdir)
		io.WriteString(bufout, header2)

		p := markdown.NewParser(markdownMode)
		p.Markdown(f, markdown.ToHTML(bufout))

		io.WriteString(bufout, footer)
		bufout.Flush()
	}
}

func main() {
	flag.Parse()

	if len(*rootDir) == 0 {
		log.Fatal("Cannot locate root")
	}

	l, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatal(err)
	}

	if addr, ok := l.Addr().(*net.TCPAddr); ok {
		launch(fmt.Sprintf("http://localhost:%d/", addr.Port))
		log.Fatal(http.Serve(l, &rootHandler{}))
	} else {
		log.Fatal(`Listen("tcp") does not return a *TCPAddr`)
	}
}
