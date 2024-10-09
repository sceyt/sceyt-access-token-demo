# Token Generator Demo Go

This is a sample Golang application that demonstrates how to provide JWT access tokens signed with an RSA key for the Sceyt Chat API.


## Prerequisites

- Golang 1.16 or later

## Setup

After cloning the repository, navigate to the project directory and install the required packages using the following commands:

1. **Navigaate to the project**

```bash
cd token-generator-demo-go
```

2. **Install dependencies**

```bash
go mod tidy
```

3. **Copy the example .env file**

```bash
cp .env.example .env
```

Edit the `.env` file and set the `PRIVATE_KEY_PATH` variable to the path of the private key file.
 
```
PRIVATE_KEY_PATH=/path/to/private.pem
```
Example private.pem file will be used for demo purposes, if environment variable is not set.

## Run the Demo

Start the app using:

```bash
go run main.go
```

## Usage

To obtain a token, access the following endpoint:

```bash
http://localhost:3000/get-token?sub={{user_id}}
```

Replace `{{user_id}}` with the user ID for which you want to generate the token.

