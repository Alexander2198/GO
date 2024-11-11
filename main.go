package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func restarHandler(w http.ResponseWriter, r *http.Request) {
	// Recuperar los valores de los parámetros en la solicitud AJAX
	num1Str := r.URL.Query().Get("num1")
	num2Str := r.URL.Query().Get("num2")

	// Convertir los valores a enteros
	num1, err := strconv.Atoi(num1Str)
	if err != nil {
		http.Error(w, "Número 1 inválido", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(num2Str)
	if err != nil {
		http.Error(w, "Número 2 inválido", http.StatusBadRequest)
		return
	}

	// Realizar la resta
	resultado := num1 - num2

	// Devolver el resultado como texto
	fmt.Fprintf(w, "%d", resultado)
}

func main() {
	// Ruta para la página principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Página principal con el formulario y JavaScript para AJAX
		fmt.Fprintf(w, `<html>
			<body>
				<h1>Calculadora de Resta</h1>
				<form id="calcForm">
					<label for="num1">Número 1:</label>
					<input type="text" id="num1" name="num1"><br><br>
					<label for="num2">Número 2:</label>
					<input type="text" id="num2" name="num2"><br><br>
					<button type="button" onclick="calcularResta()">Calcular</button>
				</form>
				<h2>Resultado: <span id="resultado"></span></h2>
				
				<script>
					function calcularResta() {
						// Obtener los valores de los inputs
						const num1 = document.getElementById("num1").value;
						const num2 = document.getElementById("num2").value;
						
						// Realizar la solicitud AJAX
						fetch("/restar?num1=" + num1 + "&num2=" + num2)
							.then(response => response.text())
							.then(data => {
								// Mostrar el resultado en el span
								document.getElementById("resultado").innerText = data;
							})
							.catch(error => {
								document.getElementById("resultado").innerText = "Error en la operación";
							});
					}
				</script>
			</body>
		</html>`)
	})

	// Ruta para la lógica de la resta
	http.HandleFunc("/restar", restarHandler)

	// Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor iniciado en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}
