---
id: collect-240926-mindstudio/mindstudio/how-to-share-ai-agent-memory-across-a-team-without-exposing-private-data
title: "how-to-share-ai-agent-memory-across-a-team-without-exposing-private-data"
domain: mindstudio
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "memory", "agents", "embedding", "embeddings", "incident", "liability"]
source: docs/RAG/clean_en/mindstudio/how-to-share-ai-agent-memory-across-a-team-without-exposing-private-data.md
source_anchor: ""
source_lines: [1, 327]
sha256: 3f66c8605aa2821b0c3b18285f53a4d8e2be6f491378e7ce5ea950c34992bf98
---

# how-to-share-ai-agent-memory-across-a-team-without-exposing-private-data

<!-- source: https://www.mindstudio.ai/blog/share-ai-agent-memory-across-team -->

## Why Team AI Memory Is a Privacy Problem Worth Solving

When you build AI agents for a single user, memory is simple: store what the agent learns, retrieve it when relevant. But the moment you scale that to a team, you hit a wall.

Sales reps shouldn’t see each other’s commission structures. HR agents shouldn’t surface one employee’s performance notes to another. A support agent that “remembers” a customer complaint shouldn’t let that context bleed into an unrelated interaction.

Sharing AI agent memory across a team without exposing private data is one of the harder design problems in enterprise AI. It’s not just a technical challenge — it’s a trust challenge. Get it wrong and you’ve built a system that leaks information in subtle, hard-to-audit ways.

This guide walks through a practical architecture for doing this right: how to separate shared from private memory, how to enforce permissions at the data layer using row-level security, and how approaches like permission-mirrored repositories let teams scale agent memory safely.

## What AI Agent Memory Actually Means

Before getting into the architecture, it helps to be precise about what “memory” means for AI agents.

AI models themselves are stateless — they don’t remember anything between conversations unless you explicitly pass that context back in. “Memory” is really just a retrieval system that fetches relevant information and includes it in the prompt.

There are a few common forms:

- **Conversation history** — A log of prior messages, often stored in a database and retrieved per-session.
- **Semantic memory** — Embeddings stored in a vector database, retrieved by similarity search. Good for surfacing relevant facts, past interactions, or documents.
- **Structured knowledge** — Facts stored in relational or key-value stores (e.g., “this user prefers email over Slack”).
- **Episodic memory** — Records of past events or decisions, useful for agents that need to reason about what they’ve done before.

## Other agents ship a demo. Remy ships an app.

Real backend. Real database. Real auth. Real plumbing. Remy has it all.

Each type creates different privacy risks when shared across teams.

## The Core Problem: Shared Infrastructure, Unequal Access

Most teams build memory on shared infrastructure — a single Supabase project, a single Pinecone index, or a shared PostgreSQL database. That’s efficient. The problem is access control.

A naive implementation looks like this: every agent reads from and writes to the same memory store, tagged by user or session ID. But if that filter is applied at the application layer — inside your agent logic — rather than the database layer, it’s fragile.

An agent prompt that says “only retrieve records where user_id = current_user” is only as trustworthy as the code passing that variable. A misconfigured agent, a prompt injection attack, or a bug in your retrieval logic could cause it to return records it shouldn’t.

The fix is to push access control down to the database layer, so that unauthorized reads are structurally impossible — not just discouraged by application code.

This is where row-level security becomes essential.

## Row-Level Security: The Right Place to Enforce Privacy

Row-level security (RLS) is a database feature that filters which rows a user can see based on their identity. It’s enforced by the database engine itself, not by application logic.

Supabase makes RLS especially approachable because it’s built on top of PostgreSQL’s native RLS and integrates with its auth system. When a user authenticates, their JWT is passed to the database, and policies automatically filter what they can read or write.

### Setting Up a Memory Table with RLS in Supabase

Here’s a simplified example of a memory table for AI agents:

