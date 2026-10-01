---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-5
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "copilot", "embedding"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [831, 1046]
sha256: dfe06dab2ade960d0b7d228c8186b95af93f2261106270b0d9c6f37589288d62
---

# Outils dev + ingénierie RAG (chunk & corpus)

```bash
git switch -c feat/nouveau-chunker
# ... code + tests ...
git push -u origin feat/nouveau-chunker
gh pr create --title "feat : chunker sémantique" --body "Tests OK en local"
# la CI tourne ; tu merges quand c'est vert
```

**Badge de statut** en tête du README (optionnel mais motivant) :

```markdown
![CI](https://github.com/<ton-user>/rag-perso/actions/workflows/ci.yml/badge.svg)
```

## 22. Codespaces : un VS Code complet dans le cloud

**C'est quoi (vérifié sept 2026) :** un environnement de dev hébergé par
GitHub = un **dev container** qui tourne sur leurs serveurs, accessible
depuis le navigateur ou ton VS Code local. Même `devcontainer.json` qu'en
local, aucune config parallèle.

**Quotas vérifiés (sept 2026) :**
- Compte **Free** : **60 core-hours/mois** + 15 Go de stockage.
- Compte **Pro** : 90 core-hours + 20 Go.
- Au-delà : facturation au core-hour (2-core : $0.18/h, 4-core : $0.36/h…).
- Limite de dépenses réglable à **$0** (anti-facture-surprise).

`.devcontainer/devcontainer.json` pour ton projet :

```jsonc
{
  "name": "rag-perso",
  "image": "mcr.microsoft.com/devcontainers/python:3.12",
  "features": {
    "ghcr.io/devcontainers/features/github-cli:1": {},
    "ghcr.io/devcontainers/features/postgres-client:1": {}
  },
  "customizations": {
    "vscode": {
      "extensions": [
        "ms-python.python",
        "eamodio.gitlens",
        "github.copilot"
      ],
      "settings": {
        "python.defaultInterpreterPath": "/usr/local/bin/python"
      }
    }
  },
  "postCreateCommand": "pip install -r requirements.txt",
  "forwardPorts": [8000],
  "portsAttributes": {
    "8000": { "label": "API RAG locale", "onAutoForward": "notify" }
  }
}
```

**Quand ça sert pour TOI :**
- ✅ Tester une PR ou une idée depuis une tablette / un PC qui n'est pas le tien.
- ✅ Onboarding : quelqu'un clone, clique *Code → Codespaces → Create*,
  tout est installé via `postCreateCommand`.
- ✅ Expérimenter sans salir ton serveur (environnement jetable).
- ❌ **Pas** pour tes gros traitements de corpus : 60 h/mois fondent vite
  avec un embedding de 100k chunks, et le stockage est limité. Tes
  `collect_*` restent sur ton serveur.

**Bonnes pratiques :** `prebuilds` (repo → Settings → Codespaces) pour un
démarrage < 1 min ; **supprime** les codespaces inactifs (ils consomment
du stockage) ; ne mets jamais de secrets dans `devcontainer.json`
(utilise les *Codespaces secrets* du compte).

## 23. Pense-bête GitHub

```bash
gh repo create rag-perso --private --source=. --push  # créer + pousser
gh issue list --label collecte                        # suivre les collectes
gh secret set OPENAI_API_KEY                         # stocker une clé
gh pr create --title "feat : ..."                     # proposer une PR
gh run list --workflow=ci.yml                        # voir les runs CI
git lfs track "*.zip" && git add .gitattributes       # gros fichiers
```

- [ ] Repo privé créé, `.gitignore` + `.gitattributes` committés
- [ ] `ci.yml` vert sur le premier push
- [ ] Branch protection sur `main`
- [ ] Aucun secret dans l'historique (`git log -S "sk-"`)
- [ ] Template d'issue « collecte » créé
- [ ] `devcontainer.json` testé une fois (Codespaces)

---

## 24. tmux : pourquoi c'est vital pour tes jobs longs

`tmux` = multiplexeur de terminal : plusieurs shells dans **une** connexion
SSH, qui **survivent à la déconnexion**. Ton cas d'usage exact :

```text
Sans tmux :  tu lances collect_still.py (4 workers, 6 h) en SSH…
             …ta connexion coupe à 23h… le script MEURT.
Avec tmux :  tu lances dans une session, tu détaches (Ctrl+b d),
             tu fermes ton laptop. Le script continue sur le serveur.
             Le lendemain : tu rattaches, tout est là.
```

