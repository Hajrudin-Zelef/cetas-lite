---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-15
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [1457, 1559]
sha256: ab8bc0af016216a4c79f7025456038ca461fa888d7af7627519d600b5e57f167
---

# Guide ultra-complet — Huawei eKit AP761

| Usage | Débit par client (utile) | Clients/AP | Largeur 5 GHz | Notes |
|---|---|---|---|---|
| Capteurs IoT | < 1 Mbps | 150+ | 20 MHz | 2.4 GHz suffit souvent |
| Navigation visiteurs (terrasse) | 2–5 Mbps | 80–120 | 40 MHz | Pics le midi/week-end |
| Bureautique extérieure | 5–10 Mbps | 40–60 | 40 MHz | Visio incluse |
| Visioconférence dense | 2–5 Mbps/client mais latence critique | 20–30 | 40 MHz | WMM + CAC (chap. 54) |
| Vidéosurveillance (caméras Wi-Fi) | 4–8 Mbps/caméra montant | 10–15 caméras | 40/80 MHz | Le montant est le goulot — tester ! |
| Événementiel (fête, marché) | 1–3 Mbps | 150–200 | 40 MHz | Prévoir +50 % de marge |

**Calcul de capacité simplifié :**
```
Capacité utile d'un AP761 ≈ 300–500 Mbps agrégés (les deux radios, conditions correctes)
Clients supportables = Capacité utile / débit par client
Exemple : 400 Mbps / 5 Mbps = 80 clients en navigation
```
Toujours garder **30 % de marge** : le Wi-Fi est un médium partagé avec des pics.

**Le cas des caméras Wi-Fi (fréquent en extérieur) :** 10 caméras à 6 Mbps montant = 60 Mbps de trafic **montant** constant. Le Wi-Fi est moins bon en montant qu'en descendant (les clients ont des petites antennes et peu de puissance). **Tester avec le nombre réel de caméras** avant de valider — c'est le cas d'usage qui casse le plus de dimensionnements théoriques.

## 60. Site survey simplifié sans outils pro — méthode terrain

Pas d'analyseur de spectre à 5000 € ? La méthode smartphone suffit pour 90 % des sites PME.

**Matériel :** un smartphone Android + app d'analyse Wi-Fi gratuite (ex. WiFi Analyzer), un mètre, le plan du site imprimé.

**Protocole (à faire AP par AP, une fois installé) :**
1. **Scan passif :** depuis 4–6 points de la zone, noter pour chaque AP : SSID, canal, RSSI, largeur. Repérer les voisins forts.
2. **Mesure de couverture :** marcher en grille (tous les 10–15 m), noter le RSSI du SSID principal sur le plan. Tracer la courbe des **−67 dBm** : c'est ta zone « confort ».
3. **Test de débit :** avec iperf3 vers un serveur du LAN (pas un speedtest Internet — on teste le Wi-Fi, pas la fibre), mesurer TCP descendant/montant à 3 distances (proche, moyen, limite).
4. **Test de roaming :** ping continu en marchant entre AP (chap. 57).
5. **Photo et notes :** chaque point de mesure = une photo du lieu + le relevé. Dans 2 ans, tu comprendras ce que tu as fait.

**Modèle de fiche de mesure (à recopier dans le dossier de site) :**
```
Date : ___  AP : AP761-Cour-Est  Téléphone : ___
Point A (entrée cour) : RSSI -55 dBm, canal 36/80MHz, iperf down 320 Mbps / up 280 Mbps
Point B (milieu cour) : RSSI -64 dBm, iperf down 180 Mbps / up 150 Mbps
Point C (fond cour)   : RSSI -72 dBm, iperf down 60 Mbps / up 40 Mbps  -> limite, prévoir AP2 si usage dense
Voisins forts : Freebox-XYZ canal 1 (-75 dBm), Entreprise-Voisine canal 44 (-70 dBm)
```

**Quand passer au survey pro :** site > 20 AP, exigences voix/vidéo critiques, environnement RF hostile (aéroport, hôpital, industrie lourde). Là, un vrai survey prédictif + validation active se justifie.

## 61. Lecture d'un relevé : RSSI, SNR, bruit

Les trois chiffres qui disent tout :

| Indicateur | Bon | Limite | Mauvais |
|---|---|---|---|
| **RSSI** (signal reçu) | > −60 dBm | −60 à −70 dBm | < −70 dBm (données), < −67 dBm (voix) |
| **Bruit** (noise floor) | < −90 dBm | −85 à −90 dBm | > −85 dBm (environnement bruyant) |
| **SNR** (RSSI − bruit) | > 25 dB | 20–25 dB | < 20 dB |

