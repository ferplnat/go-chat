# go-chat (WIP)

A simple chat application written in Go using UDP. This isn't meant to be a real thing. It's just a fun project to learn Go.

## Usage
Start server:

```bash
go run cmd/server/.
```

Start client:

```bash
go run cmd/client/.
```


## TODO
- [ ] Create protocol to send and receive messages.
  - [ ] Built-in flexibility for multiple protocol versions.
  - [ ] Create generic request/response binary format.
    - [ ] Create message type based on generic request/response format.
    - [ ] Create login request/response.
  - [ ] Create encoding/decoding functions.
    - [ ] Create message parser.
    - [ ] Create message encoder.
  - [ ] Basic encryption for server communication.
    - [ ] Create encryption/decryption functions.
    - [ ] Create encryption/decryption tests.
  - [ ] PGP encryption for client communication.
    - [ ] Create PGP encryption/decryption functions.
    - [ ] Create PGP encryption/decryption tests.
  - [ ] Distinguish request types.

- [ ] Create server to maintain and share messages.
  - [ ] Create message store.
  - [ ] Encryption at rest.

- [ ] Create basic TUI for client.
  - [ ] Create basic TUI library and functions.
  - [ ] Create TUI tests.
  - [ ] Render sent messages.
  - [ ] Render received messages.

- [ ] Create client to send and receive messages.
  - [ ] Enable client to send messages.
  - [ ] Enable client to receive messages.
  - [ ] Enable client to login.
