---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-11
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [2387, 2624]
sha256: a25e3cf5f76bf1797f23b451fcd848e7827c11c39cd4bbbccdaa540ede81f06c
---

# Guide fail2ban — Le bouclier anti-brute-force

**Symptôme :** tickets au support : « je suis bloqué » — utilisateurs légitimes (dovecot, webmail, VPN).
**Cause :** `maxretry = 3` + smartphone qui réessaie en boucle.
**Solution :** seuils différenciés (section 23) : `maxretry = 10`, `bantime = 30m` sur dovecot. Communiquez auprès des utilisateurs.

### Erreur n°10 : oublier recidive (ou mal le configurer)

**Symptôme :** les mêmes IP reviennent toutes les heures, indéfiniment.
**Cause :** pas de jail recidive, ou recidive qui pointe vers un `logpath` inexistant.
**Solution :** activez recidive (section 14), vérifiez qu'il lit bien `/var/log/fail2ban.log`.

### Erreur n°11 : loglevel DEBUG oublié en production

**Symptôme :** `/var` plein à 100 %, services qui tombent.
**Cause :** `loglevel = DEBUG` laissé après un dépannage ; sur un serveur attaqué, ça génère des Go/heure.
**Solution :** remettez `INFO`, mettez en place une alerte sur l'espace disque ET sur la taille de `/var/log/fail2ban.log` :
```bash
# Alerte simple via cron :
find /var/log/fail2ban.log -size +500M -exec echo "ALERTE fail2ban.log trop gros" | mail -s alerte admin@example.com \;
```

### Erreur n°12 : tester en production sans filet

**Symptôme :** tout ce qui précède, en même temps, un vendredi à 18h.
**Cause :** reload d'une config non testée, sans console de secours, sans dead man's switch.
**Solution :** la discipline des sections 25 et 31 : `--test`, `fail2ban-regex`, bantime court pendant les essais, console de secours ouverte, jamais le vendredi soir.

---

## 41. Backend : auto, systemd, pyinotify, polling

Le **backend** est la méthode par laquelle fail2ban surveille les logs. Le choix impacte la fiabilité.

### Les backends disponibles

| Backend | Méthode | Avantages | Inconvénients |
|---|---|---|---|
| `auto` | choisit le meilleur dispo | zéro config, convient à 95 % | choix parfois surprenant |
| `systemd` | journald natif | pas de fichier à gérer, métadonnées riches | que pour les services systemd |
| `pyinotify` | inotify (événements noyau) | temps réel, peu de CPU | ne voit pas les rotations mal gérées |
| `gamin` | Gamin (similaire) | temps réel | dépendance supplémentaire |
| `polling` | relecture périodique | fonctionne partout | CPU +, latence |

Ordre de `auto` : pyinotify → gamin → polling → systemd.

### Quand forcer systemd

```ini
[postfix]
enabled      = true
backend      = systemd
journalmatch = _SYSTEMD_UNIT=postfix.service + _SYSTEMD_UNIT=postfix@-.service
```

Cas d'usage : pas de rsyslog installé (Debian minimale, containers), logs uniquement dans le journal.

Vérifier que le journal contient bien les lignes :

```bash
sudo journalctl -u postfix --since "10 min ago" | grep -i "sasl.*failed" | head -3
```

### Quand forcer polling

Rare : systèmes de fichiers réseau (NFS) où inotify ne fonctionne pas.

```ini
[monjail]
backend = polling
```

### Diagnostiquer le backend actif

```bash
sudo grep -i "backend" /var/log/fail2ban.log | tail -5
# ... INFO [sshd] Using backend auto -> polling (par ex.)
sudo fail2ban-client get sshd backend
```

---

## 42. Le fuseau horaire : la cause cachée

### Le problème

fail2ban lit un timestamp dans chaque ligne de log et le compare à l'heure système pour `findtime`. Si les deux ne sont pas dans le même fuseau :

- Logs en UTC, système en Europe/Paris (UTC+2 en été) : un échec « à 10h00 UTC » est vu comme « il y a 2h » → il sort immédiatement de la fenêtre `findtime = 10m` → **aucun ban ne se déclenche jamais**.
- L'inverse : des échecs vieux de 2h sont comptés comme récents → **bans abusifs**.

### Diagnostic

```bash
timedatectl | grep -E "Time zone|Universal time|Local time"
head -2 /var/log/auth.log
# Sep 26 10:00:01 ...  → heure locale (rsyslog utilise l'heure système par défaut)
date -u && date
```

Si `date` (locale) et les timestamps des logs diffèrent d'un nombre rond d'heures → problème de fuseau.

