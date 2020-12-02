package main

import (
    // "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
	"net/http"
	"strings"
);

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	get();
	http.HandleFunc("/", manejador)
	http.ListenAndServe(":8080", nil)
	
}

func manejador(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hola, %s, ¡este es un servidor!", r.URL.Path)
}

func get() {
    fmt.Println("1. Performing Http Get...");
	resp, err := http.Get("https://aplicacionesc.mat.sat.gob.mx/SOIANET/oia_consultarapd_cep.aspx?&pa=1673&dn=9004168&s=0&ap=2019&pad=400&ad=TIJUANA,%20B.C");
	
    if err != nil {
        log.Fatalln(err);
	}

    defer resp.Body.Close();
    bodyBytes, _ := ioutil.ReadAll(resp.Body);

    // Convert response body to string
    bodyString := string(bodyBytes);

	// remplazar el ultimo encontrado por un vacio para obtener al final el ultimo que obtendra la info necesaria

	//El tag que coincide
	tag := `<font face="Arial" color="threeddarkshadow" size="2">`
	//Posicion del ultimo tag encontrado del html
	lastTag := strings.LastIndex(bodyString, tag);
	//Nuevo bodyString sin el ultimo tag (Ya que no se requiere como el ultimo)
	bodyString = bodyString[0:lastTag];
	//Nueva posicion del tag encontrado actualizado
	lastTag = strings.LastIndex(bodyString, tag);
	//Recorte del tag a lo ultimo que contenga el html
	bodyString = bodyString[lastTag+len(tag):len(bodyString)];
	//Obtiene la posicion del primer <
	lastTag = strings.Index(bodyString, "<");
	//Recorte de lo ultimo en bodyString hasta la posicion del primer < encontrado
	estatus := bodyString[0:lastTag];

	fmt.Printf("estatus %s\n", estatus);
}