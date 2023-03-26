![Banner]()
# StarItAll

Star every repo from any github account (orgs and user accounts).

## Installation

Make sure to have [Go installed](https://golang.org/doc/install) and clone the repository.

```bash
git clone https://github.com/KaiDevrim/StarItAll.git
```

Get a GitHub oAuth token by following the steps [here](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token).
Make sure to give it at least the permission to Star repos. ![]()
Then put the token in the .env.example like so:

```
TOKEN=git_pat_
```

Rename the .env.example to just `.env`

Run

```bash
go get -u github.com/KaiDevrim/StarItAll
```

to install all the dependencies.

## Usage

Build the project then run it.

```bash
go build
```

```bash
./StarItAll kaidevrim
```

Replace the last word with any GitHub User or Organization.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
I am not going to update the project anymore, if you need to make a contribution, please open a pr or dm me on Discord. `Kai#0335`

## License

[GNU GPVL3](https://choosealicense.com/licenses/gpl-3.0/)
