---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-22-4
title: "Configuration examples"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-22.md
source_anchor: ""
source_lines: [288, 300]
sha256: 0be60d30c7dc5d0fc3c05d2ee082a25591722a8678ab0a968ec1aee473c1bab4
---

# Configuration examples

Responses to POST method are not cacheable, unless the response includes appropriate Cache-Control or Expires header fields.

#### PUT

This method requests that the enclosed entity be stored under the supplied Request-URI. If another entity exists under specified Request-URI, the enclosed entity should be considered as an updated (newer) version of that residing on the origin server. If the Request-URI is not pointing to an existing resource, the origin server should create a resource with that URI.

If the request passes through a cache and the Request-URI identifies one or more currently cached entities, those entries should be treated as stale. Responses to this method are not cacheable.

#### TRACE

This method invokes a remote, application-layer loop-back of the request message. The final recipient of the request should reflect the message received back to the client as the entity-body of a 200 (OK) response. The final recipient is either the origin server or the first proxy or gateway to receive a Max-Forwards value of 0 in the request. A TRACE request must not include an entity.

Responses to this method MUST NOT be cached.
