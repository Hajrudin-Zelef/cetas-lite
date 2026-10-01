---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-19
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [1907, 1995]
sha256: 25a38ba9ab8ee68a6f9f96640edf03da48622e1930ddd72614c4a25c8381ea0f
---

# Guide ultra-complet — Huawei eKit AP761

**Diagnostic ordonné :**
1. **Le VLAN du SSID est-il autorisé sur le trunk ?** (`port trunk allow-pass vlan X` côté AP + côté switch — chap. 34/52). C'est la cause n° 1, et de loin.
2. **Le serveur DHCP répond-il sur ce VLAN ?** Tester en filaire sur le même VLAN : si le filaire n'a pas d'IP non plus, le problème n'est pas le Wi-Fi.
3. **Le scope DHCP est-il plein ?** Un scope /24 plein à cause de baux fantômes = plus d'IP pour personne. Vérifier les baux, réduire la durée des baux sur les VLAN invités (1–4 h).
4. **Le relay DHCP (ip-helper) est-il configuré ?** Si le serveur DHCP n'est pas sur le même VLAN, il faut un relais sur la passerelle du VLAN.
5. **Le DHCP snooping bloque-t-il ?** (chap. 55) : port uplink non déclaré « trusted » = réponses DHCP jetées.

**Test décisif :** mettre une IP **statique** sur le client (dans le bon sous-réseau) : si ça marche en statique = c'est bien le DHCP. Sinon = problème VLAN/routage.

## 81. Cas 4 : débit très en-deçà des attentes

**Symptômes :** speedtest à 20 Mbps alors qu'on attend 200+.

**Checklist par probabilité :**
1. **Sur quelle bande est le client ?** (2.4 vs 5 GHz) : un client en 2.4 GHz à 20 MHz ne dépassera jamais ~70 Mbps. Forcer le test en 5 GHz (chap. 42).
2. **Le lien filaire est-il en gigabit ?** `display interface` : un port tombé à 100 Mbps (câble abîmé, injecteur 100M) bride tout l'AP à ~95 Mbps. **Cause ultra-fréquente et sous-diagnostiquée.**
3. **La largeur de canal est-elle celle prévue ?** Un 80 MHz retombé à 20 MHz (interférence, DFS) divise le débit par 4. Vérifier la largeur réelle.
4. **L'airtime est-il saturé ?** Channel utilization > 70 % = le canal est plein (voisins, clients zombies). Changer de canal ou réduire la largeur.
5. **Le client est-il le goulot ?** Tester avec **2–3 clients différents** (un vieux smartphone ne fera jamais 300 Mbps). Tester aussi en **montant** : un mauvais montant = souvent le client.
6. **Le serveur de test est-il le goulot ?** Un speedtest vers Internet mesure la fibre + le Wi-Fi. **iperf3 vers un serveur du LAN** pour isoler le Wi-Fi.
7. **Y a-t-il un client zombie ?** (cas 13) : un client à 1 Mbps en 2.4 GHz plombe toute la cellule.

**Méthode de mesure propre :** iperf3, 3 runs, noter down/up à 3 distances (chap. 60). Un seul speedtest ne prouve rien.

## 82. Cas 5 : clients 5 GHz qui restent accrochés au 2.4 GHz

**Symptômes :** le 5 GHz est vide, le 2.4 GHz est saturé, les utilisateurs se plaignent de lenteur.

**Causes et remèdes :**
1. **Puissance 2.4 GHz trop forte** par rapport au 5 GHz : le client « entend » mieux le 2.4 GHz. **Baisser le 2.4 GHz de 3–6 dB** sous le 5 GHz (chap. 40).
2. **Pas de band steering** : l'activer (chap. 42) — avec prudence sur les parcs IoT.
3. **Le client est loin** : à −75 dBm en 5 GHz, un client préfère légitimement le 2.4 GHz à −65 dBm. C'est normal : la solution est un **meilleur placement d'AP**, pas du forcing.
4. **Le client ne supporte pas le 5 GHz** : objets connectés, vieux PC. **SSID IoT dédié en 2.4 GHz** (chap. 43) pour les isoler du problème.
5. **Le 5 GHz est en DFS avec CAC** : après un redémarrage, le 5 GHz met 60 s à apparaître — les clients s'accrochent au 2.4 GHz entre-temps et y restent (sticky client). Préférer les canaux non-DFS si c'est récurrent.

**Vérification :** ratio clients 5 GHz / total > 70 % en zone couverte (chap. 42). En dessous : appliquer les remèdes ci-dessus un par un.

## 83. Cas 6 : coupures lors des déplacements (roaming)

**Symptômes :** en marchant d'une zone à l'autre, la visio coupe, l'appel tombe, le ping perd 30+ paquets.

