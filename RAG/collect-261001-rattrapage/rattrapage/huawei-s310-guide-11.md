---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-11
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "voice"]
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [2160, 2427]
sha256: d8c89e254257f1042a85f10b568456dd3823b3790f0285b5cb11f5d613d66c4b
---

# Guide ultra-complet — Huawei eKit S310

```
system-view
# Sur un port d'accès utilisateur (binding dynamique via DHCP snooping)
interface GigabitEthernet0/0/5
 ip source check user-bind enable       # IPSG
quit
save
```

> La disponibilité exacte de DAI (`arp anti-attack` / `dai enable` selon version)
> et d'IPSG sur le S310 est **à vérifier sur la version logicielle du modèle
> exact** — le datasheet annonce le binding et la défense anti-ARP, qui sont les
> briques de base. Teste en maquette avant production.

⚠️ **Ordre de déploiement :** 1) DHCP snooping, 2) vérifie que les bindings se
remplissent, 3) **ensuite seulement** IPSG/DAI. Activer IPSG sans bindings =
couper tout le monde, y compris toi.

⚠️ **Équipements en IP fixe** : sans binding statique (section 76), IPSG les
bloque. Fais la liste des IP fixes **avant** d'activer.

---

## 78. ACL : le filtrage L2/L3/L4

Les ACL du S310 filtrent sur MAC, IP, ports TCP/UDP, VLAN (valeur datasheet :
« TCP/UDP port number, protocol type, or VLAN »).

**Exemple 1 : le VLAN invités (30) ne doit joindre que la passerelle Internet,
pas les autres VLAN.**

```
system-view
acl number 3000
 description ISOLEMENT_GUEST
 rule 10 permit ip source 192.168.30.0 0.0.0.255 destination 192.168.30.1 0
 rule 20 deny ip source 192.168.30.0 0.0.0.255 destination 192.168.0.0 0.0.255.255
 rule 30 permit ip
quit
# Application sur l'interface VLANIF du VLAN invités (là où le routage se fait)
interface Vlanif30
 traffic-filter inbound acl 3000
quit
save
```

**Exemple 2 : protéger le management (complément section 117).**

```
system-view
acl number 2000
 description MGMT_SEUL_ADMIN
 rule 10 permit source 192.168.99.0 0.0.0.255
 rule 20 deny
quit
# Appliquée aux accès distants (http/ssh/snmp) selon la version
quit
save
```

**Règles de survie ACL :**

1. Les règles se lisent **dans l'ordre** : la première qui matche gagne.
2. **Fin implicite = deny** : tout ce qui n'est pas permis est bloqué.
   Termine toujours par un `permit` explicite de contrôle ou assume le deny.
3. Teste **toujours** depuis un poste du VLAN concerné après application.
4. ⚠️ Une ACL sur le VLANIF de management peut **te couper ton propre accès** :
   applique-la en étant en console, ou prévois une fenêtre de rollback.

🔧 Vérifications :

```
display acl 3000
display traffic-filter applied-record
```

---

## 79. 802.1X, MAC auth et portail : l'authentification d'accès

Le S310 annonce **IEEE 802.1X**, authentification MAC, portail, et AAA via
**RADIUS / HWTACACS** (valeur datasheet). Principe : un port ne s'ouvre qu'après
authentification de l'équipement ou de l'utilisateur.

**Exemple 802.1X minimal (avec serveur RADIUS existant) :**

```
system-view
dot1x enable
#
radius-server template RADIUS_CORP
 radius-server shared-key cipher Exemple_Cle_Radius_2026!
 radius-server authentication 192.168.50.10 1812
quit
aaa
 authentication-scheme AUTH_R
  authentication-mode radius
 quit
 domain corp.lan
  authentication-scheme AUTH_R
  radius-server RADIUS_CORP
 quit
quit
#
interface GigabitEthernet0/0/5
 dot1x enable
 dot1x port-method portbased
quit
save
```

> La syntaxe exacte varie selon la version (template RADIUS, `dot1x` vs
> `authentication`) : **à vérifier sur la version logicielle du modèle exact**
> et à tester en maquette avec un vrai serveur (NPS, FreeRADIUS...).

**Quand déployer quoi :**

| Besoin | Solution |
|---|---|
| Postes du domaine, users connus | 802.1X (EAP) |
| Imprimantes, téléphones sans 802.1X | **MAC auth** (MAB) |
| Invités / BYOD ponctuel | Portail captif |
| Rien de tout ça en place | Port security (section 72) en attendant |

