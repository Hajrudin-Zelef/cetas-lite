---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-3
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Huawei"]
dates: ["2026-09-27"]
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [186, 290]
sha256: a6706f3a334668d3744c1c05285682e0028c76a6b0209cb8468c2b40299e07ae
---

# Guide ultra-complet — Huawei eKit AP761

| Modèle | Usage | Wi-Fi | Débit max | Ports | Conso max |
|---|---|---|---|---|---|
| AP160 / AP162E | Chambre d'hôtel, mur | Wi-Fi 6 | 1.775 / 2.975 Gbps | GE | faible |
| AP263 | Bureau PME | Wi-Fi 6 | 2.975 Gbps | 2× GE | 12 W |
| **AP361** | **Bureau PME (indoor)** | **Wi-Fi 6** | **1.775 Gbps** | 1× GE | **8.8 W** [constructeur] |
| AP362 / AP362E | Bureau dense (indoor) | Wi-Fi 6 | 2.975 Gbps | 1× GE / 1× 2.5GE | 11.2 W |
| AP661 | Indoor haute densité | Wi-Fi 6 | 6.575 Gbps | 1× 2.5GE + 1× GE | 21.2 W |
| **AP761** | **Extérieur (outdoor)** | **Wi-Fi 6** | **1.775 Gbps** | **1× GE + 1× SFP** | **17.7 W** [constructeur] |
| AP371 | Indoor | **Wi-Fi 7** | 3.57 Gbps | 1× 2.5GE | à vérifier sur la fiche du modèle exact |
| AP771 | Extérieur | **Wi-Fi 7** | 3.57 Gbps | à vérifier sur la fiche du modèle exact | à vérifier sur la fiche du modèle exact |
| AP772E | Extérieur | **Wi-Fi 7** | 6.45 Gbps | 1× 2.5GE + 1× 10GE SFP+ | à vérifier sur la fiche du modèle exact |
| AP673 | Indoor tri-bande | **Wi-Fi 7** | à vérifier sur la fiche du modèle exact | à vérifier sur la fiche du modèle exact | à vérifier sur la fiche du modèle exact |

**Lecture terrain :** l'AP761 est *le* point d'accès extérieur Wi-Fi 6 de la gamme PME. Si ton besoin est « couvrir une cour, un parking, une terrasse en Wi-Fi 6 fiable », c'est lui. Si ton besoin est « Wi-Fi 7 en extérieur », ce n'est pas lui : voir chapitres 21 et 99.

## 3. Contenu de la boîte et prérequis chantier

**Dans la boîte (kit standard, à vérifier sur la fiche du modèle exact pour ta région) :**
- L'AP761 lui-même (boîtier métal blanc, 200 × 200 × 69 mm, 1.91 kg) [constructeur].
- Kit de montage mural/mât (étriers, visserie de base).
- Cache presse-étoupe pour le passage du câble Ethernet.
- Guide de démarrage rapide + carte avec QR code d'onboarding cloud.

**Ce qui n'est généralement PAS fourni :**
- L'injecteur PoE ou le switch PoE (à prévoir, chap. 25).
- Le câble Ethernet extérieur (à prévoir, chap. 24).
- Le mât ou le bras de déport (à prévoir selon le site).
- Le parafoudre Ethernet / la barrette de terre (à prévoir, chap. 26).

**Prérequis chantier avant de monter :**
- [ ] Point de fixation identifié (mur porteur ou mât), dégagé, accessible à la nacelle/échelle en sécurité.
- [ ] Chemin de câble défini jusqu'au local technique (distance, traversées).
- [ ] Source PoE identifiée : switch PoE 802.3at minimum recommandé (voir chap. 25).
- [ ] Plan d'adressage : VLAN, sous-réseau de management, passerelle, DNS, NTP.
- [ ] Compte HUAWEI eKit / accès à la plateforme cloud si mode cloud (chap. 30).
- [ ] Smartphone avec l'app HUAWEI eKit installée (Android/iOS) pour l'onboarding.
- [ ] Autorisations : copropriété, mairie (façade classée ?), consignation électrique si intervention sur TGBT.

## 4. Fiche technique constructeur — tableau de synthèse

Toutes les valeurs ci-dessous viennent du **datasheet officiel Huawei eKitEngine AP761** (vérifié le 27/09/2026). Les valeurs non vérifiées sont marquées explicitement.

