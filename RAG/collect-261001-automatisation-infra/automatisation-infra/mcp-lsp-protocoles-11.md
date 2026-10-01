---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-11
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["mcp", "agent", "agents", "claude", "open source"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [1721, 1920]
sha256: ecdf2547de001bfa0ade462740e131a2076742af74af0266e0e55c771818cfcd
---

# MCP, LSP, code-server, websearch — Guide pratique

| Symptôme | Cause | Remède |
|---|---|---|
| Aucune complétion/diagnostic | Mauvais `root_dir` | Ouvre le **dossier projet**, pas un fichier isolé |
| Python : imports non résolus | Mauvais venv/interpréteur | Sélectionne l'interpréteur (`pyrightconfig.json` / `venvPath`) |
| Go : « no packages » | Hors module | `gopls` exige `go.mod` (ou GOPATH mode) |
| C++ : diagnostics absurdes | Pas de `compile_commands.json` | Génère-le (CMake `-DCMAKE_EXPORT_COMPILE_COMMANDS=ON`, bear sinon) |
| Rust : tout est rouge après update | Toolchain désync | `rustup update`, vérifie `rust-analyzer` aligné |
| Lenteur à l'ouverture | Indexation initiale | Patience + `exclude` les gros dossiers (`node_modules`, `.venv`) |
| Deux serveurs se battent | Extension + binaire manuel | Un seul : désactive l'un des deux |
| Neovim : `cmd not found` | Binaire pas sur le PATH | Chemin absolu dans `lsp/*.lua` |

## 85. Pourquoi les agents codeurs utilisent le LSP

Un agent qui lit le code « au texte » (grep) hallucine vite : il confond
deux fonctions homonymes, rate un import indirect, invente une signature.
Avec le LSP, l'agent obtient des **faits** :

- `definition` → « ce symbole est défini **ici**, ligne N » (fini les paris) ;
- `references` → « qui appelle cette fonction ? » (impact d'un refactoring) ;
- `hover` → signature + types réels (fini les arguments inventés) ;
- `diagnostics` → « ton patch ne compile pas, voici l'erreur exacte »
  → boucle d'auto-correction.

C'est pour ça que les frameworks d'agents 2026 intègrent le LSP en natif
(ex : extensions LSP pour agents, `mcp-language-server` vu au §53) :
**le LSP transforme le code en base de faits interrogeable**, et un agent
qui interroge des faits bat un agent qui devine du texte.

💡 Pour ton RAG : indexer du code sans LSP = chunks aveugles. Avec un
passage LSP (définitions + références), tes chunks portent la structure
réelle du code — la qualité de récupération grimpe.

## 86. LSP + MCP : la jonction

`mcp-language-server` (§53) enveloppe n'importe quel serveur LSP en tools
MCP. Résultat : dans Claude Code, `mcp__lsp__definition({file, line})`
retourne la vraie définition. Ton agent dispose alors des **deux**
protocoles de ce guide : MCP pour les actions/monde, LSP pour la
compréhension du code. C'est l'architecture des codeurs agentiques sérieux
en 2026.

## 87. Pense-bête LSP

- Un langage = un serveur = un binaire sur le PATH.
- Ouvre le dossier projet, jamais un fichier seul.
- `root_markers` (`pyproject.toml`, `go.mod`, `Cargo.toml`...) = comment le
  serveur trouve la racine. S'il se trompe, tout se trompe.
- Environnement Python : le serveur doit voir **le même venv** que ton code.
- `compile_commands.json` pour C/C++ : non négociable en projet réel.
- Debug : Output (VS Code) / `:LspLog` (Neovim) — le JSON-RPC ne ment pas.
- Agent + LSP = faits ; agent sans LSP = paris.

# PARTIE VI — CODE-SERVER : VS CODE DANS LE NAVIGATEUR

> Vérification faite : « code-serve » = bien **code-server**, le projet de
> Coder qui fait tourner VS Code (open source) sur un serveur distant,
> accessible depuis n'importe quel navigateur. Dépôt : `coder/code-server`
> (anciennement `codercom/code-server`).

## 88. code-server : c'est quoi, c'est pour quoi

code-server exécute VS Code sur une machine distante (ton serveur, un VPS,
un conteneur) et n'envoie que l'interface dans ton navigateur. Le code, les
extensions, le terminal : tout tourne **côté serveur**.

Cas d'usage :
- Développer sur une machine puissante depuis un PC léger / une tablette.
- Environnement de dev **identique** pour toute l'équipe (même image Docker).
- Accéder à ton IDE depuis n'importe où sans synchroniser de fichiers.
- Faire tourner l'agent codeur (Claude Code, etc.) sur le serveur, avec
  l'IDE en frontal.

Prérequis : hôte 64 bits, ~1 Go RAM mini (2+ recommandé), GLIBC 2.17+,
connexion HTTPS ou localhost (les service workers l'exigent).

## 89. Installation Linux (script officiel)

```bash
# Script officiel — installe le binaire + le service selon ta distro
curl -fsSL https://code-server.dev/install.sh | sh

# Lancement manuel (test)
code-server --bind-addr 127.0.0.1:8080 ~/projets

# Le mot de passe est généré dans ~/.config/code-server/config.yaml
cat ~/.config/code-server/config.yaml
```

Le script détecte apt/dnf/brew et installe le paquet correspondant.
Premier lancement : ouvre `http://127.0.0.1:8080`, saisis le mot de passe
du `config.yaml`. ⚠️ Change ce mot de passe généré si la machine est
partagée.

## 90. Installation Docker (la voie propre)

```bash
mkdir -p ~/.config
docker run -d --name code-server \
  --restart unless-stopped \
  -p 127.0.0.1:8080:8080 \
  -v "$HOME/.config:/home/coder/.config" \
  -v "$HOME/projets:/home/coder/projets" \
  -u "$(id -u):$(id -g)" \
  -e "DOCKER_USER=$USER" \
  -e "PASSWORD=ChangeMoi_Vite_123!" \
  codercom/code-server:latest
```

Notes :
- `-u "$(id -u):$(id -g)"` : les fichiers créés t'appartiennent, pas à root.
- `PASSWORD` : définit le mot de passe (sinon généré dans le volume config).
- 📌 Nom d'image : `codercom/code-server` (historique, toujours valide en
  2026 à ma connaissance — à vérifier sur Docker Hub le jour J).
- Bind sur `127.0.0.1` + reverse proxy TLS devant (§91). Jamais `0.0.0.0`
  sans TLS.

## 91. Configuration : config.yaml

`~/.config/code-server/config.yaml` (valeurs fictives — adapte) :

```yaml
bind-addr: 127.0.0.1:8080
auth: password
password: "un-vrai-mot-de-passe-long-et-unique"
cert: false
```

- `auth: password` (défaut) | `none` (⚠️ uniquement derrière un SSO/proxy
  qui authentifie déjà) | `authelia`/`openid` selon version.
- `cert: true` génère un autosigné (pratique en labo, pas en prod).
- En prod : `bind-addr: 127.0.0.1:8080` + TLS terminé par Caddy/nginx.

## 92. TLS : Caddy (simple) et nginx (contrôle)

**Caddy** (certificat Let's Encrypt automatique) :
```
# /etc/caddy/Caddyfile
ide.interne.example.com {
    reverse_proxy 127.0.0.1:8080
}
```
```bash
systemctl reload caddy   # c'est tout. Vraiment.
```

**nginx** :
```nginx
server {
  listen 443 ssl;
  server_name ide.interne.example.com;
  ssl_certificate /etc/letsencrypt/live/ide.interne.example.com/fullchain.pem;
  ssl_certificate_key /etc/letsencrypt/live/ide.interne.example.com/privkey.pem;
  location / {
    proxy_pass http://127.0.0.1:8080;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";   # websocket : indispensable
    proxy_set_header X-Forwarded-Proto $scheme;
  }
}
```

⚠️ Sans le `Upgrade` websocket, le terminal intégré et certaines extensions
ne fonctionnent pas (symptôme : IDE qui charge mais terminal mort).

## 93. Service systemd persistant

```ini
# /etc/systemd/system/code-server@.service
[Unit]
Description=code-server for %i
After=network.target

[Service]
Type=simple
User=%i
ExecStart=/usr/bin/code-server --bind-addr 127.0.0.1:8080 /home/%i/projets
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```
```bash
sudo systemctl enable --now code-server@zelef
```

💡 Avec Docker, `--restart unless-stopped` suffit. Avec le binaire, ce
service systemd + le reverse proxy = déploiement prod minimal.

## 94. Accès distant sécurisé : checklist

- [ ] Bind `127.0.0.1` + TLS via reverse proxy (jamais code-server nu sur Internet)
- [ ] Mot de passe long et unique (pas celui de ta session !)
- [ ] `auth: none` interdit sauf SSO en amont (Authelia, OAuth proxy...)
- [ ] Firewall : seul le reverse proxy écoute sur 443
- [ ] Fail2ban sur les 401 du code-server si exposé (ou mieux : ne pas exposer)
- [ ] VPN (WireGuard/Tailscale) : encore mieux que l'exposition publique —
      pour un usage perso/équipe, c'est la voie royale
- [ ] Mises à jour régulières (`latest` suivi ou script de maj mensuel)

