# Dependancies
go get -u github.com/gorilla/mux 

go get go.mongodb.org/mongo-driver/bson

go get go.mongodb.org/mongo-driver/mongo

go get go.mongodb.org/mongo-driver/mongo/options

go get go.mongodb.org/mongo-driver/bson/primitive

# Swagger doc
go get -u github.com/swaggo/swag/cmd/swag

go get -u github.com/swaggo/http-swagger

swag init

# Run main.go
go run main.go

# Open URL in Browser
http://localhost:8080/swagger/index.html