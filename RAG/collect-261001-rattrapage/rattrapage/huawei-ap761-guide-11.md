---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-11
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["EU", "Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [1035, 1138]
sha256: d474c84cb07107ff6ec114213cc4e102e90ff54d469dfca5ed9cdc521c2f5764
---

# Guide ultra-complet — Huawei eKit AP761

**Méthode terrain :**
1. Scanner la zone (smartphone + app) : noter les réseaux 5 GHz forts et leurs canaux.
2. Attribuer à chaque AP un canal 80 MHz (ou 40 MHz) **non chevauchant** avec ses voisins proches.
3. **Plan 80 MHz avec 2 AP proches :** AP1 = canal 36 (bloc 36–48), AP2 = canal 52 ou 100 (bloc DFS). Jamais deux blocs 80 MHz qui se recouvrent.
4. **Plan 80 MHz avec 4+ AP :** il faudra réutiliser — espacer au maximum les réutilisations (AP1 et AP4 sur le même canal s'ils sont à l'opposé du site).

**Tableau : combien de canaux 80 MHz propres en Europe :**
| Bloc | Canaux 80 MHz possibles | DFS ? |
|---|---|---|
| 36–48 | 1 (canal 42) | Non |
| 52–64 | 1 (canal 58) | Oui |
| 100–140 | 2–3 (106, 122, 138) | Oui |
| 149–165 | 1 (155, selon pays) | Non (selon pays) |

Soit **4 à 5 canaux 80 MHz** en tout — et les voisins les utilisent aussi. D'où l'intérêt du 40 MHz quand la densité augmente (chap. 39).

## 39. Largeur de canal : 20/40/80 MHz, que choisir

**Le compromis en une phrase :** canal large = plus de débit par client, mais moins de canaux disponibles et plus de sensibilité aux interférences.

| Largeur | Débit relatif | Canaux dispo (5 GHz EU) | Quand l'utiliser |
|---|---|---|---|
| 20 MHz | ×1 (référence) | ~13–19 | 2.4 GHz toujours ; 5 GHz en zone très dense ou très bruitée |
| 40 MHz | ×2 | ~6–9 | 5 GHz : bon compromis densité/débit en extérieur partagé |
| 80 MHz | ×4 | ~4–5 | 5 GHz : débit max, peu d'AP, environnement propre |
| 160 MHz | ×8 | 1–2 (DFS) | **Non supporté par l'AP761** (max 80 MHz [constructeur]) — cité pour mémoire |

**Recommandations terrain pour l'AP761 :**
- **2.4 GHz : 20 MHz.** Toujours. (chap. 37)
- **5 GHz, site isolé (campagne, zone industrielle propre) : 80 MHz.** Tu exploites le max du Wi-Fi 6.
- **5 GHz, site partagé (ville, lotissement, zone commerciale) : 40 MHz.** Tu divises par deux ta sensibilité aux voisins et tu doubles tes choix de canaux.
- **5 GHz, zone ultra-dense ou usage voix critique : 20 MHz** sur canaux non-DFS (36–48). La fiabilité avant le débit.

**Le piège du « 80 MHz partout » :** un canal 80 MHz qui chevauche un réseau voisin fort = **pire** qu'un 40 MHz propre. Le débit réel d'un 80 MHz brouillé est inférieur à celui d'un 40 MHz propre. **Un canal propre et étroit bat toujours un canal large et sale.** C'est LA phrase à retenir de ce chapitre.

**Configuration (principe CLI — noms à valider avec `?` selon version) :**
```
[AP761-wlan-view] radio-profile name RP-5G
[AP761-wlan-radio-prof-RP-5G] channel-width 40MHz
# ou 20MHz / 80MHz selon le plan
```

## 40. Puissance d'émission : régler sans bourriner

**Valeur constructeur :** puissance combinée max **28 dBm (2.4 GHz) / 27 dBm (5 GHz)**, réglable par pas de **1 dBm**, sous réserve de la réglementation locale.

**Le réflexe à perdre : « mettre à fond pour porter loin ».** La puissance max crée trois problèmes :
1. **Asymétrie :** l'AP crie à 27 dBm, le smartphone répond à 15–20 dBm. Le client « entend » l'AP de loin mais l'AP n'entend pas le client → associations fantômes, débits minables, roaming qui ne se déclenche pas (le client croit avoir du signal).
2. **Pollution :** tu arroses les voisins qui t'arrosent en retour — tout le monde perd en airtime.
3. **Non-conformité :** la réglementation limite la PIRE (puissance + gain d'antenne). Avec 11 dBi d'antenne, la puissance conduite doit être réduite d'autant : le firmware le fait si le domaine réglementaire est correct (chap. 36) — encore une raison de le régler en premier.

**Méthode de réglage terrain :**
1. Partir de la puissance **réglementaire max** du domaine (le firmware la connaît).
2. Faire un tour de la zone avec un smartphone : noter le RSSI aux limites de la zone utile.
3. **Baisser** jusqu'à ce que le RSSI en limite de zone soit ≈ **−67 dBm** (seuil voix/données confortables) à **−70 dBm** (données OK).
4. Entre deux AP qui se recouvrent : baisser pour que le recouvrement à −67 dBm soit de **15–20 %** — assez pour le roaming, pas assez pour se marcher dessus.

**Ordres de grandeur (indicatifs) :**
| Contexte | Puissance 5 GHz typique | Pourquoi |
|---|---|---|
| Cour 30×30 m, 1 AP | 14–17 dBm | Couvrir sans arroser le quartier |
| Parking 100 m, 2 AP | 17–20 dBm | Porter le long du parking |
| Grand site isolé | 20–23 dBm | Pas de voisins, on peut pousser |

> En 2.4 GHz, mettre **3 à 6 dB de moins** qu'en 5 GHz : le 2.4 GHz porte naturellement plus loin, et on veut que les clients **préfèrent le 5 GHz** (band steering, chap. 42).

## 41. DFS en 5 GHz : ce qu'il faut savoir

**DFS (Dynamic Frequency Selection)** [constructeur] : sur les canaux 52–140, l'AP doit détecter les **radars** (météo, aviation, militaires) et libérer le canal s'il en détecte un. C'est une obligation réglementaire, pas une option.

**Le cycle DFS :**
1. **CAC (Channel Availability Check) :** avant d'émettre sur un canal DFS, l'AP écoute **60 secondes** (10 minutes sur les canaux 120–128 dans certains pays). Pendant ce temps : pas de Wi-Fi sur cette radio.
2. **Surveillance continue :** l'AP écoute en permanence les signatures radar.
3. **Détection → évacuation :** si un radar est détecté, l'AP doit quitter le canal en quelques secondes et ne pas y revenir pendant **30 minutes** (NOP — Non-Occupancy Period).

**Conséquences terrain :**
- Un AP qui **redémarre** sur un canal DFS = 60 s sans Wi-Fi 5 GHz. Sur un site critique, préférer les canaux **non-DFS (36–48)**.
- Des **coupures inexpliquées** qui reviennent par météo orageuse ou près d'un aéroport = probablement du radar (cas 8).
- L'AP761 supporte le DFS [constructeur] : c'est transparent, mais il faut savoir que ça existe pour diagnostiquer.

**Bonnes pratiques :**
- Usages critiques (voix sur Wi-Fi, contrôle d'accès, vidéosurveillance) : **canaux 36–48 uniquement**.
- Noter dans le dossier de site quels AP sont en DFS : en dépannage, c'est la première question à poser.
- Après un changement de canal DFS, vérifier que les clients se sont bien réassociés (certains vieux clients mettent du temps à suivre).

## 42. Band steering et préférence de bande

**Le problème :** les clients préfèrent naturellement le **2.4 GHz** (il porte plus loin, le signal paraît « meilleur »), alors que le **5 GHz** est plus rapide et moins encombré. Résultat : tout le monde s'entasse sur le 2.4 GHz et le 5 GHz reste vide.

**Le band steering** (fonctionnalité classique des AP Huawei, nom exact à vérifier selon version) : l'AP **retarde ou refuse** les premières tentatives d'association en 2.4 GHz pour encourager les clients bi-bande à essayer le 5 GHz.

**Réglages qui aident (sans band steering) :**
1. **Puissance 2.4 GHz < puissance 5 GHz** de 3–6 dB (chap. 40) : le client « voit » le 5 GHz aussi fort que le 2.4 GHz.
2. **SSID identique** sur les deux bandes : le client choisit librement (la plupart des clients modernes préfèrent alors le 5 GHz).
3. **Débits de base 2.4 GHz relevés** (chap. 13) : les clients loin en 2.4 GHz ne s'associent plus.

**Quand NE PAS faire de steering :**
- Objets connectés 2.4 GHz uniquement (caméras, portiers) : ils ne comprendront pas le refus et boucleront en échec d'association. **SSID dédié 2.4 GHz** pour l'IoT (chap. 43).
- Si ça crée plus de problèmes que ça n'en résout : le steering agressif casse l'association de certains clients exotiques. **Tester avec le parc réel** avant de généraliser.

**Vérification :** après activation, contrôler la répartition des clients (`display wlan sta` ou console cloud) : objectif **> 70 % des clients bi-bande sur le 5 GHz** en zone couverte par les deux bandes.

## 43. SSID : stratégie de nommage et nombre

