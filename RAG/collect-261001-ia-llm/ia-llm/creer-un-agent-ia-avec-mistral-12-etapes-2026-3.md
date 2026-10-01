---
id: collect-261001-ia-llm/ia-llm/creer-un-agent-ia-avec-mistral-12-etapes-2026-3
title: "Créer le dossier du projet"
domain: ia-llm
role: reference
task: reference
actors: ["Mistral"]
dates: []
keywords: ["agent", "agents", "mai", "mistral", "parameters"]
source: docs/RAG/collect-261001-ia-llm/creer-un-agent-ia-avec-mistral-12-etapes-2026.md
source_anchor: ""
source_lines: [186, 343]
sha256: 45c908360b6c9ee7b7af64617df5fec4a6490e68ae548bee5ef505042dfecf8a
---

# Créer le dossier du projet

```
Mistral AI a lancé son API Agents le 27 mai 2025, permettant de
construire des agents autonomes avec recherche web, exécution de
code et bibliothèque de documents. L'entreprise a par ailleurs
levé des fonds importants en 2025 avec une participation d'ASML.
Sources : mistral.ai/news/agents-api ; couverture presse 2025.
```
### Étape 7 – Poursuivre la conversation (mémoire d’état)

Pour poser une question de suivi sans répéter le contexte, on enchaîne avec `client.beta.conversations.append()` en passant l’identifiant de la conversation. L’agent se souvient de ce qui précède :

```
# suite de lancer.py
suite = client.beta.conversations.append(
    conversation_id=conversation.conversation_id,
    inputs="Et quel est l'angle de souveraineté européenne là-dedans ?",
)
for entree in suite.outputs:
    if entree.type == "message.output":
        print(entree.content)
```
Parce que l’état est conservé côté Mistral, l’agent comprend que « là-dedans » renvoie à Mistral AI et à l’API Agents évoqués au tour précédent. C’est exactement ce qui distingue un **agent IA** d’un appel de complétion isolé, où vous devriez renvoyer manuellement tout l’historique à chaque requête.

## Étapes 8 à 10 : function calling et connecteur maison

### Étape 8 – Déclarer une fonction métier

Les outils intégrés ne suffisent jamais à un vrai cas d’usage : il faut brancher votre logique métier. C’est le rôle du *function calling* – le même principe qui permet par exemple à l’*Agent Wallet* de MetaMask, ouvert en accès anticipé le 8 juin 2026, de laisser des agents IA passer des ordres on-chain via des fonctions dédiées. Nous ajoutons une fonction `convertir_devise()` qui appelle une API publique de taux de change. Créez `outils.py` :

```
# outils.py
import httpx
def convertir_devise(montant: float, source: str, cible: str) -> dict:
    """Convertit un montant d'une devise vers une autre via une API publique."""
    url = f"https://api.frankfurter.app/latest?amount={montant}&from={source}&to={cible}"
    reponse = httpx.get(url, timeout=10)
    reponse.raise_for_status()
    donnees = reponse.json()
    resultat = donnees["rates"].get(cible)
    return {"montant_converti": resultat, "devise": cible}
# Schéma JSON décrivant la fonction pour le modèle
SCHEMA_CONVERSION = {
    "type": "function",
    "function": {
        "name": "convertir_devise",
        "description": "Convertit un montant d'une devise vers une autre.",
        "parameters": {
            "type": "object",
            "properties": {
                "montant": {"type": "number", "description": "Montant à convertir"},
                "source": {"type": "string", "description": "Devise source, ex : USD"},
                "cible": {"type": "string", "description": "Devise cible, ex : EUR"},
            },
            "required": ["montant", "source", "cible"],
        },
    },
}
```
### Étape 9 – Orchestrer l’appel de fonction avec le modèle

Le modèle n’*exécute* pas votre fonction : il vous indique qu’il veut l’appeler, avec quels arguments. C’est à votre code de l’exécuter puis de renvoyer le résultat. Cette boucle est le cœur du function calling. Créez `conversion.py` :

