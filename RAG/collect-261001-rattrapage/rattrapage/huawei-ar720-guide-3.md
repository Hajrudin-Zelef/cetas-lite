---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-3
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [346, 525]
sha256: c3474baea2a5103fa91403ebf2a96c308f84dd838d5341ceae1c1475a87631c0
---

# Partie B — WAN : accès Internet

## 16. Comprendre les interfaces WAN de l'AR720

- `GigabitEthernet0/0/0` et `GigabitEthernet0/0/1` : les 2 ports combo WAN.
- `GigabitEthernet0/0/2` à `GigabitEthernet0/0/9` : les 8 ports LAN, **reconfigurables en
  WAN** si besoin (utile pour un 3e lien ou un lien dédié).
- Chaque interface physique peut porter des **sous-interfaces** (VLAN tagging) :
  `GigabitEthernet0/0/0.10`.

Convention d'adressage dans ce guide :

| Usage | Interface | IP / mode |
|---|---|---|
| WAN principal (fibre/ADSL) | GE0/0/0 | PPPoE ou DHCP ou statique selon FAI |
| WAN secours (4G ou 2e FAI) | GE0/0/1 | DHCP client |
| LAN | GE0/0/2…9 (ou Vlanif) | 192.168.10.0/24 |

Adapter à chaque site. Noter le plan d'adressage **avant** de configurer.

## 17. WAN en PPPoE (client) — le cas ADSL/fibre avec identifiants FAI

Le PPPoE se configure via une **interface Dialer**. Étapes :

1. Créer l'interface Dialer avec login/mot de passe.
2. Lier l'interface physique au Dialer (`pppoe-client dial-bundle-number`).
3. NAT sortant sur le Dialer + route par défaut.

```
[AGENCE-DAKAR-AR720]dialer-rule
[AGENCE-DAKAR-AR720-dialer-rule]dialer-rule 1 ip permit
[AGENCE-DAKAR-AR720-dialer-rule]quit
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]dialer user agence.dakar@fai.exemple
[AGENCE-DAKAR-AR720-Dialer1]dialer-group 1
[AGENCE-DAKAR-AR720-Dialer1]dialer bundle 1
[AGENCE-DAKAR-AR720-Dialer1]ppp chap user agence.dakar@fai.exemple
[AGENCE-DAKAR-AR720-Dialer1]ppp chap password cipher MotDePasseFAIFictif
[AGENCE-DAKAR-AR720-Dialer1]ppp pap local-user agence.dakar@fai.exemple password cipher MotDePasseFAIFictif
[AGENCE-DAKAR-AR720-Dialer1]ip address ppp-negotiate
[AGENCE-DAKAR-AR720-Dialer1]nat outbound 2000
[AGENCE-DAKAR-AR720-Dialer1]quit
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/0
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/0]pppoe-client dial-bundle-number 1
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/0]quit
[AGENCE-DAKAR-AR720]ip route-static 0.0.0.0 0.0.0.0 Dialer 1
```

Points d'attention :

- `dialer user` + `ppp chap user` : certains FAI n'utilisent que CHAP, d'autres que PAP.
  Configurer les deux ne gêne pas (négociation).
