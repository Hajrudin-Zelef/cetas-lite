---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-14
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "incident"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [3180, 3438]
sha256: 4ea8c4536a68fe409572995ae08d1dceeb4ef85bf0f5bcf3ba1301fcb2cb2ccf
---

# Guide fail2ban — Le bouclier anti-brute-force

```bash
# Exemple nftables avec un set d'IP par pays (listes : ipdeny.com)
sudo nft add table inet geo-filter
sudo nft add set inet geo-filter blocked-countries '{ type ipv4_addr; flags interval; }'
# Charger la liste (script à écrire, mise à jour hebdo via cron)
for ip in $(cat /root/geo/country-zone.txt); do
  sudo nft add element inet geo-filter blocked-countries "{ $ip }"
done
sudo nft add rule inet geo-filter input ip saddr @blocked-countries drop
```

> C'est du filtrage statique au firewall. fail2ban garde son rôle : la détection **comportementale** (échecs répétés), quelle que soit l'origine.

---

## 57. Corrélation avec auditd

auditd trace les appels système ; fail2ban lit des logs applicatifs. Les deux se complètent pour l'enquête post-incident.

### Scénario : une IP bannie par fail2ban, que faisait-elle ?

```bash
IP="203.0.113.45"
# 1. Quand a-t-elle été bannie ?
sudo grep "Ban $IP" /var/log/fail2ban.log
# 2. Quelles lignes de log ont déclenché ?
sudo grep "$IP" /var/log/auth.log | tail -20
# 3. Y a-t-il eu des connexions RÉUSSIES depuis cette IP ? (le vrai danger)
sudo grep "$IP" /var/log/auth.log | grep -i "accepted"
# Si oui → compromission possible → incident de sécurité, pas juste un ban.
```

### Règle auditd utile : surveiller les modifications de fail2ban

```
# /etc/audit/rules.d/fail2ban.rules
-w /etc/fail2ban/jail.local -p wa -k fail2ban-config
-w /etc/fail2ban/jail.d/ -p wa -k fail2ban-config
-w /usr/bin/fail2ban-client -p x -k fail2ban-client
```

```bash
sudo augenrules --load
sudo ausearch -k fail2ban-config --start recent
```

Un attaquant qui a pris pied tentera de **désactiver fail2ban** pour travailler tranquille. Cette règle vous en avertit.

---

## 58. Fail2ban et Docker : attention

### Le problème

Docker manipule iptables/nftables avec ses propres chaînes (`DOCKER`, `DOCKER-USER`). Par défaut, les règles fail2ban sur `INPUT` **ne voient pas** le trafic destiné aux conteneurs (qui passe par `FORWARD`).

```
Internet → INPUT (fail2ban bloque ici : OK pour sshd de l'hôte)
Internet → DOCKER-USER → FORWARD → conteneur nginx
         (fail2ban sur INPUT ne voit RIEN ici)
```

### Solution : bannir dans DOCKER-USER

Utilisez une action qui insère dans la chaîne `DOCKER-USER` (évaluée avant les règles Docker) :

```ini
[nginx-botsearch]
enabled = true
banaction = iptables-multiport[chain=DOCKER-USER]
```

Ou en nftables : créez une action qui accroche la chaîne `forward` :

```ini
# Variante : banaction nftables sur le forward
# (à adapter selon votre version de l'action nftables)
```

### Alternative simple et robuste

Faites tourner fail2ban **dans un conteneur à part** avec accès au socket Docker et aux logs des autres conteneurs (montage des volumes de logs), ou — plus simple — exposez les ports via le host et laissez fail2ban sur l'hôte avec la chaîne adaptée.

### Vérification indispensable

```bash
# Depuis l'extérieur, avec une IP de test bannie :
# 1. Le conteneur est-il injoignable ?
curl -m 5 http://VOTRE_SERVEUR/ ; echo $?
# 2. Où est la règle ?
sudo iptables -L DOCKER-USER -n | grep 203.0.113.45
```

> Si vous ne testez pas, vous avez un fail2ban décoratif sur une infra Docker.

---

## 59. Journalctl et backend systemd : recettes

### jail postfix en backend systemd complet

```ini
[postfix]
enabled      = true
backend      = systemd
journalmatch = _SYSTEMD_UNIT=postfix.service
filter       = postfix
port         = smtp,465,submission
maxretry     = 5
```

### Plusieurs unités pour un jail

