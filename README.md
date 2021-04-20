# StarItAll

Star every repo from any github account (orgs and user accounts).

## Installation

Make sure to have go installed, preferably the latest version of go.
You can install go [here](https://golang.org/doc/install).

Get a GitHub Oauth token by following the steps [here](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token).

Then put the token in the .env.example like so:

```go
TOKEN=jkdfgjjlkdfjkldfjkldfjkjkl
```

Rename the .env.example to just `.env`

Run

```go
go get -u github.com/KaiDevrim/StarItAll
```

to install all the dependencies.

## Usage

```bash
go run main.go kaidevrim #replace the last word with any user/org you want to star
```

You could alternatively build the project then run it.

```go
go build
```

in the current directory

Then

```bash
./StarItAll kaidevrim #replace the last word with any user/org you want to star
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
I am not going to update the project anymore, if you need to make a contribution, please open a pr and dm me on Discord. `Devrim#9999`

## License

[GNU GPVL3](https://choosealicense.com/licenses/gpl-3.0/)
