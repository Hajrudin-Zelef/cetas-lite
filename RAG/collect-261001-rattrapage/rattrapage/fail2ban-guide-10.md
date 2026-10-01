---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-10
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "attention"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [2193, 2386]
sha256: 34beacf29c7afb588d1003a20e5d5ac4d1876410569471dd924276b650481d1e
---

# Guide fail2ban — Le bouclier anti-brute-force

```ini
# /etc/fail2ban/action.d/webhook.conf — avec timeout et silence
[Definition]
actionban = curl -s -m 10 -X POST -H 'Content-Type: application/json' \
  -d '{"text":"🚨 [fail2ban] <name> : <ip> bannie (<failures> échecs)"}' \
  https://mattermost.exemple.fr/hooks/TOKEN >/dev/null 2>&1 || true
actionunban =
```

Le `|| true` évite que l'échec du webhook fasse échouer l'action de ban (une alerte ne doit jamais casser la protection).

---

## 37. Sauvegarde et restauration de jail.local

### Ce qu'il faut sauvegarder

| Fichier | Criticité | Note |
|---|---|---|
| `/etc/fail2ban/jail.local` | **critique** | toute votre config |
| `/etc/fail2ban/jail.d/*.local` | critique | si vous découpez |
| `/etc/fail2ban/filter.d/mon*.conf` | critique | vos filters maison |
| `/etc/fail2ban/action.d/mon*.conf` | important | vos actions maison |
| `/var/lib/fail2ban/fail2ban.sqlite3` | pratique | historique des bans |

### Script de sauvegarde

```bash
#!/bin/bash
# /usr/local/bin/backup-fail2ban.sh
DEST=/root/backups/fail2ban
DATE=$(date +%Y%m%d)
mkdir -p "$DEST"
tar -czf "$DEST/fail2ban-$DATE.tar.gz" \
  /etc/fail2ban/jail.local \
  /etc/fail2ban/jail.d/ \
  /etc/fail2ban/fail2ban.d/ \
  $(ls /etc/fail2ban/filter.d/*.conf 2>/dev/null | grep -v "^/etc/fail2ban/filter.d/$") \
  2>/dev/null
# Ne sauvegarde que jail.local + jail.d + fail2ban.d par défaut :
tar -czf "$DEST/fail2ban-essentiel-$DATE.tar.gz" \
  /etc/fail2ban/jail.local /etc/fail2ban/jail.d/ /etc/fail2ban/fail2ban.d/
ls -la "$DEST" | tail -5
```

> Simplifiez : sauvegardez `/etc/fail2ban/jail.local`, `/etc/fail2ban/jail.d/`, `/etc/fail2ban/fail2ban.d/` et vos fichiers `filter.d`/`action.d` maison. Le reste est fourni par le paquet.

### Versionner avec git (recommandé)

```bash
cd /etc/fail2ban
git init
git add jail.local jail.d/ fail2ban.d/
git commit -m "config initiale $(date +%Y-%m-%d)"
# Après chaque modification :
git diff && git commit -am "décrit le changement"
```

Un `git log` sur fail2ban vaut de l'or pendant un audit ou un dépannage (« qui a changé maxretry le 14 ? »).

### Restauration

```bash
# Restaurer les fichiers, puis :
sudo fail2ban-client --test   # valide
sudo fail2ban-client reload   # applique
sudo fail2ban-client status   # contrôle
```

---

## 38. Mise à jour de fail2ban

### Mise à jour via le paquet (cas normal)

```bash
sudo apt update && sudo apt install --only-upgrade fail2ban
```

Points d'attention :

1. **`jail.conf` peut être écrasé** : c'est normal, votre config est dans `jail.local` (qui n'est jamais touché).
2. **Nouveaux filters** : la mise à jour peut ajouter des filters dans `filter.d/`. Relisez les notes de version si vous avez surchargé un filter du même nom.
3. **Redémarrage du service** : le paquet redémarre fail2ban. Les bans actifs sont restaurés via sqlite (section 29). Vérifiez après :
   ```bash
   sudo fail2ban-client status
   sudo fail2ban-client status sshd | grep -E "Currently banned"
   ```

### Vérifier ce qui a changé

```bash
# Debian : voir les fichiers modifiés par la mise à jour
sudo apt changelog fail2ban | head -40
# ou :
zcat /usr/share/doc/fail2ban/changelog.Debian.gz | head -40
```

### Mise à jour majeure (0.11 → 1.x) : points de vigilance

- Le backend par défaut et les actions nftables ont évolué : relisez votre `banaction`.
- Certaines options dépréciées : consultez le log au redémarrage (`grep -i deprecat /var/log/fail2ban.log`).
- Testez en pré-production avant de déployer sur le parc.

### Stratégie de déploiement sur un parc

