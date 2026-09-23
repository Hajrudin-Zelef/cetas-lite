---
id: etape8-phasef-data-messaging/02-consumer-control-your-own-offsets-for-effectively-once/overview
title: "consumer: control your own offsets for effectively-once"
domain: consumer-control-your-own-offsets-for-effectively-once
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["consumer"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [597, 609]
section: "consumer: control your own offsets for effectively-once"
sha256: 9e231ad3e90b8c227d8efad5db5ae913c60ca07981c33642078425530130859c
---

# consumer: control your own offsets for effectively-once
enable.auto.commit=false
isolation.level=read_committed   # when using transactions
```

- With `enable.idempotence=true`, keep `max.in.flight.requests.per.connection <= 5` to preserve ordering guarantees `[secondary]`.

### 18.4 Redis Streams consumer group

```bash
XGROUP CREATE orders stream-orders mygroup $ MKSTREAM
XREADGROUP GROUP mygroup worker-1 COUNT 10 BLOCK 5000 STREAMS stream-orders >
XACK stream-orders mygroup <message-id>
