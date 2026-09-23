---
id: etape8-phasef-data-messaging/03-crashed-consumer-recovery/overview
title: "crashed consumer recovery:"
domain: crashed-consumer-recovery
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["consumer", "revenue"]
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [610, 631]
section: "crashed consumer recovery:"
sha256: 28fada1799811b698ec3959de7da8c58ed25af1cc0c76dfc0355286c70980486
---

# crashed consumer recovery:
XAUTOCLAIM stream-orders mygroup worker-2 60000 0-0 COUNT 10
```

- `XAUTOCLAIM` reassigns pending entries idle beyond the threshold — the mechanism that makes Streams durable work queues rather than fire-and-forget pub/sub `[independent]`.

### 18.5 dbt model with test

```sql
-- models/marts/orders_daily.sql
{{ config(materialized='incremental', unique_key='day') }}
SELECT date_trunc('day', created_at) AS day,
       count(*) AS orders,
       sum(total) AS revenue
FROM {{ ref('stg_orders') }}
{% if is_incremental() %}
WHERE created_at > (SELECT max(day) FROM {{ this }})
{% endif %}
GROUP BY 1
```

```yaml
