---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-18
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [2549, 2698]
sha256: 2deb1cf2324931a9fccf1954f51787453df35ada90a3cc96e9b2d689e7b21899
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

```huawei
system-view
# Deux interfaces WAN dans untrust :
[USG] interface GigabitEthernet 1/0/1
[USG-GigabitEthernet1/0/1] description WAN-FIBRE-PRINCIPAL
[USG-GigabitEthernet1/0/1] ip address 203.0.113.1 30
[USG-GigabitEthernet1/0/1] quit
[USG] interface GigabitEthernet 1/0/4
[USG-GigabitEthernet1/0/4] description WAN-4G-SECOURS
[USG-GigabitEthernet1/0/4] dhcp select interface      # IP via DHCP du routeur 4G
[USG-GigabitEthernet1/0/4] quit
[USG] firewall zone untrust
[USG-zone-untrust] add interface GigabitEthernet 1/0/1
[USG-zone-untrust] add interface GigabitEthernet 1/0/4
[USG-zone-untrust] quit
# Routes avec préférences : principal = 60 (défaut), secours = 100 :
[USG] ip route-static 0.0.0.0 0.0.0.0 203.0.113.2
[USG] ip route-static 0.0.0.0 0.0.0.0 GigabitEthernet 1/0/4 preference 100
# NAT Easy IP sur les deux sorties (egress-interface) :
[USG] nat-policy
[USG-policy-nat] rule name nat-wan1
[USG-policy-nat-rule-nat-wan1] source-zone trust
[USG-policy-nat-rule-nat-wan1] destination-zone untrust
[USG-policy-nat-rule-nat-wan1] egress-interface GigabitEthernet 1/0/1
[USG-policy-nat-rule-nat-wan1] source-address 192.168.10.0 24
[USG-policy-nat-rule-nat-wan1] action nat easy-ip
[USG-policy-nat-rule-nat-wan1] quit
[USG-policy-nat] rule name nat-wan2
[USG-policy-nat-rule-nat-wan2] source-zone trust
[USG-policy-nat-rule-nat-wan2] destination-zone untrust
[USG-policy-nat-rule-nat-wan2] egress-interface GigabitEthernet 1/0/4
[USG-policy-nat-rule-nat-wan2] source-address 192.168.10.0 24
[USG-policy-nat-rule-nat-wan2] action nat easy-ip
[USG-policy-nat-rule-nat-wan2] quit
[USG-policy-nat] quit
save
```

