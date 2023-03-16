# Storage of private data

This is a backend application created for the secure storage of private data in the Blockchain. This application can be used to store medical information about a person, documents and other personal data.

**The key features of app are:**
* Free to read data
* Security
* The possibility of obtaining information to authorized persons (e.g doctors)
* Easy to use

## How it works:
The user's information is pre-encrypted on the backend using a password, after which it gets into the blockchain.
Smart-contract uses a decryption algorithm similar to that on the backend, so that
the user can directly get information from the blockchain using a password.
At any time and place, absolutely free.


## Getting started

### Prerequisites
- Go: any one of the three latest major releases.

### Getting app

Create .env file and fill it using example from .envExample to configure app

### Running app

run ```go run cmd/main.go```

### Endpoints

**Save information:**

`/api/fillUserInfo?password=<PASSWORD>&address=<ADDRESS>&data=<DATA>`

- PASSWORD - password to encrypt information, do not lose it.
- ADDRESS - the address of the user's account to which the information belongs.
- DATA - information that needs to be encrypted and saved.

_Response - TX hash._

**Get information:**

`/api/getUserInfo?password=<PASSWORD>&address=<ADDRESS>`

- PASSWORD - password to decrypt information.
- ADDRESS - the address of the user whose information you want to receive.

_Response - decrypted data._
