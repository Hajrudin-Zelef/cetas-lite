---
id: collect-261001-rattrapage/rattrapage/ssh-guide-12
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [2312, 2552]
sha256: 6d0503c29a5324847e1304be75e1afc26388b64ba0cc0486a4a928b24210182b
---

# Guide SSH approfondi

`/usr/local/bin/deploy.sh` :

```bash
#!/bin/bash
# N'autorise que la commande de déploiement attendue, rien d'autre.
set -euo pipefail
case "${SSH_ORIGINAL_COMMAND:-}" in
  "deploy production")
    logger -t deploy "Déploiement production lancé par $USER"
    exec /opt/app/scripts/deploy-production.sh
    ;;
  *)
    echo "Commande non autorisée." >&2
    logger -t deploy "Tentative non autorisée : ${SSH_ORIGINAL_COMMAND:-<vide>}"
    exit 1
    ;;
esac
```

**Côté CI** : la clé privée est un secret du runner (jamais dans le dépôt),
sans passphrase (usage non interactif) mais **restreinte** comme ci-dessus.
Test :

```bash
ci$ ssh -i deploy_key deployer@prod1 "deploy production"   # OK
ci$ ssh -i deploy_key deployer@prod1 "rm -rf /"            # refusé (command=)
ci$ ssh -i deploy_key deployer@prod1                       # pas de shell (no-pty)
```

Si cette clé fuit : l'attaquant ne peut exécuter que le déploiement, depuis le
réseau CI uniquement. C'est le pattern à généraliser pour **tous** les comptes
de service (backup, supervision, batch).

## 77. Cas pratique 3 : accès à une UI d'administration via tunnel

**Contexte** : interface iDRAC d'un serveur (`idrac-srv1`, 10.10.1.99:443),
inaccessible depuis le LAN admin. Accès via `srv1` (10.10.1.10) qui, lui, est
joignable en SSH.

```
Host srv1
    HostName 10.10.1.10
    User zelef
    ProxyJump bastion
    # Tunnel persistant vers l'iDRAC, via srv1
    LocalForward 8443 10.10.1.99:443
    ExitOnForwardFailure yes
```

```bash
client$ ssh -N -f srv1
client$ ss -tlnp | grep 8443     # vérifier que le tunnel écoute
# https://localhost:8443 -> iDRAC (avertissement certificat : normal, c'est
# l'autosigné de l'iDRAC, vérifie l'empreinte via la console physique la 1re fois)
```

**Pourquoi via `srv1` et pas en direct** : l'iDRAC n'est pas un bastion SSH
fiable pour faire du `ProxyJump` (son serveur SSH est minimaliste) ; en
revanche `srv1` relaie très bien le TCP. Le `LocalForward` est résolu **côté
srv1** : `10.10.1.99` est l'adresse vue depuis `srv1`.

**Hygiène** : tunnel à la demande (`ssh -N -f` quand tu en as besoin, `ssh -O
stop` après), jamais de `GatewayPorts` ici (l'iDRAC reste accessible
uniquement depuis ton poste). Pour un besoin d'équipe permanent, c'est un
bastion web authentifié (ou un VPN), pas un tunnel artisanal.

## 78. Cas pratique 4 : backup distant avec rsync over SSH

**Contexte** : sauvegarder `/data` de `prod1` vers `backup1` chaque nuit, avec
une clé restreinte en lecture seule.

**Sur `prod1`**, utilisateur `backup` + clé restreinte (pattern `rrsync`) :

```bash
serveur# apt install -y rsync   # fournit /usr/bin/rrsync
```

`/home/backup/.ssh/authorized_keys` :

```
from="10.0.6.20",command="/usr/bin/rrsync -ro /data/",no-pty,no-agent-forwarding,no-X11-forwarding,no-port-forwarding ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIEXEMPLEdefictifNePasUtiliser backup@backup1
```

`rrsync -ro /data/` : expose **en lecture seule** (`-ro`) uniquement `/data`.
Même avec la clé, impossible d'écrire sur `prod1` ou de lire ailleurs.

**Sur `backup1`**, script cron (`/usr/local/bin/pull-backup.sh`) :

```bash
#!/bin/bash
set -euo pipefail
SRC="backup@prod1:/data/"
DST="/srv/backups/prod1/$(date +%Y-%m-%d)/"
LOG="/var/log/pull-backup.log"
mkdir -p "$DST"
rsync -a --delete --link-dest=/srv/backups/prod1/latest \
  -e "ssh -i /root/.ssh/id_backup -o IdentitiesOnly=yes -o BatchMode=yes" \
  "$SRC" "$DST" >>"$LOG" 2>&1
ln -sfn "$DST" /srv/backups/prod1/latest
```

- `--link-dest` : déduplication par hard links (snapshots quotidiens peu
  coûteux, style rsnapshot).
