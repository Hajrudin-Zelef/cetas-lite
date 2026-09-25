---
id: collect-250926-servers-hardware/servers-hardware/rpc
title: "RPC"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean4/rpc.md
source_anchor: ""
source_lines: [1, 135]
sha256: ca1910ba6de914a0568d9df0567b9f4773057bc939309e9cd8ab0e520a24db92
---

# RPC

Plugins can expose custom methods and events that run on the server and can be called by other plugins or clients.

## Define

Use `Rpc.define` to list the RPCâs methods, errors, and events.

### Validation

An RPC definition describes the shapes of input and output of methods, events and errors. It supports two schema formats:

- JSON Schema, simple and requires no dependencies.
- Any Standard Schema compliant validator
  - Zod
  - Valibot
  - ArkType

### Input and output

Each method can define an `input` schema for its argument and an `output` schema
for its return value. Leave either one out if the method does not accept or
return a value.

```
search: {
  input: {
    type: "object",
    properties: { query: { type: "string" } },
    required: ["query"],
  },
  output: {
    type: "object",
    properties: { text: { type: "string" } },
    required: ["text"],
  },
}
```
JSON Schema values are `unknown` in TypeScript, so narrow them before use.
Standard Schema infers the input and output types.

### Errors

Add an `errors` map to a method for expected failures. Each key becomes the
errorâs `type`, and its schema defines the errorâs `data`.

```
errors: {
  not_found: {
    type: "object",
    properties: { query: { type: "string" } },
    required: ["query"],
    additionalProperties: false,
  },
}
```
Error names beginning with `rpc.` are reserved by OpenCode.

### Events

Add events to the top-level `events` map. Each event has a schema for the data
sent to subscribers.

Event data must be an object. Use an empty object schema when there is no data:

```
events: {
  refreshed: {
    schema: { type: "object", additionalProperties: false },
  },
}
```
Scalars, arrays, `null`, and `undefined` are not valid event data.

## Implement

Register the implementation inside `setup`:

After registering the RPC, the same plugin can call it through `ctx.rpc(Acme)`.

The second argument includes `signal` for cancellation and `context.error(...)`
for declared errors. You can return or throw the error.

One plugin can register more than one RPC. Disposing the registration removes it.

## Call

Once the RPC is registered, it can be called over HTTP or from another plugin.

### HTTP

Create an OpenCode client, then use `client.rpc` to create a subclient for the
RPC:

```
import { OpenCode } from "@opencode/client"
import { Acme } from "opencode-acme-plugin/rpc"
const client = OpenCode.make({
  baseUrl: "http://localhost:4096",
})
const acme = client.rpc(Acme)
const result = await acme.search({ query: "hello" })
```
### Plugin

Plugins already have an OpenCode client. For example, a TUI plugin can create
the same RPC subclient from `context.client`:

### Subscribe

Use `events.on` for a callback and unsubscribe when the listener is no longer
needed:

```
const unsubscribe = acme.events.on("updated", (event) => {
  console.log(event.type, event.location.directory, event.data.text)
})
unsubscribe()
```
For an async iterable, use `events.subscribe("updated")`:

```
for await (const event of acme.events.subscribe("updated")) {
  console.log(event.data.text)
}
```
Event subscriptions are live only, so disconnected subscribers miss events.

- Subscribe with the local name, such as `updated` .
- The received type is prefixed, such as `rpc.acme.updated` .
- Each event includes `data` and`location` .
- Plugin unload closes its subscriptions.

External clients receive events from every location, so check `event.location`
when needed. The server plugin still needs to be configured and running.
