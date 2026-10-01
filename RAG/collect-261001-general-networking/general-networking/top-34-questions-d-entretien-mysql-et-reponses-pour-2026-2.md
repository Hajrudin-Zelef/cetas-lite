---
id: collect-261001-general-networking/general-networking/top-34-questions-d-entretien-mysql-et-reponses-pour-2026-2
title: "top-34-questions-d-entretien-mysql-et-reponses-pour-2026"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: ["2024-01-01", "2024-12-31", "9999-12-31"]
keywords: []
source: docs/RAG/collect-261001-general-networking/top-34-questions-d-entretien-mysql-et-reponses-pour-2026.md
source_anchor: ""
source_lines: [130, 307]
sha256: 8df5f7030c078bc0031d09061a9c65a372c70b141f0d243cef31a3482e0eaa9c
---

# top-34-questions-d-entretien-mysql-et-reponses-pour-2026

```
CREATE TEMPORARY TABLE temp_employees (
    id INT,
    name VARCHAR(50)
);
INSERT INTO temp_employees VALUES (1, 'John Doe');
SELECT * FROM temp_employees;
```
### 11. Qu'est-ce qu'une sous-requête en MySQL ? Expliquez avec un exemple.

Une sous-requête (ou requête imbriquée) est une requête incluse dans une autre. Elle permet de décomposer des opérations complexes en étapes plus simples. Par exemple, pour trouver les employés gagnant au-dessus de la moyenne :

```
SELECT first_name, last_name, salary
FROM employees
WHERE salary > (
    SELECT AVG(salary)
    FROM employees
);
```
Décomposition :

1. 
La requête interne `SELECT AVG(salary) FROM employees` calcule d'abord le salaire moyen.
2. 
La requête externe utilise ensuite cette moyenne pour trouver les employés qui la dépassent.

### 12. Comment utiliser l'instruction INSERT pour ajouter des données à une table ? Des bonnes pratiques ?

On utilise l'instruction `INSERT` pour ajouter des données. La syntaxe de base :

```
INSERT INTO table_name (column1, column2, ...) 
VALUES (value1, value2, ...); 
```
Quelques bonnes pratiques à suivre avec `INSERT` :

1. 
Lister explicitement les colonnes. Le code est plus clair et reste robuste si la structure évolue.
2. 
Pour les colonnes `AUTO_INCREMENT` (comme les IDs), ne pas les inclure dans l'`INSERT` . MySQL les gère pour éviter les doublons.
3. 
Être cohérent dans l'usage des guillemets pour les chaînes. Les simples quotes sont souvent préférées, mais les deux fonctionnent.
4. 
Pour insérer plusieurs lignes, privilégier une seule instruction pour de meilleures performances.

### 13. À quoi sert l'attribut AUTO_INCREMENT dans MySQL ?

L'attribut `AUTO_INCREMENT` génère des nombres uniques et séquentiels pour une colonne, en général la clé primaire.

Exemple de création d'une table avec une colonne `AUTO_INCREMENT` :

```
CREATE TABLE employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    department VARCHAR(50)
);
```
Et insertion de lignes :

```
INSERT INTO employees (name, department) VALUES ('John Doe', 'Sales');
INSERT INTO employees (name, department) VALUES ('Jane Smith', 'Marketing');
```
### 14. Qu'est-ce qu'une vue dans MySQL ?

Une vue est une requête enregistrée qui se comporte comme une table virtuelle. Elle permet d'encapsuler une requête complexe sous un nom et de la réutiliser comme une table. Ainsi, vous n'avez pas à retaper la requête entière à chaque fois.

Par exemple, pour simplifier l'interrogation des employés avec le nom de leur service, vous pouvez créer une vue :

```
CREATE VIEW employee_details AS
SELECT 
    e.id,
    e.name,
    d.department_name,
    e.salary
FROM 
    employees e
JOIN 
    departments d ON e.department_id = d.department_id;
```
Vous pouvez ensuite interroger la vue `employee_details` comme une table :

`SELECT * FROM employee_details;`
En revanche, l'insertion et la mise à jour via des vues sont limitées. La plupart sont en lecture seule et masquent l'accès direct aux tables, renforçant ainsi la sécurité. Les vues peuvent aussi ralentir certaines requêtes, car la requête sous-jacente est exécutée à chaque accès.

## Amélioration de SQL pour les débutants

## Questions MySQL intermédiaires

Dans cette section, nous abordons des sujets de niveau intermédiaire. Ces questions servent principalement à évaluer votre connaissance des types de données et de la structure MySQL.

### 15. Que sont les tables versionnées par le système et comment fonctionnent-elles ?

