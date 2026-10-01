---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-8
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [1713, 1965]
sha256: ccafdd42af17ee0a4bf7a2a1dcef166bc0cadf4e0c8465ac8ce5598a6fa8407d
---

# Guide fail2ban — Le bouclier anti-brute-force

**Quand c'est dangereux :**
- Sur `sshd` ou `dovecot` avec des humains : un faux positif permanent = ticket au support + intervention manuelle.

### Stratégie recommandée : l'escalade

Plutôt que du permanent brutal, utilisez l'escalade automatique :

```ini
[DEFAULT]
bantime = 1h
# Escalade : 1h → 2h → 4h → ... → max 4 semaines
bantime.increment = true
bantime.factor = 2
bantime.maxtime = 4w
bantime.rndtime = 10m   # petite randomisation (évite les patterns prévisibles)
```

Combiné au jail `recidive` (ban 1 semaine tous ports après 5 récidives), vous obtenez une réponse graduée sans administration manuelle.

---

## 29. Bans persistants : la base sqlite

### Comment ça marche

```ini
# /etc/fail2ban/fail2ban.conf (défaut, ne pas toucher sauf besoin)
[Definition]
dbfile = /var/lib/fail2ban/fail2ban.sqlite3
dbpurgeage = 1d
```

- Chaque ban est enregistré dans sqlite avec son timestamp.
- Au (re)démarrage, fail2ban **restaure les bans non expirés**.
- `dbpurgeage = 1d` : les bans expirés depuis plus d'1 jour sont purgés de la base.

### Vérifier la base

```bash
sudo sqlite3 /var/lib/fail2ban/fail2ban.sqlite3 "SELECT name, COUNT(*) FROM bans GROUP BY name;"
# sshd|12
# recidive|3

sudo sqlite3 /var/lib/fail2ban/fail2ban.sqlite3 \
  "SELECT ip, jail, timeofban, bantime FROM bans LIMIT 5;"
```

### Désactiver la persistance (rare)

```ini
# fail2ban.conf — si vous voulez repartir de zéro à chaque boot
dbfile = :memory:
```

> À éviter en production : après un reboot (mise à jour noyau un dimanche à 3h), tous les attaquants seraient de nouveau libres de frapper pendant que vous dormez.

### Sauvegarder la base

Elle fait partie de la sauvegarde système, mais elle est **reconstruite automatiquement** si absente. Pas critique, mais pratique.

---

## 30. Logs et niveaux de debug

### Où et quoi

| Log | Contenu | Config |
|---|---|---|
| `/var/log/fail2ban.log` | bans, unbans, erreurs, démarrage des jails | `logtarget` dans fail2ban.conf |
| `journalctl -u fail2ban` | idem si `logtarget = SYSTEMD-JOURNAL` | — |
| Les logs surveillés | les attaques elles-mêmes | `logpath` de chaque jail |

### Niveaux de log

```ini
# /etc/fail2ban/fail2ban.d/debug.local — temporaire !
[Definition]
loglevel = DEBUG
```

| Niveau | Usage |
|---|---|
| `INFO` | production (défaut raisonnable) |
| `NOTICE` | production verbeuse (bans/unbans visibles) |
| `WARNING` | erreurs uniquement |
| `DEBUG` | dépannage : montre chaque ligne matchée, chaque décision |

> `DEBUG` est **très verbeux** : ne le laissez jamais en production (le log grossit de plusieurs Go/jour sur un serveur attaqué). Activez, reproduisez, désactivez.

### Lire efficacement le log

```bash
# Les 20 derniers bans
sudo grep " Ban " /var/log/fail2ban.log | tail -20

# Les erreurs (config, regex, actions)
sudo grep -iE "error|exception|failed" /var/log/fail2ban.log | tail -20

# Activité d'un jail sur la dernière heure
sudo awk -v d="$(date -d '1 hour ago' '+%Y-%m-%d %H')" \
  '$0 > d' /var/log/fail2ban.log | grep "\[sshd\]" | tail -30

# Compter les bans par jail aujourd'hui
sudo grep "$(date +%Y-%m-%d).* Ban " /var/log/fail2ban.log \
  | awk -F'[][]' '{print $2}' | sort | uniq -c | sort -rn
```

### Augmenter temporairement le debug d'un jail

```bash
# Passer tout le démon en DEBUG sans éditer de fichier :
sudo fail2ban-client set loglevel DEBUG
# ... reproduire le problème ...
sudo fail2ban-client set loglevel INFO
sudo grep -i "dovecot" /var/log/fail2ban.log | tail -40
```

---

## 31. Troubleshooting des logs : méthode

Quand « fail2ban ne bannit pas », suivez cette méthode dans l'ordre. 95 % des cas se résolvent aux étapes 1-3.

### Étape 1 : le jail tourne-t-il ?

