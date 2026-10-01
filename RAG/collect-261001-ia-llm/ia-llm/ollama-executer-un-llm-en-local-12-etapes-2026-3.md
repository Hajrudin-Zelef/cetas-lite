---
id: collect-261001-ia-llm/ia-llm/ollama-executer-un-llm-en-local-12-etapes-2026-3
title: "macOS (via Homebrew)"
domain: ia-llm
role: reference
task: reference
actors: ["Mistral", "OpenAI"]
dates: []
keywords: ["agent", "agents", "embeddings", "lora", "mistral"]
source: docs/RAG/collect-261001-ia-llm/ollama-executer-un-llm-en-local-12-etapes-2026.md
source_anchor: ""
source_lines: [152, 281]
sha256: 55e84aa09823beaf414a6c8306008c15555707fe1ea19231c2c10cc3984f04de
---

# macOS (via Homebrew)

Au-delà du chat interactif, la vraie puissance d’Ollama réside dans son **API REST locale**, exposée par défaut sur `http://localhost:11434`. Toute application – script, service, outil no-code – peut y envoyer des requêtes HTTP. Deux points d’entrée dominent : `/api/generate` pour une complétion simple et `/api/chat` pour un dialogue avec historique.

```
curl http://localhost:11434/api/generate -d '{
  "model": "llama3.2",
  "prompt": "Pourquoi exécuter un LLM en local ? Réponds en une phrase.",
  "stream": false
}'
```
Le point d’entrée `/api/chat` accepte un tableau `messages` avec des rôles (`system`, `user`, `assistant`), exactement comme les API conversationnelles classiques. C’est celui à privilégier pour bâtir un assistant qui conserve le fil de la discussion.

```
curl http://localhost:11434/api/chat -d '{
  "model": "llama3.2",
  "messages": [
    { "role": "system", "content": "Tu réponds toujours en français." },
    { "role": "user", "content": "Explique le RGPD en une phrase." }
  ],
  "stream": false
}'
```
Le paramètre `stream` est essentiel. À `true` (valeur par défaut), Ollama renvoie une suite d’objets JSON token par token, idéal pour afficher la réponse au fil de l’eau. À `false`, vous recevez un unique objet complet, plus simple à parser dans un script. Les autres points d’entrée incluent `/api/tags` (lister les modèles), `/api/show` (détails d’un modèle), `/api/ps` (modèles chargés) et `/api/embed` (embeddings). La référence complète figure dans la documentation officielle de l’API.

## Étape 6 – Utiliser l’API compatible OpenAI

Voici la fonctionnalité qui change tout pour les développeurs : Ollama expose une **API compatible OpenAI** sous le préfixe `/v1`. Concrètement, n’importe quel code écrit pour la bibliothèque `openai` fonctionne avec Ollama en changeant uniquement l’URL de base. Vos applications existantes basculent vers l’IA locale sans réécriture.

```
from openai import OpenAI
client = OpenAI(
    base_url="http://localhost:11434/v1",
    api_key="ollama",  # valeur ignorée, mais requise par la bibliothèque
)
reponse = client.chat.completions.create(
    model="llama3.2",
    messages=[{"role": "user", "content": "Bonjour, qui es-tu ?"}],
)
print(reponse.choices[0].message.content)
```
Le point d’entrée `/v1/chat/completions` reproduit fidèlement le format d’OpenAI, y compris le *streaming* et les sorties JSON. Un point d’entrée `/v1/embeddings` existe également. Cette compatibilité explique pourquoi tant d’outils – LangChain, LlamaIndex, ou des extensions d’éditeur – fonctionnent immédiatement avec Ollama : ils croient parler à OpenAI.

L’astuce est précieuse pour migrer en douceur. Vous développez en local gratuitement, puis basculez vers une API distante en production en modifiant deux lignes. Pour aller plus loin dans la création d’agents, notre guide Créer un agent IA avec Mistral applique le même principe côté API.

## Étape 7 – Intégrer Ollama en Python

Pour un contrôle plus fin, Ollama publie une bibliothèque Python officielle (version 0.6.2) et une bibliothèque JavaScript (0.6.3). Installez la première avec pip ; elle encapsule l’API REST dans des fonctions idiomatiques.

`pip install ollama`
L’appel le plus courant est `ollama.chat()`. Il prend un nom de modèle et une liste de messages, et renvoie un dictionnaire dont le contenu se lit via `reponse["message"]["content"]`.

