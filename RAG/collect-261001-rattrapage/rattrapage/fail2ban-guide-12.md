---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-12
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["apache", "attention"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [2625, 2927]
sha256: 0004f9cdbd7f329edc9198804ec13786b4fc4e11d167c0008bded7b9c6ef32c6
---

# Guide fail2ban — Le bouclier anti-brute-force

`copytruncate` : le fichier garde le même inode, fail2ban ne perd jamais le fil. (Petit risque de perdre quelques lignes pendant la troncature — acceptable pour des logs de sécurité.)

---

## 46. Fail2ban derrière un proxy inverse

### Le problème (rappel erreur n°8)

```
Client (198.51.100.7) → nginx reverse proxy → backend Apache
                                              logs : "connexion depuis 10.0.0.5 (le proxy)"
```

fail2ban sur le backend voit l'IP du proxy. Le bannir = couper le proxy = couper tout le monde.

### Solution nginx (côté proxy ET backend)

Sur le **proxy** : transmettre la vraie IP.

```nginx
proxy_set_header X-Real-IP $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
```

Sur le **backend** : utiliser la vraie IP dans les logs.

```nginx
# backend nginx
set_real_ip_from 10.0.0.0/8;      # le proxy interne
real_ip_header X-Forwarded-For;
real_ip_recursive on;
log_format main '$remote_addr - ...';   # $remote_addr = vraie IP client maintenant
```

### Solution Apache (backend)

```apache
# a2enmod remoteip
RemoteIPHeader X-Forwarded-For
RemoteIPTrustedProxy 10.0.0.5
LogFormat "%a %l %u %t \"%r\" %>s %b" combined_realip
```

`%a` loggue alors la vraie IP client, et les filters apache standards fonctionnent.

### Où placer fail2ban : proxy ou backend ?

| Option | Avantages | Inconvénients |
|---|---|---|
| Sur le proxy (recommandé) | voit les vraies IP directement, bloque tôt | un seul point de config |
| Sur le backend | protège aussi en cas de bypass du proxy | config real_ip obligatoire |
| Les deux | défense en profondeur | double administration |

> Recommandation : fail2ban sur le **proxy** pour les jails web (nginx-*), et sur chaque serveur pour sshd (chaque machine est scannée directement).

### En attendant la correction : whitelistez le proxy

```ini
[DEFAULT]
ignoreip = 127.0.0.1/8 ::1 10.0.0.5   # IP du reverse proxy
```

Ça évite la catastrophe, mais ça **désactive** la protection (toutes les IP passent par le proxy). C'est un pansement, pas un traitement.

---

## 47. Tester en sécurité : la discipline du dry-run

Il n'y a pas de vrai mode « dry-run » dans fail2ban, mais voici l'équivalent opérationnel.

### Protocole de test d'un nouveau jail

```bash
# 1. Créer le jail avec une action INOFFENSIVE (log seul)
# /etc/fail2ban/jail.d/test.local
[nginx-test]
enabled  = true
filter   = nginx-botsearch
logpath  = /var/log/nginx/access.log
maxretry = 3
action   = %(action_)s
           # action_ avec banaction remplacée par un simple log :
banaction = dummy-log

# /etc/fail2ban/action.d/dummy-log.conf
[Definition]
actionban   = echo "[DRY-RUN] bannirait <ip> (jail <name>)" >> /var/log/fail2ban-dryrun.log
actionunban = echo "[DRY-RUN] débannirait <ip> (jail <name>)" >> /var/log/fail2ban-dryrun.log
```

```bash
# 2. Recharger, attendre 24-48h
sudo fail2ban-client reload

# 3. Analyser ce qui AURAIT été banni
cat /var/log/fail2ban-dryrun.log
# Des IP légitimes dedans ? → ajustez le filter/seuils avant de passer en vrai ban.
```

### Test de charge d'un filter

```bash
# Combien de lignes matchent sur 7 jours de logs ?
zcat /var/log/nginx/access.log* 2>/dev/null | \
  sudo fail2ban-regex - /etc/fail2ban/filter.d/nginx-botsearch.conf | tail -2
```

Si le filter matche 50 000 lignes/semaine dont 40 000 légitimes (ex: votre appli mobile qui appelle `/.well-known/`), revoyez la regex avant d'activer le ban.

---

## 48. Scripts d'administration utiles

### Rapport quotidien des bans (cron)

`/usr/local/bin/fail2ban-daily-report.sh` :

```bash
#!/bin/bash
# Rapport quotidien fail2ban → mail à l'admin
DEST="admin@example.com"
HOST=$(hostname -f)
TMP=$(mktemp)

{
  echo "Rapport fail2ban du $(date +%Y-%m-%d) — $HOST"
  echo "=================================================="
  echo
  fail2ban-client status
  echo
  for j in $(fail2ban-client status | grep "Jail list" | sed 's/.*Jail list://;s/,//g'); do
    echo "--- Jail: $j ---"
    fail2ban-client status "$j" | grep -E "Currently|Total"
    echo
  done
  echo "Top 10 des IP bannies (24h) :"
  grep "$(date -d 'yesterday' +%Y-%m-%d).* Ban " /var/log/fail2ban.log | \
    awk '{print $NF}' | sort | uniq -c | sort -rn | head -10
  echo
  echo "Top 10 des IP bannies (7j) :"
  grep " Ban " /var/log/fail2ban.log | awk '{print $NF}' | \
    sort | uniq -c | sort -rn | head -10
} > "$TMP"

mail -s "[fail2ban] rapport quotidien $HOST" "$DEST" < "$TMP"
rm -f "$TMP"
```

```cron
# /etc/cron.d/fail2ban-report
0 7 * * * root /usr/local/bin/fail2ban-daily-report.sh
```

### Débannir une IP partout (tous les jails)

```bash
#!/bin/bash
# /usr/local/bin/f2b-unban-all.sh
IP="$1"
[ -z "$IP" ] && { echo "Usage: $0 <IP>"; exit 1; }
for j in $(sudo fail2ban-client status | grep "Jail list" | sed 's/.*Jail list://;s/,//g'); do
  sudo fail2ban-client set "$j" unbanip "$IP" 2>/dev/null && echo "débannie de $j"
done
```

### Exporter toutes les IP bannies (pour audit)

```bash
for j in $(sudo fail2ban-client status | grep "Jail list" | sed 's/.*Jail list://;s/,//g'); do
  echo "## $j"
  sudo fail2ban-client status "$j" | grep -A100 "Banned IP list" | tail -n +2
done > /root/fail2ban-banned-$(date +%Y%m%d).txt
```

### Compter les bans par heure (détecter les campagnes)

```bash
sudo grep " Ban " /var/log/fail2ban.log | \
  awk '{print $1" "$2}' | cut -d: -f1 | sort | uniq -c | sort -rn | head -24
# 2026-09-26 03  →  47 bans entre 3h et 4h : campagne nocturne typique
```

---

## 49. Intégration Ansible

### Rôle minimal

```yaml
# roles/fail2ban/tasks/main.yml
- name: Installer fail2ban
  apt:
    name: fail2ban
    state: present
    update_cache: true

- name: Déployer jail.local
  template:
    src: jail.local.j2
    dest: /etc/fail2ban/jail.local
    owner: root
    group: root
    mode: "0644"
  notify: reload fail2ban

- name: Déployer les filters maison
  copy:
    src: "{{ item }}"
    dest: "/etc/fail2ban/filter.d/{{ item }}"
    mode: "0644"
  loop:
    - myapp.conf
  notify: reload fail2ban

- name: Activer et démarrer fail2ban
  systemd:
    name: fail2ban
    enabled: true
    state: started
```

```yaml
# roles/fail2ban/handlers/main.yml
- name: reload fail2ban
  command: fail2ban-client reload
  # Plutôt qu'un restart : les bans actifs sont conservés
```

### Template jail.local avec variables par groupe

```jinja2
# roles/fail2ban/templates/jail.local.j2
[DEFAULT]
bantime  = {{ fail2ban_bantime | default('1h') }}
findtime = {{ fail2ban_findtime | default('10m') }}
maxretry = {{ fail2ban_maxretry | default(5) }}
ignoreip = 127.0.0.1/8 ::1 {{ fail2ban_ignoreip_extra | default('') }}
banaction = {{ fail2ban_banaction | default('nftables-multiport') }}

[sshd]
enabled = true

{% if 'webservers' in group_names %}
[nginx-http-auth]
enabled = true
port    = http,https
filter  = nginx-http-auth
logpath = /var/log/nginx/error.log

[nginx-botsearch]
enabled  = true
filter   = nginx-botsearch
logpath  = /var/log/nginx/access.log
maxretry = 3
findtime = 1d
bantime  = 1w
{% endif %}

{% if 'mailservers' in group_names %}
[postfix]
enabled = true
filter  = postfix
logpath = /var/log/mail.log

[dovecot]
enabled  = true
filter  = dovecot
logpath = /var/log/mail.log
maxretry = 10
bantime  = 30m
{% endif %}

[recidive]
enabled   = true
filter    = recidive
logpath   = /var/log/fail2ban.log
banaction = %(banaction_allports)s
bantime   = 1w
findtime  = 1d
maxretry  = 5
```

### Vérification post-déploiement

```yaml
- name: Vérifier les jails actifs
  command: fail2ban-client status
  register: f2b_status
  changed_when: false

- name: Afficher
  debug:
    var: f2b_status.stdout_lines
```

---

## 50. Intégration avec GeoIP

### Cas d'usage

Vous n'avez aucun utilisateur légitime hors d'Europe ? Vous pouvez **alerter** (ou bannir) les IP de pays inattendus. Attention : la géolocalisation d'IP est approximative, et les attaquants utilisent des VPS européens. C'est un **signal**, pas une preuve.

