---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-6
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [967, 1133]
sha256: 96cc56e09a8a36fc5693db834e6380981bc6cc041d86066266b3e3fbe6b0dd22
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

Ça évite l'élection DR/BDR inutile et accélère la montée.

## 48. OSPF : sécuriser et optimiser

```
[AGENCE-DAKAR-AR720]ospf 1 router-id 10.255.0.11
[AGENCE-DAKAR-AR720-ospf-1]silent-interface GigabitEthernet 0/0/0
[AGENCE-DAKAR-AR720-ospf-1]silent-interface Dialer 1
[AGENCE-DAKAR-AR720-ospf-1]area 0
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]authentication-mode md5 1 cipher CleOSPFfictive
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]quit
[AGENCE-DAKAR-AR720-ospf-1]quit
```

- `silent-interface` : ne pas envoyer de Hello OSPF vers Internet (sécurité + CPU).
- `authentication-mode md5` : authentifier les voisins (même clé des deux côtés).
- Les timers Hello/Dead par défaut (10 s / 40 s) conviennent ; ne les toucher que si le
  lien est instable (satellite, 4G) — et **identiques des deux côtés**.

## 49. Redistribution : injecter les statiques dans OSPF

Pour annoncer la route par défaut (ou des statiques) aux voisins OSPF :

```
[AGENCE-DAKAR-AR720]ospf 1
[AGENCE-DAKAR-AR720-ospf-1]default-route-advertise always
[AGENCE-DAKAR-AR720-ospf-1]import-route static
[AGENCE-DAKAR-AR720-ospf-1]quit
```

- `default-route-advertise always` : annonce 0.0.0.0/0 même si le routeur n'a pas lui-même
  de défaut (utile sur un concentrateur). Sans `always`, il ne l'annonce que s'il a une
  route par défaut active.
- Filtrer avec une route-policy pour ne redistribuer que ce qu'il faut (voir section 50).

## 50. Route-policy : filtrer proprement

```
[AGENCE-DAKAR-AR720]ip ip-prefix PLAGE-AGENCES index 10 permit 192.168.0.0 16 greater-equal 24 less-equal 24
[AGENCE-DAKAR-AR720]route-policy FILTRE-OSPF permit node 10
[AGENCE-DAKAR-AR720-route-policy]if-match ip-prefix PLAGE-AGENCES
[AGENCE-DAKAR-AR720-route-policy]quit
[AGENCE-DAKAR-AR720]ospf 1
[AGENCE-DAKAR-AR720-ospf-1]import-route static route-policy FILTRE-OSPF
[AGENCE-DAKAR-AR720-ospf-1]quit
```

Principe : `ip-prefix` = la liste de ce qu'on autorise, `route-policy` = le filtre
appliqué à la redistribution. Sans filtre, on risque d'injecter des routes parasites
(ou des boucles) dans tout le domaine.

## 51. Exemple multi-site simple : agence ↔ siège via VPN

Topologie :

- Siège : 192.168.0.0/24, routeur 10.255.0.1.
- Agence Dakar : 192.168.10.0/24, AR720 10.255.0.11.
- Tunnel GRE over IPSec entre les deux (section 58), OSPF par-dessus.

Côté agence, après avoir monté le tunnel :

```
[AGENCE-DAKAR-AR720]ospf 1 router-id 10.255.0.11
[AGENCE-DAKAR-AR720-ospf-1]area 0
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]network 192.168.10.0 0.0.0.255
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]network 172.16.100.0 0.0.0.3
[AGENCE-DAKAR-AR720-ospf-1-area-0.0.0.0]quit
[AGENCE-DAKAR-AR720-ospf-1]quit
```

où 172.16.100.0/30 est le réseau du tunnel GRE. Le siège déclare son LAN + le même
réseau de tunnel. Résultat : chaque site apprend les routes de l'autre dynamiquement,
et l'ajout d'une 3e agence ne demande qu'à la configurer elle (pas à toucher les autres).

## 52. PBR (Policy-Based Routing) : router selon la source

Parfois on veut qu'un VLAN sorte par un WAN différent (ex. invités → WAN2, bureautique
→ WAN1). C'est du routage par politique :

```
[AGENCE-DAKAR-AR720]acl number 3001
[AGENCE-DAKAR-AR720-acl-adv-3001]rule 5 permit ip source 192.168.20.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-adv-3001]quit
[AGENCE-DAKAR-AR720]policy-based-route INVITES permit node 5
[AGENCE-DAKAR-AR720-policy-based-route-INVITES-5]if-match acl 3001
[AGENCE-DAKAR-AR720-policy-based-route-INVITES-5]apply ip-address next-hop 192.168.8.1
[AGENCE-DAKAR-AR720-policy-based-route-INVITES-5]quit
[AGENCE-DAKAR-AR720]interface Vlanif 20
[AGENCE-DAKAR-AR720-Vlanif20]ip policy-based-route INVITES
[AGENCE-DAKAR-AR720-Vlanif20]quit
```