Les 3 niveaux : **session** (un projet) → **fenêtre** (un onglet)
→ **panneau** (split d'écran). Le **préfixe** par défaut est `Ctrl+b`
(suivi d'une autre touche).

## 25. Installation de tmux

```bash
# Debian/Ubuntu
sudo apt update && sudo apt install -y tmux

# Vérifier la version (3.x requis pour les popups / plugins modernes)
tmux -V
# tmux 3.4  ← exemple

# macOS (ton poste éventuel)
brew install tmux
```

> **À vérifier** sur ton serveur : si la version packagée est < 3.2,
> certains plugins (extrakto, popups) ne marcheront pas — compile depuis
> les sources ou prends un backport.

## 26. Sessions : créer, lister, rattacher, tuer

```bash
tmux new -s rag                    # crée la session "rag" et s'y attache
tmux new -s collecte -d            # crée en détaché (sans s'attacher)
tmux ls                            # liste les sessions
tmux attach -t rag                 # rattache la session "rag"
tmux attach -t rag -d              # rattache en détachant les autres clients
tmux kill-session -t rag           # tue UNE session
tmux kill-server                   # tue TOUT (attention !)
```

Dans tmux : `Ctrl+b d` = **detach** (détacher, la session continue).
`Ctrl+b s` = menu de choix de session. `Ctrl+b $` = renommer la session.

**Convention de nommage pour toi :**

```text
rag        → dev courant (éditeur + tests)
collecte   → jobs de scraping longs (collect_*.py)
db         → psql + monitoring
```

## 27. Fenêtres : des onglets dans la session

| Action | Raccourci | Commande |
|---|---|---|
| Nouvelle fenêtre | `Ctrl+b c` | `tmux new-window -t rag` |
| Renommer | `Ctrl+b ,` | `tmux rename-window scraper` |
| Suivante / précédente | `Ctrl+b n` / `Ctrl+b p` | |
| Aller à la n°3 | `Ctrl+b 3` | `tmux select-window -t :3` |
| Fermer | `Ctrl+b &` (confirme) | `exit` dans le shell |
| Liste des fenêtres | `Ctrl+b w` | |

```bash
# Créer une fenêtre nommée qui lance directement une commande
tmux new-window -t collecte -n "lot3" "python3 scripts/collect_still.py --input listes/lot3.txt"
```

## 28. Panneaux : splitter l'écran

| Action | Raccourci |
|---|---|
| Split horizontal (haut/bas) | `Ctrl+b "` |
| Split vertical (gauche/droite) | `Ctrl+b %` |
| Naviguer entre panneaux | `Ctrl+b` + flèches (ou `o` pour cycler) |
| Zoom plein écran / dézoomer | `Ctrl+b z` |
| Fermer le panneau | `Ctrl+b x` (confirme) |
| Rééquilibrer | `Ctrl+b` `Espace` (cycle les layouts) |

```bash
# En ligne de commande (utile dans des scripts)
tmux split-window -h -t rag        # split vertical
tmux split-window -v -t rag        # split horizontal
tmux select-pane -t rag:0.1       # focus panneau 1 de la fenêtre 0
```

**Layout type « collecte » :**

```text
┌──────────────────────┬──────────────────────┐
│ collect_still.py     │ tail -f logs/        │
│ (4 workers, tourne)  │ collecte_lot3.log    │
├──────────────────────┼──────────────────────┤
│ htop                 │ psql : vérif chunks  │
└──────────────────────┴──────────────────────┘
```

## 29. Copier-coller dans tmux

```text
1. Ctrl+b [          → mode copie
2. Espace            → début de sélection (ou 'v' si mode vi, voir §30)
3. Flèches / PgUp    → étendre la sélection
4. Entrée            → copier dans le buffer tmux
5. Ctrl+b ]          → coller
```

Avec la souris activée (`set -g mouse on` dans `.tmux.conf`, voir §30),
sélectionner à la souris copie directement. Pour le presse-papiers système
sur un serveur distant, le plus fiable reste : sélection → `Shift`+sélection
du terminal (bypass tmux) → `Ctrl+Shift+C`.

## 30. Persistance : tmux-resurrect + tmux-continuum (vérifié)

**Le problème :** un reboot du serveur tue le serveur tmux → sessions perdues.
**La solution (vérifiée, plugins maintenus en 2026) :**

1. **TPM** (gestionnaire de plugins) :

