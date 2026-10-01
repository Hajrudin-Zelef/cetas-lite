---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-2
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI", "OpenRouter"]
dates: []
keywords: ["agent", "agents", "claude", "open source", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [142, 341]
sha256: 2ca0e8dcfbbdc619b1b2285ab2e87e0cbe00c50d20cad82b628ceccf67b688be
---

# Les agents codeurs IA

Repère simple : **terminal = scriptable et CI-friendly**, **éditeur = confort
visuel et revues de diff**, **OpenClaw = agent qui vit en tâche de fond** et
que tu pilotes depuis ton téléphone.

## 5. Prérequis communs avant d'installer quoi que ce soit

Checklist à cocher une fois pour toutes :

- [ ] **Git installé et configuré** (`git --version` ≥ 2.40). Tous les agents
      travaillent mieux — et plus sûrement — dans un dépôt git.
- [ ] **Node.js 20+** (22+ recommandé pour Codex CLI) si tu installes via npm.
- [ ] **Un dossier de laboratoire** : `mkdir -p /tmp/lab-agent && cd /tmp/lab-agent && git init`.
- [ ] **Une clé API ou un abonnement** selon l'outil (détaillé par outil).
- [ ] **`.gitignore` correct** : aucun `.env`, aucune clé, aucun `id_rsa` ne
      doit traîner dans un dossier où un agent va lire/écrire.
- [ ] **Sauvegarde** : sur un vrai projet, commit propre AVANT de lancer un
      agent. `git status` doit être vide. C'est ton bouton « annuler ».

```bash
# Hygiène de départ : vérifie ce que tu as
git --version && node --version && npm --version && python3 --version

# Crée ton bac à sable
mkdir -p /tmp/lab-agent && cd /tmp/lab-agent && git init -q
cat > .gitignore <<'EOF'
.env
*.pem
*.key
id_rsa*
__pycache__/
node_modules/
EOF
git add .gitignore && git commit -qm "init lab"
echo "bac à sable prêt : $(pwd)"
```

## 6. Le fichier AGENTS.md : ton contrat avec tous les agents

Presque tous les outils de ce guide lisent `AGENTS.md` à la racine du projet
(Claude Code lit aussi `CLAUDE.md`, Kilo Code lit `AGENTS.md` en priorité).
C'est le fichier le plus rentable de tout ton outillage : 30 lignes bien
écrites valent des centaines d'euros de tokens gaspillés.

```markdown
# AGENTS.md — projet: inventaire-parc

## Contexte
App interne de gestion du parc (copieurs, onduleurs). Python 3.12, FastAPI,
PostgreSQL 16. Déploiement sur Debian 12 via Ansible.

## Commandes (à utiliser telles quelles)
- Installer : `pip install -r requirements.txt`
- Tests : `pytest -q`  (doivent passer avant tout commit)
- Lint : `ruff check . && ruff format --check .`
- Lancer en dev : `uvicorn app.main:app --reload`

## Conventions (obligatoires)
- Français pour les messages de commit et la doc, anglais pour le code.
- Pas de nouvelle dépendance sans validation humaine explicite.
- Les secrets passent par variables d'environnement, jamais en dur.
- Schéma SQL : toute migration via Alembic, jamais de SQL à la main en prod.

## Interdictions formelles
- Ne JAMAIS modifier `migrations/` déjà appliquées.
- Ne JAMAIS lancer `pytest` avec `--lf` seul : toujours la suite complète.
- Ne pas toucher à `legacy/` (code gelé, lecture seule).
```

Règles de rédaction :

1. **Impératif, court, vérifiable.** « Les tests doivent passer » > un
   paragraphe sur la qualité.
2. **Commandes copiables.** L'agent exécute bêtement ce que tu écris.
3. **Interdictions explicites.** L'agent ne devine pas tes tabous.
4. **Un AGENTS.md par périmètre** : racine (général) + sous-dossiers
   (`backend/AGENTS.md`, `ansible/AGENTS.md`) pour les règles locales.
5. **Relis-le tous les mois** : un AGENTS.md périmé fait faire n'importe quoi
   très efficacement.

> 💡 Pour ton RAG : un `AGENTS.md` versionné dans chaque repo documentaire te
> servira aussi de contrat quand tu demanderas à un agent de nettoyer ou
> d'indexer tes corpus.

## 7. Grille de lecture des fiches outils

Chaque outil est présenté avec le même plan pour comparer vite :

1. **C'est quoi / pour qui** — en deux phrases.
2. **Installation pas à pas** — commandes réelles, vérifiées sept. 2026.
3. **Premier workflow commenté** — un scénario complet, du prompt au résultat.
4. **Points forts / faiblesses honnêtes** — ce que la doc marketing ne dit pas.
5. **Tarifs vérifiés** — avec la mention « à vérifier » quand ça bouge.
6. **Verdict Zelef** — quand l'outil a du sens pour un sysadmin qui construit
   un RAG (avis assumé, pas sponsorisé).

## 8. Avertissement sécurité (à lire avant la première installation)

Un agent codeur, par construction :

- **lit tes fichiers** (y compris ceux que tu as oubliés : `.env`, dumps,
  clés) ;
- **exécute des commandes** (avec tes droits utilisateur, parfois sudo si tu
  le laisses faire) ;
- **envoie du contenu à un serveur distant** (le modèle) — ton code voyage.

Donc, avant d'installer :

1. 🔒 **Ne mets jamais de secret** dans un dossier de travail d'un agent
   (ni `.env` réel, ni clé API en dur, ni mot de passe). Utilise des `.env.example`.
2. 🔒 **Ne lance pas un agent en root** ni avec un utilisateur qui a les clés
   de la prod.
3. 🔒 **Vérifie le mode réseau** : certains agents ont un mode sandbox sans
   réseau ; active-le quand tu n'en as pas besoin.
4. 🔒 **Lis la politique de données** du fournisseur : ton code peut servir à
   entraîner des modèles (désactivable chez la plupart, mais pas partout —
   voir section 99).

La section 99 détaille tout ça. Si tu ne retiens qu'une phrase : **un agent
est un stagiaire brillant, rapide, et sans aucun sens des conséquences —
tu restes le responsable.**

---

# PARTIE A — LES CLI AGENTIQUES

---

## 9. OpenCode : c'est quoi, pour qui

**OpenCode** (opencode.ai) est un agent codeur **open source (MIT)**,
**agnostique du modèle** : au lieu d'être marié à un seul fournisseur
(Anthropic pour Claude Code, OpenAI pour Codex), il se branche sur **n'importe
quel fournisseur** via clé API — OpenAI, Anthropic, Google, OpenRouter,
Ollama en local, etc.

Pour qui :

- Tu veux **un seul outil** et changer de modèle selon la tâche (un pas cher
  pour explorer, un costaud pour coder).
- Tu veux du **100 % local** possible (Ollama) pour du code sensible.
- Tu refuses le verrouillage fournisseur et tu aimes l'open source
  (le projet dépasse les 95 000 stars GitHub — ordre de grandeur : énorme,
  très actif).
- Tu vis dans le terminal (TUI intégrée) mais tu veux aussi des plugins
  Neovim/VS Code.

Ce n'est PAS l'idéal si tu veux du « zéro configuration » : le choix du
modèle, c'est la liberté ET la responsabilité (qualité très variable selon le
modèle branché).

