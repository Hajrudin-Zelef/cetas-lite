---
id: collect-240926-misc/misc/mimo-2-6-vs-mimo-v2-6-meme-version-avec-en-plus-un-nom-piege-2
title: "mimo-2-6-vs-mimo-v2-6-meme-version-avec-en-plus-un-nom-piege"
domain: orcarouter
role: reference
task: reference
actors: ["Xiaomi"]
dates: []
keywords: ["benchmark", "gpu", "pricing", "training"]
source: docs/RAG/clean_en/misc/mimo-2-6-vs-mimo-v2-6-meme-version-avec-en-plus-un-nom-piege.md
source_anchor: ""
source_lines: [53, 71]
sha256: 5a478e10fd32cb39e68f61bb8a615da79d5d44ce27f3b7aee757b89bd995ab30
---

# mimo-2-6-vs-mimo-v2-6-meme-version-avec-en-plus-un-nom-piege

• What does `config.json` say about the drafter? The Flash card describes a five-layer speculative drafter predicting seven tokens per pass; the config sets `num_nextn_predict_layers` to 3. The runtime reads the config.

• Which identifier are you actually sending? Xiaomi's platform lists explicit lowercase identifiers — `mimo-v2.6-pro`, `mimo-v2.6-flash`, `mimo-v2.6-pro-ultraspeed`. There is no plain `mimo-v2.6` among them. Xiaomi's own documentation also notes that users migrated from the invite-only test need to change the model name to continue, which is the same lesson delivered by the vendor.

## Where a router has a place, and where it doesn't

The honest position on this release is that the MiMo-V2.6 checkpoints are a self-hosting decision first, before they are an API decision. They are MIT-licensed downloads first; the choice between the 309B checkpoint and the 1.02T one is a choice between two GPU bills more than between two measuring instruments. We do not route any Xiaomi model, and nothing here should be read as suggesting otherwise — there is no MiMo-V2.6-Pro route on our platform, and the availability mention in Xiaomi's own materials points to Xiaomi's own channels.

What a router is actually for in this situation is the other half of the decision. If you are self-hosting a 1.02-trillion-parameter checkpoint to serve a workload that only sometimes needs frontier capability, the useful architecture is a cheap local path for the bulk of traffic and a hosted path for the tail — and it is on the hosted path that **a single API key for 200+ models, at provider list price passed through with 0% margin** does something concrete, because the frontier models you would benchmark MiMo-V2.6-Pro against change their pricing without warning, and a pass-through means the change reaches you the same day. **Automatic failover** is the other half: an unproven or freshly released model is exactly what you want on a path that can switch rather than on a production path of its own.

This is a statement about how to place this release in a stack, not a claim about its hosting.

## What to watch

Three things would change how this release reads.

An independent run of DeepSWE or Terminal Bench on MiMo-V2.6-Flash would do more for the family than any new announcement, because Flash is the member with no third-party score and the one most teams would actually run. An explanation of the parameter-count and drafter discrepancies would remove the last reason to distrust the weight data, that is, the number that determines the hardware. And an independent attempt to reproduce the RL pipeline from the published environments is the real test of the claim Xiaomi is actually making — that the interesting artifact here is not the checkpoint but the training loop that produced it.

Until then, the practical summary of MiMo-2.6 vs MiMo-V2.6 is brief. One version, two spellings, two checkpoints, one of which bears the name of the flagship model in common usage. Type either spelling into a search bar and you will find this generation. Type either into an API call and nothing will respond, because neither is an identifier — and that, rather than the missing V, is the distinction that matters.
