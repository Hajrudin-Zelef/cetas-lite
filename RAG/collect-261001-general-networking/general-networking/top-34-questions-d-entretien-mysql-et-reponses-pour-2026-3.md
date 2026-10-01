---
id: collect-261001-general-networking/general-networking/top-34-questions-d-entretien-mysql-et-reponses-pour-2026-3
title: "top-34-questions-d-entretien-mysql-et-reponses-pour-2026"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2024-12-24"]
keywords: []
source: docs/RAG/collect-261001-general-networking/top-34-questions-d-entretien-mysql-et-reponses-pour-2026.md
source_anchor: ""
source_lines: [308, 507]
sha256: 3482b814fc9602dd3ef7ef97637910b192d9fefbb8514ede70b89e5cf49f3db8
---

# top-34-questions-d-entretien-mysql-et-reponses-pour-2026

- 
La fonction `LENGTH()` affiche le nombre de caractères d'un nom.
- 
`UPPER()` et`LOWER()` convertissent le texte en majuscules ou minuscules.
- 
`CONCAT()` assemble le prénom et le nom dans une même colonne.
- 
`SUBSTRING()` extrait des parties spécifiques d'un texte. Par exemple, isoler le mois à partir d'une date de naissance.

Exemple de requête :

```
SELECT 
    UPPER(first_name) AS upper_name,
    CONCAT(first_name, ' ', last_name) AS full_name,
    SUBSTRING(birthdate, 6, 2) AS birth_month,
    TRIM(last_name) AS trimmed_last_name,
    REPLACE(first_name, 'a', '@') AS replaced_name
FROM employees;
```
Cette requête :

- 
Convertit les prénoms en majuscules.
- 
Combine prénom et nom en un nom complet.
- 
Extrait le mois de la colonne `birthdate` .
- 
Supprime les espaces superflus des noms.
- 
Remplace toutes les occurrences de "a" par "@" dans les prénoms.

### 20. Comment mettre à jour une ligne précise en SQL ?

Utilisez l'instruction `UPDATE` avec une clause `WHERE` pour identifier l'enregistrement à modifier. 

Par exemple, pour mettre à jour le genre du film « Inception » (2010) en « Sci-Fi » :

```
UPDATE movies
SET genre = 'Sci-Fi'
WHERE movie_title = 'Inception' AND year = 2010;
```
Ici, `UPDATE movies` précise la table visée, et la clause `WHERE` cible la ligne où le titre est « Inception » et l'année « 2010 ». 

## Questions MySQL avancées

Les questions avancées évaluent votre capacité à gérer des scénarios MySQL complexes et donnent un aperçu de votre façon de décider.

### 21. Qu'est-ce qu'un trigger dans MySQL ? Comment l'implémenter ?

Un trigger (déclencheur) est un ensemble d'actions exécutées lorsqu'un événement survient en base. Les triggers peuvent s'exécuter avant ou après des événements comme `INSERT`, `UPDATE` ou `DELETE`. 

Par exemple, supposons une table `orders` où sont ajoutées de nouvelles commandes. Nous pouvons créer un trigger qui journalise chaque commande dans une table `order_history` :

```
CREATE TRIGGER after_order_insert
AFTER INSERT ON orders
FOR EACH ROW
BEGIN
    INSERT INTO order_history (order_id, action, timestamp)
    VALUES (NEW.order_id, 'inserted', NOW());
END;
```
Après l'exécution du trigger, la table `order_history` est automatiquement mise à jour :

| **history_id** | **order_id** | **action** | **timestamp** | 
| **1** | 1 | inserted | 2024-12-24 10:00:00 | 
| **2** | 2 | inserted | 2024-12-24 11:00:00 | 

### 22. Pourquoi l'ajout d'un index accélère-t-il les requêtes SQL ?

Sans index, la base doit scanner chaque ligne pour trouver une entrée. Un index s'apparente à une table des matières, permettant d'accéder directement aux lignes pertinentes. L'ajout d'un index réduit donc le temps de recherche et accélère les requêtes.

Les index s'appuient généralement sur des structures de données comme les B-arbres (B-trees) ou des tables de hachage, qui facilitent recherches, lookups et parcours par plage.

Exemple de création d'index :

```
-- Sans index :
SELECT * FROM employees WHERE last_name = 'Smith';
-- Ajout d'un index sur la colonne last_name :
CREATE INDEX idx_last_name ON employees(last_name);
-- Avec l'index, la base localise rapidement les lignes portant 'Smith' en last_name.
```
Les index ont aussi des inconvénients :