```ini
[dovecot]
enabled      = true
backend      = systemd
journalmatch = _SYSTEMD_UNIT=dovecot.service
filter       = dovecot
```

### Filtrer par niveau dans le journal

```ini
# Ne regarder que les erreurs et plus grave
journalmatch = _SYSTEMD_UNIT=nginx.service + PRIORITY=3
```

### Tester le journalmatch

```bash
sudo journalctl _SYSTEMD_UNIT=postfix.service --since "1 hour ago" | \
  sudo fail2ban-regex - /etc/fail2ban/filter.d/postfix.conf | tail -2
```

`fail2ban-regex` accepte `-` pour lire depuis stdin : parfait pour tester un flux journald.

---

## 60. Recette de jail.local « gold »

La configuration que je déploierais sur un serveur standard exposé (web + ssh). À adapter (surtout `ignoreip`).

```ini
# /etc/fail2ban/jail.local — recette "gold" (Debian 12/Ubuntu 24.04)
# ADAPTEZ ignoreip AVANT TOUT RELOAD.

[DEFAULT]
bantime  = 1h
findtime = 10m
maxretry = 5

# Escalade automatique du bantime (1h, 2h, 4h... max 4 semaines)
bantime.increment = true
bantime.factor    = 2
bantime.maxtime   = 4w
bantime.rndtime   = 10m

# >>> À PERSONNALISER : vos réseaux, votre supervision, votre bastion <<<
ignoreip = 127.0.0.1/8 ::1
           192.168.10.0/24
           10.20.0.0/16
           203.0.113.10
           198.51.100.5

banaction         = nftables-multiport
banaction_allports = nftables-allports

# ---------------------------------------------------------------- SSH
[sshd]
enabled = true
port    = ssh
filter  = sshd
logpath = /var/log/auth.log
backend = auto

# ---------------------------------------------------------------- Web : nginx
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
enabled  = true
port     = http,https
filter   = nginx-badbot
logpath  = /var/log/nginx/access.log
maxretry = 2
bantime  = 1w

# ---------------------------------------------------------------- Récidivistes
[recidive]
enabled   = true
filter    = recidive
logpath   = /var/log/fail2ban.log
banaction = %(banaction_allports)s
bantime   = 1w
findtime  = 1d
maxretry  = 5
# Alerte mail uniquement pour les récidivistes (bruit minimal)
action    = %(banaction_allports)s
            %(action_mwl)s[name=recidive, dest=admin@example.com, sender=fail2ban@%(hostname)s]
```

Mise en service :

```bash
sudo cp /etc/fail2ban/jail.local /root/jail.local.bak-$(date +%Y%m%d)
# ... éditez ...
sudo fail2ban-client --test && sudo fail2ban-client reload
sudo fail2ban-client status
```

---

## 61. Checklist de mise en production

À cocher avant d'activer fail2ban sur un serveur de production.

### Avant le reload

- [ ] `ignoreip` contient : loopback, LAN admin, IP publique du bureau, bastion, sondes de supervision
- [ ] `fail2ban-client --test` passe sans erreur
- [ ] Chaque filter testé avec `fail2ban-regex` (`--print-all-matched` relu)
- [ ] Chaque `logpath` existe et contient des lignes récentes
- [ ] Le `port` de chaque jail correspond au port réel du service (`ss -tlnp`)
- [ ] `banaction` cohérente avec le firewall actif (nftables/iptables/ufw)
- [ ] Console hors-bande ouverte (ou dead man's switch programmé, section 25)

### Juste après le reload

- [ ] `fail2ban-client status` : tous les jails attendus sont listés
- [ ] `fail2ban-client status <jail>` : `File list` pointe vers le bon fichier
- [ ] `/var/log/fail2ban.log` : pas d'ERROR dans les 50 dernières lignes
- [ ] Les règles firewall existent : `nft list table inet f2b-table` ou `iptables -L f2b-sshd -n`
- [ ] Test de ban contrôlé depuis une IP de test (pas la vôtre) : 6 échecs → ban visible

### 24h après

- [ ] Relire `/var/log/fail2ban.log` : des bans ? des erreurs ?
- [ ] Vérifier l'absence de faux positifs (IP internes bannies ?)
- [ ] Le rapport quotidien (section 48) arrive bien par mail

---

## 62. Checklist de revue hebdomadaire

(15 minutes, le lundi avec le café.)

