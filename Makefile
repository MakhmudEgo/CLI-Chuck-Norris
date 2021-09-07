NAME = joker

all:
	go build -o $(NAME) cmd/kek/main.go
clean:
	rm -rfv *.txt $(NAME)
