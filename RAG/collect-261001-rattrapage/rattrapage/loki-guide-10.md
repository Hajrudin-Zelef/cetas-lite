---
id: collect-261001-rattrapage/rattrapage/loki-guide-10
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-25"]
keywords: ["agent", "agents", "arr", "attention", "memory"]
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [2385, 2612]
sha256: ab371aff657682505188b67a563738c788a8d0735f102cdb0913731ba548dfda
---

# Grafana Loki — Le guide complet

- [ ] `auth_enabled: true` + tenants (section 55)
- [ ] TLS partout (section 56) ou réseau chiffré (WireGuard)
- [ ] Reverse proxy avec authentification (section 57)
- [ ] Pare-feu : 3100/9080 accessibles uniquement aux sources légitimes
      (`ufw`, nftables)
- [ ] Service systemd durci : `NoNewPrivileges`, `ProtectSystem=strict`
      (section 11)
- [ ] Pas d'interface d'admin exposée : le port 3100 sert aussi les
      endpoints `/config`, `/runtime` — filtrez via le proxy si besoin
- [ ] Mises à jour suivies (section 64) — les CVE sur les dépendances Go
      sont régulières
- [ ] Sauvegarde chiffrée des chunks si données sensibles (section 62)
- [ ] Journaliser les accès au proxy (qui a requêté quoi, quand)

---

## 59. Supervision de Loki : ses propres métriques

Loki expose ses métriques au format Prometheus sur `/metrics`
(port 3100). Ajoutez-le comme cible Prometheus :

```yaml
# prometheus.yaml
scrape_configs:
  - job_name: loki
    static_configs:
      - targets: ['loki.local:3100']
  - job_name: promtail
    static_configs:
      - targets: ['srv01:9080', 'srv02:9080']   # un par serveur
```

Métriques clés à surveiller :

| Métrique | Signification | Seuil d'alerte |
|---|---|---|
| `loki_ingester_memory_streams` | flux actifs en mémoire | croissance anormale = cardinalité |
| `loki_distributor_lines_received_total` | lignes reçues/s | chute = problème d'ingestion |
| `rate(loki_request_duration_seconds_count[5m])` | requêtes/s par route | — |
| `loki_ingester_chunk_age_seconds` | âge des chunks | — |
| `loki_boltdb_shipper_compactor_running` | compactor actif | doit être > 0 |
| `promtail_targets_active` | fichiers suivis par Promtail | chute = agent tombé |
| `promtail_file_bytes_total` | octets lus | stagnation = fichier bloqué |
| `process_resident_memory_bytes{job="loki"}` | RAM de Loki | vs limite du serveur |

---

## 60. Alertes sur la santé de Loki

Règles Prometheus (pas Loki !) pour superviser Loki lui-même :

```yaml
# /etc/prometheus/rules/loki-health.yaml
groups:
  - name: loki-sante
    interval: 1m
    rules:
      - alert: LokiIngesterMemoireElevee
        expr: process_resident_memory_bytes{job="loki"} > 8 * 1024 * 1024 * 1024
        for: 10m
        labels: {severity: warning}
        annotations:
          summary: "Loki consomme plus de 8 Go de RAM"

      - alert: LokiIngestionEnChute
        expr: rate(loki_distributor_lines_received_total[5m]) == 0
        for: 10m
        labels: {severity: critical}
        annotations:
          summary: "Aucune ligne reçue depuis 10 min — ingestion en panne ?"

      - alert: PromtailAgentTombe
        expr: up{job="promtail"} == 0
        for: 5m
        labels: {severity: warning}
        annotations:
          summary: "Promtail tombé sur {{ $labels.instance }}"

      - alert: LokiExplosionCardinalite
        expr: |
          deriv(loki_ingester_memory_streams[15m]) > 100
        for: 15m
        labels: {severity: critical}
        annotations:
          summary: "Les flux actifs explosent — vérifier les labels"
```

---

## 61. Sauvegarde : chunks sur filesystem

Avec le stockage **filesystem** (petit déploiement), la sauvegarde =
copie des répertoires de données **à froid** (Loki arrêté) ou via
snapshot filesystem (LVM/btrfs/ZFS) à chaud.

```bash
# Méthode simple et sûre : arrêt + rsync + redémarrage
sudo systemctl stop loki
sudo rsync -a --delete /var/lib/loki/ /mnt/sauvegarde/loki-$(date +%F)/
sudo systemctl start loki

# Méthode à chaud avec snapshot LVM (exemple) :
sudo lvcreate --size 10G --snapshot --name loki-snap /dev/vg0/lv-loki
sudo mount -o ro /dev/vg0/loki-snap /mnt/loki-snap
sudo rsync -a /mnt/loki-snap/ /mnt/sauvegarde/loki-$(date +%F)/
sudo umount /mnt/loki-snap
sudo lvremove -y /dev/vg0/loki-snap
```

