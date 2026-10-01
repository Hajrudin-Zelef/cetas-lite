---
id: collect-261001-rattrapage/rattrapage/netbox-guide-11
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [2110, 2308]
sha256: 9d8086778a5a02509192f5ffc7f69f45806ab3f4bf2b3df7b804f5b54586fb62
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

```bash
chmod 700 /opt/netbox/backup-netbox-db.sh
# Cron quotidien à 2h du matin
echo "0 2 * * * root /opt/netbox/backup-netbox-db.sh >> /var/log/netbox/backup.log 2>&1" \
  > /etc/cron.d/netbox-backup
```

Vérifications automatiques à ajouter : taille du dump > 0, test `pg_restore --list`
sur le dernier dump (intégrité), copie vers un stockage **externe** (NAS, S3, autre site).

> ⚠️ Un dump sur le même disque que la base ne protège que des erreurs logiques,
> pas d'une panne disque. La règle 3-2-1 s'applique : 3 copies, 2 supports, 1 hors site.

---

## 65. Sauvegarde : fichiers `media/` et `configuration.py`

```bash
#!/bin/bash
# /opt/netbox/backup-netbox-files.sh  (700)
set -euo pipefail
BACKUP_DIR="/var/backups/netbox"
DATE=$(date +%Y%m%d-%H%M%S)

# media/ : pièces jointes, images, scripts et rapports personnalisés
tar -czf "$BACKUP_DIR/netbox-media-$DATE.tar.gz" \
    -C /opt/netbox/netbox/netbox media

# configuration sensible (chiffrée au repos si possible)
tar -czf "$BACKUP_DIR/netbox-config-$DATE.tar.gz" \
    -C /opt/netbox/netbox/netbox configuration.py \
    -C /opt/netbox gunicorn.py

find "$BACKUP_DIR" -name 'netbox-media-*' -mtime +30 -delete
find "$BACKUP_DIR" -name 'netbox-config-*' -mtime +30 -delete
echo "OK $DATE"
```

Plan de sauvegarde complet (à documenter dans votre PRA) :

| Élément | Méthode | Fréquence | Rétention |
|---|---|---|---|
| Base PostgreSQL | `pg_dump -Fc` + `.sql.gz` | quotidienne 2h | 30 jours |
| `media/` | tar.gz | quotidienne 2h05 | 30 jours |
| `configuration.py`, `gunicorn.py` | tar.gz | à chaque changement + hebdo | 90 jours |
| Copie hors site | rsync vers NAS / S3 | quotidienne 3h | selon politique |
| Snapshot VM | hyperviseur | hebdo | 4 semaines |

---

## 66. Restauration : procédure complète (à tester chaque trimestre)

> ⚠️ Testez la restauration sur une VM de LAB, jamais en production la première fois.
> Chronométrez-la : c'est votre RTO réel.

```bash
# 1. Arrêter les services (évite toute écriture pendant la restauration)
systemctl stop netbox netbox-rq

# 2. Recréer une base vierge
sudo -u postgres psql -c "DROP DATABASE netbox;"
sudo -u postgres psql -c "CREATE DATABASE netbox OWNER netbox;"

# 3. Restaurer le dump (format custom)
sudo -u postgres pg_restore -d netbox -U netbox /var/backups/netbox/netbox-20260926-020000.dump

# 4. Restaurer les fichiers media/
tar -xzf /var/backups/netbox/netbox-media-20260926-020000.tar.gz \
    -C /opt/netbox/netbox/netbox/
chown -R netbox:netbox /opt/netbox/netbox/netbox/media

# 5. Redémarrer et vérifier
systemctl start netbox netbox-rq
sleep 5
curl -sk -o /dev/null -w "%{http_code}\n" https://localhost/   # 200 ou 302 attendu
sudo -iu netbox /opt/netbox/venv/bin/python \
  /opt/netbox/netbox/netbox/manage.py check
```

Checklist post-restauration : connexion admin OK, un site/une baie/une IP vérifiés
au hasard, un appel API avec un token existant, les files RQ vides et actives.

---

## 67. Supervision de NetBox lui-même

NetBox est un outil critique : s'il est en panne, l'équipe travaille à l'aveugle.
Supervisez-le comme une application de production :

| Sonde | Méthode | Seuil d'alerte |
|---|---|---|
| HTTPS | check HTTP sur `/login/` → 200 | critique si KO |
| Certificat TLS | expiration < 21 jours | warning, < 7 j critique |
| Processus gunicorn | `systemctl is-active netbox` | critique si inactif |
| File RQ | `llen rq:queue:default` | warning si > 50 |
| Espace disque | `/var/lib/postgresql`, `/var/backups` | warning 80 %, critique 90 % |
| Fraîcheur sauvegarde | âge du dernier dump | critique si > 26 h |
| PostgreSQL | connexion + `SELECT 1` | critique si KO |

Exemple avec Zabbix (agent) — `UserParameter` :

