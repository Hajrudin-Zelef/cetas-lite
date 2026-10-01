---
id: collect-261001-rattrapage/rattrapage/ssh-guide-6
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [1079, 1305]
sha256: 6cc3dcdf8ff57f77e415512ac562b5b3c2c8b80f66e8c029bbb78fdf4ca5cf1d
---

# Guide SSH approfondi

```
# Autoriser ce -R précis, sur l'interface publique, port dédié
GatewayPorts clientspecified
PermitListen 18080
```

Côté client :

```bash
client$ ssh -R 0.0.0.0:18080:localhost:5000 web1
# Le partenaire accède à http://web1.example.com:18080
```

Garde-fous :

1. `PermitListen` borne les ports possibles (pas de squat du 80/443).
2. Protège l'applicatif exposé (authentification, pas de debug).
3. **Temporaire** : ce n'est pas une architecture d'hébergement. Pour du
   durable, reverse proxy (nginx/traefik) + TLS + authentification.
4. Journalise : `LogLevel VERBOSE` côté serveur trace les ouvertures de tunnels.

## 35. Tunnels : bonnes pratiques et pièges

Checklist avant d'ouvrir un tunnel en production :

- [ ] `ExitOnForwardFailure yes` : échec franc si le port est occupé.
- [ ] Bind sur `127.0.0.1` par défaut ; `0.0.0.0` seulement si besoin réel,
      avec `GatewayPorts` maîtrisé côté serveur.
- [ ] Pas de tunnel permanent sans supervision : un tunnel mort = un service
      mort silencieusement. Préfère `autossh` (reconnexion auto) :

```bash
client$ autossh -M 0 -N -f -o ServerAliveInterval=30 -o ServerAliveCountMax=3 \
    -L 8080:localhost:9090 web1
```

- [ ] Documente chaque tunnel persistant (qui, quoi, pourquoi, jusqu'à quand).
- [ ] Ne fais jamais transiter de données sensibles par un tunnel vers un
      serveur que tu n'administres pas.

Pièges classiques :

| Symptôme | Cause probable |
|---|---|
| `bind: Address already in use` | Port local occupé (autre tunnel, appli) |
| Tunnel « ouvert » mais rien ne répond | `ExitOnForwardFailure` absent + port occupé |
| Le tunnel coupe après quelques minutes | NAT/firewall idle → `ServerAliveInterval` |
| `-R` inaccessible depuis l'extérieur | `GatewayPorts no` côté serveur |

## 36. SCP : copies simples

SCP reste le plus rapide pour un fichier ou une arborescence simple :

```bash
client$ scp rapport.pdf web1:/tmp/                    # vers le distant
client$ scp web1:/var/log/syslog ./                   # depuis le distant
client$ scp -r ./site/ web1:/var/www/                 # récursif
client$ scp -P 2222 -i ~/.ssh/id_ed25519 f.bin web1:/tmp/   # port + clé
client$ scp -3 posteA:/f.bin posteB:/tmp/             # distant -> distant via ton poste
```

Note moderne : depuis OpenSSH 9.0, `scp` utilise le protocole SFTP en
sous-main (plus de `scp -O` sauf vieux serveurs). Les options restent les
mêmes à l'usage.

Limites : pas de reprise sur échec, pas de différentiel, permissions parfois
approximatives. Pour du sérieux, voir rsync (section 38).

## 37. SFTP : transferts interactifs et batch

SFTP = FTP sécurisé par SSH (à ne pas confondre avec FTPS). Session
interactive :

```bash
client$ sftp web1
sftp> lpwd / pwd          # répertoires local / distant
sftp> lls / ls
sftp> put rapport.pdf /tmp/
sftp> mput *.log
sftp> get /var/log/syslog ./
sftp> mkdir /tmp/depot && chmod 750 /tmp/depot
sftp> bye
```

Mode batch (scripts) :

```bash
client$ sftp -b batch.txt web1
# batch.txt :
#   put rapport.pdf /tmp/
#   chmod 640 /tmp/rapport.pdf
#   bye
client$ echo "put rapport.pdf /tmp/" | sftp -b - web1
```

Côté serveur, pour un accès **fichiers uniquement** (sans shell), le chroot
SFTP est la bonne architecture (voir cas 12, section 73) :

```
Match Group sftp-only
    ChrootDirectory /srv/sftp/%u
    ForceCommand internal-sftp
    AllowTcpForwarding no
    X11Forwarding no
```

## 38. rsync over SSH

L'outil roi des synchronisations : différentiel (ne transfère que les
changements), reprise, préservation des permissions.

