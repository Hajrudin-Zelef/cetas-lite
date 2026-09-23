---
id: etape8-phasef-data-messaging/01-producer-idempotence-durability-baseline/overview
title: "producer: idempotence + durability baseline"
domain: producer-idempotence-durability-baseline
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [592, 596]
section: "producer: idempotence + durability baseline"
sha256: ba3b966990a9b073d5796581c248f31adea28333bf8000a3eb7acd08b6dc6a26
---

# producer: idempotence + durability baseline
enable.idempotence=true
acks=all
retries=2147483647
max.in.flight.requests.per.connection=5