1. MaJ sur un serveur pilote.
2. `fail2ban-client --test` + 48h d'observation.
3. Déploiement Ansible sur le parc (section 55), par vagues.
4. Vérification centralisée : le monitoring (sections 34-35) doit montrer les jails actifs partout.

---

## 39. Erreurs classiques (1/2) : les 6 premières

### Erreur n°1 : se bannir soi-même

**Symptôme :** plus d'accès SSH au serveur juste après un reload.
**Cause :** votre IP n'est pas dans `ignoreip`, et vos tests (ou votre client qui réessaie) ont dépassé `maxretry`.
**Solution immédiate :** console hors-bande → `fail2ban-client set sshd unbanip VOTRE_IP`.
**Prévention :** sections 24 et 25. `ignoreip` complet + console de secours + dead man's switch.

### Erreur n°2 : le filter ne matche pas le format de log

**Symptôme :** `Currently failed: 0` alors que les attaques sont visibles dans le log.
**Causes fréquentes :**
- `log_format` nginx personnalisé sans les champs attendus (section 10).
- Application qui loggue en JSON alors que le filter attend du texte.
- Backend qui lit le mauvais fichier (auth.log vide, tout est dans le journal).
**Solution :** `fail2ban-regex` avec `--print-all-missed`, comparez ligne par ligne (section 16).

### Erreur n°3 : mauvais logpath

**Symptôme :** jail actif, 0 échec détecté, le fichier de log est vide ou n'existe pas.
**Exemples réels :**
- proftpd qui loggue dans `/var/log/auth.log` au lieu de `/var/log/proftpd/proftpd.log`.
- postfix sur un système sans rsyslog : tout est dans `journalctl`, `/var/log/mail.log` n'existe pas.
**Solution :** trouvez où l'application écrit VRAIMENT :
```bash
sudo lsof -p $(pgrep -f "dovecot/auth" | head -1) 2>/dev/null | grep log
# ou plus simple : provoquez un échec et cherchez où il apparaît
sudo grep -r "authentication failed" /var/log/ 2>/dev/null | head -5
```

### Erreur n°4 : backend systemd vs auto

**Symptôme :** sur un système récent, le jail ne voit rien bien que le log existe.
**Cause :** `backend = auto` choisit `pyinotify`, mais les logs arrivent via journald sans passer par le fichier (ou l'inverse).
**Solution :** forcez le backend adapté :
```ini
[monjail]
backend = systemd
# et pour cibler l'unité :
journalmatch = _SYSTEMD_UNIT=monservice.service
```
Voir section 45 pour le détail des backends.

### Erreur n°5 : fuseau horaire incohérent

**Symptôme :** les bans se déclenchent avec un retard bizarre, ou `findtime` semble ne pas fonctionner (les échecs ne sont pas comptés ensemble).
**Cause :** les logs sont en UTC, le système en Europe/Paris (ou l'inverse). fail2ban compare les timestamps du log avec l'heure système.
**Solution :** harmonisez (section 46) :
```bash
timedatectl   # vérifiez "Time zone" et "Universal time"
head -1 /var/log/auth.log   # quel fuseau dans les logs ?
```

### Erreur n°6 : le ban ne bloque rien (mauvais port)

**Symptôme :** `Ban 203.0.113.45` dans le log, mais l'attaquant continue.
**Causes :**
- `port = ssh` alors que sshd écoute sur 2222 → la règle firewall bloque le port 22, pas le 2222.
- Action `iptables` alors que le système utilise nftables en natif (règles créées au mauvais endroit).
- Docker qui bypasse les chaînes INPUT.
**Solution :** vérifiez le port réel (`ss -tlnp | grep sshd`), vérifiez la règle créée (section 19), testez depuis l'extérieur.

---

## 40. Erreurs classiques (2/2) : les 6 suivantes

### Erreur n°7 : modifier jail.conf au lieu de jail.local

**Symptôme :** après `apt upgrade`, toute la config a disparu.
**Cause :** `jail.conf` est écrasé par le paquet.
**Solution :** tout migrer dans `jail.local` (section 4). Pour récupérer : si vous avez un backup ou un git, restaurez ; sinon, réécrivez (douloureux, mais formateur).

### Erreur n°8 : bannir les IP de Cloudflare / du reverse proxy

**Symptôme :** des centaines d'utilisateurs légitimes bloqués d'un coup, ou le site entier inaccessible.
**Cause :** les logs voient l'IP du proxy/CDN, fail2ban bannit le proxy/CDN.
**Solution :** `set_real_ip_from` + `real_ip_header` (nginx, section 32/52) ou `RemoteIPHeader` (Apache). Et mettez les IP du proxy en `ignoreip` en attendant.

### Erreur n°9 : maxretry trop bas sur un service à humains

