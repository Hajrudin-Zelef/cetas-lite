---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-10
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei", "United States"]
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [901, 1034]
sha256: 24a7b723d21f022b3ee299892ff65dbeda592d9e68164eed831317ef96b40f87
---

# --- Port uplink en trunk ---
[AP761] interface gigabitethernet 0/0/0
[AP761-GigabitEthernet0/0/0] port link-type trunk
[AP761-GigabitEthernet0/0/0] port trunk pvid vlan 10
[AP761-GigabitEthernet0/0/0] port trunk allow-pass vlan 10 20 30
[AP761-GigabitEthernet0/0/0] quit

# --- WLAN : profil SSID ---
[AP761] wlan
[AP761-wlan-view] ssid-profile name SP-CORP
[AP761-wlan-ssid-prof-SP-CORP] ssid ENTREPRISE
[AP761-wlan-ssid-prof-SP-CORP] quit

# --- WLAN : profil sécurité WPA2-PSK ---
[AP761-wlan-view] security-profile name SEC-CORP
[AP761-wlan-sec-prof-SEC-CORP] security wpa2 psk pass-phrase ChangeMoi_2026! aes
[AP761-wlan-sec-prof-SEC-CORP] quit

# --- WLAN : profil VAP (SSID + sécu + VLAN + forwarding) ---
[AP761-wlan-view] vap-profile name VAP-CORP
[AP761-wlan-vap-prof-VAP-CORP] ssid-profile SP-CORP
[AP761-wlan-vap-prof-VAP-CORP] security-profile SEC-CORP
[AP761-wlan-vap-prof-VAP-CORP] service-vlan vlan-id 20
[AP761-wlan-vap-prof-VAP-CORP] forward-mode direct-forward
[AP761-wlan-vap-prof-VAP-CORP] quit

# --- Domaine réglementaire (France = FR) ---
[AP761-wlan-view] regulatory-domain-profile name DOM-FR
[AP761-wlan-view] ap-group name default
[AP761-wlan-ap-group-default] regulatory-domain-profile DOM-FR

# --- Attacher le VAP aux deux radios ---
[AP761-wlan-ap-group-default] vap-profile VAP-CORP wlan 1 radio 0
[AP761-wlan-ap-group-default] vap-profile VAP-CORP wlan 1 radio 1
[AP761-wlan-ap-group-default] quit
[AP761-wlan-view] quit

# --- NTP + sauvegarde ---
[AP761] ntp-service unicast-server 192.168.10.1
[AP761] quit
<AP761> save
```

**Vérification :**
```
<AP761> display wlan vap all       # le VAP doit être « up »
<AP761> display current-configuration | include ssid
```
Puis test avec un smartphone : association, IP obtenue, navigation, débit (speedtest local de préférence — un iperf vers un serveur du LAN, pas vers Internet, pour isoler le Wi-Fi).

> 📌 **Avertissement version :** les noms de vues (`wlan-ssid-prof-…`, `vap-profile`, `forward-mode`) sont la syntaxe classique des AP Huawei VRP. Selon ta version logicielle, certains mots-clés peuvent différer : **valide chaque bloc avec `?` avant de l'appliquer**, et consulte le guide de configuration de ta version.

## 35. Sauvegarde de la configuration initiale

**Règle d'or :** une configuration non sauvegardée **n'existe pas**. Le jour où l'AP prend la foudre, tu dois pouvoir le remplacer en 30 minutes avec un fichier.

**En mode Fat (CLI) :**
```
<AP761> save                          # écrit dans la flash (vrpcfg.zip)
<AP761> backup configuration to 192.168.10.100   # vers serveur TFTP/SFTP (selon version)
# ou via SFTP depuis un poste :
sftp admin@192.168.10.11
sftp> get vrpcfg.zip ./sauvegardes/AP761-Cour-Est_2026-09-27.zip
```

**En mode Cloud :** la configuration vit dans le cloud — elle est versionnée par la plateforme. **Exporter quand même** une copie après chaque changement majeur (fonction d'export du portail — à vérifier sur la fiche du modèle exact), et noter la version firmware associée.

**Convention de nommage des sauvegardes :**
```
sauvegardes/
├── AP761-Cour-Est_2026-09-27_v1-initiale.zip
├── AP761-Cour-Est_2026-10-15_v2-wpa3.zip
└── AP761-Parking_2026-09-27_v1-initiale.zip
```
Un fichier = `NomAP_AAAA-MM-JJ_vN-motif`. Stockage : dossier partagé d'équipe + copie hors site. **Tester la restauration une fois** (chap. 73) : une sauvegarde jamais testée est une promesse, pas une sécurité.

## 36. Configuration radio : concepts de base

Avant de toucher aux réglages, les 5 paramètres qui gouvernent une radio Wi-Fi, et ce qu'ils font vraiment :

| Paramètre | Ce que c'est | Effet si mal réglé |
|---|---|---|
| **Canal** | La fréquence d'émission | 2 AP sur le même canal = partage de l'airtime, débit divisé |
| **Largeur de canal** | 20 / 40 / 80 MHz : la « largeur de l'autoroute » | Trop large = plus de débit mais plus d'interférences et moins de canaux dispo |
| **Puissance TX** | La « voix » de l'AP (en dBm) | Trop fort = l'AP crie loin mais n'entend pas les clients faibles (asymétrie) ; trop faible = trous de couverture |
| **Débits de base** | Les débits minimums autorisés | Trop bas = clients zombies qui monopolisent l'airtime (chap. 13) |
| **Domaine réglementaire** | Le pays (FR, etc.) | Mauvais pays = canaux interdits utilisés ou puissance non conforme |

**Principe cardinal de la radio Wi-Fi :** c'est un médium **partagé et à écoute avant émission** (CSMA/CA). Deux AP sur le même canal ne se « brouillent » pas comme deux radios FM : ils **se partagent** le temps de parole. Le but du plan de canaux n'est pas d'éviter tout chevauchement à tout prix, c'est de **minimiser le partage** là où le trafic est fort.

**Ordre de configuration recommandé :**
1. Domaine réglementaire (chap. 30/34) — en premier, toujours.
2. Canaux (chap. 37–38).
3. Largeurs (chap. 39).
4. Puissances (chap. 40).
5. Débits de base et band steering (chap. 13, 42).

## 37. Plan de canaux 2.4 GHz — méthode terrain

La bande 2.4 GHz (2.412–2.472 GHz en Europe) est **minuscule** : 13 canaux de 20 MHz qui se chevauchent, dont **seulement 3 ne se recouvrent pas : 1, 6, 11**. C'est la règle la plus importante du Wi-Fi, et elle n'a pas changé depuis 1999.

**Le plan en 20 MHz (le seul raisonnable en 2.4 GHz) :**
```
Canaux :  1   2   3   4   5   6   7   8   9   10  11  12  13
               └─────┘       └─────┘       └─────────┘
              chevauchement    1, 6, 11 = les 3 canaux propres
