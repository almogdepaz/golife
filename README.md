# golife

this is an implementation of the game of life
it has an upgrade which allows you to inject the games "next generation" logic


to run the normal mode use

    go run ./cmd/normal/main.go 3 3 30 "1 1 0 1 0 0 1 0 1"

plague mode
    in plague mode the rules change so that

    1. Any dead cell with a single live neighbor lives on to the next generation.
    2. Any live cell with no horizontal or vertical live neighbors dies.

to run plague mode use

    go run ./cmd/plague/main.go 3 3 6 30 "1 1 0 1 0 0 1 0 1"

to implement your own logic use SetNextFunc to determine the games "next generation" logic



setting flag -rand=true crates a random seed