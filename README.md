# cetas-lite

Assistant IA auto-hébergé — **un seul binaire Go**, interface web intégrée, vos données chez vous.

cetas-lite réunit dans un binaire unique et multi-plateforme l'essentiel d'un poste de travail IA :
chat multi-modèles, mode agent avec outils, recherche web, mémoire, base documentaire locale et
coffre chiffré pour vos clés. Aucun service externe n'est requis.

© Marexsoft Corporation — Fondateur : Kouassi Marius. Tous droits réservés.

## Fonctionnalités

- **Chat** : streaming en temps réel, reconnexion sans perte, arrêt/reprise, sessions multiples,
  régénération, export Markdown/JSON.
- **Modèles** : fournisseurs cloud (API OpenAI-compatibles) et moteurs **locaux** (découverte
  automatique). Les *alias* masquent les modèles bruts et gèrent un **pool avec repli automatique**.
- **Mode agent** : lecture/écriture de fichiers, recherche dans le projet, exécution de commandes en
  **sandbox**, plan de tâches, approbations pour les actions sensibles.
- **Recherche web** : intégrée (native au fournisseur ou via outils), avec panneau de sources.
- **Mémoire** : pages Markdown persistantes, indexées et recherchables.
- **Base documentaire locale (RAG)** : vos documents sont indexés localement et injectés dans les
  réponses ; la recherche web est évitée quand la base suffit. Aucun appel réseau.
- **Multi-agents** : plusieurs agents en parallèle, chacun avec son fil, ses approbations et son
  espace de travail isolé.
- **Terminal intégré**, pièces jointes (images, PDF, texte/HTML), lecture et dictée via le navigateur.
- **Coffre chiffré** : les clés d'API sont stockées chiffrées au repos, jamais en clair.
- **Extensible** : serveurs MCP, outils HTTP personnalisés, plugins externes.

## Prérequis

- **Go 1.27.1 ou plus récent** — voir `go.mod` (`go 1.27.1`). Un Go ≥ 1.21 peut télécharger la
  toolchain requise automatiquement (nécessite un accès réseau).
- **make** (GNU Make). Sur Windows : WSL, ou `make` via Chocolatey/Scoop.
- **git** — optionnel : sert uniquement à dériver le numéro de version (`make build` retombe sur
  `dev` sans lui).

## Installation

### A. Binaire précompilé (recommandé)

Les binaires Linux / Windows / macOS (amd64 et arm64) sont publiés à chaque tag `v*` sur la page
**Releases** du dépôt.

```bash
# Linux/macOS — remplacer <os>-<arch> par la plateforme voulue (ex. linux-amd64)
chmod +x cetas-lite-<os>-<arch>
./cetas-lite-<os>-<arch> serve
```

Sur Windows, `cetas-lite-windows-amd64.exe` se lance par double-clic (ou `.\cetas-lite-windows-amd64.exe serve`).

> Si aucune release n'est encore publiée, utiliser la voie B (depuis les sources).

### B. Depuis les sources

```bash
git clone https://github.com/Hajrudin-Zelef/cetas-lite.git
cd cetas-lite
make build
```

## Vérifier

Lancer le serveur puis ouvrir l'interface :

```bash
# Linux/macOS
CETAS_LITE_HOME="$HOME/.cetas-lite" ./bin/cetas-lite serve
```

```powershell
# Windows (PowerShell)
$env:CETAS_LITE_HOME="$env:USERPROFILE\.cetas-lite"; .\bin\cetas-lite.exe serve
```

- Interface : <http://127.0.0.1:8787>
- Santé du service : <http://127.0.0.1:8787/api/health>

`make run` équivaut à `make build` + lancement.

## Premier démarrage

1. **Compte** : au premier lancement, le premier compte créé est le vôtre (mode *bootstrap*) ; ensuite
   l'inscription se ferme. Pour la rouvrir : `CETAS_LITE_REGISTRATION_OPEN=true`.
2. **Fournisseurs** : renseigner les clés d'API depuis l'interface (ou via la CLI `cetas-lite keys set`).
3. **Clés chiffrées** : les clés sont stockées chiffrées. Pour qu'elles soient déchiffrées au
   démarrage du serveur, définir `CETAS_LITE_VAULT_PASSWORD` (sinon seuls les moteurs locaux sont
   enregistrés).

