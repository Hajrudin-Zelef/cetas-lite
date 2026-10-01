---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-12
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Anthropic", "Baidu", "Google", "Microsoft", "Perplexity"]
dates: []
keywords: ["mcp", "agent", "agents", "claude", "embeddings", "perplexity"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [1921, 2072]
sha256: 304ab5dfc39b3da463c3a5f92d091d5f998e2443a6cd65efe7f283379bf577b9
---

# MCP, LSP, code-server, websearch — Guide pratique

🔒 Rappel : code-server = un shell complet dans ton navigateur. Le niveau
de protection doit être celui d'un accès SSH, pas celui d'un blog.

## 95. Extensions et marketplace

code-server utilise le marketplace **Open VSX** (pas celui de Microsoft,
licence oblige). 99 % des extensions courantes y sont (Python, Go,
rust-analyzer, Docker, GitLens...). Installation : comme dans VS Code
(`Ctrl+Maj+X`) ou en CLI :

```bash
code-server --install-extension ms-python.python
```

⚠️ Quelques extensions propriétaires (ex : certaines de Microsoft) refusent
de tourner hors VS Code officiel. Teste avant de t'engager.

## 96. Cas d'usage pour Zelef

1. **Dev distant des scripts RAG** : ton `collect_*.py` tourne sur le
   serveur qui a la bande passante, tu codes depuis n'importe quel poste.
2. **Console d'administration** : code-server + terminal intégré + tes
   guides en Markdown preview = poste d'astreinte dans le navigateur.
3. **Agent codeur hébergé** : Claude Code tourne dans le terminal de
   code-server, sur la machine qui a les accès (VPN, creds) — pas sur ton
   laptop.
4. **Formation / pair** : partage d'écran d'un workspace identique pour
   montrer une manip à un technicien.

## 97. Limites et alternatives

| Limite | Détail | Alternative |
|---|---|---|
| Pas le marketplace MS | Quelques extensions manquantes | `code tunnel` (officiel MS) |
| Latence | Frappe distante si loin du serveur | Serveur proche (même DC/pays) |
| Ressources | VS Code + extensions = gourmand | 2 vCPU / 4 Go mini par user actif |
| 1 user par instance | Pas de multi-tenant natif | **Coder** (la plateforme, au-dessus de code-server) |

- **`code tunnel`** (Microsoft, intégré à VS Code desktop) : expose ton VS
  Code local via le cloud MS — zéro infra, mais dépendance Microsoft.
- **Coder** : la plateforme complète (workspaces éphémères en Terraform,
  code-server en app) — si toute l'équipe doit avoir un IDE distant.

💡 Pour un usage perso/équipe réduite : code-server + Docker + Caddy +
WireGuard = simple, robuste, sans dépendance externe.

# PARTIE VII — WEBSEARCH : DONNER LE WEB AUX AGENTS

## 98. Pourquoi un agent a besoin de recherche web

Le modèle a une **date de coupure** : il ne connaît pas les sorties
récentes, les CVE de la semaine, les changements d'API. Sans recherche web,
il hallucine du plausible sur du récent. Avec un outil `web_search`, l'agent
boucle : chercher → lire → synthétiser → vérifier.

Pour ton RAG, c'est le complément du corpus local : **corpus = ton savoir
stable et vérifié ; websearch = le monde qui bouge**. Un bon agent fait les
deux : d'abord `search_docs` (ta vérité terrain), puis `web_search` (les
mises à jour), puis il **croise**.

⚠️ Le web n'est pas une source de vérité : c'est une source à **citer et
vérifier**. L'agent doit toujours donner ses sources (URLs) et signaler les
contradictions avec ton corpus.

## 99. Anatomie d'une API de recherche « agent-ready »

| Critère | Question à poser |
|---|---|
| Index | Propre (Brave, Exa, Mojeek) ou revente Google/Bing (Serper, SerpAPI) ? |
| Sortie | Snippets bruts ou contenu nettoyé « LLM-ready » (Tavily, Exa) ? |
| Extraction | Peut-elle renvoyer le **contenu** des pages (`/extract`) ? |
| Fraîcheur | Paramètre `time_range` / tri par date ? |
| Prix | Au 1 000 requêtes ? Free tier ? Facturation au token ? |
| Rate limits | Requêtes/seconde et /minute du plan visé ? |
| SDK / MCP | Client Python officiel ? Serveur MCP officiel ? |
| ToS | Scraping Google revendu = risque contractuel (Serper & co) |

💡 Le critère n°1 pour un agent : **la qualité du contenu renvoyé par appel**.
Une API chère qui rend du Markdown propre avec citations (Tavily/Exa) coûte
souvent **moins cher au final** qu'une API pas chère dont les snippets
obligent à 3 appels de plus.

