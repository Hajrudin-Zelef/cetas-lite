---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-11
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [1577, 1737]
sha256: fcda03729d2037e72edb02e0f9a4240fb810829eba91d03d0d8081a896d0a61d
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

- **AC1 (AC6605)** : contrôleur WLAN, GE0/0/1 = 192.168.100.1/24 (VLAN 100).
- **SW1 (S3700)** : VLAN 100 (ports GE0/0/1-3 en access VLAN 100) + VLAN 10 (service, 192.168.10.0/24).
- **AP1, AP2 (AP6010DN-AGN)** : branchés sur SW1, obtiennent leur IP en DHCP (ou statique) dans le VLAN 100, découvrent l'AC.
- **STA1, STA2** : stations Wi-Fi simulées (objets STA d'eNSP), à associer au SSID.
- SSID : `LABO-ENTREPRISE`, WPA2-PSK, clé `WifiLabo2026!`, VLAN de service 10.

> **Note de cadrage** : eNSP émule des AP Wi-Fi 4/5 (AP6010DN), pas vos AP361/AP761 (Wi-Fi 6). La **logique Fit AP** (profils SSID/sécurité/VAP, découverte CAPWAP, radios) est identique : ce TP entraîne aux concepts que vous retrouverez sur le vrai matériel (voir vos guides AP361/AP761 pour les spécificités Wi-Fi 6).

### Énoncé

1. Adresser le VLAN 100 (gestion) : AC1 en 192.168.100.1/24, AP en DHCP ou statique dans 192.168.100.0/24.
2. Sur AC1 : activer le WLAN (`wlan`), déclarer les AP (par MAC), les nommer AP1/AP2.
3. Créer les profils : `security-profile` (WPA2-PSK), `ssid-profile` (nom du SSID), `vap-profile` (lie les deux + VLAN de service 10 + mode de transfert).
4. Créer le groupe AP / domaine de régulation, associer le VAP aux radios des AP.
5. Vérifier que les AP passent à l'état **normal** (tunnel CAPWAP établi).
6. Associer STA1/STA2 au SSID avec la clé, vérifier l'obtention d'IP (DHCP du VLAN 10 via R1 ou via AC).
7. Tester la connectivité Wi-Fi ↔ filaire.

### Correction pas à pas

**Réseau de gestion (VLAN 100) :**

```
# SW1
[SW1]vlan batch 10 100
[SW1]interface GigabitEthernet 0/0/1
[SW1-GigabitEthernet0/0/1]port link-type access
[SW1-GigabitEthernet0/0/1]port default vlan 100
# Idem GE0/0/2, GE0/0/3 (vers AC et AP)
```

**AC1 — base :**

```
[AC1]interface GigabitEthernet 0/0/1
[AC1-GigabitEthernet0/0/1]undo portswitch
[AC1-GigabitEthernet0/0/1]ip address 192.168.100.1 24
[AC1]wlan
[AC1-wlan-view]quit
```

**Déclaration des AP (récupérer leurs MAC via `display ap all` après qu'ils ont démarré, ou les lire sur l'étiquette eNSP) :**

```
[AC1]wlan
[AC1-wlan-view]ap-id 1 ap-mac 00e0-fc12-3456
[AC1-wlan-ap-1]ap-name AP1
[AC1-wlan-ap-1]quit
[AC1-wlan-view]ap-id 2 ap-mac 00e0-fc65-4321
[AC1-wlan-ap-2]ap-name AP2
[AC1-wlan-ap-2]quit
```

> En lab, les MAC des AP eNSP sont visibles dans les propriétés de l'équipement ou via `display ap all` (état idle avant configuration). Sur le vrai matériel, on utilise souvent l'authentification par **numéro de série (SN)**.

**Profils WLAN :**

```
# 1. Profil de sécurité : WPA2-PSK
[AC1-wlan-view]security-profile name SEC-LABO
[AC1-wlan-sec-prof-SEC-LABO]security wpa2 psk pass-phrase WifiLabo2026! aes
[AC1-wlan-sec-prof-SEC-LABO]quit

# 2. Profil SSID
[AC1-wlan-view]ssid-profile name SSID-LABO
[AC1-wlan-ssid-prof-SSID-LABO]ssid LABO-ENTREPRISE
[AC1-wlan-ssid-prof-SSID-LABO]quit

# 3. Profil VAP : lie SSID + sécurité + VLAN de service + transfert direct
[AC1-wlan-view]vap-profile name VAP-LABO
[AC1-wlan-vap-prof-VAP-LABO]service-vlan vlan-id 10
[AC1-wlan-vap-prof-VAP-LABO]ssid-profile SSID-LABO
[AC1-wlan-vap-prof-VAP-LABO]security-profile SEC-LABO
[AC1-wlan-vap-prof-VAP-LABO]forward-mode direct-forward
[AC1-wlan-vap-prof-VAP-LABO]quit
```

**Domaine de régulation + groupe AP :**

```
[AC1-wlan-view]regulatory-domain-profile name REG-FR
[AC1-wlan-regulatory-domain-prof-REG-FR]country-code FR
[AC1-wlan-regulatory-domain-prof-REG-FR]quit

[AC1-wlan-view]ap-group name GROUPE-LABO
[AC1-wlan-ap-group-GROUPE-LABO]regulatory-domain-profile REG-FR
[AC1-wlan-ap-group-GROUPE-LABO]vap-profile VAP-LABO wlan 1 radio all
[AC1-wlan-ap-group-GROUPE-LABO]quit

# Rattacher les AP au groupe :
[AC1-wlan-view]ap-id 1
[AC1-wlan-ap-1]ap-group GROUPE-LABO
[AC1-wlan-ap-1]quit
[AC1-wlan-view]ap-id 2
[AC1-wlan-ap-2]ap-group GROUPE-LABO
[AC1-wlan-ap-2]quit
[AC1]save
```

**Vérifier l'état des AP :**

```
[AC1]display ap all
# État "nor" (normal) = tunnel CAPWAP établi, profils poussés, radios actives.
# Les états intermédiaires (config, configFailed...) indiquent où ça bloque.
```

**Associer les STA :** double-clic sur STA1 > onglet WLAN > SSID `LABO-ENTREPRISE`, sécurité WPA2-PSK, clé `WifiLabo2026!`. La STA doit s'associer (état connecté visible dans eNSP) puis obtenir une IP du VLAN 10 (prévoir un serveur DHCP sur ce VLAN, ex. R1 du TP4 ou un pool sur l'AC).

