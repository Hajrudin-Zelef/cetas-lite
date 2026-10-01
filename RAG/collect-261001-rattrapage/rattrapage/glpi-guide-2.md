---
id: collect-261001-rattrapage/rattrapage/glpi-guide-2
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "apache", "open source"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [14, 98]
sha256: dcbe824c15cd4a463778e3ac871472d84f9b563859869ed2241f01f3ab7f0c3e
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

1. [À quoi sert GLPI (et pourquoi c'est le cœur de votre organisation)](#1-à-quoi-sert-glpi-et-pourquoi-cest-le-cœur-de-votre-organisation)
2. [Concepts fondamentaux : ITSM, CMDB, entités, profils](#2-concepts-fondamentaux--itsm-cmdb-entités-profils)
3. [Architecture technique de GLPI 10.x](#3-architecture-technique-de-glpi-10x)
4. [Prérequis serveur détaillés (Debian/Ubuntu)](#4-prérequis-serveur-détaillés-debianubuntu)
5. [Installation pas à pas — pile LAMP (Apache + PHP 8 + MariaDB)](#5-installation-pas-à-pas--pile-lamp-apache--php-8--mariadb)
6. [Installation pas à pas — variante nginx + PHP-FPM](#6-installation-pas-à-pas--variante-nginx--php-fpm)
7. [Assistant d'installation web : écran par écran](#7-assistant-dinstallation-web--écran-par-écran)
8. [Configuration post-installation : vérifications critiques](#8-configuration-post-installation--vérifications-critiques)
9. [Tâches planifiées (cron) : le cœur battant de GLPI](#9-tâches-planifiées-cron--le-cœur-battant-de-glpi)
10. [Entités : l'arborescence multi-sites](#10-entités--larborescence-multi-sites)
11. [Profils et habilitations : qui fait quoi](#11-profils-et-habilitations--qui-fait-quoi)
12. [Utilisateurs : création, import LDAP/AD, règles](#12-utilisateurs--création-import-ldapad-règles)
13. [Groupes : équipes, escalades, astreintes](#13-groupes--équipes-escalades-astreintes)
14. [Notifications par courriel : configuration complète](#14-notifications-par-courriel--configuration-complète)
15. [Le ticket : cycle de vie complet](#15-le-ticket--cycle-de-vie-complet)
16. [Catégories de tickets : la taxonomie du SAV](#16-catégories-de-tickets--la-taxonomie-du-sav)
17. [Urgence, impact, priorité : la matrice](#17-urgence-impact-priorité--la-matrice)
18. [Gabarits de tickets et champs obligatoires](#18-gabarits-de-tickets-et-champs-obligatoires)
19. [Tickets récurrents : la maintenance planifiée](#19-tickets-récurrents--la-maintenance-planifiée)
20. [Actions automatiques et règles métier sur tickets](#20-actions-automatiques-et-règles-métier-sur-tickets)
21. [SLA : niveaux de service, escalades, pénalités](#21-sla--niveaux-de-service-escalades-pénalités)
22. [Suivis, tâches, solutions : bien documenter](#22-suivis-tâches-solutions--bien-documenter)
23. [Sondages de satisfaction](#23-sondages-de-satisfaction)
24. [Base de connaissances : rédiger des articles utiles](#24-base-de-connaissances--rédiger-des-articles-utiles)
25. [Le parc : vision d'ensemble (CMDB)](#25-le-parc--vision-densemble-cmdb)
26. [Ordinateurs : fiches, composants, liaisons](#26-ordinateurs--fiches-composants-liaisons)
27. [Moniteurs, imprimantes et périphériques](#27-moniteurs-imprimantes-et-périphériques)
28. [Fiches copieurs : le métier (marque, modèle, série, compteur, contrat)](#28-fiches-copieurs--le-métier-marque-modèle-série-compteur-contrat)
29. [Périphériques réseau : switchs, routeurs, bornes Wi-Fi, onduleurs](#29-périphériques-réseau--switchs-routeurs-bornes-wi-fi-onduleurs)
30. [Cartouches et consommables : gérer le stock de toner](#30-cartouches-et-consommables--gérer-le-stock-de-toner)
31. [Logiciels, licences et versions](#31-logiciels-licences-et-versions)
32. [Réservations de matériel](#32-réservations-de-matériel)
33. [FusionInventory : présentation et architecture](#33-fusioninventory--présentation-et-architecture)
34. [Installation de l'agent FusionInventory](#34-installation-de-lagent-fusioninventory)
35. [Inventaire automatique : réseau et SNMP](#35-inventaire-automatique--réseau-et-snmp)
36. [Règles d'import et de déduplication](#36-règles-dimport-et-de-déduplication)
37. [Gestion des contrats et des fournisseurs](#37-gestion-des-contrats-et-des-fournisseurs)
38. [Budgets et lignes budgétaires](#38-budgets-et-lignes-budgétaires)
39. [Projets et tâches : piloter les chantiers](#39-projets-et-tâches--piloter-les-chantiers)
40. [Rapports et statistiques : le tableau de bord du chef de service](#40-rapports-et-statistiques--le-tableau-de-bord-du-chef-de-service)
41. [Recherche, filtres et vues enregistrées](#41-recherche-filtres-et-vues-enregistrées)
42. [Sécurité — comptes et mots de passe](#42-sécurité--comptes-et-mots-de-passe)
43. [Sécurité — durcissement des profils](#43-sécurité--durcissement-des-profils)
44. [Sécurité — HTTPS/TLS de bout en bout](#44-sécurité--httpstls-de-bout-en-bout)
45. [Sécurité — durcissement serveur et GLPI](#45-sécurité--durcissement-serveur-et-glpi)
46. [Supervision de GLPI : ce qu'il faut surveiller](#46-supervision-de-glpi--ce-quil-faut-surveiller)
47. [Sauvegarde : stratégie complète (fichiers + base)](#47-sauvegarde--stratégie-complète-fichiers--base)
48. [Sauvegarde automatisée : scripts prêts à l'emploi](#48-sauvegarde-automatisée--scripts-prêts-à-lemploi)
49. [Restauration : procédure testée pas à pas](#49-restauration--procédure-testée-pas-à-pas)
50. [Mise à jour de GLPI : méthode sans casse](#50-mise-à-jour-de-glpi--méthode-sans-casse)
51. [Plugins : lesquels installer, lesquels éviter](#51-plugins--lesquels-installer-lesquels-éviter)
52. [API REST : authentification et premiers appels](#52-api-rest--authentification-et-premiers-appels)
53. [API REST : exemples concrets (tickets, parc, compteurs)](#53-api-rest--exemples-concrets-tickets-parc-compteurs)
54. [Dépannage — page blanche / erreur 500](#54-dépannage--page-blanche--erreur-500)
55. [Dépannage — le cron ne tourne pas](#55-dépannage--le-cron-ne-tourne-pas)
56. [Dépannage — les courriels ne partent pas](#56-dépannage--les-courriels-ne-partent-pas)
57. [Dépannage — 7 autres cas concrets](#57-dépannage--7-autres-cas-concrets)
58. [10 erreurs classiques (et comment les éviter)](#58-10-erreurs-classiques-et-comment-les-éviter)
59. [Cas pratique 1 — organiser le SAV d'une entreprise de maintenance copieurs](#59-cas-pratique-1--organiser-le-sav-dune-entreprise-de-maintenance-copieurs)
60. [Cas pratique 2 — tickets d'intervention copieurs de A à Z](#60-cas-pratique-2--tickets-dintervention-copieurs-de-a-à-z)
61. [Cas pratique 3 — planning des techniciens](#61-cas-pratique-3--planning-des-techniciens)
62. [Cas pratique 4 — stock pièces et consommables](#62-cas-pratique-4--stock-pièces-et-consommables)
63. [Cas pratique 5 — migration depuis Excel/ancien outil](#63-cas-pratique-5--migration-depuis-excelancien-outil)
64. [Cas pratique 6 — revue hebdomadaire du chef de service](#64-cas-pratique-6--revue-hebdomadaire-du-chef-de-service)
65. [Cas pratique 7 — intégrer les guides existants dans la base de connaissances](#65-cas-pratique-7--intégrer-les-guides-existants-dans-la-base-de-connaissances)
66. [Pense-bête de poche (une page)](#66-pense-bête-de-poche-une-page)
67. [Glossaire](#67-glossaire)
68. [Quiz — 10 questions + réponses](#68-quiz--10-questions--réponses)
69. [Pour aller plus loin](#69-pour-aller-plus-loin)
70. [Checklist de mise en production](#70-checklist-de-mise-en-production)
71. [Annexe A — paramètres PHP recommandés](#71-annexe-a--paramètres-php-recommandés)
72. [Annexe B — configuration Apache de référence](#72-annexe-b--configuration-apache-de-référence)
73. [Annexe C — configuration nginx de référence](#73-annexe-c--configuration-nginx-de-référence)
74. [Annexe D — modèle de gabarit « Intervention copieur »](#74-annexe-d--modèle-de-gabarit--intervention-copieur-)
75. [Annexe E — modèle de rapport hebdomadaire SAV](#75-annexe-e--modèle-de-rapport-hebdomadaire-sav)

---

# PARTIE I — CONCEPTS ET ARCHITECTURE

## 1. À quoi sert GLPI (et pourquoi c'est le cœur de votre organisation)

GLPI (Gestionnaire Libre de Parc Informatique) est une application web open source
(licence GPL) qui remplit **deux fonctions majeures** :

