---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-16
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [1560, 1666]
sha256: aac5ea4e4999ed4f277d3bb0c7ccbb75c1281bbf3d670f9f015885d87db06815
---

# Guide ultra-complet — Huawei eKit AP761

**Méthode de classification (à faire avant toute contre-mesure) :**
1. Le WIDS signale un AP inconnu avec un SSID proche du tien ou branché sur ton réseau.
2. **Vérifier s'il est sur ton filaire :** chercher son adresse MAC dans les tables MAC de tes switches. Si oui = rogue naïf (ou pire) **interne** → intervention physique.
3. **S'il n'est pas sur ton filaire :** c'est un voisin (légitime) ou un attaquant externe. Observer : usurpe-t-il ton SSID ? Émet-il fort près de tes locaux ?
4. **Ne jamais contre-attaquer un AP voisin légitime** : le risque juridique et relationnel dépasse le bénéfice.

**Procédure terrain « rogue naïf » (le cas à 90 %) :**
1. Identifier le port du switch où apparaît sa MAC (table MAC du switch).
2. Remonter le brassage jusqu'à la prise murale / la zone.
3. Aller voir : 9 fois sur 10, c'est un boîtier personnel branché « pour dépanner ».
4. Débrancher, expliquer, noter l'incident.
5. **Traiter la cause :** pourquoi le salarié a branché ça ? (Zone non couverte ? Procédure trop lente pour avoir du Wi-Fi ?) — un rogue naïf est souvent un **signal de besoin non satisfait**.

## 65. Analyse de spectre (mode Fit uniquement)

