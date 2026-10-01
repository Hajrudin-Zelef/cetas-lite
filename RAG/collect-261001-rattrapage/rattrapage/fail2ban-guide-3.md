---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-3
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "apache", "attention"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [424, 675]
sha256: 6b54de9034d3432f7d553ad770dba10fc4bb6a62f232fcef93980a31e07ece3a
---

# Guide fail2ban — Le bouclier anti-brute-force

```
Status for the jail: sshd
|- Filter
|  |- Currently failed:	0
|  |- Total failed:	0
|  `- File list:	/var/log/auth.log
`- Actions
   |- Currently banned:	0
   |- Total banned:	0
   `- Banned IP list:
```

> **Point de vigilance immédiat** : `ignoreip` ci-dessus contient des exemples. Remplacez `192.168.1.0/24` et `10.0.0.0/8` par **vos** réseaux réels (section 24). Un `ignoreip` mal rempli = risque de vous bannir vous-même.

### Les unités de temps acceptées

| Suffixe | Signification | Exemple |
|---|---|---|
| (aucun) | secondes | `bantime = 3600` |
| `s` | secondes | `60s` |
| `m` | minutes | `10m` |
| `h` | heures | `1h` |
| `d` | jours | `7d` |
| `w` | semaines | `2w` |
| `y` | années | `1y` (= 365j) |

On peut combiner : `bantime = 1d12h` (= 36 heures). Et `bantime = -1` signifie **permanent** (section 28).

---

## 6. bantime, findtime, maxretry expliqués

Ces trois paramètres forment le **triangle de décision**. Bien les comprendre évite 90 % des erreurs de tuning.

### Définitions précises

- **`maxretry`** : nombre d'échecs (lignes matchées par le filter) qui déclenche le ban.
- **`findtime`** : fenêtre glissante pendant laquelle on compte les échecs. Un échec vieux de plus de `findtime` sort du compteur.
- **`bantime`** : durée du bannissement une fois déclenché.

### Exemple déroulé

Config : `maxretry = 5`, `findtime = 10m`, `bantime = 1h`. L'IP 203.0.113.45 échoue à :

```
08:00, 08:01, 08:02, 08:03   → 4 échecs, rien ne se passe
08:20                         → l'échec de 08:00 a plus de 10 min : il sort du compteur
                               compteur = 4 (08:01, 08:02, 08:03, 08:20)
08:21                         → compteur = 5 → BAN jusqu'à 09:21
```

### Les erreurs de raisonnement classiques

| Idée reçue | Réalité |
|---|---|
| « `maxretry = 3` c'est plus sûr » | Sur un service avec des utilisateurs humains (webmail, VPN), 3 fautes de frappe arrivent vite → faux positifs. Réservez les seuils bas aux services sans humains (SSH root). |
| « `findtime` long = plus strict » | Oui, mais aussi plus de faux positifs : un utilisateur qui se trompe 5 fois dans la journée se fait bannir. |
| « `bantime` long = mieux » | Pour les botnets à IP tournantes, un bantime de 24h ne change rien (l'IP ne reviendra jamais) et remplit la table firewall. Pour les attaquants persistants, c'est utile. |

### Valeurs de départ recommandées (à affiner section 23)

| Contexte | maxretry | findtime | bantime |
|---|---|---|---|
| SSH exposé (pas d'humains légitimes sauf vous) | 5 | 10m | 1h |
| Service web avec utilisateurs | 10 | 10m | 30m |
| Mail (postfix/dovecot) | 5 | 10m | 1h |
| Récidivistes (jail recidive) | 3 | 1d | 1w |

---

## 7. Le jail sshd en détail

C'est le jail le plus important : **tout serveur Linux exposé est scanné en SSH en permanence**.

### Configuration recommandée

```ini
[sshd]
enabled  = true
port     = ssh
filter   = sshd
logpath  = /var/log/auth.log
backend  = auto
maxretry = 5
findtime = 10m
bantime  = 1h
```

Sur les systèmes avec systemd et `sshd` qui loggue dans le journal :

```ini
[sshd]
enabled  = true
backend  = systemd
# logpath inutile avec le backend systemd
```

> `backend = auto` choisit tout seul (pyinotify → gamin → polling → systemd). En pratique, `auto` convient dans 95 % des cas. Ne forcez `systemd` que si `/var/log/auth.log` est vide alors que les tentatives existent (section 45).

### Ce que détecte le filter sshd

Le filter `sshd.conf` couvre, entre autres :

- `Failed password for ... from <HOST>` — mot de passe invalide
- `Failed publickey for ... from <HOST>` — clé invalide
- `Invalid user ... from <HOST>` — utilisateur inexistant
- `Connection closed by authenticating user ... <HOST> [preauth]` — abandon pendant l'authentification
- `Disconnected from authenticating user ... <HOST> [preauth]`
- `Received disconnect from <HOST> ... [preauth]` — certains scanners
- `error: PAM: Authentication failure for ... from <HOST>`
- `maximum authentication attempts exceeded for ... from <HOST>`

### Le mode agressif

`filter.d/sshd.conf` propose un mode `aggressive` qui bannit aussi les scanners qui ne tentent même pas de s'authentifier (probes de version, connexions fermées) :

```ini
[sshd]
enabled = true
filter  = sshd[mode=aggressive]
```

> À réserver aux serveurs très exposés : le mode agressif augmente les faux positifs (un simple `nmap` ou un monitoring Nagios qui ouvre le port 22 peut se faire bannir).

### Vérifier que le jail voit bien les attaques

```bash
# Générer des échecs volontaires depuis une machine de test (PAS depuis votre poste prod !)
for i in 1 2 3 4 5 6; do ssh -o ConnectTimeout=3 fauxuser@SERVEUR; done

