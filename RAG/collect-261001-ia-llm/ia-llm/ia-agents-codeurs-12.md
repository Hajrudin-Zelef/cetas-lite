---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-12
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "ByteDance", "Google", "Lambda", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agents", "apache", "aws", "claude", "copilot", "gemini", "mcp", "open source", "qwen"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [1957, 2115]
sha256: 0a1b05aeedf157c8279ea5fe8a4102d427aa75733183867bdad565d5e215fc98
---

# Les agents codeurs IA

**Continue** (continue.dev) est une extension **open source (Apache 2.0)**
pour **VS Code et JetBrains** : chat, complétion, édition multi-fichiers, et
(un plus récent) **mode Agent** — le tout **BYOK** (n'importe quel modèle :
OpenAI, Anthropic, Ollama local…).

Installation :

1. Marketplace VS Code/JetBrains → **« Continue »**.
2. Au premier lancement : choisis ton fournisseur et colle ta clé (ou
   Ollama local pour du 100 % privé).
3. Config dans `~/.continue/config.yaml` (modèles, règles, MCP).

```yaml
# ~/.continue/config.yaml — extrait (schéma indicatif)
models:
  - name: "Mon modèle chat"
    provider: openrouter
    model: qwen/qwen3-coder
    apiKey: ${OPENROUTER_API_KEY}   # variable d'environnement, jamais en dur
tabAutocompleteModel:
  provider: ollama
  model: qwen2.5-coder:1.5b         # petit modèle local pour le Tab : gratuit et privé
```

Points forts : **reste sur ton VS Code stock** (pas de fork), 100 % OSS,
BYOK total, complétion locale possible (Ollama), MCP en mode Agent.
Faiblesses : moins « magique » que Cursor (pas d'indexation cloud poussée),
le mode Agent est plus jeune que Cline/Kilo, la qualité dépend du modèle
branché.

**Verdict** : le meilleur choix si ton cahier des charges dit « VS Code
officiel + open source + je choisis mon modèle ». Excellent couple avec un
petit modèle local pour la complétion (0 €, privé) et un gros modèle cloud
pour le chat/agent.

## 75. Supermaven : la complétion ultra-rapide

**Supermaven** (supermaven.com) est une extension de **complétion IA**
(VS Code, JetBrains, Neovim) réputée pour sa **vitesse** et sa **fenêtre de
contexte énorme** (1M tokens annoncés — elle « voit » très large pour des
suggestions pertinentes).

Repère important 2026 : **Supermaven a été racheté par Anysphere (Cursor)**
— sa technologie alimente désormais le Tab de Cursor. L'extension
standalone existe toujours (à vérifier sur supermaven.com — les offres
évoluent après un rachat).

Pour qui : tu veux **uniquement** de la complétion fulgurante, sans agent,
sans chat, sans changer d'éditeur. Le « métier » pur de la section 73.

## 76. GitHub Copilot : le standard historique

**GitHub Copilot** reste en 2026 **l'outil le plus déployé** (intégré à
VS Code, JetBrains, GitHub lui-même). Il a grandi : complétion + **chat** +
**mode agent** (avec choix parmi quelques modèles).

Tarifs indicatifs (sept. 2026 — à vérifier sur github.com/features/copilot) :

| Plan | Prix |
|---|---|
| Pro | **~10 $/mois** |
| Pro+ | **~39 $/mois** |
| Business / Enterprise | par siège (plus cher) |
| Gratuit | 2 000 complétions/mois + 50 requêtes premium/mois (étudiants/enseignants/mainteneurs OSS : gratuit) |

Points forts : intégration GitHub/VS Code la plus profonde du marché,
disponibilité entreprise (SSO, conformité), prix d'entrée bas, choix de
modèles (limité mais curé). Faiblesses : modèles imposés (pas de BYOK
généralisé), moins bon sur les longues sessions autonomes que les CLI
dédiés, facturation « requêtes premium » parfois confuse.

**Verdict** : si ton entreprise est standardisée GitHub, c'est le choix
évident et le moins cher. Pour un indépendant qui veut le meilleur agent,
les CLI des parties A vont plus loin.

## 77. Tabnine : la complétion « privacy-first »

