---
id: collect-261001-rattrapage/rattrapage/sql-guide-14
title: "Guide SQL complet — du SELECT à l'optimisation"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/sql_guide.md
source_anchor: ""
source_lines: [2668, 2814]
sha256: 29ee857513a2e8be1060ae71477b0a36fc49fcb42160161c47b0044f4c87917d
---

# Guide SQL complet — du SELECT à l'optimisation

| Terme | Définition courte |
|---|---|
| ACID | Atomicité, Cohérence, Isolation, Durabilité : garanties d'une transaction |
| Agrégat | Fonction réduisant un ensemble de lignes (COUNT, SUM, AVG, MIN, MAX) |
| Alias | Nom temporaire donné à une table/colonne dans une requête (AS) |
| Anti-join | Requête « les X sans Y » (NOT EXISTS / LEFT JOIN ... IS NULL) |
| Autovacuum | Processus PG qui nettoie et met à jour les stats automatiquement |
| B-tree | Structure d'index par défaut, efficace pour =,<,>,BETWEEN, tri |
| Cardinalité | Nombre de valeurs distinctes d'une colonne |
| CASCADE | Propagation d'une suppression/mise à jour via les FK |
| CHECK | Contrainte validant une règle métier sur les données |
| CTE | Common Table Expression : requête nommée avec WITH |
| Curseur (keyset) | Pagination par « après l'id X » au lieu d'OFFSET |
| DDL / DML / DQL / DCL / TCL | Familles d'ordres SQL (section 2) |
| Deadlock | Deux transactions qui s'attendent mutuellement → l'une est tuée |
| DISTINCT | Déduplication des lignes du résultat |
| Dump | Export logique d'une base (SQL ou custom) |
| EXPLAIN | Plan d'exécution estimé (ANALYZE = réel) |
| Fan-out | Multiplication des lignes par jointure de 2 tables « plusieurs » |
| Fenêtre (fonction) | Calcul sur un ensemble lié à la ligne courante, sans réduire les lignes |
| FK (clé étrangère) | Colonne référençant la PK d'une autre table |
| FULL OUTER JOIN | Garde toutes les lignes des deux tables (émulé sous MySQL) |
| GROUP BY | Regroupement avant agrégation |
| HAVING | Filtre sur les groupes (après agrégation) |
| IDENTITÉ | Colonne auto-incrémentée standard (GENERATED ... AS IDENTITY) |
| Index | Structure accélérant les recherches (B-tree, partiel, d'expression...) |
| Index couvrant | Index contenant toutes les colonnes requises (Index Only Scan) |
| Injection SQL | Exécution de code via une entrée concaténée dans la requête |
| Isolation (niveau) | Degré de visibilité entre transactions concurrentes |
| JOIN | Combinaison de tables sur une condition |
| LATERAL | Sous-requête du FROM référençant les tables précédentes (PG) |
| Logique à 3 valeurs | TRUE / FALSE / UNKNOWN (NULL) |
| Matérialisée (vue) | Vue dont le résultat est stocké et rafraîchi explicitement |
| MVCC | Multi-Version Concurrency Control : lecteurs jamais bloqués par écrivains (PG) |
| N+1 | Antipattern : 1 requête + N requêtes unitaires en boucle |
| Normalisation | Découpage en tables pour éliminer redondances (1NF → 3NF) |
| NOT NULL | Contrainte interdisant l'absence de valeur |
| NULL | Absence/inconnu de valeur — ni 0 ni '' |
| OFFSET | Saut de N lignes (pagination, à éviter en profondeur) |
| Optimiseur | Composant choisissant le plan d'exécution |
| ORm | Object-Relational Mapping (couche applicative) |
| Partitionnement | Découpage physique d'une grosse table (par date, etc.) |
| PIVOT | Transformation lignes → colonnes (via CASE en SQL standard) |
| PK (clé primaire) | Identifiant unique et non nul d'une table |
| Plan d'exécution | Stratégie choisie par l'optimiseur (lisible via EXPLAIN) |
| Redo/WAL | Journal des modifications garantissant la durabilité |
| Requête corrélée | Sous-requête référençant la requête externe |
| REPEATABLE READ | Niveau d'isolation : vision figée pendant la transaction |
| ROLLBACK | Annulation d'une transaction |
| ROLLUP/CUBE | Sous-totaux multi-niveaux dans un GROUP BY |
| Sargable | Prédicat pouvant utiliser un index (Search ARGument ABLE) |
| SAVEPOINT | Point de rollback partiel dans une transaction |
| Schéma | Espace de noms regroupant des objets (public, ...) |
| Sélectivité | Proportion de lignes filtrées (basse = bon candidat index) |
| Semi-join | « les X ayant au moins un Y » (EXISTS / IN) |
| Séquence | Compteur servant aux auto-incréments (PG) |
| SLA | Service Level Agreement : engagement de délai/qualité |
| Sous-requête | SELECT imbriqué (WHERE, FROM, SELECT) |
| Statistiques | Infos sur les données guidant l'optimiseur (ANALYZE) |
| Surrogate key | Clé primaire technique (id auto), vs clé naturelle métier |
| THREE-valued logic | Voir « Logique à 3 valeurs » |
| Transaction | Lot atomique d'ordres (BEGIN/COMMIT/ROLLBACK) |
| Trigger | Code exécuté automatiquement sur INSERT/UPDATE/DELETE |
| UNIQUE | Contrainte d'unicité (plusieurs NULL permis en standard) |
| UPSERT | INSERT ou UPDATE si conflit (ON CONFLICT / ON DUPLICATE KEY) |
| VACUUM | Nettoyage PG des lignes mortes (libère de la place, fige les stats) |
| Verrou (lock) | Mécanisme d'exclusion mutuelle entre transactions |
| Vue | Requête stockée, réévaluée à chaque usage |

---

## 78. Quiz — 10 questions (réponses en section 79)

**Q1.** Que renvoie `SELECT * FROM interventions WHERE date_cloture = NULL;` ?
a) Les interventions en cours — b) Rien, toujours — c) Erreur de syntaxe