```bash
# /etc/zabbix/zabbix_agentd.d/netbox.conf
UserParameter=netbox.https, curl -sk -o /dev/null -w "%{http_code}" https://localhost/login/
UserParameter=netbox.rq.default, redis-cli -n 0 llen rq:queue:default
UserParameter=netbox.backup.age, echo $(( $(date +%s) - $(stat -c %Y /var/backups/netbox/netbox-*.dump | sort -n | tail -1) ))
```

> 💡 Ajoutez NetBox lui-même dans NetBox (section 28) avec le tag `supervise-zabbix` :
> la boucle est bouclée — l'outil qui documente la supervision est supervisé.

---

## 68. Montée de version de NetBox (4.x → 4.y)

> ⚠️ Ne sautez jamais de version majeure sans lire les *release notes* : chaque
> version 4.x peut exiger une version Python/PostgreSQL/Redis minimale différente
> et des migrations spécifiques.

Procédure type (exemple 4.2.3 → 4.3.0) :

```bash
# 0. SAUVEGARDE COMPLÈTE + snapshot VM (obligatoire)
# 1. Télécharger la nouvelle version (utilisateur netbox)
sudo -iu netbox
cd /opt/netbox
wget https://github.com/netbox-community/netbox/archive/refs/tags/v4.3.0.tar.gz -O netbox-4.3.0.tar.gz
tar -xzf netbox-4.3.0.tar.gz

# 2. Arrêter les services
systemctl stop netbox netbox-rq

# 3. Nouveau venv ou mise à jour des dépendances ? -> suivre les release notes.
#    En général : réutiliser le venv et mettre à jour requirements :
/opt/netbox/venv/bin/pip install -r /opt/netbox/netbox-4.3.0/requirements.txt

# 4. Recopier la configuration et les personnalisations
cp /opt/netbox/netbox/netbox/configuration.py /opt/netbox/netbox-4.3.0/netbox/netbox/
# + scripts/, reports/, media/ si personnalisés :
cp -r /opt/netbox/netbox/netbox/scripts/* /opt/netbox/netbox-4.3.0/netbox/netbox/scripts/ 2>/dev/null || true
cp -r /opt/netbox/netbox/netbox/reports/* /opt/netbox/netbox-4.3.0/netbox/netbox/reports/ 2>/dev/null || true

# 5. Basculer le lien symbolique
ln -sfn /opt/netbox/netbox-4.3.0 /opt/netbox/netbox

# 6. Migrations + statiques (dans le venv)
cd /opt/netbox/netbox
/opt/netbox/venv/bin/python netbox/manage.py migrate
/opt/netbox/venv/bin/python netbox/manage.py collectstatic --no-input

# 7. Redémarrer et vérifier
systemctl start netbox netbox-rq
/opt/netbox/venv/bin/python netbox/manage.py check
curl -sk -o /dev/null -w "%{http_code}\n" https://localhost/login/
```

En cas d'échec : rebasculez le lien symbolique vers l'ancienne version, restaurez
le dump pré-upgrade, redémarrez. D'où l'importance de l'étape 0.

---

## 69. Dépannage : méthode générale

Avant de chercher l'erreur précise, appliquez toujours ce triptyque :

1. **Les services tournent-ils ?**
   `systemctl status netbox netbox-rq postgresql redis-server nginx --no-pager`
2. **Que disent les logs ?** (dans cet ordre)
   - `/var/log/netbox/gunicorn-error.log` (erreurs applicatives),
   - `journalctl -u netbox -n 50` (service systemd),
   - `/var/log/nginx/error.log` (frontal),
   - `journalctl -u postgresql -n 30` (base).
3. **Qu'est-ce qui a changé ?** (mise à jour, `configuration.py` modifié, disque plein,
   certificat expiré — 80 % des pannes viennent d'un changement récent).

```bash
# Kit de survie en une commande : état global
echo "=== SERVICES ==="; systemctl is-active netbox netbox-rq postgresql redis-server nginx
echo "=== ECOUTE ==="; ss -ltnp | grep -E '8001|443'
echo "=== DISQUE ==="; df -h / /var/lib/postgresql /var/backups | awk '{print $1,$5,$6}'
echo "=== DERNIER DUMP ==="; ls -lt /var/backups/netbox/*.dump 2>/dev/null | head -1
```

Les sections 70 à 79 détaillent les 10 erreurs classiques rencontrées en production.

---

## 70. Erreur n°1 : « Bad Request (400) » juste après l'installation

**Symptômes** : nginx répond, mais NetBox affiche « Bad Request (400) ».

**Causes possibles** (par ordre de probabilité) :

1. Le nom d'hôte utilisé n'est pas dans `ALLOWED_HOSTS` (section 10). Vous accédez
   via l'IP `10.10.0.15` mais seuls `netbox.lan-entreprise.fr` sont autorisés.
2. En-tête `Host` non transmis par nginx (`proxy_set_header Host $host;` manquant).

**Diagnostic / remède** :