💡 Pour une **bascule intelligente** (pas seulement « si l'interface tombe »), couple avec **NQA** (section 114) : si le ping via la fibre échoue, la route est retirée → bascule sur la 4G même si le lien physique reste UP (cas du FAI « UP mais noir »).

## 165. Cas pratique 3 : interconnexion de 3 sites en IPSec (hub-and-spoke)

**Contexte** : siège (hub) + 2 agences (spokes). Chaque agence monte un tunnel vers le siège ; le trafic inter-agences passe par le hub.

- Sur chaque spoke : config IPSec standard vers le hub (section 59).
- Sur le hub : **2 jeux** de config (peer/proposal/policy/ACL) — un par spoke.
- ⚠️ **Pas de NAT** entre les réseaux des tunnels (exemptions `no-nat` sur chaque nœud).
- Routage : routes statiques vers les LAN distants via les interfaces Tunnel (ou OSPF, section 71).
- 📋 **Documenter** : tableau des PSK (dans le coffre !), IP publiques, LAN, n° de politique — un hub-and-spoke mal documenté devient ingérable à 5 sites.

## 166. Cas pratique 4 : Wi-Fi invités isolé avec portail

**Contexte** : invités + prestataires, accès Internet uniquement, **aucun accès** au LAN/DMZ.

```huawei
system-view
[USG] interface GigabitEthernet 1/0/5.100
[USG-GigabitEthernet1/0/5.100] vlan-type dot1q 100
[USG-GigabitEthernet1/0/5.100] ip address 192.168.100.1 24
[USG-GigabitEthernet1/0/5.100] quit
[USG] firewall zone name GUEST
[USG-zone-guest] set priority 20
[USG-zone-guest] add interface GigabitEthernet 1/0/5.100
[USG-zone-guest] quit
# DHCP pour les invités :
[USG] dhcp enable
[USG] interface GigabitEthernet 1/0/5.100
[USG-GigabitEthernet1/0/5.100] dhcp select interface
[USG-GigabitEthernet1/0/5.100] dhcp server ip-range 192.168.100.10 192.168.100.200
[USG-GigabitEthernet1/0/5.100] dhcp server gateway-list 192.168.100.1
[USG-GigabitEthernet1/0/5.100] dhcp server dns-list 8.8.8.8
[USG-GigabitEthernet1/0/5.100] quit
[USG] security-policy
# GUEST→Internet : web + DNS + UTM stricte :
[USG-policy-security] rule name guest-internet
[USG-policy-security-rule-guest-internet] source-zone GUEST
[USG-policy-security-rule-guest-internet] destination-zone untrust
[USG-policy-security-rule-guest-internet] service http https dns
[USG-policy-security-rule-guest-internet] profile url-profile-guest
[USG-policy-security-rule-guest-internet] profile av-profile-lan
[USG-policy-security-rule-guest-internet] action permit
[USG-policy-security-rule-guest-internet] quit
# GUEST→trust/dmz : DENY explicite (avec default deny c'est redondant, mais explicite = lisible) :
[USG-policy-security] rule name guest-deny-interne
[USG-policy-security-rule-guest-deny-interne] source-zone GUEST
[USG-policy-security-rule-guest-deny-interne] destination-zone trust
[USG-policy-security-rule-guest-deny-interne] action deny
[USG-policy-security-rule-guest-deny-interne] quit
[USG-policy-security] quit
# NAT pour les invités :
[USG] nat-policy
[USG-policy-nat] rule name nat-guest
[USG-policy-nat-rule-nat-guest] source-zone GUEST
[USG-policy-nat-rule-nat-guest] destination-zone untrust
[USG-policy-nat-rule-nat-guest] source-address 192.168.100.0 24
[USG-policy-nat-rule-nat-guest] action nat easy-ip
[USG-policy-nat-rule-nat-guest] quit
[USG-policy-nat] quit
save
```

## 167. Cas pratique 5 : paire HA pour un site critique (synthèse)

Reprendre : sections 97–106 pour la config, 101 pour le test, 104 pour l'upgrade.
📋 **Spécificités à ne pas oublier** : licences **identiques** des deux côtés, certificats SSL VPN **installés des deux côtés**, supervision des **deux** nœuds (pas seulement l'actif), **test de bascule semestriel** au planning.

## 168. Cas pratique 6 : migration d'un ancien firewall vers l'USG6000

1. **Inventaire** de l'existant : règles, NAT, VPN, routes — tout, même les règles « temporaires » de 2019.
2. **Traduction** : chaque règle → couple de zones + critères USG. C'est le moment de **nettoyer** (règles 0-hit = on ne migre pas).
3. **Maquette** : monter l'USG en labo, rejouer les flux critiques.
4. **Bascule** : en fenêtre de maintenance, avec **rollback** (rebrancher l'ancien = 5 minutes si le câblage est prêt).
5. **Recette** : checklist section 163 + tests utilisateurs pilotes.
6. 📋 Garder l'ancien firewall **1 mois** sous la main (éteint, config sauvegardée) avant recyclage.

---
---

# BLOC P — ERREURS CLASSIQUES À NE PAS COMMETTRE

## 169. Erreur 1 : `default action permit` oublié en prod

Mettre `permit` « pour les tests » et oublier de repasser en `deny`. **Conséquence** : firewall passoire pendant des mois. **Prévention** : ne JAMAIS mettre permit en prod, même temporairement — à la place, ajouter une règle `permit` large **avec une date de fin** (time-range) et une alerte.

## 170. Erreur 2 : l'interface sans zone

Interface UP, IP configurée, mais **jamais ajoutée** à une zone → aucun trafic. **Réflexe** : `display zone` dès que « l'interface est UP mais rien ne passe » (section 24).

## 171. Erreur 3 : le NAT sans la politique (ou l'inverse)

DNAT configuré mais pas de règle security-policy → le paquet traduit est **droppé**. **Réflexe** : tout DNAT s'accompagne **toujours** d'une règle de politique (section 48).

## 172. Erreur 4 : l'exemption NAT oubliée pour le VPN

Le NAT easy-ip attrape aussi le trafic du tunnel → VPN « monté mais muet ». **Réflexe** : `no-nat` **avant** la règle NAT, pour chaque réseau distant (section 52).

## 173. Erreur 5 : la PSK avec une faute de frappe

Deux côtés « identiques » sauf un caractère → phase 1 impossible, logs cryptiques. **Prévention** : copier-coller depuis le coffre, **jamais** de saisie manuelle, re-saisir des deux côtés au moindre doute (section 147).

## 174. Erreur 6 : UTM activée partout, modèle sous-dimensionné

Activer AV+IPS+URL+SSL-inspection sur toutes les règles d'un USG6330 avec 1 Gbit/s de trafic → **effondrement**. **Prévention** : dimensionner sur le **débit UTM** (section 73), appliquer les profils **de façon ciblée** (matrice section 84).

## 175. Erreur 7 : upgrade sans sauvegarde ni rollback

« Ça va bien se passer » → ça ne se passe pas bien → pas de backup → nuit blanche. **Prévention** : checklist section 123, **à chaque fois**, sans exception.

## 176. Erreur 8 : HA jamais testée

