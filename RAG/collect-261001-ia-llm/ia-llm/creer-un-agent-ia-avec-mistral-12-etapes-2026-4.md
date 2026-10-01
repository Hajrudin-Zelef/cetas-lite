---
id: collect-261001-ia-llm/ia-llm/creer-un-agent-ia-avec-mistral-12-etapes-2026-4
title: "Créer le dossier du projet"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["agent", "agents", "aws", "mcp", "mistral", "model context protocol"]
source: docs/RAG/collect-261001-ia-llm/creer-un-agent-ia-avec-mistral-12-etapes-2026.md
source_anchor: ""
source_lines: [344, 433]
sha256: 49ab01ddb7c1a49d2cee5d892a81f6eac8ce4f0221ec762cdf8c6f7845f0da4a
---

# Créer le dossier du projet

```
# main.py
from config import client, MODELE_ORCHESTRATEUR, MODELE_LEGER
from outils import SCHEMA_CONVERSION
def construire_agents():
    redacteur = client.beta.agents.create(
        model=MODELE_LEGER,
        name="Redacteur",
        description="Met en forme des résumés de veille.",
        instructions="Produis un titre, trois puces, une conclusion. En français.",
    )
    orchestrateur = client.beta.agents.create(
        model=MODELE_ORCHESTRATEUR,
        name="VeilleBot",
        description="Agent de veille tech francophone.",
        instructions=(
            "Analyste de veille tech. Recherche web pour les faits récents, "
            "conversion pour les devises. Délègue la mise en forme au rédacteur. "
            "Réponds en français et cite tes sources."
        ),
        tools=[{"type": "web_search"}, SCHEMA_CONVERSION],
        handoffs=[redacteur.id],
    )
    return orchestrateur
def boucle_chat(agent):
    conversation = None
    print("VeilleBot prêt. Tapez 'quitter' pour sortir.\n")
    while True:
        question = input("Vous : ").strip()
        if question.lower() in {"quitter", "exit", "q"}:
            break
        try:
            if conversation is None:
                conversation = client.beta.conversations.start(
                    agent_id=agent.id, inputs=question
                )
            else:
                conversation = client.beta.conversations.append(
                    conversation_id=conversation.conversation_id, inputs=question
                )
            for entree in conversation.outputs:
                if entree.type == "message.output":
                    print("VeilleBot :", entree.content, "\n")
        except Exception as e:
            print("Erreur :", e, "\n")
if __name__ == "__main__":
    agent = construire_agents()
    boucle_chat(agent)
```
Lancez-le avec `python main.py`. Vous disposez désormais d’un **agent IA** conversationnel complet : recherche web, fonction maison, mémoire d’état et délégation, le tout sur l’infrastructure européenne de Mistral. Le code source complet de référence du SDK est disponible sur le dépôt GitHub officiel mistralai.

## Pièges courants à éviter

La construction d’un **agent IA** réserve quelques chausse-trappes récurrentes. Voici les cinq erreurs les plus fréquentes et comment les éviter.

- **Coder la clé API en dur.** C’est la fuite de secret la plus banale et la plus coûteuse. Utilisez toujours un`.env` ignoré par Git. Un secret poussé sur GitHub doit être révoqué immédiatement, même dans un dépôt privé.
- **Confondre exécution et demande d’outil.** Le modèle ne lance jamais votre fonction. Il renvoie un`tool_call` . Si vous oubliez d’exécuter la fonction et de renvoyer le résultat, l’agent reste bloqué en attente et produit une réponse incohérente.
- **Renvoyer l’historique manuellement avec l’API conversations.** Contrairement à la complétion, l’API Agents gère l’état. Réinjecter tout l’historique à chaque tour double votre consommation de tokens et désynchronise le contexte.
- **Choisir un modèle surdimensionné pour chaque tâche.** Utiliser`mistral-medium-latest` pour reformuler une phrase est un gaspillage. Routez les tâches simples vers`mistral-small-latest` .
- **Ignorer les limites de débit (rate limits).** En boucle agentique, un agent peut générer des dizaines d’appels par requête. Sans gestion des erreurs 429, votre application plantera en charge. Implémentez un back-off exponentiel.

## Dépannage : les erreurs les plus fréquentes

