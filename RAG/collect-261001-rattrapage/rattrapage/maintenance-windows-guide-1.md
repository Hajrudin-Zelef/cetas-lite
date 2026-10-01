---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-1
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [1, 181]
sha256: 4df96d1086d0a5e76e3772125c439bce76d647f6f985f1b029f046745cdcd214
---

# Maintenance et exploitation Windows en entreprise

**Guide technique ultra-complet — serveurs & postes de travail**
Versions couvertes : **Windows Server 2019 / 2022 / 2025** — **Windows 11 23H2 / 24H2**
Public : administrateurs systèmes, techniciens d'exploitation, chefs de service.
Langue : français. Ton : direct, opérationnel, sans fioritures.

> **Avertissement** : les commandes de ce guide modifient parfois le système en profondeur
> (registre, BCD, pilotes, stratégies). Testez toujours sur un anneau pilote avant la
> production. Les exemples utilisent des noms fictifs (`srv-fic-01`, `contoso.local`,
> `10.0.0.0/8`) : adaptez-les à votre infrastructure.

---

## Sommaire

1. Pourquoi ce guide et comment l'utiliser
2. Périmètre : versions couvertes et conventions
3. Inventaire : la base de toute exploitation
4. Plan de maintenance : le calendrier annuel
5. Checklist quotidienne (serveurs)
6. Checklist quotidienne (postes)
7. Checklist hebdomadaire
8. Checklist mensuelle
9. Checklist trimestrielle et annuelle
10. Patch Tuesday : principe et calendrier
11. Anneaux de déploiement : test, pilote, production
12. Stratégie de ciblage par criticité
13. WSUS : architecture et dimensionnement
14. WSUS : installation et configuration
15. WSUS : maintenance du serveur (Server Cleanup Wizard)
16. Windows Update for Business : principes
17. WUfB : stratégies de report (quality/feature)
18. Intune vs GPO : pilotage des mises à jour
19. Fenêtres de maintenance et heures d'activité
20. Rollback : désinstaller un correctif
21. DISM : comprendre le magasin de composants
22. DISM : /ScanHealth, /CheckHealth, /RestoreHealth
23. DISM : monter et réparer une image WIM
24. SFC /scannow : mode d'emploi
25. CHKDSK et vérification du système de fichiers
26. Nettoyage disque : cleanmgr et Disk Cleanup
27. Storage Sense : configuration centralisée
28. Dossiers temporaires et caches à purger
29. WinSxS : analyse et nettoyage
30. Journaux IIS et fichiers de trace : rotation
31. Politique d'espace disque : seuils et alertes
32. Event Viewer : prise en main
33. Get-WinEvent : interroger les journaux en PowerShell
34. wevtutil : l'outil en ligne de commande
35. Filtres XPath et vues personnalisées
36. Abonnements d'événements (WEF) et collecteur
37. Journal Système : ce qu'on y cherche
38. Journal Sécurité : audit et alertes
39. Journaux Active Directory : DS, DNS Server, DFSR
40. Journaux DHCP, NPS et autres rôles
41. Rétention et archivage des journaux
42. Planificateur de tâches : concepts
43. Créer une tâche en PowerShell
44. Déclencheurs avancés et conditions
45. Comptes d'exécution : SYSTEM, comptes de service, gMSA
46. gMSA : création et utilisation
47. Audit et hygiène des tâches planifiées
48. Disques : gestion avec PowerShell et MMC
49. Partitionnement : bonnes pratiques
50. Quotas NTFS
51. Optimisation : défragmentation HDD, TRIM SSD
52. Santé des disques : SMART
53. Storage Spaces : concepts et administration
54. Performance : les outils (Gestionnaire des tâches, Resource Monitor, perfmon)
55. Compteurs clés : CPU
56. Compteurs clés : mémoire
57. Compteurs clés : disque
58. Compteurs clés : réseau
59. Analyseur de performances : jeux de collecteurs
60. PAL : Performance Analysis of Logs
61. Scripts de collecte de performances
62. WinRE : accéder à l'environnement de récupération
63. WinRE : activer, désactiver, personnaliser
64. Réparation du démarrage : diagnostic
65. BCD : bcdedit en pratique
66. bootrec : /fixmbr, /fixboot, /rebuildbcd
67. Restauration du système et clichés instantanés
68. Réinitialisation et réinstallation propre
69. BSOD : lire un écran bleu
70. Configurer les fichiers de vidage (dump)
71. WinDbg : installation et prise en main
72. Analyser un minidump : !analyze -v
73. Codes d'arrêt courants : tableau de référence
74. Driver Verifier : traquer un pilote fautif
75. Dépannage réseau : méthode générale
76. ipconfig : lecture et cas d'usage
77. Test-Connection, ping, pathping
78. Tracert et diagnostic de routage
79. Test-NetConnection : le couteau suisse
80. netstat et Get-NetTCPConnection
81. DNS côté client : nslookup et Resolve-DnsName
82. DNS : vider le cache, vérifier les suffixes
83. DHCP côté client : renouvellement et diagnostic
84. Wi-Fi : diagnostic de base
85. Cas pratique 1 : pas d'accès réseau du tout
86. Cas pratique 2 : adresse APIPA 169.254.x.x
87. Cas pratique 3 : DNS ne résout plus
88. Cas pratique 4 : un seul site inaccessible
89. Cas pratique 5 : lenteurs réseau intermittentes
90. Cas pratique 6 : partage SMB inaccessible
91. Cas pratique 7 : imprimante réseau hors ligne
92. Cas pratique 8 : VPN qui ne se connecte plus
93. Cas pratique 9 : port fermé / application injoignable
94. Cas pratique 10 : conflit d'adresse IP
95. Cas pratique 11 : profil réseau Public au lieu de Domaine
96. Cas pratique 12 : débit anormalement bas
97. Cas pratique 13 : RDP impossible vers un serveur
98. Cas pratique 14 : double pile IPv4/IPv6 capricieuse
99. Cas pratique 15 : proxy / PAC qui casse tout
100. AD/GPO côté client : gpupdate et gpresult
101. RSOP et journaux du service de stratégie de groupe
102. nltest : diagnostiquer la relation au domaine
103. Test-ComputerSecureChannel et réinitialisation
104. Temps : w32tm, la cause n°1 des échecs Kerberos
105. Jonction au domaine : procédure et dépannage
106. Monitoring : ce qu'on supervise sur Windows
107. Seuils d'alerte recommandés
108. Zabbix : agent2 et templates Windows
109. Prometheus : windows_exporter
110. Centralisation des journaux (syslog / Graylog / Loki)
111. PowerShell : fondamentaux d'administration
112. PSRemoting et WinRM : mise en place sécurisée
113. Invoke-Command : exécution à distance
114. Sessions persistantes et fichiers distants
115. Inventaire matériel en PowerShell
116. Inventaire logiciels installés
117. Inventaire des correctifs installés
118. Rapport espace disque multi-serveurs
119. Rapport comptes AD expirés / inactifs
120. Rapport certificats qui expirent
121. Exécution planifiée des scripts de rapport
122. PRA/PCA : RTO, RPO et scénarios
123. Stratégie de sauvegarde Windows
124. Windows Server Backup en pratique
125. VSS : clichés instantanés de volumes
126. Sauvegarde des contrôleurs de domaine (état du système)
127. Tests de restauration : procédure
128. Documentation d'exploitation : le DTI
129. Procédures : standard de rédaction
130. Gestion des changements
131. Passation et continuité (angle chef de service)
132. 20 erreurs classiques d'exploitation Windows
133. Pense-bête des commandes PowerShell
134. Pense-bête des commandes CMD / classiques
135. Quiz : 10 questions
136. Réponses du quiz
137. Glossaire
138. Pour aller plus loin

---

## 1. Pourquoi ce guide et comment l'utiliser

Ce guide couvre le **cycle de vie opérationnel** d'un parc Windows : du patch mensuel au
dépannage d'un écran bleu, en passant par la supervision et la reprise d'activité.
Il s'adresse au technicien qui intervient comme au chef de service qui organise.

**Mode d'emploi :**

- En intervention : utilisez le sommaire et le pense-bête (§133-134).
- En routine : appliquez les checklists (§5-9) telles quelles, adaptez les seuils.
- En projet : les sections WSUS/WUfB (§13-18), monitoring (§106-110) et PRA (§122-127)
  servent de cahier des charges.
- Chaque commande PowerShell est autonome : copiable, modifiable, réutilisable dans
  vos scripts. Les noms de machines, domaines et IP sont fictifs.

**Règle d'or du chef de service** : on n'exploite bien que ce qu'on connaît.
Sans inventaire à jour (§3), sans supervision (§106) et sans documentation (§128),
toute maintenance est du bricolage. Ce guide commence donc par l'organisation,
pas par la technique.

---

## 2. Périmètre : versions couvertes et conventions

