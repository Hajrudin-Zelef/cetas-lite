---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/wave-10-netbox-labs-portfolio-expansion-ai-discovery-and-ing
title: "Wave 10 — NetBox Labs portfolio expansion: AI, discovery, and ingestion"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: ["2025-04", "2025-07", "2026-02", "2026-02-10", "2026-05-24"]
keywords: ["agent", "agentic", "agents", "claude", "copilot", "funding", "governance"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [321, 379]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: 38219fc6cbd4eb65bffd382e1f1818c0742a36029e1adeafde3b61d1e1f13269
---

# Wave 10 — NetBox Labs portfolio expansion: AI, discovery, and ingestion

## Wave 10 — NetBox Labs portfolio expansion: AI, discovery, and ingestion

### 10.1 NetBox Copilot — GA 2026-02-10 (agentic AI) `[vendor-reported]`

Source: https://www.globenewswire.com/news-release/2026/02/10/3235154/0/en/NetBox-Labs-Announces-General-Availability-of-NetBox-Copilot-Delivering-Enterprise-Ready-AI-Grounded-in-Accurate-Infrastructure-Data.html

- On 2026-02-10 (New York), NetBox Labs announced **general availability of NetBox Copilot**, "an interactive AI agent embedded directly into the NetBox platform".
- Positioning: as AI agents move from experimentation to production operations, most agents "lack the accurate operational context required to be trusted in production environments"; NetBox is framed as "a comprehensive semantic map for infrastructure" that grounds Copilot with complete infrastructural context — "agentic AI… providing trustworthy answers, reliable automation and enterprise-ready governance — all conducted in natural language".
- GA expanded Copilot from inquiry into **workflow execution**: teams can validate data completeness/quality ("Which devices are missing IP addresses?"), investigate changes for troubleshooting/compliance ("Who changed this prefix last week?"), assess impact before maintenance ("What depends on this switch?"), plan capacity ("How much rack space is available in Denver?"), and build automation workflows/playbooks ("Create an Ansible playbook to deploy switch configs").
- Target: "thousands of NetBox users", expanding access to infrastructure data to security/operations teams, executives, and non-IT teams.
- **Naming note:** a July 2025 SiliconANGLE report described NetBox Labs' agentic AI operations tool as **"NetBox Operator"** (https://siliconangle.com/2025/07/14/netbox-labs-central-nervous-system-ai-data-centers-gets-35m-funding/). The February 2026 GA announcement calls the product **NetBox Copilot**. Whether this is a rename of the same product or a distinct offering is not confirmed in the sources found — registered as a naming question in the gaps log `[unverified]`.

### 10.2 NetBox Discovery and the Diode ingestion stack `[vendor-reported][official]`

- **NetBox Discovery**: observability agent for network/device discovery; included with NetBox Assurance; agent extensions in Standard and Premium bundles for Professional/Enterprise tiers (per the April 2025 Assurance announcement).
- **Diode**: source-available project from NetBox Labs offering a subset of NetBox Assurance functionality; shares a common ingestion API with Assurance via the **Diode SDK** — an alternative interface for sending data to NetBox with built-in idempotence, automatic ordering, and other capabilities for high-performance integrations `[vendor-reported]`.
- Architecture implication: Diode gives integrators a purpose-built, idempotent ingestion path into NetBox as an alternative to raw REST bulk writes (which gained `?background=true` async processing in v4.7) `[official]` on the v4.7 API part.

### 10.3 NetBox Cloud tiers and free tier `[vendor-reported]`

- NetBox Cloud is the SaaS offering (SOC 2 compliant infrastructure); dev/staging environments included at Professional tier and above; Assurance appears as an add-on in Cloud navigation `[vendor-reported]`.
- A **free NetBox Cloud** tier is advertised (https://netboxlabs.com/products/free-netbox-cloud/) — plan limits and feature gates not captured in this wave (see gaps log) `[unverified]` on specifics.
- NetBox Enterprise: self-managed enterprise offering; airgapped installations offered; bundled discovery/assurance integrations (Forward Networks, IP Fabric, Slurp'it) reduce adoption barriers `[vendor-reported]`.

---

## Wave 11 — Deployment & operations: NetBox and Nautobot in production

### 11.1 NetBox deployment topologies `[official][secondary]`

- **Reference Docker deployment:** the community maintains `netbox-community/netbox-docker` (referenced as the official Docker project in third-party deployment docs) `[secondary]`. Standard service stack observed across deployments: `netbox` (gunicorn/uvicorn app server, healthcheck on `/login/`), `netbox-worker` (RQ background worker), `netbox-housekeeping` (scheduled cleanup), `postgres`, `redis` (task queue), `redis-cache` (application cache) `[secondary]`.
- **Version requirements (v4.7, 2026):** PostgreSQL 15+, Redis 6.0+, Python 3.12–3.14, Django 6.1 `[official]`.
- **Image variants in the wild (2026):**
  - `linuxserver/docker-netbox` (`lscr.io/linuxserver/netbox`) — x86-64 and arm64 images; env-var configuration (DB/Redis connection, remote-auth, CSRF trusted origins); WebUI on port 8000. Source: https://github.com/linuxserver/docker-netbox/blob/HEAD/README.md
  - `11notes/docker-netbox` — hardened compose: postgres 16, redis 7.4.2, NetBox 4.3.3, read-only containers, cron-driven backup job. Source: https://github.com/11notes/docker-netbox/blob/HEAD/README.md
  - `swissmakers/netbox-plus` — Podman Compose variant: UBI 9 base, Python 3.12, PostgreSQL 16, Redis 7, RQ worker with high/default/low queues; rootless-friendly. Source: https://github.com/swissmakers/netbox-plus/blob/HEAD/docker/README.md
  - Homelab K3s deployment (2026-05-24): NetBox 4.6.1 on K3s behind Traefik, written by Terraform via `e-breuninger/netbox` provider `[secondary]`. Source: https://github.com/xiiisins/homelab/blob/HEAD/docs/services/netbox.md
- **Diode-integrated deployment example:** an agent-cloud stack builds on netbox-docker with the `netboxlabs-diode-netbox-plugin`, OAuth2 client secrets mounted at `/run/secrets`, PostgreSQL 18 + Valkey 9.0 (separate task-queue and cache instances), 12 compose containers — illustrating the NetBox + Diode discovery pipeline in production-shaped form `[secondary]`. Source: https://github.com/uhstray-io/agent-cloud/blob/HEAD/platform/services/netbox/deployment/CLAUDE.md

### 11.2 NetBox upgrade & backup operations (2026 notes) `[official]`

- Upgrade path uses the `upgrade.sh` script plus Django management commands (`migrate`, `post_upgrade`); v4.7 requires a planned maintenance window: the django-mptt→ltree migration takes `ACCESS EXCLUSIVE` locks and backfills rows (minutes on large DBs), then `rebuild_config_context_cache` issues one `UPDATE` per device/VM; the migration is not reversible in practice `[official]`.
- **v4.7.1 restore caveat:** restoring a `pg_dump` taken on v4.7.0 recreates the DB without ltree cascade triggers — upgrade reinstalls them but does not repair already-stale hierarchical paths; a documented repair procedure exists; plugins with their own ltree models need migration steps `[official]`.
- Background worker discipline on upgrade: webhook jobs still enqueued when workers restart fail with `TypeError` (v4.7) — allow queues to drain before upgrading; global search index updates are now deferred to a background task `[official]`.
- Python/dependency floor moves in 2026: v4.5 dropped Python 3.10/3.11 (requires 3.12–3.14); v4.7 dropped PostgreSQL 14 and Redis 5.x; SSO deployments should re-test auth (social-auth major upgrades) `[official]`.

### 11.3 Nautobot deployment notes `[official][secondary]`

- Reference deployment: Docker Compose and Kubernetes (Helm); documented K8s workflow with Flux, custom image layering apps (Golden Config + Nornir plugin via `requirements.txt`), and `nautobot_config.py` management. Nautobot 2.4 added a Kubernetes Job Execution and Job Queue data model `[official]`. Source: https://networktocode.com/blog/deploying-nautobot-to-kubernetes-03/
- **Nautobot 3.1 upgrade notes:** requires PostgreSQL ≥ 14 (drops 12.x/13.x); Django 5.2 unified `STORAGES` setting replaces `DEFAULT_FILE_STORAGE`/`STATICFILES_STORAGE` and Nautobot's `STORAGE_BACKEND`/`STORAGE_CONFIG`/`JOB_FILE_IO_STORAGE`; Python 3.14 supported in the app cookiecutter `[official]`.
- Jobs run on workers with scheduling, approval workflows (3.0), console logging (3.1), and cancel (3.2) — the operational surface for automation is the Jobs UI/API rather than shell scripts `[official]`.

### 11.4 Auth & enterprise integration `[official][secondary]`

- NetBox: SSO via python-social-auth (OIDC/SAML; v6.0/v5.1 in v4.7 — retest before upgrading), LDAP, remote-auth headers; object-level permissions with constraints; `render_config` permission (v4.5+) `[official]`.
- Nautobot: SSO, approval workflows for jobs/changes (3.0), Data Validation Engine in core (3.0), secrets groups for credential management `[official]`.

---

