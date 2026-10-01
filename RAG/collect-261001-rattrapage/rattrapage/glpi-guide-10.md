---
id: collect-261001-rattrapage/rattrapage/glpi-guide-10
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-20"]
keywords: ["apache", "attention", "incident"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [1509, 1739]
sha256: 5e41622f9f071417bd678dac686e8ec9e13e242ff7516edb7c3b55019348485f
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

- [ ] Le profil déclarant ne voit **que** sa propre entité (pas de récursivité).
- [ ] Les techniciens ne peuvent pas **supprimer** (seulement mettre au rebut).
- [ ] Seuls Admin/Superviseur gèrent **utilisateurs, profils, entités, règles**.
- [ ] L'accès à `Configuration` est réservé à l'Admin.
- [ ] Les **fournisseurs externes** (s'ils ont un accès) : profil dédié, entité
      limitée, pas d'accès aux budgets ni aux autres clients.
- [ ] Désactivez les **comptes par défaut** inutilisés (`tech`, `normal`,
      `post-only` si vous avez créé vos propres profils).

### Audit

`Administration > Journaux` : historique de **toutes** les modifications
(qui, quoi, quand). En cas d'incident, c'est votre piste d'audit.
Conservez les journaux ≥ 1 an (purge paramétrable, attention aux obligations
contractuelles/RGPD).

---

## 44. Sécurité — HTTPS/TLS de bout en bout

GLPI **ne doit jamais** circuler en HTTP clair en production (mots de passe,
données clients).

### Certificat

- En interne : certificat d'une **CA interne** (ou Let's Encrypt si exposé).
- Déposez cert + clé : `/etc/ssl/certs/glpi.crt`, `/etc/ssl/private/glpi.key`
  (permissions `640`, propriétaire `root:ssl-cert`, www-data dans le groupe).

### Redirection HTTP → HTTPS (Apache)

```apache
<VirtualHost *:80>
    ServerName glpi.entreprise.lan
    Redirect permanent / https://glpi.entreprise.lan/
</VirtualHost>
```

### En-têtes de sécurité (Apache, vhost 443)

```apache
Header always set Strict-Transport-Security "max-age=31536000; includeSubDomains"
Header always set X-Content-Type-Options "nosniff"
Header always set X-Frame-Options "SAMEORIGIN"
Header always set Referrer-Policy "strict-origin-when-cross-origin"
Header always set Content-Security-Policy "default-src 'self' https:; img-src 'self' data:;"
```

### Côté GLPI

- `Configuration > Générale` : URL de l'application en `https://...`.
- `php.ini` : `session.cookie_secure = 1`, `session.cookie_httponly = 1`,
  `session.cookie_samesite = "Lax"`.
- Testez avec un scanner (testssl.sh en interne) après chaque changement.

---

## 45. Sécurité — durcissement serveur et GLPI

### Système

- Mises à jour automatiques de sécurité (`unattended-upgrades` sur Debian/Ubuntu).
- Pare-feu : 443 ouvert vers le LAN, SSH restreint, **pas d'exposition directe
  sur Internet** sans WAF/VPN (préférez un VPN pour l'accès distant des techniciens).
- Sauvegardes chiffrées et **testées** (partie IX).

### GLPI applicatif

- Dossiers `config/` et `files/` **hors racine web** (section 8.4).
- `install/` supprimé (section 8.1).
- Droits : fichiers `644`, dossiers `755`, propriétaire `www-data` — jamais de `777`.
- ` marketplace` : n'installez que des plugins **compatibles avec votre version
  exacte** et maintenus (section 51).
- Désactivez l'affichage des erreurs PHP en production :
  `display_errors = Off`, `log_errors = On` (les stack traces en page blanche
  révèlent des chemins et versions — voir section 54).

### Fichiers sensibles à surveiller

| Fichier | Risque si exposé |
|---|---|
| `config/config_db.php` | Identifiants MariaDB |
| `files/_log/*.log` | Chemins, erreurs, parfois données |
| Dumps de sauvegarde dans `files/` | Base complète ! → stocker hors web |

---

# PARTIE IX — SUPERVISION ET SAUVEGARDE

## 46. Supervision de GLPI : ce qu'il faut surveiller

GLPI supervise le parc ; **qui supervise GLPI ?** Votre Zabbix/Nagios/Prometheus.

### Points de contrôle

| Sonde | Seuil d'alerte | Comment |
|---|---|---|
| HTTPS `https://glpi.entreprise.lan` | code ≠ 200 | check_http |
| Cron GLPI (dernière exécution) | > 10 min | script section 9.4 |
| File de notifications `queuednotification` | > 50 en attente depuis > 30 min | requête SQL |
| Espace disque `/var/lib/glpi` | > 80 % | les pièces jointes grossissent vite |
| MariaDB joignable + réplication (si) | — | check_mysql |
| Validité certificat TLS | < 30 jours | check_http --ssl |
| Temps de réponse page d'accueil | > 5 s | indicateur de saturation |

### Requêtes SQL utiles pour la supervision

```sql
-- Notifications en file d'attente (anormal si ancien)
SELECT COUNT(*) AS en_attente,
       TIMESTAMPDIFF(MINUTE, MIN(date_creation), NOW()) AS plus_ancien_min
FROM glpidb.glpi_queuednotifications;

-- Tâches cron en erreur
SELECT name, lastrun, lastcode FROM glpidb.glpi_crontasks
WHERE state != 1 OR lastcode != 0;

-- Tickets critiques ouverts sans assigné
SELECT COUNT(*) FROM glpidb.glpi_tickets
WHERE status NOT IN (5,6) AND priority >= 4
  AND (users_id_recipient = 0);
```

> Les statuts GLPI : 1=Nouveau, 2=En cours (attribué), 3=En cours (planifié),
> 4=En attente, 5=Résolu, 6=Clos. Vérifiez dans votre version
> (`glpi_configs` / interface).

---

## 47. Sauvegarde : stratégie complète (fichiers + base)

**Règle 3-2-1** : 3 copies, 2 supports différents, 1 hors site.

### Ce qu'il faut sauvegarder

| Quoi | Où | Fréquence |
|---|---|---|
| Base `glpidb` (dump) | `mysqldump` | **Quotidienne** |
| `/var/lib/glpi/files` (pièces jointes, documents) | rsync/tar | Quotidienne (incrémentale) |
| `/var/lib/glpi/config` + vhost Apache/nginx | tar | À chaque changement |
| Code `/var/www/glpi` | tar versionné | À chaque mise à jour |

### Rétention type

- Quotidiennes : 14 jours en local.
- Hebdomadaires : 8 semaines sur NAS.
- Mensuelles : 12 mois hors site (chiffrées).

> **RGPD :** les sauvegardes contiennent des données personnelles (noms,
> courriels). Chiffrez-les (`gpg`/`openssl`) et limitez l'accès au support
> de sauvegarde.

---

## 48. Sauvegarde automatisée : scripts prêts à l'emploi

### Script `/usr/local/bin/backup-glpi.sh`

```bash
#!/bin/bash
set -euo pipefail
DEST="/srv/backup/glpi"
DATE=$(date +%F)
RETENTION_LOCAL=14
DB="glpidb"; DBUSER="glpi"; DBPASS="REMPLACER"  # ou fichier ~/.my.cnf 600 !

mkdir -p "$DEST"
# 1. Dump base
mysqldump --single-transaction --routines --triggers \
  -u"$DBUSER" -p"$DBPASS" "$DB" | gzip > "$DEST/glpidb-$DATE.sql.gz"
# 2. Fichiers applicatifs
tar czf "$DEST/glpi-files-$DATE.tar.gz" -C / var/lib/glpi
# 3. Config web + code (léger)
tar czf "$DEST/glpi-conf-$DATE.tar.gz" /etc/apache2/sites-available/glpi.conf
# 4. Rotation locale
find "$DEST" -name 'glpidb-*.sql.gz' -mtime +$RETENTION_LOCAL -delete
find "$DEST" -name 'glpi-files-*.tar.gz' -mtime +$RETENTION_LOCAL -delete
# 5. Copie vers NAS (exemple)
rsync -a --delete "$DEST/" nas:/backups/glpi/ || echo "WARN: rsync NAS échoué"
echo "OK backup GLPI $DATE"
```

```bash
sudo chmod 700 /usr/local/bin/backup-glpi.sh
# cron root : tous les jours à 2h
echo "0 2 * * * root /usr/local/bin/backup-glpi.sh >> /var/log/backup-glpi.log 2>&1" \
  | sudo tee /etc/cron.d/backup-glpi
```

### Variante sans mot de passe en clair

`/root/.my.cnf` (permissions `600`) :

```ini
[client]
user = glpi
password = MOT_DE_PASSE_FORT
host = localhost
```

puis `mysqldump --defaults-extra-file=/root/.my.cnf ...` (plus sûr que `-p`).

---

##  49. Restauration : procédure testée pas à pas

**Une sauvegarde non testée = pas de sauvegarde.** Testez la restauration
**trimestriellement** sur une machine isolée.

### Procédure

```bash
# 1. Stopper le cron GLPI et le web (éviter les écritures)
sudo systemctl stop apache2   # ou nginx

# 2. Restaurer la base
zcat /srv/backup/glpi/glpidb-2026-09-20.sql.gz \
  | mysql -u root -p glpidb_restauree   # d'abord sur une base de test !

# 3. Restaurer les fichiers
sudo tar xzf /srv/backup/glpi/glpi-files-2026-09-20.tar.gz -C /
sudo chown -R www-data:www-data /var/lib/glpi

# 4. Redémarrer et vérifier
sudo systemctl start apache2
# → page de connexion OK, connexion admin OK, ticket récent présent,
#    pièce jointe téléchargeable, cron relancé
```

### Checklist de validation post-restauration

