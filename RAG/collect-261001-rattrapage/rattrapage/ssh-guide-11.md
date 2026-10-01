---
id: collect-261001-rattrapage/rattrapage/ssh-guide-11
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [2114, 2311]
sha256: ae1b303e3f0c4c923afccab1536ff107f6f194779907994a8bd2b1d5991090f0
---

# Guide SSH approfondi

- La **publique** n'existe pas : régénère-la depuis la privée
  (`ssh-keygen -y -f ~/.ssh/id_ed25519 > ~/.ssh/id_ed25519.pub`), ou pointe
  `-i` vers la bonne.
- `Permission denied` pendant le `ssh-copy-id` : normal si le serveur
  n'accepte que les clés et que tu n'en as pas encore déployé — c'est
  l'œuf et la poule. Solutions : mot de passe temporaire (puis le couper),
  console (iDRAC/IPMI), ou déploiement par Ansible/cloud-init.
- Le serveur refuse l'écriture : home en lecture seule, quota, ou
  `AuthorizedKeysFile` pointant ailleurs (`sshd -T | grep authorizedkeysfile`).
- Doublons : `ssh-copy-id` ne déduplique pas ; relancé 3 fois = 3 lignes
  identiques. Nettoie avec `sort -u` (en vérifiant avant d'écraser).

## 70. Cas 9 : le tunnel ne répond pas / port déjà utilisé

```
bind [127.0.0.1]:8080: Address already in use
channel_setup_fwd_listener_tcpip: cannot listen to port: 8080
```

1. **Port déjà pris** : un autre tunnel, ou l'appli elle-même. `ss -tlnp |
   grep 8080` pour identifier le squatteur.
2. **Tunnel « démarré » mais mort** : sans `ExitOnForwardFailure yes`, SSH
   continue sans le tunnel si le bind échoue → tu crois que ça marche. Ajoute :

```
Host tunnel-*
    ExitOnForwardFailure yes
```

3. **La destination vue par le serveur est fausse** : dans
   `-L 8080:localhost:9090`, `localhost` est résolu **par le serveur**. Si le
   service écoute sur une autre interface du serveur, utilise son IP ou
   `127.0.0.1` explicite.
4. **`AllowTcpForwarding no`** côté serveur (section 46) : le tunnel est
   refusé silencieusement ou avec `open failed: administratively prohibited`.
   Vérifie `sshd -T | grep allowtcpforwarding` et les blocs `Match`.
5. **Tunnel `-R` inaccessible** : `GatewayPorts` (section 31).

## 71. Cas 10 : multiplexage, `ControlPath` trop long

```
muxserver_listen: link mux listener: path too long
```

Les sockets Unix ont une limite de ~107 caractères. `~/.ssh/sockets/%r@%h:%p`
avec un home profond + un FQDN long la dépasse.

Solutions :

```
Host *
    # %C = hash court de (user, host, port) : compact et unique
    ControlPath ~/.ssh/sockets/%C
    # ou, radical : /tmp
    # ControlPath /tmp/ssh-%C
```

- Vérifie que `~/.ssh/sockets/` existe (`mkdir -p`).
- `%C` évite aussi les caractères spéciaux des FQDN dans le nom de fichier.
- Diagnostic : `ssh -O check cible` → `Master running` ou message d'erreur
  explicite.

## 72. Cas 11 : X11 forwarding sans `DISPLAY`

Symptôme : `ssh -X web1`, puis `xclock` → `Error: Can't open display:`.

1. **Côté client** : `echo $DISPLAY` **avant** ssh doit être non vide. En SSH
   depuis un terminal sans X (ou Wayland sans XWayland), il n'y a rien à
   forwarder.
2. **`ForwardX11 yes`** actif ? `ssh -G web1 | grep -i forwardx11`.
3. **Côté serveur** : `X11Forwarding yes` dans `sshd_config` + paquet
   `xauth` installé (`command -v xauth`). Sans `xauth`, sshd ne peut pas
   créer le cookie d'autorisation et le forwarding échoue silencieusement.
4. Après connexion, `echo $DISPLAY` doit donner quelque chose comme
   `localhost:10.0`. Si vide : relis les points 2-3, et `ssh -v` (ligne
   `Requesting X11 forwarding` / `X11 forwarding request failed`).

Rappel section 40 : `-X` plutôt que `-Y`, et jamais en `Host *`.

## 73. Cas 12 : SFTP en chroot qui échoue

Config type qui échoue mystérieusement :

```
Match Group sftp-only
    ChrootDirectory /srv/sftp/%u
    ForceCommand internal-sftp
```

Checklist du chroot SFTP :

1. **Propriété** : chaque composant du chemin (`/srv`, `/srv/sftp`, `/srv/sftp/alice`)
   doit appartenir à **root:root** et ne pas être accessible en écriture au
   groupe/autres. Sinon : `fatal: bad ownership or modes for chroot directory`.
   Le répertoire **inscriptible** par l'utilisateur est un **sous-répertoire**
   (ex. `/srv/sftp/alice/upload` en `alice:sftp-only`).
2. `Subsystem sftp internal-sftp` doit exister dans `sshd_config` (le binaire
   externe ne marche pas en chroot sans les libs).
3. L'utilisateur ne doit pas avoir de shell valide nécessaire : avec
   `ForceCommand internal-sftp`, un shell `/bin/false` ou `/usr/sbin/nologin`
   est OK et même recommandé.
4. Tester : `sftp alice@serveur` puis `journalctl -u ssh -f` en parallèle —
   le message `fatal:` donne le composant fautif exact.

```bash
serveur# mkdir -p /srv/sftp/alice/upload
serveur# chown root:root /srv/sftp /srv/sftp/alice
serveur# chmod 755 /srv/sftp /srv/sftp/alice
serveur# chown alice:sftp-only /srv/sftp/alice/upload
```

## 74. Les erreurs classiques (tableau récapitulatif)

| Message / symptôme | Cause la plus probable | Section |
|---|---|---|
| `Permission denied (publickey)` | Clé non déployée / permissions / `from=` | 62 |
| `Host key verification failed` | Clé d'hôte changée (réinstall, IP réattribuée) | 63 |
| Connexion bloquée sur `Connecting to...` | Firewall / routage / sshd arrêté | 64 |
| Session qui freeze en inactivité | NAT idle → `ServerAliveInterval` | 64 |
| `Too many authentication failures` | Trop de clés dans l'agent, `IdentitiesOnly` absent | 67 |
| `UNPROTECTED PRIVATE KEY FILE` | Permissions > 600 sur la clé privée | 68 |
| `Connection refused` | sshd arrêté ou n'écoute pas ce port/IP | 64 |
| `Connection timed out` | Firewall qui drop, mauvaise IP, routage | 64 |
| `No route to host` | Routage / hôte éteint / mauvais réseau | 64 |
| `open failed: administratively prohibited` | `AllowTcpForwarding no` ou `PermitOpen` | 70 |
| `Bad owner or permissions on ~/.ssh/config` | Config lisible par d'autres → `chmod 600` | 68 |
| Prompt mot de passe malgré la clé | `PasswordAuthentication` encore actif + clé refusée (lire `-v`) | 62 |
| `Agent admitted failure to sign` | Agent sans la clé ou clé corrompue → `ssh-add -D` + `ssh-add` | 18 |
| `WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED` | Voir cas 2 | 63 |
| `ssh: Could not resolve hostname` | DNS / faute de frappe dans l'alias | 10 |

---

# CAS PRATIQUES COMMENTÉS

## 75. Cas pratique 1 : bastion d'équipe

**Contexte** : 15 serveurs en 10.10.0.0/16, un bastion public
`bastion.example.com`, 4 admins. Objectif : accès simple, sans forwarding
d'agent, avec traçabilité.

**`~/.ssh/config` de chaque admin** (déployé par Ansible, `config.d/20-prod.conf`) :

```
Host bastion
    HostName bastion.example.com
    User zelef
    Port 2222
    IdentityFile ~/.ssh/id_ed25519
    IdentitiesOnly yes

# Tout le LAN passe par le bastion, sans exception
Host 10.10.*.*
    User zelef
    IdentityFile ~/.ssh/id_ed25519
    IdentitiesOnly yes
    ProxyJump bastion
    ServerAliveInterval 60
    ServerAliveCountMax 3

# Alias mémorisable pour les serveurs fréquents
Host web1
    HostName 10.10.1.10
Host db1
    HostName 10.10.2.10
    User dba
```

Usage : `ssh web1`, `scp dump.sql db1:/tmp/`, `rsync -e ssh` — tout passe par
le bastion **sans y laisser de clés** (ProxyJump, pas de `-A`).

**Côté bastion** (`sshd_config.d/`) : `PasswordAuthentication no`,
`AllowGroups ssh-users`, `AllowTcpForwarding yes` (nécessaire au ProxyJump),
fail2ban actif, `auth.log` forwardé au SIEM. Le bastion **ne stocke aucune clé
privée d'admin** : il ne fait que relayer.

**Traçabilité** : `LogLevel VERBOSE` → chaque `Accepted` loggue l'empreinte de
la clé → on sait qui s'est connecté où et quand, sans agent de supervision
supplémentaire.

## 76. Cas pratique 2 : clé restreinte pour déploiement (`command=`)

**Contexte** : un runner CI doit déployer l'application sur `prod1`, sans
jamais obtenir un shell.

**Sur `prod1`**, utilisateur dédié `deployer` (shell `/bin/false`, groupe
`deploy`) :

```bash
serveur# useradd -m -s /bin/false -G deploy deployer
```

`/home/deployer/.ssh/authorized_keys` (une seule ligne) :

```
from="10.0.5.0/24",command="/usr/local/bin/deploy.sh",no-pty,no-agent-forwarding,no-X11-forwarding,no-port-forwarding ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIEXEMPLEdefictifNePasUtiliser ci@runner-01
```

