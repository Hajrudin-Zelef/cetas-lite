# INDEX — RAG

Corpus RAG de référence pour Cetas. Chaque corpus est une partition exacte de sa source, avec index et manifest.

Un corpus `delta` ne contient que les faits nouveaux/corrigés d'une source déjà couverte par un corpus `base` ; il ne la remplace pas. `delta_of` indique la base visée.

| corpus | titre | relation | fichiers | source | index |
|---|---|---|---|---|---|
| `briefing-ia-2026` | AI News 2026 — Reference Dossier |  | 126 | `docs/RAG/briefing-ia-2026-en.md` | [INDEX](briefing-ia-2026/INDEX.md) |
| `briefing-general-tech-2026` | General Tech News 2026 — Hardware, Infrastructure & Consumer Tech |  | 112 | `docs/RAG/briefing-general-tech-2026-en.md` | [INDEX](briefing-general-tech-2026/INDEX.md) |
| `ai-industry-kb-2026` | AI Industry Knowledge Base 2026 | base | 123 | `docs/RAG/ai-industry-knowledge-base-2026.md` | [INDEX](ai-industry-kb-2026/INDEX.md) |
| `ai-industry-kb-2026-wave6` | AI Industry Knowledge Base 2026 — Wave 6 Consolidation | delta de `ai-industry-kb-2026` | 136 | `docs/RAG/ai-industry-knowledge-base-2026-wave6.md` | [INDEX](ai-industry-kb-2026-wave6/INDEX.md) |
| `frontier-models-2026` | Frontier AI Models 2026 — Vague 1 (EN) |  | 12 | `docs/RAG/Grands titres IA modèlesEN.md` | [INDEX](frontier-models-2026/INDEX.md) |
| `labs-grok-platforms-2026` | Labs, Grok, Tools & Platforms 2026 — Vague 2 (EN) |  | 17 | `docs/RAG/Labos, Grok, outils & plateformes_EN.md` | [INDEX](labs-grok-platforms-2026/INDEX.md) |
| `open-local-models-2026` | Open / Local AI Models 2026 (EN) |  | 26 | `docs/RAG/Modèles IA open  locauxEN.md` | [INDEX](open-local-models-2026/INDEX.md) |
| `tools-platforms-2026` | AI Tools & Platforms 2026 (Step 2) |  | 8 | `docs/RAG/Outils & plateformes IAEN.md` | [INDEX](tools-platforms-2026/INDEX.md) |

Voir aussi [README](README.md) et `manifest.json`.