```bash
# Syntaxe de base
client$ rsync -avz --progress ./site/ web1:/var/www/site/
# -a archive (permissions, dates, liens), -v verbeux, -z compression

# Options SSH personnalisées via -e
client$ rsync -avz -e "ssh -p 2222 -i ~/.ssh/id_ed25519" ./site/ web1:/var/www/

# À travers un bastion
client$ rsync -avz -e "ssh -J bastion" ./site/ web1-interne:/var/www/

# Supprimer côté distant ce qui n'existe plus en local (DANGEREUX sans --dry-run)
client$ rsync -avz --delete --dry-run ./site/ web1:/var/www/site/
```

Le `/` final sur la source change tout : `./site/` synchronise le **contenu**,
`./site` crée le **répertoire** lui-même. En cas de doute : `--dry-run`
d'abord.

Pour les gros transferts : `--partial --progress` (reprise), `--bwlimit=10000`
(Ko/s, pour ne pas saturer un lien de prod en journée).

## 39. Optimiser les transferts (compression, algorithmes)

- `-z` (rsync/scp) ou `Compression yes` : utile sur liens lents ou données
  compressibles (logs, textes). **Inutile voire contre-productif** en LAN
  10 Gb/s ou sur données déjà compressées (zip, jpg, vidéos) : ça consomme du
  CPU pour rien.
- Algorithme de chiffrement rapide pour les gros volumes :
  `chacha20-poly1305` est souvent le plus rapide sans AES-NI ; avec AES-NI,
  `aes128-gcm` gagne. Test :

```bash
client$ scp -c chacha20-poly1305 gros-fichier.bin web1:/tmp/   # si dispo
# ou en config :
Host transfert-rapide
    HostName web1
    Ciphers chacha20-poly1305@openssh.com
```

- `NoneSwitch`/`none` cipher : existe en patch uniquement, **à proscrire**.
- Paralléliser : plusieurs `rsync` sur des sous-répertoires > un seul flux,
  sur les liens à forte latence (chaque flux TCP a sa fenêtre).

## 40. X11 forwarding : usage et prudence

Permet d'afficher une application graphique distante sur ton poste :

```bash
client$ ssh -X web1          # forwarding X11 "non fiable" (restrictions)
client$ ssh -Y web1          # forwarding "fiable" (pleins droits) — À ÉVITER
```

```
Host appli-graphique
    HostName web1
    ForwardX11 yes
    ForwardX11Trusted no     # équivalent -X, jamais -Y par défaut
```

**Prudence** (c'est un vecteur d'attaque réel) :

- `-Y`/`ForwardX11Trusted yes` donne au serveur distant un accès quasi total à
  ton serveur X local (capture d'écran, injection de touches). Ne l'utilise
  que vers des machines de confiance absolue, et jamais en `Host *`.
- Le X11 forwarding transmet aussi des données non chiffrées côté serveur X
  local si celui-ci écoute en TCP : sur un poste moderne c'est du socket Unix,
  OK.
- Alternative moderne : pour une appli graphique, préfère un accès via
  navigateur + tunnel `-L` (section 33), ou RDP/VNC chiffré.

Côté serveur : `X11Forwarding no` par défaut dans le durcissement (section 46),
`yes` uniquement sur les serveurs qui en ont besoin.

## 41. `sshd_config` : structure et rechargement

Fichier : `/etc/ssh/sshd_config` (+ `/etc/ssh/sshd_config.d/*.conf` sur les
versions récentes — **préfère ce répertoire** pour tes réglages, le fichier
principal restant celui du paquet).

Règles inverses du client : **la première valeur obtenue gagne aussi**, mais
les blocs `Match` doivent être **en fin de fichier** (tout ce qui suit un
`Match` en fait partie jusqu'au prochain `Match`).

```bash
serveur# sshd -T    # affiche la config EFFECTIVE (après fusion) — outil n°1
serveur# sshd -t    # teste la syntaxe avant de recharger (obligatoire !)
serveur# systemctl reload ssh     # recharge sans couper les sessions existantes
```

**Procédure de modification sans risque** (à afficher dans ton équipe) :

1. `cp /etc/ssh/sshd_config.d/hardening.conf{,.bak-$(date +%F)}`
2. Éditer.
3. `sshd -t` → doit être silencieux (silence = OK).
4. `systemctl reload ssh`.
5. **Tester depuis une 2e session déjà ouverte** avant de fermer la 1re.
   (Garde toujours une session de secours ouverte quand tu touches au SSH.)

## 42. Durcissement : authentification

Le cœur du durcissement. Fichier `/etc/ssh/sshd_config.d/10-auth.conf` :

```
# --- Authentification : clés uniquement ---
PermitRootLogin no
# Variante : autoriser root UNIQUEMENT par clé (transition, à éviter à terme)
# PermitRootLogin prohibit-password
PasswordAuthentication no
ChallengeResponseAuthentication no
UsePAM yes
PubkeyAuthentication yes
```

Détails :