Les tables versionnées par le système conservent l'historique complet des modifications d'une table. Comme elles gardent les versions précédentes de chaque ligne, on peut les utiliser pour l'audit et la récupération de données.

Elles fonctionnent en ajoutant deux colonnes supplémentaires — `StartTime` et `EndTime` — qui enregistrent la période de validité de chaque ligne. Lors des insertions, mises à jour ou suppressions, ces horodatages sont ajustés :

- 
Insertion : une nouvelle ligne est ajoutée avec un `StartTime` défini à l'horodatage courant et un`EndTime` à`9999-12-31 23:59:59` — la valeur`DATETIME` maximale de MySQL, utilisée comme sentinelle pour indiquer "ligne actuellement active".
- 
Mise à jour : le `EndTime` de la ligne d'origine est mis à l'horodatage courant pour la marquer comme non valide, puis une nouvelle ligne mise à jour est créée avec`StartTime` à l'instant courant et`EndTime` "infini".
- 
Suppression : le `EndTime` de la ligne existante est mis à l'horodatage courant, indiquant qu'elle n'est plus valide.

Avec la clause `FOR SYSTEM_TIME` de SQL, vous pouvez interroger l'état de la table à un instant donné ou sur une période. Par exemple :

- 
`FOR SYSTEM_TIME AS OF '2024-01-01'` : récupère l'état de la table tel qu'il était le 1er janvier 2024.
- 
`FOR SYSTEM_TIME BETWEEN '2024-01-01' AND '2024-12-31'` : affiche toutes les lignes valides sur cet intervalle.

### 16. Qu'est-ce qu'une transaction MySQL et comment l'utiliser ?

Une transaction est un ensemble d'opérations exécutées comme une unité. Elle garantit l'intégrité des données en imposant que toutes les opérations réussissent ou échouent ensemble.

Exemple d'utilisation :

```
START TRANSACTION;
UPDATE accounts SET balance = balance - 500 WHERE account_id = 1;
UPDATE accounts SET balance = balance + 500 WHERE account_id = 2;
COMMIT; -- Enregistre définitivement
-- ou
ROLLBACK; -- Annule les changements
```
### 17. Qu'est-ce qu'une contrainte par défaut dans MySQL ? Comment définir une valeur par défaut pour une colonne ?

Une contrainte par défaut attribue une valeur à une colonne quand aucune valeur explicite n'est fournie lors d'un `INSERT`. Elle garantit la validité de la colonne si l'utilisateur l'omet à la saisie.

Exemple de création avec valeur par défaut :

```
CREATE TABLE employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    status VARCHAR(10) DEFAULT 'active'
);
```
Vous pouvez ensuite insérer une ligne sans préciser `status` :

`INSERT INTO employees (name) VALUES ('John Doe');`
Cette approche réduit le risque de `NULL` ou de données invalides sur des colonnes critiques et simplifie les requêtes en évitant de gérer les cas par défaut dans le code.

| **Champ** | **Type** | **Null** | **Clé** | **Défaut** | **Extra** | 
| **id** | INT | NO | PRI | NULL | AUTO_INCREMENT | 
| **name** | VARCHAR(50) | YES |  | NULL |  | 
| **status** | VARCHAR(10) | YES |  | active |  | 

Cette commande est utile car :

- Elle aide les développeurs à comprendre le schéma avant d'écrire des requêtes.
- Elle sert au débogage, notamment sur des bases méconnues.
- Elle permet d'identifier rapidement les contraintes, comme les clés primaires ou les valeurs par défaut.

### 18. Quelle est la différence entre CHAR et VARCHAR dans MySQL ?

Les deux stockent des chaînes, mais diffèrent dans la gestion de l'espace :

- 
`CHAR(n)` stocke toujours exactement`n` caractères, en complétant avec des espaces si besoin. Longueur fixe, ce qui peut être légèrement plus rapide pour des colonnes de longueur uniforme (codes pays, statuts, etc.).
- 
`VARCHAR(n)` ne stocke que les caractères saisis, jusqu'à`n` . Il économise de l'espace pour les données à longueur variable, avec un léger surcoût de gestion de longueur.

```
CREATE TABLE example (
    country_code CHAR(2),      -- Toujours 2 caractères, ex. 'US', 'UK'
    email VARCHAR(255)         -- Longueur variable jusqu'à 255
);
```
Règle pratique : utilisez `CHAR` pour les valeurs à longueur fixe, `VARCHAR` pour le reste.

### 19. Comment utiliser les fonctions de chaîne en SQL pour traiter du texte ?

Différentes fonctions de chaîne aident à manipuler des noms et autres textes. Par exemple :

