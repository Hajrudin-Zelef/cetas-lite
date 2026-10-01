---
id: collect-261001-ia-llm/ia-llm/creer-un-agent-ia-avec-mistral-12-etapes-2026-2
title: "Créer le dossier du projet"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Mistral", "Moonshot"]
dates: []
keywords: ["agent", "agents", "aws", "kimi", "mistral"]
source: docs/RAG/collect-261001-ia-llm/creer-un-agent-ia-avec-mistral-12-etapes-2026.md
source_anchor: ""
source_lines: [50, 185]
sha256: 382fb0001c50e2c0c6609fd531f68f08c1521fc765c18087ec7940b0779bb953
---

# Créer le dossier du projet

## Ce que nous allons construire : l’agent de veille tech

Le projet fil rouge est un **agent IA** de veille technologique baptisé `VeilleBot`, dans l’esprit des agents d’analyse autonome qui se multiplient sur le marché – à l’image de *TracisAI*, l’agent d’analyse de logs que l’éditeur BTM Inc. a lancé sur AWS Marketplace le 27 juin 2026. Concrètement, il devra : (1) recevoir une question d’un utilisateur, par exemple « Quelles sont les dernières annonces de Mistral cette semaine ? » ; (2) décider seul d’utiliser la recherche web pour trouver l’information ; (3) appeler une fonction maison `convertir_devise()` quand une valorisation ou un prix est mentionné dans une devise étrangère ; (4) mémoriser le fil de la conversation pour répondre à des questions de suivi ; et (5) déléguer à un sous-agent « rédacteur » la mise en forme d’un résumé final propre.

Ce périmètre couvre les quatre piliers d’un **agent IA** de production : outils intégrés, outils maison, état/mémoire et orchestration multi-agents. Le code est volontairement modulaire pour que vous puissiez remplacer la veille tech par votre propre cas d’usage (support client, analyse financière, automatisation interne). Tout l’intérêt d’un agent souverain est là : vos données et votre logique métier restent sous votre contrôle, sur une infrastructure européenne.

## Prérequis et versions

Avant de démarrer, assurez-vous de disposer de l’environnement suivant. Les versions indiquées sont celles testées pour ce tutoriel ; des versions plus récentes fonctionneront dans la plupart des cas.

| Composant | Version testée | Rôle | 
|---|---|---|
| Python | 3.11 ou supérieur | Langage du projet | 
| SDK `mistralai` | dernière version (pip) | Client officiel de La Plateforme | 
| `python-dotenv` | dernière version | Chargement de la clé API | 
| `httpx` | dernière version | Appel de notre API de devises | 
| Compte Mistral | – | Accès à console.mistral.ai | 
| Clé API Mistral | – | Authentification | 

Côté connaissances, une familiarité de base avec Python (fonctions, dictionnaires, gestion d’exceptions) suffit. Aucune expérience préalable en machine learning n’est requise : tout le « cerveau » est délégué aux modèles Mistral. Si vous débutez avec les LLM en local et la confidentialité RGPD, notre tutoriel Ollama est un complément utile pour comprendre l’alternative auto-hébergée – la startup française H Company a d’ailleurs lancé en juin 2026 sa famille de modèles d’agents « computer-use » locaux Holo 3.1, déclinée de 0,8 à 35 milliards de paramètres, qui illustre bien cette tendance à faire tourner l’agent chez soi. Dans la même veine des modèles ouverts, le chinois Moonshot a dévoilé mi-juillet 2026 son modèle de codage agentique Kimi K3, pesant environ 2 800 milliards de paramètres, dont les poids ouverts ont été promis pour le 27 juillet 2026.

## Étapes 1 à 3 : configuration de l’environnement

### Étape 1 – Créer la clé API sur La Plateforme

Rendez-vous sur `console.mistral.ai`, créez un compte (ou connectez-vous), puis ouvrez la section « API Keys ». Cliquez sur « Create new key », nommez-la `veille-bot` et copiez immédiatement la valeur affichée : elle ne sera plus jamais visible en clair. Mistral propose un palier d’expérimentation gratuit pour démarrer, ainsi que des paliers payants à l’usage ; consultez la page de tarification officielle pour les montants à jour, car les prix évoluent et nous ne reproduisons ici aucun chiffre non vérifié.

