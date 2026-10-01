---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-14
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [1364, 1456]
sha256: b1cc45c50d069c4bce82b2752f8bb66c6b198efdc56b3ea032df49464981cc86
---

# Guide ultra-complet — Huawei eKit AP761

| Classe (AC) | Usage | Comportement |
|---|---|---|
| AC_VO (voix) | Téléphonie sur Wi-Fi, interphonie | Accès prioritaire au canal, files séparées |
| AC_VI (vidéo) | Visioconférence, caméras | Prioritaire après la voix |
| AC_BE (best effort) | Navigation, mail | Le tout-venant |
| AC_BK (background) | Mises à jour, sauvegardes | Sacrifié en premier en cas de congestion |

**Ce que le WMM fait vraiment :** il donne **plus de chances d'accéder au canal** aux trames prioritaires (temps d'attente plus court, backoff réduit). Il ne **réserve** pas de bande passante : si 50 clients font de la visio en même temps, le WMM ne créera pas de capacité magique.

**Bonnes pratiques :**
1. **Marquer en amont** : le WMM se base sur les marques DSCP/802.1p des trames. Si ton réseau ne marque pas (téléphones qui marquent EF, switch qui fait confiance), le WMM ne sert à rien. **Vérifier la chaîne de marquage de bout en bout.**
2. **Limiter le débit des classes basses** sur le SSID invité (rate limiting par client — fonction classique, nom à vérifier selon version) : un invité qui fait ses mises à jour Windows ne doit pas écraser la visio du salarié.
3. **CAC (Call Admission Control)** [constructeur] : limiter le nombre d'appels voix simultanés par radio pour garantir la qualité — à activer si tu fais de la téléphonie sur Wi-Fi en extérieur (interphonie de parking, par exemple).

**Ordre de grandeur :** un appel voix G.711 = ~100 kbps avec les en-têtes ; une visio correcte = 1–2 Mbps. Le dimensionnement voix sur Wi-Fi extérieur reste un exercice délicat : prévoir des tests réels (chap. 60).

## 55. ACL et DHCP snooping / DAI / IPSG

L'AP761 embarque des fonctions de **sécurité du plan filaire** [constructeur] : DHCP snooping, DAI (Dynamic ARP Inspection), IPSG (IP Source Guard), ACL. Ce sont des fonctions de switch, présentes sur l'AP pour protéger son port.

| Fonction | Ce qu'elle fait | Quand l'activer |
|---|---|---|
| **DHCP snooping** | Ne laisse passer les réponses DHCP que des serveurs de confiance | Toujours : tue les serveurs DHCP pirates (box 4G d'un visiteur branchée au réseau) |
| **DAI** | Vérifie que les ARP correspondent aux baux DHCP | Avec le snooping : tue l'ARP spoofing |
| **IPSG** | Ne laisse passer que les IP attribuées par DHCP | Réseaux invités/IoT : tue le vol d'IP |
| **ACL** | Filtre L2/L3/L4 sur le port | Complément au firewall pour des règles locales |

**En pratique sur un AP :** ces fonctions protègent surtout le **trafic ponté localement** (forward-mode direct). Si tout ton trafic remonte au contrôleur ou au firewall, c'est là-bas que se joue l'essentiel — mais activer le snooping sur l'AP coûte peu et protège le segment local.

**Attention :** mal configuré (serveur DHCP non déclaré « trusted »), le DHCP snooping **casse tout le DHCP** du segment. **Toujours déclarer les ports/uplinks de confiance en premier, tester, puis activer.**

## 56. Itinérance : 802.11k, 802.11v, 802.11r

L'AP761 supporte **802.11k, 802.11v et 802.11r** [revendeurs — à vérifier sur la fiche du modèle exact pour le détail d'implémentation]. Les trois ont des rôles différents, et il faut les comprendre pour ne pas les confondre :

| Standard | Nom | Rôle | Analogie |
|---|---|---|---|
| **802.11k** | Radio Resource Management | L'AP donne au client la **liste des AP voisins** avec leurs canaux et leur charge | Le GPS qui te montre les stations-service autour |
| **802.11v** | BSS Transition Management | L'AP peut **suggérer** au client de migrer vers un meilleur AP (et le client peut demander conseil) | Le copilote qui dit « sors à la prochaine, c'est mieux » |
| **802.11r** | Fast BSS Transition | **Accélère la ré-authentification** lors du roaming (clés pré-négociées) | Le télépéage au lieu du ticket à chaque barrière |

