---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-15
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [2556, 2700]
sha256: cc9785ec26d1e5021652654049930ad4d4d60c49206367d3c3f6044ebe213b61
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

1. Console → hostname → mots de passe → `save`.
2. Firmware : vérifier/valider la version, upgrade si besoin.
3. WAN (PPPoE/DHCP/statique) → tester le ping Internet depuis le routeur.
4. LAN + DHCP → tester depuis un PC.
5. NAT (`nat outbound` + ACL) → tester la navigation.
6. Zones + politiques → re-tester (le firewall peut tout casser).
7. VPN → `display ike sa` / `display ipsec sa` → ping inter-sites.
8. QoS si besoin, supervision (NTP/SNMP/syslog), durcissement.
9. `save` + sauvegarde externe + dossier de site à jour.

## 126. Dossier de site : ce qu'il doit contenir

- Plan d'adressage (VLAN, IP, DHCP, réservations).
- Identifiants FAI (dans le coffre, pas dans le dossier partagé).
- IP WAN, passerelles, DNS.
- PSK VPN (coffre), paramètres IKE/IPSec.
- Version firmware, numéro de série, date d'achat, garantie.
- Sauvegarde de config datée.
- Résultats des tests (débit, basculement, temps de coupure).
- Contacts FAI + contrat.

## 127. Les 10 erreurs qui reviennent tout le temps

1. Oublier `save` → config perdue au reboot.
2. `nat outbound` sur la mauvaise interface (GE au lieu du Dialer).
3. Exclusion NAT oubliée pour le trafic VPN.
4. Politique untrust→local sans IKE → tunnel qui ne monte pas.
5. Crypto ACL non miroir → phase 2 KO.
6. `dhcp enable` oublié → DHCP muet.
7. Réseau oublié dans l'ACL du NAT → VLAN sans Internet.
8. `silent-interface` oublié → Hello OSPF vers Internet.
9. Horloge fausse → certificats/logs inutilisables.
10. Pas de `tcp adjust-mss` sur PPPoE → sites qui ne chargent pas.

---
---

# Partie P — Glossaire

## 128. Glossaire (A–Z)

- **ACL (Access Control List)** : liste de règles qui filtrent le trafic (ici : qui peut
  être NATé, quoi mettre dans le tunnel…).
- **AR720** : routeur de branche Huawei, série NetEngine AR700, 2 WAN combo + 8 LAN,
  2 slots SIC.
- **BPDU** : trames d'échange du protocole STP entre switchs.
- **CAR (Committed Access Rate)** : mécanisme de limitation de débit (policing).
- **CHAP/PAP** : protocoles d'authentification PPP (CHAP = challenge, plus sûr ; PAP =
  mot de passe en clair).
- **Combo (port)** : port RJ45 + cage SFP partageant la même interface logique.
- **Dialer (interface)** : interface logique VRP pour le PPPoE.
- **DMZ** : zone réseau intermédiaire pour les équipements exposés ou isolés (ici aussi
  les invités).
- **DPD (Dead Peer Detection)** : détection d'un pair IPSec mort.
- **DSCP** : marquage de qualité de service dans l'en-tête IP.
- **GRE** : protocole de tunnelisation (encapsulation IP dans IP).
- **IKE (Internet Key Exchange)** : protocole de négociation des clés IPSec (phase 1) ;
  v1 et v2.
- **IPSec** : suite de protocoles de chiffrement/authentification au niveau IP (ESP/AH).
- **L2TP** : protocole de tunnel de niveau 2, souvent combiné à IPSec pour les nomades.
- **LLQ/PQ** : file d'attente prioritaire (voix).
- **MSS (Maximum Segment Size)** : taille max d'un segment TCP ; à ajuster sous la MTU.
- **MTU** : taille maximale d'un paquet sur un lien.
- **NAPT** : NAT avec translation de ports (plusieurs privés → une publique).
- **NAT server** : redirection d'un port public vers un serveur interne (port mapping).
- **NetStream** : export de flux façon NetFlow (Huawei).
- **NQA** : sondes de qualité/disponibilité (ping, jitter…) + objet `track`.
- **OSPF** : protocole de routage dynamique à état de liens.
- **PBR (Policy-Based Routing)** : routage selon des critères (source…) au lieu de la
  table de routage.
- **PFS (Perfect Forward Secrecy)** : renouvellement des clés Diffie-Hellman en phase 2.
- **PPPoE** : PPP sur Ethernet (authentification FAI sur ADSL/fibre).
- **PSK (Pre-Shared Key)** : clé partagée pour authentifier un tunnel IPSec.
- **SA (Security Association)** : association de sécurité négociée (IKE phase 1 / IPSec
  phase 2).
