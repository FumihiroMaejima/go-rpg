module go-rpg

go 1.16

replace go-rpg/controllers/usersController => ./src/controllers/usersController

replace go-rpg/db => ./src/db

replace go-rpg/repositories/usersRepository => ./src/repositories/usersRepository

replace go-rpg/util => ./src/util

replace go-rpg/services/usersService => ./src/services/usersService

replace go-rpg/web => ./src/web

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	go-rpg/controllers/usersController v0.0.0-00010101000000-000000000000 // indirect
	go-rpg/db v0.0.0-00010101000000-000000000000
	go-rpg/repositories/usersRepository v0.0.0-00010101000000-000000000000 // indirect
	go-rpg/services/usersService v0.0.0-00010101000000-000000000000 // indirect
	go-rpg/util v0.0.0-00010101000000-000000000000
	go-rpg/web v0.0.0-00010101000000-000000000000
)
