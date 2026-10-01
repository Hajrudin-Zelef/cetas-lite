---
id: collect-261001-ia-llm/ia-llm/kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques-2
title: "kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques"
domain: ia-llm
role: reference
task: reference
actors: ["Moonshot", "OpenAI", "United States"]
dates: []
keywords: ["kimi", "agent", "alignment", "decode", "parameters", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques.md
source_anchor: ""
source_lines: [111, 221]
sha256: b631280ac1cd6a67d60881f82e2401c4dee0440fa0cd99632fcdc4c2c77c5f7c
---

# kimi-k3-fonctionnalites-benchmarks-api-et-5-exemples-pratiques

```
first = client.chat.completions.create(
    model="kimi-k3",
    messages=messages,
    tools=TOOLS,
    tool_choice="required",
    max_completion_tokens=2500,
)
assistant_message = first.choices[0].message
messages.append(assistant_message)
for tool_call in assistant_message.tool_calls or []:
    args = json.loads(tool_call.function.arguments)
    messages.append({"role": "tool", "tool_call_id": tool_call.id, "content": run_tool(tool_call.function.name, args)})
```
Le modèle a appelé les deux outils avec le bon code produit, puis a renvoyé un récapitulatif de commande propre en JSON : cinq claviers mécaniques à 89 $ unitaires, un total de 445 $, et un indicateur de stock à true. Deux détails font la différence en pratique : vous devez réinsérer le message complet de l'assistant dans la conversation avant d'ajouter les résultats d'outil, et vous ne devez parser que le `content` pour le JSON, jamais le champ de raisonnement. La paire d'appels a coûté moins d'un cent au total.

Appels d'outils et sortie JSON structurée. Image de l'auteur.

## Exemple 3 : charger des outils dynamiquement

Si vous avez des dizaines d'outils, envoyer toutes leurs définitions à chaque requête gaspille des jetons et encombre le prompt. Kimi K3 vous permet d'injecter une définition d'outil en cours de conversation via un message `system` qui contient un champ `tools` et aucun `content`. L'outil devient disponible à partir de ce point, ce qui maintient les grands catalogues d'outils hors de votre préfixe mis en cache jusqu'à ce qu'un outil soit réellement nécessaire.

```
messages = [
    {"role": "user", "content": "Convert 100 US dollars to euros at a rate of 0.92."},
    {"role": "system", "tools": [{
        "type": "function",
        "function": {
            "name": "convert_currency",
            "description": "Convert an amount from one currency to another",
            "parameters": {
                "type": "object",
                "properties": {"amount": {"type": "number"}, "rate": {"type": "number"}},
                "required": ["amount", "rate"],
            },
        },
    }]},
]
completion = client.chat.completions.create(model="kimi-k3", messages=messages)
print(completion.choices[0].message.tool_calls)
```
K3 a pris en compte l'outil fraîchement chargé et a appelé `convert_currency` avec un montant de 100 et un taux de 0,92, comme prévu. Gardez en tête que le serveur ne conserve pas cette définition pour vous : renvoyez le message system lors des requêtes suivantes si vous voulez que l'outil reste disponible. C'était l'appel le moins cher de la série, à environ deux dixièmes de cent.


Appel d'un outil de conversion dynamique. Image de l'auteur.

## Exemple 4 : réduire le coût des longs contextes avec le cache

C'est ici que la fenêtre d'un million de jetons devient vraiment pratique. Le cache de contexte est automatique : pas d'ID de cache ni de durée de vie à gérer. Vous envoyez un long préfixe, vous le conservez strictement identique aux requêtes suivantes, et la partie répétée est facturée au tarif cache hit plutôt qu'au tarif cache miss. Pour rendre l'écart visible, j'ai utilisé une base de connaissances d'environ 33 000 jetons et posé une question dessus.

```
knowledge = Path("knowledge_base.md").read_text(encoding="utf-8")
completion = client.chat.completions.create(
    model="kimi-k3",
    messages=[
        {"role": "system", "content": knowledge},
        {"role": "user", "content": "What is the rated payload of the Atlas robot?"},
    ],
    max_completion_tokens=600,
)
```
La première fois, rien n'était en cache et la requête a coûté environ 9,9 cents pour quelque 33 000 jetons d'entrée. Une fois le préfixe vu, la même requête a touché le cache sur les 32 512 jetons de préfixe et coûté environ 1,1 cent, soit près d'un facteur neuf. La raison : l'écart de prix ; l'entrée en cache est facturée 0,30 $ par million de jetons contre 3,00 $ hors cache. Un point à noter : l'écriture en cache est asynchrone, donc le hit ne se voit pas sur un appel immédiatement consécutif. Il apparaît sur une requête ultérieure ; exécuter le script deux fois à une minute d'intervalle montre d'abord le miss, puis le hit.

Coût d'un cache miss vs cache hit. Image de l'auteur.

## Exemple 5 : détecter des bugs de mise en page sur une capture d'écran

La vision est native dans K3, et l'API offre un moyen simple de l'utiliser, même si elle n'accepte pas d'URL d'image publique. Vous envoyez l'image en data URL base64 et faites du `content` un tableau d'objets, une partie pour l'image, une pour le texte. J'ai rendu un petit tableau de bord avec quelques bugs de mise en page volontaires, enregistré une capture d'écran et demandé à K3 ce qui n'allait pas.

*Le tableau de bord avec des bugs de mise en page volontaires. Image de l'auteur.*

```
import base64
from pathlib import Path
image_data = base64.b64encode(Path("broken_dashboard.png").read_bytes()).decode()
completion = client.chat.completions.create(
    model="kimi-k3",
    messages=[{
        "role": "user",
        "content": [
            {"type": "image_url", "image_url": {"url": f"data:image/png;base64,{image_data}"}},
            {"type": "text", "text": "List the layout and alignment problems you can see, and give a short CSS fix for each."},
        ],
    }],
    max_completion_tokens=3500,
)
print(completion.choices[0].message.content)
```
K3 a bien lu l'image. Il a repéré la carte plus basse que la rangée et qui chevauche sa voisine, le badge posé sur un nombre (il a même mal lu le 3 910 masqué en 5 910, preuve du bug), l'espace irrégulier avant la dernière carte, la barre qui déborde vers la carte du dessus, ainsi que l'info-bulle posée sur les barres, et a proposé une correction CSS courte pour chacune, comme regrouper les cartes dans une grille. En revanche, il a omis le sous-titre à très faible contraste : la vision capte mieux ce qui saute aux yeux que les détails discrets. L'appel a coûté environ deux cents.

## Limites de Kimi K3

Les exemples API se sont bien passés, mais quelques aspérités méritent d'être signalées pour éviter les surprises. J'ai rencontré la plupart directement.

- 
Seul `reasoning_effort="max"` est disponible pour l'instant, vous ne pouvez donc pas encore réduire le raisonnement pour économiser.
- 
Les paramètres d'échantillonnage sont figés. Des valeurs comme `temperature` ,`top_p` et les pénalités sont verrouillées ; ne les joignez pas aux requêtes au lieu d'essayer de les ajuster.
- 
La sortie peut devenir longue et coûteuse. Limitez `max_completion_tokens` , comme dans les exemples, et validez toute boucle d'agent.
- 
Les URL d'image publiques ne sont pas prises en charge via l'API ; prévoyez le base64 ou des fichiers téléversés pour la vision.

Rien de tout cela n'est bloquant, mais ces points influencent la façon d'utiliser le modèle. Le coût de sortie est celui que je surveillerais en priorité.

## Conclusion

Au fil de mes essais, deux choses ressortent. L'appel d'outils et la sortie structurée ont fonctionné sans relance, et le cache a compté davantage que prévu : réutiliser le même long préfixe a rendu peu coûteuse la répétition d'une grosse requête. Pour l'analyse à l'échelle d'un dépôt, les appels répétés à long contexte ou l'ingénierie multimodale, K3 est un choix raisonnable par défaut ; pour un chat rapide et à bas coût ou un contrôle fin de l'échantillonnage, un modèle plus petit sera plus simple. Les détails sur les poids ouverts et la licence, que j'ai signalés plus haut, devraient être clarifiés après la sortie du 27 juillet.

Pour approfondir les schémas utilisés dans ces exemples, notre cours Developing AI Systems with the OpenAI API couvre le function calling et le raccordement des modèles à des outils externes en Python.

