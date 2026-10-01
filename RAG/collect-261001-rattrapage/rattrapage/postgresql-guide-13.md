---
id: collect-261001-rattrapage/rattrapage/postgresql-guide-13
title: "PostgreSQL en production — Guide complet pour sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention", "incident"]
source: docs/RAG/collect-261001-rattrapage/postgresql_guide.md
source_anchor: ""
source_lines: [2684, 2829]
sha256: bcf9ab2179982e54fec340f04e18c73e96429cd336754f904a0b344daba4131b
---

# PostgreSQL en production — Guide complet pour sysadmin

| Terme | Définition |
|---|---|
| ACID | Atomicité, Cohérence, Isolation, Durabilité : garanties transactionnelles |
| Autovacuum | Processus automatique de nettoyage (VACUUM) et stats (ANALYZE) |
| Bloat | Gonflement d'une table/index par les versions mortes non nettoyées |
| Checkpoint | Point où toutes les pages sales sont écrites sur disque |
| Cluster | Instance PostgreSQL : un postmaster + ses bases + sa conf |
| GUC | Grand Unified Configuration : un paramètre de `postgresql.conf` |
| HOT | Heap-Only Tuple : UPDATE sans nouvelle entrée d'index (optimisation) |
| Lag | Retard d'un standby par rapport au primaire |
| LSN | Log Sequence Number : position dans le WAL |
| MVCC | Contrôle de concurrence multi-version (section 4) |
| PITR | Restauration à un instant T via WAL archivé |
| RLS | Row-Level Security : filtrage des lignes par rôle |
| RPO / RTO | Perte de données max / durée d'indisponibilité max acceptées |
| SCRAM-SHA-256 | Méthode d'authentification par mot de passe moderne |
| Slot (réplication) | Réserve le WAL pour un standby/outil (attention aux slots morts) |
| Split-brain | Deux primaires simultanés : corruption logique assurée |
| Standby | Réplica en lecture seule (streaming) |
| Timeline | Historique d'un cluster après chaque promotion (PITR) |
| TOAST | Stockage externe des grosses valeurs (text/jsonb volumineux) |
| VACUUM | Recyclage des versions mortes + gel des XID |
| WAL | Write-Ahead Log : journal des modifications, base de la réplication et du PITR |
| Wraparound | Dépassement du compteur de transactions : incident critique si l'âge XID explose |
| XID | Identifiant de transaction (surveillance : `age(relfrozenxid)`) |

---

## 83. Quiz : 10 questions pour valider

**Q1.** Pourquoi ne faut-il jamais couper `autovacuum = off` en production ?
**Q2.** Quelle différence entre `\copy` et `COPY` ?
**Q3.** Un `GRANT SELECT ON ALL TABLES IN SCHEMA public` suffit-il pour que
l'appli lise les tables créées demain ? Sinon, que manque-t-il ?
**Q4.** Votre `pg_wal` fait 10× `max_wal_size` : quelles sont les deux
causes les plus probables, et dans quel ordre vérifiez-vous ?
**Q5.** Pourquoi `CREATE INDEX CONCURRENTLY` est-il obligatoire en
production, et quelle est sa contrainte d'usage ?
**Q6.** Un standby affiche `canceling statement due to conflict with
recovery` : expliquez et donnez deux remèdes.
**Q7.** Quelle est la différence entre une mise à jour mineure (17.4→17.5)
et majeure (17→18) en termes de procédure ?
**Q8.** Pourquoi `timestamptz` plutôt que `timestamp` en production ?
**Q9.** Citez trois paramètres `postgresql.conf` qui exigent un
redémarrage (`context = 'postmaster'`).
**Q10.** Après un `DELETE` accidentel à 14h32, détaillez les grandes étapes
d'une restauration PITR sans écraser la production.

<details>
<summary><b>Réponses</b></summary>