# Sur le serveur, observer :
sudo tail -f /var/log/fail2ban.log
# ... NOTICE [sshd] Ban 198.51.100.23
```

---

## 8. Le jail sshd : variantes et cas particuliers

### SSH sur un port non standard

Si sshd écoute sur le port 2222 :

```ini
[sshd]
enabled = true
port    = 2222
filter  = sshd
logpath = /var/log/auth.log
```

> Le paramètre `port` sert à l'action firewall : il détermine quel(s) port(s) sont bloqués. Si vous vous trompez de port, le ban ne bloquera rien d'utile.

### Plusieurs ports SSH

```ini
[sshd]
enabled = true
port    = ssh,2222
```

### SSH derrière un bastion : ne bannir que l'extérieur

Si votre serveur n'est joignable en SSH que depuis le bastion (bonne pratique), le jail sshd devient presque inutile — mais gardez-le : il protège contre un bastion compromis ou une erreur de firewall.

### Journal auth.log vs journald

Sur Debian 12+ / Ubuntu 22.04+, rsyslog écrit toujours `/var/log/auth.log` par défaut. Vérifiez :

```bash
ls -la /var/log/auth.log
# Si le fichier n'existe pas ou reste vide :
sudo journalctl -u ssh --since "10 min ago" | head
# → alors utilisez backend = systemd dans le jail
```

### Dropbear (routeurs, embarqué)

```ini
[dropbear]
enabled  = true
port     = ssh
filter   = dropbear
logpath  = /var/log/auth.log
maxretry = 5
```

---

## 9. nginx-http-auth : protéger le basic auth

Le basic auth HTTP (pop-up login/mot de passe) est une cible classique : interfaces d'admin, Kibana, zones `/admin`.

```ini
[nginx-http-auth]
enabled  = true
port     = http,https
filter   = nginx-http-auth
logpath  = /var/log/nginx/error.log
maxretry = 5
findtime = 10m
bantime  = 1h
```

Ce que le filter détecte dans `/var/log/nginx/error.log` :

```
2026/09/26 10:22:11 [error] 1234#1234: *5678 user "admin" was not found in "/etc/nginx/.htpasswd",
  client: 203.0.113.78, server: intranet.example.com,
  request: "GET /admin/ HTTP/1.1", host: "intranet.example.com"
```

> Point d'attention : c'est bien `error.log` (pas `access.log`) que nginx utilise pour les échecs d'authentification basique. Si votre `error.log` est verbeux, le filter ne matche que les lignes d'authentification, pas les autres erreurs.

### Variante Apache

```ini
[apache-auth]
enabled  = true
port     = http,https
filter   = apache-auth
logpath  = /var/log/apache2/error.log
maxretry = 5
```

---

## 10. nginx-badbot et nginx-botsearch

Deux jails complémentaires pour l'hygiène d'un serveur web public.

### nginx-botsearch : les scans de vulnérabilités

Détecte les requêtes vers des chemins qui n'existent pas et qui ressemblent à des scans (`/wp-admin`, `/phpmyadmin`, `/.env`, `/actuator`, etc.) :

```ini
[nginx-botsearch]
enabled  = true
port     = http,https
filter   = nginx-botsearch
logpath  = /var/log/nginx/access.log
maxretry = 3
findtime = 1d
bantime  = 1w
```

> Seuils volontairement sévères : aucun utilisateur légitime ne demande `/.git/config` ou `/wp-login.php` sur un site qui n'est pas WordPress. `maxretry = 3` sur 1 jour, c'est raisonnable.

### nginx-badbot : les user-agents pourris

Détecte les bots connus pour être malveillants via leur User-Agent :

