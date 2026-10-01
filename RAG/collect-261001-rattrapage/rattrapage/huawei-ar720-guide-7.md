---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-7
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [1134, 1319]
sha256: c68fe50a86cdf74b9f1ce68fabd8842f780f92b8b7c49426a5f34a66287bd05f
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

```
[SIEGE]ike proposal 10
[SIEGE-ike-proposal-10]encryption-algorithm aes-256
[SIEGE-ike-proposal-10]dh group14
[SIEGE-ike-proposal-10]authentication-algorithm sha2-256
[SIEGE-ike-proposal-10]integrity-algorithm hmac-sha2-256
[SIEGE-ike-proposal-10]quit
[SIEGE]ike peer AGENCE-DAKAR
[SIEGE-ike-peer-AGENCE-DAKAR]ike-version 2
[SIEGE-ike-peer-AGENCE-DAKAR]remote-address 197.155.10.34
[SIEGE-ike-peer-AGENCE-DAKAR]pre-shared-key cipher CleIPSecFictive2026!
[SIEGE-ike-peer-AGENCE-DAKAR]proposal 10
[SIEGE-ike-peer-AGENCE-DAKAR]quit
[SIEGE]ipsec proposal TRANSFO-SIEGE
[SIEGE-ipsec-proposal-TRANSFO-SIEGE]esp authentication-algorithm sha2-256
[SIEGE-ipsec-proposal-TRANSFO-SIEGE]esp encryption-algorithm aes-256
[SIEGE-ipsec-proposal-TRANSFO-SIEGE]quit
[SIEGE]acl number 3100
[SIEGE-acl-adv-3100]rule 5 permit ip source 192.168.0.0 0.0.0.255 destination 192.168.10.0 0.0.0.255
[SIEGE-acl-adv-3100]quit
[SIEGE]ipsec policy CARTE-AGENCE 10 isakmp
[SIEGE-ipsec-policy-isakmp-CARTE-AGENCE-10]security acl 3100
[SIEGE-ipsec-policy-isakmp-CARTE-AGENCE-10]ike-peer AGENCE-DAKAR
[SIEGE-ipsec-policy-isakmp-CARTE-AGENCE-10]proposal TRANSFO-AGENCE
[SIEGE-ipsec-policy-isakmp-CARTE-AGENCE-10]pfs dh-group14
[SIEGE-ipsec-policy-isakmp-CARTE-AGENCE-10]quit
[SIEGE]interface GigabitEthernet 0/0/0
[SIEGE-GigabitEthernet0/0/0]ipsec policy CARTE-AGENCE
[SIEGE-GigabitEthernet0/0/0]quit
```

Checklist de symétrie avant de tester : même IKE version, même proposal (AES/SHA/DH),
même PSK, ACL miroir, PFS identique, politique appliquée sur l'interface WAN des deux
côtés.

## 57. Vérifier un tunnel IPSec : les commandes qui comptent

```
display ike sa
display ipsec sa
display ike peer
```

- `display ike sa` : phase 1. On doit voir l'état `ESTABLISHED` (ou équivalent selon
  version). Rien ici = problème phase 1 (connectivité UDP 500/4500, PSK, proposal).
- `display ipsec sa` : phase 2. Des SA avec compteurs qui augmentent = trafic chiffré
  qui passe.
- Test applicatif : `ping -a 192.168.10.1 192.168.0.1` (forcer la source dans le LAN
  pour que le trafic matche la crypto ACL).

Si le tunnel ne monte pas : voir cas n°1 et n°2 du dépannage (Partie K).

## 58. GRE over IPSec : le tunnel qui accepte OSPF et le multicast

GRE encapsule n'importe quel protocole (dont OSPF multicast) dans de l'IP ; IPSec
chiffre le tout. Configuration côté agence (tunnel 172.16.100.0/30) :

```
[AGENCE-DAKAR-AR720]interface Tunnel 0/0/1
[AGENCE-DAKAR-AR720-Tunnel0/0/1]ip address 172.16.100.2 255.255.255.252
[AGENCE-DAKAR-AR720-Tunnel0/0/1]tunnel-protocol gre
[AGENCE-DAKAR-AR720-Tunnel0/0/1]source Dialer 1
[AGENCE-DAKAR-AR720-Tunnel0/0/1]destination 197.155.20.10
[AGENCE-DAKAR-AR720-Tunnel0/0/1]ospf network-type p2p
[AGENCE-DAKAR-AR720-Tunnel0/0/1]quit
[AGENCE-DAKAR-AR720]acl number 3101
[AGENCE-DAKAR-AR720-acl-adv-3101]rule 5 permit gre source 197.155.10.34 destination 197.155.20.10
[AGENCE-DAKAR-AR720-acl-adv-3101]quit
```

Puis une `ipsec policy` dont la `security acl` est l'ACL 3101 (trafic GRE entre les deux
IP publiques), appliquée sur l'interface WAN. Côté siège : miroir (source/destination
inversées, IP tunnel 172.16.100.1).

