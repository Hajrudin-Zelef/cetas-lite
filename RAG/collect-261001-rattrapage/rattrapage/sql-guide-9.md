---
id: collect-261001-rattrapage/rattrapage/sql-guide-9
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2020-01-01", "2026-09-01"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [1698, 1913]
sha256: ee57c72ec5af01cf585d5368423e3f4463fe922362457d331b4af57bdd1dc43a
---

# Guide SQL complet — du SELECT à l'optimisation

-- Récupérer l'id généré (PG)
INSERT INTO interventions (equipement_id, technicien_id, type_interv, priorite, description)
VALUES (103, 4, 'preventif', 4, 'Nettoyage + remplacement kit maintenance')
RETURNING id, date_ouverture;
-- MySQL : utiliser LAST_INSERT_ID() après l'INSERT
```

Règles :

- Liste de colonnes **explicite** : protège contre l'ajout d'une colonne.
- `DEFAULT` pour forcer la valeur par défaut : `VALUES (109, 'SRV-002', ..., DEFAULT)`.
- Gros imports : `COPY` (PG) / `LOAD DATA INFILE` (MySQL), pas des millions d'INSERT
  unitaires (facteur ×100 en vitesse).

---

## 50. INSERT ... SELECT — copier/transformer en masse

```sql
-- Archiver les interventions clôturées de plus de 3 ans
INSERT INTO interventions_archive
    (id, equipement_id, technicien_id, date_ouverture, date_cloture, description, cout_eur)
SELECT id, equipement_id, technicien_id, date_ouverture, date_cloture, description, cout_eur
FROM interventions
WHERE date_cloture < CURRENT_DATE - INTERVAL '3 years';
-- MySQL : WHERE date_cloture < CURDATE() - INTERVAL 3 YEAR

-- Peupler une table de reporting
INSERT INTO rapport_mensuel (mois, nb_interv, cout_total)
SELECT DATE_TRUNC('month', date_ouverture), COUNT(*), SUM(cout_eur)
FROM interventions
GROUP BY 1;
```

C'est transactionnel : tout ou rien (section 54). Pour des volumes énormes,
découper par lots (ex. par mois) pour ne pas gonfler le journal de transactions.

---

## 51. UPDATE — modifier (avec prudence)

```sql
-- ✅ Ciblé par PK
UPDATE equipements
SET statut = 'en_maintenance', localisation = 'Atelier'
WHERE id = 102;

-- Mise à jour calculée
UPDATE onduleurs
SET autonomie_min = autonomie_min - 2
WHERE date_dernier_test < CURRENT_DATE - INTERVAL '1 year';
-- MySQL : WHERE date_dernier_test < CURDATE() - INTERVAL 1 YEAR

-- UPDATE avec jointure : répercuter une info
-- PostgreSQL :
UPDATE interventions i
SET sla_respecte = TRUE
FROM equipements e
WHERE e.id = i.equipement_id AND i.date_cloture IS NOT NULL AND e.statut <> 'reforme';

-- MySQL :
UPDATE interventions i
JOIN equipements e ON e.id = i.equipement_id
SET i.sla_respecte = TRUE
WHERE i.date_cloture IS NOT NULL AND e.statut <> 'reforme';
```

**Discipline UPDATE/DELETE** (à afficher au-dessus du clavier) :

1. Écrire d'abord le `SELECT` avec le même WHERE → vérifier le périmètre.
2. Encadrer dans une transaction : `BEGIN; UPDATE...; SELECT...;` puis `COMMIT` ou `ROLLBACK`.
3. Ne **jamais** exécuter sans WHERE sur une table métier (ou alors `WHERE TRUE` écrit
   volontairement pour marquer l'intention... et relu deux fois).

---

## 52. DELETE — supprimer (avec encore plus de prudence)

```sql
-- Supprimer les interventions d'un équipement réformé (après archivage !)
BEGIN;
SELECT COUNT(*) FROM interventions WHERE equipement_id = 109;  -- contrôle
DELETE FROM interventions WHERE equipement_id = 109;
SELECT COUNT(*) FROM interventions WHERE equipement_id = 109;  -- 0 attendu
COMMIT;   -- ou ROLLBACK si le contrôle échoue
```

```sql
-- DELETE avec sous-requête : purger les brouillons sans équipement
DELETE FROM interventions
WHERE equipement_id NOT IN (SELECT id FROM equipements WHERE id IS NOT NULL);
-- (mieux : NOT EXISTS, section 31)

-- PostgreSQL : RETURNING pour auditer ce qu'on supprime
DELETE FROM interventions WHERE date_cloture < DATE '2020-01-01'
RETURNING id, date_ouverture, cout_eur;
```

Alternative sûre au DELETE physique : **suppression logique** (`actif = FALSE`,
`date_suppression`) — conserve l'historique, annulable, mais penser à filtrer
`WHERE actif` dans toutes les requêtes (ou utiliser une vue, section 56).

---

## 53. UPSERT — insérer ou mettre à jour

Synchroniser un inventaire (ex. import GLPI quotidien) sans doublons :

```sql
-- PostgreSQL : ON CONFLICT
INSERT INTO equipements (id, reference, type_equip, marque, modele, statut)
VALUES (101, 'UPS-001', 'onduleur', 'APC', 'Easy UPS 3S 40 kVA', 'en_service')
ON CONFLICT (id) DO UPDATE
SET statut = EXCLUDED.statut,
    modele  = EXCLUDED.modele;
