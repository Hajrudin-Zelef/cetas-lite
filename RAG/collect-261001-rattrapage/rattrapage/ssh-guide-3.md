---
id: collect-261001-rattrapage/rattrapage/ssh-guide-3
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [398, 617]
sha256: 16ee42e07badb82a25b1ea31829d388cb57059f9f733a384556012f1855e60d0
---

# Guide SSH approfondi

- `HostName` accepte un nom DNS, une IPv4 ou IPv6.
- `User` : l'utilisateur distant. Ne mets jamais `root` ici par défaut ;
  crée un bloc dédié explicite si tu en as besoin (avec un commentaire).
- `Port` : utile quand le DSI impose un port non standard ou quand plusieurs
  sshd tournent derrière une même IP (NAT).

Astuce : `%h` et `%p` dans `ProxyCommand`/`ProxyJump` reprennent l'hôte et le
port calculés — pratique pour les modèles génériques (section 27).

## 12. `IdentityFile` : une clé par hôte

```
Host web1
    HostName 192.0.2.10
    IdentityFile ~/.ssh/id_ed25519_web

Host github-perso
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_ed25519_perso

Host *
    IdentitiesOnly yes
```

- `IdentityFile` peut être répété : SSH essaie chaque clé dans l'ordre.
- **`IdentitiesOnly yes`** : n'essaie **que** les clés configurées (pas celles
  de l'agent). Sans ça, SSH propose toutes les clés de l'agent, ce qui peut
  déclencher `Too many authentication failures` (section 67) et divulgue
  quelles clés tu possèdes.
- Séparer les clés par usage (perso / pro / client X) limite l'impact d'une
  compromission et clarifie l'inventaire (section 59).

## 13. Options client par hôte : les plus utiles

```
Host web1
    HostName 192.0.2.10
    # Rester connecté à travers les NAT/firewalls qui coupent les sessions idle
    ServerAliveInterval 60      # ping toutes les 60 s
    ServerAliveCountMax 3       # 3 échecs -> coupe proprement
    # Ne pas transmettre l'agent par défaut (voir section 20)
    ForwardAgent no
    # Ne pas faire suivre X11 par défaut
    ForwardX11 no
    # Compression pour les liens lents (pas en LAN)
    Compression yes
    # Timeout de connexion : ne pas attendre 2 minutes sur un hôte mort
    ConnectTimeout 10
    # Journaliser dans un fichier dédié pour cet hôte sensible
    LogLevel VERBOSE
```

Tableau des options client les plus rentables :

| Option | Effet | Quand l'utiliser |
|---|---|---|
| `ServerAliveInterval` / `ServerAliveCountMax` | keepalive | Toujours (NAT, Wi-Fi) |
| `ConnectTimeout` | délai max de connexion | Scripts, inventaires |
| `IdentitiesOnly yes` | limite les clés essayées | Toujours avec plusieurs clés |
| `StrictHostKeyChecking accept-new` | TOFU strict | Scripts non interactifs |
| `BatchMode yes` | jamais de prompt | Cron, Ansible |
| `LogLevel` | verbosité | `VERBOSE` pour auditer |
| `AddKeysToAgent yes` | ajoute la clé à l'agent à la 1re utilisation | Confort + passphrase (section 18) |

## 14. `Match`, `Include` et organisation du fichier

`Match` applique un bloc selon des **critères** (pas seulement le nom) :

```
# Appliquer à tous les hôtes sauf le bastion lui-même
Match host !bastion
    ProxyJump bastion

# Options différentes selon le réseau d'origine
Match originalhost 192.168.*.*
    StrictHostKeyChecking no
    UserKnownHostsFile /dev/null
```

`Include` découpe la config en fichiers (pratique en équipe) :

```
Include ~/.ssh/config.d/*.conf
```

Organisation recommandée :

```
~/.ssh/
├── config                  # squelette : Include + Host * global
├── config.d/
│   ├── 00-defaults.conf    # défauts globaux
│   ├── 10-perso.conf       # machines perso
│   ├── 20-prod.conf        # production (généré/déployé par Ansible)
│   └── 30-clients.conf     # un fichier par client
└── known_hosts
```

Avantage : le fichier `20-prod.conf` peut être **déployé et versionné** par
l'équipe, sans toucher aux alias personnels de chacun.

## 15. Générer une clé ed25519

Ed25519 est le choix par défaut moderne : 256 bits, rapide, clés courtes.

```bash
client$ ssh-keygen -t ed25519 -C "zelef@poste-prod-$(date +%Y-%m-%d)" -f ~/.ssh/id_ed25519
Generating public/private ed25519 key pair.
Enter passphrase (empty for no passphrase): ********
Enter same passphrase again: ********
Your identification has been saved in /home/zelef/.ssh/id_ed25519
Your public key has been saved in /home/zelef/.ssh/id_ed25519.pub
The key fingerprint is:
SHA256:AbCdEfGhIjKlMnOpQrStUvWxYz0123456789ABc zelef@poste-prod-2026-09-26
```