```
create table agent_memory (
  id uuid primary key default gen_random_uuid(),
  team_id uuid not null,
  owner_id uuid references auth.users(id),
  visibility text check (visibility in ('private', 'team', 'public')),
  content text,
  embedding vector(1536),
  created_at timestamptz default now()
);
```
The `visibility` field is key. A record marked `private` should only be readable by its `owner_id`. A record marked `team` should be readable by any user in the same `team_id`. Public records are available to all.

Enforce this with RLS policies:

```
-- Enable RLS
alter table agent_memory enable row level security;
-- Private: only the owner can see it
create policy "Private records: owner only"
  on agent_memory for select
  using (
    visibility = 'private' and owner_id = auth.uid()
  );
-- Team: any member of the same team can see it
create policy "Team records: same team"
  on agent_memory for select
  using (
    visibility = 'team' and team_id in (
      select team_id from team_members where user_id = auth.uid()
    )
  );
```
Now even if an agent retrieves records using a raw SQL query, the database enforces the policy. The agent cannot see what it’s not supposed to see — regardless of how the query is written.

### Applying RLS to Vector Search

If you’re doing semantic search over embeddings, RLS applies there too. Supabase supports `pgvector` for similarity search, and you can wrap your vector search in a function that respects RLS automatically.

```
create function search_agent_memory(query_embedding vector(1536), match_count int)
returns setof agent_memory
language sql security invoker
as $$
  select *
  from agent_memory
  order by embedding <=> query_embedding
  limit match_count;
$$;
```
## Remy is new. The platform isn't.

Remy is the latest expression of years of platform work. Not a hastily wrapped LLM.

Because this uses `security invoker`, the function runs with the calling user’s permissions. RLS policies apply automatically — the function will only return rows the user is allowed to see.

This is the right pattern. The privacy guarantee lives at the data layer, not the application layer.

## Designing Shared vs. Private Memory Layers

Once you have RLS enforcing access at the database level, you can design a deliberate memory architecture with distinct layers.

### Layer 1: Private Memory

Private memory belongs to an individual user or agent session. No one else on the team can read it.

Use cases:

- A user’s personal preferences or communication style
- Draft content the user hasn’t shared yet
- Sensitive conversation history
- Individual performance notes

In your memory table, these rows have `visibility = 'private'` and `owner_id` set to the user.

### Layer 2: Team Memory

Team memory is shared across a defined group. Everyone in the team can read it, and (depending on your write policies) certain roles may be able to write to it.

Use cases:

- Shared knowledge bases (product FAQs, company policies)
- Aggregated learnings from multiple agents
- Resolved customer issues that are safe to share
- Standard operating procedures that agents reference

These rows have `visibility = 'team'` and are scoped to a `team_id`.

### Layer 3: Workspace or Global Memory

Some memory should be available across the entire organization — brand guidelines, approved messaging, compliance rules. This is the equivalent of a company knowledge base.

These rows have `visibility = 'public'` or a separate global flag. Write access to global memory should be tightly restricted (admin-only or through a controlled process).

### Separating Read and Write Permissions

Memory layers need different policies for reading vs. writing. A common mistake is making team memory writable by all team members, which leads to agents polluting shared memory with low-quality or incorrect data.

A safer pattern:

- All team members can **read** team memory.
- Only designated agents or admins can **write** to team memory.
- Writing to team memory from an agent requires a review or confidence threshold before the data is committed.

You can implement this with separate RLS policies for `select`, `insert`, and `update` operations.

## Permission-Mirrored GitHub Repos for Agent Knowledge

For teams that version-control their agent configurations, prompts, or knowledge bases, GitHub repos offer another layer of memory management — and GitHub’s existing permission system can be mirrored into your agent architecture.

The idea: your agent’s knowledge base lives in a Git repository. Access to that knowledge base mirrors the access controls in GitHub.

### Why This Works

GitHub already has a mature permission model: organizations, teams, and repository-level access controls. If your company already manages document access via GitHub (common for engineering teams), you can reuse that structure for agent memory.

Here’s how to set it up:

