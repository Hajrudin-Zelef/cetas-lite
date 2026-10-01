---
id: collect-261001-general-networking/general-networking/mysql-vs-mongodb-choisir-la-bonne-base-de-donnees-pour-votre-projet-2
title: "mysql-vs-mongodb-choisir-la-bonne-base-de-donnees-pour-votre-projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-general-networking/mysql-vs-mongodb-choisir-la-bonne-base-de-donnees-pour-votre-projet.md
source_anchor: ""
source_lines: [67, 141]
sha256: 624adea42bef99031a22f6edcb6cbbdb0a80dd1fe723113aa3748e2ba80d27d6
---

# mysql-vs-mongodb-choisir-la-bonne-base-de-donnees-pour-votre-projet

Le langage de requête de MongoDB(MQL) travaille sur des documents de type JSON dans des collections plutôt que dans des tableaux rigides. Vous récupérez les enregistrements avec `find()`, vous les filtrez avec des opérateurs de comparaison (`$gt`, `$and`, `$or`), vous ne projetez que les champs dont vous avez besoin, et vous les ordonnez et les paginez avec `.sort()`, `.limit()`, et `.skip()`. 

Joignez des collections à l'aide de `$lookup`, et utilisez les étapes du pipeline d'agrégation (`$match`, `$group`, `$sort`, `$lookup`) pour transformer et combiner des documents dans des flux de travail à plusieurs étapes. 

Illustrons ces différences par des exemples. Supposons tout d'abord que nous disposions d'un tableau contenant des étudiants et leurs notes. Dressons la liste des meilleures notes, de la plus élevée à la plus basse.

Dans MySQL, cette opération peut se présenter comme suit :

```
SELECT name, grade
FROM students
WHERE grade > 90
ORDER BY grade DESC;
```
En MQL, vous écririez une requête similaire à celle-ci.

```
db.students.find({ grade: { $gt: 90 } },
                                        { name: 1, grade: 1 })  # only project name and grade
                   .sort({ grade: -1 })
```
Chaque approche tire parti de son modèle sous-jacent - les tableaux pour SQL et les documents pour MQL - pour obtenir une récupération puissante des données.

### Critères de performance

MongoDB offre généralement des performances d'écriture plus rapides que MySQL, car il écrit des documents entiers sans vérifier un schéma rigide. Même lorsque MongoDB utilise la validation de schéma ou des index uniques, il est généralement plus rapide pour les insertions simples, car il écrit l'intégralité du document en une seule fois.

Les performances de lecture varient selon les cas d'utilisation. En raison de son modèle centré sur les documents, MongoDB excelle dans la recherche de documents individuels ou de petits lots par identifiant. My SQL excelle pour les jointures complexes entre plusieurs tableaux ou les grandes agrégations grâce à son optimiseur basé sur les coûts et à ses stratégies d'indexation matures.

## Stratégies de montée en charge de MySQL vs MongoDB

Il existe deux stratégies opposées pour la mise à l'échelle : la stratégie verticale et la stratégie horizontale. La mise à l'échelle verticale consiste à améliorer une machine unique en ajoutant des unités centrales, de la mémoire vive ou des disques plus rapides pour gérer une charge accrue.

Cette approche est simple en ce sens qu'elle ne nécessite pas de grappes ou d'équilibreurs de charge. Cependant, elle est limitée par le matériel et peut être coûteuse. Il y a un seul point de défaillance : Si le serveur tombe en panne, tout s'arrête.

En revanche, la mise à l'échelle horizontale consiste à ajouter des machines pour partager la charge de travail. Il vous permet de vous développer à moindre coût (pas besoin d'un énorme serveur onéreux) et garantit la poursuite des opérations en cas de défaillance d'une machine. Cependant, comme elle nécessite davantage de coordination, comme l'équilibrage de la charge et la synchronisation des données, elle peut entraîner une surcharge du réseau.

### Limites de la mise à l'échelle verticale de MySQL

