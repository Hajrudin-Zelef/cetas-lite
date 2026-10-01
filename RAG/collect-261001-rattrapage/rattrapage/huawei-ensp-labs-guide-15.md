---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-15
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attribution", "valuation"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [2177, 2287]
sha256: 03e57b3aad062e77aaa842ba5d2ec95ec09dbe9ff7653ab006e4ee9a86abf2be
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

1. `ipconfig /all` : adresse, **masque**, **passerelle**, **DNS** — une IP sans passerelle/DNS = pas d'Internet (TP4).
2. Ping de la **passerelle** (couche 3 locale OK ?), puis ping d'une IP publique (routage/NAT OK ?), puis ping d'un **nom de domaine** (DNS OK ?).
3. Côté réseau : `display ip routing-table` (défaut présente ?), `display nat session all` (translation active ?), `display firewall session table` (politique USG ?).
4. Remonter par couches (méthode du TP12) sans sauter d'étape.
</details>

---

## 16. Glossaire

| Terme | Définition |
|---|---|
| **AC** (Access Controller) | Contrôleur WLAN centralisant la gestion des AP en mode Fit. |
| **ACL** (Access Control List) | Liste de règles de filtrage (permit/deny) ; sert aussi à sélectionner du trafic (NAT, IPSec, QoS). |
| **AP Fit / Fat** | Fit : AP "léger" piloté par un AC (tunnel CAPWAP). Fat : AP autonome. |
| **BDR** (Backup Designated Router) | Routeur OSPF de secours du DR sur un segment multi-accès. |
| **BPDU** | Trames d'échange STP/RSTP entre switches (élection root, états de ports). |
| **Broadcast domain** | Ensemble d'équipements recevant les diffusions de niveau 2 ; un VLAN = un broadcast domain. |
| **CAPWAP** | Protocole de tunnel entre AP Fit et AC (contrôle + données selon le mode). |
| **CHAP** | Protocole d'authentification PPP par challenge/response (ne transmet pas le mot de passe en clair). |
| **CIST** | Instance STP unique couvrant tous les VLAN (mode de base sur eNSP). |
| **DHCP DORA** | Discover, Offer, Request, Ack : les 4 étapes d'attribution d'adresse. |
| **DR** (Designated Router) | Routeur OSPF élu sur un segment multi-accès pour centraliser les échanges. |
| **Easy IP** | NAPT utilisant l'adresse IP de l'interface de sortie (cas standard PME). |
| **Edge port** | Port STP/RSTP vers un équipement terminal, passant immédiatement en forwarding (équivalent du portfast). |
| **ESP** | Protocole 50 : charge utile chiffrée d'IPSec (phase 2). |
| **Floating static** | Route statique de backup avec une préférence élevée, inactive tant que la route principale existe. |
| **Forward-mode** | Mode de transfert du trafic Wi-Fi : `direct-forward` (bridgé localement par l'AP) ou `tunnel-forward` (remonté à l'AC). |
| **IKE** | Protocole de négociation de la phase 1 IPSec (UDP 500, UDP 4500 avec NAT-T). |
| **Interesting traffic** | Trafic sélectionné par ACL pour être protégé par le tunnel IPSec. |
| **LSA / LSDB** | Annonces d'état de lien (Link State Advertisement) et base de données OSPF construite à partir d'elles. |
| **NAPT** | NAT avec translation de ports : plusieurs IP privées partagent une seule IP publique. |
| **NAT server** | Port mapping statique : expose un service interne vers l'extérieur (IP publique:port → IP privée:port). |
| **PVID** | VLAN natif d'un port : VLAN attribué aux trames non taggées entrantes. |
| **Router-ID** | Identifiant 32 bits d'un routeur OSPF (format d'adresse IP) ; doit être unique. |
| **RSTP** | Rapid STP (802.1w) : convergence rapide (~quelques secondes) via proposal/agreement. |
| **Silent-interface** | Interface OSPF qui annonce son réseau sans y envoyer de Hello (vers les LAN). |
| **SSID** | Nom du réseau Wi-Fi diffusé par les AP. |
| **STP** | Spanning Tree Protocol (802.1D) : élimine les boucles de niveau 2 en bloquant des ports redondants. |
| **Trunk 802.1Q** | Lien transportant plusieurs VLAN via taggage des trames. |
| **VAP** | Profil virtuel liant SSID + sécurité + VLAN de service, diffusé par les radios des AP. |
| **VRP** | Versatile Routing Platform : le système d'exploitation des équipements Huawei (équivalent d'IOS chez Cisco). |
| **Wildcard (masque inversé)** | Masque OSPF/ACL inversé (ex. /24 → 0.0.0.255). |
| **WPA2-PSK** | Authentification Wi-Fi par clé pré-partagée (Pre-Shared Key). |
| **Zone (USG)** | Domaine de sécurité du firewall (trust, untrust, dmz, local) ; les politiques s'expriment entre zones. |

