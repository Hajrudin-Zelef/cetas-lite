---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-25
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [2573, 2668]
sha256: c17db1d36f57c659fabcf6af7bb6f08ccf12f1e64626e74b88899f50c97a7a87
---

# Guide ultra-complet — Huawei eKit AP761

**RADIUS** : serveur d'authentification pour le 802.1X. Chap. 48.

**Radio** : ici : l'émetteur-récepteur Wi-Fi (l'AP761 en a 2 : 2.4 GHz et 5 GHz).

**Retries (réessais)** : trames réémises — un taux élevé signale interférences ou mauvais signal. Chap. 67.

**Roaming** : passage d'un client d'un AP à l'autre en se déplaçant. Chap. 56–57.

**Rogue AP** : point d'accès non autorisé. Chap. 64.

**RSSI** : puissance du signal reçu, en dBm (négatif — plus proche de 0 = mieux). Chap. 61.

**SAE** : authentification WPA3 remplaçant le PSK, résistante au brute-force hors ligne. Chap. 46.

**SFP** : module optique enfichable — sur l'AP761 : port SFP GE pour uplink fibre [constructeur]. Chap. 6.

**SNR** : rapport signal/bruit — **l'indicateur qui gouverne vraiment la qualité**. Chap. 61.

**SSID** : le nom du réseau Wi-Fi. 16 max par radio sur l'AP761 [constructeur]. Chap. 43.

**STA isolation** : voir Isolation client.

**Sticky client** : client qui reste accroché à un AP lointain au lieu de roamer.

**STBC / CDD / CSD** : techniques de diversité d'antennes pour fiabiliser la réception [constructeur].

**Surtension (protection)** : 6 kV sur les ports Ethernet de l'AP761 [constructeur] — contre les transitoires, pas la foudre directe. Chap. 26.

**Syslog** : centralisation des journaux. Chap. 69.

**TWT** : Target Wake Time — négociation des cycles de sommeil des objets connectés (autonomie) [constructeur]. Chap. 12.

**Uplink** : le lien filaire de l'AP vers le réseau (RJ45 ou fibre).

**VAP** : profil virtuel d'AP — l'association SSID + sécurité + VLAN sur une radio (terminologie Huawei). Chap. 34.

**VLAN** : réseau local virtuel — segmentation par SSID. Chap. 52.

**VRP** : le système d'exploitation des équipements Huawei (l'interface CLI de ce guide).

**WIDS/wIPS** : détection / prévention d'intrusion sans fil [constructeur]. Chap. 63.

**Wi-Fi 6 (802.11ax)** : la génération de l'AP761 — OFDMA, MU-MIMO, 1024-QAM, BSS coloring. Chap. 12.

**Wi-Fi 7 (802.11be)** : MLO, 320 MHz, 4096-QAM, multi-RU — **non supporté par l'AP761**. Chap. 14–21.

**WLAN** : réseau local sans fil.

**WMM** : Wi-Fi Multimedia — les 4 classes de priorité QoS (voix, vidéo, best effort, background) [constructeur]. Chap. 54.

**WPA2 / WPA3** : protocoles de sécurité Wi-Fi. Chap. 44–48.

**Zombie (client)** : client accroché avec un signal très faible et un débit minime, qui monopolise l'airtime de toute la cellule. Chap. 90.

## 113. Quiz : 10 questions

Réponds sans regarder les réponses (chap. 114). L'objectif n'est pas le 10/10, c'est de repérer tes angles morts.

**Q1.** L'AP761 est-il un point d'accès Wi-Fi 7 ? Si non, quelles sont ses caractéristiques radio réelles (standard, bandes, MIMO, débit max) ?

**Q2.** L'AP761 a deux ports filaires : un RJ45 GE et un SFP GE. Peut-on les utiliser simultanément (l'un en uplink, l'autre en downlink) ? Que se passe-t-il si les deux sont branchés ?

**Q3.** Ton switch ne fournit que du PoE 802.3af sur le port de l'AP761. L'AP démarre. Quel est le problème, et que dit le constructeur à ce sujet ?

**Q4.** Tu dois couvrir une cour rectangulaire de 60 × 40 m avec des AP761. Les antennes sont directionnelles (65° horizontal). Où places-tu les AP et comment les orientes-tu ? Pourquoi ne peux-tu pas compter sur « l'arrière » de l'AP ?

**Q5.** En 2.4 GHz, quels sont les seuls canaux à utiliser en Europe, et pourquoi ? Que penses-tu du 40 MHz en 2.4 GHz ?

**Q6.** Un client se plaint de coupures de son appel Wi-Fi tous les jours vers 14h, et tu constates que l'AP a changé de canal 5 GHz. Explique le mécanisme et propose deux solutions.

