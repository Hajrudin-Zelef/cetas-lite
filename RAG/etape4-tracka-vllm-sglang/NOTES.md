# NOTES — corpus `etape4-tracka-vllm-sglang`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**. La source contient
6 titres H1 : un entête de document (`00-front-matter`) puis trois rapports et leurs
séparateurs. Les séparateurs de type `# PART n — …` (titre seul, sans texte) sont
**rattachés au rapport suivant** : dossiers `01-part-1-vllm`, `02-part-2-sglang`,
`03-part-3-synthesis-vllm-vs-sglang-2026-state`.

Source sans ancre HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

« Step 4 — Track A : frameworks de serving d'inférence, vLLM + SGLang »
(1er février → 22 septembre 2026).

- vLLM : versions 2026 (dont v0.30.0), architecture, quantification, benchmarks
  débit/latence, déploiement, écosystème, support entreprise, incertitudes ;
- SGLang (LMSYS) : versions, architecture, formats/quantification, déploiement,
  positionnement vs vLLM, questions ouvertes ;
- synthèse vLLM vs SGLang (état 2026).

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
