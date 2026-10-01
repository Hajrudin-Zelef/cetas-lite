---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-15
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: ["embeddings"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [2780, 2974]
sha256: ff80e75c7d453b178b89aada616540046a6d7f9b608a5abc729378f5462a807c
---

# Outils dev + ingénierie RAG (chunk & corpus)

Ton pipeline en 5 passes est déjà bon — formalise-le :

```text
listes/lotX.txt
  → collect_failed.py   (requests)      → data/lotX_p1.jsonl
  → collect_pw.py       (requests)      → data/lotX_p2.jsonl
  → collect_still.py    (4 workers)     → data/lotX_p3.jsonl
  → collect_final.py    (APIs)          → data/lotX_p4.jsonl
  → collect_pw2.py      (Playwright §53) → data/lotX_p5.jsonl
  → failed_dead.txt                     → étape cloud (Leo)
```

Chaque passe écrit le **même format** (`articles_*.json` :
`url/title/file/text/method`) + met à jour un **manifeste** :

```json
{
  "lot": "huawei_edoc_3",
  "date": "2026-09-27",
  "urls_total": 452,
  "passes": [
    {"script": "collect_failed.py", "ok": 380, "restants": 72},
    {"script": "collect_pw.py", "ok": 45, "restants": 27}
  ],
  "failed_dead": "listes/huawei_edoc_3_dead.txt"
}
```

## 95. Nettoyage : strip headers/footers (ton cas Huawei !)

Les PDF convertis (`pdftotext`) et les pages scrapées contiennent des
**répétitions** : en-têtes/pieds de page (« NetEngine AR1000V Command
Reference », numéros de page), menus de navigation, sommaires. Ces
répétitions **polluent les embeddings** (le même texte revient dans 600
chunks) et **gonflent la facture**.

```python
"""Nettoyage : supprime headers/footers répétitifs d'un texte multi-pages."""
import re
from collections import Counter

def strip_repeated_lines(pages: list[str], seuil: float = 0.5) -> list[str]:
    """Supprime les lignes qui apparaissent dans > seuil des pages
    (typiquement : titres de doc répétés, numéros de page)."""
    n = len(pages)
    if n < 3:
        return pages
    compteur = Counter()
    for p in pages:
        for ln in set(p.splitlines()):       # set : 1 vote par page
            s = ln.strip()
            if s:
                compteur[s] += 1
    parasites = {ln for ln, c in compteur.items() if c / n > seuil}
    # Garde-fou : ne jamais supprimer une ligne longue (vrai contenu)
    parasites = {ln for ln in parasites if len(ln) < 120}
    propres = []
    for p in pages:
        propres.append("\n".join(
            ln for ln in p.splitlines() if ln.strip() not in parasites))
    return propres

def strip_pdf_artifacts(text: str) -> str:
    """Nettoie les artefacts pdftotext : numéros de page isolés, etc."""
    text = re.sub(r"\n\s*\d+\s*\n", "\n", text)          # "  2317  " seul sur sa ligne
    text = re.sub(r"[ \t]{2,}", " ", text)              # espaces multiples
    text = re.sub(r"\n{3,}", "\n\n", text)              # lignes vides en trop
    # Mots coupés en fin de ligne par pdftotext : "vpn-instancename"
    text = re.sub(r"(\w)-\n(\w)", r"\1\2", text)         # à vérifier sur échantillon !
    return text.strip()
```

> **Valide sur échantillon** : affiche 5 pages avant/après et relis.
> Le `(\w)-\n(\w)` peut fusionner à tort (« porte-monnaie » coupé
> volontairement). Ton nettoyage des CLI Huawei a déjà rencontré ce cas
> (micro-défauts de dé-hyphenation).

## 96. Normalisation

```python
import unicodedata, re

def normalize(text: str) -> str:
    # 1. Unicode : NFC (é = 1 caractère, pas e + accent)
    text = unicodedata.normalize("NFC", text)
    # 2. Guillemets/tirets typographiques → ASCII (optionnel, mais aide
    #    la cohérence des embeddings sur corpus mixte FR/EN)
    #    À évaluer : pour du FR soigné, tu peux aussi les GARDER.
    # 3. Espaces insécables, tabulations
    text = text.replace("\u00a0", " ").replace("\t", " ")
    # 4. Whitespace
    text = re.sub(r"[ ]{2,}", " ", text)
    text = re.sub(r"\n{3,}", "\n\n", text)
    return text.strip()

def detect_encoding(path) -> str:
    """Les vieux PDF/CSV ne sont pas toujours en UTF-8."""
    import chardet  # pip install chardet — à vérifier : ou charset-normalizer
    raw = open(path, "rb").read(100000)
    return chardet.detect(raw)["encoding"] or "utf-8"
```

**Encodage** : décode TOUJOURS en UTF-8 en sortie (`errors="replace"` en
dernier recours, mais loggue les fichiers concernés — un `�` dans un chunk
est un signal de corruption).

## 97. Déduplication exacte (rappel + au niveau document)

§87 faisait la dédup au niveau **chunk**. Fais-la aussi au niveau
**document** (avant chunking) : deux collectes du même article (URL avec
et sans `?utm_source=...`) = même contenu.

```python
import hashlib, json

def doc_hash(text: str) -> str:
    return hashlib.sha256(text.encode("utf-8")).hexdigest()

def normalize_url(url: str) -> str:
    """Canonicalise pour la dédup : minuscules, sans query tracking, sans / final."""
    from urllib.parse import urlparse, urlunparse
    p = urlparse(url.lower())
    return urlunparse((p.scheme, p.netloc, p.path.rstrip("/"), "", "", ""))
```

## 98. Near-dup : MinHash expliqué + exemple

**Le problème :** deux documents à 95 % identiques (même article, intro
différente ; doc versionnée v2/v3) passent la dédup exacte mais polluent
le retrieval (5 copies du même fait dans le top-k — étude 2026 : 5 copies
= même score qu'**1** document, contre 5 documents **divers** bien
meilleurs).

**MinHash, l'idée en 30 secondes :**
1. **Shingling** : découpe le doc en k-grammes de mots (ex. : 5 mots
   consécutifs) → un *ensemble* de shingles.
2. **MinHash** : hachage qui résume l'ensemble en une courte *signature*
   (ex. : 128 entiers). Deux docs similaires ont des signatures similaires.
3. **LSH (banding)** : découpe les signatures en bandes pour ne comparer
   que les paires probablement similaires → évite le O(n²).

```bash
pip install datasketch   # lib standard pour MinHash/SimHash (vérifié : existe)
```

```python
"""Near-dup avec MinHash + LSH (datasketch)."""
from datasketch import MinHash, MinHashLSH

def shingles(text: str, k: int = 5) -> list[str]:
    mots = text.lower().split()
    return [" ".join(mots[i:i + k]) for i in range(len(mots) - k + 1)]

def signature(text: str, num_perm: int = 128) -> MinHash:
    mh = MinHash(num_perm=num_perm)
    for s in shingles(text):
        mh.update(s.encode("utf-8"))
    return mh

def find_near_dups(docs: list[dict], seuil: float = 0.85) -> list[tuple[int, int, float]]:
    """docs = [{"id":..., "text":...}]. Retourne les paires (i, j, similarité)."""
    lsh = MinHashLSH(threshold=seuil, num_perm=128)
    sigs = {}
    for i, d in enumerate(docs):
        mh = signature(d["text"])
        sigs[i] = mh
        lsh.insert(f"doc_{i}", mh)
    paires = set()
    for i, mh in sigs.items():
        for key in lsh.query(mh):
            j = int(key.split("_")[1])
            if j > i:
                # Similarité de Jaccard estimée via les signatures
                paires.add((i, j, round(mh.jaccard(sigs[j]), 3)))
    return sorted(paires)

# Politique : pour chaque paire near-dup, garde le doc le PLUS RÉCENT
# (ou le plus complet), archive l'autre avec la raison dans le manifeste.
```

**Réglages :** `threshold=0.85` (Jaccard), `num_perm=128` (précision),
shingles de 5 mots. **Piège (§121) :** scope par `section`/titre —
« Installer sur Linux » vs « Installer sur Windows » ont un corps quasi
identique mais un sens différent (vérifié : article dev.to 2026).

## 99. SimHash : l'alternative légère

**SimHash** = une seule empreinte de 64 bits par document ; deux docs
proches ont des empreintes qui diffèrent de peu de bits (distance de
Hamming). Moins précis que MinHash+LSH, mais **minuscule en mémoire**
(parfait pour un premier tri sur 100k+ docs).

```python
from datasketch import SimHash

