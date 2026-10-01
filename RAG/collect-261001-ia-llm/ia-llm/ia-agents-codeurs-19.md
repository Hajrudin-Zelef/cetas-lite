---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-19
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "DeepSeek", "Meta", "Microsoft", "OpenAI", "SpaceX", "xAI"]
dates: ["2026-09-02", "2026-09-03", "2026-09-12", "2026-09-14", "2026-09-15", "2026-09-20", "2026-09-21", "2026-09-22", "2026-09-23", "2026-09-24", "2026-09-25", "2026-09-27", "2026-09-29"]
keywords: ["agent", "agents", "astra", "aws", "claude", "copilot", "deepseek", "gpt-6", "gpu", "grok", "grok 4", "leaderboard"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [3001, 3141]
sha256: 86119b988726d823fbdbb45f9cc04200020eba1dbc6dd813258b5ee89451e20b
---

# Les agents codeurs IA

- **Codex CLI 0.156.0 (22/09/2026)** : voice mode activé par défaut
  (bascule `F8`, réglages `/voice`) ; interface plein écran optionnelle
  `/tui` (recherche dans le transcript, sélection souris, diagrammes
  Mermaid, notation math) ; tableau de bord d'usage `/usage` ; sessions
  worktree activées par défaut. Patch **0.156.1** (même jour) : **GPT-6
  Sol et GPT-6 Luna** ajoutés au sélecteur de modèles, Luna recommandé
  comme repli économique en cas de rate limit.
- **GPT-6 Astra** (début sept. 2026) : OpenAI annonce pour Codex un mode
  de **notes persistantes entre fenêtres de contexte** — les détails
  accumulés survivent sans compaction brutale, et les anciennes fenêtres
  restent interrogeables. Expérimental au lancement, promis comme défaut
  « dans les prochaines semaines ». À surveiller : c'est la première
  réponse officielle du constructeur au problème de la « mémoire de
  session » que ce guide contourne par fichiers (section 106).
- **Feuille de route Codex (annoncée par OpenAI)** : support Windows de
  l'app desktop (aujourd'hui macOS), **Automations** avec déclencheurs
  cloud (agents en tâche de fond permanente, pas seulement PC allumé),
  **plan mode** en lecture seule avant exécution, **personnalités**
  configurables (`/personality`).
- **Codex Cloud nouvelle génération — RUMEUR non confirmée** : analyse
  par reverse-engineering (RuntimeWire, 20/09/2026) du build desktop
  9922 montrant un parcours « Cloud » refondu (feature-gated) séparant
  « Cloud (Legacy) », avec préparation d'environnement pilotée par
  l'agent via **Tailscale**, secrets injectés par proxy et identités
  cloud (Azure/AWS/GCP). Non annoncé par OpenAI : ne pas planifier
  dessus.
- Sources : *Oday Bakkour*, roundup du 24/09/2026 ; *codex-timeline*
  (GitHub, vérifié 23/09/2026) ; *9to5Mac*, 03/09/2026 ; *VentureBeat*
  (roadmap Codex) ; *RuntimeWire*, 20/09/2026.

### 113.3 GitHub Copilot — sandboxing local (public preview)

Annonce GitHub du **23/09/2026** : **sandboxing local** dans l'app
Copilot (public preview). Politiques par projet sur trois axes :
système de fichiers (lecture/écriture, lecture seule, dossiers
interdits), réseau (internet sortant, réseau local), identifiants
(Git HTTPS, GitHub CLI). Désactivé par défaut (`Sandbox new sessions`
par projet, ou `/sandbox on` en session) ; les admins Entreprise
peuvent imposer plus strict ; si l'OS ne peut appliquer la politique,
la session échoue **fermée** plutôt que de tourner sans sandbox.

- Pour ton cadre de sécurité (section 92) : c'est la version « produit »
  du confinement que ce guide fait à la main (utilisateur dédié,
  conteneurs). Limite connue : **sessions locales uniquement**, pas
  cloud/remote pour l'instant.
- La semaine du 21/09/2026 ajoute aussi à Copilot : **Claude Opus 5.5**,
  **GPT-6 Sol**, **GPT-6 Luna**, **Grok 4.7** (selon plans) ;
  **OpenTelemetry** pour tracer l'activité des agents dans tes outils
  de supervision existants ; **assisted approvals** (approbation auto
  des appels à faible risque) en public preview ; édition d'un message
  antérieur avec **rembobinage** de la conversation et des fichiers.
- Sources : *GitHub Changelog*, 23/09/2026 (« Local sandboxing in the
  GitHub Copilot app ») et 25/09/2026 (« Copilot weekly releases —
  September 21 ») ; *Nandann*, guide sandboxing, 24/09/2026.

### 113.4 Cline — Desktop (bêta) et v4.1.21