**Diagnostic :**
1. **Y a-t-il un trou de couverture ?** Faire le tour avec un relevé RSSI (chap. 60) : si le client passe par une zone < −75 dBm entre les deux AP, il n'y a pas de roaming qui tienne — il faut un AP ou un repositionnement.
2. **Le 802.11r est-il actif ?** Sans r, chaque roaming en 802.1X coûte 200–500 ms (chap. 56). L'activer et **tester avec le parc réel**.
3. **Les AP sont-ils dans le même domaine de mobilité ?** Même SSID/sécurité/VLAN/groupe (chap. 57). En Fat avec des AP indépendants, le roaming est best-effort.
4. **Le client est-il sticky ?** Certains clients (vieux Android, IoT) ne roament jamais proprement. Tester avec 2–3 modèles différents pour isoler.
5. **Y a-t-il un changement de VLAN ?** Même SSID mais VLAN différent selon l'AP = changement de sous-réseau = sessions coupées. Uniformiser.

**Test de référence** (chap. 57) : ping continu en marchant. Objectif : **< 5 pings perdus** par transition en PSK avec 802.11r, **< 10** en 802.1X.

## 84. Cas 7 : interférences 2.4 GHz (micro-ondes, Bluetooth, caméras)

**Symptômes :** débit en dents de scie, retries élevés, problèmes **à heures fixes** (le midi = le micro-ondes du food-truck ; le soir = les caméras du voisin).

**Diagnostic :**
1. **Corréler avec l'heure :** un problème qui apparaît tous les jours à 12h00 n'est pas un problème Wi-Fi, c'est un **micro-ondes**. Noter les horaires précis pendant une semaine.
2. **Scanner :** les réseaux Wi-Fi voisins ne sont que la partie visible. Les **non-Wi-Fi** (micro-ondes, Bluetooth, caméras analogiques 2.4 GHz, ZigBee) n'apparaissent pas dans les scans Wi-Fi mais pourrissent le canal pareil.
3. **Tester en coupant les suspects** un par un (chap. 65).
4. **Vérifier le SNR**, pas seulement le RSSI (chap. 61) : un bon RSSI avec un mauvais SNR = interférence.

**Remèdes par source :**
| Source | Remède |
|---|---|
| Four à micro-ondes | Éloigner l'AP / changer de canal ne sert à rien (pollution large) → **déplacer l'AP** ou blinder |
| Caméra sans fil du voisin | Passer en 5 GHz les clients critiques ; discuter avec le voisin (canal) |
| Bluetooth (enceintes) | Le Bluetooth saute de fréquence : impact modéré, mais réel à proximité immédiate |
| Son propre IoT 2.4 GHz | SSID IoT sur canal 1/6/11 dédié, séparé du SSID principal |

**Si rien ne marche :** basculer les usages critiques en **5 GHz** (plus de canaux, moins de pollution non-Wi-Fi) et réserver le 2.4 GHz à l'IoT.

## 85. Cas 8 : radar DFS — l'AP change de canal tout seul

**Symptômes :** coupures brèves et simultanées de tous les clients 5 GHz, l'AP est sur un canal différent après, les logs montrent `radar detected` / `channel change`.

**C'est un comportement NORMAL et OBLIGATOIRE** (chap. 41), pas une panne. Le diagnostic consiste à le confirmer puis à décider :

1. **Confirmer :** logs (`display logbuffer` — chercher `DFS`, `radar`, `channel switch`), corrélation avec la météo (orages) ou la proximité d'un aéroport/météo radar.
2. **Si c'est occasionnel** (quelques fois par an) : ne rien faire, c'est la vie en DFS.
3. **Si c'est fréquent** (hebdomadaire) : **migrer vers les canaux non-DFS 36–48** (chap. 38). On perd du spectre, on gagne en stabilité — pour un usage pro, la stabilité gagne toujours.
4. **Si c'est quotidien à heure fixe** : suspecter un faux positif (un équipement local qui ressemble à un radar) — investiguer la source (chap. 65).

**À noter dans le dossier de site** quels AP sont en DFS : dans 6 mois, quand ça recommencera, tu sauras où regarder en 30 secondes.

## 86. Cas 9 : PoE insuffisant — fonctions restreintes en 802.3af

**Symptômes :** l'AP démarre, le Wi-Fi « marche », mais : débit plafonné, USB/BLE inactifs, puissance radio qui ne monte pas, messages `power insufficient` dans les logs.

**Rappel constructeur :** en **802.3af**, l'AP761 fonctionne en **mode restreint** [datasheet]. Le 802.3af fournit 12.95 W utiles, l'AP en veut 17.7 W max.