- 
Écritures ralenties : les opérations `INSERT` ,`UPDATE` et`DELETE` sont plus lentes car l'index doit être mis à jour à chaque changement.
- 
Coût de stockage : les index consomment de l'espace disque supplémentaire.

### 23. Quels types de données utiliser pour le poids et le prix d'un produit, et pourquoi ?

Pour le poids, `DECIMAL` est en général le choix le plus sûr. `FLOAT` et `REAL` stockent des décimales mais utilisent l'arithmétique en virgule flottante, source de petits écarts d'arrondi.

Pour des poids nécessitant de la précision (expédition, inventaire), `DECIMAL(8, 3)` offre un contrôle précis avec 3 décimales sans mauvaises surprises d'arrondi. `FLOAT` n'est acceptable que si une petite marge d'erreur est tolérée.

### 24. Comment trouver des doublons en SQL avec une fonction fenêtre ?

Voici comment repérer des doublons avec la fonction `ROW_NUMBER()` :

```
WITH DuplicateCheck AS (
    SELECT product_name, 
           category,
           ROW_NUMBER() OVER(
               PARTITION BY product_name, category 
               ORDER BY id
           ) AS row_num
    FROM sales
)
SELECT *
FROM DuplicateCheck
WHERE row_num > 1;
```
Explication :

1. `ROW_NUMBER()` numérote chaque ligne du résultat. 

2. `PARTITION BY` regroupe par `product_name` et `category`. 

3. Dans chaque groupe, la numérotation commence à 1.

4. Toute ligne avec `row_num` > 1 est un doublon. 

Par exemple, pour ces enregistrements :

- Product A, Category X, row_num = 1
- Product A, Category X, row_num = 2 (doublon)
- Product B, Category Y, row_num = 1

La requête retournera la deuxième ligne, car son `row_num` est supérieur à 1.

### 25. Comment créer et utiliser une procédure stockée avec paramètres dans MySQL ? Donnez un exemple.

Les procédures stockées permettent d'enregistrer et réutiliser des requêtes complexes pour rendre les opérations plus efficaces et maintenables. Voici un exemple avec paramètres.

Supposons une base d'étudiants et la nécessité de filtrer par âge. Créons une procédure prenant un paramètre d'âge :

D'abord, une procédure simple avec paramètre d'entrée :

```
CREATE PROCEDURE get_student_info(IN age INT)
BEGIN
    SELECT * FROM student WHERE student.age = age;
END;
```
Pour l'exécuter, on la `CALL` avec l'âge souhaité :

`CALL get_student_info(21);`
On peut aller plus loin avec des paramètres de sortie. Par exemple, compter les étudiants d'un âge donné :

```
CREATE PROCEDURE count_students_by_age(IN age INT, OUT student_count INT)
BEGIN
    SELECT COUNT(*) INTO student_count FROM students WHERE students.age = age;
END;
```
Pour récupérer le résultat :

```
SET @count = 0;
CALL count_students_by_age(21, @count);
SELECT @count AS total_students;
```
### 26. Pourquoi l'intégrité référentielle est-elle importante en base ?

L'intégrité référentielle assure la cohérence des relations entre tables. La clé étrangère garantit que les valeurs d'une table correspondent à une valeur unique de la table référencée.

Exemple concret : vous gérez une base e-commerce avec une table `Customers` et une table `Orders`. Chaque commande doit appartenir à un client réel. L'intégrité référentielle, via les clés étrangères, s'assure que :

- On ne peut pas créer une commande pour un client inexistant.
- On ne peut pas supprimer un client ayant des commandes (sauf règle explicite sur le devenir de ces commandes).
- On ne peut pas modifier l'ID d'un client ayant des commandes existantes.

Ainsi, en créant une contrainte de clé étrangère :

```
ALTER TABLE Orders
ADD FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID);
```
La base applique automatiquement ces règles :

- 
Chaque `CustomerID` dans`Orders` doit exister dans`Customers` .
- 
Toute tentative de violation (insertion d'un `CustomerID` invalide) est rejetée.

Cela évite des incohérences critiques, telles que des commandes impossibles à rattacher à des clients, ou des rapports lacunaires.

## Questions MySQL pour administrateurs de bases de données

Si vous postulez spécifiquement à un poste d'administrateur de bases de données, voici des questions possibles.

### 27. Pourquoi une application à grande échelle recourt-elle au sharding ? Quels sont les défis associés ?

Le sharding permet de répartir un gros volume de données sur plusieurs serveurs. Chaque shard contient une partie des données. En répartissant la charge, on évite du matériel haut de gamme. La vitesse et l'évolutivité s'améliorent, mais il existe des défis :

