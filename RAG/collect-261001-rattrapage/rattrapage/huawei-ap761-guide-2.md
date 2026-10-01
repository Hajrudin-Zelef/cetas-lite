---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-2
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [38, 185]
sha256: 4ffffa23c636e950e995266568b59a4a48593bd912768ad23e3dbc01fed25274
---

# Guide ultra-complet — Huawei eKit AP761

1. Ce que ce guide couvre (et ne couvre pas)
2. Positionnement de l'AP761 dans la gamme eKit
3. Contenu de la boîte et prérequis chantier
4. Fiche technique constructeur — tableau de synthèse
5. Hardware : boîtier, dimensions, robustesse
6. Hardware : ports filaires (GE RJ45 + SFP, mode combo)
7. Hardware : alimentation PoE — 802.3at/af, 17.7 W
8. Hardware : LED d'état et bouton reset
9. Hardware : antennes intégrées directionnelles
10. Hardware : BLE 5.2 et USB IoT
11. Les 3 modes de fonctionnement : Fat, Fit, Cloud
12. Wi-Fi 6 (802.11ax) : ce que l'AP761 fait vraiment
13. Rétrocompatibilité Wi-Fi 5/4/anciens clients
14. Le Wi-Fi 7 (802.11be) expliqué simplement
15. MLO — Multi-Link Operation, en clair
16. Canaux 320 MHz : pourquoi l'outdoor eKit n'en a pas
17. 4096-QAM (4K-QAM) : +20 % de débit, à quel prix
18. Preamble puncturing et multi-RU
19. 6 GHz et réglementation : le point pour l'Europe
20. Ce que l'AP761 ne fait PAS (tableau Wi-Fi 6 vs Wi-Fi 7)
21. La gamme eKit Wi-Fi 7 réelle : AP771, AP772E, AP371, AP673
22. Installation physique : où poser un AP extérieur
23. Montage mural et sur mât — procédure pas à pas
24. Câblage : choix du câble, distance, étanchéité
25. PoE : budget, injecteurs, switches — calculs
26. Protection foudre et mise à la terre
27. Checklist pré-installation chantier
28. Première mise en route : les 3 chemins
29. Onboarding via l'app HUAWEI eKit (QR code)
30. Mode cloud : créer un site, adopter l'AP
31. Mode Fat : accès web local et CLI
32. Mode Fit : raccordement à un AC (WAC)
33. Changer de mode : fat ↔ fit ↔ cloud
34. Configuration initiale minimale (recette)
35. Sauvegarde de la configuration initiale
36. Configuration radio : concepts de base
37. Plan de canaux 2.4 GHz — méthode terrain
38. Plan de canaux 5 GHz — méthode terrain
39. Largeur de canal : 20/40/80 MHz, que choisir
40. Puissance d'émission : régler sans bourriner
41. DFS en 5 GHz : ce qu'il faut savoir
42. Band steering et préférence de bande
43. SSID : stratégie de nommage et nombre
44. Sécurité : panorama WPA/WPA2/WPA3 sur l'AP761
45. WPA2-PSK : configuration pas à pas
46. WPA3-SAE : configuration et limites
47. Mode transition WPA2/WPA3 : pour qui, pourquoi
48. 802.1X / WPA2-Enterprise avec RADIUS
49. Portail captif (Portal) pour invités
50. Filtrage MAC : utile ou placebo ?
51. 802.11w (PMF) : protection des trames de management
52. VLAN : segmentation par SSID
53. Isolation client (STA isolation) dans un VLAN
54. QoS / WMM : prioriser la voix et la vidéo
55. ACL et DHCP snooping / DAI / IPSG
56. Itinérance : 802.11k, 802.11v, 802.11r
57. Fast roaming en pratique sur l'AP761
58. Planification : combien d'AP pour quelle surface
59. Règles de pouce densité / débit par usage
60. Site survey simplifié sans outils pro — méthode terrain
61. Lecture d'un relevé : RSSI, SNR, bruit
62. Erreurs classiques de placement d'AP extérieur
63. WIDS/wIPS : détection d'intrusion sans fil
64. Rogue AP : détecter, classifier, contenir
65. Analyse de spectre (mode Fit uniquement)
66. Supervision via la plateforme cloud eKit
67. Indicateurs Wi-Fi à surveiller (tableau)
68. SNMP : v1/v2c/v3 en mode Fat
69. Syslog : centraliser les journaux
70. NTP : pourquoi c'est critique
71. Alertes et seuils : que configurer
72. Sauvegarde de configuration : méthodes
73. Restauration de configuration
74. Firmware : cycle de vie et bonnes pratiques
75. Mise à jour firmware — procédure pas à pas
76. Rollback : revenir en arrière proprement
77. Dépannage : méthode générale (les 5 couches)
78. Cas 1 : l'AP ne s'allume pas / pas de PoE
79. Cas 2 : l'AP démarre mais n'apparaît pas dans le cloud
80. Cas 3 : SSID visible mais pas d'IP (DHCP)
81. Cas 4 : débit très en-deçà des attentes
82. Cas 5 : clients 5 GHz qui restent accrochés au 2.4 GHz
83. Cas 6 : coupures lors des déplacements (roaming)
84. Cas 7 : interférences 2.4 GHz (micro-ondes, Bluetooth, caméras)
85. Cas 8 : radar DFS — l'AP change de canal tout seul
86. Cas 9 : PoE insuffisant — fonctions restreintes en 802.3af
87. Cas 10 : SFP branché mais pas de lien
88. Cas 11 : WPA3 activé, certains clients ne se connectent plus
89. Cas 12 : portail captif qui ne s'affiche pas
90. Cas 13 : un client consomme tout l'airtime (slow client)
91. Cas 14 : l'AP surchauffe / redémarre en plein soleil
92. Cas 15 : rogue AP détecté — que faire concrètement
93. Cas 16 : après mise à jour firmware, comportements bizarres
94. Cas 17 : itinérance entre AP761 et AP indoor
95. Cas 18 : je veux du Wi-Fi 7 — que faire avec mes AP761 ?
96. Comparatif AP761 vs AP361 — tableau décisionnel
97. Quand choisir l'AP361 plutôt que l'AP761
98. Quand choisir l'AP761 plutôt que l'AP361
99. Et le Wi-Fi 7 outdoor ? AP771 vs AP772E
100. Tableau décisionnel global : quel AP eKit pour quel besoin
101. Maintenance préventive : planning annuel
102. Contrôle visuel et nettoyage
103. Contrôle électrique et PoE
104. Contrôle radio (audit annuel)
105. Gestion du cycle de vie et fin de vie
106. Durcissement : checklist sécurité
107. Durcissement : management (SSH, SNMP, web)
108. Durcissement : Wi-Fi (chiffrement, PMF, WIDS)
109. Pense-bête de poche — page 1 : specs
110. Pense-bête de poche — page 2 : CLI express
111. Pense-bête de poche — page 3 : dépannage express
112. Glossaire (A–Z)
113. Quiz : 10 questions
114. Quiz : réponses commentées
115. Cas pratique 1 : cour d'école / cour d'entreprise
116. Cas pratique 2 : parking et contrôle d'accès
117. Cas pratique 3 : entrepôt / zone logistique
118. Cas pratique 4 : terrasse d'hôtel / camping
119. Cas pratique 5 : relais entre deux bâtiments (point à point)
120. Pour aller plus loin : docs, outils, formations

