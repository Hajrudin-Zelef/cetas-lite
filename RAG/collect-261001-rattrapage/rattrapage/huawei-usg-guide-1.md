---
id: collect-261001-rattrapage/rattrapage/huawei-usg-guide-1
title: "Huawei USG — Guide CLI complet (USG6000 series, V500R005 / V600)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_usg_guide.md
source_anchor: ""
source_lines: [1, 246]
sha256: 0c20dcd14687ebc05cc63e3d933ceb402f9033bced2a08fe62d0764d7699b6f9
---

# Huawei USG — Guide CLI complet (USG6000 series, V500R005 / V600)

> Synthèse rédigée à partir de la documentation officielle Huawei (Command
> Reference & Configuration Guide USG6000). Couvre l'essentiel du CLI VRP
> appliqué aux pare-feu USG : installation, zones, politiques, NAT, VPN,
> UTM, HA, supervision et maintenance. Vérifiez les variantes selon votre
> version logicielle exacte (`display version`).

---

## 1. Gammes et concepts

- **USG6300 / 6500 / 6600 / 6700** : pare-feu « NGFW » (Next-Generation
  Firewall), OS = VRP (même base que routeurs/switches Huawei).
- Le concept central du USG est la **zone de sécurité** : tout le trafic
  est contrôlé *entre zones*, jamais directement entre interfaces.
- Zones prédéfinies : `local` (100), `trust` (85), `dmz` (50),
  `untrust` (5). Le chiffre = priorité (plus haut = plus de confiance).
- Par défaut, **tout trafic interzone est refusé** : il faut des règles
  explicites pour autoriser.
- Deux moteurs de politique coexistent :
  - `security-policy` (mode NGFW moderne, recommandé) ;
  - `policy interzone` (mode legacy).
- Les profils UTM (IPS, antivirus, filtrage URL…) s'appliquent **dans**
  les règles `security-policy` via `profile`.

---

## 2. Premier accès

- Console : 9600 bauds, 8N1. Identifiants par défaut :
  `admin` / `Admin@123` (changement obligatoire à la première connexion).
- Interface de management : `GigabitEthernet0/0/0` = **192.168.0.1/24**.
- Web UI : `https://192.168.0.1:8443` (activée par défaut).
- SSH activé après configuration (voir §13).

```shell
# Passer en mode configuration
system-view
[USG] sysname USG-SIEGE

# Voir la version et le matériel
display version
display device
display device fan
display device power
```

---

## 3. Bases du CLI VRP

```shell
system-view              # user view <USG> -> system view [USG]
quit                     # remonte d'un niveau
return                   # retour direct en user view (Ctrl+Z)
display current-configuration   # config active
display saved-configuration     # config sauvegardée
save                     # sauvegarde (toujours après modifs !)
compare configuration    # diff entre active et sauvegardée
display this             # config du niveau courant
?                        # aide contextuelle partout
display history-command   # historique
```

- `display` = afficher (jamais `show`), `undo` = annuler une commande.
- Les commandes s'abrègent : `dis cur`, `sys`, etc.
- Niveaux de commande : 0 (visite) à 3 (management) ; `super` pour monter.

---

## 4. Interfaces

```shell
system-view
[USG] interface GigabitEthernet 0/0/1
[USG-GigabitEthernet0/0/1] ip address 192.168.1.1 24
[USG-GigabitEthernet0/0/1] description LAN_SIEGE
[USG-GigabitEthernet0/0/1] quit

# Sous-interface VLAN (trunk 802.1Q)
[USG] interface GigabitEthernet 0/0/1.10
[USG-GigabitEthernet0/0/1.10] vlan-type dot1q 10
[USG-GigabitEthernet0/0/1.10] ip address 192.168.10.1 24
[USG-GigabitEthernet0/0/1.10] quit

# Agrégat de liens
[USG] interface Eth-Trunk 1
[USG-Eth-Trunk1] mode lacp-static
[USG-Eth-Trunk1] trunkport GigabitEthernet 0/0/3 to 0/0/4
[USG-Eth-Trunk1] ip address 10.0.0.1 30
[USG-Eth-Trunk1] quit

# Vérifications
display interface brief
display ip interface brief
display interface GigabitEthernet 0/0/1
```

### PPPoE (fibre/ADSL grand public)

```shell
[USG] interface Dialer 1
[USG-Dialer1] dialer user u1
[USG-Dialer1] dialer-group 1
[USG-Dialer1] dialer bundle 1
[USG-Dialer1] ppp chap user monlogin password cipher MonPass
[USG-Dialer1] ip address ppp-negotiate
[USG-Dialer1] quit
[USG] interface GigabitEthernet 0/0/2
[USG-GigabitEthernet0/0/2] pppoe-client dial-bundle-number 1
[USG-GigabitEthernet0/0/2] quit
[USG] dialer-rule
[USG-dialer-rule] dialer-rule 1 ip permit
```

---

## 5. Zones de sécurité

