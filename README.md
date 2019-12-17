# Ably Number Generation

## Server Usage

The two servers are written in go and can be built from the root of the repository with:

<code>make build</code>

This should produce two binaries. `stateless-server` and `stateful-server`

To run a stateful server on localhost and port 9000 run:

<code>./stateful-server -port=9001</code>

To run a stateless server on localhost and port 9001 run:

<code>./stateless-server -port=9000</code>

## Client Usage

There is a single python client in the `client` directory that can be used for both stateful and stateless connections.

To run a stateless client to a server on port 9001:

<code>python3 client.py stateless 9001</code>

To run a stateful client to a server on port 9000:

<code>python3 client.py stateful 9000</code>

The client will randomly choose a sequence length between 0 and 0xffff. To manually choose a sequence length use `-l` option.

<code>python3 client.py stateful 9001 -l 3</code>

## Protocol

Both stateful and stateless protocols are built on gRPC to allow easy error handling and streaming support.

### Stateless

The proto definitions of the stateless protocol are defined in `protos/stateless/stateless.proto`.

Initial connections and re-connections both use the `GenerateSequenceRequest` RPC. Re-connections must specify the number in the sequence that they last received in the `restartFrom` field. The server will then stream from the next number in the sequence.

Once a stream is established the server will stream a number to the client every 1 second. The initial number is randomly chosen between 0 and 0xff and subsequent numbers are twice the previous number.

The client must keep track of how many numbers it has received and close the connection once it has received sufficient.

### Stateful

The proto definitions of the stateful protocol are defined in `protos/stateful/stateful.proto`.

Initial connections use the `GenerateSequenceRequest` RPC. Clients must specify a UUID and the desired sequence length.

If an established stream drops clients should use the `ReconnectSequenceRequest` RPC to re-connect. Clients must specify the connection UUID and the index of the last received item. The server will then stream from the next item.

The server will initiate a PRNG for the connection and stream a number from it every 1 second. 

When the server streams the final number in the sequence it will set the `finalItem` flag to `true` and include a checksum of the sequence.

The client is responsible for closing the connection on receiving this final message.

#### PRNG

Each stateful connection has a PRNG created for it seeded by the unix time at connection. 

The PRNG implementation is provided by https://golang.org/pkg/math/rand.

#### Checksum

The checksum is calculated by taking the SHA256 hash of the concatenated UTF8 string formats of each element in the sequence.

#### State Store

For stateful connections the seed used to generate the PRNG and the requested sequence length are stored in a key value store keyed on the connection UUID.

The seed allows the PRNG to be reconstructed on reconnection so the same numbers can be generated across reconnections.

The sequence length being stored is a convenience for client.

The store has a built in expiry on entries of 30 seconds.