-- Variantes : ON CONFLICT (reference) DO NOTHING  (ignorer les déjà-là)

-- MySQL/MariaDB : ON DUPLICATE KEY UPDATE
INSERT INTO equipements (id, reference, type_equip, marque, modele, statut)
VALUES (101, 'UPS-001', 'onduleur', 'APC', 'Easy UPS 3S 40 kVA', 'en_service')
ON DUPLICATE KEY UPDATE
    statut = VALUES(statut),
    modele = VALUES(modele);
-- MySQL 8.0.20+ : VALUES() déprécié → alias : ... AS n ON DUPLICATE KEY UPDATE statut = n.statut
```

⚠️ `ON CONFLICT` exige une contrainte d'unicité réelle (PK/UNIQUE) sur la cible.
Sans index unique, pas d'arbitre → erreur. C'est voulu : l'upsert n'a de sens
que si « déjà-là » est défini sans ambiguïté.

---

## 54. Transactions et ACID

Une transaction = un lot d'ordres **atomique** : tout est validé ou rien.

```sql
-- Transfert d'un équipement d'un site à l'autre + traçabilité : indissociable
BEGIN;
UPDATE equipements SET localisation = 'Annexe B' WHERE id = 103;
INSERT INTO interventions (equipement_id, technicien_id, type_interv, priorite, description)
VALUES (103, 3, 'installation', 3, 'Déménagement vers annexe B');
COMMIT;      -- tout est validé
-- En cas de problème : ROLLBACK;  (annule tout depuis BEGIN)
```

ACID :

| Lettre | Sens | Concrètement |
|---|---|---|
| **A**tomicité | tout ou rien | panne en milieu de transaction → rien n'est validé |
| **C**ohérence | contraintes respectées | une transaction ne laisse jamais la base incohérente |
| **I**solation | transactions concurrentes isolées | niveaux réglables (section 55) |
| **D**urabilité | validé = persisté | survit au crash (journal WAL / redo log) |

Points de vigilance :

- `SAVEPOINT` : rollback partiel dans une grosse transaction.
- Sous MySQL, certaines instructions (DDL : CREATE/ALTER/DROP/TRUNCATE) provoquent
  un **commit implicite** → découper les migrations.
- Ne jamais laisser une transaction ouverte « en attendant » (verrous conservés →
  blocages en cascade, voir section 69).

```sql
BEGIN;
UPDATE onduleurs SET nb_batteries = 32 WHERE equipement_id = 101;
SAVEPOINT sp1;
UPDATE onduleurs SET puissance_kva = -5 WHERE equipement_id = 101;  -- CHECK va échouer
ROLLBACK TO sp1;    -- annule juste la 2e mise à jour
COMMIT;            -- la 1re est conservée
```

---

## 55. Niveaux d'isolation — comprendre les anomalies

| Niveau | Dirty read | Non-repeatable read | Phantom read | Usage |
|---|---|---|---|---|
| READ UNCOMMITTED | possible | possible | possible | quasi jamais (PG le traite comme READ COMMITTED) |
| READ COMMITTED | non | possible | possible | **défaut PG/MySQL** |
| REPEATABLE READ | non | non | possible (PG : non) | rapports cohérents |
| SERIALIZABLE | non | non | non | criticité max, coût max |

```sql
-- Rapport mensuel cohérent : figer la vision des données
BEGIN ISOLATION LEVEL REPEATABLE READ;
SELECT SUM(cout_eur) FROM interventions WHERE date_ouverture >= '2026-09-01';
-- ... autres requêtes du rapport : mêmes données, même si ça bouge en face ...
COMMIT;

-- Forcer au niveau session (avec parcimonie)
SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
```

En pratique sysadmin : le défaut (READ COMMITTED) convient à 95 % des cas.
Passer en REPEATABLE READ pour les exports/reportings multi-requêtes qui doivent
« voir » la même photo. SERIALIZABLE : réservé aux invariants critiques
(ex. ne jamais dépasser un stock) — avec gestion des retries (erreur 40001).

---

## 56. Vues — requêtes stockées

```sql
-- Vue : le parc "vu" par les techniciens (sans les salaires, sans les réformés)
CREATE OR REPLACE VIEW v_parc_actif AS
SELECT e.id, e.reference, e.type_equip, e.marque, e.modele, e.localisation, e.statut,
       o.puissance_kva, o.autonomie_min
FROM equipements e
LEFT JOIN onduleurs o ON o.equipement_id = e.id
WHERE e.statut <> 'reforme';

-- Usage : comme une table
SELECT * FROM v_parc_actif WHERE type_equip = 'onduleur' ORDER BY puissance_kva DESC;
```

