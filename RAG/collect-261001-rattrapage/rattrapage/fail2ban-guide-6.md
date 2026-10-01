---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-6
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [1211, 1443]
sha256: 9e809a52e78824f9626d0bcfd4fdd9817d2f0ce11d61f23d03d886a24dbb9c38
---

# Guide fail2ban — Le bouclier anti-brute-force

```bash
sudo iptables -L f2b-sshd -n --line-numbers
# Chain f2b-sshd (1 references)
# num  target  prot  source        destination
# 1    REJECT  all   203.0.113.45  0.0.0.0/0  reject-with icmp-port-unreachable
```

### REJECT vs DROP

- `iptables-multiport` utilise **REJECT** par défaut : l'attaquant reçoit un ICMP d'erreur, sa connexion échoue vite, il passe à autre chose.
- `iptables` tout court (action de base) utilise **DROP** : les paquets sont jetés silencieusement, l'attaquant attend le timeout.

> En pratique, REJECT fait perdre moins de temps aux scanners (ils abandonnent vite) et rend vos propres tests plus lisibles (connexion refusée immédiatement plutôt que timeout).

---

## 20. Action nftables en détail

Sur les systèmes récents, nftables est le backend natif. Comprendre ce que fail2ban y fait aide au dépannage.

### Inspection

```bash
# Voir toutes les tables fail2ban
sudo nft list tables | grep f2b

# Détail d'une chaîne de jail
sudo nft list chain inet f2b-table f2b-sshd

# Compter les IP bannies par jail
sudo nft list set inet f2b-table f2b-sshd 2>/dev/null | tr ',' '\n' | grep -c "203\."
```

### Fichier d'action

`/etc/fail2ban/action.d/nftables.conf` définit les variantes :

| Action | Portée |
|---|---|
| `nftables` | un seul port (`<port>`) |
| `nftables-multiport` | liste de ports (`<port>` = `http,https`) |
| `nftables-allports` | tous les ports |

```ini
# Extrait de nftables-common.conf
actionban = nft add element <table> <chain> { <ip> }
actionunban = nft delete element <table> <chain> { <ip> }
```

### Conflit avec docker ou d'autres outils nftables

Docker manipule aussi nftables/iptables. Si fail2ban et Docker cohabitent :

- fail2ban crée sa propre table `f2b-table` : pas de conflit direct.
- Mais l'ordre des hooks compte : les règles Docker s'appliquent dans leurs propres chaînes.
- **Testez** : bannissez une IP de test et vérifiez qu'elle est bien bloquée vers le conteneur.

### Nettoyage manuel (si fail2ban a planté)

```bash
# Si des règles orphelines subsistent après un crash :
sudo nft delete table inet f2b-table
sudo fail2ban-client restart
```

---

## 21. Action ufw

Si votre serveur utilise UFW (`sudo ufw enable`), utilisez l'action `ufw` pour rester cohérent.

```ini
[DEFAULT]
banaction = ufw
```

Ce que fait l'action :

```bash
# actionban :
sudo ufw insert 1 deny from 203.0.113.45 to any
# actionunban :
sudo ufw delete deny from 203.0.113.45 to any
```

Vérification :

```bash
sudo ufw status numbered | head -20
```

### Limites de l'action ufw

- Elle insère des règles **en tête** (`insert 1`) : les bans priment sur vos règles `allow`. C'est voulu.
- Sur un serveur avec beaucoup de bans simultanés, `ufw` (qui régénère tout le ruleset à chaque changement) est **plus lent** que nftables direct. Au-delà de ~500 bans actifs, préférez nftables.
- Ne mélangez pas `banaction = ufw` et des règles iptables manuelles : UFW ne les verra pas.

---

## 22. Actions multiples et combinées

Un jail peut déclencher **plusieurs actions** : bannir + alerter, par exemple.

### Syntaxe

```ini
[sshd]
enabled = true
action  = %(action_)s
          %(action_mail)s[name=sshd, dest=root@example.com]
```

`%(action_)s` = l'action de bannissement par défaut (définie par `banaction`).
`%(action_mail)s` = bannir + envoyer un mail via `action.d/mail.conf`.

### Actions d'alerte disponibles

| Action | Effet |
|---|---|
| `action_` | ban seul (défaut) |
| `action_mw` | ban + mail avec extrait de log (whois) |
| `action_mwl` | ban + mail avec whois + lignes de log |
| `action_cf_mwl` | ban + mail + ban Cloudflare (section 32) |

Exemple complet :

```ini
[sshd]
enabled = true
action  = %(action_)s
          %(action_mwl)s[name=sshd, dest=admin@example.com, sender=fail2ban@srv-web.example.com]
```