- `BatchMode=yes` : le cron ne doit jamais bloquer sur un prompt.
- La clé privée est sur `backup1` uniquement, `600`, sans passphrase (compte
  de service), restreinte côté `prod1` comme ci-dessus.

## 79. Cas pratique 5 : jump host multi-sites

**Contexte** : 3 sites (Paris, Lyon, usine), chacun avec son bastion, un admin
nomade. Objectif : une config unique, des alias stables, du multiplexage pour
la fluidité.

```
# --- Bastions ---
Host bastion-paris
    HostName bastion.paris.example.com
    User zelef
Host bastion-lyon
    HostName bastion.lyon.example.com
    User zelef
Host bastion-usine
    HostName bastion.usine.example.com
    User zelef
    Port 2222

# --- Sites : motif par site ---
Host *.paris
    ProxyJump bastion-paris
Host *.lyon
    ProxyJump bastion-lyon
Host *.usine
    ProxyJump bastion-usine

# --- Défauts communs ---
Host *.paris *.lyon *.usine
    User zelef
    IdentityFile ~/.ssh/id_ed25519
    IdentitiesOnly yes
    ServerAliveInterval 60
    ControlMaster auto
    ControlPath ~/.ssh/sockets/%C
    ControlPersist 10m

# --- Alias métier ---
Host pve-paris
    HostName pve1.paris
Host sw-usine
    HostName 10.30.0.2.usine
    User admin
```

Usage : `ssh pve-paris`, `ssh sw-usine`. Le motif `*.usine` + `ProxyJump`
fait le reste. Le multiplexage (`ControlPersist 10m`) rend la 2e connexion
instantanée même à travers deux sauts — critique quand on administre 30
switches d'usine à la chaîne.

**Exploitation** : ce fichier est versionné (dépôt d'équipe), le bloc
« bastions » est généré par l'inventaire. Un nouveau site = 4 lignes.

---

# AIDE-MÉMOIRE

## 80. Pense-bête de poche : commandes

```bash
# Connexion de base
ssh zelef@web1
ssh -p 2222 zelef@web1
ssh -J bastion web1-interne              # via un bastion
ssh -J b1,b2 cible                       # deux sauts

# Diagnostic
ssh -v cible / ssh -vvv cible            # verbosité
ssh -G cible                             # config effective (LE réflexe)
sshd -T                                  # config serveur effective
sshd -t && systemctl reload ssh          # tester puis recharger

# Clés
ssh-keygen -t ed25519 -C "qui@ou-quand" -f ~/.ssh/id_ed25519
ssh-keygen -l -f cle.pub                 # empreinte
ssh-keygen -R hote                       # purger known_hosts
ssh-keygen -p -f ~/.ssh/id_ed25519       # changer la passphrase
ssh-copy-id -i cle.pub user@hote

# Agent
eval "$(ssh-agent -s)" ; ssh-add -t 8h ~/.ssh/id_ed25519
ssh-add -l / -d cle / -D
ssh -O check cible / ssh -O stop cible   # multiplexage

# Tunnels
ssh -L 8080:localhost:9090 web1          # local  : poste -> serveur
ssh -R 9000:localhost:8000 web1          # distant: serveur -> poste
ssh -D 1080 bastion                      # SOCKS
ssh -N -f -L 8080:localhost:9090 web1    # tunnel en fond, sans shell
autossh -M 0 -N -f -L ...                # tunnel persistant auto-réparé

# Transferts
scp -r ./dir web1:/tmp/
sftp web1
rsync -avz --progress -e "ssh -J bastion" ./dir/ cible:/dest/
rsync -avz --delete --dry-run ./dir/ cible:/dest/   # simuler avant --delete

# Serveur : qui est là / qui a essayé
who ; last -n 20 ; lastb -n 20
ss -tnp | grep :22
grep -E "Accepted|Failed" /var/log/auth.log | tail
journalctl -u ssh -f
fail2ban-client status sshd
```

## 81. Pense-bête : `ssh_config` minimal propre

```
# ~/.ssh/config — squelette recommandé
Include ~/.ssh/config.d/*.conf

Host *
    ServerAliveInterval 60
    ServerAliveCountMax 3
    ConnectTimeout 10
    HashKnownHosts yes
    StrictHostKeyChecking accept-new
    IdentitiesOnly yes
    ForwardAgent no
    ForwardX11 no
    GSSAPIAuthentication no
    AddKeysToAgent yes
    ControlMaster auto
    ControlPath ~/.ssh/sockets/%C
    ControlPersist 10m
    ExitOnForwardFailure yes
    LogLevel INFO
```

Avec `mkdir -p ~/.ssh/sockets && chmod 700 ~/.ssh && chmod 600 ~/.ssh/config`.
Tout le reste (bastions, alias, tunnels) va dans `config.d/`.

## 82. Pense-bête : `sshd_config` durci minimal

