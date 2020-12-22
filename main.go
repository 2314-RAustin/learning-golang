package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
);

func main() {
	getStatusFromPageOldSATByPedimento("9004168");
	getStatusFromPageNewSATByNumberOfIntegration("58343656");
	http.HandleFunc("/", manejador)
	http.ListenAndServe(":8080", nil)
}

func manejador(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hola, %s, ¡este es un servidor!", r.URL.Path)
}

func getStatusFromPageOldSATByPedimento(numPedimento string){
	fmt.Println("Obteniendo datos de la pagina de SOIANET...");
	resp, err := http.Get("https://aplicacionesc.mat.sat.gob.mx/SOIANET/oia_consultarapd_cep.aspx?&pa=1673&dn="+numPedimento+"&s=0&ap=2019&pad=400&ad=TIJUANA,%20B.C");

	if err != nil {
        log.Fatalln(err);
	} else {
		respApiString := convertRespFromHttResponseToString(resp);
	
		//El tag que coincide
		tag := `<font face="Arial" color="threeddarkshadow" size="2">`
		//Posicion del ultimo tag encontrado del html
		lastTag := strings.LastIndex(respApiString, tag);
		//Nuevo respApiString sin el ultimo tag (Ya que no se requiere como el ultimo)
		respApiString = respApiString[0:lastTag];
		//Nueva posicion del tag encontrado actualizado
		lastTag = strings.LastIndex(respApiString, tag);
		//Recorte del tag a lo ultimo que contenga el html
		respApiString = respApiString[lastTag+len(tag):len(respApiString)];
		//Obtiene la posicion del primer <
		lastTag = strings.Index(respApiString, "<");
		//Recorte de lo ultimo en respApiString hasta la posicion del primer < encontrado
		estatus := respApiString[0:lastTag];
		fmt.Println("Estatus %s", estatus);
	}
}

func getStatusFromPageNewSATByNumberOfIntegration(numIntegration string) {
	fmt.Println("Obteniendo datos de la pagina de DODA...");
	resp, err := http.Get("https://pecem.mat.sat.gob.mx/app/qr/ce/faces/pages/mobile/validadorqr.jsf?D1=16&D2=1&D3="+numIntegration);
	if err != nil {
        log.Fatalln(err);
	} else {
		asterisks := "***";
		respApiString := convertRespFromHttResponseToString(resp);
		positionInit := strings.Index(respApiString, asterisks);
		positionFinal := strings.LastIndex(respApiString, asterisks);
		status := respApiString[positionInit+len(asterisks):positionFinal];
		fmt.Println("Estatus %s", status);
	}
}

func convertRespFromHttResponseToString(respApi *http.Response) string {
	return convertRespApiToString(convertRespApiToBytes(respApi));
	
}

func convertRespApiToBytes(respApi *http.Response) []uint8 {
	defer respApi.Body.Close();
	bodyBytes, _ := ioutil.ReadAll(respApi.Body);
	return bodyBytes
}

func convertRespApiToString(respApiUnit8 []uint8) string {
	return string(respApiUnit8);
}