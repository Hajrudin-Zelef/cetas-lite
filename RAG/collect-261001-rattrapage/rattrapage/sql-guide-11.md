---
id: collect-261001-rattrapage/rattrapage/sql-guide-11
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-01", "2026-09-26", "2026-09-30", "2026-10-01"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [2110, 2302]
sha256: 93bf20d57fa74ee6dfc30800093451333fd53ab75d1794bdeddd27ef1c6c2e3d
---

# Guide SQL complet — du SELECT à l'optimisation

```sql
-- PostgreSQL
SELECT pg_cancel_query(12345);      -- poli : annule la requête
SELECT pg_terminate_backend(12345); -- brutal : tue la connexion
-- MySQL
KILL 12345;            -- idem
KILL QUERY 12345;      -- tue juste la requête
```

---

## 62. Sécurité — comprendre l'injection SQL

L'injection = concaténer une entrée utilisateur **dans** la requête. Exemple d'horreur :

```python
# ❌ JAMAIS ÇA (exemple Python)
saisie = input("Référence : ")          # l'attaquant tape : ' OR '1'='1
query = f"SELECT * FROM equipements WHERE reference = '{saisie}'"
# Requête réelle : SELECT * FROM equipements WHERE reference = '' OR '1'='1'
# → renvoie TOUTE la table. Variante : '; DROP TABLE employes; --
```