```bash
sudo fail2ban-client status monjail
# Si "Sorry but the jail 'monjail' does not exist" → le jail n'est pas chargé.
# Vérifiez : enabled = true ? faute de frappe dans le nom ? reload effectué ?
sudo fail2ban-client status | grep "Jail list"
```

### Étape 2 : le log est-il lu ?

```bash
sudo fail2ban-client status monjail | grep -A2 "File list"
# File list: /var/log/auth.log
ls -la /var/log/auth.log   # le fichier existe ? non vide ? lisible par root ?
```

Si `Currently failed` reste à 0 alors que le log contient des échecs → le filter ne matche pas (étape 3).

### Étape 3 : le filter matche-t-il ?

```bash
sudo fail2ban-regex /var/log/auth.log /etc/fail2ban/filter.d/sshd.conf | tail -2
# Lines: 5000 lines, 0 ignored, 0 matched, 5000 missed
# → 0 matched = le filter ne voit rien. Vérifiez le format de date (section 46)
#    et le format des lignes (section 18).
```

### Étape 4 : l'action s'exécute-t-elle ?

```bash
sudo grep -E "Ban |action.*error" /var/log/fail2ban.log | tail -10
# Si vous voyez des erreurs après "Ban" → l'action firewall échoue.
# Causes : nftables vs iptables mal choisi, ufw inactif, droits.
```

Vérifiez la règle réellement créée :

```bash
sudo nft list table inet f2b-table 2>/dev/null | head -20
# ou :
sudo iptables -L f2b-sshd -n 2>/dev/null | head -20
```

### Étape 5 : le ban est-il effectif réseau ?

```bash
# Depuis la machine attaquante (ou un conteneur de test) :
curl -m 5 -v http://VOTRE_SERVEUR/ ; echo "code: $?"
ssh -o ConnectTimeout=5 user@VOTRE_SERVEUR ; echo "code: $?"
```

### Grille de diagnostic rapide

| Symptôme | Cause probable | Section |
|---|---|---|
| Jail absent de `status` | `enabled` oublié / faute de frappe / pas de reload | 31 |
| `Currently failed: 0` malgré des attaques | filter ne matche pas / mauvais logpath | 16, 46 |
| `Total failed` augmente mais aucun ban | seuils trop hauts / findtime trop court | 6, 23 |
| `Ban` dans le log mais IP toujours joignable | mauvaise action / mauvais port / Docker | 19, 20 |
| Ban immédiat au démarrage | `ignoreip` incomplet (c'est vous) | 24, 25 |
| Tout fonctionnait, puis plus rien | rotation de log mal gérée / backend | 45, 51 |

---

## 32. Intégration Cloudflare : bannir à l'edge

Si votre site est derrière Cloudflare, bannir l'IP sur votre serveur ne suffit pas toujours : autant la bannir **chez Cloudflare**, au bord du réseau.

### Principe

```
Attaquant → Cloudflare (edge) → votre serveur (origin)
                                        │
                              fail2ban détecte
                                        │
                                        ▼
                              API Cloudflare : "bloque cette IP"
                                        │
                                        ▼
                              Attaquant bloqué AVANT d'atteindre l'origin
```

### Prérequis

1. Un token API Cloudflare (modèle : *Edit zone firewall* ou le modèle « Firewall »).
2. L'ID de votre zone.

### Action Cloudflare

`/etc/fail2ban/action.d/cloudflare.conf` est fournie. Configuration :

```ini
# /etc/fail2ban/action.d/cloudflare.local
[Definition]
cfuser  = admin@example.com
cftoken = VOTRE_TOKEN_API_CLOUDFLARE
cfzoneid = VOTRE_ZONE_ID
```

Puis dans le jail :

```ini
[nginx-botsearch]
enabled = true
action  = %(action_)s
          cloudflare[name=nginx-botsearch]
```

### Le problème de l'IP réelle (important !)

Derrière Cloudflare, vos logs nginx voient **l'IP de Cloudflare**, pas celle du visiteur. Il faut restaurer la vraie IP :

```nginx
# /etc/nginx/nginx.conf
set_real_ip_from 173.245.48.0/20;
set_real_ip_from 103.21.244.0/22;
# ... (toutes les plages Cloudflare : https://www.cloudflare.com/ips/)
real_ip_header CF-Connecting-IP;
# ou : real_ip_header X-Forwarded-For;
```

> Sans `set_real_ip_from`, fail2ban bannira les IP de Cloudflare lui-même — catastrophique. C'est l'erreur classique n°8.

### Alternative sans API : bannir localement quand même

Si vous ne voulez pas d'intégration API, le ban local reste utile : il protège l'origin contre les attaques directes (bypass du CDN en attaquant l'IP du serveur). Mais il ne bloquera pas via le CDN.

---

