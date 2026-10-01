---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-37
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "Cohere", "Nvidia", "vLLM"]
dates: []
keywords: ["amd", "awq", "blackwell", "cohere", "dram", "embedding", "embeddings", "fine-tuning", "gpu", "hbm3", "kv cache", "llama"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [2712, 2804]
sha256: 5ca02992ea81ef4d782b917e59ef1221ec214ad5e8e3bda89ca1f490218639f4
---

# IA — Le grand dossier

**Impact du contexte** (KV cache, ordres de grandeur pour un 7–8B) : 8K → +0,5–1 Go ; 32K → +2–4 Go ; 128K → +10 Go et plus. Un modèle qui « tient » à 4K peut OOM à 32K : **dimensionner avec le contexte cible, pas avec les poids seuls**.

### 6.2. Les plateformes : Mac (mémoire unifiée) vs NVIDIA vs CPU

| Plateforme | Modèle(s) repère | Mémoire adressable | Bande passante | Débit indicatif (8B Q4, 1 utilisateur) | Prix indicatif (sept. 2026) | Verdict |
|---|---|---|---|---|---|---|
| **RTX 5090** (Blackwell) | Tour + carte | **32 Go GDDR7** (mur dur) | **1 792 Go/s** | ~140–190 tok/s | ~2 500–3 800 $ (marché ; MSRP 1 999 $) | Le plus rapide en <32B ; 575 W ; **ne tient pas un 70B Q4** |
| **RTX 4090** (Ada) | Tour + carte | **24 Go GDDR6X** | 1 008 Go/s | ~100–130 tok/s | ~1 500–2 200 $ (occasion ; +24 % depuis mars, crise DRAM) | Le meilleur rapport perf/prix <27B ; le plus testé |
| **RTX 6000 Ada** | Workstation pro | **48 Go GDDR6** | 960 Go/s | ~90–110 tok/s (à vérifier) | ~7 000 $+ (à vérifier) | 70B Q4 sur une seule carte, budget pro |
| **2× RTX 3090/4090** | Tour double GPU | **48 Go** (2×24) | 2× ~1 000 Go/s (PCIe entre elles) | 70B Q4 : ~15–30 tok/s (à vérifier) | ~2 500–4 500 $ (occasion) | Le 70B « bricolé » : ça marche, plus lent |
| **Mac Studio M5 Ultra** | Tout-en-un Apple | **96–512 Go unifiés** | **~1 200 Go/s** | 8B : ~95 tok/s ; 70B Q4 : plafond théorique ~28 tok/s | dès **5 499 $** (96 Go) | **Le seul desktop qui fait du 70B+ en silence** (<110 W) |
| **Mac Studio M4 Max** | Tout-en-un Apple | 36–128 Go unifiés | 546 Go/s (à vérifier) | 8B : ~70–80 tok/s ; 70B Q4 : ~12 tok/s | 1 999–3 799 $ | 32B confortable, 70B limite en 128 Go |
| **NVIDIA DGX Spark** | Mini-station NVIDIA | **128 Go unifiés** | 273 Go/s | 7B Q4 : ~78 tok/s | 3 999 $ | 120B Q4 tient ; moins rapide que le Mac à mémoire égale |
| **AMD RX 7900 XTX** | Tour + carte | 24 Go GDDR6 | 960 Go/s | ~100–130 tok/s (7B) | ~700–1 100 $ (occasion) | Prix/VRAM imbattable ; **ROCm moins mature** (edge cases llama.cpp) |
| **CPU seul** (serveur/PC, AVX-512) | — | RAM système (64 Go+) | ~50–100 Go/s | 8B Q4 : ~10–20 tok/s ; 70B : 1–5 tok/s | 0 $ (existant) | Embeddings, batch nuit, petits modèles ; **pas d'interactif 70B** |
| **H100 80 Go** (serveur) | Serveur GPU | 80 Go HBM3 | ~3 350 Go/s | 70B Q4 : ~80–120 tok/s (à vérifier) | ~150–450 k$ le nœud 8× / 2,5–3,3 $/h en cloud | Prod multi-utilisateurs (vLLM) |

**Enseignements** :

1. **Le mur des 32 Go** : la 5090 est la carte la plus rapide du marché… et elle ne fait pas tenir un 70B Q4 (43 Go). Vitesse sans capacité = 32B max en full-GPU.
2. **La mémoire unifiée change la donne** : sur Mac/Spark, CPU+GPU partagent la même mémoire → un 70B Q4 (43 Go) tient dans 64 Go unifiés là où il faut 2 GPU côté NVIDIA. En échange, la bande passante est 2–3× moindre → moins de tok/s.
3. **Pas de CUDA sur Mac** : pas de vLLM « classique », pas d'entraînement PyTorch standard → écosystème **MLX** (qui mûrit vite mais reste un sous-ensemble). Si ton workflow touche à autre chose que l'inférence texte (fine-tuning, multimodal lourd), NVIDIA reste obligatoire.
4. **L'occasion a un prix** : crise DRAM en 2026 — une 3090 d'occasion ~1 252 $, une 4090 ~2 150–2 350 $ (vérifié sept. 2026). Calculer l'amortissement sur ces prix réels, pas sur les MSRP.

### 6.3. Trois configurations types (budgets 2026)

| Profil | Config | Coût matériel | Fait tourner | Débit typique |
|---|---|---|---|---|
| **A. Poste dev** | RTX 4090 24 Go (ou 5090 32 Go) + 64 Go RAM | ~2 000–3 500 $ | 8B–32B Q4 full-GPU ; embeddings ; 70B en offload partiel (lent) | 100–190 tok/s (8B) |
| **B. Serveur d'équipe** | 2× RTX 4090 48 Go ou 1× RTX 6000 Ada 48 Go, vLLM | ~4 500–8 000 $ | 70B Q4 servi à 5–10 utilisateurs ; 32B AWQ très rapide | 15–40 tok/s/utilisateur (70B) |
| **C. Station silencieuse** | Mac Studio M4 Max 128 Go (ou M5 Ultra 96 Go) | ~3 800–5 500 $ | 70B Q4 entièrement en mémoire ; 100B+ Q4 en 128 Go | ~12–28 tok/s (70B) |

**Coût électrique** (ordre de grandeur, France ~0,20 €/kWh) : 4090 à 450 W en charge ≈ 0,09 €/h ; Mac Studio à ~40–110 W ≈ 0,01–0,02 €/h. Sur 2 000 h/an : ~180 € vs ~30 €. Le Mac se rembourse en silence et en watts, la NVIDIA en tok/s.

### 6.4. Méthode de dimensionnement (checklist)

1. **Choisir le modèle** (qualité requise) → taille en B.
2. **Choisir la quantification** : Q4_K_M par défaut ; Q8 si la VRAM le permet et que la tâche est exigeante (code, maths) ; jamais en dessous de Q4 pour ≤14B en prod.
3. **Additionner** : poids (tableau §6.1) + KV cache (contexte cible) + 15 % d'overhead + 8 Go pour l'OS (Mac).
4. **Comparer** à la mémoire rapide disponible. Si ça dépasse de <30 % : offload partiel CPU (llama.cpp `--n-gpu-layers`) en dégradé, ou baisser le contexte. Si ça dépasse largement : modèle plus petit, ou Mac/2×GPU, ou cloud.
5. **Valider** par `llama-bench` / un test de charge réel avant d'acheter : la théorie ne remplace pas une mesure.

---

## 7. Cas d'usage professionnels concrets

### 7.1. RAG local : indexer la doc technique sans qu'elle sorte du site

C'est **ton** cas d'usage (le RAG de Zelef). Architecture minimale viable :

```
[Docs PDF/md] -> [découpage/chunking] -> [embeddings locaux] -> [vecteurs]
[Question] -> [embeddings locaux] -> [recherche] -> [top-k chunks]
  -> [prompt : question + chunks] -> [LLM local ou cloud] -> [réponse + sources]
```

**Choix des composants** (éprouvés, pas de hype) :

| Brique | Option locale (coût nul) | Option cloud (qualité max) | Recommandation |
|---|---|---|---|
| Embeddings | `nomic-embed-text` (Ollama, 274 Mo) ou `Qwen3-Embedding-8B` (vLLM) | `text-embedding-3-small` (0,02 $/1M, 1536 dims) — **ton choix actuel** | Hybride : local en dev/qualif, cloud en prod si l'eval le justifie |
| Base vectorielle | **Qdrant** / **pgvector** (Postgres) / Chroma | Pinecone, Weaviate Cloud | Qdrant ou pgvector : tu sais déjà opérer Postgres |
| LLM génération | Qwen3-32B Q4 (vLLM, 4090) ou 8B (Ollama) | `triage`/`raisonnement` via LiteLLM (§4) | **Routage par criticité** : local par défaut, escalade cloud si confiance basse |
| Reranker | `Qwen3-Reranker` ou Jina Reranker local | Cohere Rerank / Voyage | Le rerank local apporte souvent +5–15 pts de précision pour ~0 $ |

**Script d'indexation minimal** (Python, `sentence-transformers` + Qdrant — remplace par ton pipeline réel) :

```python
# index_local.py — indexation 100 % locale (aucune donnée ne sort)
from sentence_transformers import SentenceTransformer
from qdrant_client import QdrantClient
from qdrant_client.models import Distance, VectorParams, PointStruct
from pypdf import PdfReader
import glob, uuid

model = SentenceTransformer("nomic-ai/nomic-embed-text-v1", trust_remote_code=True)
qdrant = QdrantClient(host="qdrant-interne.lan", port=6333)
qdrant.recreate_collection("doc-tech", vectors_config=VectorParams(size=768, distance=Distance.COSINE))

def chunks(text, size=800, overlap=120):
    return [text[i:i+size] for i in range(0, len(text), size-overlap)]

points = []
for pdf in glob.glob("/data/docs/**/*.pdf", recursive=True):
    text = "\n".join(p.extract_text() or "" for p in PdfReader(pdf).pages)
    for c in chunks(text):
        points.append(PointStruct(id=str(uuid.uuid4()),
                                  vector=model.encode(c).tolist(),
                                  payload={"source": pdf, "text": c}))
qdrant.upsert("doc-tech", points)
print(f"{len(points)} chunks indexés — 0 appel API, 0 €.")
```