⚠️ **Ne déploie pas 802.1X un vendredi à 17h.** Prévois une phase pilote
(quelques ports), un bypass d'urgence documenté, et un moyen de revenir en
arrière port par port.

---

## 80. Protections CPU et anti-attaques : ce qui tourne en fond

Le datasheet annonce : défense contre **DoS** (SYN flood, Land, Smurf, ICMP
flood), attaques **ARP/ICMP**, **apprentissage ARP strict** (anti-épuisement
de la table), CPU defense.

```
system-view
# Exemples de principes (syntaxe à vérifier sur la version exacte)
arp anti-attack entry-check fixed-mac enable
# Limitation du nombre de MAC apprises par port (complément port-security)
quit
save
```

✅ La plupart de ces protections sont **actives par défaut** ou à activer en
une commande. Vérifie leur état (`display cpu-defend statistics`,
`display arp anti-attack statistics` — noms à vérifier sur la version exacte)
lors de ton audit initial.

**Blackhole MAC (bloquer un équipement malveillant partout, immédiatement) :**

```
system-view
mac-address blackhole 00aa-bbcc-ddee vlan 10
quit
save
```

🔧 Pour lever : `undo mac-address blackhole 00aa-bbcc-ddee vlan 10`.
C'est le « bouton rouge » quand tu as identifié une MAC source d'attaque :
rapide, radical, réversible.

---

## 81. LLDP : voir qui est branché où (sans se déplacer)

**LLDP** (Link Layer Discovery Protocol) fait s'annoncer les équipements entre
eux. Le S310 supporte **LLDP et LLDP-MED** (valeur datasheet) — LLDP-MED sert
aussi à la voice VLAN auto avec les téléphones.

```
system-view
lldp enable
quit
save
```

🔧 **La commande la plus rentable du guide :**

```
display lldp neighbor brief
```

Tu vois, port par port : le **nom**, le **modèle** et le **port distant** de
l'équipement en face (switch, AP, téléphone...). Fini le « c'est quoi ce câble
qui part vers le plafond ? ».

✅ **Active LLDP partout, tout le temps.** Coût : nul. Gain : cartographie
automatique du réseau. Tes outils de supervision (et l'app eKit) s'en servent
pour dessiner la topologie.

---

## 82. SNMP : v2c pour commencer, v3 en cible

Le S310 parle **SNMPv1/v2c/v3** (valeur datasheet). v2c = simple mais communauté
en clair ; **v3 = chiffré et authentifié**, l'objectif.

**SNMPv2c (mise en route rapide) :**

```
system-view
snmp-agent
snmp-agent community read cipher Exemple_Communaute_Lecture_2026! mib-view iso-view
snmp-agent sys-info version v2c
quit
save
```

**SNMPv3 (recommandé en production) :**

```
system-view
snmp-agent
snmp-agent sys-info version v3
snmp-agent group v3 GROUP_ADMIN privacy
snmp-agent usm-user v3 USER_SUP privacy
snmp-agent usm-user v3 USER_SUP authentication-mode sha Exemple_Auth_2026!
snmp-agent usm-user v3 USER_SUP privacy-mode aes128 Exemple_Priv_2026!
snmp-agent usm-user v3 USER_SUP group GROUP_ADMIN
quit
save
```

> Les mots de passe ci-dessus sont **fictifs** : génère les tiens (20+ caractères
> aléatoires) et stocke-les au coffre. La syntaxe exacte USM est **à vérifier sur
> la version logicielle du modèle exact**.

**Ce qu'on supervise en priorité :** état des ports, erreurs/compteurs, CPU,
mémoire, température, état PoE, traps de changement d'état.

🔧 Test depuis le superviseur : `snmpwalk -v3 ... sysDescr`. Si ça ne répond pas :
ACL (section 78), communauté, pare-feu, ou `snmp-agent` oublié.

---

## 83. Syslog : centraliser les logs (sinon ils ne servent à rien)

Les logs en local (`display logbuffer`) sont écrasés par rotation. Envoie-les
vers un **serveur syslog** (souvent le superviseur lui-même).

```
system-view
info-center enable
info-center loghost 192.168.50.20 facility local7
info-center source default channel 2 log level informational state on
quit
save
```

🔧 Vérifications :

```
display info-center
display logbuffer
```

**Niveaux utiles :** `debugging` (trop verbeux en prod), `informational`
(le bon compromis), `warning` (alertes), `critical`.

✅ **À faire dès le jour 1**, avec le NTP (section 24) : des logs horodatés et
centralisés, c'est la moitié du dépannage. Sans ça, tu enquêtes à l'aveugle.

---

## 84. Supervision via eKit cloud : ce que ça apporte

