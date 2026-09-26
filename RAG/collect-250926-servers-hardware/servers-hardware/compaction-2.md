---
id: collect-250926-servers-hardware/servers-hardware/compaction-2
title: "Compaction"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["pruning"]
source: docs/RAG/clean4/compaction.md
source_anchor: ""
source_lines: [149, 190]
sha256: 102e4f97ae17a7100174e974f6424671972310569d21945b51ec81622890e68a
---

# Compaction

```
recent user message
recent assistant message
tool output (limited to 2,000 characters)
attachment descriptor (embedded data omitted)
```
On later local compactions, OpenCode updates the previous summary, carries its retained recent context forward, and then selects a new tail. Completed checkpoints appear to the model as historical conversation, not new instructions. Running and failed checkpoints are excluded from model context.

## Instructions

Instruction updates and conversation compaction are tracked separately. Before each actual provider attempt, OpenCode checks live instruction sources before it delivers pending input. Later changes appear to the model as chronological system messages, while request construction renders the stored baseline.

`initial instructions â instruction update â completed checkpoint â new baseline`
When compaction completes, the instruction values already accepted at that exact point become the new baseline. Compaction does not reread sources or emit a new instruction update. Moving a session keeps this state, so instructions changed at the destination appear chronologically. A committed revert clears it, and the next model attempt performs one complete source read. See Instructions for source ordering and update behavior.

## Limits

| Limit | Result | 
|---|---|
| Model | Compaction requires a resolvable model. There is no separate compaction model or fallback model. | 
| Catalog | Automatic scheduling requires a positive catalog context limit; manual compaction and overflow recovery do not. | 
| Summary | Generation can fail when its prompt and output allowance do not fit, the model returns no summary, or the provider fails. | 
| History | Automatic and overflow compaction require older context that can be replaced. An overflow can remain when there is no compressible history or fixed instructions and tool schemas dominate the request. | 
| Recovery | Overflow recovery retries only once per model step. Heuristic token estimates cannot prevent every provider-specific overflow. | 
| Storage | Earlier messages remain stored even when excluded from active model context. | 

For example, compaction cannot create room when almost the entire request is a fixed system prompt and tool schemas:

`128k context = 120k fixed instructions and tools + 8k conversation`
## Migration

V1 also used tail-turn and pruning behavior. V2 instead uses checkpoint-based
compaction and `compaction.keep.tokens`:

```
{
  "compaction": {
    "keep": { "tokens": 15000 }
  }
}
```
The settings and behavior on this page apply to V2.