**Vérifications côté AC :**

```
[AC1]display ap all                    # AP en "nor"
[AC1]display station all               # clients associés : MAC, SSID, AP, radio, VLAN
[AC1]display vap all                   # VAP diffusés par AP/radio
[AC1]display ssid-profile name SSID-LABO
```

### Pièges classiques

1. **AP qui reste en `idle`** : pas de connectivité IP avec l'AC (VLAN 100 mal configuré, trunk oublié) → le tunnel CAPWAP ne monte pas. Vérifier le ping AP→AC.
2. **AP en `configFailed`** : souvent un profil référencé qui n'existe pas (faute de frappe dans le nom du ssid-profile) → `display ap config-info ap-id 1` pour le détail.
3. **`vap-profile ... wlan 1 radio all` oublié** : les profils existent mais ne sont pas diffusés → le SSID n'apparaît pas. C'est l'oubli n°1.
4. **Mauvais `service-vlan`** : clients associés mais sans IP / sans connectivité → le VLAN de service ne correspond pas au VLAN du DHCP/routeur.
5. **`forward-mode`** : en `tunnel-forward` (centralisé, défaut sur certains modèles), tout le trafic client remonte à l'AC ; en `direct-forward`, l'AP bridge localement vers le switch. En lab les deux marchent, mais le mode doit être cohérent avec le câblage (en tunnel-forward, pas besoin du VLAN 10 sur le lien AP→switch).
6. **Clé PSK avec caractères spéciaux** : en lab, rester sur de l'alphanumérique pour éviter les problèmes de saisie côté STA.
7. **Régulation** : sans `country-code`, les radios peuvent rester éteintes (puissance/canaux non autorisés). Toujours configurer le domaine de régulation.

### Barème indicatif (20 points)

| Critère | Points |
|---|---|
| VLAN 100 + adressage AC/AP | 3 |
| AP déclarés (MAC) et nommés | 3 |
| Les 3 profils (sécurité, SSID, VAP) corrects et liés | 5 |
| Groupe AP + domaine de régulation + VAP sur les radios | 3 |
| AP en état normal, STA associées avec IP | 4 |
| Vérifications (`display ap all`, `display station all`) | 2 |

### Durée estimée
**1 h 30**.

### Fiche animateur — points à insister
- La logique **Fit AP** : l'AP est "bête", l'AC est "intelligent" (configuration centralisée, tunnel CAPWAP). C'est exactement l'architecture de vos AP361/AP761 gérés en central.
- Les **profils** sont des briques réutilisables : un même `security-profile` peut servir à plusieurs SSID, un même `vap-profile` à plusieurs groupes d'AP. C'est la clé d'un déploiement propre à l'échelle d'une entreprise.
- `display ap all` est LE réflexe : l'état de l'AP dit tout (idle = réseau, config = en cours, nor = OK, configFailed = profil).
- Lien avec vos guides AP361/AP761 : la syntaxe des profils est la même famille ; les différences portent sur le Wi-Fi 6 (OFDMA, etc.), pas sur la logique Fit.
- Question piège : "Un client voit le SSID mais n'obtient pas d'IP : où chercher ?" → service-vlan, DHCP du VLAN, forward-mode. "Un client ne voit pas le SSID ?" → VAP non lié aux radios, AP pas en normal.

---

## 13. TP11 — USG : zones, politiques, NAT sortant

### Objectif
Prendre en main un pare-feu Huawei USG : zones de sécurité, politiques interzones, NAT sortant — la base de l'administration de vos USG6000.

### Prérequis
TP1 (VLAN), TP5 (routage), TP7 (NAT).

### Topologie