**Fait constructeur important :** sur l'AP761, l'**analyse de spectre** (identification des sources d'interférences : Bluetooth, micro-ondes, téléphones sans fil, ZigBee, vidéo/audio sans fil, babyphones) n'est supportée qu'en **mode Fit** [datasheet]. En Fat ou Cloud, tu n'y as pas accès.

**Ce que ça apporte :** au lieu de voir « le canal est bruyant », tu vois « c'est un four à micro-ondes » ou « c'est du Bluetooth ». Ça change le diagnostic : on ne change pas de canal contre un micro-ondes (il pollue large), on **déplace l'AP** ou on **blinde la source**.

**Si tu es en Cloud/Fat et que tu soupçonnes une interférence non-Wi-Fi :**
1. Méthode pauvre : éteindre successivement les suspects (le micro-ondes du food-truck, la caméra sans fil du voisin) et observer le bruit.
2. Méthode smartphone : certaines apps montrent la « qualité » du canal au-delà des réseaux Wi-Fi visibles.
3. Méthode pro : louer/emprunter un analyseur de spectre portable pour une journée (rentable sur un gros site qui rame).
4. **Basculer temporairement un AP en Fit** vers un AC de test si tu en as un — lourd, mais c'est la seule analyse de spectre native.

**Sources d'interférences typiques en extérieur (mémo) :** fours à micro-ondes, caméras vidéo sans fil 2.4 GHz, Bluetooth (enceintes, kits), ZigBee (domotique), téléphones DECT (normalement 1.9 GHz en Europe — vérifier si modèle exotique), drones FPV en 5.8 GHz (!), liaisons vidéo de sécurité.

## 66. Supervision via la plateforme cloud eKit

En mode cloud, la supervision est **gratuite et intégrée** : c'est l'argument n° 1 du cloud pour un chef de service multi-sites.

**Ce que la console cloud montre (selon version — à vérifier sur la fiche du modèle exact) :**
- **État des AP :** en ligne/hors ligne, uptime, version firmware, utilisation CPU/mémoire.
- **Clients :** nombre par AP/SSID, débits, RSSI, historique.
- **Radio :** canal utilisé, largeur, puissance, taux d'utilisation (channel utilization).
- **Alertes :** AP hors ligne, rogue détecté, interférences, seuil de clients dépassé.
- **Cartographie :** position des AP sur plan (si renseignée).

**Organisation multi-sites :**
```
Compte principal (toi)
├── Site « Siège » (3 AP)
├── Site « Usine Nord » (8 AP)
└── Site « Dépôt Sud » (2 AP)
Techniciens : accès par site, pas au compte global.
```

**Les alertes à configurer en priorité :**
1. **AP hors ligne** (le plus critique — souvent = coupure PoE ou câble arraché).
2. **Rogue AP détecté** (chap. 64).
3. **Utilisation canal > 70 %** prolongée (signe de saturation ou d'interférence).
4. **Firmware obsolète** (rappel de mise à jour).

**Limite honnête :** la supervision cloud dépend d'Internet et des serveurs Huawei. Pour les sites critiques, **doubler avec du SNMP local** (chap. 68) : en cas de coupure Internet, le cloud est aveugle mais ton superviseur local voit toujours.

## 67. Indicateurs Wi-Fi à surveiller (tableau)

Le tableau de bord minimal d'un parc Wi-Fi extérieur. Relever **une fois par semaine** en routine, **en continu** sur les sites critiques.

| Indicateur | Où le voir | Seuil d'alerte | Ce que ça veut dire si dépassé |
|---|---|---|---|
| AP hors ligne | Cloud / SNMP | > 0 | PoE, câble, AP mort (cas 1) |
| Channel utilization 5 GHz | Cloud / CLI | > 70 % prolongé | Saturation ou interférence (cas 4, 7) |
| Channel utilization 2.4 GHz | Cloud / CLI | > 60 % prolongé | Le 2.4 GHz sature vite — migrer vers 5 GHz |
| Clients par AP | Cloud / CLI | > 80 % du dimensionnement | Ajouter un AP ou rééquilibrer (chap. 58) |
| RSSI moyen des clients | Cloud | < −70 dBm | Couverture limite ou clients trop loin |
| Taux de réessais (retries) | CLI / SNMP | > 15–20 % | Interférences ou clients faibles |
| CRC errors filaire | Switch / CLI | > 0 croissant | Câble/connecteur défectueux (chap. 24) |
| CPU / mémoire AP | Cloud / CLI | > 80 % prolongé | Surcharge (trop de clients/fonctions) |
| Uptime | Cloud / SNMP | Redémarrage inattendu | Surchauffe, PoE instable (cas 14, cas 1) |
| Rogues détectés | WIDS / Cloud | > 0 non classifié | À classifier (chap. 64) |

**Rituel hebdo (15 min) :** ouvrir la console, vérifier les AP hors ligne, les alertes, un coup d'œil aux utilizations canal. **Rituel mensuel (1 h) :** tendances (la cour se remplit-elle ? un canal se dégrade-t-il ?), revue des firmwares, mise à jour du dossier de site.

## 68. SNMP : v1/v2c/v3 en mode Fat

**Supporté en mode Fat : SNMP v1/v2c/v3** [constructeur]. C'est la porte d'entrée vers ton superviseur (Zabbix, PRTG, Centreon — tu as un guide Zabbix : `~/workspace/user/files/zabbix_guide.md`).

**Configuration (principe CLI) :**
```
[AP761] snmp-agent
[AP761] snmp-agent sys-info version v3
[AP761] snmp-agent group v3 GRP-ADMIN privacy
[AP761] snmp-agent usm-user v3 USER-SUP auth-mode sha CleAuthForte_2026! priv-mode aes128 ClePrivForte_2026!
[AP761] snmp-agent group v3 GRP-ADMIN acl 2000
# (ACL 2000 = n'autoriser que le superviseur — à créer)
```

**Règles :**
- **Toujours SNMPv3** (auth + priv) : le v1/v2c envoie la community **en clair** — sur un réseau extérieur, c'est une faute.
- **Restreindre par ACL** : seul le superviseur interroge l'AP.
- **Ce qu'on supervise :** uptime, état des interfaces, clients associés, utilisation CPU/mémoire, traps (AP down, rogue détecté).
- Les **OID spécifiques** (clients Wi-Fi, stats radio) dépendent de la MIB Huawei : charger la MIB du firmware **exact** dans le superviseur — à vérifier sur la fiche du modèle exact.

**Traps vs polling :** le polling (toutes les 1–5 min) pour les tendances, les **traps** pour l'immédiat (AP down). Configurer les deux : un AP qui ne répond plus au polling ET n'envoie plus de traps = vraiment mort, pas juste un trou SNMP.

## 69. Syslog : centraliser les journaux

Les logs de l'AP sont la **boîte noire** : en dépannage, c'est souvent là que se trouve l'explication (désauthentifications en boucle, changements de canal DFS, pertes CAPWAP).

**Configuration (principe CLI) :**
```
[AP761] info-center enable
[AP761] info-center loghost 192.168.10.100 facility local7
[AP761] info-center source default channel loghost log level informational
```

