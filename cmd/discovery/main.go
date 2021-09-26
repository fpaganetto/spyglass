package main

var app *application

func main() {
	cnf := NewConfig()
	app = &application{
		config: cnf,
	}

	app.initK8Client()
	app.start()
}