- `-C` : commentaire = **traçabilité** (qui, quelle machine, quand). Toujours
  le renseigner : dans 2 ans, tu sauras à quoi correspond cette clé.
- `-f` : nom explicite, surtout si tu gères plusieurs clés.
- **Permissions** : `600` sur la privée (`ssh-keygen` le fait), `644` sur la
  publique. SSH refuse une clé privée trop permissive (section 68).

Vérifier une clé existante :

```bash
client$ ssh-keygen -l -f ~/.ssh/id_ed25519.pub
256 SHA256:AbCdEfGhIjKlMnOpQrStUvWxYz0123456789ABc zelef@poste-prod-2026-09-26 (ED25519)
client$ ssh-keygen -y -f ~/.ssh/id_ed25519   # régénère la publique depuis la privée
```

## 16. RSA 4096 : quand s'en servir encore

```bash
client$ ssh-keygen -t rsa -b 4096 -C "zelef@legacy" -f ~/.ssh/id_rsa_4096
```

Garde RSA 4096 pour :

- les **vieux équipements** (switches, iLO/iDRAC anciens, appliances) qui ne
  connaissent pas ed25519 ;
- les **vieux serveurs** (OpenSSH < 6.5 côté serveur, rare aujourd'hui).

Règles :

- **4096 bits minimum** (2048 est le plancher tolérable, 1024 = à détruire).
- Ne génère plus de DSA (`ssh-dss`) ni d'ECDSA : DSA est désactivé par défaut
  depuis OpenSSH 7.0, ECDSA n'apporte rien face à ed25519.
- Si un équipement n'accepte que RSA 1024 ou DSA : c'est un **signal d'alerte**
  (firmware à mettre à jour, équipement à remplacer), pas une config à banaliser.

## 17. Passphrase : la choisir et la gérer

Une clé **sans passphrase** = quiconque lit le fichier se connecte partout où
la clé est autorisée. Sur un poste de travail, c'est inacceptable.

- Mets **toujours** une passphrase sur les clés d'un poste (portable++).
- Exceptions légitimes (clés **sans** passphrase) : comptes de service pour
  scripts/cron — mais alors clé **restreinte** (`command=`, `from=`, section 24)
  et jamais sur un poste utilisateur.
- Changer la passphrase sans changer la clé :

```bash
client$ ssh-keygen -p -f ~/.ssh/id_ed25519
```

- Une passphrase n'a pas besoin d'être compliquée à taper 50 fois par jour :
  c'est `ssh-agent` (section 18) qui la mémorise pour la session. Choisis-la
  **longue** (4-5 mots), pas « complexe ».

## 18. `ssh-agent` et `ssh-add`

L'agent garde les clés déchiffrées en mémoire : tu tapes la passphrase **une
fois** par session, plus jamais ensuite.

```bash
client$ eval "$(ssh-agent -s)"
Agent pid 1234
client$ ssh-add ~/.ssh/id_ed25519
Enter passphrase for /home/zelef/.ssh/id_ed25519: ********
Identity added: /home/zelef/.ssh/id_ed25519 (zelef@poste-prod-2026-09-26)
client$ ssh-add -l        # clés chargées
256 SHA256:AbCd... (ED25519)
client$ ssh-add -t 8h ~/.ssh/id_ed25519   # durée de vie limitée : bonne hygiène
client$ ssh-add -d ~/.ssh/id_ed25519      # retirer une clé
client$ ssh-add -D                        # tout vider (verrouillage de poste)
```

Bonnes pratiques :

- `-t <durée>` : limite la durée de vie en mémoire (ex. `8h` = journée de
  travail). Si le poste est compromis, la fenêtre d'abus est bornée.
- `AddKeysToAgent yes` dans `ssh_config` : la première utilisation d'une clé
  l'ajoute à l'agent automatiquement (avec `AddKeysToAgent 2h` pour la durée).
- `ssh-add -D` en verrouillant ton poste le soir, ou via le script de
  verrouillage d'écran.

## 19. Lancer l'agent au démarrage

**Linux + systemd (user service)** — `~/.config/systemd/user/ssh-agent.service` :

```ini
[Unit]
Description=SSH key agent

[Service]
Type=simple
Environment=SSH_AUTH_SOCK=%t/ssh-agent.socket
ExecStart=/usr/bin/ssh-agent -D -a $SSH_AUTH_SOCK

[Install]
WantedBy=default.target
```

```bash
client$ systemctl --user enable --now ssh-agent
# Dans ~/.bashrc ou ~/.profile :
export SSH_AUTH_SOCK="$XDG_RUNTIME_DIR/ssh-agent.socket"
```

