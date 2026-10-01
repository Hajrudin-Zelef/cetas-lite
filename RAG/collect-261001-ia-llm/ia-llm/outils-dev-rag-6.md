---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-6
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [1047, 1213]
sha256: 57214f28956097134b000eabf87eb1b39ce9bb463ccfa171d50bc28815e060e4
---

# Outils dev + ingénierie RAG (chunk & corpus)

```bash
git clone https://github.com/tmux-plugins/tpm ~/.tmux/plugins/tpm
# puis dans tmux : Ctrl+b I  (installer les plugins déclarés)
#                  Ctrl+b U  (mettre à jour)
```

2. **tmux-resurrect** — sauvegarde/restauration manuelle :
   - `Ctrl+b Ctrl+s` : **sauve** (fenêtres, panneaux, layouts, cwd, programmes)
   - `Ctrl+b Ctrl+r` : **restaure**
3. **tmux-continuum** — sauvegarde **auto toutes les 15 min** + restauration
   auto au démarrage de tmux.

**`.tmux.conf` complète et commentée** (`~/.tmux.conf`) :

```bash
# ── Base ─────────────────────────────────────────────────────
set -g default-terminal "screen-256color"   # couleurs correctes
set -g history-limit 50000                  # scrollback généreux
set -g mouse on                             # souris : scroll, sélection, resize
set -g base-index 1                         # fenêtres numérotées dès 1
setw -g pane-base-index 1
set -g renumber-windows on                  # renumérote après fermeture
setw -g mode-keys vi                        # mode copie façon vi

# ── Préfixe : on garde Ctrl+b (standard) ─────────────────────
# (certains passent à Ctrl+a ; ne change que si tu sais pourquoi)

# ── Splits plus intuitifs ────────────────────────────────────
bind | split-window -h -c "#{pane_current_path}"  # Ctrl+b | : vertical
bind - split-window -v -c "#{pane_current_path}"  # Ctrl+b - : horizontal
# -c "#{pane_current_path}" : le nouveau panneau hérite du dossier courant

# ── Recharger la config ──────────────────────────────────────
bind r source-file ~/.tmux.conf \; display "Config rechargée !"

# ── Synchronisation des panneaux (dangereux mais utile) ─────
bind S setw synchronize-panes \; display "Sync panneaux : #{?pane_synchronized,ON,OFF}"
# Quand ON, chaque frappe va dans TOUS les panneaux (ex. : apt sur N serveurs)

# ── Plugins (TPM) ────────────────────────────────────────────
set -g @plugin 'tmux-plugins/tpm'
set -g @plugin 'tmux-plugins/tmux-sensible'    # defaults sains
set -g @plugin 'tmux-plugins/tmux-resurrect'  # sauvegarde sessions
set -g @plugin 'tmux-plugins/tmux-continuum'  # sauvegarde auto
set -g @plugin 'tmux-plugins/tmux-yank'       # copie → presse-papiers

# ── Réglages resurrect / continuum ───────────────────────────
set -g @resurrect-capture-pane-contents 'on'  # sauve aussi le CONTENU des panneaux
set -g @continuum-restore 'on'                # restaure auto au démarrage
set -g @continuum-save-interval '15'          # minutes (défaut : 15)
# Programmes à restaurer tels quels (ex. : ne pas essayer de relancer vim à moitié) :
set -g @resurrect-processes 'ssh psql python3 tail htop'

# ── TPM en DERNIÈRE ligne (obligatoire) ──────────────────────
run '~/.tmux/plugins/tpm/tpm'
```

Après édition : `tmux source-file ~/.tmux.conf` ou `Ctrl+b r`.

## 31. Workflow « jobs longs » : tes scripts de scraping

```bash
# 1. Session dédiée, détachée dès le départ
tmux new -s collecte -d

# 2. Fenêtre par passe du pipeline
tmux new-window -t collecte -n "passe1" "python3 scripts/collect_failed.py --input listes/lot3.txt --out data/lot3_p1.jsonl 2>&1 | tee logs/lot3_p1.log"
tmux new-window -t collecte -n "passe2"

# 3. Détache-toi, éteins tout : ça tourne sur le serveur
# 4. Le lendemain :
tmux attach -t collecte
# Ctrl+b w → voir l'état de chaque fenêtre
# Si un worker a planté : relance-le dans sa fenêtre, l'historique est là
```

**Avec `tee` + log** : tu gardes la trace même si le panneau défile.
**Avec resurrect** : même après un reboot du serveur, `tmux attach`
retrouve ta mise en page (mais **pas** les processus en cours — voir §32).