**R1.** Sans autovacuum, les versions mortes du MVCC s'accumulent (bloat),
les statistiques se périment (mauvais plans) et l'âge des XID augmente
jusqu'au **wraparound** (arrêt forcé du cluster). C'est le poumon du moteur.
**R2.** `\copy` s'exécute côté **client** (fichiers de votre poste,
droits de votre user) ; `COPY` s'exécute côté **serveur** (fichiers du
serveur, réservé au superuser / `pg_read_server_files`).
**R3.** Non : `ON ALL TABLES` ne couvre que les tables existantes. Il faut
`ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO ...`
(et le refaire par rôle créateur si besoin).
**R4.** (1) `archive_command` en échec (`pg_stat_archiver.failed_count` > 0),
(2) slot de réplication inactif qui retient le WAL
(`pg_replication_slots`). Vérifier l'archiver d'abord, puis les slots.
**R5.** `CREATE INDEX` simple prend un verrou `SHARE`/`ACCESS EXCLUSIVE`
qui bloque les écritures ; `CONCURRENTLY` ne bloque pas. Contraintes :
2 passages (plus long), **interdit dans un bloc transactionnel**,
ne peut pas être combiné avec d'autres DDL dans la même transaction.
**R6.** Le standby applique le WAL et doit annuler une requête qui
détiendrait un snapshot trop vieux (`max_standby_streaming_delay`
dépassé). Remèdes : augmenter `max_standby_streaming_delay`, ou
`hot_standby_feedback = on` (au prix d'un possible bloat sur le primaire).
**R7.** Mineure : `apt upgrade` + restart, même format de données.
Majeure : `pg_upgrade` (ou dump/restore, ou réplication logique),
avec `--check`, sauvegarde testée, `ANALYZE` après, plan de retour.
**R8.** `timestamptz` stocke un instant absolu (UTC) et l'affiche selon le
`TimeZone` de la session : pas d'ambiguïté lors des changements
d'heure ni entre sites. `timestamp` (sans fuseau) est ambigu.
**R9.** Exemples : `shared_buffers`, `max_connections`, `port`,
`listen_addresses`, `ssl`, `shared_preload_libraries`, `wal_level`,
`max_wal_senders`, `archive_mode`.
**R10.** 1) Figer les écritures si possible. 2) Restaurer la base physique
sur un serveur **isolé**. 3) `recovery_target_time` juste avant 14h32 +
`restore_command` vers les archives WAL. 4) Démarrer, vérifier les
données. 5) Réintégrer **uniquement** les lignes manquantes
(`INSERT ... SELECT ... WHERE NOT EXISTS`). 6) Post-mortem et durcissement
des droits du batch.
</details>

---

## 84. Pour aller plus loin

**Documentation officielle :**

- https://www.postgresql.org/docs/current/ — la référence absolue,
  par version (remplacez `current` par `17`, `16`, `15`).
- Wiki PostgreSQL : https://wiki.postgresql.org/ (tuning, réplication).

**Outils cités dans ce guide :**

| Outil | Usage |
|---|---|
| PgBouncer | Pool de connexions (au-delà de ~200 connexions) |
| Barman | Sauvegarde physique + WAL + PITR industrialisés |
| Patroni | Failover automatique (avec etcd) |
| pgbadger | Rapports HTML des logs PostgreSQL |
| pg_partman | Partitionnement automatique |
| `check_postgres` / `postgres_exporter` | Supervision Nagios / Prometheus |
| PostGIS | Extension géographique (SIG) |
| pg_trgm | Recherche floue / similarité de texte |

**Livres :**

- *PostgreSQL 17 Administration Cookbook* — Packt (recettes opérationnelles).
- *The Art of PostgreSQL* — Dimitri Fontaine (SQL avancé, modélisation).

**Prochaines étapes conseillées :**

1. Monter un labo : primaire + standby + Barman sur 3 VM Debian.
2. Simuler chaque scénario de la section 58 (game day).
3. Brancher `pg_stat_statements` + exporter Prometheus sur la maquette.
4. Tester `pg_upgrade` 16→17 sur une copie avant la prochaine EOL.

---

## 85. Journal d'exploitation : modèle de fiche incident

```markdown
# Fiche incident — <date> <heure>
## Symptôme
## Détection (alerte ? utilisateur ?)
## Diagnostic (commandes passées, résultats)
## Cause racine
## Actions correctives
## Actions préventives (droits, alertes, doc à jour ?)
## RTO / RPO constatés
```

> 💡 Tenez ce journal dans votre wiki d'équipe. Dans six mois, c'est lui
> qui fera la différence entre « on a déjà vu ça » et une nouvelle
> nuit blanche.

---

*Fin du guide — bon courage, et que vos checkpoints soient espacés
et vos WAL bien archivés.*
