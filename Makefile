CXX = g++
CXXFLAGS = -Wall -Werror -Wextra -pedantic -std=c++17 -g -fsanitize=address
LDFLAGS =  -fsanitize=address

SRC = 
OBJ = $(SRC:.cc=.o)
EXEC = go run app/main.go

run:
	go run app/main.go

test:
	go test ./Server/handler -cover
	go test ./Server/usecase -cover
	go test ./Server/repository -cover