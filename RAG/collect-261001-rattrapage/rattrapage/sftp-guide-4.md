---
id: collect-261001-rattrapage/rattrapage/sftp-guide-4
title: "Guide SFTP complet — OpenSSH, chroot, supervision et automatisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/sftp_guide.md
source_anchor: ""
source_lines: [697, 998]
sha256: b95f30278c05c9c3871ec813b950e2860dc45aad4178294f814dc57b3c035954
---

# Guide SFTP complet — OpenSSH, chroot, supervision et automatisation

> `-l INFO` donne les open/close sans les détails ; `-l VERBOSE` est le bon compromis.
> Évitez `-l DEBUG` en production (volume énorme).

---

## 31. Séparer les logs SFTP (rsyslog)

`/etc/rsyslog.d/40-sftp.conf` :

```
# Les logs du subsystem SFTP dans un fichier dédié
:programname, isequal, "sftp-server" /var/log/sftp.log
& stop
```

```bash
sudo systemctl restart rsyslog
sudo logrotate --debug /etc/logrotate.d/sftp 2>/dev/null  # tester si conf créée
```

`/etc/logrotate.d/sftp` :

```
/var/log/sftp.log {
    weekly
    rotate 12
    compress
    missingok
    notifempty
    create 0640 root adm
    sharedscripts
    postrotate
        /usr/lib/rsyslog/rsyslog-rotate 2>/dev/null || true
    endscript
}
```

---

## 32. Exploiter les logs : qui a transféré quoi ?

One-liners d'exploitation quotidienne :

```bash
# Transferts du jour (uploads = WRITE, downloads = READ)
sudo grep "sftp-server" /var/log/sftp.log | grep "$(date +%b' '%d)" | grep -E 'open|close'

# Volumes par fichier aujourd'hui
sudo awk '/close/ {print $NF, $(NF-2)}' /var/log/sftp.log | sort | uniq -c

# Tous les uploads d'un utilisateur sur 7 jours
sudo grep "sftp-server" /var/log/sftp.log | grep 'flags WRITE' | grep -B5 "sftp_acme"

# Fichiers supprimés (opération remove)
sudo grep -i "remove" /var/log/sftp.log
```

Script de rapport quotidien (`/usr/local/sbin/rapport_sftp.sh`) : voir § 71.

---

## 33. Supervision : ce qu'il faut surveiller

| Indicateur | Seuil d'alerte | Moyen |
|---|---|---|
| Service ssh actif | down | systemd / sonde |
| Espace disque `/srv` | > 80 % | quota + `df` |
| Quota par utilisateur | > 85 % | `repquota` / `xfs_quota` |
| Échecs d'authentification | > 20 / 5 min / IP | fail2ban (§ 55) |
| Transferts anormaux | volume > 3× la moyenne | rapport quotidien (§ 32) |
| Certificats/clés expirés | J-30 | inventaire des clés (§ 19) |
| Âge des fichiers en dépôt | > rétention définie | `find -mtime` |
| Intégrité `sshd_config` | modification inattendue | AIDE / checksum |

---

## 34. Sonde de service (check simple)

```bash
#!/bin/bash
# /usr/local/sbin/check_sftp.sh — sonde applicative SFTP
SFTP_USER="svc_sonde"
SFTP_HOST="sftp.entreprise.lan"
SFTP_KEY="/etc/sonde/id_sftp_sonde"
echo "ls" | sftp -b - -i "$SFTP_KEY" -o ConnectTimeout=10 \
  -o BatchMode=yes "$SFTP_USER@$SFTP_HOST" >/dev/null 2>&1
if [ $? -eq 0 ]; then echo "SFTP OK"; exit 0
else echo "SFTP KO"; exit 2; fi
```

- Compte `svc_sonde` en lecture seule sur un chroot dédié.
- `BatchMode=yes` : échoue proprement sans interaction (pas de prompt).
- À brancher sur Zabbix/Nagios/Prometheus (user parameter / exporter / node).

---

## 35. Alertes disque et quota

```bash
#!/bin/bash
# Alerte si /srv > 80 % ou si un quota dépasse 85 %
SEUIL=80
USAGE=$(df --output=pcent /srv | tail -1 | tr -dc '0-9')
[ "$USAGE" -ge "$SEUIL" ] && \
  echo "/srv à $USAGE %" | mail -s "[ALERTE] Disque SFTP" admin@entreprise.lan
```

Pour les quotas individuels, parsez `repquota -u /srv` en cron quotidien et alertez par utilisateur.

---

## 36. Rotation et conservation des journaux

- `auth.log` : rotation hebdomadaire par défaut (logrotate) — conservez 12 mois si vos obligations
  (bancaires, clients) l'exigent.
- `/var/log/sftp.log` : voir § 31 (12 semaines par défaut, à ajuster).
- Pensez à **centraliser** : transférez les logs vers un serveur syslog/SIEM
  (immuabilité en cas de compromission du serveur SFTP).

```bash
# rsyslog : envoi vers le collecteur (TLS recommandé)
# /etc/rsyslog.d/50-forward.conf
action(type="omfwd" target="logs.entreprise.lan" port="6514" protocol="tcp")
```

---

## 37. Pare-feu : n'ouvrir que le nécessaire

