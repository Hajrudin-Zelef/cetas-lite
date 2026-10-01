---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-14
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [2387, 2555]
sha256: 73b044a9fad563e430591e5c08c2a5239bccd5a584be1fbf0e88ca6065b63be7
---

# --- Supervision ---
[AGENCE-DAKAR-AR720]ntp-service unicast-server 192.168.0.5
# Serveur NTP du siège, joignable via le tunnel (configuré ci-dessous).
[AGENCE-DAKAR-AR720]snmp-agent
[AGENCE-DAKAR-AR720]snmp-agent sys-info version v3
[AGENCE-DAKAR-AR720]snmp-agent group v3 supervision-group privacy read-view iso-view
[AGENCE-DAKAR-AR720]snmp-agent mib-view included iso-view iso
[AGENCE-DAKAR-AR720]snmp-agent usm-user v3 supervision-user group supervision-group acl 2001
[AGENCE-DAKAR-AR720]snmp-agent usm-user v3 supervision-user authentication-mode sha MotDePasseAuthFictif privacy-mode aes128 MotDePassePrivFictif
[AGENCE-DAKAR-AR720]info-center enable
[AGENCE-DAKAR-AR720]info-center loghost 10.0.0.101
[AGENCE-DAKAR-AR720]acl number 2001
[AGENCE-DAKAR-AR720-acl-basic-2001]rule 5 permit source 10.0.0.100 0
[AGENCE-DAKAR-AR720-acl-basic-2001]quit
[AGENCE-DAKAR-AR720]acl number 2005
[AGENCE-DAKAR-AR720-acl-basic-2005]rule 5 permit source 192.168.10.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-basic-2005]rule 10 permit source 10.0.0.0 0.0.255.255
[AGENCE-DAKAR-AR720-acl-basic-2005]quit

# --- Durcissement de base ---
[AGENCE-DAKAR-AR720]undo telnet server enable
[AGENCE-DAKAR-AR720]dns resolve
[AGENCE-DAKAR-AR720]dns server 8.8.8.8
[AGENCE-DAKAR-AR720]dns server 1.1.1.1

# --- SAUVEGARDE ---
# <AGENCE-DAKAR-AR720>save
# Ne jamais oublier. Vérifier ensuite display saved-configuration.
```

Ordre de vérification après application : `display pppoe-client session summary` →
`display ip routing-table` → ping 8.8.8.8 depuis un PC → `display ike sa` /
`display ipsec sa` → `save`.

## 121. Cas pratique n°2 — Site avec tunnel GRE over IPSec + OSPF

Scénario : deux sites qui doivent router dynamiquement. On reprend l'agence ci-dessus et
on ajoute le tunnel GRE chiffré par IPSec, avec OSPF par-dessus.

Côté agence (complément à la config du cas n°1) :

```
# --- Tunnel GRE ---
[AGENCE-DAKAR-AR720]interface Tunnel 0/0/1
[AGENCE-DAKAR-AR720-Tunnel0/0/1]ip address 172.16.100.2 255.255.255.252
[AGENCE-DAKAR-AR720-Tunnel0/0/1]tunnel-protocol gre
[AGENCE-DAKAR-AR720-Tunnel0/0/1]source Dialer 1
[AGENCE-DAKAR-AR720-Tunnel0/0/1]destination 197.155.20.10
[AGENCE-DAKAR-AR720-Tunnel0/0/1]mtu 1400
[AGENCE-DAKAR-AR720-Tunnel0/0/1]tcp adjust-mss 1360
[AGENCE-DAKAR-AR720-Tunnel0/0/1]ospf network-type p2p
[AGENCE-DAKAR-AR720-Tunnel0/0/1]quit
# MTU réduite : GRE (24 o) + IPSec (~60 o) mangent de la place.

# --- IPSec qui chiffre le GRE ---
[AGENCE-DAKAR-AR720]acl number 3101
[AGENCE-DAKAR-AR720-acl-adv-3101]rule 5 permit gre source 197.155.10.34 destination 197.155.20.10
[AGENCE-DAKAR-AR720-acl-adv-3101]quit
# (ike proposal/peer, ipsec proposal et ipsec policy : voir sections 55-56,
#  avec security acl 3101 au lieu de 3100, appliquée sur Dialer 1.)