### Cas typiques

| Source de log | Fuseau habituel | Solution |
|---|---|---|
| rsyslog local | heure système | OK si système à l'heure |
| Application Java/Node | souvent UTC | configurer l'app en heure locale, ou accepter le décalage |
| Conteneur Docker | UTC par défaut | monter `/etc/localtime` dans le conteneur |
| Log applicatif avec `Z` | UTC explicite | fail2ban gère le `Z` s'il est dans le format de date |

### Bonne pratique

Mettez **tout en UTC** sur les serveurs (recommandation sysadmin classique) :

```bash
sudo timedatectl set-timezone UTC
```

fail2ban, rsyslog et le système sont alors cohérents. Les humains convertissent à la lecture (les outils le font pour vous).

---

## 43. IPv6 avec fail2ban

### Ça marche nativement

`<HOST>` matche les IPv6. Rien à configurer pour la détection.

### Points d'attention spécifiques

1. **`ignoreip` doit inclure `::1`** (loopback v6) — sinon un service local en v6 peut se faire bannir.
2. **Les actions firewall doivent gérer v6** : `iptables-multiport` ne gère que v4 ! Pour v6 :
   ```ini
   [DEFAULT]
   banaction = nftables-multiport   # gère v4 + v6 nativement
   ```
   Avec nftables, pas de problème. Avec iptables classique, il faudrait aussi `ip6tables` — d'où l'intérêt de nftables.
3. **Les attaquants v6 ont des /64 entiers** : bannir une seule IPv6 est parfois inutile (l'attaquant en a 2^64). Pour les cas extrêmes, certaines actions permettent de bannir le /64 (personnalisation avancée, hors scope ici — à noter comme piste).

### Vérifier

```bash
# Forcer un ban de test en IPv6 (depuis une machine v6) puis :
sudo nft list set inet f2b-table f2b-sshd | grep -i ":"
```

---

## 44. Bannir via ipset pour la performance

Quand le nombre de bans simultanés dépasse ~1000, les chaînes iptables linéaires ralentissent. `ipset` stocke les IP dans une table de hachage : O(1) au lieu de O(n).

### Installation et configuration

```bash
sudo apt install -y ipset
```

```ini
[DEFAULT]
banaction = iptables-ipset-proto6
```

Variantes : `iptables-ipset-proto6` (tous protocoles, ports multiples), `iptables-ipset-proto6-allports`.

### Vérification

```bash
sudo ipset list | head -30
# Name: f2b-sshd
# Type: hash:net
# Members:
# 203.0.113.45
```

### Quand en avez-vous besoin ?

| Bans simultanés | Action conseillée |
|---|---|
| < 200 | nftables-multiport (défaut moderne) |
| 200 – 2000 | nftables-multiport tient encore très bien |
| > 2000 | ipset ou nftables avec sets (les sets nftables sont déjà en hachage) |

> Bonne nouvelle : les **sets nftables** utilisés par `nftables-multiport` sont déjà des tables de hachage. Sur les systèmes récents, vous avez la performance d'ipset sans rien faire.

---

## 45. Rotation de logs et fail2ban

### Le problème

`logrotate` compresse/renomme `/var/log/auth.log` → `auth.log.1`. Si fail2ban suit l'ancien descripteur de fichier, il ne voit plus les nouvelles attaques.

### Pourquoi ça marche (en général)

- Avec `pyinotify`/`systemd`, fail2ban détecte la rotation et rouvre le fichier.
- Les configurations logrotate de Debian/Ubuntu utilisent `copytruncate` ou `create` avec un `postrotate` qui notifie — compatibles.

Vérifiez votre logrotate :

```bash
cat /etc/logrotate.d/rsyslog | head -20
# /var/log/auth.log {
#   rotate 4
#   weekly
#   ...
#   postrotate
#     /usr/lib/rsyslog/rsyslog-rotate
#   endscript
# }
```

### Tester après une rotation forcée

```bash
sudo logrotate -f /etc/logrotate.d/rsyslog
sleep 5
sudo fail2ban-client status sshd | grep -E "File list|Currently"
# Si le jail ne suit plus le nouveau fichier : restart du jail
sudo fail2ban-client reload sshd
```

### Le cas des logs applicatifs maison

Si votre application écrit dans `/var/log/myapp/auth.log`, créez son logrotate :

```
# /etc/logrotate.d/myapp
/var/log/myapp/*.log {
  daily
  rotate 14
  compress
  delaycompress
  missingok
  notifempty
  copytruncate
}
```

