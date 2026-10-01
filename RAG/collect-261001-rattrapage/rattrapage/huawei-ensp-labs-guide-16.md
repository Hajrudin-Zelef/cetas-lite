---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-16
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "arr", "cost", "ethernet", "intel"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [2288, 2512]
sha256: 17f8db9cf0e516d830e5aedc597a91f2d69a78c95f2e65332adf39c7966bb1af
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

```
display version                  # version VRP, modèle, uptime
display current-configuration    # configuration complète
display saved-configuration      # configuration sauvegardée
display interface brief          # état de toutes les interfaces (up/down)
display ip interface brief       # adresses IP des interfaces
display ip routing-table         # table de routage
display mac-address              # table MAC (par VLAN sur switch)
display vlan [vlan-id]           # VLAN et ports membres
display port vlan                # synthèse ports : type, PVID, VLAN autorisés
display arp                      # table ARP
save                             # sauvegarder la configuration
reset saved-configuration        # effacer la config sauvegardée
reboot                           # redémarrer (après reset saved-configuration = reset usine)
```

### Switching

```
vlan batch 10 20 30              # créer plusieurs VLAN
port link-type access | trunk | hybrid
port default vlan 10             # VLAN d'un port access (PVID)
port trunk allow-pass vlan 10 20 # VLAN autorisés sur un trunk
port-group group-member Ethernet 0/0/1 to Ethernet 0/0/4  # config groupée
stp mode rstp | stp enable       # activer RSTP
stp root primary | secondary     # forcer l'élection
stp priority 4096                # priorité manuelle (multiple de 4096)
stp edged-port enable            # port vers terminal (jamais inter-switch)
display stp brief                # rôles et états des ports
```

### DHCP

```
dhcp enable                      # activer le service DHCP (global)
ip pool NOM                      # créer un pool
 network 192.168.10.0 mask 24   # réseau distribué
 gateway-list 192.168.10.254    # passerelle
 dns-list 8.8.8.8               # DNS
 excluded-ip-address 192.168.10.1 192.168.10.10  # exclusions
 lease day 1 hour 0              # durée du bail
dhcp select global               # sur l'interface : utiliser les pools globaux
dhcp relay server-ip 192.168.100.1  # relais vers un serveur distant
display ip pool [name NOM]       # pools et baux
```

### Routage

```
ip route-static 192.168.20.0 24 10.0.12.2 [preference 100]  # route statique
ip route-static 0.0.0.0 0.0.0.0 10.0.12.2                    # route par défaut
undo portswitch                  # passer un port GE en mode routé (L3)
ospf 1 router-id 1.1.1.1         # process OSPF + Router-ID
 area 0
  network 10.0.12.0 0.0.0.3      # déclaration avec wildcard
ospf dr-priority 100             # priorité DR (sur l'interface)
ospf cost 100                    # coût manuel (sur l'interface)
default-route-advertise always   # injecter la défaut dans OSPF
display ospf peer brief          # voisins OSPF
display ospf lsdb                # base des LSAs
```

### NAT

```
acl 2000
 rule permit source 192.168.10.0 0.0.0.255   # qui a droit au NAT
interface GigabitEthernet 0/0/2
 nat outbound 2000                            # Easy IP (NAPT) en sortie
 nat server protocol tcp global 198.51.100.1 8080 inside 192.168.10.100 80  # port mapping
display nat session all          # sessions actives
display nat server               # mappings statiques
```

### PPPoE

```
# Serveur :
interface Virtual-Template 1
 ppp authentication-mode chap
 ip address 198.51.100.2 30
 remote address pool POOL-PPPOE
interface GigabitEthernet 0/0/2
 pppoe-server bind virtual-template 1
# Client :
dialer-rule
 dialer-rule 1 ip permit
interface Dialer 1
 dialer user client-lab
 dialer-group 1
 dialer bundle 1
 ppp chap user client-lab
 ppp chap password cipher huawei123
 ip address ppp-negotiate
interface GigabitEthernet 0/0/2
 pppoe-client dial-bundle-number 1
display pppoe-client session summary
```

### IPSec

```
ike proposal 10
 encryption-algorithm aes-256
 authentication-algorithm sha2-256
 dh group14
ike peer R2
 pre-shared-key simple Cle-IPSec-Lab-2026
 ike-proposal 10
 remote-address 100.64.1.2
ipsec proposal PROP-A
 esp authentication-algorithm sha2-256
 esp encryption-algorithm aes-256
acl 3000
 rule permit ip source 192.168.10.0 0.0.0.255 destination 192.168.20.0 0.0.0.255
ipsec policy POLICY-A 10 isakmp
 security acl 3000
 ike-peer R2
 proposal PROP-A
interface GigabitEthernet 0/0/2
 ipsec policy POLICY-A
display ike sa
display ipsec sa
```