```shell
system-view
# Zones prédéfinies : ajouter les interfaces
[USG] firewall zone trust
[USG-zone-trust] set priority 85
[USG-zone-trust] add interface GigabitEthernet 0/0/1
[USG-zone-trust] quit

[USG] firewall zone untrust
[USG-zone-untrust] set priority 5
[USG-zone-untrust] add interface GigabitEthernet 0/0/2
[USG-zone-untrust] quit

[USG] firewall zone dmz
[USG-zone-dmz] set priority 50
[USG-zone-dmz] add interface GigabitEthernet 0/0/3
[USG-zone-dmz] quit

# Zone personnalisée (ex : invités)
[USG] firewall zone name GUEST
[USG-zone-GUEST] set priority 60
[USG-zone-GUEST] add interface GigabitEthernet 0/0/1.20
[USG-zone-GUEST] quit

display zone
display firewall zone trust
```

- La zone `local` = le pare-feu lui-même (ping, SSH, web vers le USG).
- Une interface ne peut appartenir qu'à **une seule** zone.

---

## 6. Security-policy (cœur du USG)

```shell
system-view
[USG] security-policy
[USG-policy-security] rule name LAN_TO_WAN
[USG-policy-security-rule-LAN_TO_WAN] source-zone trust
[USG-policy-security-rule-LAN_TO_WAN] destination-zone untrust
[USG-policy-security-rule-LAN_TO_WAN] source-address 192.168.1.0 24
[USG-policy-security-rule-LAN_TO_WAN] destination-address any
[USG-policy-security-rule-LAN_TO_WAN] action permit
[USG-policy-security-rule-LAN_TO_WAN] quit

# Autoriser le ping / l'admin vers le pare-feu lui-même
[USG-policy-security] rule name ADMIN_TO_FW
[USG-policy-security-rule-ADMIN_TO_FW] source-zone trust
[USG-policy-security-rule-ADMIN_TO_FW] destination-zone local
[USG-policy-security-rule-ADMIN_TO_FW] service ping ssh https
[USG-policy-security-rule-ADMIN_TO_FW] action permit
[USG-policy-security-rule-ADMIN_TO_FW] quit
[USG-policy-security] quit

# Vérifications
display security-policy rule all
display security-policy rule name LAN_TO_WAN
```

- Les règles se lisent **dans l'ordre** ; la première qui matche gagne.
- `action deny` = refuser (loggé si `logging` activé sur la règle).
- Sans règle correspondante : trafic **refusé** par défaut.
- Le trafic *retour* d'une session autorisée est automatiquement accepté
  (pare-feu stateful) : inutile de créer la règle inverse.

---

## 7. Objets : adresses, services, plannings

```shell
system-view
# Groupe d'adresses
[USG] ip address-set SERVEURS_DMZ
[USG-address-set-SERVEURS_DMZ] address 0 192.168.50.0 mask 24
[USG-address-set-SERVEURS_DMZ] quit

# Groupe de services
[USG] ip service-set WEB
[USG-service-set-WEB] service 0 protocol tcp destination-port 80
[USG-service-set-WEB] service 1 protocol tcp destination-port 443
[USG-service-set-WEB] quit

# Plage horaire (ex : heures ouvrées)
[USG] time-range BUREAU 08:00 to 18:00 daily

# Utilisation dans une règle
[USG] security-policy
[USG-policy-security] rule name WAN_TO_DMZ_WEB
[USG-policy-security-rule-WAN_TO_DMZ_WEB] source-zone untrust
[USG-policy-security-rule-WAN_TO_DMZ_WEB] destination-zone dmz
[USG-policy-security-rule-WAN_TO_DMZ_WEB] destination-address address-set SERVEURS_DMZ
[USG-policy-security-rule-WAN_TO_DMZ_WEB] service service-set WEB
[USG-policy-security-rule-WAN_TO_DMZ_WEB] schedule time-range BUREAU
[USG-policy-security-rule-WAN_TO_DMZ_WEB] action permit
```

---

## 8. NAT source (accès Internet)

```shell
system-view
# Pool d'adresses publiques (PAT)
[USG] nat address-group SNAT_POOL
[USG-address-group-SNAT_POOL] mode pat
[USG-address-group-SNAT_POOL] section 0 202.10.10.10 202.10.10.20
[USG-address-group-SNAT_POOL] quit

[USG] nat-policy
[USG-policy-nat] rule name SNAT_LAN
[USG-policy-nat-rule-SNAT_LAN] source-zone trust
[USG-policy-nat-rule-SNAT_LAN] destination-zone untrust
[USG-policy-nat-rule-SNAT_LAN] source-address 192.168.1.0 24
[USG-policy-nat-rule-SNAT_LAN] action source-nat address-group SNAT_POOL
[USG-policy-nat-rule-SNAT_LAN] quit
[USG-policy-nat] quit

# Variante Easy-IP : utilise l'IP de l'interface de sortie
[USG-policy-nat] rule name SNAT_EASY
[USG-policy-nat-rule-SNAT_EASY] source-zone trust
[USG-policy-nat-rule-SNAT_EASY] destination-zone untrust
[USG-policy-nat-rule-SNAT_EASY] source-address 192.168.1.0 24
[USG-policy-nat-rule-SNAT_EASY] action source-nat easy-ip
```

