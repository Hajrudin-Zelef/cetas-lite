---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-4
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "distribution"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [676, 958]
sha256: f92713019cbcdbcd0bcfa68a4c6a6460e5fad2b49907dbc5c2d8d0398e3e61f0
---

# Guide fail2ban — Le bouclier anti-brute-force

```ini
[nginx-badbot]
enabled  = true
port     = http,https
filter   = nginx-badbot
logpath  = /var/log/nginx/access.log
maxretry = 2
bantime  = 1w
```

### Adapter le format de log nginx

Les filters nginx supposent le `combined` log format par défaut. Si vous avez personnalisé `log_format`, vérifiez que les champs `$remote_addr`, `$request` et `$http_user_agent` sont présents, sinon les filters ne matcheront pas (erreur classique n°2, section 39).

```nginx
# /etc/nginx/nginx.conf — format compatible fail2ban
log_format main '$remote_addr - $remote_user [$time_local] "$request" '
                '$status $body_bytes_sent "$http_referer" '
                '"$http_user_agent"';
access_log /var/log/nginx/access.log main;
```

---

## 11. postfix : relais et SASL

Un serveur mail exposé subit des tentatives d'authentification SASL en continu. Deux jails complémentaires.

```ini
[postfix]
enabled  = true
port     = smtp,465,submission
filter   = postfix
logpath  = /var/log/mail.log
maxretry = 5
findtime = 10m
bantime  = 1h
```

```ini
[postfix-sasl]
enabled  = true
port     = smtp,465,submission,imap,imaps,pop3,pop3s
filter   = postfix[mode=auth]
logpath  = /var/log/mail.log
maxretry = 5
findtime = 10m
bantime  = 1h
```

Ce que ça détecte :

```
# postfix (générique) :
Sep 26 11:00:01 srv-mail postfix/smtpd[3210]: NOQUEUE: reject: RCPT from unknown[203.0.113.90]:
  554 5.7.1 <victim@example.com>: Relay access denied; ...

# postfix[mode=auth] (SASL) :
Sep 26 11:00:02 srv-mail postfix/smtpd[3210]: warning: unknown[203.0.113.90]:
  SASL LOGIN authentication failed: authentication failure
```

> Sur les systèmes avec rsyslog, les logs postfix sont dans `/var/log/mail.log`. Avec systemd pur, utilisez `backend = systemd` et `journalmatch = _SYSTEMD_UNIT=postfix.service`.

### Cas particulier : postfix derrière un anti-spam

Si un relais (type Proxmox Mail Gateway) est devant, les logs voient l'IP du relais, pas celle de l'attaquant. **Ne pas activer** le jail postfix dans ce cas, ou configurez `postscreen`/`smtpd` pour logger la vraie IP (via `X-Forwarded-For` équivalent SMTP : en-têtes `Received`).

---

## 12. dovecot : POP3/IMAP sous pression

```ini
[dovecot]
enabled  = true
port     = pop3,pop3s,imap,imaps,submission,465,sieve
filter   = dovecot
logpath  = /var/log/mail.log
maxretry = 5
findtime = 10m
bantime  = 1h
```

Détecte :

```
Sep 26 11:05:00 srv-mail dovecot: imap-login: Disconnected: Connection closed:
  auth failed, 1 attempts in 2 secs (auth failed, 1 attempts):
  user=<victim@example.com>, method=PLAIN, rip=203.0.113.91, lip=192.0.2.10
```

### Le piège des clients mail mal configurés

C'est LE service où les faux positifs sont fréquents : un smartphone avec un ancien mot de passe tente de se reconnecter **toutes les 5 minutes**, indéfiniment. Recommandations :

- `maxretry = 10` plutôt que 5 pour dovecot si vous avez des utilisateurs humains.
- `bantime = 30m` plutôt que 1h : un utilisateur bloqué une heure appelle le support.
- Communiquez : « si votre téléphone redemande le mot de passe en boucle, appelez avant d'être bloqué ».

### dovecot agressif

```ini
[dovecot]
enabled = true
filter  = dovecot[mode=aggressive]
```

Le mode agressif matche aussi les échecs `pop3-login`/`managesieve-login` avec des patterns supplémentaires. À tester avec `fail2ban-regex` d'abord.

---

## 13. proftpd et vsftpd

Le FTP est un protocole legacy ; s'il est encore exposé, il est pilonné.

```ini
[proftpd]
enabled  = true
port     = ftp,ftp-data,ftps,ftps-data
filter   = proftpd
logpath  = /var/log/proftpd/proftpd.log
maxretry = 5
findtime = 10m
bantime  = 1h
```

```ini
[vsftpd]
enabled  = true
port     = ftp,ftp-data,ftps,ftps-data
filter   = vsftpd
logpath  = /var/log/vsftpd.log
maxretry = 5
```

