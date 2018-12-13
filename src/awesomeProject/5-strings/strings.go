package main

import "strconv"

func main(){
	// Data type String
	// Los Strings son una secuencia de Bytes (Slice de Bytes)
	// Are indexables
	// Are inmutables

	var cadena = "Hola mundo"
	// length of a string
	println(len(cadena))

	println(cadena[2]) // utf-8 format, this sentence must print 108

	println(cadena[0:4]) // slice --> the last index are not inclusive
	println(cadena[:4]) // same result of the previous sentence

	//cadena[2] = "a"  the strings are immutable

	cadena = cadena + "nuevamente"
	println(cadena) // Aquí la cadena no ha cambiado, sólo dejó de apuntar a ese valor para apuntar a otro

	// To keep the format, use this quotation mark
	cadena = `<html> 
				<head>
					<meta charset="utf-8">
						<title></title>
				</head>
				<body>
				</body>
			   </html>`

	println(cadena)

	cadena = "Hola mundo \"25\"" // caracteres especiales con el slash invertido
	println(cadena)

	edad := 29

	// Conversion de tipos
	cadena = "La edad es "
	println(cadena)
}
