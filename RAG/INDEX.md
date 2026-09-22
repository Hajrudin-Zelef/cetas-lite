# INDEX — RAG

Corpus RAG de référence pour Cetas. Chaque corpus est une partition exacte de sa source, avec index et manifest.

Un corpus `delta` ne contient que les faits nouveaux/corrigés d'une source déjà couverte par un corpus `base` ; il ne la remplace pas. `delta_of` indique la base visée.

| corpus | titre | relation | fichiers | source | index |
|---|---|---|---|---|---|
| `briefing-ia-2026` | AI News 2026 — Reference Dossier |  | 126 | `docs/RAG/briefing-ia-2026-en.md` | [INDEX](briefing-ia-2026/INDEX.md) |
| `briefing-general-tech-2026` | General Tech News 2026 — Hardware, Infrastructure & Consumer Tech |  | 112 | `docs/RAG/briefing-general-tech-2026-en.md` | [INDEX](briefing-general-tech-2026/INDEX.md) |
| `ai-industry-kb-2026` | AI Industry Knowledge Base 2026 | base | 123 | `docs/RAG/ai-industry-knowledge-base-2026.md` | [INDEX](ai-industry-kb-2026/INDEX.md) |
| `ai-industry-kb-2026-wave6` | AI Industry Knowledge Base 2026 — Wave 6 Consolidation | delta de `ai-industry-kb-2026` | 136 | `docs/RAG/ai-industry-knowledge-base-2026-wave6.md` | [INDEX](ai-industry-kb-2026-wave6/INDEX.md) |

Voir aussi [README](README.md) et `manifest.json`.