La mise à l'échelle verticale de MySQL se fait en deux étapes : d'abord, vous mettez à niveau le matériel, puis vous réglez les paramètres de configuration pour utiliser efficacement ces nouvelles ressources. En ce qui concerne le matériel, vous pouvez ajouter des cœurs de processeur, augmenter la mémoire vive et utiliser des composants de stockage et de réseau plus rapides. Une fois que la machine est suffisamment puissante, utilisez ces paramètres de configuration :

- `innodb_buffer_pool_size` . Augmentez la mémoire pour qu'elle corresponde à votre RAM.
- `innodb_log_file_size` ,`innodb_log_buffer_size` : augmentez-les pour mettre en lot davantage d'opérations d'écriture.
- `max_connections` ,`table_open_cache, thread_cache_size` . Augmentez-les pour garder plus de clients et de tableaux au chaud dans la mémoire.

Cependant, l'échelonnement vertical a des limites pratiques. À un moment donné, une machine unique atteindra ses limites en termes de cœurs de CPU ou de mémoire. Il existe un point de rendement décroissant pour l'investissement en matériel ; doubler le prix que vous payez pour les cœurs de l'unité centrale, par exemple, ne double pas la puissance de l'unité centrale.

Les charges de travail à forte intensité d'écriture peuvent rencontrer des problèmes avec des lignes ou des index verrouillés. Une défaillance du serveur peut mettre l'ensemble de la base de données hors ligne.

MySQL peut utiliser et utilise des techniques horizontales, mais cela nécessite une configuration supplémentaire par rapport au modèle de partage intégré de MongoDB.

### La mise à l'échelle horizontale de MongoDB via le sharding

MongoDB utilise une stratégie de mise à l'échelle horizontale, appelée "sharding", qui divise une grande collection en morceaux plus petits sur la base d'une "clé de shard" (par exemple, des plages d'identifiants de documents). Chaque morceau est stocké sur un serveur différent.

Un groupe de serveurs de configuration tient à jour une carte indiquant les zones contenant les différents morceaux. Les routeurs de requêtes utilisent ce répertoire pour envoyer les requêtes des clients aux unités de stockage appropriées. Un équilibreur automatique redistribue les morceaux au fur et à mesure que les données et les schémas de trafic changent, ce qui garantit qu'aucune unité n'est surchargée. Augmentez la capacité en ajoutant des serveurs de stockage.

Cette approche est idéale pour les applications soumises à des charges imprévisibles ou massives, telles que les médias sociaux, les flux vidéo et le commerce électronique pendant les grandes ventes. Les charges de stockage et d'interrogation étant réparties sur de nombreuses machines, le système reste réactif et résilient en cas de pics de trafic.

## Considérations relatives à la sécurité et à la conformité

Pour tout déploiement de base de données, il est essentiel de veiller à ce que les données restent sécurisées et conformes aux réglementations sectorielles. Comparons MySQL et MongoDB en termes d'authentification, de cryptage et de prise en charge de normes telles que HIPAA et GDPR.

### Authentification et cryptage

#### MySQL

MySQL sécurise les données en combinant l'authentification de l'utilisateur, les connexions cryptées et le cryptage au niveau du disque. Chaque utilisateur doit se connecter avec un compte et un mot de passe uniques et ne se voit accorder que les privilèges précis dont il a besoin, soit individuellement, soit en fonction de son rôle. Toutes les données en transit peuvent être protégées par TLS/SSL, imposées comme SSL uniquement, et validées par des certificats de serveur. Les données au repos sont cryptées par le moteur InnoDB. Les clés de chiffrement peuvent être changées sans interruption de service.

#### MongoDB

MongoDB sécurise les données par l'authentification, le cryptage et l'audit. Chaque utilisateur se connecte avec un nom d'utilisateur et un mot de passe uniques (via SCRAM, certificats x.509, Kerberos ou LDAP) et se voit attribuer des autorisations par le biais de rôles, intégrés ou personnalisés. Le trafic en transit est protégé par TLS/SSL, et les données au repos sont protégées par le moteur de stockage de WiredTiger. Des journaux d'audit détaillés enregistrent les tentatives de connexion et les modifications de données afin de faciliter les contrôles de sécurité et la conformité.

### Conformité