Quoi sauvegarder impérativement :

| Répertoire | Contenu |
|---|---|
| `/var/lib/loki/chunks` | les logs eux-mêmes |
| `/var/lib/loki/tsdb-index` | l'index actif |
| `/var/lib/loki/tsdb-cache` | cache (optionnel) |
| `/var/lib/loki/wal` | chunks non flushés (critique !) |
| `/var/lib/loki/rules` | règles d'alerting |
| `/etc/loki/` | configuration |

> ⚠️ Sans le **WAL**, les logs des ~30 dernières minutes (chunks en
> mémoire) sont perdus en cas de crash. C'est la pièce la plus critique
> avec les chunks.

---

## 62. Sauvegarde : stockage objet (S3)

Avec S3/MinIO, la sauvegarde = **réplication du bucket** (versioning +
réplication inter-région ou vers un second MinIO).

Stratégie recommandée :

1. **Versioning** activé sur le bucket `loki-chunks` (protection contre
   les suppressions accidentelles du compactor).
2. **Réplication** vers un bucket secondaire (autre site / autre MinIO).
3. **Lifecycle** : transition vers stockage froid après 90 jours si le
   provider le permet (les vieux chunks sont rarement relus).
4. Sauvegarder **à part** : `/etc/loki/`, les règles, et le WAL local
   (lui n'est pas sur S3 !).

```bash
# Exemple MinIO : réplication avec mc (ponctuelle, hors versioning natif)
mc mirror --overwrite --remove \
  minio-principal/loki-chunks \
  minio-secours/loki-chunks
```

> 💡 Avec le stockage objet, Loki lui-même est **stateless** (hors WAL) :
> on peut reconstruire un Loki neuf qui relit le bucket. Testez cette
> reconstruction une fois par an (voir section 63).

---

## 63. Restauration après sinistre

### Scénario : disque du Loki filesystem mort, sauvegarde rsync disponible

```bash
# 1. Réinstaller Loki (section 10-11)
# 2. Restaurer les données AVANT le premier démarrage
sudo systemctl stop loki
sudo rsync -a /mnt/sauvegarde/loki-2026-09-25/ /var/lib/loki/
sudo chown -R loki:loki /var/lib/loki
# 3. Restaurer la config
sudo cp /mnt/sauvegarde/loki-etc/loki.yaml /etc/loki/
# 4. Démarrer et vérifier
sudo systemctl start loki
curl -s http://localhost:3100/ready
# 5. Tester une requête sur une vieille période
curl -s 'http://localhost:3100/loki/api/v1/query_range?query={job="syslog"}&start=2026-08-01T00:00:00Z&end=2026-08-02T00:00:00Z' | head -c 500
```

### Scénario : stockage objet intact, Loki à reconstruire

```bash
# 1. Nouveau serveur, même loki.yaml (mêmes buckets, mêmes préfixes d'index)
# 2. Démarrer : Loki relit l'index TSDB depuis l'objet
# 3. Le WAL est perdu → trou de ~30 min de logs récents : acceptable,
#    les agents Promtail ne renvoient PAS l'historique (positions déjà avancées)
```

> ⚠️ Point honnête : après un sinistre total **sans WAL**, les logs en
> mémoire au moment du crash sont définitivement perdus. Pour les logs
> critiques (audit, sécurité), envisagez un double push Promtail vers
> deux Loki indépendants (`clients:` avec 2 URLs).

---

## 64. Mise à jour de Loki

Procédure sans surprise (binaire systemd) :

```bash
# 1. Lire les notes de version ! (breaking changes, migrations de schéma)
#    https://github.com/grafana/loki/releases

# 2. Sauvegarde avant (section 61)
sudo systemctl stop loki
sudo rsync -a /var/lib/loki/ /mnt/sauvegarde/loki-avant-maj/

# 3. Remplacer le binaire
cd /tmp
LOKI_VERSION="3.4.1"   # exemple
curl -sSLO "https://github.com/grafana/loki/releases/download/v${LOKI_VERSION}/loki-linux-amd64.zip"
unzip -o loki-linux-amd64.zip
sudo install -m 0755 loki-linux-amd64 /usr/local/bin/loki

# 4. Vérifier la config (dépréciations)
loki -config.file=/etc/loki/loki.yaml -verify-config 2>&1 | head -30
# ou : loki --print-config-stderr 2>/dev/null | grep -i deprecat

# 5. Redémarrer et surveiller
sudo systemctl start loki
sudo journalctl -u loki -f
curl -s http://localhost:3100/ready
```

Points d'attention par version :

| Montée | Vigilance |
|---|---|
| 2.9 → 3.x | schéma v13 recommandé, TSDB par défaut, flags renommés |
| Changement de schéma | ajouter une **nouvelle entrée** `schema_config` avec `from:` futur (section 66), jamais modifier l'existante |
| Ruler | vérifier que les règles s'évaluent toujours (`/ruler` API) |

---

