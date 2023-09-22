# Token Generator Demo Nodejs

This is a sample express.js application that demonstrates how to provide JWT access tokens signed with an RSA key for the Sceyt Chat API.


## Prerequisites

- Node.js
- npm or yarn

## Setup

After cloning the repository, navigate to the project directory and install the required packages using the following commands:

1. **Navigaate to the project**

```bash
cd token-generator-demo-nodejs
```

2. **Install dependencies**

```bash
npm install
```
or

```bash
yarn install
```

## Run the Demo

Start the app using:

```bash
npm start
```

or

```bash
yarn start
```

## Usage

To obtain a token, access the following endpoint:

```bash
http://localhost:3000/get-token
```

## Security Note

This is a sample application. In real-world scenarios, always ensure you keep the `private.pem` file secure. Exposing the private key might jeopardize the security of the service. When deploying in production, consider employing environment variables, secure vaults, or other mechanisms to safeguard the key.

