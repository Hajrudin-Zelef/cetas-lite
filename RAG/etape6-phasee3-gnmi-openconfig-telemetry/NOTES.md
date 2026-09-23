# NOTES — corpus `etape6-phasee3-gnmi-openconfig-telemetry`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-gnmi-openconfig-telemetry` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 6 — Phase E3: gNMI / OpenConfig / Model-Driven Telemetry & Programmability` — 751 lignes, 14 chunks.

- Wave 1 — OpenConfig 2026: models, working group, and vendor adoption
- Wave 2 — gNMI protocol: specification, RPCs, and subscription modes
- Wave 3 — Vendor implementations: streaming telemetry & gNMI/gNOI support (2026)
- Wave 4 — Open-source collectors: gnmic, Telegraf, Prometheus exporters, pmacct
- Wave 5 — Streaming telemetry pipeline architectures
- Wave 6 — NETCONF/RESTCONF vs gNMI in 2026; YANG Suite
- Wave 7 — Real-time topology tools and NetBox/Nautobot integration
- Wave 8 — AI-era telemetry: RoCE fabrics, PFC/ECN watch, congestion visibility
- Wave 9 — Coverage audit, open items, and close
- Wave 10 — OpenConfig model catalog: key models, versions, and coverage gaps
- Wave 11 — gNMI deep-dive: encodings, auth, CLI patterns, target discovery
- Wave 12 — Vendor matrix: gNMI/OpenConfig support at a glance (2026)
- Wave 13 — Security: TLS, certs, auth for telemetry planes
- Wave 14 — Adjacent tooling and cross-references
- Wave 15 — Final audit: scope coverage, gaps, and close
- Wave 16 — Arista deep-dive: EOS streaming, CloudVision NetDL, AVA, AI observability
- Wave 17 — gNMI Set: config management over gRPC
- Wave 18 — Deployment playbook: building the stack in practice
- Wave 19 — gRPC tunnel dial-out: the third pattern
- Wave 20 — OpenConfig model-structure walkthroughs (operator's map)
- Wave 21 — Nokia EDA and event-driven network operations
- Wave 22 — Troubleshooting playbook: telemetry failures
- Wave 23 — OpenConfig 2026 process: how the models evolve
- Wave 24 — Outlook and final notes
- Wave 25 — Juniper JTI dial-out configuration anatomy
- Wave 26 — Collector configuration patterns (Telegraf, Fluent Bit, JTI-native)
- Wave 27 — LLDP topology: the OpenConfig data model and reconciliation loop
- Wave 28 — NETCONF vs RESTCONF vs gNMI: operation-by-operation
- Wave 29 — Prometheus metric patterns for network telemetry
- Wave 30 — Closing supplement: what changed since Wave 15's audit
- Wave 31 — Quick-reference card: ports, transports, defaults
- Wave 32 — Glossary and acronym map (telemetry domain)
- Wave 33 — Cisco IOS-XR MDT dial-out and Dell OS10 sensor-path anatomies
- Wave 34 — Research method note and provenance ledger
- Wave 35 — Operator's gNMI path cheat sheet (OpenConfig-first, vendor fallback)
- Wave 36 — Errata and conflict register (detail)
- Document control

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
