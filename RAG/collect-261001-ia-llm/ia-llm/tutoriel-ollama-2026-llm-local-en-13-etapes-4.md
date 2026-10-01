---
id: collect-261001-ia-llm/ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes-4
title: "Vérifier la version du pilote NVIDIA"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["nvidia", "agents", "amd", "dpo", "embedding", "embeddings", "fine-tuning", "gpu", "llama", "mai"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes.md
source_anchor: ""
source_lines: [254, 381]
sha256: 015d25567a95c6d3cd75c711cea158b0f744c1558f99c7804c6b30f3efa578c7
---

# Vérifier la version du pilote NVIDIA

| Quantification | Bits effectifs | Taille (8B) | Perplexité vs FP16 | Recommandation | 
|---|---|---|---|---|
| Q2_K | ~2,6 | 3,0 Go | +15 % (dégradé) | Éviter sauf urgence VRAM | 
| Q3_K_M | ~3,4 | 3,8 Go | +5 % | Acceptable, 4 Go VRAM | 
| Q4_K_M | ~4,5 | 4,7 Go | +1,5 % | Sweet spot par défaut | 
| Q5_K_M | ~5,5 | 5,4 Go | +0,7 % | Si VRAM disponible | 
| Q6_K | ~6,5 | 6,1 Go | +0,2 % | Quasi-FP16 | 
| Q8_0 | ~8,5 | 8,0 Go | +0,05 % | Production critique | 
| F16 | 16 | 14,9 Go | Référence | Recherche, fine-tuning | 

En pratique pour un développeur français, la règle est simple : si votre modèle Q4_K_M tient en VRAM avec votre fenêtre de contexte habituelle, restez sur Q4_K_M. Si vous avez 30 à 50 % de VRAM libre, montez à Q5_K_M ou Q6_K pour un gain de qualité mesurable sur les tâches de raisonnement. Réservez Q8_0 aux cas où la qualité est critique (génération de contrats juridiques, code médical) et où vous pouvez vous permettre 70 % de RAM en plus. Pour télécharger une quantification spécifique :

```
# Tags spécifiques par quantification
ollama pull llama3.1:8b-instruct-q4_K_M    # par défaut
ollama pull llama3.1:8b-instruct-q5_K_M    # +15 % qualité
ollama pull llama3.1:8b-instruct-q8_0      # quasi-FP16
ollama pull llama3.1:8b-instruct-fp16      # référence
# Comparer les tailles
ollama list | grep llama3.1
# llama3.1:8b-instruct-q4_K_M    4.7 GB
# llama3.1:8b-instruct-q5_K_M    5.4 GB
# llama3.1:8b-instruct-q8_0      8.0 GB
# llama3.1:8b-instruct-fp16     14.9 GB
```
## Étape 9 : Générer des embeddings pour un RAG local

L’endpoint `/api/embed` introduit en v0.6.2 permet de générer des embeddings vectoriels en batch, brique fondamentale pour construire un système RAG (Retrieval-Augmented Generation) 100 % local. Les modèles d’embeddings recommandés en 2026 sont `nomic-embed-text` (768 dimensions, 137 Mo) pour les usages généraux et `mxbai-embed-large` (1024 dimensions, 670 Mo) pour la production. Pour le français, `jina-embeddings-v2-base-fr` donne d’excellents résultats sur les documents juridiques et techniques.

```
# Télécharger un modèle d'embeddings
ollama pull nomic-embed-text
# Générer un embedding (batch)
curl http://localhost:11434/api/embed -d '{
  "model": "nomic-embed-text",
  "input": [
    "Le RGPD est entré en vigueur le 25 mai 2018.",
    "La CNIL est l autorité de contrôle française.",
    "Le DPO est le délégué à la protection des données."
  ]
}'
# Réponse :
# {
#   "model": "nomic-embed-text",
#   "embeddings": [[0.123, -0.456, ...], [...], [...]],
#   "total_duration": 234000000,
#   "load_duration": 1200000,
#   "prompt_eval_count": 36
# }
# Avec Python + ChromaDB
pip install chromadb==0.5.13 ollama==0.4.2
# rag_setup.py
import ollama
import chromadb
client = chromadb.PersistentClient(path="./rag_db")
collection = client.get_or_create_collection("docs_juridiques")
docs = [
    "Article 5 RGPD : licéité, loyauté, transparence du traitement.",
    "Article 17 RGPD : droit à l effacement (droit à l oubli).",
    "Article 25 RGPD : protection des données dès la conception."
]
for i, doc in enumerate(docs):
    emb = ollama.embed(model="nomic-embed-text", input=doc)
    collection.add(
        ids=[f"doc_{i}"],
        embeddings=[emb["embeddings"][0]],
        documents=[doc]
    )
# Requête
query = "Comment supprimer mes données personnelles ?"
q_emb = ollama.embed(model="nomic-embed-text", input=query)
results = collection.query(query_embeddings=[q_emb["embeddings"][0]], n_results=2)
print(results["documents"])
```
Une fois les embeddings stockés dans ChromaDB (ou Qdrant, ou Milvus), le pattern RAG complet consiste à : (1) embeddifier la question utilisateur, (2) récupérer les *k* documents les plus proches en similarité cosinus, (3) injecter ces documents dans le prompt système d’un modèle de génération, (4) répondre à l’utilisateur. Cette architecture permet de répondre sur des documents privés sans les envoyer dans le cloud, conformité RGPD garantie par construction.

## Étape 10 : Déployer Ollama dans un conteneur Docker

Pour un déploiement production reproductible, containeriser Ollama est la voie royale. L’image officielle `ollama/ollama` est maintenue activement et supporte CPU, GPU NVIDIA (via nvidia-container-toolkit) et GPU AMD (via ROCm). Le volume persistant `/root/.ollama` stocke les modèles téléchargés et survit aux recréations du conteneur. Cette pratique évite de re-télécharger 40 Go à chaque redéploiement.

```
# docker-compose.yml — déploiement GPU NVIDIA
version: "3.9"
services:
  ollama:
    image: ollama/ollama:0.6.2
    container_name: ollama
    restart: unless-stopped
    ports:
      - "11434:11434"
    volumes:
      - ollama_data:/root/.ollama
    environment:
      - OLLAMA_HOST=0.0.0.0:11434
      - OLLAMA_KEEP_ALIVE=24h
      - OLLAMA_NUM_PARALLEL=4
      - OLLAMA_MAX_LOADED_MODELS=2
    deploy:
      resources:
        reservations:
          devices:
            - driver: nvidia
              count: all
              capabilities: [gpu]
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:11434/api/tags"]
      interval: 30s
      timeout: 10s
      retries: 3
volumes:
  ollama_data:
# Démarrer
docker compose up -d
# Vérifier que le GPU est détecté
docker compose exec ollama nvidia-smi
# Tirer un modèle dans le conteneur
docker compose exec ollama ollama pull llama3.1:8b
# Tester l API depuis l hôte
curl http://localhost:11434/api/tags
```
Les variables `OLLAMA_NUM_PARALLEL` et `OLLAMA_MAX_LOADED_MODELS` méritent d’être ajustées selon votre charge. `OLLAMA_NUM_PARALLEL=4` autorise 4 requêtes simultanées par modèle (utile pour servir une équipe), tandis que `OLLAMA_MAX_LOADED_MODELS=2` limite à 2 le nombre de modèles distincts en VRAM. Si vous servez à la fois un modèle de génération (Llama 3.1 8B) et un d’embeddings (nomic-embed-text), cette valeur doit être ≥ 2 sinon le système swappe en permanence entre les deux, écroulant les performances.

**Piège fréquent #4** : sur Docker Desktop Windows avec WSL2, le passthrough GPU exige des étapes supplémentaires : installer le pilote NVIDIA pour WSL côté Windows, vérifier `wsl --update`, et activer « Use NVIDIA GPU acceleration » dans les paramètres Docker Desktop. Sans ces étapes, Ollama tombe silencieusement sur le CPU et vous voyez des performances 10 à 20× plus lentes que prévu, sans message d’erreur explicite.

## Étape 11 : Intégrer Ollama avec LangChain et LlamaIndex

Les frameworks d’orchestration LLM comme LangChain et LlamaIndex offrent des intégrations Ollama de première classe. Ces librairies abstraient la mécanique des prompts, gèrent la mémoire conversationnelle, et permettent de chaîner plusieurs appels LLM (« agents », « chains »). Pour un développeur français qui construit un chatbot interne ou un assistant de support, c’est généralement le bon niveau d’abstraction : ni le bas niveau de l’API HTTP, ni la rigidité d’une plateforme SaaS.

