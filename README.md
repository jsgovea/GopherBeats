# Welcome

## Install go

To install Go on windows download the latest release 
[offical page](https://go.dev/dl/go1.21.6.windows-amd64.msi)

Restart the terminal/cmd. Verify the installation by typing in your cmd
```bash 
go version
```

### Run program
Clone the repository
```bash
git clone https://github.com/jsgovea/GopherBeats.git
```

To ensure you have the dependencies for the project by running
```bash
go mod download
```
If nothing prints to the console it means everything is installed. 

Once the required packages are installed run the program
```bash
go run main.go
```

Use your preferred browser to open the site
```bash
http://localhost:8080/
```

## Project structure
A regular MVC pattern for simplicity. 

```
.
├── controllers               # Application logic
├── models                    # Database models
├── static                    # Everything related to JS & CSS goes here
├── utils                     # Tools and utilities
├── views                     # Server-rendered HTML 
├── main.go                   # Compiles the application and run it
└── README.md
```