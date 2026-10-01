---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-9
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI", "Z.ai"]
dates: ["2026-06-16"]
keywords: ["agent", "agentic", "agents", "claude", "copilot", "glm", "mcp", "open source", "research", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [1226, 1360]
sha256: f8bdd15c426efddc52b90840f0b6dfbe912c488971b746da7dcde97a80e2753b
---

# Concepts : agents IA, agentic, autonomie

1. **Jamais en root** : lance-le sous ton utilisateur, ou mieux sous un compte
   dédié (montage 1, section 52). `auto_run: false` en permanence.
2. **Lis le code avant de valider** : `verbose: true`, et prends 10 secondes par
   bloc. Si tu ne comprends pas ce qu'il fait, refuse.
3. **Safe mode activé** (`ask` minimum), en sachant que c'est expérimental.
4. **Données sensibles** : ne lui fais pas traiter de fichiers contenant des
   secrets en clair si le modèle est cloud (les données partent chez le provider).
   Modèle local (Ollama/LM Studio) si la confidentialité l'exige.
5. **Sandbox pour l'inconnu** : fichier téléchargé d'internet à analyser ?
   Conteneur jetable (montage 2, section 52), pas ta machine.
6. **Pas d'accès réseau aveugle** : s'il doit télécharger quelque chose, vérifie
   l'URL et le contenu avant.
7. **Sauvegardes** : la règle des incidents (sections 44-45) s'applique ici aussi :
   ce que l'agent peut toucher doit être sauvegardé et restaurable.

En une phrase : traite Open Interpreter comme un collègue très rapide qui a les
mêmes droits que toi — c'est-à-dire avec exactement la prudence que ça mérite.

---

# PARTIE E — « cowork » : le nom ambigu

## 65. « cowork » : pourquoi ce nom est un problème

Recherche effectuée fin septembre 2026 : **il n'existe pas UN outil canonique
appelé « cowork »**. Le nom désigne au moins cinq réalités différentes, dans
trois écosystèmes. Si quelqu'un te dit « utilise cowork », demande **lequel** :
sans nom d'éditeur, la consigne est inexploitable. Ci-dessous l'inventaire vérifié.

## 66. Candidat n°1 : Claude Cowork (Anthropic)

- **Quoi** : fonctionnalité « agentic desktop » de Claude, annoncée en
  **research preview en janvier 2026** (d'après la documentation communautaire).
- **Principe** : au lieu du terminal (Claude Code, pour les devs), Cowork s'adresse
  aux « knowledge workers » via l'app desktop : l'agent accède à des **dossiers
  locaux** et manipule des fichiers (organisation, génération de documents,
  synthèse), **sans exécution de code**.
- **Accès relevé** : abonnement Pro/Max, macOS uniquement (au stade de la preview).
- **Différence clé vs Claude Code** : pas d'exécution de code, sandbox dossier.
- **Statut à vérifier** : une research preview de janvier 2026 a pu évoluer
  (GA ? abandon ? renommage ?). Revérifie avant tout plan qui en dépend.

## 67. Candidat n°2 : Microsoft Copilot Cowork

- **Quoi** : agent IA intégré à **Microsoft 365 Copilot**, présenté comme
  « il ne répond pas, il fait le travail » : tu décris un résultat
  (« prépare ma réunion client de demain »), il planifie et exécute à travers
  Outlook, Teams, Word, Excel, SharePoint, avec points de validation.
- **Disponibilité relevée** : GA le **16 juin 2026** pour les clients M365 Copilot
  (activation par l'admin, désactivé par défaut).
- **Modèle éco relevé** : licence M365 Copilot (~30 $/utilisateur/mois) + facturation
  à l'usage en « Copilot Credits » (~0,01 $/crédit, plafonds configurables).
- **Angle entreprise** : gouvernance Entra ID, audit trails, human-in-the-loop —
  c'est le « cowork » le plus « DSI-compatible » de la liste.

## 68. Candidat n°3 : Cowork (fonctionnalité de BrowserOS)

- **Quoi** : fonctionnalité du navigateur agentique **BrowserOS** (open source) :
  l'agent combine **automatisation du navigateur + accès fichiers locaux**
  (7 outils : lire/écrire/éditer des fichiers, exécuter des commandes shell
  dans un dossier sandboxé, recherche regex, etc.).
- **Principe** : « décris une tâche complexe, l'agent la fait de bout en bout » —
  recherche web puis sauvegarde du rapport dans ton dossier, dans la même session.
- **Public** : plutôt tech, open source, auto-hébergeable.

## 69. Candidat n°4 : CoWork OS (projet open source)

- **Quoi** : « personal AI super app » open source (**cowork-os/cowork-os**) :
  hub d'agents avec GUI, exécutions planifiées, passerelle multi-canaux
  (WhatsApp, Telegram, Discord, Slack), 30+ providers LLM dont Ollama en local,
  automatisation navigateur via Playwright, garde-fous configurables.
- **Positionnement** : équipe d'agents personnels tournant sur ta machine,
  architecture local-first, BYOK (bring your own key).
- **À noter** : projet communautaire ambitieux ; avant adoption, vérifier
  maturité, sécurité du code et maintenance réelle (audit indispensable —
  c'est un agent avec accès local + messageries).

## 70. « cowork » : synthèse et recommandation

| Nom | Éditeur / nature | Public | Statut relevé (sept 2026) |
|---|---|---|---|
| Claude Cowork | Anthropic | knowledge workers, macOS | research preview (janv. 2026) — à revérifier |
| Copilot Cowork | Microsoft | entreprises M365 | GA 16/06/2026, facturation à l'usage |
| Cowork (BrowserOS) | open source | tech | fonctionnalité du navigateur agentique |
| CoWork OS | open source communautaire | power users | en développement actif — auditer avant usage |
| CoWork (chuk_chat) | plan seul, pas de code | — | document de planification (août 2026), pas un produit |

**Recommandation** : si ton besoin est « un agent qui manipule mes fichiers en
langage naturel », les deux pistes sérieuses sont **Claude Cowork** (si tu es
dans l'écosystème Anthropic/macOS) et **Copilot Cowork** (si tu es en entreprise
Microsoft 365). Pour un sysadmin qui veut du contrôle local : ni l'un ni l'autre
ne remplace un agent maison sandboxé (partie G) ou Open Interpreter (partie D).
Et surtout : **n'installe jamais un « cowork » trouvé au hasard** — avec 5
homonymes, le risque de télécharger le mauvais (voire un malware) est réel.
Vérifie l'éditeur, l'URL officielle et les signatures.

---

# PARTIE F — « zcode » : identifié, c'est ZCode de Z.ai

## 71. « zcode » : identification (recherche sept 2026)

Le nom **« zcode »** correspond à **ZCode**, l'environnement de développement
agentique de **Z.ai (Zhipu AI)** — le laboratoire chinois derrière les modèles
**GLM** (modèles ouverts/ouverts en poids, ex. GLM-5.x). Identification croisée
sur plusieurs sources indépendantes (cartographies d'agents, docs d'intégrateurs,
wiki de recherche), toutes convergentes fin septembre 2026.

En une phrase : **ZCode = le « Claude Code / Codex » officiel de l'écosystème GLM**,
mais avec une philosophie différente : une application centrée sur la **tâche**
(pas sur l'éditeur), pensée pour des sessions longues et pilotables à distance.

## 72. ZCode : ce que c'est (fonctionnalités relevées)

D'après les sources relevées (docs officielles rapportées + analyses tierces,
versions 3.x, été 2026) :

- **3 interfaces, un seul runtime agent** : application **desktop** (Electron ;
  Windows, macOS, Linux), **interface web**, et **CLI terminal** (`zcode`).
- **Modèle de référence : GLM** (GLM-5.2 / 5.3 selon version), contexte long
  (1M tokens revendiqué), co-design modèle+harness.
- **Multi-provider** : GLM en natif + tout endpoint compatible OpenAI ou Anthropic
  en provider personnalisé.
- **Sessions longues « Goal »** : tâches à long horizon avec vérification,
  pilotables à distance (WeChat, Feishu, Telegram d'après les descriptions).
- **Extensibilité style Claude Code** : skills, subagents, hooks, MCP,
  configuration via `AGENTS.md`, marketplace de plugins.
- **Dev distant** : SSH / WSL / Docker ; automatisations planifiées.
- **Tarification relevée** : harness gratuit ; l'usage des modèles GLM passe par
  les formules « coding plan » de Z.ai (détails et prix : à vérifier sur place).

## 73. ZCode : versions et open-sourcing (point de vigilance)

Points relevés fin septembre 2026 — à manier avec prudence car ça bouge vite :