**Q7.** Quelle est la différence entre WPA2-PSK, WPA3-SAE et le mode transition ? Dans quel cas utilises-tu chacun ?

**Q8.** Un seul smartphone ancien reste connecté en 2.4 GHz à −85 dBm, et depuis, tout le monde se plaint que « le Wi-Fi est lent ». Explique le mécanisme physique et donne deux remèdes côté AP.

**Q9.** Le WIDS te signale un « rogue AP » avec un SSID proche du tien. Détaille ta procédure avant toute contre-mesure.

**Q10.** La direction te demande « quand passe-t-on au Wi-Fi 7 en extérieur ? ». Construis une réponse argumentée en 3 points (besoin, parc clients, calendrier), en citant les vrais modèles eKit concernés.

## 114. Quiz : réponses commentées

**R1.** **Non.** L'AP761 est un AP **extérieur Wi-Fi 6 (802.11ax)** [constructeur] : bi-bande **2.4 GHz 2×2 + 5 GHz 2×2**, **1.775 Gbps** max agrégés (0.575 + 1.2 Gbps), canaux 20/40/80 MHz, 1024-QAM max. Pas de MLO, pas de 320 MHz, pas de 6 GHz, pas de 4096-QAM. Les modèles eKit Wi-Fi 7 d'extérieur sont l'**AP771** (3.57 Gbps) et l'**AP772E** (6.45 Gbps). *Si tu as répondu « Wi-Fi 7 » : relis l'avertissement en tête de guide — c'est l'erreur que ce guide existe pour éviter.*

**R2.** **Non.** Ce sont des ports **combo** : un seul actif à la fois, et **le SFP optique est prioritaire** s'il est branché [constructeur]. Brancher le SFP « pour tester » pendant que le cuivre fonctionne fait basculer l'AP sur la fibre — et si elle n'est pas prête, l'AP disparaît. L'alimentation reste **toujours** sur le RJ45 (PoE-In) : la fibre ne transporte pas d'énergie.

**R3.** Le 802.3af fournit ~12.95 W utiles, l'AP en veut 17.7 W max : le constructeur précise que **l'AP fonctionne en mode restreint (fonctions limitées)** en 802.3af [datasheet]. Symptômes : débit plafonné, USB/BLE coupés, puissance radio bridée. **Solution : passer en 802.3at (PoE+)** ou injecteur 30 W, et réserver 25 W par AP sur le budget du switch.

**R4.** Les AP se placent **en bordure de la zone, pointés vers l'intérieur**, secteur de 65° couvrant la cour — typiquement 2 AP en vis-à-vis ou en angles opposés selon la forme. On ne compte pas sur l'arrière car les antennes sont **directionnelles** : l'arrière est sourd (10–11 dBi concentrés vers l'avant [constructeur]). Hauteur 4–8 m, léger downtilt.

**R5.** **Canaux 1, 6 et 11** : les seuls 3 canaux 20 MHz qui ne se recouvrent pas en Europe (2.412–2.472 GHz). Le 40 MHz en 2.4 GHz est **à proscrire** en environnement partagé : il occupe 2 des 3 canaux propres, double le débit théorique mais divise la capacité du voisinage — et en extérieur, le voisinage est partout.

**R6.** **Mécanisme : DFS.** L'AP est sur un canal 52–140 ; à 14h (météo ? radar ?), il détecte une signature radar et **doit légalement quitter le canal** (chap. 41) → coupure pour tous les clients 5 GHz le temps de la bascule. **Solutions :** (1) migrer les usages critiques vers les canaux **non-DFS 36–48** ; (2) vérifier les logs pour confirmer (`radar detected`), noter l'AP en DFS dans le dossier de site, et si c'est quotidien, investiguer une source locale (chap. 65).

**R7.** **WPA2-PSK** : clé partagée, le standard PME — simple, mais la clé fuitée expose tout le monde, et vulnérable au brute-force hors ligne si clé faible. **WPA3-SAE** : résiste au brute-force hors ligne (échange par essai en ligne), impose AES+PMF — mais **exclut les clients non compatibles**. **Transition** : accepte les deux en parallèle — le mode de **migration** : modernes en SAE, anciens en WPA2, avec un plan pour passer en WPA3-only quand le parc suit.

**R8.** **Mécanisme : le client zombie et l'airtime.** Le Wi-Fi partage le **temps**, pas le débit : un client à 1 Mbps occupe le canal ~100× plus longtemps qu'un client à 100 Mbps pour la même quantité de données → toute la cellule ralentit. **Remèdes :** (1) **relever les débits de base** (interdire 1/2/5.5/11 Mbps en 2.4 GHz) pour empêcher l'association à débit minable ; (2) activer l'**airtime fairness** si supporté, ou limiter le débit par client.

