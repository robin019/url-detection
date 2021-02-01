# load environment variables written in .env file
export $(xargs < .env)

# compile and execute main.go
go run main.go