Ce qui est en jeu : lecture de toutes les données, modification, suppression,
voire exécution système selon la configuration. **Toute** entrée (formulaire, API,
fichier d'import, nom de fichier !) est suspecte.

Défenses en profondeur :

1. **Requêtes paramétrées** (section 63) — la défense n°1, non négociable.
2. Principe du moindre privilège : le compte applicatif n'a que les droits nécessaires
   (SELECT sur certaines tables, jamais DROP/ALTER ; voir section 64).
3. Ne jamais exposer les messages d'erreur SQL bruts à l'utilisateur final.
4. Validation des entrées côté application (type, longueur, format) — complément, pas substitut.
5. Pas de construction dynamique de SQL avec des identifiants (noms de tables/colonnes)
   venus de l'utilisateur ; si inévitable : liste blanche stricte.

---

## 63. Requêtes paramétrées — la bonne pratique (exemples)

Le principe : la requête est envoyée **une fois** avec des placeholders,
les valeurs **séparément** — le SGBD ne les confond jamais avec du code.

```python
# ✅ Python + psycopg (PostgreSQL) — placeholder %s
import psycopg
with psycopg.connect("dbname=sql_formation") as conn:
    with conn.cursor() as cur:
        cur.execute(
            "SELECT * FROM equipements WHERE reference = %s AND statut = %s",
            ("UPS-001", "en_service"),   # valeurs : jamais concaténées
        )
        rows = cur.fetchall()

# ✅ Python + mysql-connector — placeholder %s aussi (ne pas confondre avec le % Python !)
# cur.execute("SELECT * FROM equipements WHERE id = %s", (101,))
```

```php
<?php
// ✅ PHP + PDO — placeholders nommés
$pdo = new PDO($dsn, $user, $pass, [PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION]);
$stmt = $pdo->prepare("SELECT * FROM interventions WHERE technicien_id = :tech AND priorite <= :prio");
$stmt->execute(['tech' => 2, 'prio' => 2]);
$rows = $stmt->fetchAll();
?>
```

```bash
# ✅ Shell : psql avec variables (pas d'interpolation bash dans le SQL !)
psql -d sql_formation -v ref="UPS-001" -c "SELECT * FROM equipements WHERE reference = :'ref';"
# ⚠️ :'ref' (avec apostrophes) = valeur quotée ; :ref sans apostrophes = identifiant
```

Règle : **si vous voyez une concaténation de variable dans du SQL, c'est un bug de sécurité.**
Chercher `+ "` / `f"` / `. "` dans le code = audit express.

---

## 64. Droits et rôles — moindre privilège

```sql
-- Rôle applicatif lecture seule (dashboard Grafana, exports)
CREATE ROLE lecture_seule LOGIN PASSWORD 'mot_de_passe_solide';
GRANT CONNECT ON DATABASE sql_formation TO lecture_seule;
GRANT USAGE ON SCHEMA public TO lecture_seule;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO lecture_seule;
-- Pour les futures tables :
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO lecture_seule;

-- Rôle technicien : lire + insérer des interventions, pas toucher au reste
CREATE ROLE technicien LOGIN PASSWORD '...';
GRANT SELECT ON equipements, employes TO technicien;
GRANT SELECT, INSERT, UPDATE ON interventions TO technicien;
-- MySQL :
-- CREATE USER 'technicien'@'%' IDENTIFIED BY '...';
-- GRANT SELECT ON sql_formation.equipements TO 'technicien'@'%';

-- Révoquer
REVOKE DELETE ON interventions FROM technicien;
DROP ROLE IF EXISTS ancien_prestataire;
```

Bonnes pratiques :

- Un compte **par usage** (appli, dashboard, humain, backup) — jamais le compte
  `postgres`/`root` dans une appli.
- Le compte de **sauvegarde** : SELECT + droits de verrouillage, rien d'autre.
- Auditer périodiquement : PG `\du` / `SELECT * FROM information_schema.role_table_grants;`,
  MySQL `SHOW GRANTS FOR 'user'@'host';`.

---

## 65. Sauvegarde — rappels opérationnels

Le DBA/sysadmin retient : **on ne sauvegarde pas des fichiers, on sauvegarde la capacité
à restaurer**. Tester la restauration, sinon la sauvegarde n'existe pas.

```bash
# PostgreSQL : dump logique (SQL rejouable)
pg_dump -d sql_formation -F p -f sauvegarde_$(date +%F).sql        # format texte
pg_dump -d sql_formation -F c -f sauvegarde_$(date +%F).dump       # format custom (compressé)

# Restauration
psql -d sql_formation_restauree -f sauvegarde_2026-09-26.sql
pg_restore -d sql_formation_restauree sauvegarde_2026-09-26.dump

# MySQL / MariaDB
mysqldump --single-transaction --routines --triggers sql_formation > sauvegarde_$(date +%F).sql
mysql sql_formation_restauree < sauvegarde_2026-09-26.sql
```

| Option | Pourquoi |
|---|---|
| `--single-transaction` (mysqldump) | dump cohérent sans verrouiller (InnoDB) |
| `--routines --triggers` | sinon procédures et triggers **perdus** |
| `-F c` (pg_dump custom) | compressé + restauration sélective (`-t table`) |

Stratégie minimale viable :

- [ ] Dump logique **quotidien** (rétention 7 j) + **hebdo** (rétention 4 sem.) + **mensuel** (1 an).
- [ ] Dump **avant** chaque migration/ALTER sur table métier.
- [ ] Test de restauration **mensuel** sur une base à part, chronométré (connaître son RTO).
- [ ] Sauvegarde **hors site** (règle 3-2-1 : 3 copies, 2 supports, 1 hors site).
- [ ] Chiffrer les dumps contenant des données personnelles/salaires.
- [ ] Bonus PG : archivage WAL + `pg_basebackup` pour du point-in-time-recovery (PITR).

---

## 66. Dix erreurs classiques (et comment les éviter)

**Erreur 1 — `= NULL` au lieu de `IS NULL`.**
Résultat : zéro ligne, sans erreur. Réflexe : NULL ne se compare pas, il se teste.

**Erreur 2 — `NOT IN` avec un NULL dans la liste.**
Résultat : zéro ligne. Réflexe : `NOT EXISTS` (section 31).

**Erreur 3 — Oublier le WHERE sur UPDATE/DELETE.**
Résultat : table entière modifiée. Réflexe : transaction + SELECT de contrôle d'abord (sections 51-52).

**Erreur 4 — Division entière.**
`COUNT(*) * 100 / COUNT(*)` = 0 sous PG. Réflexe : `* 100.0` ou CAST (section 15).

**Erreur 5 — `GROUP BY` incomplet accepté par MySQL laxiste.**
Résultat : valeurs arbitraires silencieuses. Réflexe : `ONLY_FULL_GROUP_BY` activé, GROUP BY complet (section 20).

**Erreur 6 — LEFT JOIN + WHERE sur la table de droite.**
Résultat : INNER JOIN déguisé. Réflexe : condition dans le `ON` (section 24).

**Erreur 7 — `COUNT(*)` après LEFT JOIN pour compter les enfants.**
Résultat : 1 au lieu de 0 pour les parents sans enfant. Réflexe : `COUNT(table_droite.id)` (section 24).

**Erreur 8 — BETWEEN sur TIMESTAMP pour un mois.**
`BETWEEN '2026-09-01' AND '2026-09-30'` rate le 30 septembre après minuit.
Réflexe : `>= '2026-09-01' AND < '2026-10-01'` (section 6).

**Erreur 9 — `SELECT *` en production / dans une vue.**
Résultat : colonnes surprises après ALTER, perfs dégradées. Réflexe : liste explicite (section 4).

**Erreur 10 — Concaténer des variables dans le SQL.**
Résultat : injection SQL (section 62). Réflexe : requêtes paramétrées, toujours (section 63).

**Erreur 11 — Comparer des dates en texte non ISO.**
`'26/09/2026' > '01/10/2026'` est TRUE en tri alphabétique ! Réflexe : type DATE + format ISO (section 16).

**Erreur 12 — Oublier que les agrégats ignorent NULL.**
`AVG(cout_eur)` sur (100, NULL) = 100, pas 50. Réflexe : COALESCE si la sémantique l'exige (section 19).

---

## 67. Piège : les doublons silencieux

