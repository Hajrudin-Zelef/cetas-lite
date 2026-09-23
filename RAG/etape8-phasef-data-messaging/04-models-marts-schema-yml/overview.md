---
id: etape8-phasef-data-messaging/04-models-marts-schema-yml/overview
title: "models/marts/schema.yml"
domain: models-marts-schema-yml
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape8_phaseF_data_messaging.md
source_anchor: ""
source_lines: [632, 668]
section: "models/marts/schema.yml"
sha256: 23a0bb437f69ae77a56973bf77130bd015ff4affc030fa0ff7421d17473d4613
---

# models/marts/schema.yml
models:
  - name: orders_daily
    tests:
      - dbt_utils.unique_combination_of_columns:
          combination_of_columns: [day]
```

### 18.6 Airflow 3 TaskFlow DAG

```python
from airflow.sdk import dag, task
from datetime import datetime

@dag(schedule="@daily", start_date=datetime(2026, 1, 1), catchup=False)
def etl_sales():
    @task
    def extract() -> list[dict]:
        return fetch_from_api()          # runs at execution time

    @task
    def transform(rows: list[dict]) -> list[dict]:
        return [clean(r) for r in rows]  # XCom passes rows implicitly

    @task
    def load(rows: list[dict]) -> None:
        warehouse.insert(rows)

    load(transform(extract()))           # dependencies are implicit

etl_sales()
```

### 18.7 Alembic autogenerate workflow

```bash
alembic revision --autogenerate -m "add is_admin to users"