Avantages : OSPF passe, on peut router dynamiquement, le MTU se règle une fois sur le
tunnel (`mtu 1400` + `tcp adjust-mss 1360` recommandés sur l'interface Tunnel).

## 59. L2TP : serveur d'accès distant pour les nomades

L2TP permet à un PC/telephone avec un client VPN natif de se connecter au réseau de
l'agence/du siège. Configuration serveur L2TP over IPSec (recommandé) côté AR720 :

```
[AGENCE-DAKAR-AR720]l2tp enable
[AGENCE-DAKAR-AR720]l2tp-group 1
[AGENCE-DAKAR-AR720-l2tp-group-1]tunnel password cipher MotDePasseTunnelFictif
[AGENCE-DAKAR-AR720-l2tp-group-1]allow l2tp virtual-template 1 remote client1
[AGENCE-DAKAR-AR720-l2tp-group-1]quit
[AGENCE-DAKAR-AR720]interface Virtual-Template 1
[AGENCE-DAKAR-AR720-Virtual-Template1]ppp authentication-mode chap
[AGENCE-DAKAR-AR720-Virtual-Template1]ip address 192.168.50.1 255.255.255.0
[AGENCE-DAKAR-AR720-Virtual-Template1]quit
[AGENCE-DAKAR-AR720]aaa
[AGENCE-DAKAR-AR720-aaa]local-user nomade1 password cipher MotDePasseNomadeFictif
[AGENCE-DAKAR-AR720-aaa]local-user nomade1 service-type ppp
[AGENCE-DAKAR-AR720-aaa]quit
[AGENCE-DAKAR-AR720]ip pool vpn-nomades
[AGENCE-DAKAR-AR720-ip-pool-vpn-nomades]network 192.168.50.0 mask 255.255.255.0
[AGENCE-DAKAR-AR720-ip-pool-vpn-nomades]gateway-list 192.168.50.1
[AGENCE-DAKAR-AR720-ip-pool-vpn-nomades]quit
[AGENCE-DAKAR-AR720]interface Virtual-Template 1
[AGENCE-DAKAR-AR720-Virtual-Template1]ppp ipcp dns 192.168.10.1
[AGENCE-DAKAR-AR720-Virtual-Template1]remote address pool vpn-nomades
[AGENCE-DAKAR-AR720-Virtual-Template1]quit
```

Côté client (Windows/Android/iOS) : VPN L2TP/IPSec avec PSK + login/mot de passe PPP.
Ports à ouvrir sur le WAN : UDP 500, UDP 4500, UDP 1701 (voir NAT server section 38 +
politique WAN→Local section 66).

## 60. L2TP : vérification et sessions actives

```
display l2tp tunnel
display l2tp session
display ppp user
```

On y voit les tunnels montés, les sessions PPP et les utilisateurs connectés. Pour
déconnecter un utilisateur : `cut connection` (syntaxe exacte selon version — **à
vérifier sur la fiche du modèle exact**) ou désactiver son compte AAA.

## 61. IPSec avec certificats : principe et étapes

Le PSK partagé, c'est bien pour 2–3 sites. Au-delà (ou pour une exigence de sécurité),
on passe aux certificats X.509 :

1. Déployer ou utiliser une CA (interne : AD CS, ou appliance ; jamais d'auto-signé
   bricolé sans procédure).
2. Générer une clé RSA sur l'AR720 (`pki` / `rsa local-key-pair create` — syntaxe selon
   version VRP).
3. Créer une demande de certificat (CSR), la faire signer par la CA.
4. Importer le certificat + la CA sur le routeur.
5. Dans le `ike peer`, remplacer `pre-shared-key` par l'authentification par certificat
   (`certificate` + `rsa-signature` selon version).

**Avertissement version** : les commandes PKI varient sensiblement entre VRP V200 et
V300. Suivre le guide « PKI Configuration » de la version exacte (`display version`).
Points critiques : horloge exacte (NTP obligatoire, sinon le certificat paraît expiré),
et renouvellement suivi (un certificat expiré = tous les tunnels tombent en même temps).

## 62. Dépannage tunnel : la méthode en 5 minutes

Ordre de diagnostic, toujours le même :

1. **Connectivité IP** : `ping` l'IP publique distante. Si ça ne répond pas, le problème
   n'est pas IPSec.
2. **Phase 1** : `display ike sa`. Vide → vérifier UDP 500/4500 ouverts, PSK identique,
   proposals compatibles, `ike-version` identique.
3. **Phase 2** : `display ipsec sa`. Phase 1 OK mais pas de phase 2 → crypto ACL non
   miroir, transform-set incompatible, PFS différent.
4. **Trafic** : compteurs SA à 0 → exclusion NAT oubliée, route manquante vers le
   réseau distant, politique de sécurité qui bloque.
5. **Logs** : `display logbuffer` + debug ciblé (`debugging ike`, `debugging ipsec` —
   avec parcimonie, voir section 100).

## 63. Keepalive et DPD : détecter un pair mort

```
[AGENCE-DAKAR-AR720]ike peer SIEGE
[AGENCE-DAKAR-AR720-ike-peer-SIEGE]dpd type periodic
[AGENCE-DAKAR-AR720-ike-peer-SIEGE]quit
```

Le DPD (Dead Peer Detection) envoie des sondes périodiques : si le siège ne répond plus,
la SA est nettoyée et renégociée au prochain trafic. Sans DPD, un tunnel peut rester
« à moitié monté » indéfiniment après une coupure. Recommandé sur tous les tunnels,
surtout avec des liens 4G instables.

---
---

# Partie G — Sécurité

## 64. Le modèle de zones Huawei : comprendre avant de configurer

VRP utilise des **zones de sécurité** : chaque interface appartient à une zone, et le
trafic est filtré **entre zones** (interzone), pas par interface. Zones par défaut :

| Zone | Usage typique |
|---|---|
| `trust` | LAN interne |
| `untrust` | WAN / Internet |
| `dmz` | Serveurs exposés |
| `local` | Le routeur lui-même |

