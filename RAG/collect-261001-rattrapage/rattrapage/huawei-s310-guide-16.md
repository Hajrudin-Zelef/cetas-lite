---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-16
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [3205, 3410]
sha256: 895b054ad06e84a6554a1c01b0588c7518223283703df05a6a651d92b6efdf19
---

# Guide ultra-complet — Huawei eKit S310

| Pièce | Quantité conseillée | Pourquoi |
|---|---|---|
| Switch S310 identique (ou modèle supérieur) | 1 pour ~10 switches en prod | Un ventilateur intégré HS = SAV = jours sans réseau sinon |
| Modules SFP/SFP+ compatibles | 2 par type utilisé | L'uplink fibre est un point unique de défaillance |
| Jarretières fibre + cuivre | assortiment | 80 % des pannes couche 1 |
| Cordon secteur + sangle | 1–2 | Ça disparaît mystérieusement |
| Câble console + adaptateur USB | 2 kits | La porte de secours |
| Ventilateur externe / clim d'appoint | selon local | Canicule : un switch qui chauffe, c'est un switch qui meurt |

✅ **Le switch de secours doit avoir :** la **même version logicielle** que la
prod (ou être mis à jour avant stockage) et une **copie récente de la config**
du switch qu'il remplace. Un spare avec un firmware de 3 ans d'écart = 2 heures
de mise à jour le jour J. Vérifie-le **une fois par an** (il boote ?).

---

## 120. Durcissement (1/4) : comptes et mots de passe

1. **Compte admin nominatif par personne**, pas de compte « admin » partagé :
   ```
   system-view
   aaa
    local-user TECH_MARIE password cipher Exemple_MotDePasse_Fictif_2026!
    local-user TECH_MARIE privilege level 3
    local-user TECH_MARIE service-type ssh https
   quit
   save
   ```
   > Identifiants **fictifs** : génère des mots de passe uniques de 16+ caractères
   > au coffre (Bitwarden, Keepass...).
2. **Niveau de privilège** : 0 (visite) à 3 (admin) — donne le minimum nécessaire.
   Le stagiaire n'a pas besoin du niveau 3.
3. **Désactive les comptes inutilisés** immédiatement au départ d'un technicien.
4. **Renouvelle** les mots de passe admin au moins une fois par an (et après
   chaque départ).
5. ⚠️ Ne stocke jamais un mot de passe **en clair** dans un fichier partagé, un
   post-it ou un ticket. Coffre d'entreprise, point.

---

## 121. Durcissement (2/4) : SSH only, bye-bye Telnet et HTTP

```
system-view
# --- SSH actif, Telnet interdit ---
stelnet server enable
undo telnet server enable
# --- HTTPS actif, HTTP pur coupé ---
http secure-server enable
undo http server enable
# --- Timeout de session : pas de console laissée ouverte ---
user-interface vty 0 4
 authentication-mode aaa
 protocol inbound ssh
 idle-timeout 10 0
quit
user-interface console 0
 idle-timeout 10 0
quit
save
```

> La syntaxe exacte (`stelnet server enable`, `protocol inbound ssh`) est
> **à vérifier sur la version logicielle du modèle exact**.

✅ **Après :** teste que SSH fonctionne **avant** de fermer ta session console.
Te couper ton propre accès distant, c'est le bizutage classique.

---

## 122. Durcissement (3/4) : restreindre l'accès au management

**Principe :** l'administration (SSH/HTTPS/SNMP) n'est joignable que depuis le
VLAN/poste d'administration, jamais depuis les VLAN utilisateurs ou invités.

```
system-view
acl number 2000
 description ACCES_MGMT_ADMIN_ONLY
 rule 10 permit source 192.168.99.0 0.0.0.255     # VLAN management
 rule 20 deny
quit
# Application aux services (selon version : http / ssh / snmp)
http acl 2000
ssh server acl 2000
snmp-agent acl 2000
quit
save
```

> L'application d'ACL aux services (`http acl`, `ssh server acl`) est **à vérifier
> sur la version logicielle du modèle exact**. À défaut : ACL sur le VLANIF de
> management (section 78) — mais appliquée **en étant en console**.

**Et SNMP :** passe en **v3** (section 82), change les communautés v2c par défaut
si v2c il y a, et ne diffuse les traps que vers le superviseur.

---

## 123. Durcissement (4/4) : services, bannières, divers

- [ ] **Bannière légale** : `header login information "Acces reserve..."` —
  dissuasif et utile juridiquement.