## Application bureau (Windows)

```bash
make desktop
# -> bin/cetas-lite-desktop-windows-amd64.exe : double-clic, sans console.
```

Le mode bureau démarre le serveur en local et ouvre une fenêtre native **WebView2** ; la fermeture
de la fenêtre arrête proprement le serveur.

## Configuration (variables d'environnement)

| Variable | Rôle | Défaut |
|---|---|---|
| `CETAS_LITE_HOME` | Répertoire de données | `os.UserConfigDir()/cetas-lite` |
| `CETAS_LITE_ADDR` | Adresse d'écoute HTTP | `127.0.0.1:8787` |
| `CETAS_LITE_REGISTRATION_OPEN` | Inscription : non défini = *bootstrap* (1er compte puis fermé) · `true` = ouvert · `false` = fermé | *(bootstrap)* |
| `CETAS_LITE_TRUST_PROXY` | Faire confiance à `X-Forwarded-For`/`X-Real-IP` (derrière un reverse proxy) | `false` |
| `CETAS_LITE_SANDBOX` | Isolation des commandes : `none` · `auto` · `bwrap` | `auto` |
| `CETAS_LITE_ALLOW_SCRIPT` | Autoriser l'exécution de scripts par l'agent | `false` |
| `CETAS_LITE_RAG_DIR` | Dossier de la base documentaire locale (vide = désactivée) | `$CETAS_LITE_HOME/rag` |
| `CETAS_LITE_VAULT_PASSWORD` | Mot de passe déchiffrant les clés stockées | *(vide)* |
| `CETAS_PEPPER` | Secret serveur complémentaire du coffre | *(vide)* |
| `CETAS_LITE_OLLAMA_URL`, `CETAS_LITE_LMSTUDIO_URL`, `CETAS_LITE_LLAMACPP_URL` | Points d'accès des moteurs locaux | *(vide)* |

## Base documentaire locale (RAG)

Déposez vos documents dans `$CETAS_LITE_HOME/rag/` (ou `CETAS_LITE_RAG_DIR`) :

- **corpus structurés** : un sous-dossier avec un `manifest.json` (métadonnées et facettes) ;
- **fichiers bruts** : tout `.md`/`.txt` est indexé tel quel (un fichier = un passage).

L'index est construit **localement, en mémoire** — aucun appel réseau, aucune dépendance externe.
Dossier vide ⇒ fonctionnalité inactive (coût nul).

## Sécurité

- Authentification multi-utilisateurs (JWT) ; inscription contrôlée (*bootstrap* par défaut).
- Clés d'API **chiffrées au repos** ; jamais journalisées ni renvoyées en clair.
- Commandes de l'agent exécutées en **sandbox** (isolation système optionnelle) et confinées à
  l'espace de travail.
- Interface web servie localement ; aucun asset tiers n'est chargé par défaut.

## Build & distribution

```bash
make build     # binaire local -> bin/cetas-lite
make cross     # linux/windows/darwin x amd64/arm64 -> bin/
make desktop   # application bureau Windows
make test      # go test ./... -race
make ci        # gofmt + vet + tests
```

Binaire statique (`CGO_ENABLED=0`), sans dépendance d'exécution.

## Dépannage

- **`make: command not found`** → installer GNU Make (ou utiliser WSL sous Windows).
- **`go: go.mod requires go >= 1.27.1`** → mettre Go à jour (ou autoriser le téléchargement de la
  toolchain : `GOTOOLCHAIN=auto`).
- **Le port 8787 est déjà utilisé** → changer d'adresse : `CETAS_LITE_ADDR=127.0.0.1:8899 ./bin/cetas-lite serve`.
- **Mes clés d'API ne sont pas prises en compte** → définir `CETAS_LITE_VAULT_PASSWORD` avant de
  lancer le serveur.
- **Rien ne s'affiche dans le navigateur** → vérifier `http://127.0.0.1:8787/api/health` et les logs
  du serveur.

## Licence

Propriétaire — tous droits réservés © Marexsoft Corporation.