### WLAN (AC)

```
wlan                                     # entrer en vue WLAN
ap-id 1 ap-mac 00e0-fc12-3456            # déclarer un AP
 ap-name AP1
security-profile name SEC-LABO
 security wpa2 psk pass-phrase Cle123! aes
ssid-profile name SSID-LABO
 ssid LABO-ENTREPRISE
vap-profile name VAP-LABO
 service-vlan vlan-id 10
 ssid-profile SSID-LABO
 security-profile SEC-LABO
 forward-mode direct-forward
regulatory-domain-profile name REG-FR
 country-code FR
ap-group name GROUPE-LABO
 regulatory-domain-profile REG-FR
 vap-profile VAP-LABO wlan 1 radio all
ap-id 1
 ap-group GROUPE-LABO
display ap all
display station all
```

### USG (zones et politiques)

```
firewall zone trust
 add interface GigabitEthernet 0/0/1
policy interzone trust untrust outbound
 policy 1
  action permit
  policy source 192.168.10.0 mask 24
display zone
display policy interzone trust untrust outbound
display firewall session table
```

---

## Annexe B — Modèle de fiche de lab vierge

> À dupliquer pour chaque nouveau TP créé par l'équipe.

```markdown
# TPXX — Titre

## Objectif
(1-2 phrases : ce que le stagiaire saura faire à la fin)

## Prérequis
(TP précédents, notions)

## Topologie
(Schéma textuel : équipements, interconnexions, adressage)

## Énoncé
(Étapes numérotées, sans les commandes)

## Correction pas à pas
(Commandes VRP complètes par équipement)

## Vérifications
(display ... attendus)

## Pièges classiques
(3-6 pièges avec diagnostic)

## Barème indicatif (20 points)
(Tableau critère / points)

## Durée estimée
## Fiche animateur
(Points à insister, questions pièges, lien terrain)
```

---

*Fin du guide — Labs eNSP pour former son équipe. Bon courage à l'équipe, et que les pings soient avec vous.*

---

## Annexe C — Dépannage de l'installation eNSP (pense-bête formateur)

| Symptôme | Cause probable | Solution |
|---|---|---|
| Équipements affichés `???` après démarrage | Images VirtualBox non enregistrées / antivirus bloquant | Réinstaller eNSP **en administrateur**, désactiver l'antivirus pendant l'install |
| Message "VirtualBox version too low" | Version de VirtualBox incompatible | Installer **VirtualBox 5.2.x** (ex. 5.2.44), chemin sans accents |
| Les nœuds démarrent puis s'arrêtent aussitôt | Hyper-V / WSL2 actif | `bcdedit /set hypervisorlaunchtype off` + désactiver Hyper-V + reboot |
| Erreur VT-x / AMD-V | Virtualisation désactivée dans le BIOS | Activer Intel VT-x / AMD-V (SVM Mode) dans le BIOS/UEFI |
| La capture Wireshark ne démarre pas | WinPcap absent ou mal installé | Installer **WinPcap 4.1.3** (pas seulement Npcap) |
| eNSP ne détecte pas VirtualBox | Ordre d'installation incorrect | Toujours installer dans l'ordre : VirtualBox → WinPcap → eNSP → reboot |
| Console CLI noire / pas de prompt | Équipement pas encore booté | Attendre 1 à 2 min (premier boot long), vérifier l'icône verte |
| Lenteurs avec 6+ équipements | RAM insuffisante | 16 Go conseillés ; ne démarrer que les équipements du TP en cours |
| Le .topo s'ouvre mais les configs ont disparu | `save` jamais exécuté sur les équipements | Réflexe : `save` en fin de chaque TP avant d'enregistrer le .topo |
| Conflit avec Docker Desktop | Docker utilise WSL2/Hyper-V | Choisir : Docker **ou** eNSP sur le poste (ou 2 postes / dual-boot) |

> **Checklist de validation d'un poste stagiaire** (5 minutes) : eNSP lancé en admin → palette complète visible → lab "hello world" (2 routeurs) démarré → ping OK → capture Wireshark OK sur le lien → `hello_world.topo` enregistré et rouvert. Si ces 5 points passent, le poste est prêt pour tout le parcours.
