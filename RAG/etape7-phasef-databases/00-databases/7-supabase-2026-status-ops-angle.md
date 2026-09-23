---
id: etape7-phasef-databases/00-databases/7-supabase-2026-status-ops-angle
title: "7. Supabase — 2026 status (ops angle)"
domain: step-7-phase-f-databases-cache-operations-angle
role: deep-dive
task: reference
actors: ["Anthropic", "Google", "Stripe"]
dates: ["2026-06"]
keywords: ["claude", "compute", "cost", "governance", "memory", "open source", "pricing", "valuation"]
source: docs/RAG/etape7_phaseF_databases.md
source_anchor: ""
source_lines: [371, 415]
section: "Step 7 — Phase F: Databases & Cache (Operations Angle)"
sha256: 3df52faf94c2646636d9dd1977d89d0f5aa4c95cba79517a0ffae6080a9bce94
---

# 7. Supabase — 2026 status (ops angle)

## 7. Supabase — 2026 status (ops angle)

- **Fundraising/valuation**: in June 2026 Supabase raised **$500M at a $10.5B
  valuation** (investors cited: GIC, Accel, Coatue, Stripe); valuation doubled
  in eight months, ~5x in fourteen. Database launches grew **>600% in one
  year**, and **>60% of new databases were launched by an AI tool**
  (Claude Code named the single largest contributor) — a cost-governance
  signal for teams [independent].
- **Architecture** (self-described "open-source Firebase alternative"):
  managed **PostgreSQL** + **GoTrue** (auth) + **PostgREST** (auto REST API)
  + **Realtime** (subscriptions) + **Storage** (S3-compatible) + Edge
  Functions + pgvector + Row Level Security [secondary].
- **Pricing, verified against 2026 published pages** (USD, monthly, no annual
  discount) [secondary]:
  - **Free $0**: 2 projects/org, 500 MB DB, 50K MAU, 5 GB egress, 1 GB file
    storage, 500K edge-function invocations, unlimited API requests.
    **Projects pause after 7 days of inactivity** — unsuitable for
    production uptime.
  - **Pro $25/org/mo**: 100K MAUs, 8 GB DB, 250 GB egress, 100 GB storage,
    2M invocations, daily backups with **7-day retention**, $10 compute
    credits; usage overages on top (egress **$0.09/GB**, extra MAUs
    **$0.00325/MAU**).
  - **Team $599/mo**: same resource envelope as Pro; adds SOC 2 + ISO 27001,
    SSO, 14-day backups, 28-day log retention, priority support/SLA.
  - **Enterprise**: custom — dedicated support manager, uptime SLAs, BYO
    cloud, HIPAA, PITR add-on.
  - Real-world production bills commonly land **$35–200/mo** once compute,
    egress and MAU overages stack [secondary].
- **Self-hosting**: the full stack is open source (supabase/supabase repo)
  and Docker-Compose self-hostable; ops cost is running Postgres +
  supporting services yourself, losing managed backups/pooling/upgrades
  [secondary].
- **Alternatives** (2026): **Appwrite**, **PocketBase** (single-binary
  Go, SQLite-backed — far simpler ops), **Nhost**, **Firebase** (Google),
  **Convex**; Grobase-style comparison tables pair Supabase Pro ($25) as the
  mid-market anchor [secondary].
- **Ops caveats**: free-tier pausing; egress as the surprise cost driver;
  AI-generated projects multiplying orgs/projects (sprawl); RLS misconfig
  as the #1 self-inflicted data-exposure vector — RLS policies must be part
  of code review/CI [secondary].

---

## 8. Redis / Valkey and the in-memory ecosystem

