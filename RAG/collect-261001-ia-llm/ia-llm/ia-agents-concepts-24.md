---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-24
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Meta", "OpenAI", "Z.ai"]
dates: ["2025-12-18", "2026-09-11", "2026-09-13", "2026-09-21", "2026-09-23", "2026-09-27"]
keywords: ["agent", "agentic", "agents", "apache", "chatgpt", "claude", "gemini", "glm", "mcp", "muse", "open source", "valuation"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [3360, 3479]
sha256: 07ccea1ca91b46d2c8aaed116653cff95e6e969628d58f8b049180777cd6d236
---

# Concepts : agents IA, agentic, autonomie

- **Annonce** : OpenAI a publié dans son centre d'aide (11/09/2026, relayé
  par la presse le 11/09) le **retrait programmé des Custom GPTs le
  11 décembre 2026** (source : GSMDome, 4 jours avant le 27/09 ; Tablet News).
- **Remplacement** : migration vers **Plugins + Skills** (pas vers Projects).
  L'outil « Migrate to plugin » apparaît sous « My GPTs ».
- **Ce qui migre** : les instructions du GPT deviennent une **Skill** ; les
  fichiers de connaissances deviennent des **fichiers de référence** ; les
  applis connectées peuvent être incluses dans le Plugin. ChatGPT pourra
  sélectionner automatiquement une Skill quand une demande correspond à sa
  description.
- **Ce qui ne migre pas** : conversations existantes, conversation starters,
  modèle sélectionné, paramètres de partage, **custom actions** (à reconstruire
  — OpenAI évoque un serveur MCP custom).
- **Création de nouveaux Custom GPTs** : déjà coupée sur les comptes
  personnels (Free, Go, Plus, Pro) en sept. 2026 ; fin planifiée le
  **26 octobre** pour les workspaces Enterprise, dates « susceptibles de
  changer » selon OpenAI — à vérifier sur ton compte.
- **Si tu utilises un Custom GPT** : sauvegarde maintenant instructions,
  fichiers, prompts d'exemple (les brouillons non publiés ne migrent pas).

### 136.2. Meta Muse : adresse e-mail dédiée, Muse Charm, extensions

- **Adresse e-mail propre à Muse** : annoncée par Alexandr Wang (chief AI
  officer) à **Meta Connect 2026 (23/09/2026)**. Objet : donner à l'agent une
  identité opérationnelle (confirmations, tickets, relances). **Pas de date de
  sortie communiquée** — annoncé, pas encore disponible au 27/09/2026.
- **Muse Charm** : objet de poche comparé à un Tamagotchi, **en
  développement** (Connect 2026). Meta promet plus de détails « plus tard
  cette année » — pas de date, pas de prix : à vérifier.
- **Extensions géographiques et plateformes** : Canada ajouté après les
  États-Unis (Connect 2026). Extension aux autres apps du groupe (WhatsApp,
  Instagram) évoquée dans la presse comme étape logique du déploiement,
  **sans annonce datée au 27/09/2026** — à vérifier.
- **Rien d'autre d'officialisé** sur la feuille de route produit de Muse au
  27/09/2026 (pas de calendrier public de versions).

### 136.3. Open Interpreter : 0.0.43 (maintenance, visible fin sept. 2026)

- Les **release notes 0.0.43** sont visibles dans le dépôt
  `openinterpreter/openinterpreter` (fichier `RELEASE_NOTES.md`, consulté
  via recherche le 27/09/2026 ; **pas de date de publication vérifiée** dans
  les notes elles-mêmes).
- Contenu : maintenance release calée sur la base de compatibilité Codex
  `rust-v0.154.0` ; presets de modèles Google `gemini-3.x-flash` ;
  preset `zai-zcode` (endpoint Messages Z.AI) ; `interpreter acp` expose la
  session Harness.
- Rappel : le dépôt indiquait un dernier commit le 13/09/2026 (section 131).
  Le projet reste actif mais publie des maintenance releases, pas de
  refonte annoncée.

### 136.4. ZCode / Z.ai : open source le 21/09/2026, GLM-5.3 sorti