```
import ollama
reponse = ollama.chat(
    model="llama3.2",
    messages=[{"role": "user", "content": "Donne-moi 3 noms pour un chatbot RGPD."}],
)
print(reponse["message"]["content"])
```
Pour une expérience réactive, activez le *streaming* : la fonction devient un générateur que vous parcourez pour afficher chaque fragment dès qu’il arrive. C’est ce qui donne l’effet « machine à écrire » des interfaces de chat modernes.

```
import ollama
flux = ollama.chat(
    model="llama3.2",
    messages=[{"role": "user", "content": "Écris un haïku sur le code."}],
    stream=True,
)
for partie in flux:
    print(partie["message"]["content"], end="", flush=True)
print()
```
La bibliothèque expose aussi `ollama.generate()`, `ollama.embed()`, `ollama.list()` et `ollama.pull()`, ce qui permet de tout piloter depuis Python : télécharger un modèle, l’interroger, gérer la mémoire. Si vous débutez en Python pour le web, notre tutoriel Flask montre comment exposer ces appels derrière une petite API.

## Étape 8 – Créer un modèle personnalisé avec un Modelfile

Un **Modelfile** est à Ollama ce qu’un Dockerfile est à Docker : un fichier texte qui décrit comment construire un modèle dérivé. Vous partez d’un modèle de base, vous lui imposez une personnalité (via `SYSTEM`), vous réglez ses paramètres d’inférence, et vous obtenez un modèle réutilisable d’une seule commande.

Créez un fichier nommé `Modelfile` avec le contenu suivant. Ici, nous fabriquons « Léa », une assistante technique francophone à la température basse (réponses plus déterministes) et au contexte élargi.

```
FROM llama3.2
# Paramètres d'inférence
PARAMETER temperature 0.6
PARAMETER num_ctx 8192
PARAMETER top_p 0.9
# Personnalité de l'assistant
SYSTEM """
Tu es Léa, une assistante technique francophone.
Tu réponds toujours en français, de manière concise et précise.
Si tu n'es pas sûre, tu le dis plutôt que d'inventer.
"""
```
Construisez ensuite le modèle, puis exécutez-le comme n’importe quel autre. Léa conservera systématiquement sa consigne système, sans que vous ayez à la répéter à chaque requête.

```
ollama create lea -f Modelfile
# transferring model data... success
ollama run lea
# >>> Présente-toi.
# Bonjour, je suis Léa, votre assistante technique francophone.
```
Les directives les plus utiles d’un Modelfile sont `FROM` (modèle de base), `SYSTEM` (consigne permanente), `PARAMETER` (température, `num_ctx` pour la taille du contexte, `num_predict` pour la longueur maximale), `TEMPLATE` (format du prompt) et `ADAPTER` (pour appliquer un adaptateur LoRA finement entraîné). Vous pouvez ainsi packager un assistant métier – juridique, support, documentation interne – et le distribuer à toute une équipe.

## Étape 9 – Générer des sorties structurées en JSON

Les LLM sont bavards par nature, ce qui complique leur intégration dans du code. Depuis fin 2024, Ollama prend en charge les **sorties structurées** : vous fournissez un schéma JSON et le modèle est contraint de produire une réponse strictement conforme. Fini les expressions régulières fragiles pour extraire une donnée – vous recevez du JSON valide à coup sûr.

En Python, le plus élégant consiste à décrire la structure avec Pydantic, puis à passer son schéma au paramètre `format`. Le modèle remplit alors les champs attendus.

```
from pydantic import BaseModel
import ollama
class Produit(BaseModel):
    nom: str
    prix: float
    en_stock: bool
reponse = ollama.chat(
    model="llama3.2",
    messages=[{"role": "user",
               "content": "Décris un clavier mécanique à 89,99 euros, en stock."}],
    format=Produit.model_json_schema(),
)
produit = Produit.model_validate_json(reponse["message"]["content"])
print(produit)
# nom='Clavier mécanique' prix=89.99 en_stock=True
```
Cette technique transforme un LLM local en moteur d’extraction fiable : analyse de factures, classification de tickets, structuration de données issues de texte libre. Combinée à la confidentialité du traitement local, elle est particulièrement adaptée aux données personnelles soumises au RGPD, qui ne doivent pas transiter par un service tiers.

Via l’API REST, le même résultat s’obtient en passant le schéma JSON dans le champ `format` de la requête. Choisissez un modèle suffisamment capable (8B et plus) : les très petits modèles respectent moins fidèlement les schémas complexes.

