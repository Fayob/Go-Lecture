package main

import "fmt"

func main() {
	name := "Fayob World"

	temp := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Testing Template</title>
		</head>
		<body>
		<h1>` + name + `</h1>
		</body>
		</html>
	`

	fmt.Println(temp)
}