1. **Store memory as files in a structured repo** — Markdown files, JSON documents, or structured YAML configs that the agent can retrieve.
2. **Use GitHub Teams to define access** — Org members, teams, and roles already exist. Mirror these into your memory retrieval logic.
3. **Sync permissions to your database** — Use GitHub’s API or a webhook to sync team memberships into a`team_members` table in Supabase. This keeps your RLS policies current without manual maintenance.
4. **Index only what users have access to** — When building your vector index for semantic search, only embed documents that the requesting user has read access to in GitHub.

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

This approach is especially useful for engineering or product teams where the knowledge base is already code-adjacent — architecture decision records, runbooks, incident postmortems, API documentation.

### Avoiding Stale Permissions

The main risk with this pattern is stale data. If someone’s access is revoked in GitHub but the memory store isn’t updated immediately, they might still see records they shouldn’t.

Mitigate this by:

- Syncing GitHub team memberships via webhook (real-time) rather than scheduled jobs.
- Setting a short TTL on cached permission lookups.
- Running a periodic reconciliation job that audits the memory store against current GitHub permissions and flags or removes orphaned records.

## Handling Memory Writes Safely

Reading is only half the problem. Writes are where most teams get into trouble.

When an AI agent writes to memory — storing a new fact, updating a record, saving a conversation summary — it needs to make the right call about visibility. If your agent defaults everything to `team` visibility, private information can leak into shared contexts.

### The Principle of Conservative Defaults

Set the default visibility to `private`. Agents should only escalate visibility when there’s a clear reason.

A sensible escalation path:

1. Agent generates a memory item.
2. Default: mark as `private` .
3. If the content matches criteria for team sharing (e.g., it’s a resolved FAQ, a verified fact, a standard process), flag it for promotion.
4. Promotion to `team` or`global` requires either human approval or an automated confidence check.

This prevents agents from silently leaking sensitive context into shared spaces.

### Classifying Memory at Write Time

You can use a lightweight classification step before writing to memory. Pass the content through a classifier that determines:

- Does this contain personal or sensitive information (names, financial data, health info)?
- Is this specific to one user or generalizable to the team?
- What visibility level is appropriate?

This classification can itself be an AI step — a small, fast model that outputs a structured decision before the write is committed.

## Building This in MindStudio

MindStudio’s no-code agent builder is a practical place to implement this kind of memory architecture without writing a full backend from scratch.

You can connect MindStudio agents directly to Supabase using its built-in integrations, then build the read/write logic visually. An agent workflow might look like:

1. **Receive user input** — the agent gets a query or task.
2. **Retrieve relevant memory** — call the Supabase function (with RLS in place) to fetch semantically similar records the current user is allowed to see.
3. **Generate a response** — pass retrieved context plus the user input to your chosen model.
4. **Write to memory** — run a classification step, determine visibility, write the record.

The RLS enforcement happens at the Supabase level, so MindStudio agents can’t accidentally bypass it — the database simply won’t return unauthorized rows.

For teams that want to automate memory promotion (moving private memories to team-visible), you can build a separate MindStudio workflow that runs on a schedule: it scans private memories flagged for review, applies your classification logic, and either promotes or discards them. No code required for the orchestration layer.

### Built like a system. Not vibe-coded.

Remy manages the project — every layer architected, not stitched together at the last second.

MindStudio also supports custom JavaScript functions if you need more fine-grained control — for example, a function that calls the GitHub API to verify a user’s team membership before allowing a memory read.

You can start building for free at mindstudio.ai.

## Common Mistakes to Avoid

Even teams with good intentions get this wrong. Here are the patterns that cause the most problems.

### Filtering at the Application Layer Only

If your only privacy control is a `where user_id = ?` clause in your agent code, you’re one bug away from a data leak. Push enforcement to the database using RLS. Treat application-layer filtering as a performance optimization, not a security control.

### Treating All Memory as Equal

Not all memory needs the same protection. Customer PII, compensation data, and health information need different handling than product documentation. Classify your memory types and apply appropriate controls to each.

### Ignoring Embedding Leakage

Even if you restrict access to the raw text of a memory record, its embedding can reveal information. If a bad actor can query your vector store and observe which records are retrieved as similar to a probe query, they can infer content they’re not supposed to see.

Mitigate this by:

