set -e

# load environment variables written in .env file
if [ -f ".env" ]; then
  export $(xargs < .env)
fi

# create log directory if not exists
mkdir log -p

# compile and execute main.go
go run main.go