---

## 17. Pour aller plus loin

### 17.1. Prolongements directs des TP (niveau 2 du parcours)

1. **OSPF multi-zones** : ajouter une area 1 derrière R2, observer les LSA inter-zones et le rôle de l'ABR.
2. **VRRP** : deux routeurs en passerelle redondante pour un même LAN (IP virtuelle, master/backup, preempt).
3. **Eth-Trunk (agrégation de liens)** : agréger 2 liens entre SW1 et SW2 (`interface Eth-Trunk`), mesurer le débit et tester la résilience.
4. **MSTP** : deux instances (VLAN 10 → instance 1, VLAN 20 → instance 2) avec des roots différents → équilibrage de charge des VLAN sur les liens.
5. **DHCP snooping + DAI** : sécuriser le VLAN 10 contre les serveurs DHCP pirates.
6. **IPSec avec NAT-T** : placer un NAT entre R1 et R2 (topologie du TP7+TP9 combinés) et activer la traversée de NAT.
7. **GRE over IPSec** : encapsuler du multicast (ou de l'OSPF !) dans du GRE lui-même protégé par IPSec — le classique des interconnexions dynamiques.
8. **PPPoE + IPv6** : attribuer un préfixe IPv6 via IPCP6/DHCPv6-PD sur la session du TP8.
9. **WLAN avancé** : deuxième SSID invité (portail captif simulé), limitation de débit par SSID, band steering.
10. **USG avancé** : politiques avec plages horaires, NAT bidirectionnel, VPN IPSec terminé sur l'USG lui-même (au lieu du routeur).

### 17.2. Du lab au terrain : transposer sur vos équipements

| Concept du lab (eNSP) | Équivalent terrain | Différences à connaître |
|---|---|---|
| AR2220 (lab) | AR720 | Interfaces, débits, modules 4G/USB ; CLI VRP quasi identique |
| S3700 (L2) / S5700 (L3) | S310 | Le S310 fait du L3 léger ; vérifier la fiche exacte (routage inter-VLAN, DHCP) |
| USG5500 (`policy interzone`) | USG6000 (`security-policy`) | Même logique de zones ; syntaxe des règles nommées sur USG6000 |
| AC6605 + AP6010DN | AC + AP361/AP761 | Wi-Fi 6 sur le terrain (OFDMA, MU-MIMO) ; logique Fit AP et profils identiques |
| STA eNSP | Smartphones/PC portables | Comportements roaming réels, à tester sur site |

### 17.3. Ressources

- **Vos guides d'équipe** (déjà en votre possession) : `huawei_ap361_guide.md`, `huawei_ap761_guide.md`, `huawei_s310_guide.md`, `huawei_ar720_guide.md`, `huawei_usg6000_guide.md`, `huawei_ekit_guide.md`, `huawei_esight_guide.md` — chacun détaille le matériel réel correspondant aux concepts des TP.
- **Documentation officielle Huawei** (support.huawei.com) : *Configuration Guides* VRP des AR/S/USG pour la syntaxe exacte par version.
- **Fichiers du parcours** : conserver un dossier par TP avec `sujet.md`, `topologie.topo` (vierge) et `correction.md` — ce guide en est la matière première.

### 17.4. Organisation conseillée de la formation d'équipe

| Semaine | Séance | Durée |
|---|---|---|
| S1 | Installation + prise en main + TP1 | 2 h 30 |
| S2 | TP2 + TP3 | 2 h 15 |
| S3 | TP4 | 1 h 30 |
| S4 | TP5 + TP6 | 2 h 30 |
| S5 | TP7 + TP8 | 2 h 45 |
| S6 | TP9 + TP10 | 3 h |
| S7 | TP11 + TP12 (évaluation) | 3 h |
| S8 | Quiz final + débrief + plan de progression individuel | 1 h |

> **Total : ~18 h** de formation, découpables en sessions d'1 h 30 à 3 h. Le TP12 noté + le quiz final donnent une évaluation à 50 points (30 + 20) convertible en appréciation par équipier.

---

## Annexe A — Aide-mémoire des commandes VRP

### Vues et navigation

```
<Huawei>system-view              # passer en vue système
[Huawei]sysname R1               # nommer l'équipement
Ctrl+Z                           # retour en vue utilisateur depuis n'importe où
quit                             # remonter d'un niveau
?                                # aide contextuelle
display this                     # config de la vue courante
```

### Informations et vérifications