- **21 septembre 2026** : Z.ai **open-source ZCode** sous **Apache 2.0**
  (dépôt `zai-org/ZCode` : app desktop Electron, client web, Agent CLI,
  backend, runtime). Sources : Reuters (21/09/2026), RuntimeWire.
  Contexte : après un rapport (18/09) montrant que le client empaquetait des
  workspaces complets (dont l'historique Git) vers Alibaba Cloud OSS —
  correctifs en 3.14.0, excuses, invitation à des audits indépendants.
  Le dépôt public a démarré avec **2 commits seulement** (drop consolidé,
  pas l'historique de dev).
- **À venir (promis par Z.ai, sans date)** : publication du **rapport complet
  d'évaluation de sécurité** indépendante (Reuters, 21/09/2026) — « will be
  released soon ». Au 27/09/2026 : pas encore vu publiquement.
- **GLM-5.3** : sorti (Z.ai affirme qu'il « approche Anthropic Mythos » sur
  la détection de vulnérabilités), après **2 semaines de retard volontaire
  pour raisons de sécurité** — premier lab chinois à retarder explicitement
  une sortie de modèle pour la sûreté (Reuters).
- **Rien d'autre d'officialisé** sur la feuille de route ZCode/GLM au
  27/09/2026.

### 136.5. MCP : feuille de route 2026 (sans dates fermes)

- Publiée sur le blog officiel **modelcontextprotocol.io** (posts 2026) :
  **Streamable HTTP sans état** (scale horizontale), primitive **Tasks**
  (opérations longues asynchrones), **MCP Server Cards** (métadonnées
  `.well-known`), **registre centralisé** de serveurs (style npm),
  **auth intégrée SSO** et pistes d'audit.
- Statut : feuille de route officielle, **pas de calendrier de sortie
  ferme** au 27/09/2026 — à surveiller via le blog MCP et le dépôt de spec.
- Déjà acté (spec nov. 2025) : Streamable HTTP remplace SSE comme transport
  recommandé.

### 136.6. Google : ADK 2.0 / Genkit Agents API (sortis juillet 2026)

- **1er juillet 2026** : Google publie l'**Agents API** de **Genkit**
  (TypeScript et Go) — historique de messages, boucles d'outils, streaming,
  persistance, interrupts humain-dans-la-boucle, délégation multi-agents
  derrière une interface `chat()` unique (blog développeurs Google).
- **1er juillet 2026** : Google détaille **ADK 2.0** — séparation entre
  routage d'exécution déterministe (Workflows en graphe) et étapes LLM
  probabilistes, état reprenable (blog développeurs Google).
- **Rien d'officialisé** au 27/09/2026 sur la suite (ADK 3.0, Genkit GA)
  — pas de dates annoncées.

### 136.7. Anthropic : rien d'officialisé au 27/09/2026

- Aucune annonce de sortie à venir repérée au 27/09/2026 sur les skills
  (format ouvert depuis le 18/12/2025, cf. section 135), Claude Code ou
  l'Agent SDK. Le standard Agent Skills reste la trajectoire publique
  (gouvernance : Linux Foundation / Agentic AI Foundation).
- Rumeurs : **aucune rumeur crédible repérée et retenue** — cette section
  n'en invente pas. « RUMEUR non confirmée » : néant au 27/09/2026.

### 136.8. Comment maintenir cette section à jour

- **MCP** : blog `modelcontextprotocol.io` + dépôt de spec (roadmap 2026).
- **Open Interpreter** : `RELEASE_NOTES.md` du dépôt GitHub (commits/tags).
- **ZCode** : dépôt `zai-org/ZCode` (releases) + compte X officiel.
- **OpenAI** : centre d'aide (migration Custom GPTs → Plugins/Skills).
- **Meta Muse** : annonces Connect / blog Meta (adresse e-mail, Charm).
- Règle : une annonce n'entre ici que si elle est **officielle et datée** ;
  sinon elle reste dehors ou marquée « RUMEUR non confirmée ».

---

*Extension du 27/09/2026 — sections 134, 135, 136. Les prix, disponibilités*
*et dates marqués « à vérifier » évoluent vite : revérifie avant d'agir.*

## 137. Jev (TypeSafe AI) : la « decision layer » System One des agents

