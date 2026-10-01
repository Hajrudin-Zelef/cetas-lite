---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-8
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [1108, 1296]
sha256: d730bceb7dbae8095522bef8e80aa98f5b7968822da1b02313ea008f1e8be35b
---

# Huawei eKit AP361 — Guide ultra-complet

```
<AP361> system-view
[AP361] wlan
[AP361-wlan-view] ssid-profile name PROF-BUREAU
[AP361-wlan-ssid-prof-PROF-BUREAU] ssid STE-Bureau
[AP361-wlan-ssid-prof-PROF-BUREAU] quit
[AP361-wlan-view] security-profile name SEC-BUREAU
[AP361-wlan-ssid-prof-SEC-BUREAU] security wpa3-sae
[AP361-wlan-ssid-prof-SEC-BUREAU] wpa3 sae password cipher Exemple-Fictif-2026!Changez-Moi
[AP361-wlan-ssid-prof-SEC-BUREAU] quit
[AP361-wlan-view] vap-profile name VAP-BUREAU
[AP361-wlan-vap-prof-VAP-BUREAU] ssid-profile PROF-BUREAU
[AP361-wlan-vap-prof-VAP-BUREAU] security-profile SEC-BUREAU
[AP361-wlan-vap-prof-VAP-BUREAU] service-vlan vlan-id 20
[AP361-wlan-vap-prof-VAP-BUREAU] quit
```

> ⚠️ Noms de commandes **indicatifs** : `ssid-profile`, `security-profile`,
> `vap-profile`, `service-vlan` existent dans l'univers WLAN Huawei (AC/AP),
> mais les mots-clés exacts (`wpa3-sae`, etc.) varient selon la version.
> En mode cloud : faites-le dans l'app.

## 68. Rotation des clés PSK : procédure

1. Générez la nouvelle clé (20+ caractères aléatoires, stockée au coffre).
2. Appliquez-la **hors heures ouvrées** (chaque client devra se reconnecter).
3. Communiquez : affichez la nouvelle clé à l'accueil (invités), envoyez-la aux
   utilisateurs (bureau) via le canal interne habituel.
4. Vérifiez le lendemain : clients encore en échec = clé mal recopiée
   (le classique).
5. Fréquence : **trimestrielle** pour le SSID bureau, **mensuelle** pour invités.

## 69. Filtrage MAC : utile ou pas ?

- **Pas une sécurité** : une adresse MAC se spoofe en 2 minutes.
- **Utile en complément** : pour du matériel fixe connu (imprimantes, caméras)
  sur un SSID IoT, ça ajoute une barrière anti-erreur (pas anti-attaquant).
- **Coûteux à maintenir** : chaque changement de téléphone = une entrée à
  ajouter. En pratique : à réserver aux parcs stables.

## 70. PMF (Protected Management Frames / 802.11w)

- Protège les trames de management (désauthentification, désassociation) contre
  la falsification → anti « déauth attack ».
- **Requis** par WPA3, **optionnel** en WPA2.
- En WPA2 : passez en « capable/optional » si vos clients le supportent
  (la plupart des clients récents oui). Si un vieux client coince : repassez en
  « disabled » **pour ce SSID** et notez-le.

## 71. Checklist SSID/sécurité

- [ ] ≤ 4 SSID par AP
- [ ] SSID diffusés (pas de SSID caché sauf justification)
- [ ] Bureau : WPA3-SAE ou transition, clé 20+ caractères au coffre
- [ ] Invité : SSID + VLAN dédiés, portail, isolation, débit limité
- [ ] IoT : SSID 2,4 GHz only, WPA2-PSK, VLAN isolé
- [ ] PMF : required (WPA3) / optional (WPA2)
- [ ] Rotation des clés planifiée (trimestrielle/mensuelle)

## 72. Erreurs classiques SSID

| Erreur | Symptôme | Correction |
|---|---|---|
| WPA3 pur + vieux clients | Certains PC/imprimantes ne s'associent plus | Mode transition |
| Clé avec caractères exotiques | Saisie impossible sur IoT (pas de clavier) | Clé alphanumérique simple pour IoT |
| Trop de SSID | Débit global en berne | Fusionner, max 4 |
| Même clé partout depuis 3 ans | — (risque silencieux) | Rotation planifiée |
| SSID bureau sans VLAN | Invités sur le LAN si clé fuitée | VLAN dédiés (§73) |

---
---

# H. VLAN : SSID → VLAN, TRUNK VERS LE SWITCH

## 73. Le principe : un SSID = un VLAN (en général)

Le VLAN sépare les populations **au niveau filaire**, dès la sortie de l'AP :

```
Client Wi-Fi --(SSID STE-Bureau)--> AP361 --(tag VLAN 20)--> Switch --(routage/firewall)--> LAN/Internet
Client Wi-Fi --(SSID STE-Invite)--> AP361 --(tag VLAN 30)--> Switch --(routage/firewall)--> Internet seul
```

**Plan de VLAN type PME** :

