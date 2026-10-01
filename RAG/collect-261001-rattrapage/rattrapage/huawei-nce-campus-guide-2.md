---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-2
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [12, 158]
sha256: bcf11c2a92ac91ab1ee7b43ba515414257bed3a18b52b926c897dd7398ef3a52
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

1. Ce qu'est iMaster NCE-Campus — définition officielle
2. La famille iMaster NCE : NCE-Campus, NCE-Fabric, NCE-WAN, NCE-IP
3. Pourquoi Huawei pousse NCE côté campus — la logique produit
4. ADN : Autonomous Driving Network — ce que ça veut dire concrètement
5. Relation avec eSight — le positionnement officiel de Huawei
6. eSight reste-t-il valable ? — réponse honnête
7. Fin de vie eSight — ce qui est connu et ce qui ne l'est pas (à vérifier)
8. iMaster NCE-Campus vs Huawei eKit — la question du petit site
9. Qui a besoin de NCE-Campus — profils types
10. Qui n'en a PAS besoin — quand eSight ou eKit suffisent
11. Architecture globale — vue à 10 000 mètres
12. Les trois couches logiques : management, control, analysis
13. Le contrôleur SDN au cœur du système
14. Southbound interfaces : comment NCE parle aux équipements
15. NETCONF/YANG — le protocole de configuration structurée
16. SNMP — la télémétrie et le legacy
17. CAPWAP — la gestion des AP AirEngine
18. HTTP/2, HTTPS et autres canaux southbound
19. Base de données et stockage — que garde NCE ?
20. iMaster NCE-CampusInsight — l'analyseur (composant distinct)
21. Le composant d'authentification — déploiement distribué
22. Haute disponibilité — modes de déploiement
23. Dimensionnement : petit site (ordre de grandeur)
24. Dimensionnement : site moyen (ordre de grandeur)
25. Dimensionnement : grand campus / multi-sites (ordre de grandeur)
26. Modèle de licence — principe général (device-day)
27. Licence plateforme vs licence équipement
28. Licences par équipement géré — comment elles se consomment
29. Période de grâce et comportement en cas de dépassement
30. Mode MSP — multi-tenant et licence commune
31. iMaster NCE-CampusInsight — souscriptions associées
32. Licence de 60 devices incluse dans eSight — comparaison de logique
33. Coût — ordre de grandeur et pourquoi on ne donne pas de prix ici
34. Checklist avant achat de licences
35. Architecture matérielle requise — serveur / VM
36. Système d'exploitation supporté
37. Prérequis réseau (ports, flux, DNS, NTP)
38. Prérequis d'installation — checklist complète
39. Obtenir le package d'installation
40. Installation pas à pas — grandes étapes
41. Premier login — comptes par défaut et politique
42. Configuration initiale — assistant de démarrage
43. Création de l'organisation : tenant, régions, sites
44. Intégration DNS / NTP / certificats
45. Sauvegarde initiale — avant toute chose
46. Checklist post-installation
47. Échecs d'installation fréquents et remèdes
48. Mise à jour du contrôleur lui-même
49. Notion de site dans NCE-Campus
50. Les équipements gérés : S310, AP361, AP761, AR720, USG6000
51. ZTP — principe du zero-touch provisioning
52. ZTP via DHCP (option 148/option personnalisée) — détail
53. ZTP via e-mail d'activation — détail
54. ZTP via scan (QR / code) — détail
55. Onboarding manuel — quand le ZTP n'est pas possible
56. Ajouter un S310 — procédure détaillée
57. Ajouter un AP361 — procédure détaillée
58. Ajouter un AP761 — procédure détaillée
59. Ajouter un AR720 — procédure détaillée
60. Vérifier qu'un équipement est bien géré — checklist
61. États des équipements dans NCE (normal, alarme, hors ligne, non enregistré)
62. Templates de configuration — concepts
63. Variables dans les templates — paramétrage par site/équipement
64. VLAN — conception et déploiement via NCE
65. Politiques d'accès filaire — 802.1X : principes
66. 802.1X : rôles (supplicant, authentificateur, serveur)
67. 802.1X : méthodes EAP courantes
68. MAC-auth et portail captif — les alternatives
69. Autorisation dynamique : VLAN, ACL, QoS par utilisateur
70. QoS filaire — modèles et déploiement
71. Gestion WLAN centralisée — principes
72. Le modèle WAC + AP : comment NCE pilote le Wi-Fi
73. Profils radio — conception (canaux, puissance, bandes)
74. SSID — conception et bonnes pratiques
75. Planification radio : RRM, DCA, TPC
76. Roaming — 802.11k/v/r et pratiques Huawei
77. Sécurité WLAN : WPA2/WPA3, 802.1X Wi-Fi, PPSK
78. NCE vs eKit pour le WLAN — quand NCE apporte un plus
79. Invités et BYOD — portail captif via NCE
80. IoT sur le WLAN — segmentation
81. Automation — vue d'ensemble
82. ZTP détaillé — séquence complète d'un onboarding automatique
83. Templates et variables — guide pratique
84. Déploiement en masse — vagues et fenêtres de maintenance
85. Mise à jour firmware en masse — procédure
86. Politique de mise à jour : manuelle vs automatique
87. Rollback — revenir en arrière après un déploiement raté
88. API northbound RESTful — principes
89. Authentification API : tokens (POST /controller/v2/tokens)
90. Cas d'usage API : topologie, métriques, alarmes
91. Intégration avec des outils tiers (Zabbix, GLPI, scripts)
92. Assurance et supervision — vue d'ensemble
93. Monitoring temps réel — tableaux de bord
94. Cartographie réseau (digital map) — topologie
95. Détection d'anomalies — ce que fait l'IA/ML
96. Analyse de cause racine (root cause analysis)
97. Rapports — types et planification
98. Alertes — configuration mail/SMS
99. SLA management — NQA et mesures actives
100. Maintenance des équipements depuis NCE
101. Comptes et rôles — modèle RBAC
102. Rôles prédéfinis et rôles personnalisés
103. Authentification des administrateurs — locale, RADIUS, LDAP/AD
104. Journal d'audit — traçabilité des actions
105. Durcissement du contrôleur — checklist
106. Certificats — gestion et renouvellement
107. Sauvegardes NCE — stratégie
108. Restauration — procédure et tests
109. Sécurité des flux southbound
110. Stratégie de migration eSight → NCE — vue d'ensemble
111. Coexistence eSight + NCE — mode recommandé
112. Migration par site — méthode pas à pas
113. Ce qui est repris de eSight — inventaire
114. Ce qui N'est PAS repris — limites honnêtes
115. Checklist de migration complète
116. Risques de la migration et mitigations
117. Planning type de migration (exemple 3 sites)
118. Retour en arrière — plan B
119. Comparaison eSight vs NCE-Campus — tableau général
120. Comparaison détaillée par fonction
121. Coût comparé — TCO sur 5 ans (méthode)
122. Complexité comparée
123. Cas d'usage : quand choisir quoi — arbre de décision
124. Cas pratique 1 — L'onboarding ZTP qui échoue (DHCP)
125. Cas pratique 2 — Un AP361 qui ne remonte pas dans NCE
126. Cas pratique 3 — Un AP761 extérieur qui reste hors ligne
127. Cas pratique 4 — Un S310 qui boucle en enregistrement
128. Cas pratique 5 — Un AR720 qui ne télécharge pas sa config
129. Cas pratique 6 — Un template qui casse la production
130. Cas pratique 7 — Licence dépassée : que se passe-t-il ?
131. Cas pratique 8 — Mise à jour firmware massive qui échoue à mi-parcours
132. Cas pratique 9 — Roaming Wi-Fi dégradé après migration NCE
133. Cas pratique 10 — 802.1X qui rejette tous les utilisateurs un lundi matin
134. Cas pratique 11 — Le contrôleur NCE ne répond plus
135. Cas pratique 12 — Faux positifs d'alarmes après un orage
136. Cas pratique 13 — Un site distant perd le lien vers NCE
137. Cas pratique 14 — Conflit entre config manuelle CLI et NCE
138. Cas pratique 15 — Certificat expiré : tout s'arrête
139. Cas pratique 16 — Restauration après sinistre du serveur NCE
140. Cas pratique 17 — Montée en charge : NCE sature
141. Cas pratique 18 — Intégration RADIUS/LDAP qui casse l'authentification admin
142. Cas pratique 19 — eKit et NCE se disputent un AP
143. Cas pratique 20 — Audit de sécurité : durcir en urgence
144. Cas pratique 21 — Migration eSight → NCE d'un site pilote
145. Cas pratique 22 — Rapport mensuel pour la direction
146. Cas pratique 23 — Supervision croisée NCE + Zabbix
147. Cas pratique 24 — PPSK pour les invités d'un hôtel
