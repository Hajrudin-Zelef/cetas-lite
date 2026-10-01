---
id: collect-261001-ia-llm/ia-llm/tutoriel-n8n-2026-workflow-ia-auto-heberge-en-13-etapes-3
title: "Verifier la version de Node.js (20 ou superieur requis)"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Mistral", "OpenAI"]
dates: []
keywords: ["agent", "agents", "gemini", "mistral"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-n8n-2026-workflow-ia-auto-heberge-en-13-etapes.md
source_anchor: ""
source_lines: [188, 293]
sha256: 99330a6b4fa9b63acc441f6dd72189be36d1878d4a7190854e78f7e8b0034d59
---

# Verifier la version de Node.js (20 ou superieur requis)

## Étape 8 : Le nœud Code – transformer les données en JavaScript

**Étape 8.** Le tableau d’IDs brut n’est pas exploitable tel quel. Le **nœud Code** est la signature de n8n : il permet d’écrire du JavaScript (ou du Python) arbitraire pour manipuler les données. Ajoutez un nœud Code après le HTTP Request et limitez-vous aux 5 premiers articles pour ne pas saturer l’API ni le LLM :

```
// Nœud Code – mode "Run Once for All Items"
// Récupère le tableau d'IDs et garde les 5 premiers
const ids = $input.first().json;
const topFive = ids.slice(0, 5);
// Retourne un item par identifiant d'article
return topFive.map((id) => {
  return {
    json: {
      articleId: id,
      url: `https://hacker-news.firebaseio.com/v0/item/${id}.json`,
    },
  };
});
```
Exécutez ce nœud : il produit désormais 5 items distincts, chacun contenant un `articleId` et une `url`. Ajoutez ensuite un second nœud HTTP Request paramétré pour récupérer le détail de chaque article. Utilisez une **expression n8n** pour injecter dynamiquement l’URL de l’item courant :

```
# Dans le second nœud HTTP Request, champ URL :
={{ $json.url }}
# Sortie attendue par article (extrait) :
{
  "by": "developer42",
  "title": "n8n leve 2,3 Md$ pour l'automatisation IA",
  "url": "https://example.com/article",
  "score": 412
}
```
La syntaxe `={{ ... }}` est le moteur d’expressions de n8n : tout ce qui est entre doubles accolades est évalué comme du JavaScript au moment de l’exécution. C’est ce qui permet de chaîner dynamiquement les nœuds sans écrire de glue code fastidieux. À ce stade, votre workflow récupère le titre et l’URL des 5 meilleurs articles du jour.

## Étape 9 : Ajouter un agent IA (OpenAI, Mistral ou Ollama)

**Étape 9.** Nous arrivons au cœur du projet : faire résumer ces articles par un grand modèle de langage. n8n intègre **LangChain** et propose des nœuds IA dédiés, dont le puissant nœud **AI Agent** et des nœuds de modèle de chat pour OpenAI, Anthropic, Mistral, Google Gemini et Ollama. Depuis la version **2.35.0**, publiée le **11 août 2026**, n8n embarque également un **AI Assistant** et un **Agent Builder** natifs qui accélèrent la conception de ces nœuds directement dans l’éditeur ; un mois plus tard, en septembre 2026, la version stable est passée à **2.38.5**, avec une bêta **2.39.1** déjà accessible aux testeurs selon la documentation officielle de n8n. Ajoutez un nœud « Message a model » (ou « Basic LLM Chain ») et reliez-y un nœud de modèle de chat.

Configurez d’abord les **identifiants** (credentials). Pour OpenAI, collez votre clé `sk-...`. Pour rester dans une logique 100 % européenne, préférez **Mistral AI** ; pour une confidentialité absolue et un coût nul, pointez le nœud Ollama vers votre instance locale. Voici la configuration pour une instance Ollama tournant sur la même machine :

```
# Identifiants du noeud "Ollama Chat Model" dans n8n
Base URL: http://host.docker.internal:11434
Model:    llama3.1:8b
# Verifier qu'Ollama repond bien depuis l'hote :
$ curl -s http://localhost:11434/api/tags | head -c 80
{"models":[{"name":"llama3.1:8b","model":"llama3.1:8b", ...
```
Notez l’adresse `host.docker.internal` : depuis l’intérieur du conteneur n8n, c’est ainsi qu’on atteint un service (Ollama) tournant sur la machine hôte. Rédigez ensuite le **prompt** du modèle en injectant le titre de l’article via une expression :

```
Tu es un assistant de veille technologique francophone.
Resume en deux phrases claires, en francais, l'article suivant
et indique pourquoi il est pertinent pour un public technique europeen.
Titre : {{ $json.title }}
URL : {{ $json.url }}
Score : {{ $json.score }}
```
Exécutez : pour chaque article, le LLM renvoie un résumé concis en français. Vous venez de construire un véritable **agent IA d’automatisation**. Cette approche reprend les principes des frameworks multi-agents comme CrewAI, mais sans écrire une seule ligne d’orchestration : tout se fait visuellement dans n8n.

## Étapes 10 & 11 : Logique conditionnelle et envoi du résultat

**Étape 10.** Tous les articles ne méritent pas une notification. Ajoutez un nœud **IF** pour ne conserver que les articles dont le score dépasse un seuil, par exemple 200. Le nœud IF crée deux sorties – *true* et *false* – et n’achemine vers la suite que les items satisfaisant la condition.

```
# Configuration du noeud IF
Condition 1:
  Value 1:   ={{ $json.score }}
  Operation: Number > (superieur a)
  Value 2:   200
# Pour des branchements multiples, utilisez le noeud "Switch"
# qui route les items vers N sorties selon une regle.
```
**Étape 11.** Reliez la sortie *true* à un nœud d’envoi. n8n propose des nœuds **Send Email** (SMTP), **Slack**, **Discord** ou **Telegram**. Pour Slack, créez les identifiants OAuth ou utilisez un webhook entrant, puis composez le message en agrégeant les résumés :

```
# Noeud Slack – champ "Text"
=*Veille tech du jour* :star:
{{ $json.title }}
{{ $json.text }}
:link: {{ $json.url }}
```
Vous pouvez aussi insérer un nœud Code final pour fusionner les 5 résumés en un seul message récapitulatif, plus agréable à lire qu’une rafale de notifications. À ce stade, le workflow est fonctionnellement complet : il se déclenche, récupère, transforme, résume par IA, filtre et notifie.

## Étape 12 : Tester, activer et planifier le workflow complet

**Étape 12.** Avant de l’activer, lancez une exécution manuelle complète avec le bouton « Test workflow ». n8n exécute tous les nœuds dans l’ordre et colore en vert ceux qui réussissent, en rouge ceux qui échouent. Cliquez sur chaque nœud pour inspecter ses données d’entrée/sortie – c’est la méthode de débogage la plus efficace dans n8n.

Une fois le test concluant, basculez l’interrupteur **Active** en haut à droite. Le workflow s’exécutera désormais automatiquement selon votre planning (chaque matin à 8 h dans notre cas). Consultez l’onglet **Executions** pour voir l’historique : chaque exécution est journalisée avec son statut, sa durée et ses données, ce qui facilite l’audit et le diagnostic.

```
# Exemple de sortie console (logs du conteneur) lors d'une execution
$ docker compose logs -f n8n
[info] Workflow execution started   workflowId=2 executionId=147
[info] Node "HTTP Request" finished  items=5
[info] Node "AI Agent" finished      items=5
[info] Node "IF" finished            true=3 false=2
[info] Node "Slack" finished         items=3
[info] Workflow execution finished   status=success duration=4.2s
```
Astuce : activez l’option « Save successful executions » uniquement si vous avez besoin de l’historique complet. Sur un workflow très fréquent, conserver toutes les exécutions réussies peut faire gonfler rapidement votre base PostgreSQL. Réglez la rétention via la variable `EXECUTIONS_DATA_MAX_AGE` (en heures).

## Étape 13 : Passage en production – queue mode, Redis et HTTPS

**Étape 13.** Une instance de test sur `localhost` ne suffit pas en production. Trois chantiers s’imposent : la **scalabilité**, la **sécurité** et la **haute disponibilité**. Pour absorber de gros volumes d’exécutions, n8n propose le **queue mode** : le processus principal délègue les exécutions à des *workers* dédiés via une file **Redis**. Voici les variables d’environnement clés :

