---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-15
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "datacenter"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [3439, 3714]
sha256: b64a7532bea687960a72789af33b930a1bfaeb2a0098eeb9705a3477d53c6d4f
---

# Guide fail2ban — Le bouclier anti-brute-force

- [ ] `sudo grep " Ban " /var/log/fail2ban.log | wc -l` : volume de la semaine vs semaine précédente (pic ?)
- [ ] Top 10 des IP bannies : des récidivistes ? → envisager ban permanent manuel ou remontée AbuseIPDB
- [ ] `sudo fail2ban-client status --all` : tous les jails actifs ? (`Number of jail` == attendu)
- [ ] Y a-t-il eu des tickets « je suis bloqué » ? → ajuster seuils ou `ignoreip`
- [ ] Espace disque `/var/log/` : `fail2ban.log` ne grossit pas anormalement (loglevel DEBUG oublié ?)
- [ ] Le jail `recidive` a-t-il banni ? (si oui : attaques persistantes en cours)

---

## 63. Checklist de revue mensuelle

- [ ] Relire `jail.local` (ou `git log`) : encore pertinent ? nouveaux services à protéger ?
- [ ] Vérifier les mises à jour fail2ban en attente (`apt list --upgradable | grep fail2ban`)
- [ ] Contrôler `ignoreip` : l'équipe a-t-elle changé ? nouvelles sondes ? IP du bureau modifiée ?
- [ ] Tester un débannissement et un bannissement manuel (la procédure d'urgence est-elle connue de l'astreinte ?)
- [ ] Revoir les seuils avec les stats du mois (section 23 : méthode de tuning)
- [ ] Sauvegarde de `/etc/fail2ban/` à jour ? (section 37)
- [ ] Corréler avec les autres sources : les IP bannies apparaissent-elles aussi dans les alertes WAF / IDS ?

---

## 64. Cas pratique : protéger un serveur web exposé

### Contexte

Serveur `srv-web` : Debian 12, nginx + PHP-FPM, site vitrine + espace client avec basic auth sur `/admin`. IP publique directe (pas de CDN).

### Étape 1 : inventaire des surfaces

```bash
ss -tlnp | grep -E "nginx|sshd"
# 0.0.0.0:80, 0.0.0.0:443 (nginx), 0.0.0.0:22 (sshd)
ls -la /var/log/nginx/   # access.log + error.log présents ?
```

### Étape 2 : jail.local

```ini
[DEFAULT]
bantime  = 1h
findtime = 10m
maxretry = 5
ignoreip = 127.0.0.1/8 ::1 192.168.10.0/24 203.0.113.10 198.51.100.5
banaction = nftables-multiport
banaction_allports = nftables-allports

[sshd]
enabled = true

[nginx-http-auth]
enabled = true
port    = http,https
filter  = nginx-http-auth
logpath = /var/log/nginx/error.log

[nginx-limit-req]
enabled = true
port    = http,https
filter  = nginx-limit-req
logpath = /var/log/nginx/error.log
maxretry = 10
findtime = 1m
bantime  = 10m
# ^ bannit les IP qui se prennent des 503 par le rate-limit nginx :
#   complément idéal au limit_req de nginx.

[nginx-botsearch]
enabled  = true
filter   = nginx-botsearch
logpath  = /var/log/nginx/access.log
maxretry = 3
findtime = 1d
bantime  = 1w

[nginx-badbot]
enabled  = true
filter   = nginx-badbot
logpath  = /var/log/nginx/access.log
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

### Étape 3 : durcir nginx en amont (défense en profondeur)

```nginx
# Rate limiting : 10 req/s par IP, burst 20
limit_req_zone $binary_remote_addr zone=login:10m rate=10r/s;

server {
  location /admin/ {
    auth_basic "Espace client";
    auth_basic_user_file /etc/nginx/.htpasswd;
    limit_req zone=login burst=20 nodelay;
  }
  # Bloquer les chemins de scan les plus courants (avant fail2ban)
  location ~* (\.env|\.git|wp-login\.php|phpmyadmin) {
    return 444;   # ferme la connexion sans répondre
  }
}
```

### Étape 4 : validation

```bash
sudo fail2ban-client --test && sudo fail2ban-client reload
# Depuis une VM de test : forcer 6 mauvais logins sur /admin → vérifier le ban
curl -u faux:mauvais https://srv-web/admin/  # x6
sudo grep "Ban " /var/log/fail2ban.log | tail -3
```

---

## 65. Cas pratique : durcir un Proxmox VE

### Contexte

Nœud Proxmox VE 8.x : interface web sur le port 8006, SSH sur 22, exposé sur un réseau d'administration (jamais sur Internet directement — mais fail2ban reste utile contre les erreurs internes et les mouvements latéraux).

### Jail pour l'interface web Proxmox (port 8006)

Proxmox loggue les échecs d'authentification web dans `/var/log/daemon.log` (pveproxy) :

```
Sep 26 12:00:01 pve pvedaemon[1234]: authentication failure; rhost=203.0.113.60 user=root@pam msg=no such user
```

Filter maison `/etc/fail2ban/filter.d/proxmox.conf` :

```ini
[INCLUDES]
before = common.conf

