---
id: collect-261001-rattrapage/rattrapage/grafana-guide-1
title: "Guide Grafana — Dashboards, visualisation et alerting"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-rattrapage/grafana_guide.md
source_anchor: ""
source_lines: [1, 163]
sha256: 1032e8a96f9349883768398ddd5612c33a5d93bfa65c3052e354cbebe15b827a
---

# Guide Grafana — Dashboards, visualisation et alerting

> **Public :** chefs de service systèmes, sysadmins, équipes d'exploitation.
> **Objectif :** installer, configurer, exploiter et superviser Grafana en production :
> dashboards lisibles, alerting fiable, provisioning versionné, sauvegardes qui restaurent.
> **Versions couvertes :** Grafana 10.x / 11.x (OSS), sur Debian 12 / Ubuntu 22.04 / 24.04.
> **Convention :** les commandes sont à exécuter en root sauf mention contraire.
> Aucun mot de passe réel dans ce guide : remplacez les valeurs `<A_COMPLETER>`.

---

## Sommaire

1. Introduction et périmètre du guide
2. Concepts : datasources
3. Concepts : dashboards, panels, variables
4. Concepts : alerting unifié
5. Architecture de Grafana
6. Installation sur Debian/Ubuntu (dépôt officiel)
7. Installation via Docker (aperçu)
8. Premier lancement et première connexion
9. Tour d'horizon de l'interface
10. Configuration `grafana.ini` — serveur et chemins
11. Configuration — base de données interne
12. Configuration — journalisation
13. Configuration — sécurité applicative (secret, cookies, sessions)
14. Configuration — HTTPS/TLS natif
15. Configuration — SMTP (notifications e-mail)
16. Datasource Prometheus — ajout et réglages
17. PromQL de base pour vos dashboards
18. PromQL avancé (rate, histogrammes, label_replace)
19. Datasource Loki — ajout et réglages
20. LogQL de base
21. LogQL avancé (parsing, filtres, métriques issues des logs)
22. Datasource Zabbix (plugin)
23. Datasource MySQL — métriques métier
24. Datasource PostgreSQL — métriques métier
25. Datasource InfluxDB (introduction)
26. Datasources : tests et bonnes pratiques
27. Panels : time series
28. Panels : gauge (jauge)
29. Panels : stat (grandes valeurs)
30. Panels : table
31. Panels : heatmap
32. Panels : bar chart, pie chart, state timeline, status history
33. Panels : text, dashboard list, alert list
34. Options de panel : unités, min/max, seuils, couleurs
35. Overrides de champs (field overrides)
36. Variables et templates — le principe
37. Variables : les 7 types en détail
38. Chaînage de variables, regex et multi-sélection
39. Transformations (merge, join, filter, calculate field…)
40. Annotations (événements superposés aux courbes)
41. Liens et drill-down (data links, dashboard links)
42. Alerting unifié : concepts (règles, instances, états)
43. Créer une règle d'alerte — exemple commenté pas à pas
44. Expressions d'alerte (reduce, math, resample, classic condition)
45. Contact points : e-mail, Slack, webhook, Telegram
46. Notification policies : routage, regroupement, répétition
47. Silences et mute timings (plages de maintenance)
48. Bonnes pratiques d'alerting (ce qui évite la fatigue d'alerte)
49. Provisioning as code : le principe
50. Provisioning des datasources (YAML)
51. Provisioning des dashboards (YAML)
52. Provisioning de l'alerting (contact points, policies, règles)
53. Workflow Git pour le provisioning
54. Organisations, équipes et dossiers
55. Rôles : Viewer, Editor, Admin, Server Admin
56. RBAC fin (permissions granulaires)
57. Authentification : basique, LDAP, OAuth (aperçu)
58. Service accounts et tokens d'API
59. HTTPS en production via reverse proxy Nginx
60. Durcissement sécurité — checklist
61. Superviser Grafana lui-même (métriques internes)
62. Sauvegarde : base SQLite
63. Sauvegarde : base PostgreSQL
64. Sauvegarde : provisioning, plugins, stratégie 3-2-1
65. Restauration — procédure complète
66. Mise à jour de Grafana — procédure sans casse
67. Usage quotidien — routine de l'exploitant
68. Bonnes pratiques de dashboarding (lisibilité)
69. Exemple : dashboard supervision serveurs (node_exporter)
70. Exemple : dashboard supervision réseau
71. Exemple : dashboard supervision onduleur (lien métier)
72. Kiosque / TV murale (playlists, kiosk mode)
73. Dépannage : la datasource ne répond pas
74. Dépannage : les alertes ne partent pas
75. Dépannage : 10 cas concrets résolus
76. Les 10 erreurs classiques (et comment les éviter)
77. Cas pratiques commentés (4 scénarios d'exploitation)
78. Pense-bête de poche (commandes, requêtes, raccourcis)
79. Glossaire
80. Quiz — 10 questions + réponses
81. Pour aller plus loin
82. Annexe A — `grafana.ini` de référence commenté
83. Annexe B — checklist de mise en production

---

## 1. Introduction et périmètre du guide

Grafana est un outil open source de **visualisation et d'observabilité** : il ne collecte
pas les données lui-même, il les interroge là où elles vivent (Prometheus, Loki,
Zabbix, MySQL, InfluxDB…) et les présente sous forme de dashboards, avec un
moteur d'alerting unifié.

Ce guide couvre la chaîne complète, de l'installation sur Debian/Ubuntu jusqu'à
l'exploitation quotidienne en production :

- installer proprement depuis le dépôt officiel (pas de binaire téléchargé à la main) ;
- configurer `grafana.ini` de façon sûre et reproductible ;
- brancher les datasources les plus courantes en entreprise (Prometheus, Loki,
  Zabbix, MySQL/PostgreSQL, InfluxDB) ;
- construire des dashboards **lisibles** (panels, variables, transformations) ;
- mettre en place un **alerting fiable** (règles, contact points, routage, silences) ;
- gérer Grafana **as code** (provisioning YAML versionné dans Git) ;
- sécuriser (rôles, HTTPS, durcissement), sauvegarder, mettre à jour, dépanner.

**Ce que ce guide ne couvre pas :** le développement de plugins maison, Grafana
Cloud (offre SaaS), ni l'administration avancée des backends eux-mêmes
(cluster Prometheus Thanos/Mimir, cluster Loki) — ceux-ci sont mentionnés en
piste dans la section 81.

**Pré-requis :** un serveur Debian 12 ou Ubuntu 22.04/24.04, 2 vCPU / 4 Go RAM
minimum pour un usage d'équipe, un accès réseau vers vos datasources, et des
bases de Linux (systemd, apt, nginx).

---

## 2. Concepts : datasources

Une **datasource** est une connexion nommée vers un backend de données.
Grafana ne stocke pas vos métriques : chaque panel exécute une requête vers la
datasource au moment de l'affichage.

| Datasource | Type de données | Langage de requête | Usage typique |
|---|---|---|---|
| Prometheus | Métriques (séries temporelles) | PromQL | Supervision serveurs, conteneurs, applicatif |
| Loki | Logs | LogQL | Corrélation logs ↔ métriques |
| Zabbix (plugin) | Métriques + événements | Requêtes Zabbix API | Parcs déjà sous Zabbix |
| MySQL / PostgreSQL | Tables relationnelles | SQL | Métriques métier (commandes, tickets, SLA) |
| InfluxDB | Métriques (séries temporelles) | InfluxQL / Flux | IoT, historique existant |
| Elasticsearch | Logs / documents | Lucene / PPL | Logs centralisés existants |
| TestData | Données factices | — | Maquettes, formation |

Points clés à retenir :

- **Une datasource = un backend + des droits.** Grafana interroge le backend avec
  les identifiants configurés dans la datasource, pas avec ceux de l'utilisateur
  connecté (sauf configuration spécifique).
- **Le mode d'accès** : `Server` (le backend Grafana interroge la datasource —
  recommandé, les secrets restent côté serveur) ou `Browser` (le navigateur de
  l'utilisateur interroge directement — à éviter sauf cas particulier, car cela
  expose l'URL et potentiellement les identifiants).
- **La datasource par défaut** (`default: true`) est pré-sélectionnée dans les
  nouveaux panels.
- Une datasource peut être **provisionnée** (définie en YAML, versionnée dans
  Git) : c'est la pratique recommandée en production (voir sections 49-50).

---

## 3. Concepts : dashboards, panels, variables

