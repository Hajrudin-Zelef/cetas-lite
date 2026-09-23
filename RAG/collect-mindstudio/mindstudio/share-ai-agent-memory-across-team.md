---
id: collect-mindstudio/mindstudio/share-ai-agent-memory-across-team
title: "How to Share AI Agent Memory Across a Team Without Exposing Private Data"
domain: mindstudio
role: reference
task: article
actors: []
dates: ["2026-09-23"]
keywords: ["agent", "memory", "agents", "embedding", "embeddings", "incident", "liability"]
source: docs/RAG/Collect RAG/02_mindstudio/share-ai-agent-memory-across-team.md
source_anchor: ""
source_lines: [1, 80]
sha256: a5d80b37b4d78e515aab4cb5729f797e7cdc59528ccd9882f2cfecc702ff6001
---

# How to Share AI Agent Memory Across a Team Without Exposing Private Data

## Metadata

- **Source** : https://www.mindstudio.ai/blog/share-ai-agent-memory-across-team
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article presents a practical architecture for sharing AI agent memory across a team while protecting private data — one of the harder design problems in enterprise AI. It's not just technical; it's a trust challenge: get it wrong and you've built a system that leaks information in subtle, hard-to-audit ways. The solution: separate shared from private memory, enforce permissions at the data layer with row-level security (RLS), and use permission-mirrored repositories.

What agent memory actually means: models are stateless; "memory" is a retrieval system fetching relevant info into the prompt. Forms: conversation history (log of prior messages), semantic memory (embeddings in a vector DB, similarity search), structured knowledge (facts in relational/key-value stores), episodic memory (records of past events/decisions). Each type creates different privacy risks when shared.

The core problem: shared infrastructure (single Supabase project, Pinecone index, or PostgreSQL DB) with unequal access. A naive implementation filters at the application layer ("only retrieve where user_id = current_user" in agent code) — fragile, because a misconfigured agent, prompt injection, or a bug could return records it shouldn't. The fix: push access control down to the database layer so unauthorized reads are structurally impossible.

Row-level security (RLS): database-engine-enforced row filtering based on identity — not application logic. Supabase exposes PostgreSQL's native RLS integrated with its auth (JWT passed to the DB; policies filter reads/writes). Example memory table: agent_memory (id, team_id, owner_id, visibility ['private','team','public'], content, embedding vector(1536), created_at). RLS policies: private → owner_id = auth.uid(); team → team_id in (select team_id from team_members where user_id = auth.uid()). Even a raw SQL query can't bypass the policies. Vector search: wrap pgvector similarity search in a function with security invoker so RLS applies automatically — the function only returns rows the user is allowed to see.

Designing shared vs private memory layers: (1) Private memory (visibility='private', owner_id) — personal preferences, draft content, sensitive conversation history, individual performance notes. (2) Team memory (visibility='team', scoped team_id) — shared knowledge bases (FAQs, policies), aggregated learnings, resolved customer issues, SOPs. (3) Workspace/global memory (visibility='public' or global flag) — brand guidelines, approved messaging, compliance rules; write access tightly restricted (admin-only). Separate read/write permissions: all members read team memory; only designated agents/admins write; agent writes to team memory require a review or confidence threshold. Implement via separate RLS policies for select/insert/update.

Permission-mirrored GitHub repos: the agent's knowledge base lives in a Git repo; access mirrors GitHub's permission model. Setup: store memory as files (Markdown/JSON/YAML) in a structured repo; use GitHub Teams to define access; sync permissions to the database (GitHub API or webhook → team_members table in Supabase) to keep RLS current; index only documents users have access to. Especially useful for engineering/product teams with code-adjacent knowledge (ADRs, runbooks, incident postmortems, API docs). Avoid stale permissions: sync via webhook (real-time) rather than scheduled jobs; short TTL on cached permission lookups; periodic reconciliation job auditing against current permissions.

Handling writes safely — the principle of conservative defaults: default visibility to private; only escalate when there's a clear reason. Escalation path: agent generates memory → default private → if content matches team-sharing criteria (resolved FAQ, verified fact, standard process) flag for promotion → promotion requires human approval or automated confidence check. Classify memory at write time with a lightweight classifier (an AI step — small fast model): does it contain personal/sensitive info (names, financial, health)? Is it user-specific or generalizable? What visibility level is appropriate?

Common mistakes: filtering only at the application layer (one bug away from a leak; treat app-layer filtering as a performance optimization, not a security control); treating all memory as equal (PII, compensation, health info need different handling than product docs); ignoring embedding leakage (RLS the vector search so similarity only considers authorized rows; avoid returning raw similarity scores that could be used to infer private data); over-sharing by default (shared memory is a liability surface).

Auditing and monitoring: log every read/write (timestamp, user ID, record ID, operation, visibility level) in a separate append-only table agents can't write to. Review for high-volume reads by a single user (scraping), out-of-scope reads (misconfiguration), unusual write patterns. Regulated industries may require this for compliance.

## Key points

- Team AI memory privacy is a data-layer problem, not an application-layer problem — enforce access with row-level security.
- Three memory tiers: private (owner), team (scoped to a group), global (organization-wide); default to private.
- Supabase RLS + pgvector gives semantic search that respects access controls (security invoker function — similarity only considers authorized rows).
- Permission-mirrored GitHub repos reuse GitHub's existing permission model for agent knowledge; sync memberships via webhooks to keep RLS current.
- Writes matter as much as reads: conservative defaults + write-time classification + review/confidence threshold for promotion.
- Avoid embedding leakage: RLS the similarity search; don't return raw similarity scores.
- Audit logs (append-only) for every memory read/write; watch for scraping and misconfiguration patterns.

## Technical data / figures

| Memory form | Storage/retrieval |
|---|---|
| Conversation history | Message log in DB, per-session |
| Semantic memory | Embeddings in vector DB, similarity search |
| Structured knowledge | Facts in relational/key-value stores |
| Episodic memory | Records of past events/decisions |

| Memory tier | Visibility | Use cases |
|---|---|---|
| Private | owner_id only | Preferences, drafts, sensitive history, performance notes |
| Team | team_id scope | Shared KBs, aggregated learnings, resolved issues, SOPs |
| Global | public/global flag | Brand guidelines, compliance rules (restricted writes) |

| Table field | Example |
|---|---|
| agent_memory schema | id, team_id, owner_id, visibility ('private'/'team'/'public'), content, embedding vector(1536), created_at |
| RLS private policy | visibility='private' and owner_id = auth.uid() |
| RLS team policy | visibility='team' and team_id in (team_members for user) |
| Vector search | security invoker function ordering by embedding <=> query_embedding |

| Write-safety step | Detail |
|---|---|
| Default visibility | private |
| Promotion criteria | resolved FAQ, verified fact, standard process |
| Promotion gate | human approval or automated confidence check |
| Write-time classification | small fast model outputs structured visibility decision |

## Why this source matters for the RAG

Essential reference for enterprise RAG/agent-memory deployments that must enforce per-user access: RLS-backed retrieval (including vector search) prevents unauthorized rows from ever reaching the model, and the private/team/global tiering plus write-time classification addresses data-privacy requirements. Provides concrete PostgreSQL/Supabase schema and policy patterns reusable for any shared retrieval system.

## Related context from the article

- Application-layer filtering is a performance optimization, not a security control.
- Embedding leakage: observing which records are retrieved for a probe query can reveal content.
- GitHub permission mirroring + webhook-synced team membership keeps RLS policies current.
- MindStudio connects agents to Supabase visually; RLS still enforces at the DB level.
