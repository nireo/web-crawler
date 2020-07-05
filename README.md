# Web crawler

> A small web crawler that goes through all the links on a site and then indexes those sites.

## Running code

There is two "modes" in this program; the server which hosts results and the crawler which indexes new sites. Running the code is like with any Go program:

```
go run main.go
```

But there are two flags: startingUrl and display. These can be provided as follows:

```
go run main.go -start="https://github.com/nireo" -display
```

If no starting address is specified the program will host the server. Also if display flag is apparent new indexing results will _not_ be displayed, since they are displayed by default.

## Contributing

You can create a pull request if you're interested in contributing to the project. This is highly encouraged!