# --- OSPF par-dessus ---
[AGENCE-DAKAR-AR720]interface LoopBack 0
[AGENCE-DAKAR-AR720-LoopBack0]ip address 10.255.0.11 255.255.255.255
[AGENCE-DAKAR-AR720-LoopBack0]quit
[AGENCE-DAKAR-AR720]ospf 1 router-id 10.255.0.11
[AGENCE-DAKAR-AR720-ospf-1]silent-interface Dialer 1
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]area 0
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]network 192.168.10.0 0.0.0.255
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]network 172.16.100.0 0.0.0.3
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]network 10.255.0.11 0.0.0.0
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]authentication-mode md5 1 cipher CleOSPFfictive
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]quit
[AGENCE-DAKAR-AR720-ospf-1]quit
```

Vérifications : `display ipsec sa` (trafic GRE chiffré) → `display ospf peer` (voisin
Full) → `display ip routing-table protocol ospf` (routes du siège apprises).

## 122. Cas pratique n°3 — Exposer un serveur : NAT server + politique

Scénario : caméra/NVR interne 192.168.30.10 à rendre accessible depuis Internet sur le
port 8080, WAN en IP statique 197.155.10.34.

```
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/1
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/1]ip address 197.155.10.34 255.255.255.252
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/1]nat server protocol tcp global 197.155.10.34 8080 inside 192.168.30.10 8080
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/1]quit
[AGENCE-DAKAR-AR720]security-policy
[AGENCE-DAKAR-AR720-policy-security]rule name WAN-vers-NVR
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-NVR]source-zone untrust
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-NVR]destination-zone trust
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-NVR]destination-address 192.168.30.10 32
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-NVR]service protocol tcp destination-port 8080
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-NVR]action permit
[AGENCE-DAKAR-AR720-policy-security-rule-WAN-vers-NVR]quit
[AGENCE-DAKAR-AR720-policy-security]quit
```

Test : depuis un téléphone en 4G (jamais depuis le LAN — le NAT hairpin peut fausser le
test), ouvrir `http://197.155.10.34:8080`. Vérifier `display nat server` et
`display firewall session table`.

## 123. Cas pratique n°4 — QoS voix sur lien contraint

Scénario : lien montant 20 Mbit/s, 10 téléphones IP (VLAN 30), bureautique VLAN 10.
Objectif : la voix ne saccade jamais, même quand quelqu'un envoie un gros fichier.

```
[AGENCE-DAKAR-AR720]acl number 3003
[AGENCE-DAKAR-AR720-acl-adv-3003]rule 5 permit udp source 192.168.30.0 0.0.0.255 destination-port range 10000 20000
[AGENCE-DAKAR-AR720-acl-adv-3003]quit
[AGENCE-DAKAR-AR720]traffic classifier VOIX
[AGENCE-DAKAR-AR720-classifier-VOIX]if-match acl 3003
[AGENCE-DAKAR-AR720-classifier-VOIX]quit
[AGENCE-DAKAR-AR720]traffic classifier DATA
[AGENCE-DAKAR-AR720-classifier-DATA]if-match any
[AGENCE-DAKAR-AR720-classifier-DATA]quit
[AGENCE-DAKAR-AR720]traffic behavior PRIORITE-VOIX
[AGENCE-DAKAR-AR720-behavior-PRIORITE-VOIX]queue pq
[AGENCE-DAKAR-AR720-behavior-PRIORITE-VOIX]quit
[AGENCE-DAKAR-AR720]traffic behavior LIMITE-DATA
[AGENCE-DAKAR-AR720-behavior-LIMITE-DATA]car cir 16000 pir 18000 green pass yellow pass red discard
[AGENCE-DAKAR-AR720-behavior-LIMITE-DATA]quit
[AGENCE-DAKAR-AR720]traffic policy QOS-VOIX
[AGENCE-DAKAR-AR720-trafficpolicy-QOS-VOIX]classifier VOIX behavior PRIORITE-VOIX
[AGENCE-DAKAR-AR720-trafficpolicy-QOS-VOIX]classifier DATA behavior LIMITE-DATA
[AGENCE-DAKAR-AR720-trafficpolicy-QOS-VOIX]quit
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]traffic-policy QOS-VOIX outbound
[AGENCE-DAKAR-AR720-Dialer1]quit
```

Logique : la voix passe en priorité absolue ; le reste est plafonné à 16 Mbit/s sur un
lien de 20, ce qui laisse toujours de la place à la voix. **Ordre des classifiers** :
le plus spécifique d'abord (VOIX avant DATA), sinon tout matche `if-match any`.

---
---

# Partie O — Pense-bête de poche

## 124. Les 20 commandes à connaître par cœur

```
display version                    # modèle, VRP, uptime
display current-configuration      # config en cours
display saved-configuration        # config sauvegardée
save                               # sauvegarder !
display ip interface brief         # état + IP des interfaces
display interface GE 0/0/0         # détail un port
display ip routing-table           # table de routage
display ospf peer                  # voisins OSPF
display ike sa                     # phase 1 IPSec
display ipsec sa                   # phase 2 IPSec
display pppoe-client session summary
display nat session table          # sessions NAT actives
display nat outbound               # règles NAT sortant
display firewall session table     # sessions firewall
display security-policy all        # politiques
display cpu-usage / display memory
display logbuffer                  # derniers logs
display clock / display ntp-service status
reboot                             # (vue utilisateur)
```

## 125. Séquence de mise en service (ordre chronologique)

