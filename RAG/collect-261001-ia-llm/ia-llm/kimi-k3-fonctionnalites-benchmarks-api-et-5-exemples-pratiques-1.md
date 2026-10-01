---
id: collect-261001-ia-llm/ia-llm/kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques-1
title: "kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Moonshot", "OpenAI"]
dates: []
keywords: ["benchmarks", "kimi", "agent", "claude", "cost", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques.md
source_anchor: ""
source_lines: [1, 110]
sha256: ee9fbdf6cf067b54841635ee038783e2dd16cfee7be834b94738f66f010eb7d1
---

# kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques

Cursus

La course aux modèles ouverts a de nouveau bougé le 16 juillet 2026, lorsque Moonshot AI a lancé Kimi K3, un modèle de 2,8 billions de paramètres, doté d'une fenêtre de contexte d'un million de jetons et de la vision native. C'est le plus grand modèle ouvert publié par Moonshot, largement au-delà de Kimi K2 en taille, et le premier qu'ils décrivent comme relevant de la classe des 3 billions de paramètres.

Si vous cherchez l'histoire du lancement, l'analyse d'architecture, les graphiques de benchmarks, les comparaisons avec Claude, GPT et les autres laboratoires chinois, ainsi que la liste des limites de Moonshot, notre article de blog sur Kimi K3 couvre tout cela. Ce tutoriel est le pendant pratique : comment y accéder et comment il se comporte à l'usage. Je passe en revue cinq petits exemples, quatre via l'API où je montre l'usage réel des jetons et le coût, et deux dans l'application web kimi.com. Ensemble, ils montrent comment K3 gère :

- L'appel d'outils et le retour d'un JSON strict
- Le chargement à la volée d'une définition d'outil
- La réduction du coût des longs contextes grâce au cache automatique
- La lecture d'une capture d'écran et la correction de la mise en page
- La création d'un tableau de bord interactif à partir d'un seul prompt

Les quatre exemples API ont été exécutés le 17 juillet 2026 avec le modèle `kimi-k3` et ont coûté environ 11 cents à froid, ou quelques cents après l'activation du cache.

## Comment accéder à Kimi K3

Le moyen le plus rapide d'essayer le modèle est kimi.com, où l'application web et les applications mobiles utilisent Kimi K3 pour des tâches d'agent générales sans configuration.

Pour des travaux plus lourds comme des rapports et des tableaux de bord, il y a Kimi Work, une application de bureau.

Si vous vivez dans le terminal, Kimi Code est un agent de codage à installer via npm sous `@moonshot-ai/kimi-code`, et vous choisissez le modèle avec la commande `/model`. L'usage de K3 dans Kimi Code nécessite un abonnement payant, et la fenêtre d'un million de jetons exige un palier supérieur.

Ce tutoriel se concentre sur l'API brute et l'application web, mais l'agent en ligne de commande est disponible si vous en avez besoin.

K3 ne remplace toutefois pas ses aînés. Le tableau ci-dessous montre comment se répartit la gamme actuelle.

| **Modèle** | **Fenêtre de contexte** | **Idéal pour** | 
| `kimi-k3` | 1 48 576 jetons | Travaux phares : long codage, vision, connaissances | 
| `kimi-k2.7-code` | 262 144 jetons | Codage dédié, avec une option haute vitesse plus rapide | 
| `kimi-k2.6` | 262 144 jetons | Conversation générale texte, image et vidéo | 

En bref, K3 est le modèle à privilégier quand une tâche mêle code, outils, documents et images, ou lorsque vous avez réellement besoin de la fenêtre d'un million de jetons. Pour la génération de code pure où la vitesse prime sur le contexte, `kimi-k2.7-code` reste le choix le plus judicieux : ne supposez pas que le dernier modèle est toujours le meilleur pour votre cas.

## Configuration de l'API Kimi K3

L'API est compatible avec le SDK OpenAI, donc si vous l'avez déjà utilisé, presque rien de ce qui suit ne vous surprendra. Vous avez besoin de Python 3.9 ou plus et d'une clé API.

### Étape 1 : générer une clé API

Commencez par vous connecter à la plateforme Kimi et ouvrez la page API Keys dans la console. Créez une clé, copiez-la une fois et stockez-la en lieu sûr, car vous ne la reverrez pas. Prévoyez aussi un petit crédit sur le compte pour passer des appels ; pour tout ce tutoriel, quelques dollars suffisent largement.

*Création d'une clé API Kimi K3. Image de l'auteur.*

### Étape 2 : installer le SDK

Installez ensuite le SDK OpenAI dans votre environnement. Une seule commande suffit.

`python -m pip install --upgrade "openai>=1.0"`
Cela télécharge la bibliothèque cliente utilisée par le reste des exemples, sans installation spécifique à Kimi.

### Étape 3 : stocker la clé et initialiser le client

Il vaut mieux lire la clé depuis une variable d'environnement que de la coller dans votre code. Définissez `MOONSHOT_API_KEY` dans votre shell ou un fichier `.env` , puis pointez le client vers l'URL de base de Moonshot.

```
import os
from openai import OpenAI
client = OpenAI(
    api_key=os.environ["MOONSHOT_API_KEY"],
    base_url="https://api.moonshot.ai/v1",
)
```
Les deux seules différences par rapport à une configuration OpenAI standard sont le `base_url` et le nom du modèle, à savoir `kimi-k3`. Avec cela en place, vous êtes prêt à faire un premier appel.

### Étape 4 : effectuer votre premier appel

Passons à une première requête. J'ai demandé au modèle de se présenter en une phrase, ce qui a donné un petit moment de franchise.

```
completion = client.chat.completions.create(
    model="kimi-k3",
    messages=[{"role": "user", "content": "Introduce Kimi K3 in one sentence."}],
    max_completion_tokens=800,
)
print(completion.choices[0].message.content)
```
La réponse fut un refus poli de deviner : le modèle a indiqué ne pas disposer d'informations fiables sur Kimi K3, ayant été entraîné avant sa propre sortie, et m'a renvoyé vers les annonces de Moonshot. Rappel utile : un modèle ne se connaît pas lui-même. L'appel API que je viens d'effectuer coûte environ sept dixièmes de cent. Notez le plafond `max_completion_tokens` que je fixe à chaque appel dans ce tutoriel pour éviter que des sorties verbeuses ne fassent grimper la facture.

*Premier retour de l'API Kimi K3. Image de l'auteur.*

## Exemple 1 : raisonnement en streaming et réponse finale

K3 raisonne systématiquement, et l'API renvoie ce raisonnement sur un canal séparé de la réponse. En mode flux, chaque fragment peut contenir `reasoning_content`, le `content` final, ou les deux, ce qui vous permet d'afficher la pensée et la réponse séparément.

```
stream = client.chat.completions.create(
    model="kimi-k3",
    messages=[{"role": "user", "content": "A bat and a ball cost $1.10 together. The bat costs $1.00 more than the ball. How much is the ball?"}],
    max_completion_tokens=1200,
    stream=True,
    stream_options={"include_usage": True},
)
for chunk in stream:
    if not chunk.choices:
        continue
    delta = chunk.choices[0].delta
    reasoning = getattr(delta, "reasoning_content", None)
    if reasoning:
        print(reasoning, end="", flush=True)
    if delta.content:
        print(delta.content, end="", flush=True)
```
Le modèle a d'abord diffusé son raisonnement : il a reconnu la question de la batte et de la balle comme le classique Cognitive Reflection Test, a signalé la réponse intuitive fausse de 0,10 $, puis a posé l'algèbre pour arriver à 0,05 $ et vérifier que 1,05 $ plus 0,05 $ font bien 1,10 $. La séparation est précieuse : dans une application réelle, vous montrez le `content` à vos utilisateurs et conservez le `reasoning_content` pour les journaux, car afficher le raisonnement brut en production est rarement souhaitable. Cet appel a utilisé 488 jetons de sortie et coûté moins d'un cent.

*Raisonnement en streaming puis réponse finale. Image de l'auteur.*

## Exemple 2 : appel d'outils avec sortie structurée

Kimi K3 est le modèle de la gamme qui prend en charge `tool_choice="required"`, ce qui force au moins un appel d'outil pendant un tour. Pratique lorsque vous voulez que le modèle récupère des données avant de répondre plutôt que de deviner. Ici, je lui ai fournis deux faux outils, une consultation de prix et un contrôle de stock, j'ai imposé un appel d'outil, exécuté les outils en local, puis demandé le résultat en JSON strict via `response_format`.

