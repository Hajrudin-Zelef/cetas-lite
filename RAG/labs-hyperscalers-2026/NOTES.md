# NOTES — corpus `labs-hyperscalers-2026`

## Génération

Mode auto (`RAG/_tools/build_rag.py`) avec `first_is_content: true` (H1 unique = contenu),
cibles **45–90 lignes**. Dossier `00-labs-hyperscalers` (`folder_name`), fichiers par titre H2.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 3 — Labs & Hyperscalers » (1er février → 22 septembre 2026), rapport consolidé.
19 sections H2 :

§1 Anthropic · §2 OpenAI · §3 Google/DeepMind · §4 Meta · §5 Apple · §6 Microsoft ·
§7 Amazon · §8 xAI/Grok · §9 Mistral AI · §10 labs indépendants · §11 Nvidia NIM ·
§12 Groq · §13 hyperscalers d'inférence (Cerebras, SambaNova, Together AI, Fireworks AI,
Nebius, CoreWeave) · §14 désambiguïsation « Cabreras » + FreeLLMAPI ·
§15 chronologie maîtresse · §16 comparaisons inter-labs · §17 tracker M&A ·
§18 journal d'incertitudes · §19 métadonnées de collecte.

C'est le corpus le plus volumineux de la série (3 006 lignes, 199 sous-titres H3).

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