Alternative sans tmux pour un job unique : `nohup python3 script.py > log 2>&1 &`
— mais tu perds l'interactivité et la mise en page. tmux gagne dès que tu
as 2+ choses à surveiller.

## 32. Pièges tmux (8)

1. **`tmux kill-server` par réflexe** : tue TOUTES les sessions, y compris
   ton scraping de 6 h. Préfère `tmux kill-session -t <nom>`.
2. **Resurrect ne restaure pas les processus en cours** (vérifié) : après un
   reboot, la mise en page revient mais `collect_still.py` ne redémarre pas
   tout seul. Relance explicite obligatoire.
3. **La première sauvegarde continuum arrive 15 min après le démarrage** :
   reboot dans les 15 premières minutes = rien à restaurer.
4. **La ligne `run '~/.tmux/plugins/tpm/tpm'` doit rester la DERNIÈRE**
   du `.tmux.conf`, sinon les plugins ne chargent pas.
5. **Bootstrap TPM sur une machine neuve** : si un vieux serveur tmux tourne
   avec une ancienne config, `~/.tmux/plugins/tpm/bin/install_plugins`
   échoue (« FATAL: Tmux Plugin Manager not configured »). Tue le serveur
   (`tmux kill-server`), relance, puis installe.
6. **Scrollback vs logs** : `history-limit 50000` aide, mais pour un job de
   6 h, redirige vers un fichier (`tee`) — le scrollback est volatil.
7. **`Ctrl+b` avalé par un programme** (ex. : dans `less`, `vim`) : appuie
   deux fois sur le préfixe pour l'envoyer au programme, ou utilise les
   commandes `tmux ...` depuis un autre shell.
8. **Copier-coller via SSH + tmux imbriqués** : si tu es en tmux local ET
   distant, le préfixe va au tmux local. Détache le local ou utilise
   `tmux -L distant` (socket distinct).

## 33. Pense-bête tmux

```bash
tmux new -s rag            # créer
tmux ls                    # lister
tmux attach -t rag         # rattacher
# Dans tmux :
# Ctrl+b d  détacher | Ctrl+b c  fenêtre | Ctrl+b % / "  splits
# Ctrl+b z  zoom | Ctrl+b [  copie | Ctrl+b ]  colle
# Ctrl+b Ctrl+s  sauver (resurrect) | Ctrl+b I  installer plugins
```

---

## 34. ngrok : le principe des tunnels

Un **tunnel** expose un service qui tourne sur ta machine (ou ton serveur)
vers Internet via une URL publique, **sans ouvrir de port** sur ta box ni
avoir d'IP publique. Schéma :

```text
Internet ──HTTPS──▶ edge ngrok ──tunnel chiffré (sortant)──▶ ngrok (ton PC)
                                                              └──▶ localhost:8000
```

La connexion est **sortante** : ça traverse NAT et CGNAT sans configuration
routeur. Cas d'usage : tester un webhook, partager une démo, exposer une API
locale à un service cloud, donner à un agent distant l'accès à ton service.

## 35. Installation de ngrok (Linux, vérifié)

```bash
# Méthode 1 : dépôt APT officiel (recommandé, vérifié sur ngrok.com/download)
curl -sSL https://ngrok-agent.s3.amazonaws.com/ngrok.asc \
  | sudo tee /etc/apt/trusted.gpg.d/ngrok.asc >/dev/null \
  && echo "deb https://ngrok-agent.s3.amazonaws.com bookworm main" \
  | sudo tee /etc/apt/sources.list.d/ngrok.list \
  && sudo apt update && sudo apt install ngrok
# ⚠ "bookworm" = Debian 12. Adapte (trixie pour Debian 13) — à vérifier.

# Méthode 2 : snap
sudo snap install ngrok

# Méthode 3 : binaire seul, sans sudo
curl -fsSL https://bin.ngrok.com/c/bNyj1mQVY4c/ngrok-v3-stable-linux-amd64.tgz | tar -xz
./ngrok version
```

```bash
# Authentification (une fois) — token sur dashboard.ngrok.com/get-started/your-authtoken
ngrok config add-authtoken "VOTRE_TOKEN_ICI"
```

> Le token est un secret : ne le mets ni dans un script versionné ni dans
> une capture d'écran. `ngrok config add-authtoken` le stocke dans
> `~/.config/ngrok/ngrok.yml`.

## 36. Premier tunnel : exposer un service local