- **SIC (Smart Interface Card)** : carte d'extension (4G, E1, voix…) des slots AR720.
- **SNMP** : protocole de supervision ; v3 = version chiffrée/authentifiée.
- **STP** : protocole anti-boucle de niveau 2.
- **Track** : objet VRP qui suit l'état d'une sonde NQA et pilote les routes.
- **VLAN** : réseau local virtuel (segmentation logique).
- **Vlanif** : interface virtuelle VRP portant l'IP d'un VLAN.
- **VRP (Versatile Routing Platform)** : système d'exploitation des routeurs Huawei.
- **VTY** : lignes virtuelles pour l'accès distant (SSH/telnet).
- **WSIC** : slot pour carte Wi-Fi intégrée (option selon référence).

---
---

# Partie Q — Quiz (10 questions)

## 129. Questions

1. Quelle commande sauvegarde la configuration en cours pour qu'elle survive au reboot ?
2. Sur l'AR720, combien y a-t-il de ports WAN combo fixes et de ports LAN fixes ?
3. Pourquoi faut-il appliquer le `nat outbound` sur l'interface Dialer (et pas sur
   GE0/0/0) quand le WAN est en PPPoE ?
4. Un tunnel IPSec monte en phase 1 mais pas en phase 2. Citez deux causes probables.
5. À quoi sert l'exclusion NAT (règle `deny` dans l'ACL du `nat outbound`) dans un
   contexte VPN site-à-site ?
6. Que fait la commande `track 1 nqa admin pppoe-check` liée à une route statique ?
7. Pourquoi configure-t-on `tcp adjust-mss` sur une interface Dialer PPPoE ?
8. Quelle politique de sécurité faut-il impérativement pour que l'IKE aboutisse au
   routeur ?
9. Un voisin OSPF reste bloqué en ExStart. Quelle est la cause la plus probable ?
10. Pourquoi ne faut-il jamais exposer SNMP v1/v2c sur le WAN, et que faut-il utiliser
    à la place ?

## 130. Réponses

1. `save` (en vue utilisateur `<AR720>`). Sans ça, la config est perdue au reboot.
2. 2 ports WAN combo (RJ45/SFP) + 8 ports LAN RJ45 (reconfigurables en WAN).
3. Parce que c'est l'interface Dialer qui porte l'IP publique négociée en PPP ; le NAT
   doit s'appliquer sur l'interface de sortie qui a l'adresse publique.
4. Crypto ACL non miroir entre les deux sites ; transform-set (chiffrement/hash)
   incompatible ; PFS différent ; politique IPSec appliquée sur la mauvaise interface.
5. À empêcher que le trafic destiné au site distant soit traduit (NATé) avant d'entrer
   dans le tunnel — sinon le tunnel ne reconnaît plus le trafic et rien ne passe.
6. Elle retire automatiquement la route statique (ici la route par défaut via le WAN
   principal) quand la sonde NQA échoue, permettant à la route de secours de prendre
   le relais.
7. Parce que l'en-tête PPPoE réduit la MTU utile (1492) ; sans ajustement du MSS TCP,
   certains flux se fragmentent ou se bloquent (pages web qui ne chargent pas).
8. Une règle `untrust → local` qui autorise UDP 500 et UDP 4500 (IKE / NAT-T) vers le
   routeur lui-même.
9. Une MTU différente entre les deux interfaces (classique), ou un router-id dupliqué.
10. Parce que la communauté transite en clair (interception + usurpation faciles) ;
    utiliser SNMPv3 (authentification + chiffrement), restreint par ACL au superviseur.

---
---

# Partie R — Pour aller plus loin

## 131. Monter en gamme et autour du produit

- **Grand frère AR730** : mêmes slots, 6 Gbit/s sortants, 1 port 10GE optique en plus,
  1200 terminaux — si l'AR720 sature (CPU, débit), c'est la marche suivante logique
  (**à valider avec le fournisseur** selon les prix et la dispo).
- **SD-WAN Huawei** : pour un parc de dizaines d'agences, l'orchestration centralisée
  (iMaster NCE) change la vie par rapport au CLI site par site.
- **Contrôleur WLAN intégré** : l'AR720 peut gérer des AP Huawei (CAPWAP) — pratique
  pour une agence avec 2–5 bornes sans contrôleur dédié.

## 132. Documentation officielle à garder sous la main

