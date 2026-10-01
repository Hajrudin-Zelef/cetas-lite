---
id: collect-261001-ia-llm/ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes-5
title: "Vérifier la version du pilote NVIDIA"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["nvidia", "agent", "attention", "embeddings", "flash attention", "gpu", "llama"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes.md
source_anchor: ""
source_lines: [382, 459]
sha256: 8c9f2074b6e2af98d15fe276a0b93cbe2b68d9b1052ecd322d1dd82c2ca6ccc1
---

# Vérifier la version du pilote NVIDIA

```
# Installation
pip install langchain==0.3.7 langchain-ollama==0.2.0 langchain-chroma==0.1.4
# rag_langchain.py — RAG complet en 30 lignes
from langchain_ollama import ChatOllama, OllamaEmbeddings
from langchain_chroma import Chroma
from langchain.text_splitter import RecursiveCharacterTextSplitter
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.runnables import RunnablePassthrough
from langchain_core.output_parsers import StrOutputParser
# 1. Charger et découper le document
with open("rgpd.txt", "r", encoding="utf-8") as f:
    text = f.read()
splitter = RecursiveCharacterTextSplitter(chunk_size=800, chunk_overlap=100)
chunks = splitter.split_text(text)
# 2. Embeddings et vector store
embeddings = OllamaEmbeddings(model="nomic-embed-text")
vectorstore = Chroma.from_texts(chunks, embeddings, persist_directory="./db")
retriever = vectorstore.as_retriever(search_kwargs={"k": 4})
# 3. LLM et prompt
llm = ChatOllama(model="llama3.1:8b", temperature=0.2, num_ctx=8192)
prompt = ChatPromptTemplate.from_template(
    """Tu es un assistant juridique français.
Réponds uniquement avec le contexte fourni.
Contexte :
{context}
Question : {question}
Réponse :"""
)
# 4. Chaîne RAG
chain = (
    {"context": retriever, "question": RunnablePassthrough()}
    | prompt
    | llm
    | StrOutputParser()
)
# 5. Question
print(chain.invoke("Que dit l article 17 sur le droit à l oubli ?"))
```
LlamaIndex propose une approche similaire mais orientée « moteur de recherche augmenté ». Pour un projet d’analyse documentaire massive (par exemple indexer 10 000 PDF de jurisprudence), LlamaIndex offre des structures d’index hiérarchiques plus efficaces que LangChain. Pour un chatbot conversationnel ou un agent multi-étapes, LangChain garde l’avantage. Les deux frameworks supportent Ollama nativement via leurs packages dédiés `llama-index-llms-ollama` et `langchain-ollama`.

## Étape 12 : Optimiser les performances (Flash Attention, GPU, threading)

Une fois Ollama opérationnel, plusieurs leviers permettent d’extraire 20 à 40 % de performance supplémentaire. Le premier est **Flash Attention**, activé par défaut depuis v0.6.2 sur GPU compatibles, mais qui peut être forcé via `OLLAMA_FLASH_ATTENTION=1`. Le second est le KV-cache quantifié (`OLLAMA_KV_CACHE_TYPE=q8_0`), qui réduit de moitié la VRAM consommée par le contexte, libérant de la place pour des fenêtres plus larges.

```
# Variables d'optimisation à définir avant ollama serve
export OLLAMA_FLASH_ATTENTION=1
export OLLAMA_KV_CACHE_TYPE=q8_0
export OLLAMA_NUM_PARALLEL=4
export OLLAMA_KEEP_ALIVE=24h
export OLLAMA_CONTEXT_LENGTH=8192
export OLLAMA_MAX_LOADED_MODELS=2
# Sur Linux avec systemd
sudo systemctl edit ollama
# Ajouter dans la section [Service] :
# Environment="OLLAMA_FLASH_ATTENTION=1"
# Environment="OLLAMA_KV_CACHE_TYPE=q8_0"
sudo systemctl daemon-reload
sudo systemctl restart ollama
# Vérifier l utilisation GPU pendant inférence
watch -n 1 nvidia-smi
# Mesurer le débit réel
time curl http://localhost:11434/api/generate -d '{
  "model": "llama3.1:8b",
  "prompt": "Compte de 1 à 100",
  "stream": false,
  "options": {"num_predict": 500}
}' | jq '.eval_count / (.eval_duration / 1000000000)'
```
Pour un usage CPU-only (machines sans GPU), la commande `num_thread` dans les options de requête, ou la variable `OLLAMA_NUM_THREADS`, contrôle le parallélisme. La règle empirique : utilisez le nombre de cœurs physiques (pas logiques avec hyperthreading) moins 1, pour laisser un cœur au système. Sur un Ryzen 9 7950X (16 cœurs physiques), `OLLAMA_NUM_THREADS=15` est optimal. Au-delà, le surcoût de synchronisation annule les gains.

**Piège fréquent #5** : sur les machines hybrides (CPU + GPU faible), Ollama peut décider de répartir le modèle entre VRAM et RAM (mode « partial offload »), entraînant des performances catastrophiques. Forcez tout le modèle sur GPU avec `"options": {"num_gpu": -1}` (toutes les couches), ou choisissez explicitement le nombre de couches GPU avec `"num_gpu": 32` pour les 32 premières couches. Si vous voyez « offloading X/33 layers » dans les logs et que X < nombre total, soit votre VRAM est insuffisante, soit votre quantification est trop lourde.

## Étape 13 : Projet complet — Assistant RAG pour PME française

Pour boucler ce tutoriel, voici un projet complet qui assemble toutes les briques précédentes : un assistant interne pour PME française qui répond aux questions des employés sur les procédures RH internes, en s’appuyant uniquement sur les documents PDF stockés localement. Le projet utilise FastAPI pour l’interface HTTP, Ollama pour le LLM et les embeddings, ChromaDB pour le vector store, et PyPDF pour l’extraction de texte. L’ensemble tourne sur une machine avec 16 Go RAM et un GPU RTX 4060 (8 Go VRAM).

