---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/part-8
title: "Step 8 — Phase C: Backend Frameworks & APIs (part 8)"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["agent", "guardrails"]
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [384, 391]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: 535ebdad437e9a69f051f548214f0f1910ac6017e9c5d10d3a3d8a937b2a261c
---

# Step 8 — Phase C: Backend Frameworks & APIs (part 8)

- Python: SQLAlchemy 2.0.x is the default; Django ORM if you're on Django; Tortoise ORM for pure-async niche.
- TypeScript: Prisma for schema-DSL + guardrails (+ AI-agent skills in 7.9/8); Drizzle for SQL transparency + edge.
- .NET: EF Core 10 — no serious alternative for relational; Dapper for micro-ORM raw speed.
- Java: Hibernate ORM 7.x (blocking) / Hibernate Reactive 3.x (non-blocking); Spring Data JPA, Micronaut Data, Quarkus Panache as the repository facades.
- The 2026 through-line: JSON-column mapping (EF Core 10, SQLAlchemy, Hibernate) and vector-search support (EF Core 10, Prisma extension packs) — ORMs are absorbing the document/AI workload instead of ceding it to ODMs.

---