[Definition]
failregex = ^%(__prefix_line)sauthentication failure; rhost=<HOST> user=\S+ msg=.*$
ignoreregex =
```

Test :

```bash
sudo fail2ban-regex /var/log/daemon.log /etc/fail2ban/filter.d/proxmox.conf --print-all-matched
```

Jail `/etc/fail2ban/jail.d/proxmox.local` :

```ini
[proxmox]
enabled  = true
port     = 8006
filter   = proxmox
logpath  = /var/log/daemon.log
maxretry = 5
findtime = 10m
bantime  = 1h
```

### Jail SSH (standard)

```ini
[sshd]
enabled = true
port    = 22
maxretry = 5
```

### Points d'attention Proxmox

1. **Cluster** : chaque nœud a son fail2ban. Une IP bannie sur le nœud 1 ne l'est pas sur le nœud 2. Pour un blocage cluster-wide, il faut une action qui propage (ssh vers les autres nœuds, ou firewall Proxmox au niveau datacenter — voir ci-dessous).
2. **Firewall intégré Proxmox** : Proxmox a son propre firewall (datacenter > Firewall). Vous pouvez AUSSI y mettre des règles, mais ne mélangez pas : fail2ban gère le dynamique, le firewall Proxmox le statique.
3. **Ne bannissez jamais l'IP d'un autre nœud du cluster** : mettez toutes les IP du cluster en `ignoreip` ! Un nœud qui se fait bannir par un autre = split-brain et migration HA en échec.
4. **Port 8006** : l'action doit bloquer le port 8006 (`port = 8006`), sinon le ban ne sert à rien.

```ini
[DEFAULT]
# Sur un nœud Proxmox en cluster à 3 nœuds (exemple) :
ignoreip = 127.0.0.1/8 ::1 10.30.0.11 10.30.0.12 10.30.0.13 192.168.10.0/24
```

### Vérification

```bash
sudo fail2ban-client status proxmox
# Provoquez 6 échecs de login sur https://pve:8006 depuis une IP de test,
# puis vérifiez le ban :
sudo nft list table inet f2b-table | grep -A5 "f2b-proxmox"
```

---

## 66. Cas pratique : sécuriser un bastion SSH

### Contexte

Un bastion est LE point d'entrée : il mérite le niveau maximal.

### Configuration

```ini
[sshd]
enabled  = true
port     = 22
filter   = sshd[mode=aggressive]
logpath  = /var/log/auth.log
maxretry = 3
findtime = 10m
bantime  = 4h

[recidive]
enabled   = true
filter    = recidive
logpath   = /var/log/fail2ban.log
banaction = %(banaction_allports)s
bantime   = 4w
findtime  = 1d
maxretry  = 3
```

### Mesures complémentaires (le bastion ne vit pas de fail2ban seul)

```bash
# /etc/ssh/sshd_config — durcissement bastion
PermitRootLogin no
PasswordAuthentication no        # clés uniquement
ChallengeResponseAuthentication no
AllowUsers admin-bastion          # liste fermée
MaxAuthTries 3
LoginGraceTime 30
```

- `PasswordAuthentication no` : sans mot de passe, le brute-force devient inutile — fail2ban ne sert plus qu'à calmer les scanners.
- 2FA via `libpam-google-authenticator` pour les accès sensibles.

### Particularité : les utilisateurs du bastion sont des admins

`ignoreip` doit contenir **tous** les réseaux d'où les admins se connectent (bureau, VPN, astreinte). Un admin banni du bastion un dimanche à 3h du matin, c'est une indisponibilité de l'infogérance.

---

## 67. Cas pratique : protéger un webmail Roundcube

### Contexte

Roundcube sur `webmail.exemple.com` : des centaines d'utilisateurs humains, smartphones en arrière-plan. **Faux positifs garantis** si on tune comme SSH.

### Logs Roundcube

Activez le log d'échecs (`config/config.inc.php`) :

```php
$config['log_driver'] = 'file';
$config['log_session_id'] = false;
```

Les échecs ressemblent à (selon version) :

```
[26-Sep-2026 13:00:01 +0200]: <abc123> IMAP Error: Login failed for victim@example.com
  from 203.0.113.70. AUTHENTICATE PLAIN: Authentication failed.
```

### Filter maison

```ini
# /etc/fail2ban/filter.d/roundcube.conf
[INCLUDES]
before = common.conf