**Tabnine** (tabnine.com) : complétion IA, 15+ éditeurs, argument historique
= **confidentialité** (modèles locaux, entraînement personnalisé sur ton
codebase sans qu'il quitte ton infra).

Pour qui : entreprise avec une politique « le code ne sort pas » stricte,
ou dev qui veut de la complétion sans cloud. En 2026, l'argument reste
valable mais la concurrence « locale » s'est étoffée (Ollama + Continue,
Tabby…).

## 78. Amazon Q Developer : le choix AWS

**Amazon Q Developer** (ex-CodeWhisperer) : assistant pour VS Code et
JetBrains — complétion, revue de code, génération de doc et de tests, et
sa spécialité : les **montées de version Java** et l'intégration profonde
**AWS** (il connaît les API AWS par cœur).

Tarif indicatif : **~19 $/utilisateur/mois** (Pro), avec un palier gratuit
limité (50 interactions chat/mois…). **À vérifier** sur aws.amazon.com/q.

Pour qui : tu vis dans AWS (CDK, Lambda, IAM…) — aucun autre assistant ne
connaît aussi bien cet écosystème. Sinon, passe ton chemin.

## 79. JetBrains AI : pour les IDE JetBrains

**JetBrains AI** : l'assistant intégré à **tous les IDE JetBrains**
(IntelliJ, PyCharm…). Si tu codes en Python sous PyCharm (ton cas pour les
scripts d'inventaire !), c'est la voie native : complétion, chat, refactos
qui exploitent l'indexation sémantique de l'IDE (meilleure compréhension du
typage que les extensions VS Code).

Note : **Junie** est l'agent autonome de JetBrains (évoqué dans les
comparatifs 2026) — si tu es « team JetBrains », le couple IDE + JetBrains
AI + Junie est cohérent.

## 80. Zed : l'éditeur rapide qui branche tout le monde

**Zed** (éditeur open source, écrit en Rust — rapidité légendaire) mérite sa
place : via l'**ACP (Agent Client Protocol)**, tu peux y brancher **n'importe
quel agent externe** (Claude Code, OpenCode…). Philosophie : le meilleur
éditeur possible + **tes** agents, pas ceux de l'éditeur.

Pour qui : tu veux la vitesse et l'ouverture, et tu as déjà un CLI agentique
préféré. « L'éditeur à surveiller » selon les comparatifs 2026.

## 81. Windsurf, Trae, Antigravity : les éditeurs challengers

- **Windsurf** : fork VS Code « agent-first » (anciennement Codeium). Supporte
  MCP (limite : ~100 outils). Positionné entre Cursor et VS Code+extensions.
- **Trae** (ByteDance) : assistant adaptatif, **offre gratuite généreuse**
  historiquement (accès à des modèles frontier), paliers payants agressifs
  (Lite ~3 $/mois, Pro ~10 $/mois — à vérifier sur trae.ai). Intéressant
  pour tester à moindre coût.
- **Antigravity** (Google) : la plateforme « agent-first » de Google
  (antigravity.google), workspace multi-agents. Paliers Pro ~19,99 $/mois,
  Ultra ~100 $/mois (à vérifier). À suivre si tu es dans l'écosystème Google.

Règle : **ne multiplie pas les éditeurs**. Teste un challenger 2 semaines
maximum, puis tranche — chaque changement d'éditeur coûte des semaines de
productivité.

## 82. Aider : le CLI git-first pour les puristes

**Aider** (aider.chat) : CLI agentique en **Python**, **open source**, pensé
« git-first » : il **committe** automatiquement chaque étape avec des
messages propres, fonctionne avec ta clé API ou des modèles locaux, et
tourne même **entièrement offline**.

Pour qui : le sysadmin/dev qui vit dans le terminal, aime git plus que tout,
et veut un agent **sobre, prévisible, privé**. Moins « magique » que Claude
Code, mais d'une fiabilité appréciable pour des tâches bien définies.
Excellent pour apprendre la boucle agentique côté terminal (pendant de
Cline côté éditeur).

```bash
pip install aider-chat
export ANTHROPIC_API_KEY="sk-ant-fictive"
cd /tmp/lab-agent
aider --model sonnet --message "ajoute une fonction de retry avec backoff exponentiel"
```

## 83. Gemini CLI, Kiro, Devin, Jules : les autres à connaître