```bash
# UFW : autoriser SSH/SFTP uniquement
sudo ufw allow 22/tcp
sudo ufw enable
sudo ufw status verbose
```

Durcissement par source (recommandé pour les partenaires à IP fixe) :

```bash
# nftables : n'autoriser que les IP partenaires sur le port 22
sudo nft add rule inet filter input tcp dport 22 \
  ip saddr { 203.0.113.10, 198.51.100.0/24 } accept
sudo nft add rule inet filter input tcp dport 22 drop
```

Et dans `sshd_config`, double verrou par utilisateur :

```
Match User sftp_banque
    AllowUsers sftp_banque@203.0.113.10
```

> `AllowUsers`/`DenyUsers`/`AllowGroups` se placent dans la partie **globale**,
> pas dans un bloc `Match` (restriction d'OpenSSH).

---

## 38. fail2ban : bannir les attaquants

```bash
sudo apt install -y fail2ban
```

`/etc/fail2ban/jail.d/sshd.local` :

```ini
[sshd]
enabled = true
port = 22
maxretry = 5
findtime = 600
bantime = 3600
```

Commandes :

```bash
sudo fail2ban-client status sshd
sudo fail2ban-client set sshd banip 203.0.113.99     # bannir manuellement
sudo fail2ban-client set sshd unbanip 203.0.113.99
```

> Les comptes SFTP n'ayant pas de mot de passe, le brute-force par mot de passe est inopérant ;
> fail2ban protège surtout les comptes admin restants et réduit le bruit.

---

## 39. Pas de shell, pas de tunnel, pas de X11 (rappel)

Le bloc `Match Group sftpusers` doit contenir (voir § 10) :

```
ForceCommand internal-sftp
AllowTcpForwarding no
AllowAgentForwarding no
X11Forwarding no
PermitTunnel no
```

Vérifiez qu'un compte SFTP ne peut pas exécuter de commande :

```bash
ssh -i ~/.ssh/id_sftp_acme sftp_acme@sftp.entreprise.lan 'id'
# attendu : "This service allows sftp connections only." puis fermeture
```

---

## 40. Utilisateurs fictifs : ne jamais utiliser de vrais noms

Dans toute la documentation et les exemples de ce guide, les utilisateurs
(`sftp_acme`, `sftp_banque`, `svc_copieur`…) et les domaines (`entreprise.lan`,
`203.0.113.0/24`, `198.51.100.0/24`) sont **fictifs** (plages de documentation
RFC 5737 / RFC 2606). Adaptez à votre réalité sans jamais exposer de vrais
identifiants dans des exemples partagés.

---

## 41. Mode batch : sftp -b pour les scripts

Fichier de commandes `envoi.txt` :

```
cd depot
put /var/exports/rapport_quotidien.csv
bye
```

Exécution :

```bash
sftp -b envoi.txt -i ~/.ssh/id_sftp_acme -o BatchMode=yes sftp_acme@sftp.entreprise.lan
```

- `-o BatchMode=yes` : échoue au lieu de demander un mot de passe (indispensable en cron).
- Code retour non nul si une commande échoue → testable avec `$?`.
- Pour des batchs complexes, préférez `lftp` (voir § 42) ou un script Python/paramiko.

---

## 42. lftp : miroir, reprise et robustesse

```bash
sudo apt install -y lftp
```

```bash
# Miroir montant (upload) avec reprise et suppression distante des fichiers disparus
lftp -u sftp_acme, sftp://sftp.entreprise.lan <<'EOF'
mirror --reverse --continue --delete --verbose /var/exports/ /depot/
bye
EOF
```

Options clés :

| Option | Effet |
|---|---|
| `--reverse` | Envoi local → distant (sans = téléchargement) |
| `--continue` | Reprise des fichiers partiels |
| `--delete` | Supprime côté distant ce qui n'existe plus en local |
| `--only-newer` | Ne transfère que les fichiers plus récents |
| `--parallel=4` | 4 transferts simultanés |
| `--exclude` | Exclure un motif |

Authentification par clé avec lftp :

```bash
# ~/.lftprc ou variables d'environnement
# lftp utilise l'agent SSH ou la clé par défaut ; sinon :
lftp -u sftp_acme sftp://sftp.entreprise.lan -e "set sftp:connect-program 'ssh -i /chemin/cle'; mirror ..."
```

---

## 43. Automatisation par cron (comptes techniques)

```bash
# /etc/cron.d/envoi_sftp — envoi quotidien à 2h05
SHELL=/bin/bash
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
5 2 * * * svc_export /usr/local/bin/envoi_quotidien.sh >>/var/log/envoi_sftp.log 2>&1
```

`/usr/local/bin/envoi_quotidien.sh` :

```bash
#!/bin/bash
set -euo pipefail
CLE=/etc/svc_export/id_ed25519
HOTE=sftp.partenaire.lan
USER=sftp_expediteur
SRC=/var/exports/a_envoyer
[ -f "$CLE" ] || { echo "clé absente"; exit 1; }
sftp -b /dev/stdin -i "$CLE" -o BatchMode=yes -o ConnectTimeout=15 "$USER@$HOTE" <<EOF
cd depot
mput $SRC/*.csv
bye
EOF
logger -t envoi_sftp "envoi quotidien terminé : $?"
```

