---
id: vague2-nerdykings/nerdykings/deepseek-thinking-visual-primitive
title: "DeepSeek 2026 : La Révolution Des Agents IA Multimodaux"
domain: nerdykings
role: reference
task: article
actors: ["DeepSeek"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "deepseek", "agentic", "attention", "memory", "multimodal", "reasoning"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/deepseek-thinking-visual-primitive.md
source_anchor: ""
source_lines: [1, 42]
sha256: 8d62107c74ba85e3ba7d312951edd1f3d2ab446000944359f7f6e201d12001da
---

# DeepSeek 2026 : La Révolution Des Agents IA Multimodaux

## Metadata

- **Source** : https://www.nerdykings.com/blog/deepseek-thinking-visual-primitive.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article covers DeepSeek's approach to making multimodal AI better at spatial reasoning. Current multimodal models are impressive at describing what they see, but collapse on advanced spatial reasoning — counting objects, following trajectories, identifying the correct zone of an interface. Today's models reason with phrases like "the object at the top left" or "the blue button under the menu," which becomes ambiguous in dense images (10 buttons, 3 menus). When counting 20 similar objects, text descriptions easily lead to omissions or double-counting. DeepSeek calls this a reference problem: the model sees elements but has no reliable way to maintain a precise reference — like counting all students in a class photo without using your finger.

The solution is "thinking with visual primitives." Instead of forcing the model to describe everything in words, DeepSeek lets it use visual markers directly in its reasoning: bounding boxes (frames around an object, letting the model literally say "I mean that one") and coordinate points (precise points to follow a path, connection, or trajectory — useful for mazes, diagrams, or locating the exact click zone). Importantly, these coordinates are not appended as a final annotation but integrated directly into the reasoning chain: the model alternates between text and visual markers, explaining, pointing, continuing, framing, then reaching its answer.

This benefits AI agents, which must act on images (click the right place, fill the right field, open the right menu). A small reference error can break the whole task. With visual primitives, the model can directly link its reasoning to the image and keep track of what it is looking at. It also helps debugging: when a visual agent fails, you can inspect its visual logic via the boxes and points generated during reasoning. On architecture, DeepSeek worked on efficiency: the image is split into patches, then aggressively compressed (neighboring patches grouped into a single visual token), combined with a mechanism called Compressed Sparse Attention (CSA) that greatly lightens the model's internal memory. Limitations: aggressive compression loses fine details (tiny text, micro-objects, thin lines); the model doesn't always decide alone when to use primitives (prompt guidance sometimes needed); and the approach suits organized 2D structures (interfaces, documents), while complex 3D scenes remain largely unsolved.

## Key points

- Multimodal models describe well but fail at spatial reasoning (counting, trajectories, UI zones).
- The "reference problem": models see elements but can't reliably maintain precise references.
- Solution: "thinking with visual primitives" — bounding boxes and coordinate points.
- Primitives are integrated into the reasoning chain, not appended as annotations.
- Enables more reliable agentic UI interaction and visual debugging.
- Efficiency: patch compression into single visual tokens + Compressed Sparse Attention (CSA).
- Limitations: loss of fine detail, need for prompt guidance, 3D scenes unsolved.

## Technical data / figures

| Item | Detail |
|---|---|
| Primitive types | Bounding boxes, coordinate points |
| Integration | Directly in reasoning chain (not post-annotation) |
| Efficiency mechanism | Patch compression + Compressed Sparse Attention (CSA) |
| Best-fit domains | Organized 2D structures (UI, documents) |
| Weak domains | 3D complex scenes, fine text/micro-objects |

## Why this source matters for the RAG

This source explains a concrete technique for improving visual reasoning and multimodal agent reliability, with clear architecture details and limitations. It is valuable for RAG corpora on multimodal AI, GUI agents, and computer-use agents.
