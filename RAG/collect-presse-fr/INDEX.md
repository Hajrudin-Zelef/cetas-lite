# INDEX — Presse FR — IA & tech (2026)

Corpus `collect-presse-fr` · **9 fiches** · 430 lignes · ~6727 mots · **une fiche = un chunk**, copiée verbatim de `docs/RAG/Collect RAG/05_presse_fr`.

## Mode d'emploi

1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur `domain`, `task`, `actors`, `dates` ou `keywords`.
2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité thématique auto-suffisante avec un en-tête YAML.
3. Pour un événement répété dans plusieurs sections, préférer le fichier marqué `canonical_for` (voir la table Événements canoniques).

## Domaines (dossiers → fichiers)

### `presse-fr/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Adoption de l'IA : la France s'installe à la 4ᵉ place du classement mondial](presse-fr/adoption-ia-france-4e-place.md) | 1–53 | reference | article |
| 02 | [Claude Opus 5.5 face à GPT-6 Sol et Luna : trois nouvelles IA d'un coup, et des prix qui chutent](presse-fr/claude-opus-5-5-gpt-6-sol-luna-lesnumeriques.md) | 1–51 | reference | article |
| 03 | [Claude Opus 5.5 et GPT-6 Sol sortent le même jour : Anthropic et OpenAI ont du mal à « ralentir » l'IA](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md) | 1–64 | reference | article |
| 04 | [Aux États-Unis, les data centers se négocient dans le secret](presse-fr/etats-unis-data-centers-secret-1.md) | 1–54 | reference | article |
| 05 | [Aux États-Unis, les data centers se négocient dans le secret](presse-fr/etats-unis-data-centers-secret-2.md) | 55–55 | reference | article |
| 06 | [GPT-6 Astra : le modèle le plus intelligent et le mieux aligné au monde](presse-fr/gpt-6-astra-koul.md) | 1–70 | reference | article |
| 07 | [Meta lance Muse Spark 1.3, axé sur le code et le contexte long](presse-fr/meta-muse-spark-1-3.md) | 1–50 | reference | article |
| 08 | [OpenAI gagne du terrain avec GPT-6 Astra, Anthropic prépare sa riposte : un nouveau Claude se profile](presse-fr/openai-gpt-6-astra-anthropic-riposte-zdnet.md) | 1–33 | reference | article |
| 09 | [OpenAI lance GPT-6 Sol et Luna, avec des tarifs API divisés par deux](presse-fr/openai-gpt-6-sol-luna-tarifs.md) | 1–54 | reference | article |

## Par tâche

- **article** — [Adoption de l'IA : la France s'installe à la 4ᵉ place du classement mondial](presse-fr/adoption-ia-france-4e-place.md), [Claude Opus 5.5 face à GPT-6 Sol et Luna : trois nouvelles IA d'un coup, et des prix qui chutent](presse-fr/claude-opus-5-5-gpt-6-sol-luna-lesnumeriques.md), [Claude Opus 5.5 et GPT-6 Sol sortent le même jour : Anthropic et OpenAI ont du mal à « ralentir » l'IA](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md), [Aux États-Unis, les data centers se négocient dans le secret](presse-fr/etats-unis-data-centers-secret-1.md), [Aux États-Unis, les data centers se négocient dans le secret](presse-fr/etats-unis-data-centers-secret-2.md), [GPT-6 Astra : le modèle le plus intelligent et le mieux aligné au monde](presse-fr/gpt-6-astra-koul.md), [Meta lance Muse Spark 1.3, axé sur le code et le contexte long](presse-fr/meta-muse-spark-1-3.md), [OpenAI gagne du terrain avec GPT-6 Astra, Anthropic prépare sa riposte : un nouveau Claude se profile](presse-fr/openai-gpt-6-astra-anthropic-riposte-zdnet.md), [OpenAI lance GPT-6 Sol et Luna, avec des tarifs API divisés par deux](presse-fr/openai-gpt-6-sol-luna-tarifs.md)

## Par acteur

- **AWS** (2) — [presse-fr/etats-unis-data-centers-secret-1.md](presse-fr/etats-unis-data-centers-secret-1.md), [presse-fr/gpt-6-astra-koul.md](presse-fr/gpt-6-astra-koul.md)
- **Alibaba** (1) — [presse-fr/adoption-ia-france-4e-place.md](presse-fr/adoption-ia-france-4e-place.md)
- **Anthropic** (6) — [presse-fr/claude-opus-5-5-gpt-6-sol-luna-lesnumeriques.md](presse-fr/claude-opus-5-5-gpt-6-sol-luna-lesnumeriques.md), [presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md), [presse-fr/gpt-6-astra-koul.md](presse-fr/gpt-6-astra-koul.md), [presse-fr/meta-muse-spark-1-3.md](presse-fr/meta-muse-spark-1-3.md), [presse-fr/openai-gpt-6-astra-anthropic-riposte-zdnet.md](presse-fr/openai-gpt-6-astra-anthropic-riposte-zdnet.md), [presse-fr/openai-gpt-6-sol-luna-tarifs.md](presse-fr/openai-gpt-6-sol-luna-tarifs.md)
- **Applied Digital** (1) — [presse-fr/etats-unis-data-centers-secret-1.md](presse-fr/etats-unis-data-centers-secret-1.md)
- **DeepSeek** (1) — [presse-fr/adoption-ia-france-4e-place.md](presse-fr/adoption-ia-france-4e-place.md)
- **EU** (1) — [presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md)
- **Google** (2) — [presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md), [presse-fr/meta-muse-spark-1-3.md](presse-fr/meta-muse-spark-1-3.md)
- **Hugging Face** (1) — [presse-fr/gpt-6-astra-koul.md](presse-fr/gpt-6-astra-koul.md)
- **Meta** (2) — [presse-fr/etats-unis-data-centers-secret-1.md](presse-fr/etats-unis-data-centers-secret-1.md), [presse-fr/meta-muse-spark-1-3.md](presse-fr/meta-muse-spark-1-3.md)
- **Microsoft** (2) — [presse-fr/adoption-ia-france-4e-place.md](presse-fr/adoption-ia-france-4e-place.md), [presse-fr/gpt-6-astra-koul.md](presse-fr/gpt-6-astra-koul.md)
- **Moonshot** (1) — [presse-fr/adoption-ia-france-4e-place.md](presse-fr/adoption-ia-france-4e-place.md)
- **OpenAI** (6) — [presse-fr/claude-opus-5-5-gpt-6-sol-luna-lesnumeriques.md](presse-fr/claude-opus-5-5-gpt-6-sol-luna-lesnumeriques.md), [presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md), [presse-fr/gpt-6-astra-koul.md](presse-fr/gpt-6-astra-koul.md), [presse-fr/meta-muse-spark-1-3.md](presse-fr/meta-muse-spark-1-3.md), [presse-fr/openai-gpt-6-astra-anthropic-riposte-zdnet.md](presse-fr/openai-gpt-6-astra-anthropic-riposte-zdnet.md), [presse-fr/openai-gpt-6-sol-luna-tarifs.md](presse-fr/openai-gpt-6-sol-luna-tarifs.md)
- **OpenRouter** (1) — [presse-fr/adoption-ia-france-4e-place.md](presse-fr/adoption-ia-france-4e-place.md)
- **SpaceX** (1) — [presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md)
- **United States** (2) — [presse-fr/etats-unis-data-centers-secret-1.md](presse-fr/etats-unis-data-centers-secret-1.md), [presse-fr/etats-unis-data-centers-secret-2.md](presse-fr/etats-unis-data-centers-secret-2.md)
- **xAI** (1) — [presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md)

## Par date

- **2023-03** — [presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md)
- **2023-12** — [presse-fr/etats-unis-data-centers-secret-1.md](presse-fr/etats-unis-data-centers-secret-1.md)
- **2025-03** — [presse-fr/etats-unis-data-centers-secret-1.md](presse-fr/etats-unis-data-centers-secret-1.md)
- **2025-12** — [presse-fr/etats-unis-data-centers-secret-1.md](presse-fr/etats-unis-data-centers-secret-1.md)
- **2026-08** — [presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md), [presse-fr/meta-muse-spark-1-3.md](presse-fr/meta-muse-spark-1-3.md)
- **2026-09** — [presse-fr/adoption-ia-france-4e-place.md](presse-fr/adoption-ia-france-4e-place.md), [presse-fr/claude-opus-5-5-gpt-6-sol-luna-lesnumeriques.md](presse-fr/claude-opus-5-5-gpt-6-sol-luna-lesnumeriques.md), [presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md), [presse-fr/etats-unis-data-centers-secret-1.md](presse-fr/etats-unis-data-centers-secret-1.md), [presse-fr/gpt-6-astra-koul.md](presse-fr/gpt-6-astra-koul.md), [presse-fr/meta-muse-spark-1-3.md](presse-fr/meta-muse-spark-1-3.md), [presse-fr/openai-gpt-6-sol-luna-tarifs.md](presse-fr/openai-gpt-6-sol-luna-tarifs.md)
- **2026-09-23** — [presse-fr/adoption-ia-france-4e-place.md](presse-fr/adoption-ia-france-4e-place.md), [presse-fr/claude-opus-5-5-gpt-6-sol-luna-lesnumeriques.md](presse-fr/claude-opus-5-5-gpt-6-sol-luna-lesnumeriques.md), [presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md](presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md), [presse-fr/etats-unis-data-centers-secret-1.md](presse-fr/etats-unis-data-centers-secret-1.md), [presse-fr/gpt-6-astra-koul.md](presse-fr/gpt-6-astra-koul.md), [presse-fr/meta-muse-spark-1-3.md](presse-fr/meta-muse-spark-1-3.md), [presse-fr/openai-gpt-6-astra-anthropic-riposte-zdnet.md](presse-fr/openai-gpt-6-astra-anthropic-riposte-zdnet.md), [presse-fr/openai-gpt-6-sol-luna-tarifs.md](presse-fr/openai-gpt-6-sol-luna-tarifs.md)

## Carte de couverture (lignes source)

| plage | fichier |
|---|---|
| 1–53 | collect-presse-fr/presse-fr/adoption-ia-france-4e-place.md |
| 1–51 | collect-presse-fr/presse-fr/claude-opus-5-5-gpt-6-sol-luna-lesnumeriques.md |
| 1–64 | collect-presse-fr/presse-fr/claude-opus-5-5-gpt-6-sol-numerama.md |
| 1–54 | collect-presse-fr/presse-fr/etats-unis-data-centers-secret-1.md |
| 55–55 | collect-presse-fr/presse-fr/etats-unis-data-centers-secret-2.md |
| 1–70 | collect-presse-fr/presse-fr/gpt-6-astra-koul.md |
| 1–50 | collect-presse-fr/presse-fr/meta-muse-spark-1-3.md |
| 1–33 | collect-presse-fr/presse-fr/openai-gpt-6-astra-anthropic-riposte-zdnet.md |
| 1–54 | collect-presse-fr/presse-fr/openai-gpt-6-sol-luna-tarifs.md |

