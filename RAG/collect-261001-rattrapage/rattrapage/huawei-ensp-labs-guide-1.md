---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-1
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Huawei"]
dates: []
keywords: ["amd", "arr", "attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [1, 109]
sha256: 07707579ed86f06c120aba872a4ff6a516974f75053a2341d2b8ce05af83678e
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

> **Public cible** : équipe systèmes & énergies manipulant en production des AP361, AP761, switches S310, routeurs AR720 et pare-feu USG6000.
> **Objectif pédagogique** : faire monter l'équipe en compétence sur la CLI VRP Huawei (concepts identiques sur AR/S/USG), le switching, le routage, le NAT, la sécurité et le WLAN, sans toucher au réseau de production.
> **Outil** : Huawei eNSP (Enterprise Network Simulation Platform), simulateur gratuit. Non maintenu officiellement par Huawei, mais reste l'outil standard de lab pour les équipements de campus/PME.
> **Langue** : français. Les commandes VRP sont identiques quel que soit le modèle.

---

## Sommaire

- [1. Installation d'eNSP sur Windows](#1-installation-densp-sur-windows)
- [2. Prise en main d'eNSP](#2-prise-en-main-densp)
- [3. TP1 — VLAN de base sur S3700](#3-tp1--vlan-de-base-sur-s3700)
- [4. TP2 — Trunk 802.1Q entre deux switches](#4-tp2--trunk-8021q-entre-deux-switches)
- [5. TP3 — STP/RSTP : boucles et élection du root](#5-tp3--stprstp--boucles-et-élection-du-root)
- [6. TP4 — Serveur DHCP sur routeur AR](#6-tp4--serveur-dhcp-sur-routeur-ar)
- [7. TP5 — Routage statique et floating static](#7-tp5--routage-statique-et-floating-static)
- [8. TP6 — OSPF mono-zone](#8-tp6--ospf-mono-zone)
- [9. TP7 — NAT : NAPT et port mapping](#9-tp7--nat--napt-et-port-mapping)
- [10. TP8 — PPPoE client et serveur](#10-tp8--pppoe-client-et-serveur)
- [11. TP9 — IPSec site-à-site](#11-tp9--ipsec-site-à-site)
- [12. TP10 — WLAN : AC + AP en mode Fit](#12-tp10--wlan--ac--ap-en-mode-fit)
- [13. TP11 — USG : zones, politiques, NAT sortant](#13-tp11--usg--zones-politiques-nat-sortant)
- [14. TP12 — Dépannage guidé : la topologie cassée](#14-tp12--dépannage-guidé--la-topologie-cassée)
- [15. Quiz final (10 questions corrigées)](#15-quiz-final-10-questions-corrigées)
- [16. Glossaire](#16-glossaire)
- [17. Pour aller plus loin](#17-pour-aller-plus-loin)
- [Annexe A — Aide-mémoire des commandes VRP](#annexe-a--aide-mémoire-des-commandes-vrp)
- [Annexe B — Modèle de fiche de lab vierge](#annexe-b--modèle-de-fiche-de-lab-vierge)

**Conventions utilisées dans ce guide :**

- Les blocs ` ``` ` contiennent des commandes VRP à saisir sur l'équipement indiqué.
- `<Huawei>` = vue utilisateur (user view), `[Huawei]` = vue système (system view), `[Huawei-GigabitEthernet0/0/1]` = vue d'interface.
- `[R1]`, `[SW1]`, `[FW1]`, `[AC1]` indiquent l'équipement concerné.
- `#` en début de ligne = commentaire explicatif, jamais une commande.
- Les adresses IP utilisées suivent le plan : `192.168.x.0/24` par site, avec `.254` = passerelle (routeur).

---

## 1. Installation d'eNSP sur Windows

### 1.1. Ce qu'est eNSP (et ce qu'il n'est pas)

eNSP (Enterprise Network Simulation Platform) est le simulateur réseau gratuit de Huawei, équivalent de Cisco Packet Tracer / GNS3. Il émule de vrais firmwares VRP :

- **Routeurs** : AR201, AR1220, AR2220, AR2240, AR3260 (modèles ARxxxx, proches dans l'esprit de vos AR720).
- **Switches** : S3700 (L2), S5700 (L3, routage inter-VLAN), S9300, CE6800, CE12800.
- **Sécurité** : USG5500, USG6000V (logique identique à vos USG6000, commandes interzone).
- **WLAN** : AC6005/AC6605 + points d'accès AP (AP6010DN, AP2050...) — logique Fit AP identique à vos AP361/AP761.
- **Terminaux** : PC, Laptop, Server, Client, STA (station Wi-Fi), Cloud (pont vers le réseau réel de la machine hôte).

**Limites à connaître avant de commencer :**

- eNSP **n'est plus officiellement maintenu** par Huawei : la dernière version courante est la **V1R2.00.510** (aussi appelée 1.3.00.100). Il n'y aura pas de correctif futur.
- Il **fonctionne uniquement sous Windows** (7/8/10/11).
- Il repose sur **VirtualBox** pour la virtualisation et **WinPcap** pour la capture.
- Les équipements émulés correspondent à des gammes des années 2010 : les concepts (VLAN, OSPF, NAT, zones USG, WLAN) sont **identiques** à vos équipements récents, mais les interfaces physiques et certains détails (noms de vues, limites) diffèrent.
- eNSP ne supporte pas : le PoE réel, le Wi-Fi 6/7 (les AP émulés sont du Wi-Fi 4/5), le stacking iStack complet, EVPN, SD-WAN.

### 1.2. Prérequis matériels et logiciels

| Élément | Recommandation | Commentaire |
|---|---|---|
| OS | Windows 10 ou 11 (64 bits) | Windows 7/8 fonctionnent aussi |
| CPU | 4 cœurs minimum, virtualisation VT-x/AMD-V activée dans le BIOS | Vérifiable avec `systeminfo` (rechercher "Virtualisation activée") |
| RAM | 8 Go minimum, 16 Go conseillés | Chaque équipement émulé consomme 200 à 500 Mo |
| Disque | 5 Go libres (SSD conseillé) | Les images sont incluses dans le package |
| VirtualBox | **5.2.x** (ex. 5.2.44) | Les versions 6.x/7.x provoquent des erreurs de démarrage des nœuds ; eNSP vérifie la version et refuse parfois les versions "trop basses" ou "trop hautes" |
| WinPcap | **4.1.3** | Indispensable pour la capture Wireshark ; Npcap *peut* fonctionner en mode "compatibilité WinPcap" mais WinPcap 4.1.3 reste la valeur sûre |
| Wireshark | Dernière version stable | Optionnel mais très utile pour les TP (capture sur les liens) |
| Droits | Compte administrateur | Installation + exécution |

> **Point d'attention 2025-2026 (vérifié)** : la communauté contourne les problèmes d'installation de deux façons :
> 1. **Couple éprouvé** : Windows 10/11 + VirtualBox 5.2.x + WinPcap 4.1.3 + eNSP V1R2.00.510, le tout lancé **en administrateur**. C'est la combinaison la plus fiable.
> 2. **ISO pré-configurée communautaire** : certains formateurs diffusent une VM VirtualBox contenant eNSP déjà installé (à importer en `.ova`). Pratique pour déployer vite sur plusieurs PC d'équipe, mais à n'utiliser que depuis une source de confiance.
> Les symptômes classiques d'une mauvaise installation : équipements affichés `???` dans la topologie, nœuds qui restent rouges, erreur "VirtualBox version too low" (solution : réinstaller la 5.2.x, chemin d'installation sans caractères accentués ni espaces exotiques).

### 1.3. Conflit Hyper-V / WSL2 : à désactiver impérativement

VirtualBox 5.2.x **ne cohabite pas** avec l'hyperviseur natif de Windows (Hyper-V, WSL2, Credential Guard, Device Guard, isolation du noyau). Symptôme : les équipements eNSP démarrent puis s'arrêtent, ou VirtualBox affiche des erreurs VT-x.

**Procédure (en administrateur) :**

```
# 1. Désactiver les fonctionnalités Windows (GUI) :
# Panneau de configuration > Programmes > Activer/désactiver des fonctionnalités Windows
# Décocher : Hyper-V, Plateforme d'hyperviseur Windows, Plateforme de machine virtuelle,
#            Sous-système Windows pour Linux (si WSL2 n'est pas nécessaire)

# 2. En invite de commandes administrateur, forcer la désactivation de l'hyperviseur :
bcdedit /set hypervisorlaunchtype off

# 3. Redémarrer le PC.
```

> Si vous avez besoin de WSL2/Docker sur la même machine, sachez que c'est **l'un ou l'autre** avec VirtualBox 5.2. Pour un poste de formation dédié eNSP, désactivez tout : c'est le plus simple.

### 1.4. Installation pas à pas

1. **Installer VirtualBox 5.2.x** (exécuter le setup en administrateur). Laisser les options par défaut, y compris l'interface réseau hôte (Host-Only) — elle sert au "Cloud" d'eNSP.
2. **Installer WinPcap 4.1.3** (setup administrateur). Si un antivirus bloque l'installation du pilote, le désactiver temporairement.
3. **Installer eNSP V1R2.00.510** (setup administrateur). Choisir un chemin simple, ex. `C:\eNSP` (éviter les accents et les chemins trop longs).
4. **Installer Wireshark** (optionnel mais recommandé).
5. **Redémarrer** le PC après l'ensemble.

### 1.5. Premier lancement et vérifications