## 100. Tavily : fiche détaillée

- **Positionnement** : pensé pour les agents IA dès le départ.
- **Sortie** : JSON avec `content` nettoyé par résultat + `answer` synthétisée
  optionnelle (avec citations) — directement injectable dans un prompt.
- **Endpoints** : `/search` (profondeurs `basic`/`advanced`), `/extract`
  (récupère le contenu propre d'URLs), la même clé API fait les deux.
- **Prix (sept 2026, 📌 à vérifier)** : ~1 000 crédits/mois gratuits sans CB ;
  ~8 $/1k en basic, ~16 $/1k en advanced ; PAYG ~0,008 $/crédit.
- **Intégrations** : SDK Python/JS officiels, serveur **MCP officiel**,
  LangChain natif.
- **Forces** : qualité du contenu, citations, anti-bruit. **Faiblesse** : pas
  d'index propre (agrégateur), prix supérieur aux SERP brutes.
- **Idéal pour** : agent qui doit répondre avec sources, RAG d'appoint.

## 101. Exa : fiche détaillée

- **Positionnement** : recherche **neurale** (embeddings) + index propre.
- **Sortie** : JSON + **contenu des pages** (pas seulement des snippets) —
  énorme pour la qualité RAG.
- **Prix (📌 à vérifier)** : ~1 000 req/mois + 10 $ de crédits gratuits ;
  ~7 $/1k en recherche, ~12-15 $/1k en « deep » ; MCP ~150 appels/jour
  sur l'entrée de gamme.
- **Forces** : trouve des pages « à propos de X » même sans les mots-clés
  exacts ; contenu riche. **Faiblesse** : rappel par mot-clé exact en retrait
  vs Google ; plus lent.
- **Idéal pour** : recherche sémantique (« trouve des docs sur... »),
  alimentation RAG.

## 102. Brave Search API : fiche détaillée

- **Positionnement** : **index indépendant** (30 Md+ de pages), pas un
  revendeur Google — le seul avec son propre crawl à ce prix.
- **Sortie** : SERP JSON brute (`web.results[]` + `extra_snippets`) ;
  endpoint « Answers » avec contexte LLM.
- **Prix (📌 à vérifier)** : ~5 $/1k requêtes ; ~5 $/mois de crédits offerts
  (≈ 1k requêtes) — carte requise pour le dépassement.
- **Auth** : header `X-Subscription-Token`.
- **Forces** : indépendance, latence faible, bon rapport qualité/prix,
  serveur **MCP officiel**. **Faiblesse** : snippets bruts (à toi de
  nettoyer/extraire).
- **Idéal pour** : volume élevé à petit budget, usage généraliste.

## 103. Serper : fiche détaillée

- **Positionnement** : **le moins cher pour du vrai SERP Google** (scraping).
- **Prix (📌 à vérifier)** : ~0,30 à 1 $/1k requêtes selon volume ;
  2 500 requêtes gratuites **one-shot** (pas par mois).
- **Auth** : header `X-API-KEY`, POST JSON.
- **Forces** : prix imbattable, qualité Google. **Faiblesses** : revente de
  scraping (risque ToS), pas d'extraction de contenu, free tier non
  renouvelable.
- **Idéal pour** : gros volumes de requêtes simples où seul le classement
  Google compte, en acceptant le risque contractuel.

## 104. Les autres : SerpAPI, Perplexity, You.com, Firecrawl, Parallel, Linkup

| API | Prix ordre de grandeur (📌) | Note |
|---|---|---|
| **SerpAPI** | 25 $+/mois, ~250 req/mois gratuites | Multi-moteurs (Google, Bing, Baidu...), vieux, fiable, cher |
| **Perplexity Sonar** | ~5 $/1k + tokens | Réponses synthétisées avec sources, pas du SERP brut |
| **You.com** | ~100 $ de crédit d'essai, ~5 $/1k | Search + contenu de page groupés |
| **Firecrawl** | 1 000 crédits/mois gratuits, plans 16-599 $/mois | Search + scrape Markdown en un appel, MCP officiel |
| **Parallel AI** | ~1 $/1k (turbo) à 5 $/1k | Champs structurés avec citations + confiance |
| **Linkup** | ~5 €/1k | Acteur UE, réponse sourcée + brut |
| **Kagi** | ~25 $/1k | Excellente qualité, anti-SEO-spam, mais cher |
| **Jina Reader** | Gratuit (rate-limit sans clé) | `https://s.jina.ai/<url>` → texte propre ; dépanne toujours |

## 105. Les morts : à ne plus choisir en 2026