Le PBR s'applique **avant** la table de routage. Attention : le trafic PBRé ne suit plus
les routes dynamiques — à combiner avec prudence avec le basculement NQA.

---
---

# Partie F — VPN

## 53. Vue d'ensemble : quel VPN pour quel besoin

| Besoin | Solution | Section |
|---|---|---|
| Relier 2 sites (LAN à LAN) | IPSec site-à-site | 54–57 |
| Relier 2 sites + OSPF dynamique | GRE over IPSec | 58 |
| Utilisateur nomade (PC/téléphone) | L2TP over IPSec | 59–60 |
| Accès distant simple (secours) | L2TP pur (sans IPSec) — déconseillé | 59 |

Règle : **IPSec en mode tunnel** pour du site-à-site simple ; **GRE over IPSec** dès qu'on
veut du routage dynamique ou du multicast par-dessus.

## 54. IPSec site-à-site : les briques (IKE phase 1 et 2)

- **Phase 1 (IKE)** : les deux routeurs s'authentifient et négocient un canal sécurisé
  (SA ISAKMP). Paramètres : version IKE (v1/v2), chiffrement, hash, DH, PSK ou certificat.
- **Phase 2 (IPSec)** : négociation des SA de données (ESP). Paramètres : transform-set
  (ESP-AES + SHA), PFS, durée de vie.
- **Crypto ACL** : définit **quel trafic** passe dans le tunnel (le « interesting traffic »).

Les deux côtés doivent avoir des paramètres **miroir** : ce que l'un propose, l'autre
doit l'accepter. 80 % des échecs viennent d'un paramètre qui diffère.

## 55. IPSec site-à-site IKEv2 + PSK : configuration complète — côté agence (AR720)

Hypothèses : agence 192.168.10.0/24, IP WAN 197.155.10.34 (Dialer 1) ; siège
192.168.0.0/24, IP WAN 197.155.20.10. PSK fictive `CleIPSecFictive2026!`.

```
[AGENCE-DAKAR-AR720]ike proposal 10
[AGENCE-DAKAR-AR720-ike-proposal-10]encryption-algorithm aes-256
[AGENCE-DAKAR-AR720-ike-proposal-10]dh group14
[AGENCE-DAKAR-AR720-ike-proposal-10]authentication-algorithm sha2-256
[AGENCE-DAKAR-AR720-ike-proposal-10]integrity-algorithm hmac-sha2-256
[AGENCE-DAKAR-AR720-ike-proposal-10]quit
[AGENCE-DAKAR-AR720]ike peer SIEGE
[AGENCE-DAKAR-AR720-ike-peer-SIEGE]ike-version 2
[AGENCE-DAKAR-AR720-ike-peer-SIEGE]remote-address 197.155.20.10
[AGENCE-DAKAR-AR720-ike-peer-SIEGE]pre-shared-key cipher CleIPSecFictive2026!
[AGENCE-DAKAR-AR720-ike-peer-SIEGE]proposal 10
[AGENCE-DAKAR-AR720-ike-peer-SIEGE]quit
[AGENCE-DAKAR-AR720]ipsec proposal TRANSFO-AGENCE
[AGENCE-DAKAR-AR720-ipsec-proposal-TRANSFO-AGENCE]esp authentication-algorithm sha2-256
[AGENCE-DAKAR-AR720-ipsec-proposal-TRANSFO-AGENCE]esp encryption-algorithm aes-256
[AGENCE-DAKAR-AR720-ipsec-proposal-TRANSFO-AGENCE]quit
[AGENCE-DAKAR-AR720]acl number 3100
[AGENCE-DAKAR-AR720-acl-adv-3100]rule 5 permit ip source 192.168.10.0 0.0.0.255 destination 192.168.0.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-adv-3100]quit
[AGENCE-DAKAR-AR720]ipsec policy CARTE-SIEGE 10 isakmp
[AGENCE-DAKAR-AR720-ipsec-policy-isakmp-CARTE-SIEGE-10]security acl 3100
[AGENCE-DAKAR-AR720-ipsec-policy-isakmp-CARTE-SIEGE-10]ike-peer SIEGE
[AGENCE-DAKAR-AR720-ipsec-policy-isakmp-CARTE-SIEGE-10]proposal TRANSFO-AGENCE
[AGENCE-DAKAR-AR720-ipsec-policy-isakmp-CARTE-SIEGE-10]pfs dh-group14
[AGENCE-DAKAR-AR720-ipsec-policy-isakmp-CARTE-SIEGE-10]quit
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]ipsec policy CARTE-SIEGE
[AGENCE-DAKAR-AR720-Dialer1]quit
```

Et ne pas oublier l'exclusion NAT (section 40) avec une ACL qui nie le trafic vers
192.168.0.0/24 avant le `permit` général.

## 56. IPSec site-à-site : configuration miroir — côté siège

Le siège fait exactement le miroir (crypto ACL inversée, remote-address = IP de l'agence).
Si le siège est aussi un Huawei VRP, la config est symétrique :