---

## 1. Ce que ce guide couvre (et ne couvre pas)

**Ce guide couvre :**
- Le hardware réel de l'AP761, vérifié sur datasheet constructeur (sept. 2026).
- L'installation physique d'un AP extérieur (montage, PoE, foudre) — ton métier.
- La configuration via app eKit / cloud / web / CLI VRP.
- La radio Wi-Fi 6 : canaux, puissance, sécurité, VLAN, QoS, roaming.
- La supervision, la maintenance, le durcissement.
- 18 cas de dépannage terrain commentés.
- Un volet Wi-Fi 7 honnête : théorie + vrais modèles eKit concernés.

**Ce guide ne couvre PAS :**
- Le Wi-Fi 7 *sur* l'AP761 — parce qu'il n'en fait pas. Le chapitre 18 explique la migration.
- Les détails d'AirEngine (gamme entreprise) : l'AP761 est eKit (PME), la logique est simplifiée.
- Les procédures internes à ton entreprise (consignation électrique, habilitations) : elles s'appliquent en plus, pas à la place.

**Convention d'écriture :**
- `[valeur constructeur]` = valeur tirée du datasheet Huawei eKitEngine AP761.
- `« à vérifier sur la fiche du modèle exact »` = valeur non vérifiée au 27/09/2026. Ne jamais la prendre pour argent comptant sur un appel d'offres.
- Les commandes CLI sont **syntaxiquement plausibles pour VRP Huawei** (famille des AP eKit/CloudEngine). Le nom exact des vues et commandes peut varier selon la version logicielle : toujours valider avec `?` en CLI avant d'appliquer en production.

## 2. Positionnement de l'AP761 dans la gamme eKit

Huawei eKit = la gamme PME de Huawei (ex-« eKitEngine »), sans licence, gérée par app mobile ou cloud gratuit. Dans cette gamme :

