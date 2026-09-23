---
id: etape6-phasee4-validation-cicd/00-front-matter/e4-15-3-lab-platform-edition-pricing-snapshot-sept-2026
title: "E4.15.3 Lab platform edition/pricing snapshot (Sept 2026)"
domain: front-matter
role: reference
task: pricing
actors: ["AWS"]
dates: ["2026-06"]
keywords: ["pricing", "aws", "licenses", "training"]
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [596, 646]
section: "Phase E4 — Network Validation, Observability & CI/CD"
sha256: 62e248e19ba415d8f08723f659f5ac10b9241f7b4f856031e83fdb68504d2fd9
---

# E4.15.3 Lab platform edition/pricing snapshot (Sept 2026)

### E4.15.3 Lab platform edition/pricing snapshot (Sept 2026)

| Platform | Free tier | Paid tiers | Notes |
|---|---|---|---|
| Containerlab | Fully free (MIT) | — (support via Nokia) | Image licenses separate |
| EVE-NG | Freemium: 7 nodes | Pro (per vendor pricing) | Community EOL Jun 2026 |
| GNS3 | Fully free (GPL) | — | 2.2.x stable; 3.x alpha |
| CML | Free: 5 nodes | Personal $199/yr (20), Plus $349/yr (40), Enterprise (300) | 1-yr subscription |
| netlab (ipspace) | — (uses Containerlab) | Training-linked | Min clab 0.75.0 |

- Pricing rows `[secondary]`; verify before purchase `[unverified]`.

### E4.15.4 pyATS release line (2026 context)

- 24.7: 38 new parsers, 71 new APIs; SONiC ops support; Cat8200 Clean `[official]`.
- 24.8: 78 new parsers, 31 new APIs; password recovery in Clean connect stage; Cat8300/Cat9800-CL/IR1800 Clean `[official]`.
- 26.5: in use at Cisco Live 2026 (June 2026); full package list captured in §E4.7.1 `[official]`.
- Python floor: 3.8 dropped after Oct 2024 `[official]`.

### E4.15.5 EDA event-source matrix

| Source plugin | Event origin | Typical network use |
|---|---|---|
| `webhook` | HTTP POST | Alertmanager, generic alerts |
| `kafka` | Kafka topic | High-volume telemetry/alerts |
| `alertmanager` | Prometheus AM | Firing alerts |
| `aws_cloudtrail` / `aws_sqs_queue` | AWS | Cloud network events |
| `azure_service_bus` | Azure | Cloud events |
| `mqtt` | MQTT broker | IoT/device events |
| `pg_listener` | PostgreSQL NOTIFY | DB-driven triggers |
| `file_watch` | Filesystem | Log-driven |
| `journald` | systemd journal | Host events |
| `url_check` | HTTP polling | Reachability state changes |
| `tick` | Timer | Periodic checks |
| `generic` / `range` | Static | Testing |

- As of `ansible.eda` 2.11.0 `[secondary]` (SLAC 2026).

### E4.15.6 CI/CD stage gates for network changes

| Stage | Gate | Tool examples |
|---|---|---|
| Lint | Syntax/YAML/Jinja valid | yamllint, ansible-lint |
| Static validation | No policy violation | Batfish, OPA/Conftest |
| Unit | Templates render | pytest, Jinja tests |
| Lab test | Behavior correct | Containerlab/CML + Robot/pyATS |
| Review | Human approval | PR + plan artifact |
| Canary | Limited blast radius | Staged apply |
| Post-check | Observed == intended | SuzieQ, IP Fabric |
| Remediate | Auto-fix on anomaly | EDA, StackStorm |

