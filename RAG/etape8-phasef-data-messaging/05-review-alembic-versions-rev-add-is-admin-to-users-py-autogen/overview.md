---
id: etape8-phasef-data-messaging/05-review-alembic-versions-rev-add-is-admin-to-users-py-autogen/overview
title: "REVIEW alembic/versions/<rev>_add_is_admin_to_users.py -- autogenerate is a draft"
domain: review-alembic-versions-rev-add-is-admin-to-users-py-autogen
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [669, 682]
section: "REVIEW alembic/versions/<rev>_add_is_admin_to_users.py -- autogenerate is a draft"
sha256: aab0ca169c8505e86c0e9920c0aff3c07dd1a6e77c6d91c9845b4a600f2e5cdc
---

# REVIEW alembic/versions/<rev>_add_is_admin_to_users.py -- autogenerate is a draft
alembic upgrade head
alembic current && alembic heads   # verify single head, no branches
```

### 18.8 Spark Connect thin client

```python
from pyspark.sql import SparkSession
spark = (SparkSession.builder
         .remote("sc://spark-prod:15002")   # thin client, no cluster jars locally
         .getOrCreate())
spark.sql("SELECT date_trunc('day', ts) AS day, count(*) "
          "FROM events GROUP BY 1").show()
