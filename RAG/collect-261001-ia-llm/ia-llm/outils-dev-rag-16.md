---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-16
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: ["benchmark", "embedding", "embeddings", "valuation"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [2975, 3174]
sha256: 1254a4d964ffe4fe458bd525dccc894f67b29574b062c58a4ce7abe1ba94283d
---

# Outils dev + ingénierie RAG (chunk & corpus)

def simhash_dedup(docs: list[dict], distance_max: int = 3) -> list[str]:
    """Retourne les ids à SUPPRIMER (doublons proches d'un gardé)."""
    gardes, a_supprimer = [], []
    empreintes = []
    for d in docs:
        h = SimHash(d["text"])          # 64 bits
        doublon = any(h.distance(e) <= distance_max for e in empreintes)
        if doublon:
            a_supprimer.append(d["id"])
        else:
            gardes.append(d["id"]); empreintes.append(h)
    return a_supprimer
```

**Stratégie combinée recommandée :**
1. Dédup exacte (sha256) — gratuit, obligatoire.
2. SimHash (distance ≤ 3) — premier filtre near-dup, rapide.
3. MinHash+LSH (seuil 0.85) — arbitrage fin sur les cas restants.

## 100. Versioning du corpus (git / DVC / manifestes)

Le **code** est versionné par Git. Le **corpus** a besoin de son propre
versioning :

```jsonc
// corpus/manifest.json — VERSIONNÉ dans git (léger, précieux)
{
  "version": "v3",
  "date": "2026-09-27",
  "sources": [
    {"nom": "guides", "fichiers": 32, "chunks": 6400, "hash": "sha256:…"},
    {"nom": "huawei_cli", "fichiers": 12620, "chunks": 4100, "hash": "sha256:…"},
    {"nom": "articles", "fichiers": 312, "chunks": 2900, "hash": "sha256:…"}
  ],
  "chunker": "chunk_md.py @ commit a1b2c3d, MAX_TOKENS=500",
  "embedding_model": "text-embedding-3-small",
  "total_chunks": 13400,
  "notes": "v3 : re-chunk après fix fences §82"
}
```

- Chaque chunk porte `corpus_version` (§84) → tu sais **quels** vecteurs
  invalider quand le chunking change.
- **DVC** (Data Version Control, `pip install dvc`) : si ton corpus brut
  dépasse quelques Go, DVC versionne les *pointeurs* dans Git et les
  données sur ton disque/NAS/S3. À évaluer le jour où Git LFS coince (§16).
- **Règle** : on ne modifie jamais un corpus versionné en place — on crée
  `v4` et on migre (ou on re-ingère tout, c'est pas cher : §71).

## 101. Séparer les collections (guides vs CLI vs articles — ton cas réel)

**Ne mets pas tout dans le même sac** : une question « syntaxe exacte de
`display interface brief` » doit chercher dans les CLI, pas dans les
guides narratifs.

```sql
-- La colonne source (§64) suffit pour commencer :
--   'guides' | 'huawei_cli' | 'articles' | 'fiches'

-- Recherche restreinte à une collection :
SELECT * FROM match_chunks(:q, 0.70, 10, p_source := 'huawei_cli');

-- Recherche multi-collections avec boost (le SQL reste lisible) :
SELECT *, 1 - (embedding <=> :q) AS s FROM rag_chunks
WHERE source IN ('guides', 'huawei_cli')
ORDER BY embedding <=> :q LIMIT 10;
```

**Stratégie de seuils par collection** (à calibrer §103) :
- `huawei_cli` : seuil plus bas (0.60) — les commandes sont courtes, les
  similarités cosinus naturellement plus faibles.
- `guides` : seuil 0.70-0.75 — prose riche, similarités plus hautes.
- `articles` : 0.70.

## 102. Évaluation de la qualité : jeux de questions test

**Sans évaluation, tu pilotes à l'aveugle.** Construis un
`tests/golden.jsonl` (30-50 questions dont TU connais la réponse) :

```json
{"q": "Quelle commande affiche l'état des interfaces sur un switch Huawei ?",
 "attendu_source": "huawei_cli",
 "attendu_section": "display interface",
 "mots_cles": ["display", "interface", "brief"]}
{"q": "Quelle est la consommation PoE de l'AP361 ?",
 "attendu_source": "guides",
 "attendu_section": "AP361",
 "mots_cles": ["8,8 W", "802.3af"]}
{"q": "Comment configurer iStack sur un S310 ?",
 "attendu_source": "guides",
 "mots_cles": ["iStack", "S310", "stack"]}
```

Métriques :

```python
def hit_rate(resultats: list[list[str]], attendu: list[str], k: int = 5) -> float:
    """% de questions dont la bonne source est dans le top-k."""
    hits = sum(1 for r, a in zip(resultats, attendu) if a in r[:k])
    return hits / len(attendu)

def mrr(resultats, attendu) -> float:
    """Mean Reciprocal Rank : 1/rang du premier bon résultat, moyenné."""
    s = 0.0
    for r, a in zip(resultats, attendu):
        if a in r:
            s += 1 / (r.index(a) + 1)
    return s / len(attendu)
```

**Protocole :**
1. Baseline : chunking 512/50, seuil 0.70 → note `hit_rate@5`, `MRR`.
2. Change **un** paramètre (ex. : sections vs fixe) → re-mesure.
3. Garde ce qui améliore le score **sur TON jeu**, pas sur un benchmark générique.
4. **Inspection manuelle** : lis 10 résultats par semaine — l'œil humain
   attrape ce que les métriques ratent (chunk tronqué, mauvaise section).

**Hybride vectoriel + BM25** (quand le vectoriel seul plafonne sur les
termes exacts comme les noms de commandes) :

```sql
ALTER TABLE rag_chunks ADD COLUMN search_tsv tsvector
  GENERATED ALWAYS AS (to_tsvector('french', content)) STORED;
CREATE INDEX idx_chunks_tsv ON rag_chunks USING gin (search_tsv);
-- Puis combine : filtre tsvector + tri vectoriel (voir §67).
```

## 103. Maintenance : re-embedder quand ?

| Événement | Action |
|---|---|
| Nouveau lot de documents | embed + insert (incrémental), MAJ manifeste |
| Changement de chunker | **re-embed complet** du corpus concerné |
| Changement de modèle d'embedding | re-embed complet + migration de colonne |
| Document mis à jour | delete par `source`+`doc_title`, re-insert |
| Dérive qualité (hit_rate ↓) | audit : nouveaux doublons ? nouveau vocabulaire ? |
| Rien pendant 6 mois | re-mesure le golden set (les questions évoluent) |

```sql
-- Invalider proprement une version de corpus :
DELETE FROM rag_chunks WHERE corpus_version = 'v2';
-- (puis re-ingère en v3 — jamais de mélange de versions dans l'index)
```

> Le re-embed complet de ton corpus coûte ~$0.10 (§88) : **le coût n'est
> jamais une raison de garder des vecteurs périmés.**

## 104. Pipeline d'ingestion : orchestrateur

```bash
#!/usr/bin/env bash
# scripts/ingest.sh — pipeline complet, rejouable
set -euo pipefail
VERSION="v3"
mkdir -p build logs

echo "[1/5] Nettoyage…"
python scripts/clean_corpus.py --input corpus/brut --output build/clean

echo "[2/5] Chunking…"
python scripts/chunk_md.py     --input build/clean/guides --output build/chunks_guides.jsonl --source guides
python scripts/chunk_fixed.py  --input build/clean/articles.json --output build/chunks_articles.jsonl

echo "[3/5] Dédup…"
cat build/chunks_*.jsonl | python scripts/dedup_stream.py > build/chunks_dedup.jsonl

echo "[4/5] Embeddings…"
python scripts/embed_corpus.py --input build/chunks_dedup.jsonl --output build/embedded.jsonl

echo "[5/5] Ingestion pgvector…"
python scripts/ingest_pgvector.py --input build/embedded.jsonl --version "$VERSION"

echo "OK — $(wc -l < build/embedded.jsonl) chunks en $VERSION"
```

Chaque étape lit/écrit des fichiers → **rejouable** étape par étape,
interrompable, débuggable. Loggue tout (`tee logs/ingest_$VERSION.log`).

## 105. Filtrage : langue, PII, qualité

```python
def quality_gate(text: str) -> str | None:
    """Retourne la raison du rejet, ou None si le texte est gardé."""
    if len(text.strip()) < 100:
        return "trop court"
    if len(set(text)) < 20:
        return "alphabet dégénéré (binaire/ocr cassé)"
    # Proportion de caractères bizarres (�, contrôles)
    weird = sum(1 for c in text if c == "�" or ord(c) < 9)
    if weird / max(len(text), 1) > 0.01:
        return "corruption d'encodage"
    return None

# Langue (optionnel) : pip install langdetect — à vérifier : précision sur textes courts
# PII : si ton corpus contient des tickets/incidents avec des noms,
# envisage un filtre (regex emails/télés + modèle NER) AVANT l'indexation.
```

## 106. 20 pièges du corpus

