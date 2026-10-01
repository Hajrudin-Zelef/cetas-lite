---
id: collect-261001-ia-llm/ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes-3
title: "Vérifier la version du pilote NVIDIA"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Mistral", "OpenAI"]
dates: []
keywords: ["attention", "embeddings", "gguf", "llama", "lora", "mai", "mistral", "packaging", "qwen"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes.md
source_anchor: ""
source_lines: [136, 253]
sha256: 78ab43cb82861b9b8ee44b38e75117fceaebdbe0e39ecae83b2fcedbf41ee4b4
---

# Vérifier la version du pilote NVIDIA

La variable `OLLAMA_HOST` mérite une attention particulière. Par défaut, le serveur n’écoute que sur `127.0.0.1`, ce qui empêche l’accès depuis d’autres machines du réseau. Pour exposer Ollama à votre LAN (par exemple depuis un laptop vers un serveur d’inférence sous Ubuntu), définissez `OLLAMA_HOST=0.0.0.0:11434`. Sur systemd, éditez `/etc/systemd/system/ollama.service.d/override.conf` avec `sudo systemctl edit ollama` et ajoutez `Environment="OLLAMA_HOST=0.0.0.0:11434"`. Attention : exposer Ollama publiquement sans authentification est une faille de sécurité majeure.

## Étape 5 : Interagir avec l’API HTTP native (/api/generate)

L’API HTTP est le cœur d’Ollama pour toute intégration sérieuse. Elle expose quatre endpoints principaux : `/api/generate` pour la complétion simple, `/api/chat` pour les conversations multi-tours avec historique, `/api/embed` pour générer des embeddings vectoriels, et `/api/tags` pour lister les modèles disponibles. Chaque endpoint accepte des requêtes POST JSON et peut renvoyer soit une réponse complète, soit un flux SSE (Server-Sent Events) si vous activez `"stream": true`.

```
# Complétion simple (non-streaming)
curl http://localhost:11434/api/generate -d '{
  "model": "llama3.1:8b",
  "prompt": "Explique le RGPD en 3 points",
  "stream": false,
  "options": {
    "temperature": 0.4,
    "num_ctx": 4096,
    "top_p": 0.9
  }
}'
# Réponse JSON typique :
# {
#   "model": "llama3.1:8b",
#   "created_at": "2026-04-29T10:23:45.123Z",
#   "response": "Le RGPD impose...",
#   "done": true,
#   "total_duration": 4823000000,
#   "load_duration": 12000000,
#   "prompt_eval_count": 12,
#   "eval_count": 142,
#   "eval_duration": 4500000000
# }
# Conversation multi-tours avec historique
curl http://localhost:11434/api/chat -d '{
  "model": "llama3.1:8b",
  "messages": [
    {"role": "system", "content": "Tu es un expert juridique français."},
    {"role": "user", "content": "Article 5 RGPD ?"}
  ],
  "stream": false
}'
# Lister les modèles
curl http://localhost:11434/api/tags | jq '.models[].name'
```
Les champs `total_duration`, `prompt_eval_count` et `eval_count` de la réponse sont précieux pour mesurer les performances réelles de votre déploiement. En divisant `eval_count` par `eval_duration` (en secondes), vous obtenez le débit en tokens/seconde. Le champ `load_duration` indique le temps de chargement du modèle en VRAM : il est non nul uniquement lors de la première requête après un déchargement. Pour un service production qui sert des milliers de requêtes, gardez le modèle chaud avec `"keep_alive": "24h"`.

**Piège fréquent #3** : par défaut, `num_ctx` est limité à 2048 tokens même si le modèle supporte 128k. Cette limite par défaut sert à économiser la VRAM, mais provoque une troncature silencieuse des prompts longs. Pour profiter de la pleine fenêtre de Llama 3.1 (128k), ajoutez explicitement `"options": {"num_ctx": 131072}` dans chaque requête, ou définissez `OLLAMA_CONTEXT_LENGTH=131072` avant `ollama serve`. Vérifiez que vous avez la VRAM nécessaire : un contexte de 128k peut consommer 30 à 50 Go supplémentaires.

## Étape 6 : Utiliser la compatibilité OpenAI (/v1/chat/completions)

L’une des features les plus utiles d’Ollama est sa couche de compatibilité OpenAI exposée sur `/v1`. Cela permet de remplacer un appel `OpenAI()` dans votre code Python ou JavaScript par un appel à Ollama local sans changer une seule ligne de logique métier. Le SDK officiel `openai` reconnaît le paramètre `base_url` qui redirige toutes les requêtes vers `http://localhost:11434/v1`. C’est l’intégration la plus rapide pour migrer un projet existant.

```
# Installation du SDK OpenAI (Python 3.11+)
pip install openai==1.54.3
# script openai_compat.py
from openai import OpenAI
client = OpenAI(
    base_url="http://localhost:11434/v1",
    api_key="ollama"  # Clé factice, requise mais ignorée
)
response = client.chat.completions.create(
    model="llama3.1:8b",
    messages=[
        {"role": "system", "content": "Réponds en français concis."},
        {"role": "user", "content": "Quelle est la capitale de la Belgique ?"}
    ],
    temperature=0.2,
    max_tokens=150
)
print(response.choices[0].message.content)
# Sortie : Bruxelles est la capitale de la Belgique.
# Mode streaming
stream = client.chat.completions.create(
    model="llama3.1:8b",
    messages=[{"role": "user", "content": "Compte de 1 à 5"}],
    stream=True
)
for chunk in stream:
    if chunk.choices[0].delta.content:
        print(chunk.choices[0].delta.content, end="", flush=True)
```
La compatibilité OpenAI couvre `/v1/chat/completions`, `/v1/completions`, `/v1/embeddings` et `/v1/models`. Quelques différences subsistent : Ollama ignore les paramètres `logit_bias`, `logprobs` et `n > 1` (génération multiple). Les `tool_calls` sont supportés depuis la v0.4 pour les modèles compatibles (Llama 3.1, Qwen 2.5+, Mistral Nemo), et la v0.24.0 (mai 2026) a ajouté la prise en charge de l’app Codex, élargissant l’intégration côté outils de codage assistés par IA. Pour le SDK JavaScript, le package officiel `ollama` (`npm install ollama`) offre une API typée TypeScript plus idiomatique que le détournement du SDK OpenAI.

## Étape 7 : Créer un Modelfile personnalisé

Le Modelfile est l’équivalent Ollama du Dockerfile : un fichier texte qui décrit comment construire un modèle personnalisé à partir d’un modèle de base. Les mots-clés principaux sont `FROM` (modèle parent), `PARAMETER` (température, top_p, num_ctx, etc.), `SYSTEM` (prompt système permanent), `TEMPLATE` (format de prompt Jinja), `MESSAGE` (exemples few-shot) et `ADAPTER` (chargement d’un LoRA). Cette mécanique transforme Ollama en plateforme de packaging réutilisable pour distribuer des assistants spécialisés.

```
# Fichier : Modelfile
FROM llama3.1:8b
# Paramètres d'inférence
PARAMETER temperature 0.3
PARAMETER top_p 0.85
PARAMETER num_ctx 8192
PARAMETER repeat_penalty 1.1
PARAMETER stop "<|eot_id|>"
# Prompt système permanent
SYSTEM """Tu es Juris-FR, un assistant juridique français spécialisé
en droit du numérique. Réponds toujours en français professionnel.
Cite l'article exact du Code applicable. Si tu n'es pas sûr,
indique-le clairement plutôt que d'inventer."""
# Exemples few-shot
MESSAGE user "Quel article RGPD pour le droit à l'oubli ?"
MESSAGE assistant "Article 17 du RGPD (Règlement UE 2016/679)."
# Construire le modèle
ollama create juris-fr -f Modelfile
# transferring model data
# creating layer
# writing manifest
# success
# Utiliser le modèle personnalisé
ollama run juris-fr "Que dit l'article 5 du RGPD ?"
```
L’avantage d’un Modelfile bien construit est de garantir un comportement reproductible entre environnements de développement et production. Vous pouvez versionner votre Modelfile dans Git, le distribuer à votre équipe, et chacun obtiendra le même modèle après `ollama create`. Pour les déploiements à plusieurs nœuds, vous pouvez aussi pousser votre modèle vers ollama.com (compte gratuit requis) avec `ollama push user/juris-fr`, puis le récupérer ailleurs avec `ollama pull user/juris-fr`. Cette mécanique facilite la collaboration interne sans exposer publiquement vos prompts système.

## Étape 8 : Choisir la bonne quantification (Q4_K_M, Q5_K_M, Q8_0)

La quantification est le processus qui réduit la précision des poids du modèle (par exemple de FP16 16-bits à 4-bits) pour économiser drastiquement la mémoire au prix d’une légère perte de qualité. Ollama distribue par défaut les modèles en Q4_K_M, considéré comme le meilleur compromis qualité/taille pour la majorité des usages. Comprendre les différentes quantifications GGUF est essentiel pour optimiser votre déploiement selon votre matériel.