### Étape 2 – Initialiser le projet et l’environnement virtuel

```
# Créer le dossier du projet
mkdir veille-bot && cd veille-bot
# Créer et activer un environnement virtuel
python3 -m venv .venv
source .venv/bin/activate      # Sous Windows : .venv\Scripts\activate
# Installer les dépendances
pip install mistralai python-dotenv httpx
# Vérifier l'installation du SDK
python -c "import mistralai; print('SDK Mistral OK')"
```
Sortie attendue :

`SDK Mistral OK`
### Étape 3 – Sécuriser la clé API dans un fichier .env

Ne codez **jamais** votre clé en dur dans le code source. Créez un fichier `.env` à la racine du projet et ajoutez-le immédiatement à votre `.gitignore` :

```
# Fichier .env
MISTRAL_API_KEY=votre_cle_api_ici
# Ajouter au .gitignore
echo ".venv/" >> .gitignore
echo ".env" >> .gitignore
```
Créez ensuite `config.py`, le point d’entrée qui charge la clé et instancie le client. Ce client unique sera réutilisé dans tout le projet :

```
# config.py
import os
from dotenv import load_dotenv
from mistralai import Mistral
load_dotenv()
API_KEY = os.environ.get("MISTRAL_API_KEY")
if not API_KEY:
    raise RuntimeError("MISTRAL_API_KEY est manquante. Vérifiez votre fichier .env.")
# Client global réutilisable
client = Mistral(api_key=API_KEY)
# Modèles utilisés dans le projet
MODELE_ORCHESTRATEUR = "mistral-medium-latest"
MODELE_LEGER = "mistral-small-latest"
```
## Étapes 4 à 7 : premier agent et outils intégrés

### Étape 4 – Valider la connexion avec une complétion simple

Avant de construire l’agent, vérifions que la complétion de chat fonctionne. C’est le test de fumée indispensable : s’il échoue, inutile d’aller plus loin. Créez `test_connexion.py` :

```
# test_connexion.py
from config import client, MODELE_LEGER
reponse = client.chat.complete(
    model=MODELE_LEGER,
    messages=[
        {"role": "user", "content": "Réponds en un mot : capitale de la France ?"}
    ],
)
print(reponse.choices[0].message.content)
```
Sortie attendue :

`Paris`
### Étape 5 – Créer le premier agent avec l’API Agents

Nous passons maintenant de la complétion à l’**agent IA**. La création d’un agent se fait via `client.beta.agents.create()`. Un agent est un objet persistant, doté d’instructions et d’une liste d’outils. Créez `agent.py` :

```
# agent.py
from config import client, MODELE_ORCHESTRATEUR
veille_bot = client.beta.agents.create(
    model=MODELE_ORCHESTRATEUR,
    name="VeilleBot",
    description="Agent de veille technologique francophone.",
    instructions=(
        "Tu es un analyste de veille tech. Tu réponds en français, "
        "de manière factuelle et concise. Quand une information peut "
        "être périmée, utilise la recherche web avant de répondre. "
        "Cite toujours tes sources."
    ),
    tools=[{"type": "web_search"}],
)
print("Agent créé, id =", veille_bot.id)
```
Sortie attendue (l’identifiant variera) :

`Agent créé, id = ag_01h9z8k2m4p6q8r0s2t4v6w8x0`
Notez le paramètre `tools=[{"type": "web_search"}]` : c’est lui qui donne à notre **agent IA** la capacité d’aller chercher des informations fraîches. Sans cet outil, l’agent se contenterait de sa connaissance pré-entraînée, par nature figée à sa date de coupure.

### Étape 6 – Démarrer une conversation et déclencher la recherche web

L’interaction avec un agent passe par l’API *conversations*, qui gère l’état. On démarre avec `client.beta.conversations.start()`. Créez `lancer.py` :

```
# lancer.py
from config import client
from agent import veille_bot
conversation = client.beta.conversations.start(
    agent_id=veille_bot.id,
    inputs="Quelles annonces majeures Mistral AI a-t-il faites récemment ?",
)
# Parcourir les éléments de sortie de l'agent
for entree in conversation.outputs:
    if entree.type == "message.output":
        print(entree.content)
```
L’agent décide seul d’appeler `web_search`, récupère des résultats, puis synthétise une réponse sourcée. Exemple de sortie (abrégé) :