- Applying RLS to the embedding search results (as shown above), so the similarity function only considers rows the user is allowed to see.
- Avoiding returning raw similarity scores to users in contexts where they could be used to infer private data.

### Over-Sharing by Default

The convenience of shared memory makes it tempting to share everything. But shared memory is a liability surface. Only promote information to team or global visibility when there’s a clear, auditable reason.

## Auditing and Monitoring Memory Access

Privacy controls are only as good as your ability to audit them.

At minimum, log every memory read and write with:

- Timestamp
- User ID
- Record ID
- Operation (read, write, update, delete)
- Visibility level of the record

Store these logs in a separate, append-only table — don’t let agents write to the audit log.

Review your audit logs regularly for:

- High-volume reads by a single user (potential data scraping)
- Reads of records outside a user’s normal scope (possible misconfiguration)
- Unusual write patterns that might indicate an agent is incorrectly classifying memory

For regulated industries, these logs may also be required for compliance. PostgreSQL’s audit extension or a purpose-built audit table in Supabase both work well here.

## Frequently Asked Questions

### What is the difference between shared and private AI agent memory?

Private memory is only accessible to the individual user or session that created it. Shared (or team) memory is accessible to a defined group — typically a team or the entire organization. The key design goal is to make this separation enforceable at the data layer, not just the application layer, so there’s no risk of a misconfigured agent bypassing the controls.

### How does row-level security work for AI agent memory?

Row-level security (RLS) is a database-level feature — supported natively in PostgreSQL and exposed cleanly in Supabase — that filters which rows a user can see based on their identity. When an AI agent queries memory on behalf of a user, the database checks that user’s permissions and returns only the rows they’re authorized to see. The agent cannot retrieve data it’s not permitted to access, regardless of how the query is written.

### Can vector search be used safely with team memory?

Yes, but it requires deliberate design. Standard vector similarity search doesn’t natively enforce access controls — it searches the entire index. The safe approach is to wrap your vector search in a database function that runs with the caller’s security context (using `security invoker` in PostgreSQL), so RLS policies automatically filter results. This means the similarity search only considers records the user is allowed to see.

### What happens if a team member’s role or access changes?

If access controls are enforced at the database layer with RLS policies that check current team membership, a user’s memory access automatically reflects their current role. The key is keeping your team membership data current — ideally synced in real time via webhooks from your identity provider or GitHub. Stale team membership data is the main risk vector here.

### How do I prevent AI agents from writing private data to shared memory?

Use conservative defaults: agents should write to private memory unless there’s an explicit reason to share. Add a classification step before any write that evaluates whether the content contains sensitive information and assigns the appropriate visibility level. For promotion to team or global visibility, require either human approval or a confidence threshold that the content is safe to share.

### Is it possible to share AI agent memory across teams without using a vector database?

Yes. Structured relational databases work fine for many use cases, especially when memory consists of discrete facts or structured records rather than free-form text. Row-level security applies equally to relational tables. Vector databases (with embeddings for semantic search) are only necessary when you need fuzzy, similarity-based retrieval — for example, “find memories related to this question” rather than “find the record for this user ID.”

## Key Takeaways

- **Privacy in team AI memory is a data layer problem, not an application layer problem.** Enforce access controls with row-level security, not just filtering in agent code.
- **Design three memory tiers:** private (owner only), team (scoped to a group), and global (organization-wide). Default to private.
- **Supabase’s RLS with pgvector** gives you semantic search that respects access controls — the similarity function only considers rows the user is allowed to see.
- **Permission-mirrored GitHub repos** work well for engineering teams whose knowledge base already lives in version control. Sync GitHub team memberships to your database to keep permissions current.
- **Writes are as important as reads.** Use a classification step to assign visibility before committing a memory record, and keep audit logs for everything.
- **MindStudio** lets you build and orchestrate this kind of multi-step, permission-aware agent workflow without writing backend infrastructure from scratch.

If you’re building agents for teams and want to implement memory sharing without the risk of data leaks, MindStudio is worth exploring — it connects directly to Supabase and lets you build the classification and retrieval logic visually. The permission enforcement stays where it belongs: in the database.
