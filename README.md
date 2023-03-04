# Storage of private data

This is a backend application created for the secure storage of private data in the Blockchain. This application can be used to store medical information about a person, documents and other personal data.

**The key features of Gin are:**
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