Prérequis mail : un MTA local fonctionnel (`postfix` en mode satellite, ou `ssmtp`) :

```bash
echo "test" | mail -s "test fail2ban" admin@example.com
```

### Webhook générique (exemple Slack/Mattermost)

Créez `/etc/fail2ban/action.d/webhook.conf` :

```ini
[Definition]
actionban = curl -s -X POST -H 'Content-Type: application/json' \
              -d '{"text":"[fail2ban] <name> : bannissement de <ip> (échecs: <failures>)"}' \
              https://hooks.exemple.fr/services/TOKEN
actionunban =
```

```ini
[sshd]
action = %(action_)s
         webhook[name=sshd]
```

> Attention aux secrets dans les fichiers d'action : `webhook.conf` contient un token. `chmod 640` + propriétaire root, et ne le commitez jamais en clair.

---

## 23. Tuning : bantime, findtime, maxretry par service

Il n'existe pas de valeurs universelles, mais voici des bases éprouvées en production, avec le raisonnement.

### Tableau de référence

| Jail | maxretry | findtime | bantime | Raison |
|---|---|---|---|---|
| sshd | 5 | 10m | 1h | pas d'humains légitimes en faute ; bots agressifs |
| sshd (parano) | 3 | 10m | 4h | serveurs ultra-exposés |
| nginx-http-auth | 5 | 10m | 1h | humains possibles derrière le basic auth |
| nginx-botsearch | 3 | 1d | 1w | aucun humain légitime ne scanne `/.env` |
| nginx-badbot | 2 | 10m | 1w | user-agent = signature quasi certaine |
| postfix / postfix-sasl | 5 | 10m | 1h | bots de relay + brute-force SASL |
| dovecot | 10 | 10m | 30m | **smartphones mal configurés** : seuil haut, ban court |
| proftpd / vsftpd | 5 | 10m | 1h | legacy, bots uniquement |
| recidive | 5 | 1d | 1w | deuxième ligne, tous ports |
| wordpress* | 5 | 10m | 1h | `wp-login.php` pilonné |

\* nécessite un filter wordpress (non fourni par défaut sur les vieilles versions ; voir section 17 pour l'écrire).

### Méthode de tuning en 4 étapes

1. **Démarrez avec les valeurs du tableau** (conservatrices).
2. **Observez une semaine** : `grep "Ban " /var/log/fail2ban.log | awk '{print $7}' | sort | uniq -c | sort -rn | head` — qui se fait bannir, et combien de fois ?
3. **Repérez les faux positifs** : une IP interne ? un utilisateur légitime ? → augmentez `maxretry` ou ajoutez à `ignoreip`.
4. **Repérez les attaquants persistants** : une IP qui revient après chaque unban → elle finira dans `recidive` ; si besoin, durcissez `bantime`.

### Le piège du bantime trop long

Un `bantime = 1w` sur le jail sshd semble sécurisant, mais :

- Les botnets utilisent des IP jetables : l'IP bannie ne reviendra jamais, la règle firewall reste pour rien.
- À 10 000 bans/semaine, la table nftables/iptables grossit et **ralentit chaque paquet**.
- Solution : `bantime` modéré (1h) + `recidive` agressif (1w) pour les vrais persistants.

### bantime incrémental (fail2ban ≥ 0.11)

```ini
[DEFAULT]
# Le bantime augmente à chaque récidive : 1h, 2h, 4h, 8h... (facteur 2)
bantime.increment = true
bantime.factor = 2
bantime.maxtime = 4w
# Formule : bantime * factor ^ (nombre de bans précédents)
```

Avec `bantime = 1h` : 1er ban = 1h, 2e = 2h, 3e = 4h, 4e = 8h... plafonné à 4 semaines. C'est une excellente alternative (ou un complément) au jail recidive.

> Note : `bantime.increment` utilise la base sqlite (`dbfile`) pour se souvenir des bans précédents. Vérifiez que `dbfile = /var/lib/fail2ban/fail2ban.sqlite3` est actif (défaut).

---

## 24. ignoreip : whitelist (son réseau, la supervision !)

`ignoreip` est le paramètre **le plus important pour votre tranquillité** : les IP listées ne seront JAMAIS bannies, même si elles dépassent `maxretry`.

### Qui mettre en whitelist, systématiquement

```ini
[DEFAULT]
ignoreip = 127.0.0.1/8 ::1
           192.168.10.0/24      # LAN bureautique
           10.20.0.0/16         # LAN serveurs
           203.0.113.10         # IP publique du bureau (fixe)
           198.51.100.5         # sonde Zabbix
           192.0.2.53           # sonde Prometheus/blackbox
           203.0.113.200        # bastion d'administration
```

Checklist des oublis classiques :