## 10. OpenCode : installation pas à pas

Prérequis : macOS, Linux ou Windows avec un terminal. Node.js utile mais le
script officiel gère tout.

```bash
# Méthode officielle (macOS / Linux) — vérifiée sept. 2026
curl -fsSL https://opencode.ai/install | bash

# Alternatives équivalentes
npm install -g opencode-ai
# ou : brew install opencode-ai
# Windows : npm install -g opencode-ai

# Vérifier
opencode --version
```

Premier lancement dans ton bac à sable :

```bash
cd /tmp/lab-agent
opencode
```

Au premier démarrage, OpenCode propose de **connecter un fournisseur** :

- **Option simple pour tester** : **OpenCode Zen**, l'offre hébergée
  d'OpenCode (démarrage gratuit selon la doc — à vérifier sur
  opencode.ai/docs, l'offre évolue).
- **Option BYOK** : dans la TUI, commande `/connect`, choisis le fournisseur
  (Anthropic, OpenAI, OpenRouter…) et colle ta clé API.
- **Option locale** : branche Ollama (`http://localhost:11434`) — aucun
  token facturé, qualité dépendant du modèle local.

La clé est stockée dans la config locale (`~/.config/opencode/`), pas dans
le projet. Ne la commite jamais.

Mettre à jour (le projet publie très souvent) :

```bash
# Même commande que l'installation, ou via le gestionnaire utilisé
curl -fsSL https://opencode.ai/install | bash
opencode --version
```

## 11. OpenCode : prise en main de la TUI

La TUI (Text User Interface) est le cœur de l'expérience :