- **Cline Desktop, bêta macOS/Windows annoncée le 14/09/2026** par le
  fondateur : l'agent devient une app autonome (même moteur SDK que
  l'extension). Fonctions phares : **import de sessions Claude Code /
  Codex** pour continuer avec un autre modèle, **jobs planifiés** (revue
  de PR, scans sécu, doc), marketplace intégré (plugins, serveurs MCP,
  skills). Modèles open-weight gratuits cités au lancement :
  DeepSeek-V4.1-Flash, Musespark-1.3 (offre gratuite tournante, quotas
  journaliers).
- **v4.1.21 (24/09/2026)** : correctif de fiabilité pour les modèles
  locaux (llama.cpp, Ollama, LM Studio) — compaction + un retry au lieu
  d'abandonner la tâche ; nouveau provider « ai& » ; 19 providers sans
  version épinglée basculent leur défaut vers **Claude Opus 5.5**.
- Sources : *RuntimeWire*, 14/09/2026 ; *ExplainX*, 15/09/2026 ;
  *TPS Report*, 24/09/2026 ; *morphllm.com*, leaderboard sept. 2026.

### 113.5 Cursor — agents cloud sur ton infra

Le **update de septembre 2026** (documenté par *developer-tech.com*)
ajoute l'exécution des **Cloud Agents sur infrastructure propre** :
pools de workers dynamiques mutualisés par équipe (au lieu d'un worker
par machine/dépôt), pools séparés par environnement (GPU, Mac iOS),
**opérateur Kubernetes** (warm capacity, rolling updates, rotation de
tokens) et déploiement **Cloud Run** avec autoscaler. L'inférence et la
planification restent dans le cloud Cursor ; l'exécution (terminal,
fichiers, navigateur) se fait chez toi.

- Lien avec le point Coder (section 113.7) : même mouvement de fond —
  les agents cloud deviennent **déployables on-prem**, ce qui les rend
  envisageables dans ton contexte entreprise régulé.
- « Cursor Projects » : une bêta multi-agents évoquée dans des études
  tierces autour du 10-12/09/2026, **sans annonce officielle retrouvée
  au 27/09/2026** — à classer en *RUMEUR non confirmée* en l'état.

### 113.6 Kilo Code — rien d'officialisé pour la suite

Versions en cours : extension VS Code + CLI **v7.7.7** (22/09/2026),
stable JetBrains v7.1.6 (le CLI Kilo est un fork d'OpenCode, ~5 semaines
derrière l'upstream). Rappel : Anaconda a acquis Kilo en **juillet
2026**. **Aucune annonce officielle de roadmap ou de prochaine version
retrouvée au 27/09/2026** — rien d'officialisé.

### 113.7 Coder — Agent Relay (02/09/2026)

Annonce officielle **Coder × SpaceXAI** (GlobeNewswire, 02/09/2026) :
**Coder Agent Relay**, environnement d'exécution auto-hébergé pour
agents cloud — les Cursor Cloud Agents tournent dans des workspaces
Coder sur l'infrastructure du client (code, secrets, services internes
restent sur tes machines). C'est l'argument « entreprise régulée »
(banques, défense, secteur public) : auditabilité + contrôle réseau.
Coder Agents (agent natif auto-hébergé, agnostique de modèle) reste en
bêta.

### 113.8 OpenCode — Desktop en bêta

L'app **OpenCode Desktop** existe en **bêta** (macOS arm64/x64,
Windows x64, Linux deb/rpm/AppImage ; aussi via Homebrew/Scoop) :
multi-sessions parallèles, LSP chargés automatiquement, liens de
partage de sessions, login Claude Pro/Max, 75+ providers via
Models.dev. **Aucune roadmap officielle publiée retrouvée au
27/09/2026** — rien d'officialisé pour la suite.

### 113.9 OpenClaw — directions annoncées par Steinberger

Interview (août 2026, ~4,7 M de téléchargements) : trois axes de
roadmap énoncés par Peter Steinberger — (1) l'agent **toujours actif et
synchronisé** (problème présenté comme économique — les tokens — pas
comme un problème de modèle) ; (2) interfaces **voix et multimodalité**
(démo FaceTime en hack) ; (3) **team server** (sessions visibles entre
collègues, orchestration délégable), OpenClaw se reconstruisant
lui-même. Rappel : Steinberger a rejoint OpenAI (fév. 2026), OpenClaw
reste open source via une fondation. Sources : *BigGo Finance*
(interview) ; *Outlook Business*, fév. 2026.

### 113.10 Calendrier — OpenAI DevDay (29/09/2026)

**DevDay OpenAI prévu le 29/09/2026 à San Francisco**, avec l'IA
agentique annoncée comme axe majeur (Agents API entrée en bêta
publique début septembre, Agents SDK, infrastructure managée,
MCP au programme). Deux jours après la rédaction de ce guide :
c'est le prochain point de contrôle naturel pour mettre à jour
cette section.

### 113.11 Tableau récapitulatif (statut au 27/09/2026)

