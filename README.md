# Prysm<>Geth Proxy

> **Note** ><i> This is my personal experimental project. It is WORK IN PROGRESS. Using Prysm for CL and Geth for EL.</i>
> This design doesn't make sense without [light client support](https://ethereum.stackexchange.com/questions/135812/do-light-clients-still-work-after-merge), since full node will take 20+ hours to sync, which renders this entire approach meaningless. As of March 2023, light clients in EL has completely stopped working. Depending how what/how light client support in EL is implemented, this design will change.

## Context

As of March 2023, consensus clients and execution clients require a 1 to 1 pairing. It's recommended to have both CL and EL be in same root directory, where they share a secret jwt.hex. EL will expose authenticated Engine APIs via rpc server, which CL will call.

## Problem

While this schema works fine for most users, for those running clusters of nodes, handling massive amounts of incoming transactions request, this poses a challenge. Traffic on CL is very consistent and stable, as it's only purpose is to handle consensus protocol and block p2p. Traffic on EL however is very unpredictable, and can spike 10X or 100X. A single EL can only handle so many RPS, and WSS connections. In systems that handle massive traffic and burst traffic, horizontal scaling of EL is required. Post merge, this means that you would simultaneously need to also spin up CL as well EL.

## Objective (1 of many potential solutions)

Create a proxy to enable 1 to many CL : EL system, to allow EL to easily scale horizontally. CL will strictly talk to proxy, and proxy will expose the same Engine APIs and other ETH APIs that CL will require. CL and proxy communication will be authenticated using the same JWT, as if CL<>EL. There will be no code changes required to CL and EL for this implementation to work. From CL pov, it will feel as if its talking to EL, and vice versa.

## Proxy responsibilities

1. Setup Geth's go-ethereum/rpc server, and register Eth & Engine RPC apis
2. Expose authenticated rpc ETH and Engine APIs via http and wss
3. Implement custom routing logic for CL request to downstream EL clients
4. Aggregate responses from downstream EL clients, and return single response to CL
5. Manages a set of EL clients. Handle health check, add/register, and remove EL clients

## Points of consideration

1. Geth uses RLP serialization, and the rpc library is located here go-ethereum/rpc. In order to get RPC server set up, register the neccessary RPC Apis, create httpservers, obtain JWT secret...etc I copied over source code from Geth (since some of the codes are internals or not exported)
2. Right now manually copying over the secret into my proxy. In reality, a key managment service, like AWS KMS, would be used to acquire jwt secret for CL, EL, and proxy.
3. Only supporting http and wss for now
4. This design doesn't make sense without [light client support](https://ethereum.stackexchange.com/questions/135812/do-light-clients-still-work-after-merge), as a full node will take many hours to sync, which renders this entire approach meaningless. As of March 2023, light clients in EL has completely stopped working.

## TODO:

1. Write API interceptors and routing logic
2. Write service to manage EL clients, add, remove, and health checking.