**Point fondamental : c'est le CLIENT qui décide de roamer.** L'AP propose (k, v), accélère (r), mais un client têtu restera accroché à un AP lointain. D'où l'importance des réglages côté AP : seuils de déconnexion des clients faibles (débits de base, chap. 13), puissances équilibrées (chap. 40).

**802.11r en détail :** sans 802.11r, un roaming en WPA2-Enterprise = ré-authentification complète (EAP + 4-way handshake) = **200–500 ms** de coupure — audible en voix, visible en visio. Avec 802.11r : **< 50 ms** typique. En PSK, le gain est moindre mais réel.

**Prérequis 802.11r :** les AP doivent partager un **domaine de mobilité** (même groupe, même contrôleur/cloud). En mode Fat avec des AP indépendants, le 802.11r est limité — une raison de plus de préférer le **cloud** ou le **Fit** dès qu'il y a plusieurs AP et du roaming.

## 57. Fast roaming en pratique sur l'AP761

**Checklist roaming propre :**
- [ ] Même **SSID**, même **sécurité**, même **mot de passe** (ou même RADIUS) sur tous les AP.
- [ ] **802.11r activé** sur le profil de sécurité (tester avec le parc réel — certains vieux clients bugguent avec r, voir cas 6).
- [ ] **802.11k/v activés** : aide les clients modernes à choisir vite et bien.
- [ ] Recouvrement de couverture **15–20 %** à −67 dBm entre AP voisins (chap. 40).
- [ ] Canaux **différents** entre AP voisins (chap. 37–38) : un client qui roame change aussi de canal, c'est normal.
- [ ] **VLAN cohérents** : le même SSID doit donner le même VLAN sur tous les AP, sinon le client change de sous-réseau en roamant = sessions coupées.

**Tester le roaming (méthode terrain) :**
1. Connecter un smartphone au SSID, lancer un **ping continu** vers la passerelle (`ping -t`).
2. Marcher lentement d'un AP à l'autre en regardant l'écran.
3. Noter : à quel RSSI le client lâche l'AP1 ? Combien de pings perdus pendant la bascule ? (< 5 = bon, > 20 = problème).
4. Refaire avec un appel voix ou une visio : c'est le test qui compte vraiment.

**Seuils indicatifs :** un client bien réglé roame vers −70 dBm. S'il reste accroché à −80 dBm, c'est soit qu'il ne voit pas mieux (trou de couverture), soit qu'il est têtu (désactiver le Wi-Fi/réactiver pour forcer, ou ajuster les débits de base pour le « pousser » dehors).

## 58. Planification : combien d'AP pour quelle surface

Deux logiques de dimensionnement, à combiner :

**A. Dimensionnement par COUVERTURE (le signal) :**
Un AP761 couvre un secteur de 65° (chap. 9). En pratique :
| Environnement | Portée utile 5 GHz (indicative) | Portée utile 2.4 GHz (indicative) |
|---|---|---|
| Vue directe, peu d'obstacles | 80–150 m | 150–250 m |
| Cour avec arbres, mobilier | 40–80 m | 80–150 m |
| Zone avec bâtiments intermédiaires | 20–40 m | 40–80 m |

> Ce sont des ordres de grandeur terrain, pas des valeurs constructeur. Le seul chiffre constructeur approchant est la « portée 500 m » citée par un revendeur (**à vérifier sur la fiche du modèle exact** — c'est probablement une portée max en conditions idéales, pas une portée utile).

**B. Dimensionnement par CAPACITÉ (le débit) :**
L'AP761 accepte **1024 clients max** (512/radio) [constructeur], mais c'est un maximum d'association, pas de confort. La vraie limite, c'est l'**airtime** :
| Usage | Clients confortables par AP (ordre de grandeur) |
|---|---|
| Navigation légère, IoT | 100–200 |
| Usage bureautique / visio occasionnelle | 40–80 |
| Visio dense, voix sur Wi-Fi | 20–40 |

**Règle :** on dimensionne sur le **plus contraignant** des deux. Une cour de 200 personnes en visio = 4–5 AP pour la capacité, même si 2 AP suffiraient pour la couverture.

## 59. Règles de pouce densité / débit par usage

Le tableau à sortir en réunion quand on te demande « ça va tenir ? » :

