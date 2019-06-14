package main
/**
$GOOS=js GOARCH=wasm go build -o test.wasm test.go
 */
import (
"encoding/hex"
"flag"
"fmt"
"html"
"io"
"log"
"net/http"
)

func main() {
	addr := flag.String("addr", ":5555", "server address:port")
	flag.Parse()
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/wasm", wasmHandle)

	log.Printf("listening on %q...", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, page, html.EscapeString(hex.Dump(wasmAdd)))
}

func wasmHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	n, err := w.Write(wasmAdd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	if n != len(wasmAdd) {
		http.Error(w, io.ErrShortWrite.Error(), http.StatusServiceUnavailable)
	}
}

var wasmAdd = []byte{
	0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00,
	0x01, 0x07, 0x01, 0x60, 0x02, 0x7f, 0x7f, 0x01,
	0x7f, 0x03, 0x02, 0x01, 0x00, 0x07, 0x07, 0x01,
	0x03, 0x61, 0x64, 0x64, 0x00, 0x00, 0x0a, 0x09,
	0x01, 0x07, 0x00, 0x20, 0x00, 0x20, 0x01, 0x6a,
	0x0b,
}

const page = `
<html>
	<head>
		<title>Testing WebAssembly</title>
		<script type="text/javascript">

		function fetchAndInstantiate(url, importObject) {
			return fetch(url).then(response =>
				response.arrayBuffer()
			).then(bytes =>
				WebAssembly.instantiate(bytes, importObject)
			).then(results =>
			    results.instance
			);
		}

		var mod = fetchAndInstantiate("/wasm", {});

		window.onload = function() {
			mod.then(function(instance) {
				var div = document.getElementById("wasm-result");
				div.innerHTML = "<code>add(1, 2)= " + instance.exports.add(1, 2) + "</code>";
			});
		};

		</script>
	</head>

	<body>
		<h2>WebAssembly content</h2>
		<div id="wasm-content">
			<pre>%s</pre>
		</div>

		<h2>WebAssembly</h2>
		<div id="wasm-result"><code>add(1, 2)= N/A</code></div>
	</body>
</html>
`