```
# conversion.py
import json
from config import client, MODELE_ORCHESTRATEUR
from outils import convertir_devise, SCHEMA_CONVERSION
messages = [
    {"role": "user", "content": "Convertis 13700000000 USD en EUR."}
]
# 1er appel : le modèle demande l'outil
reponse = client.chat.complete(
    model=MODELE_ORCHESTRATEUR,
    messages=messages,
    tools=[SCHEMA_CONVERSION],
    tool_choice="auto",
)
message = reponse.choices[0].message
messages.append(message)
# Si le modèle veut appeler la fonction
if message.tool_calls:
    appel = message.tool_calls[0]
    args = json.loads(appel.function.arguments)
    resultat = convertir_devise(**args)
    # Renvoyer le résultat au modèle
    messages.append({
        "role": "tool",
        "name": "convertir_devise",
        "content": json.dumps(resultat),
        "tool_call_id": appel.id,
    })
    # 2e appel : le modèle rédige la réponse finale
    finale = client.chat.complete(model=MODELE_ORCHESTRATEUR, messages=messages)
    print(finale.choices[0].message.content)
```
Sortie attendue (le montant exact dépend du taux du jour) :

```
13,7 milliards de dollars correspondent à environ 12,6 milliards
d'euros au taux de change actuel.
```
### Étape 10 – Connecter la fonction directement à l’agent

Plutôt que de gérer la boucle à la main, on peut déclarer la fonction comme outil de l’agent lui-même, aux côtés de la recherche web. L’agent décidera alors d’enchaîner web + conversion automatiquement. On met à jour `agent.py` :

```
# agent.py (version enrichie)
from config import client, MODELE_ORCHESTRATEUR
from outils import SCHEMA_CONVERSION
veille_bot = client.beta.agents.create(
    model=MODELE_ORCHESTRATEUR,
    name="VeilleBot",
    description="Agent de veille tech avec conversion de devises.",
    instructions=(
        "Tu es un analyste de veille tech francophone. Utilise la "
        "recherche web pour les faits récents et la fonction de "
        "conversion pour tout montant en devise étrangère. "
        "Réponds en français, cite tes sources."
    ),
    tools=[
        {"type": "web_search"},
        SCHEMA_CONVERSION,
    ],
)
```
Quand l’agent rencontre une demande mêlant veille et chiffres en devise – par exemple analyser la valorisation de Mistral et la convertir en euros – il orchestre désormais seul les deux outils. C’est la promesse de l’API Agents : vous décrivez les capacités, le modèle décide de l’enchaînement. Pour aller plus loin sur la mécanique de déclaration de fonctions, la documentation function calling de Mistral fait référence.

## Étapes 11 et 12 : handoffs et projet complet

### Étape 11 – Déléguer à un sous-agent (handoff)

Un **agent IA** mûr ne fait pas tout lui-même : il délègue. Le mécanisme de *handoff* permet à un agent de passer la main à un autre agent spécialisé. Ici, notre orchestrateur délègue la mise en forme finale à un agent « rédacteur » optimisé pour produire des résumés clairs avec le modèle léger. Créez `redacteur.py` :

```
# redacteur.py
from config import client, MODELE_LEGER
redacteur = client.beta.agents.create(
    model=MODELE_LEGER,
    name="Redacteur",
    description="Met en forme des résumés de veille clairs et structurés.",
    instructions=(
        "Tu reçois des notes brutes et tu produis un résumé en "
        "français : un titre, trois puces, une phrase de conclusion. "
        "Style sobre, pas de superlatifs."
    ),
)
# On déclare le handoff depuis l'orchestrateur vers le rédacteur
client.beta.agents.update(
    agent_id="AGENT_ORCHESTRATEUR_ID",   # id de VeilleBot
    handoffs=[redacteur.id],
)
```
Une fois le handoff déclaré, l’orchestrateur peut transférer la conversation au rédacteur lorsqu’il juge que la phase de recherche est terminée et qu’il faut produire la synthèse. L’utilisateur ne voit qu’une réponse fluide ; en coulisses, deux agents ont collaboré. Cette séparation des responsabilités améliore la qualité et réduit les coûts, car la mise en forme tourne sur `mistral-small-latest` plutôt que sur le modèle frontière.

### Étape 12 – Le projet complet, prêt à l’emploi

Voici le script d’entrée `main.py` qui assemble tout : création des agents, boucle de conversation interactive en ligne de commande, et gestion d’erreurs. C’est le livrable final, exécutable tel quel.