> Vérifiez TOUJOURS le `logpath` réel : selon la distribution et la configuration, proftpd peut logger dans `/var/log/auth.log`, `/var/log/proftpd/proftpd.log` ou via syslog. Un jail qui pointe vers un fichier vide ne protège rien (erreur classique n°3).

### Recommandation stratégique

Si vous pouvez, **ne pas exposer FTP du tout** : migrez vers SFTP (couvert par le jail sshd). Un jail proftpd bien configuré reste mieux que rien.

---

## 14. recidive : les récidivistes

Le jail `recidive` est la **deuxième ligne de défense** : il surveille le log de fail2ban lui-même et bannit longuement les IP déjà bannies plusieurs fois.

```ini
[recidive]
enabled  = true
filter   = recidive
logpath  = /var/log/fail2ban.log
banaction = %(banaction_allports)s
bantime  = 1w
findtime = 1d
maxretry = 5
```

Fonctionnement :

1. Une IP se fait bannir 5 fois en 1 jour par n'importe quel jail (chaque ban écrit `NOTICE [sshd] Ban 203.0.113.45` dans `/var/log/fail2ban.log`).
2. Le filter `recidive` matche ces lignes `Ban`.
3. L'IP est bannie **sur tous les ports** (`banaction_allports`) pendant **1 semaine**.

### Pourquoi `banaction_allports` ?

Un récidiviste qui a attaqué SSH attaquera probablement le web et le mail ensuite. Le bannir sur tous les ports (`iptables -A INPUT -s <ip> -j DROP` sans restriction de port) est logique à ce stade.

### Réglages selon la paranoïa

| Profil | findtime | maxretry | bantime |
|---|---|---|---|
| Standard | 1d | 5 | 1w |
| Strict (serveur très exposé) | 1d | 3 | 4w |
| Prudent (peur des faux positifs) | 2d | 10 | 3d |

> Le jail recidive ne fonctionne que si `/var/log/fail2ban.log` existe et contient les lignes de ban. Si vous avez mis `logtarget = SYSTEMD-JOURNAL`, adaptez avec `backend = systemd` (section 30).

---

## 15. Activer plusieurs jails sans conflit

### Règles de coexistence

1. **Un log, plusieurs jails : OK.** `nginx-http-auth` (sur `error.log`) et `nginx-botsearch` (sur `access.log`) cohabitent sans problème.
2. **Deux jails sur le même log avec des filters qui se chevauchent :** possible double comptage, mais chaque jail a son propre compteur — pas de conflit technique, juste un tuning à surveiller.
3. **Actions différentes par jail : OK.** Un jail peut bannir sur `http,https` pendant qu'un autre bannit sur tous les ports.

### Exemple : stack web complète

```ini
[DEFAULT]
bantime  = 1h
findtime = 10m
maxretry = 5
ignoreip = 127.0.0.1/8 ::1 192.168.10.0/24
banaction = nftables-multiport

[sshd]
enabled = true

[nginx-http-auth]
enabled = true
port    = http,https
filter  = nginx-http-auth
logpath = /var/log/nginx/error.log

[nginx-botsearch]
enabled  = true
port     = http,https
filter   = nginx-botsearch
logpath  = /var/log/nginx/access.log
maxretry = 3
findtime = 1d
bantime  = 1w

[nginx-badbot]
enabled = true
port    = http,https
filter  = nginx-badbot
logpath = /var/log/nginx/access.log
maxretry = 2
bantime  = 1w

[recidive]
enabled   = true
filter    = recidive
logpath   = /var/log/fail2ban.log
banaction = %(banaction_allports)s
bantime   = 1w
findtime  = 1d
maxretry  = 5
```

### Vérifier l'ensemble

```bash
sudo fail2ban-client status
# Status
# |- Number of jail:	5
# `- Jail list:	sshd, nginx-http-auth, nginx-botsearch, nginx-badbot, recidive

for j in sshd nginx-http-auth nginx-botsearch nginx-badbot recidive; do
  echo "=== $j ==="
  sudo fail2ban-client status "$j" | grep -E "Currently|Total|File list"
done
```

---

## 16. fail2ban-regex : tester son filtre

`fail2ban-regex` est l'outil le plus sous-utilisé de la suite. **Ne jamais mettre en production un filter non testé.**

### Syntaxe de base

```bash
sudo fail2ban-regex /var/log/auth.log /etc/fail2ban/filter.d/sshd.conf
```

Sortie (extraits) :

```
Running tests
=============

Use   failregex filter file : sshd, basedir: /etc/fail2ban
Use         log file : /var/log/auth.log
Use         encoding : UTF-8

Results
=======

Failregex: 27 total
|-  #) [# of hits] regular expression
|   1) [15] ^\s*(?:\S+ )?(?:kernel: \[...)...
...
Ignoreregex: 0 total

Date template hits:
|- [# of hits] date format
|  [2450] {^LN-BEG}ExYear(?P<_sep>-|/)?Month(?P<_sep>-|/)Day ...