**Traductions terrain :**
- RSSI −55 dBm, bruit −95 dBm → SNR 40 dB : **excellent**, 1024-QAM possible à courte distance.
- RSSI −70 dBm, bruit −95 dBm → SNR 25 dB : **correct**, 256-QAM, débits honorables.
- RSSI −70 dBm, bruit −85 dBm → SNR 15 dB : **mauvais** malgré un RSSI « acceptable » — le bruit tue. Chercher la source (chap. 62, cas 7).
- **RSSI qui fluctue de ±10 dB** au même endroit : trajets multiples (réflexions), végétation qui bouge, ou interférence intermittente — à investiguer.

**L'erreur classique :** ne regarder que le RSSI. Un client à −65 dBm dans un bruit à −80 dBm (SNR 15 dB) aura de pires débits qu'un client à −72 dBm dans un bruit à −95 dBm (SNR 23 dB). **Le SNR gouverne, pas le RSSI seul.**

## 62. Erreurs classiques de placement d'AP extérieur

La galerie des horreurs, vue sur le terrain. À vérifier sur chaque chantier :

1. **L'AP derrière un obstacle métallique** (auvent, bardage, armoire) : le métal fait cage de Faraday partielle — le secteur est haché.
2. **L'AP trop haut** (« pour voir loin ») : zone morte au pied du mât + asymétrie avec les clients (chap. 40).
3. **L'AP collé à un mur** : la moitié du secteur part dans le mur, l'autre moitié est déformée par réflexion.
4. **Deux AP dos à dos sur le même mât, même canal** : ils se partagent l'airtime au lieu de se compléter.
5. **L'AP à côté du transformateur / du variateur** : bruit électromagnétique large bande.
6. **L'AP sous un arbre** : en été avec les feuilles mouillées, le 5 GHz perd 10–20 dB. Et ça bouge avec le vent (fluctuations).
7. **Le câble qui pend** : prise au vent, usure par frottement, eau qui suit le câble.
8. **L'AP orienté vers le parking… des voisins** : on couvre la zone utile, pas la plus facile d'accès.
9. **Oublier l'hiver** : un AP posé en été derrière des arbres feuillus se retrouve à découvert en hiver (ou l'inverse : la végétation estivale n'existait pas au survey d'hiver).
10. **Pas de repérage photo** : dans 3 ans, personne ne sait quel AP est où — étiquetage + plan à jour (chap. 101).

## 63. WIDS/wIPS : détection d'intrusion sans fil

**WIDS** (détection) / **wIPS** (prévention) [constructeur] : l'AP761 surveille l'air et détecte les comportements suspects. C'est le « vigile » du Wi-Fi.

**Ce que le WIDS détecte typiquement :**
| Attaque / anomalie | Principe |
|---|---|
| **Rogue AP** | Un AP non autorisé branché sur ton réseau filaire (chap. 64) |
| **Evil twin** | Un AP pirate qui usurpe ton SSID pour intercepter les clients |
| **Désauthentification massive** | Flood de trames de désauthentification (attaque DoS) |
| **Flood de sondes / d'associations** | Tentative de saturation de l'AP |
| **Ad-hoc / Wi-Fi Direct suspects** | Réseaux pair-à-pair qui contournent l'infrastructure |
| **Clients en mode promiscuous** | Sniffing potentiel |

**wIPS vs WIDS :** le WIDS **alerte**, le wIPS **contre-attaque** (désauthentification du rogue, par exemple). La contre-attaque est à manier avec précaution : désauthentifier un AP voisin légitime parce qu'on l'a mal classifié = incident avec le voisin + possible non-conformité (brouillage actif). **Recommandation : WIDS en alerte d'abord, wIPS en contre-mesure seulement sur les rogues confirmés branchés sur ton filaire** (chap. 64).

**Prérequis :** le WIDS consomme du temps radio (l'AP doit écouter hors de son canal). Sur un AP chargé, activer le WIDS peut coûter quelques % de capacité — acceptable sur l'AP761 en usage PME, à surveiller sur les sites denses.

## 64. Rogue AP : détecter, classifier, contenir

**Définition :** un **rogue AP** = un point d'accès non autorisé. Deux cas très différents :

| Type | Exemple | Danger | Action |
|---|---|---|---|
| **Rogue « naïf »** | Un salarié branche sa box 4G/routeur perso « pour avoir du Wi-Fi dans son bureau » | Il crée une porte d'entrée non contrôlée vers ton LAN | Le trouver (il est sur ton filaire), le débrancher, sensibiliser |
| **Rogue « malveillant »** | Evil twin qui usurpe ton SSID sur le parking | Interception d'identifiants, MITM | Contenir (désauth), alerter, enquête |

