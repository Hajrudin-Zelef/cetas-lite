---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-11
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "distribution", "incident"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [2284, 2470]
sha256: af6766ae0e34b9b36e9478f9bad1578e19090cf8a2444de7b5be1a1692d594ee
---

# PostgreSQL en production — Guide complet pour sysadmin

- [ ] Sauvegarde complète **testée** avant
- [ ] `--check` sans erreur (extensions compatibles !)
- [ ] Extensions : `postgis`, `pg_stat_statements`… à jour pour la cible
- [ ] Fenêtre de maintenance + plan de retour arrière (l'ancien cluster
      reste intact avec `--link`… tant qu'on ne le supprime pas)
- [ ] `ANALYZE` sur le nouveau cluster (stats non migrées)
- [ ] Tests applicatifs non-régression avant bascule DNS/VIP

> ⚠️ Alternative sans `pg_upgrade` : réplication logique
> (section 56) 17 → 18 avec bascule applicative = coupure quasi nulle,
> au prix d'une préparation plus lourde.

---

## 70. Dépannage : méthode générale

1. **Qualifier** : panne totale ? lenteurs ? erreurs applicatives ?
   Depuis quand ? Qu'est-ce qui a changé (déploiement, batch, volume) ?
2. **Logs d'abord** : `/var/log/postgresql/`, `journalctl -u`.
3. **Métriques** : `pg_stat_activity`, espace disque, charge CPU/RAM.
4. **Hypothèse unique**, action unique, vérification.
5. **Documenter** : cause racine + remédiation dans le journal d'exploitation.

```bash
# Kit de survie : les 5 commandes à lancer en premier
pg_lsclusters                                          # le cluster tourne ?
df -h /var/lib/postgresql /var/log/postgresql          # disque plein ?
sudo -u postgres psql -c "SELECT count(*) FROM pg_stat_activity;"  # connexions ?
sudo -u postgres psql -c "SELECT * FROM pg_stat_archiver;"         # WAL archivé ?
tail -n 100 /var/log/postgresql/postgresql-17-main.log # erreurs récentes ?
```

---

## 71. Cas concret 1 : « FATAL: remaining connection slots are reserved »

**Symptômes :** les applis ne se connectent plus, `FATAL: too many clients`.

**Diagnostic :**

```sql
-- D'abord se connecter en superuser (slots réservés : superuser_reserved_connections)
SELECT state, count(*) FROM pg_stat_activity GROUP BY state;
-- Souvent : des centaines de 'idle' -> fuites de connexions applicatives
SELECT pid, usename, application_name, now() - state_change AS depuis,
       left(query, 60)
FROM pg_stat_activity WHERE state = 'idle' ORDER BY state_change LIMIT 10;
```

**Remédiation :**

1. Court terme : `SELECT pg_terminate_backend(pid)` sur les `idle`
   les plus anciens (après accord applicatif).
2. Moyen terme : `idle_in_transaction_session_timeout`,
   `statement_timeout` par rôle.
3. Fond : **PgBouncer** en mode transaction devant PG ; corriger le
   pool applicatif (fermeture des connexions).

---

## 72. Cas concret 2 : disque plein à cause du WAL

**Symptômes :** `PANIC: could not write to file "pg_wal/..."`, cluster arrêté.

**Diagnostic :**

```bash
du -sh /var/lib/postgresql/17/main/pg_wal
sudo -u postgres psql -c "SELECT * FROM pg_stat_archiver;"
# failed_count > 0 -> archive_command en panne (disque dest plein, droits...)
sudo -u postgres psql -c "SELECT slot_name, active, pg_size_pretty(pg_wal_lsn_diff(pg_current_wal_lsn(), restart_lsn)) AS retenu FROM pg_replication_slots;"
# slot inactif qui retient des Go de WAL -> standby mort non supprimé
```

**Remédiation :**

1. Réparer `archive_command` (espace/droits/réseau) → PG reprend l'archivage.
2. Slot mort : si le standby est abandonné,
   `SELECT pg_drop_replication_slot('standby_mort');` (⚠️ le standby
   devra être reconstruit).
3. Redémarrer, vérifier la reprise des checkpoints.

**Prévention :** alerte sur `pg_stat_archiver.failed_count` et sur la
taille de `pg_wal` (section 65).

---

## 73. Cas concret 3 : table énorme et requêtes soudain lentes

**Symptômes :** une requête habituellement rapide devient lente du jour
au lendemain, sans changement de code.

**Diagnostic :**

```sql
EXPLAIN (ANALYZE, BUFFERS) SELECT ... ;  -- la requête incriminée
-- Si rows estimé << rows réel :
SELECT last_analyze, last_autoanalyze, n_mod_since_analyze
FROM pg_stat_user_tables WHERE relname = 'mesures';
-- Si n_dead_tup énorme :
SELECT relname, n_live_tup, n_dead_tup FROM pg_stat_user_tables ORDER BY n_dead_tup DESC LIMIT 5;
```

**Causes fréquentes :**

1. Statistiques périmées après un gros chargement → `ANALYZE`.
2. Bloat : autovacuum dépassé → `VACUUM (VERBOSE, ANALYZE)` manuel puis
   réglage par table (section 62).
3. Plan qui a basculé (distribution changée) → statistiques ciblées
   `SET STATISTICS`, voire `pg_hint_plan` en dernier recours.

---

## 74. Cas concret 4 : « deadlock detected »

```sql
-- Le log indique les deux requêtes et les verrous :
-- ERROR: deadlock detected
-- DETAIL: Process 1234 waits for ShareLock on transaction 5678; blocked by process 4321.
```

**Cause :** deux transactions verrouillent les mêmes lignes dans un
ordre différent. **Remédiation :**

1. Toujours verrouiller **dans le même ordre** (ex. `ORDER BY id`
   avant `SELECT ... FOR UPDATE`).
2. Transactions courtes (section 45).
3. Côté appli : **retry** sur `deadlock_detected` (erreur 40P01) —
   c'est un incident normal et récupérable, pas un bug PG.

---

## 75. 10+ erreurs classiques et leur solution

| # | Erreur / symptôme | Cause probable | Solution |
|---|---|---|---|
| 1 | `FATAL: password authentication failed` | Mauvais mdp / pg_hba `md5` vs `scram` | Vérifier `pg_hba.conf` + `password_encryption` (sec. 16–17) |
| 2 | `FATAL: no pg_hba.conf entry for host...` | Règle manquante | Ajouter la ligne `host`, `reload` |
| 3 | `FATAL: remaining connection slots...` | Saturation connexions | Sec. 71 : pooler, timeouts |
| 4 | `ERROR: permission denied for table X` | `GRANT` manquant ou `DEFAULT PRIVILEGES` oubliés | Sec. 19 |
| 5 | `ERROR: relation "x" does not exist` | `search_path` / schéma / casse (`"MaTable"`) | Qualifier `schema.table`, vérifier `\dt` |
| 6 | `canceling statement due to conflict with recovery` | Standby en retard vs requêtes longues | `max_standby_*_delay`↑ ou `hot_standby_feedback` (sec. 53) |
| 7 | `PANIC: could not write to file pg_wal` | Disque plein / archivage HS | Sec. 72 |
| 8 | `WARNING: oldest xmin is far in the past` | Longue transaction bloque le vacuum | Tuer la transaction, `idle_in_transaction_session_timeout` |
| 9 | `ERROR: deadlock detected` | Ordre de verrouillage incohérent | Sec. 74 |
| 10 | `FATAL: the database system is in recovery mode` | Standby en lecture seule / recovery en cours | Normal sur standby ; sinon vérifier les logs |
| 11 | `ERROR: duplicate key value violates unique constraint` | UPSERT manquant / race applicative | `ON CONFLICT DO NOTHING/UPDATE` (sec. 28) |
| 12 | `could not serialize access due to concurrent update` | `SERIALIZABLE`/`REPEATABLE READ` + concurrence | Retry applicatif |
| 13 | Montée de version : `pg_upgrade --check` échoue | Extension incompatible | MAJ l'extension d'abord (sec. 69) |

---

## 76. Corruption : prévention et conduite à tenir

**Prévention :**

- Stockage fiable (pas de NFS pour PGDATA), `full_page_writes = on`
  (ne jamais couper), onduleur + arrêt propre (lien avec votre
  expertise UPS : un arrêt brutal répété fragilise tout).
- `data_checksums` activé à `initdb` (`--data-checksums`) : détecte
  les corruptions silencieuses disque.

```bash
# Vérifier si les checksums sont actifs
sudo -u postgres psql -tAc "SHOW data_checksums;"
```

**Conduite à tenir si corruption suspectée :**

1. **Stopper les écritures**, sauvegarder l'état (copie physique du
   PGDATA, même corrompu).
2. `pg_dump` ce qui est lisible ; restaurer sur un cluster neuf.
3. Restaurer depuis la sauvegarde physique + PITR **avant** la corruption
   si le dump échoue.
4. Ne jamais `VACUUM FULL` / `REINDEX` à l'aveugle sur un cluster
   corrompu : ça peut aggraver.

> ⚠️ La corruption est rare avec du matériel sain. 90 % des
> « corruptions » remontées sont en réalité du disque plein, des
> droits fichiers cassés, ou deux postmasters sur le même PGDATA.

---

# Partie H — Cas pratiques, pense-bête, glossaire, quiz

## 77. Cas pratique 1 : dimensionner PostgreSQL pour une appli

