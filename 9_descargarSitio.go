package main
/** Descargado de sitios web.

Nativamente Go ofrece características para trabajar con aplicaciones web.
En este archivo se muestra como descartar sitios web.
*/
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Lectura de parámetros en la línea de comandos.
	for _, url := range os.Args[1:] {
		// Realizar una petición GET sobre la URL
		resp, err := http.Get(checkPrefix(url))

		// Revisión de errores
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// Para guardar el contenido del sitio, realizamos una copia.
		//_, err = io.Copy(os.Stdout, resp.Body)

		file, err := os.Create("out.html")

		if err != nil {
			fmt.Fprintf(os.Stderr, ": %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)

		if( err != nil) {
			fmt.Fprintf(os.Stderr, "Error while fetching URL %s : %v", url, err)
		} else {
			fmt.Fprintf(os.Stdout, "URL requested: %s with code: %v\n.", url, resp.Status)
		}
	}
}

func checkPrefix(url string) string{
	// Es obligatorio que las URLs tengan el protocolo correctamente definidio.
	newURL := url
	if ! (strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		newURL = "http://" + url
	}
	return newURL
}