```

**Méthode terrain (sans analyseur de spectre) :**
1. Avec un smartphone + app d'analyse Wi-Fi (ex. « WiFi Analyzer », open source), faire un scan depuis 3–4 points de la zone.
2. Noter les réseaux voisins forts (RSSI > −80 dBm) et leurs canaux.
3. Choisir pour chaque AP le canal **1, 6 ou 11** le moins occupé par les voisins **forts**.
4. Deux AP761 à portée l'un de l'autre : **jamais le même canal** — alterner 1 / 6 / 11.

**Le 40 MHz en 2.4 GHz : à proscrire en extérieur partagé.** Le 40 MHz en 2.4 GHz occupe 2 des 3 canaux propres : tu doubles ton débit théorique mais tu divises par deux la capacité du voisinage — et en extérieur, le voisinage (lotissements, commerces) est partout. **Règle : 2.4 GHz = 20 MHz, toujours.** (Le datasheet confirme que l'AP761 supporte le 40 MHz en 2.4 GHz [constructeur] — le supporter n'est pas une raison de l'utiliser.)

**Canaux 12 et 13 :** autorisés en Europe, interdits aux USA. Si tes clients sont des objets connectés « US-first » (certaines caméras, gadgets), ils ne verront pas les canaux 12/13. **Reste sur 1/6/11** sauf raison précise.

## 38. Plan de canaux 5 GHz — méthode terrain

La bande 5 GHz européenne (5.18–5.825 GHz selon le datasheet [constructeur]) offre **beaucoup plus de canaux** — c'est là que se joue le débit.

**Canaux 20 MHz usuels en Europe (liste indicative — à vérifier sur la fiche du modèle exact, tableau « Country Code & Channel Compliance ») :**
```
U-NII-1 (sans DFS) : 36, 40, 44, 48
U-NII-2 (DFS)      : 52, 56, 60, 64
U-NII-2e (DFS)     : 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140
U-NII-3            : 149, 153, 157, 161, 165 (selon pays)
```

**Canaux non-DFS (36–48) :** pas de détection radar, pas de coupure surprise. **À privilégier** pour les usages critiques (voix, contrôle d'accès). Inconvénient : seulement 4 canaux 20 MHz (ou 2× 40 MHz, ou 1× 80 MHz) — vite saturés si les voisins s'y entassent aussi.

**Canaux DFS (52–140) :** beaucoup de spectre, mais l'AP doit **écouter 60 s avant d'émettre** (CAC — Channel Availability Check) et **quitter le canal** s'il détecte un radar (chap. 41, cas 8). À réserver aux usages non critiques ou quand les canaux propres sont saturés.

