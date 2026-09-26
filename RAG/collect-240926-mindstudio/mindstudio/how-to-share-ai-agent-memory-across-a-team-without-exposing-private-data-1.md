---
id: collect-240926-mindstudio/mindstudio/how-to-share-ai-agent-memory-across-a-team-without-exposing-private-data-1
title: "how-to-share-ai-agent-memory-across-a-team-without-exposing-private-data"
domain: mindstudio
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "memory", "agents", "embedding", "embeddings"]
source: docs/RAG/clean_en/mindstudio/how-to-share-ai-agent-memory-across-a-team-without-exposing-private-data.md
source_anchor: ""
source_lines: [1, 153]
sha256: 70546fcf741191e2692c3bc9f205089fcdcf0cca37376d7cb544ec65014bfc99
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

