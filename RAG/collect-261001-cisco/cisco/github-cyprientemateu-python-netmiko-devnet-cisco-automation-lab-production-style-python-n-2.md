---
id: collect-261001-cisco/cisco/github-cyprientemateu-python-netmiko-devnet-cisco-automation-lab-production-style-python-n-2
title: "github-cyprientemateu-python-netmiko-devnet-cisco-automation-lab-production-style-python-netdevops-a"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2026-05-18", "2026-05-21", "2026-05-26", "2026-05-31", "2026-07-02"]
keywords: []
source: docs/RAG/collect-261001-cisco/github-cyprientemateu-python-netmiko-devnet-cisco-automation-lab-production-style-python-netdevops-a.md
source_anchor: ""
source_lines: [240, 318]
sha256: c389e449e2476231d2b927dcb563150b835a10f33597b5272210d9c3f13ec213
---

# github-cyprientemateu-python-netmiko-devnet-cisco-automation-lab-production-style-python-netdevops-a

- Extract only the relevant interface block from running config
- Use subset check: verify all expected lines are present, ignore extras
- Exclude `no shutdown` from expected lines

```
2026-07-02 20:03:14 | INFO  | [EXEC-20260702-200309] | NetDevOps Framework — LIVE | SEQUENTIAL
2026-07-02 20:03:14 | WARNING | [EXEC-20260702-200309] | [device-1] [DRIFT] GigabitEthernet1/0/2 → DRIFT
2026-07-02 20:03:16 | INFO  | [EXEC-20260702-200309] | Remediation successful — GigabitEthernet1/0/2
2026-07-02 20:03:22 | INFO  | [EXEC-20260702-200309] | RESTCONF GET successful
2026-07-02 20:03:22 | INFO  | [EXEC-20260702-200309] | [device-1] RESTCONF [GigabitEthernet1/0/2] → COMPLIANT
2026-07-02 20:03:26 | INFO  | [EXEC-20260702-200309] | [device-2] [OK] GigabitEthernet1/0/2 → COMPLIANT
2026-07-02 20:03:27 | INFO  | [EXEC-20260702-200309] | [device-2] RESTCONF [GigabitEthernet1/0/2] → COMPLIANT
```
```
2026-05-31 16:35:43 | INFO  | [EXEC-20260531-163543] | NetDevOps Framework — LIVE | PARALLEL
2026-05-31 16:35:43 | INFO  | [EXEC-20260531-163543] | Connecting to devnetsandboxiosxec9k.cisco.com...
2026-05-31 16:35:43 | INFO  | [EXEC-20260531-163543] | Connecting to devnetsandboxiosxec9k.cisco.com...
2026-05-31 16:35:45 | INFO  | [EXEC-20260531-163543] | [device-2] [OK] GigabitEthernet1/0/2 → COMPLIANT
2026-05-31 16:35:45 | INFO  | [EXEC-20260531-163543] | [device-1] [OK] GigabitEthernet1/0/2 → COMPLIANT
2026-05-31 16:35:46 | INFO  | [EXEC-20260531-163543] | [device-1] Processing complete ✔
2026-05-31 16:35:46 | INFO  | [EXEC-20260531-163543] | [device-2] Processing complete ✔
2026-05-31 16:35:46 | INFO  | [EXEC-20260531-163543] | NetDevOps Framework — Execution Complete
```
```
2026-05-26 18:34:39 | INFO     | [EXEC-20260526-183439] | NetDevOps Framework — DRY-RUN MODE
2026-05-26 18:34:39 | INFO     | [EXEC-20260526-183439] | Inventory validation passed — 1 device(s), 3 interface(s)
2026-05-26 18:34:41 | INFO     | [EXEC-20260526-183439] | [OK]    GigabitEthernet1/0/2 → COMPLIANT
2026-05-26 18:34:41 | INFO     | [EXEC-20260526-183439] | [OK]    GigabitEthernet1/0/3 → COMPLIANT
2026-05-26 18:34:41 | INFO     | [EXEC-20260526-183439] | [OK]    GigabitEthernet1/0/4 → COMPLIANT
2026-05-26 18:34:41 | WARNING  | [EXEC-20260526-183439] | DRY-RUN MODE — remediation skipped, no changes pushed
2026-05-26 18:34:41 | INFO     | [EXEC-20260526-183439] | NetDevOps Framework — Execution Complete
```
```
2026-05-18 22:22:30 | WARNING  | [DRIFT] GigabitEthernet1/0/2 → DRIFT
2026-05-18 22:22:30 | WARNING  | [DRIFT] GigabitEthernet1/0/3 → DRIFT
2026-05-18 22:22:30 | WARNING  | [DRIFT] GigabitEthernet1/0/4 → DRIFT
2026-05-18 22:22:33 | INFO     | Remediation successful — GigabitEthernet1/0/2
2026-05-18 22:22:34 | INFO     | Remediation successful — GigabitEthernet1/0/3
2026-05-18 22:22:35 | INFO     | Remediation successful — GigabitEthernet1/0/4
```
```
2026-05-21 21:37:59 | INFO     | NetDevOps Framework — LIVE MODE
2026-05-21 21:38:02 | INFO     | [OK]    GigabitEthernet1/0/2 → COMPLIANT
2026-05-21 21:38:02 | INFO     | [OK]    GigabitEthernet1/0/3 → COMPLIANT
2026-05-21 21:38:02 | INFO     | [OK]    GigabitEthernet1/0/4 → COMPLIANT
2026-05-21 21:38:02 | INFO     | All interfaces COMPLIANT — no remediation needed
2026-05-21 21:38:02 | INFO     | NetDevOps Framework — Execution Complete
```
- Always backup before changes
- Validate after configuration
- Avoid manual verification
- Ensure idempotent automation
- Use structured reporting
- Externalize inventory and credentials
- Separate concerns (inventory, templating, execution)

- Scheduler as background service (systemd)
- PyATS/Genie validation (via WSL)
- Multi-vendor support
- Intent-based networking

This project demonstrates:

- Real-world network automation design
- Infrastructure as Code thinking
- Validation-first engineering approach
- Data-driven automation (YAML + Jinja2)
- Scalable modular software architecture
- Idempotent execution model (single and multi-device)
- Production-grade observability (logging + dry-run)
- Defensive automation (schema validation before execution)
- Concurrent network automation (parallel multi-device)
- API-native validation (RESTCONF dual-layer compliance)
- Full-stack NetDevOps (automation + web dashboard)
- Continuous compliance (scheduled automation jobs)

Cyprien Carlos Temateu

NetDevOps Practice Lab