Ce tableau de dépannage couvre les huit problèmes que vous rencontrerez le plus probablement en développant votre **agent IA**, avec leur cause et leur résolution.

| Symptôme | Cause probable | Résolution | 
|---|---|---|
| `401 Unauthorized` | Clé API invalide ou absente | Vérifier le `.env` et le rechargement via dotenv | 
| `429 Too Many Requests` | Limite de débit dépassée | Ajouter un back-off exponentiel, réduire la concurrence | 
| `422 Unprocessable Entity` | Schéma de fonction mal formé | Vérifier le JSON Schema (types, `required` ) | 
| `ModuleNotFoundError: mistralai` | SDK non installé dans le venv | Activer le venv puis `pip install mistralai` | 
| Agent ignore la recherche web | Outil non déclaré ou instructions vagues | Ajouter `web_search` et préciser quand l’utiliser | 
| Réponse en anglais | Instructions sans consigne de langue | Forcer « réponds en français » dans les instructions | 
| `tool_call_id` manquant | Résultat d’outil renvoyé sans identifiant | Toujours propager l’ `id` du`tool_call` | 
| Timeout sur la fonction maison | API externe lente ou injoignable | Définir un `timeout` et gérer l’exception réseau | 

## Conseils avancés pour passer en production

Une fois votre **agent IA** fonctionnel en local, plusieurs optimisations le rendent fiable à l’échelle. D’abord, le **streaming** : utilisez les variantes *stream* de l’API pour afficher la réponse au fil de l’eau, ce qui améliore drastiquement la latence perçue côté utilisateur. Ensuite, l’**observabilité** : journalisez chaque appel d’outil, sa durée et son coût en tokens. Un agent opaque est impossible à déboguer en production ; instrumentez-le dès le départ, par exemple avec une stack d’observabilité comme celle décrite dans notre tutoriel OpenTelemetry.

Côté robustesse, ajoutez systématiquement un **garde-fou sur le nombre d’itérations** : une boucle agentique peut, en cas d’instruction ambiguë, enchaîner indéfiniment des appels d’outils et faire exploser la facture. Fixez un plafond (par exemple dix appels par requête) et un budget de tokens maximal. Pensez aussi à la **bibliothèque de documents** de Mistral pour ancrer l’agent sur vos propres données internes via le RAG hébergé : c’est souvent plus pertinent que la recherche web pour un usage métier. Enfin, exploitez les **outils MCP** (Model Context Protocol) pour brancher vos systèmes existants – CRM, base de connaissances, ticketing – sans réécrire de connecteur propriétaire ; l’interopérabilité de ce type d’outillage a d’ailleurs franchi une étape le 6 août 2026, quand OpenAI, AWS, Microsoft, GitHub, Cursor et Vercel ont conjointement publié le standard *Agent Plugins 1.0.0* pour uniformiser la façon dont les agents découvrent et appellent des outils tiers.

### Sécurité, RGPD et bonnes pratiques européennes

Un avantage décisif de construire votre **agent IA** sur Mistral plutôt que sur un fournisseur américain est le traitement des données dans le cadre européen. Pour rester conforme au RGPD, minimisez les données personnelles transmises à l’agent, anonymisez ce qui peut l’être, et documentez les traitements. Évitez d’envoyer des secrets ou des identifiants dans les messages : l’agent n’en a pas besoin pour raisonner. Côté chaîne d’approvisionnement logicielle, épinglez vos versions de dépendances et auditez régulièrement votre `requirements.txt` – les attaques de la supply chain visant les paquets Python sont en hausse.

## Coûts et enjeu de souveraineté

La maîtrise des coûts d’un **agent IA** repose sur trois leviers. Le premier est le **routage des modèles** : réserver le modèle frontière à l’orchestration et au raisonnement, et déléguer le reste au modèle léger. Le deuxième est la **limitation des appels d’outils**, car chaque recherche web et chaque exécution de code consomme des ressources. Le troisième est la **mise en cache** des réponses pour les questions récurrentes. Mistral facture à l’usage, par million de tokens, avec des tarifs distincts selon le modèle ; consultez la page officielle des modèles pour les montants à jour, que nous ne reproduisons pas ici afin d’éviter toute donnée périmée.

