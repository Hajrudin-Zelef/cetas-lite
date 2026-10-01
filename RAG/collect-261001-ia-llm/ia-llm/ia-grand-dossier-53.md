---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-53
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Meta", "Microsoft", "Mistral", "OpenAI", "vLLM"]
dates: []
keywords: ["agent", "agents", "awq", "benchmarks", "claude", "embedding", "embeddings", "mcp", "mistral", "model context protocol", "multimodal", "muse"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [4190, 4308]
sha256: ba03fffb8d74677b70a33bfa06e1e20d9b7b12a26b980e44a34b86af940a15e1
---

# IA — Le grand dossier

for ch in chunk_markdown("kyocera_copieurs_guide.md"):
    emb = model.encode("search_document: " + ch["text"]).tolist()
    cur.execute("INSERT INTO docs(source, chunk, embedding) VALUES (%s,%s,%s)",
                (ch["source"], ch["text"], emb))
conn.commit()

# === 2. REQUÊTE : hybride vectoriel + BM25, fusion RRF, rerank ===
from sentence_transformers import CrossEncoder
reranker = CrossEncoder("BAAI/bge-reranker-v2-m3")

def search(question, k_vec=20, k_final=5):
    q_emb = model.encode("search_query: " + question).tolist()
    # a) vectoriel (cosine) — top 20
    cur.execute("""SELECT id, chunk, source,
                   1 - (embedding <=> %s::vector) AS sim
                   FROM docs ORDER BY embedding <=> %s::vector LIMIT 20""",
                (q_emb, q_emb))
    vec_hits = cur.fetchall()
    # b) BM25 lexical (tsvector français) — top 20 : capte "U163", "C6000", "VFI"…
    cur.execute("""SELECT id, chunk, source,
                   ts_rank(tsv, plainto_tsquery('french', %s)) AS rank
                   FROM docs WHERE tsv @@ plainto_tsquery('french', %s)
                   ORDER BY rank DESC LIMIT 20""", (question, question))
    bm25_hits = cur.fetchall()
    # c) fusion RRF (reciprocal rank fusion), k=60
    scores = {}
    for rank, h in enumerate(vec_hits):  scores[h[0]] = scores.get(h[0], 0) + 1/(60+rank+1)
    for rank, h in enumerate(bm25_hits): scores[h[0]] = scores.get(h[0], 0) + 1/(60+rank+1)
    top_ids = sorted(scores, key=scores.get, reverse=True)[:k_vec]
    # d) rerank cross-encoder → top 5
    by_id = {h[0]: h for h in vec_hits + bm25_hits}
    pairs = [(question, by_id[i][1]) for i in top_ids]
    ranked = sorted(zip(top_ids, reranker.predict(pairs)), key=lambda x: x[1], reverse=True)
    return [by_id[i] for i, _ in ranked[:k_final]]

# === 3. GÉNÉRATION avec citations ===
from openai import OpenAI
llm = OpenAI(base_url="http://localhost:8000/v1", api_key="x")  # ton vLLM local, ou API cloud

def answer(question):
    hits = search(question)
    context = "\n\n".join(f"[Source: {h[2]}]\n{h[1][:1500]}" for h in hits)
    r = llm.chat.completions.create(model="Qwen/Qwen3-8B-AWQ", temperature=0.15, max_tokens=600,
        messages=[
            {"role": "system", "content": (
                "Tu es un assistant technique. Réponds UNIQUEMENT à partir du contexte fourni. "
                "Cite tes sources entre crochets [source]. Si le contexte ne contient pas la réponse, "
                "dis-le explicitement au lieu d'inventer.")},
            {"role": "user", "content": f"Contexte:\n{context}\n\nQuestion: {question}"}])
    return r.choices[0].message.content, [h[2] for h in hits]

rep, srcs = answer("Comment sortir du mode maintenance 10871087 sur un copieur Kyocera ?")
print(rep); print("Sources:", set(srcs))
```

**Variantes par besoin** (à connaître) :

| Besoin | Stack |
|---|---|
| 100 % local, zéro cloud | Ollama (embeddings `nomic-embed-text` + LLM) + **Chroma** ou SQLite-vec + BM25 en Python (`rank-bm25`) |
| Prod PME | **Qdrant** ou **pgvector** (si tu as déjà Postgres — ton cas !) + vLLM + reranker local |
| Scale / SaaS | Pinecone/Weaviate managés + API embeddings + API LLM |
| Multimodal (PDF scannés, schémas) | OCR (Tesseract/Mistral OCR) + embeddings multimodaux (ex. : ColPali pour la recherche visuelle dans les PDF — « à vérifier » sur la maturité 2026) |
| Temps réel / agentique | **GraphRAG** (Microsoft) : graphe de connaissances + communautés — puissant mais coûteux à construire |

**Anti-patterns RAG (à afficher au mur) :**

1. Chunker en aveugle à taille fixe sans respecter la structure du document.
2. Embedder avec un modèle différent entre l'indexation et la requête (les vecteurs ne sont pas compatibles !).
3. Pas de filtre par ACL → fuite de documents confidentiels (voir 2.1.3).
4. Température haute (0.7+) sur de la réponse factuelle → hallucinations élégantes.
5. Aucun jeu d'évaluation → impossible de savoir si un changement améliore ou dégrade.
6. Index jamais rafraîchi → le RAG répond avec des docs obsolètes en toute confiance.

### 4.4. Stack agents (les « starck » qui agissent)

**À quoi ça sert** : passer du chatbot qui *répond* à l'agent qui *fait* (lit ta boîte mail,
ouvre un ticket GLPI, redémarre un service — avec garde-fous).

```
┌────────────┐   ┌──────────────┐   ┌───────────────┐   ┌────────────────┐
│  Modèle    │◀─▶│  Boucle      │◀─▶│  Outils       │◀─▶│  Mémoire       │
│  (LLM      │   │  agent       │   │  (functions,  │   │  (contexte,    │
│  + raison.)│   │  ReAct :     │   │  MCP, API,    │   │  vectorielle,  │
│            │   │  pense → agit│   │  shell, code) │   │  episodic)     │
│            │   │  → observe   │   │               │   │                │
└────────────┘   └──────────────┘   └───────────────┘   └────────────────┘
        │                │                   │
        ▼                ▼                   ▼
   Garde-fous : validation humaine, moindre privilège, audit log, limites (timeout, budget tokens)
```

| Couche | Outils réels (2026) |
|---|---|
| Frameworks | **LangChain / LangGraph** (standard de facto, graphes d'agents), **LlamaIndex** (RAG → agents), **CrewAI** (multi-agents par rôles), **AutoGen** (Microsoft, conversations multi-agents), **OpenAI Agents SDK**, **Claude Agent SDK** (Anthropic, janv. 2026), **Muse** (framework d'orchestration, cf. vol. outillage) |
| Protocole d'outils | **MCP** (Model Context Protocol, Anthropic fin 2024 → standard 2025-2026) : un serveur MCP expose outils/données, n'importe quel client compatible s'y branche. Alternative : **A2A** (Agent-to-Agent, Google) pour agents qui se parlent entre eux |
| Orchestration | n8n / Make (no-code), Temporal / Prefect (workflows durables), code maison |
| Éval d'agents | Benchmarks : SWE-bench (code), Terminal-Bench, GAIA (assistants généralistes), **τ-bench** (agents avec outils) ; en interne : taux de succès par scénario + revue humaine |
| Observabilité | **Langfuse**, **LangSmith**, Arize Phoenix : traces, coûts par run, qualité — indispensable dès qu'un agent touche à la prod |

**Exemple réel — agent minimal avec outils (boucle ReAct maison, ~40 lignes, sans framework) :**

```python
import json, subprocess
from openai import OpenAI
llm = OpenAI(base_url="http://localhost:8000/v1", api_key="x")

TOOLS = [
    {"type": "function", "function": {
        "name": "run_diag", "description": "Exécute un diagnostic réseau en lecture seule",
        "parameters": {"type": "object", "properties": {
            "cmd": {"type": "string", "enum": ["ping_gw", "dns_test", "iface_status"]}}}}}
]
def run_diag(cmd):  # whitelist stricte : JAMAIS de shell libre
    cmds = {"ping_gw": ["ping", "-c", "3", "192.168.1.1"],
            "dns_test": ["nslookup", "example.com"],
            "iface_status": ["ip", "-br", "addr"]}
    return subprocess.run(cmds[cmd], capture_output=True, text=True, timeout=15).stdout

