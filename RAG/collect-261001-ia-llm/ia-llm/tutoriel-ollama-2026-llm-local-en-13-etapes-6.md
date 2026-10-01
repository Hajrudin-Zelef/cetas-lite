---
id: collect-261001-ia-llm/ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes-6
title: "Vérifier la version du pilote NVIDIA"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Apple", "EU", "Mistral", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["nvidia", "amd", "claude", "embeddings", "gguf", "gpu", "llama", "llama.cpp", "mistral", "qwen", "safetensors", "valuation"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes.md
source_anchor: ""
source_lines: [460, 578]
sha256: f64a26bca4e52e5d068595b527b9231182197147bddc19c6aab67dab3a5f7b9e
---

# Vérifier la version du pilote NVIDIA

```
# Structure du projet
# assistant-rh/
# ├── docs/                  # PDF des procédures RH
# ├── main.py                # API FastAPI
# ├── ingest.py              # Indexation initiale
# ├── requirements.txt
# └── docker-compose.yml
# requirements.txt
fastapi==0.115.5
uvicorn==0.32.0
pypdf==5.1.0
chromadb==0.5.13
ollama==0.4.2
python-multipart==0.0.17
# ingest.py — Indexer tous les PDF du dossier docs/
import os
from pypdf import PdfReader
import ollama
import chromadb
client = chromadb.PersistentClient(path="./db")
collection = client.get_or_create_collection("rh")
def chunk_text(text, size=600, overlap=80):
    chunks = []
    start = 0
    while start < len(text):
        chunks.append(text[start:start+size])
        start += size - overlap
    return chunks
for pdf_file in os.listdir("docs"):
    if not pdf_file.endswith(".pdf"):
        continue
    reader = PdfReader(f"docs/{pdf_file}")
    full_text = "\n".join(p.extract_text() for p in reader.pages)
    chunks = chunk_text(full_text)
    for i, chunk in enumerate(chunks):
        emb = ollama.embed(model="nomic-embed-text", input=chunk)
        collection.add(
            ids=[f"{pdf_file}_{i}"],
            embeddings=[emb["embeddings"][0]],
            documents=[chunk],
            metadatas=[{"source": pdf_file, "chunk": i}]
        )
    print(f"Indexé : {pdf_file} ({len(chunks)} chunks)")
# main.py — API HTTP
from fastapi import FastAPI
from pydantic import BaseModel
import ollama
import chromadb
app = FastAPI(title="Assistant RH")
client = chromadb.PersistentClient(path="./db")
collection = client.get_collection("rh")
class Question(BaseModel):
    text: str
    k: int = 4
@app.post("/ask")
def ask(q: Question):
    q_emb = ollama.embed(model="nomic-embed-text", input=q.text)
    results = collection.query(
        query_embeddings=[q_emb["embeddings"][0]],
        n_results=q.k
    )
    context = "\n\n".join(results["documents"][0])
    sources = list({m["source"] for m in results["metadatas"][0]})
    response = ollama.chat(
        model="llama3.1:8b",
        messages=[
            {"role": "system", "content":
             "Tu es l assistant RH. Réponds uniquement avec le contexte fourni. "
             "Si la réponse n est pas dans le contexte, dis-le."},
            {"role": "user", "content":
             f"Contexte :\n{context}\n\nQuestion : {q.text}"}
        ],
        options={"temperature": 0.2, "num_ctx": 8192}
    )
    return {
        "answer": response["message"]["content"],
        "sources": sources
    }
# Lancer
# python ingest.py
# uvicorn main:app --host 0.0.0.0 --port 8000
# curl -X POST http://localhost:8000/ask -H "Content-Type: application/json" \
#      -d '{"text":"Combien de jours de congés payés par an ?"}'
```
Ce projet de 80 lignes vous donne un assistant entièrement local, conforme RGPD, capable de répondre sur vos documents internes sans qu’aucune donnée ne quitte la machine. Sur un RTX 4060, le temps de réponse moyen est de 2 à 4 secondes pour une question typique. Pour passer en production, ajoutez une authentification (JWT via FastAPI Security), un rate-limit (slowapi), un cache Redis pour les questions fréquentes, et un monitoring Prometheus/Grafana. Vous pouvez aussi remplacer Llama 3.1 8B par Mistral 7B ou Qwen 3 7B selon vos tests de qualité.