- `ip address ppp-negotiate` : l'IP est attribuée par le FAI.
- Si le FAI impose un VLAN (ex. fibre avec VLAN 10 sur l'ONT) : créer une sous-interface
  `GE0/0/0.10` avec `dot1q termination vid 10` et y attacher le client PPPoE.
- Vérification : `display pppoe-client session summary`, `display ip interface brief`.

## 18. PPPoE sur sous-interface VLAN (fibre avec tag FAI)

```
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/0.10
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/0.10]dot1q termination vid 10
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/0.10]pppoe-client dial-bundle-number 1
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/0.10]arp broadcast enable
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/0.10]quit
```

Ne pas oublier `arp broadcast enable` sur les sous-interfaces dot1q, sinon certains flux
ne passent pas. C'est un classique des tickets « ça marche en direct sur l'ONT mais pas
derrière le routeur ».

## 19. WAN en DHCP client (box FAI / modem en mode pont+DHCP)

```
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/0
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/0]ip address dhcp-alloc
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/0]nat outbound 2000
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/0]quit
```

La route par défaut est en général poussée par le serveur DHCP du FAI. Vérifier :

```
display ip routing-table
```

Si la route par défaut n'apparaît pas, l'ajouter à la main (voir section suivante).
Astuce : noter l'IP/gateway obtenues (`display ip interface brief`) pour le dossier du site.

## 20. WAN en IP statique (liaison louée / IP fixe FAI)

```
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/1
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/1]ip address 197.155.10.34 255.255.255.252
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/1]nat outbound 2000
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/1]quit
[AGENCE-DAKAR-AR720]ip route-static 0.0.0.0 0.0.0.0 197.155.10.33
```

(IP fictives — remplacer par celles du FAI.) Avec une IP statique, on maîtrise le NAT
entrant (NAT server) et les tunnels VPN : c'est le mode préféré pour un site qui héberge
des services ou termine des VPN.

## 21. Multi-WAN avec basculement : le principe

Deux liens WAN = un principal + un secours. Le basculement repose sur :

1. **Deux routes par défaut** avec des préférences différentes (la plus petite préférence
   gagne).
2. **NQA** (sonde de connectivité) qui détecte la panne réelle du lien principal
   (pas seulement la perte du lien physique).
3. **Track** qui lie la route statique au résultat NQA : si la sonde échoue, la route
   principale est retirée et la route de secours prend le relais.

Sans NQA, une panne « lien physique OK mais pas d'Internet » (panne chez le FAI) ne serait
pas détectée : le routeur garderait la route principale et tout serait coupé. C'est LE
point à ne pas rater.

## 22. Multi-WAN : configuration complète (PPPoE principal + 4G/DHCP secours)

```
[AGENCE-DAKAR-AR720]nqa test-instance admin pppoe-check
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]test-type icmp
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]destination-address ipv4 8.8.8.8
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]frequency 10
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]probe-count 3
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]timeout 2
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]start now
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]quit
[AGENCE-DAKAR-AR720]track 1 nqa admin pppoe-check
[AGENCE-DAKAR-AR720]ip route-static 0.0.0.0 0.0.0.0 Dialer 1 track 1
[AGENCE-DAKAR-AR720]ip route-static 0.0.0.0 0.0.0.0 192.168.8.1 preference 100
```

Explications ligne par ligne :

- `nqa test-instance admin pppoe-check` : crée une sonde nommée.
- `test-type icmp` + `destination-address 8.8.8.8` : ping régulier vers une cible fiable.
  (En production, préférer une cible du FAI ou du siège plutôt qu'un DNS public.)
- `frequency 10` : un test toutes les 10 s. `probe-count 3` : 3 pings par test.
- `track 1 nqa ...` : l'objet track 1 suit l'état de la sonde.
- Route principale via Dialer 1 **liée au track** : si la sonde tombe, la route est retirée.
- Route de secours via la passerelle du lien 2 (ici 192.168.8.1, la box 4G) avec
  préférence 100 (moins prioritaire que la préférence par défaut 60).

Vérification :

```
display nqa results test-instance admin pppoe-check
display track 1
display ip routing-table
```

## 23. Basculement : tester avant de déclarer « ça marche »

Procédure de test obligatoire en fenêtre de maintenance :

1. Noter la route active : `display ip routing-table | include 0.0.0.0`.
2. Débrancher le câble du WAN principal (ou couper la session PPPoE :
   `reset pppoe-client session` — attention, commande en vue utilisateur).
3. Attendre ~30 s, vérifier que la route de secours est active et qu'un ping sort.
4. Rebrancher, vérifier le retour sur le principal (le track remonte, la route
   préférée reprend).
5. Mesurer le temps de coupure réel avec des pings continus depuis un poste.

Noter le temps de basculement dans le dossier du site. Objectif raisonnable : < 60 s.
Si c'est plus long, resserrer `frequency`/`probe-count` (au prix de plus de trafic de
sonde).

## 24. WAN 4G/5G via carte SIC : principes

Si l'AR720 est équipé d'une carte cellulaire SIC (modem 4G LTE ou 5G — **à vérifier sur la
fiche du modèle exact** pour la référence de carte compatible) :

- L'interface apparaît en général comme `Cellular0/0/0` (slot 0) ou similaire.
- Insérer la carte SIM (routeur éteint de préférence, même si le hot swap est supporté).
- Configurer le profil APN du opérateur, le code PIN si nécessaire, puis une interface
  Dialer ou une IP négociée selon le mode du modem.
- Le lien cellulaire sert quasi toujours de **secours** (coût au volume, latence variable).

