---
id: collect-261001-ia-llm/ia-llm/tutoriel-n8n-2-20-workflows-ia-en-13-etapes-2026-3
title: "Nœud Code – Langage Python"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Mistral", "OpenAI", "Stripe"]
dates: []
keywords: ["agent", "claude", "gpu", "llama", "memory", "mistral", "open weights", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-n8n-2-20-workflows-ia-en-13-etapes-2026.md
source_anchor: ""
source_lines: [200, 299]
sha256: 7e603d923710aad2f1c9986c7eec336075da6d95f04eb68adf14b992822640de
---

# Nœud Code – Langage Python

n8n inclut un nœud **Postgres** de première classe qui exploite le pilote *pg* de Node.js. Créez d’abord une credential Postgres pointant vers le conteneur déjà déployé. L’hôte sera `postgres`, le port 5432, et la base `n8n` (ou créez une base dédiée `app_data`).

Avant d’insérer des données, créez la table qui accueillera les statistiques GitHub. Ajoutez un nœud **Postgres** en mode *Execute Query* en début de workflow :

```
CREATE TABLE IF NOT EXISTS github_repos (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  stars INTEGER NOT NULL,
  forks INTEGER NOT NULL,
  density NUMERIC(10, 3),
  url TEXT,
  fetched_at TIMESTAMPTZ DEFAULT NOW(),
  UNIQUE(name)
);
CREATE INDEX IF NOT EXISTS idx_stars ON github_repos (stars DESC);
```
Ajoutez ensuite un second nœud Postgres en mode *Insert* juste après le nœud Code. Mappez chaque colonne du tableau JSON entrant vers une colonne de la table. n8n propose une syntaxe d’expression `{{ $json.name }}` qui suit le format de Mustache et autocomplète les chemins disponibles.

Pour gérer les conflits sur la contrainte UNIQUE, activez l’option *Options > Mode* sur *Insert* avec la clause *ON CONFLICT*. Vous pouvez aussi écrire la requête à la main avec *Execute Query* :

```
INSERT INTO github_repos (name, stars, forks, density, url)
VALUES ({{$json.name}}, {{$json.stars}}, {{$json.forks}}, {{$json.density}}, {{$json.url}})
ON CONFLICT (name) DO UPDATE
SET stars = EXCLUDED.stars,
    forks = EXCLUDED.forks,
    density = EXCLUDED.density,
    fetched_at = NOW();
```
Cette approche *upsert* garantit que le workflow reste idempotent. Vous pouvez le relancer plusieurs fois par jour sans dupliquer les lignes. Pour des volumes plus importants, considérez Supabase comme alternative gérée à PostgreSQL avec ses propres rôles et son API auto-générée.

## Étape 7 : Exposer un Webhook et recevoir des événements externes

Le nœud **Webhook** transforme n8n en endpoint HTTP que vous pouvez appeler depuis n’importe quel service externe : Stripe, GitHub, Typeform, formulaire HTML. Créez un nouveau workflow nommé *Stripe payment handler* et glissez un nœud Webhook comme trigger.

Configurez la méthode HTTP en POST, le chemin (path) en `stripe-events` et le mode d’authentification en *Header Auth* avec une credential nommant l’en-tête `stripe-signature`. n8n affiche immédiatement deux URLs : une de test, valable tant que l’éditeur reste ouvert, et une de production qui ne fonctionne qu’une fois le workflow activé.

Pour développer localement avec un service externe comme Stripe, utilisez **ngrok** ou **localtunnel** pour exposer votre instance Docker au monde. Voici un exemple de commande Cloudflared, plus pérenne que ngrok pour des tests longs :

`cloudflared tunnel --url http://localhost:5678`
Le tunnel retourne une URL en `https://random-name.trycloudflare.com`. Mettez à jour la variable `WEBHOOK_URL` dans votre `docker-compose.yml` avec cette URL et redémarrez le conteneur. Désormais, les services tiers atteignent votre n8n local depuis Internet via HTTPS, condition sine qua non pour la plupart des webhooks Stripe, GitHub ou Shopify qui refusent HTTP.

## Étape 8 : Construire un agent IA avec le nœud AI Agent et LangChain

C’est ici que n8n 2.0 prend toute sa puissance. Le nœud **AI Agent**, basé sur LangChain.js, exécute une boucle ReAct (Reasoning + Action) avec accès à des outils. Selon les notes de Latenode, le nœud bénéficie d’une gestion améliorée des tokens et de performances accrues dans les dernières versions 2.x.

Créez un workflow *Assistant veille tech*. Glissez un **Chat Trigger**, puis un **AI Agent**. Connectez à l’AI Agent trois sous-nœuds : un modèle (OpenAI Chat Model avec GPT-4 ou Anthropic Claude), une mémoire (Window Buffer Memory) et un ou plusieurs outils.

- **Model** : OpenAI Chat Model, modèle*gpt-4.1-mini* , température 0.2
- **Memory** : Window Buffer Memory avec 10 derniers messages
- **Tool 1** : HTTP Request Tool ciblant l’API Hacker News
- **Tool 2** : Postgres Tool exécutant des SELECT sur la table`github_repos`
- **Tool 3** : Wikipedia Tool pour les définitions techniques

Le prompt système pilote le comportement de l’agent. Voici un exemple efficace pour un assistant tech français :

```
Tu es un assistant de veille technologique francophone.
Tu disposes de trois outils :
1. hacker_news_top_stories : récupère les meilleurs articles du jour.
2. github_repos_query : interroge notre base PostgreSQL.
3. wikipedia_lookup : cherche une définition.
Réponds toujours en français. Cite tes sources entre parenthèses.
Si tu n'as pas l'information, dis-le explicitement plutôt qu'inventer.
Format des réponses : trois sections – Contexte, Faits, Recommandation.
```
Activez le workflow et ouvrez le widget de chat intégré (icône en bas à droite de l’éditeur). Posez la question : « Quels sont les dépôts GitHub les plus actifs cette semaine ? ». L’agent décide d’appeler l’outil *github_repos_query*, formule une requête SQL, parse le résultat et produit une réponse structurée en français.

## Étape 9 : Intégrer Ollama pour des modèles 100 % auto-hébergés

Pour les organisations soumises à des contraintes de confidentialité – santé, droit, défense – l’appel à OpenAI ou Anthropic est rédhibitoire. n8n inclut depuis fin 2025 un nœud **Ollama Chat Model** qui se branche sur une instance Ollama locale. Pointez l’URL de base sur `http://host.docker.internal:11434` si Ollama tourne sur le même hôte, ou sur l’IP interne si vous opérez en cluster.

Pour télécharger un modèle *Mistral-Small 3.2* open weights ou *Llama 4 8B*, exécutez préalablement sur l’hôte :

```
ollama pull mistral-small:3.2
ollama pull llama4:8b
ollama list
```
Remplacez le nœud OpenAI Chat Model par Ollama Chat Model dans votre workflow AI Agent. Sélectionnez le modèle dans la liste déroulante. Les performances dépendront fortement de votre GPU : une carte RTX 4090 délivre environ 90 tokens/seconde sur Mistral-Small, contre 15 tokens/seconde sur un MacBook Pro M2 Max sans accélération GPU dédiée.

Cette architecture rejoint les bonnes pratiques décrites dans notre tutoriel Ollama Docker et permet de respecter le RGPD sans renoncer à l’IA générative. C’est la voie privilégiée par de nombreuses ETI françaises selon le baromètre Cigref 2026.

## Étape 10 : Workflow complet – tri intelligent d’e-mails avec IA

Mettons tout cela en pratique avec un workflow de production : un trieur d’e-mails entrants qui classe automatiquement chaque message reçu sur une adresse `support@` en trois catégories – urgent, commercial, spam – et qui crée un ticket Notion correspondant.

Architecture du workflow :

- **Email Trigger (IMAP)** : surveille la boîte support@
- **Set Node** : extrait sujet, expéditeur, corps en texte brut
- **AI Agent (OpenAI)** : classe en JSON {category, priority, summary}
- **Switch Node** : route selon la category
- **Notion Node** : crée une page dans la base Tickets
- **Slack Node** : notifie le canal #support pour les urgents
- **Postgres Node** : log de tous les e-mails traités

Le prompt envoyé à l’AI Agent doit être déterministe et retourner du JSON strict. Configurez le modèle avec *response_format = json_object* côté OpenAI pour éviter les hallucinations de syntaxe :