| VLAN | Nom | Usage | Accès |
|---|---|---|---|
| 1 (défaut) | — | Éviter pour la prod | — |
| 10 | MGMT | Management (AP, switchs) | Admins uniquement |
| 20 | BUREAU | Collaborateurs Wi-Fi (+ filaire) | LAN + Internet |
| 30 | INVITE | Visiteurs Wi-Fi | Internet uniquement |
| 40 | IOT | Objets connectés | Internet restreint, pas de LAN |
| 50 | VOIX | Téléphonie (optionnel) | Serveur téléphonique |

> ⚠️ **Ne laissez pas les AP et les utilisateurs sur le même VLAN** : un invité
> qui devine l'IP de l'AP ne doit pas pouvoir ouvrir sa page d'admin.

## 74. Côté AP : associer SSID → VLAN

Dans l'app eKit : `SSID > Paramètres avancés > VLAN` → saisir l'ID (ex. 20).
En CLI (indicatif) : `service-vlan vlan-id 20` dans le vap-profile (voir §67).

**VLAN natif / non tagué** : le trafic de management de l'AP lui-même transite
(en général) **non tagué** sur le port du switch → le VLAN natif du port doit
correspondre au VLAN de management (voir §75).

## 75. Côté switch : le trunk vers l'AP

Le port du switch qui alimente l'AP doit être en **trunk** : VLAN management en
natif (non tagué) + VLANs utilisateurs tagués.

**Exemple côté switch Huawei (S5735, VRP — syntaxe réelle courante)** :

```
<SW> system-view
[SW] vlan batch 10 20 30 40
[SW] interface GigabitEthernet 0/0/5        # port vers AP-01
[SW-GigabitEthernet0/0/5] port link-type trunk
[SW-GigabotEthernet0/0/5] port trunk pvid vlan 10      # VLAN natif = management
[SW-GigabitEthernet0/0/5] port trunk allow-pass vlan 10 20 30 40
[SW-GigabitEthernet0/0/5] port-isolate enable          # optionnel : isole les ports AP entre eux
[SW-GigabitEthernet0/0/5] quit
[SW] save
```

Vérification :

```
<SW> display vlan
<SW> display port trunk
<SW> display interface GigabitEthernet 0/0/5
```

> 💡 **Test décisif** : client connecté au SSID Bureau → `ipconfig` doit montrer
> une IP du sous-réseau du VLAN 20, pas du VLAN 10. Sinon le tagging est faux
> (voir §126).

## 76. Isoler le VLAN invité : ACL / firewall

Un VLAN séparé ne suffit pas : il faut **bloquer le routage** vers le LAN.
Selon votre routeur/firewall :

**Exemple d'ACL sur switch L3 Huawei (logique indicative)** :

```
[SW] acl number 3001
[SW-acl-adv-3001] rule deny ip source 192.168.30.0 0.0.0.255 destination 192.168.0.0 0.0.255.255
[SW-acl-adv-3001] rule permit ip source 192.168.30.0 0.0.0.255 destination any
```

Traduction : le VLAN 30 (invités, 192.168.30.0/24) peut aller vers Internet
mais **pas** vers les réseaux internes (192.168.0.0/16).

**Mieux** : faites-le sur le **firewall** (USG/OPNsense/pfSense) avec une vraie
règle inter-VLAN + log. L'ACL switch dépanne, le firewall administre.

## 77. DHCP par VLAN

Chaque VLAN utilisateur a besoin de son **serveur DHCP** (sauf IP statiques, rare
en Wi-Fi) :

| VLAN | Sous-réseau exemple | DHCP | Passerelle |
|---|---|---|---|
| 10 MGMT | 192.168.10.0/24 | Statique ou réservations | 192.168.10.1 |
| 20 BUREAU | 192.168.20.0/24 | Serveur Windows / routeur | 192.168.20.1 |
| 30 INVITE | 192.168.30.0/24 | Routeur/firewall (bail court : 2 h) | 192.168.30.1 |
| 40 IOT | 192.168.40.0/24 | Routeur (baux longs, réservations) | 192.168.40.1 |

**Bail court sur INVITE** (1-2 h) : les adresses se recyclent vite avec le
turnover des visiteurs. **Réservations DHCP** sur MGMT : chaque AP a toujours la
même IP (indispensable pour la supervision).

## 78. Exemple complet : 3 SSID / 3 VLAN sur un site

```
Site : SIEGE-ETAGE2 — 6 x AP361

VLAN 10 MGMT   : 192.168.10.0/24  (AP en .11 à .16, switch .2, passerelle .1)
VLAN 20 BUREAU : 192.168.20.0/24  (SSID STE-Bureau, WPA3-SAE transition)
VLAN 30 INVITE : 192.168.30.0/24  (SSID STE-Invite, portail, isolé)
VLAN 40 IOT    : 192.168.40.0/24  (SSID STE-IoT, 2,4 GHz only)

Ports switch 1-6 : trunk, PVID 10, allow 10/20/30/40
DHCP : Windows Server pour VLAN 20, firewall pour 30/40
DNS : interne pour 20, public (ou filtré) pour 30
```

## 79. Erreurs VLAN classiques