## Comparatif Ollama vs alternatives en 2026

Ollama n’est pas seul sur le marché des runtimes LLM locaux. Plusieurs alternatives méritent d’être considérées selon votre cas d’usage. Pour les déploiements à très haute charge (>100 req/s), **vLLM** offre un débit 3 à 9× supérieur grâce à PagedAttention et au continuous batching, au prix d’une complexité d’installation bien plus élevée. Pour les setups graphiques sans code, **LM Studio** propose une UI desktop conviviale mais ferme l’extensibilité. Pour les développeurs Python qui veulent le contrôle total, **llama.cpp** reste la référence bas niveau.

| Critère | Ollama | vLLM | LM Studio | llama.cpp | 
|---|---|---|---|---|
| Installation | 1 commande | pip + CUDA | App desktop | Compile C++ | 
| API HTTP native | Oui (/api/*) | Oui (OpenAI) | Oui (OpenAI) | Non (server à part) | 
| Format modèles | GGUF + tags | HF safetensors | GGUF | GGUF | 
| Débit (8B Q4) | 80-110 t/s | 200-350 t/s | 70-90 t/s | 90-120 t/s | 
| Multi-modèle | Oui | 1 par GPU | 1 chargé | 1 par instance | 
| GPU NVIDIA | Oui (CUDA) | Oui | Oui | Oui | 
| GPU AMD (ROCm) | Oui (v0.6.2+) | Limité | Beta | Oui | 
| Apple Silicon | Oui (Metal 3) | Non | Oui | Oui | 
| OpenAI compat | Oui (/v1) | Oui | Oui | Via wrapper | 
| Docker officiel | Oui | Oui | Non | Communauté | 
| Cas d’usage | Dev, RAG, équipe | Production scale | Solo no-code | Bas niveau | 

Pour 80 % des projets français, Ollama est le bon choix par défaut. Il offre le meilleur ratio simplicité/fonctionnalités, et l’écart de débit avec vLLM ne devient critique qu’au-delà de quelques dizaines de requêtes par seconde simultanées. Pour les déploiements à très grande échelle (chatbot grand public avec des milliers d’utilisateurs concurrents), basculer vers vLLM ou un service géré comme Mistral Le Chat Pro ou Anthropic Claude API devient plus pertinent économiquement.

## Sécurité, RGPD et déploiement en entreprise

Le principal argument d’Ollama en entreprise française est sa nature locale : aucune donnée ne quitte votre infrastructure, ce qui simplifie radicalement la conformité RGPD et AI Act. Cela ne signifie pas pour autant qu’Ollama soit « sécurisé par défaut ». Plusieurs précautions s’imposent en production. D’abord, ne jamais exposer le port 11434 sur Internet sans authentification : il n’y a pas de mécanisme natif d’auth dans Ollama. Placez-le derrière un reverse proxy Nginx ou Traefik avec basic auth ou OAuth2 Proxy.

Ensuite, isolez le réseau : si Ollama tourne sur un serveur dédié, utilisez un firewall (UFW, nftables) pour autoriser uniquement les IP des serveurs applicatifs autorisés. Pour les déploiements multi-tenants où plusieurs équipes partagent un même cluster, considérez un proxy intelligent comme LiteLLM qui ajoute la gestion des clés API, le rate-limiting et l’audit log. Côté logs, Ollama écrit dans `~/.ollama/logs/server.log` sur Linux : vérifiez régulièrement la taille (rotation via logrotate) et auditez les requêtes inhabituelles.

Concernant la conformité AI Act (entré en application en France selon le calendrier du Digital Omnibus IA reporté au 2 décembre 2027), Ollama lui-même n’est qu’un runtime : ce sont les modèles que vous exécutez qui peuvent tomber sous les obligations de l’AI Act (transparence, évaluation des risques). Llama 3.3 et Mistral sont publiés sous licences permissives mais leurs cartes modèles doivent être archivées si vous les utilisez en production sur des cas à haut risque (RH, santé, éducation).

## Dépannage : 8 problèmes fréquents et solutions

Voici les huit erreurs les plus fréquemment rencontrées par les utilisateurs Ollama, avec leurs solutions vérifiées. Avant de chercher plus loin, consultez toujours `~/.ollama/logs/server.log` ou `journalctl -u ollama -f` qui contiennent presque toujours la cause racine.