**Q2.** Dans `SELECT COUNT(*) FROM interventions`, que compte-t-on si `cout_eur`
contient des NULL ?
a) Toutes les lignes — b) Seulement les lignes où cout_eur est non NULL — c) Erreur

**Q3.** Quelle est la différence entre `RANK()` et `DENSE_RANK()` sur des ex æquo ?
a) Aucune — b) RANK laisse un trou après les ex æquo, DENSE_RANK non — c) L'inverse

**Q4.** Pourquoi ce LEFT JOIN se comporte comme un INNER JOIN ?
```sql
SELECT e.nom FROM employes e
LEFT JOIN interventions i ON i.technicien_id = e.id
WHERE i.priorite = 1;
```

**Q5.** Écrire la requête « les équipements sans aucune intervention » de **deux**
manières différentes.

**Q6.** Corriger : `SELECT type_interv, COUNT(*) FROM interventions HAVING COUNT(*) > 2;`
(il manque quelque chose).

**Q7.** Un rapport calcule `AVG(prix_moyen_par_mois)` sur 12 mois. Pourquoi le résultat
peut-il être faux, et que faire à la place ?

**Q8.** Citer **trois** défenses contre l'injection SQL, dans l'ordre de priorité.

**Q9.** `SELECT * FROM interventions ORDER BY date_ouverture LIMIT 10 OFFSET 100000;`
rame. Proposer une méthode plus rapide pour la page suivante.

**Q10.** Quelle commande MySQL/MariaDB garantit un dump **cohérent sans verrouiller**
les tables InnoDB, et quelles options évitent de perdre routines et triggers ?

---

## 79. Quiz — réponses détaillées

**R1. b) Rien, toujours.** `= NULL` donne UNKNOWN pour chaque ligne, et WHERE ne garde
que TRUE. Il faut `IS NULL`. C'est l'erreur n°1 (section 66).

**R2. a) Toutes les lignes.** `COUNT(*)` compte les lignes, pas les valeurs —
les NULL ne le gênent pas. C'est `COUNT(cout_eur)` qui ignorerait les NULL (section 19).

**R3. b)** Exemple avec deux 2èmes ex æquo : RANK → 1, 2, 2, 4 (trou) ;
DENSE_RANK → 1, 2, 2, 3 (continu). ROW_NUMBER → 1, 2, 3, 4 (unique arbitraire) (section 37).

**R4.** Le `WHERE i.priorite = 1` élimine les lignes où `i.*` vaut NULL (les employés
sans intervention prioritaire), car `NULL = 1` est UNKNOWN. Le LEFT JOIN est
« re-transformé » en INNER. Correction : mettre la condition dans le `ON`
(section 24).

**R5.**
```sql
-- Méthode 1 : NOT EXISTS (recommandée)
SELECT reference FROM equipements e
WHERE NOT EXISTS (SELECT 1 FROM interventions i WHERE i.equipement_id = e.id);
-- Méthode 2 : LEFT JOIN + IS NULL
SELECT e.reference FROM equipements e
LEFT JOIN interventions i ON i.equipement_id = e.id
WHERE i.id IS NULL;
```
(Section 31. `NOT IN` est la 3e méthode mais piégée par NULL.)

**R6.** Il manque le `GROUP BY` : `HAVING` filtre des groupes, il faut les former.
```sql
SELECT type_interv, COUNT(*) AS nb FROM interventions GROUP BY type_interv HAVING COUNT(*) > 2;
```
Ou avec alias CTE pour la portabilité (section 21).

**R7.** C'est la **moyenne des moyennes** : fausse sauf effectifs mensuels identiques
(section 68). À la place : recalculer sur les données brutes —
`SUM(total) / SUM(effectif)` ou `AVG()` directement sur la colonne source.