- [ ] **Désactive les services inutilisés** : ce qui n'est pas utilisé est une
  surface d'attaque (FTP/TFTP server, Telnet déjà coupé...). `display` l'état
  des services et coupe.
- [ ] **LLDP** : utile en interne ; si le switch est dans une zone publique
  accessible, évalue (il annonce le modèle et le hostname).
- [ ] **CDP** (si présent) : à désactiver en environnement non-Cisco.
- [ ] **Mots de passe chiffrés** : utilise toujours `cipher`, jamais `simple`.
- [ ] **Logs d'accès** : `display logbuffer` doit tracer les connexions —
  vérifie après durcissement.
- [ ] **Sauvegarde post-durcissement** (section 88) : un durcissement non
  sauvegardé est un durcissement qui disparaîtra au prochain reboot.

🔧 **Audit express (15 min) :** `display current-configuration | include` sur
`telnet`, `http server` (non secure), `community`, `password simple` → tout ce
qui ressort est à corriger.

---

## 124. Pense-bête de poche : les 30 commandes qui sauvent

```
# --- État général ---
display version | display device | display clock | display cpu-usage | display memory-usage
display interface brief              # l'état de tous les ports en un écran
display interface GigabitEthernet0/0/5
display logbuffer                    # les logs récents
display current-configuration | display saved-configuration
save                                 # SAUVEGARDER !

# --- VLAN ---
display vlan | display vlan 10
display port vlan | display port vlan GigabitEthernet0/0/5
display mac-address vlan 10 | display mac-address interface GigabitEthernet0/0/5

# --- STP / Trunk ---
display stp brief
display eth-trunk 10

# --- PoE ---
display poe power-state | display poe interface all

# --- IP / routage ---
display ip interface brief | display ip routing-table
ping 192.168.10.1 | tracert 8.8.8.8

# --- Sécurité ---
display dhcp snooping user-bind all
display port-security

# --- Supervision ---
display lldp neighbor brief
display acl 3000

# --- Maintenance fichiers ---
dir flash:/ | display startup

# --- Actions ---
reboot                               # en heure creuse, après save !
reset counters interface GigabitEthernet0/0/5
test cable GigabitEthernet0/0/5      # VCT : coupe brièvement le lien
```

✅ **Imprime cette section**, plastifie-la, scotche-la dans la baie. Le jour où
tu en auras besoin, tu n'auras pas le temps de chercher le PDF.

---

## 125. Glossaire

| Terme | Signification |
|---|---|
| 802.1Q | Standard de taggage des VLAN |
| 802.1X | Authentification d'accès au réseau (port-based) |
| AAA | Authentication, Authorization, Accounting |
| ACL | Access Control List — règles de filtrage |
| BPDU | Trames d'échange du Spanning Tree |
| DAI | Dynamic ARP Inspection — vérifie les ARP contre le binding |
| DSCP | Marquage de priorité dans l'en-tête IP (QoS L3) |
| DHCP snooping | Filtrage des serveurs DHCP non autorisés + table de binding |
| ERPS | Protection en anneau (ITU-T G.8032) |
| Eth-Trunk | Agrégation de liens Huawei (équivalent EtherChannel) |
| HOUP | Huawei Online Upgrade Platform (mises à jour intelligentes) |
| IPSG | IP Source Guard — filtre selon le binding IP/MAC/port |
| iStack | Empilement logique de switches Huawei en un seul |
| LACP | Protocole de négociation d'agrégation (802.3ad) |
| LLDP / LLDP-MED | Découverte des voisins / variante pour la téléphonie |
| MQC | Modular QoS : classifier → behavior → policy |
| MSTP / RSTP / STP | Spanning Tree : multi-instance / rapide / historique |
| OUI | 3 premiers octets d'une MAC = identifiant constructeur |
| PD | Powered Device — équipement alimenté en PoE |
| PSE | Power Sourcing Equipment — le switch qui fournit le PoE |
| PVID | VLAN attribué aux trames non taggées entrant par un port |
| RMON | Surveillance distante (statistiques SNMP) |
| SFP / SFP+ | Modules fibre 1G / 10G |
| SNMP | Protocole de supervision (v1/v2c/v3) |
| Storm control | Plafonnement broadcast/multicast/unicast inconnu |
| VBST | VLAN-based Spanning Tree (spécificité Huawei) |
| VLANIF | Interface virtuelle L3 rattachée à un VLAN |
| VCT | Virtual Cable Test — testeur de câble intégré |
| VRP | Versatile Routing Platform — OS des équipements Huawei |

---