| Rubrique | Valeur constructeur |
|---|---|
| Référence | 02355VFB |
| EAN | 6901443451289 |
| Standard Wi-Fi | IEEE 802.11a/b/g/n/ac/ac Wave 2/**ax (Wi-Fi 6)** |
| Bandes / MIMO | 2.4 GHz 2×2 + 5 GHz 2×2, simultanées |
| Débit max agrégé | **1.775 Gbps** (0.575 Gbps @ 2.4 GHz + 1.2 Gbps @ 5 GHz) |
| Largeurs de canal | 20 / 40 / 80 MHz |
| Modulation max | 1024-QAM (compatible 256/64/16-QAM, QPSK, BPSK) |
| Ports filaires | **1× GE RJ45 (PoE-In)** + **1× SFP GE optique** — combo : un seul actif à la fois, **l'optique est prioritaire** |
| Alimentation | PoE **IEEE 802.3at/af** — en 802.3af, **fonctions restreintes** (voir chap. 9 du dépannage, cas 9) |
| Consommation max | **17.7 W** |
| Dimensions (H×L×P) | **69 × 200 × 200 mm** |
| Poids | **1.91 kg** |
| Antennes | Intégrées **directionnelles** : 2.4 GHz **10 dBi** (65° H / 40° V), 5 GHz **11 dBi** (65° H / 20° V), BLE 5 dBi |
| Puissance TX max combinée | 2.4 GHz : **28 dBm** ; 5 GHz : **27 dBm** (sous réserve réglementation locale) |
| Pas de réglage puissance | 1 dBm |
| Clients max | **1024** (512 par radio) — le réel dépend de l'environnement |
| SSID max | **16 par radio** |
| Bluetooth | **BLE 5.2** (O&M via app CloudCampus/eKit) |
| LED | États : power-on, démarrage, fonctionnement, alarme, défaut |
| Température de fonctionnement | **−40 °C à +65 °C** (dératage : −1 °C par 300 m au-delà de 1800 m) |
| Température de stockage | −40 °C à +85 °C |
| Humidité | 0 % à 100 % HR |
| Indice de protection | **IP68** |
| Altitude | −60 m à +5000 m ; pression 53 à 106 kPa |
| Surtension ports Ethernet | **6 kV** |
| Modes | **Fit, Fat, cloud management** |
| Sécurité Wi-Fi | WEP, WPA/WPA2-PSK, WPA/WPA2-Enterprise, **WPA3-SAE**, WPA3-Enterprise, WAPI, 802.11w |
| Fonctions radio | MU-MIMO, OFDMA, BSS coloring, TWT, beamforming, MRC, STBC, CDD/CSD, LDPC, A-MPDU/A-MSDU, DFS, SST, U-APSD, ACC, CAC |
| Réseau | 802.1Q (VLAN par SSID), DHCP client, NAT, ACL, DHCP snooping, DAI, IPSG, mDNS, IPv6 SAVI |
| Management | Web (HTTP/HTTPS), Telnet/STelnet (SSHv2), SFTP, SNMP v1/v2/v3 (mode Fat), NTP, CAPWAP, PnP |
| Montage | Mural / sur mât (kit fourni) |

> 📌 **Points à vérifier sur la fiche du modèle exact** (informations vues chez des revendeurs mais non confirmées par le datasheet consulté) : présence du **port USB** (extension IoT ZigBee/RFID), portée annoncée de **500 m**, puissance TX « ajustée » à 23 dBm par radio chez certains revendeurs, absence de **Mesh**.

## 5. Hardware : boîtier, dimensions, robustesse

L'AP761 est un pavé métallique blanc de **200 × 200 mm** pour **69 mm** d'épaisseur, **1.91 kg** [constructeur]. Ce n'est pas un AP de bureau : c'est un boîtier métal IP68 pensé pour rester dehors des années.

**Ce que ça change sur le terrain :**
- **1.91 kg + prise au vent** : sur mât, prévoir un collier sérieux et un mât qui ne fléchit pas. Un AP qui oscille au vent = des variations de RSSI côté clients fixes.
- **Boîtier métal = dissipateur** : la chaleur s'évacue par le boîtier. Ne jamais l'enfermer dans un coffret étanche non ventilé « pour le protéger » — tu le cuisinerais (voir cas 14).
- **IP68** : immersion temporaire supportée sur le papier. En pratique : ce qui tue les AP extérieurs, ce n'est pas la pluie directe, c'est l'eau qui **stagne** et remonte par capillarité dans un presse-étoupe mal serré, ou la condensation interne sur des cycles chaud/froid. Règle d'or : **presse-étoupe vers le bas**, boucle d'égouttage sur le câble (goutte d'eau), jamais de gaine qui fait entonnoir vers l'AP.
- **Plage −40 °C à +65 °C** [constructeur] : tient le plein soleil méditerranéen comme l'hiver. Au-delà de 1800 m d'altitude, dératage de 1 °C par 300 m — à noter pour les sites de montagne.
- **6 kV de protection surtension sur les ports Ethernet** [constructeur] : c'est une protection interne, pas un parafoudre. Elle encaisse les transitoires, pas la foudre directe. Chapitre 26 pour la vraie protection.

**Checklist réception matériel :**
- [ ] Boîtier sans choc ni déformation (un coin enfoncé = joint IP68 compromis).
- [ ] Joints et presse-étoupe présents et souples.
- [ ] Kit de montage complet (étriers, vis, colliers).
- [ ] Étiquette lisible : modèle AP761, numéro de série, adresse MAC, QR code d'onboarding.
- [ ] Photo de l'étiquette avant montage — quand l'AP est à 8 m de haut, tu seras content de l'avoir.

## 6. Hardware : ports filaires (GE RJ45 + SFP, mode combo)

L'AP761 a **deux ports filaires** [constructeur] :

| Port | Type | Rôle |
|---|---|---|
| GE / PoE_IN | RJ45 10/100/1000 Mbps | Uplink cuivre + **alimentation PoE** |
| SFP | Cage SFP 1 Gbps optique | Uplink fibre optique |

