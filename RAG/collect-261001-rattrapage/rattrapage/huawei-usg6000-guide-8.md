---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-8
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [1069, 1223]
sha256: 29324c67e0ab596b2025e92a539212b42c448402c3b7932459a9403bd525f96f
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

```huawei
# Côté A :
[USG-A] security-policy
# Trafic LAN local vers LAN distant (via le tunnel) :
[USG-A-policy-security] rule name vpn-a-vers-b
[USG-A-policy-security-rule-vpn-a-vers-b] source-zone trust
[USG-A-policy-security-rule-vpn-a-vers-b] destination-zone untrust
[USG-A-policy-security-rule-vpn-a-vers-b] source-address 192.168.10.0 24
[USG-A-policy-security-rule-vpn-a-vers-b] destination-address 192.168.20.0 24
[USG-A-policy-security-rule-vpn-a-vers-b] action permit
[USG-A-policy-security-rule-vpn-a-vers-b] quit
# Le firewall lui-même doit accepter IKE/ESP (zone local) :
[USG-A-policy-security] rule name vpn-ike-local
[USG-A-policy-security-rule-vpn-ike-local] source-zone untrust
[USG-A-policy-security-rule-vpn-ike-local] destination-zone local
[USG-A-policy-security-rule-vpn-ike-local] source-address 198.51.100.1 32
[USG-A-policy-security-rule-vpn-ike-local] service ike esp nat-t
[USG-A-policy-security-rule-vpn-ike-local] action permit
[USG-A-policy-security-rule-vpn-ike-local] quit
[USG-A-policy-security] quit
# Exemption NAT (section 52) - CRITIQUE :
[USG-A] nat-policy
[USG-A-policy-nat] rule name no-nat-vpn-b
[USG-A-policy-nat-rule-no-nat-vpn-b] source-zone trust
[USG-A-policy-nat-rule-no-nat-vpn-b] destination-zone untrust
[USG-A-policy-nat-rule-no-nat-vpn-b] source-address 192.168.10.0 24
[USG-A-policy-nat-rule-no-nat-vpn-b] destination-address 192.168.20.0 24
[USG-A-policy-nat-rule-no-nat-vpn-b] action no-nat
[USG-A-policy-nat-rule-no-nat-vpn-b] quit
[USG-A-policy-nat] quit
save
```

## 61. IPSec : vérification et diagnostic

```huawei
[USG-A] display ike sa                       # Phase 1 : doit afficher RD/ST (ready)
[USG-A] display ipsec sa                     # Phase 2 : SPI, algorithmes, compteurs
[USG-A] display ike proposal
[USG-A] display ipsec policy
# Debug ciblé (à couper après usage !) :
[USG-A] debugging ike packet                 # échanges IKE en live
[USG-A] debugging ipsec packet               # paquets ESP
[USG-A] undo debugging all                   # TOUJOURS couper après
```

Lecture de `display ike sa` :
- **État RD** (ready) en phase 1 + SA IPSec active = tunnel OK.
- Phase 1 **absente** → problème réseau/PSK/proposal IKE (voir section 147).
- Phase 1 OK mais **pas de phase 2** → ACL ou transform-set incohérents entre les deux côtés.

## 62. IPSec : durcir la config (PFS, DPD, lifetimes)

```huawei
system-view
[USG-A] ike peer peer-site-b
[USG-A-ike-peer-peer-site-b] dpd type periodic   # Dead Peer Detection : détecte un pair mort
[USG-A-ike-peer-peer-site-b] quit
[USG-A] ipsec proposal trans-site-b
[USG-A-ipsec-proposal-trans-site-b] pfs dh-group14   # Perfect Forward Secrecy
[USG-A-ipsec-proposal-trans-site-b] quit
[USG-A] ike sa-lifetime 86400                # durée de vie phase 1 (défaut souvent 86400s)
[USG-A] ipsec sa-lifetime time-based 3600    # durée de vie phase 2
save
```

- **DPD** : indispensable — sans lui, un tunnel « à moitié mort » ne se reconstruit pas tout seul.
- **PFS** : recommandé (compromis perf/sécurité acceptable sur USG6000).
- ⚠️ Les **lifetimes doivent être cohérents** des deux côtés (pas forcément identiques, mais compatibles — le plus petit l'emporte lors de la renégociation).

## 63. SSL VPN : principe et cas d'usage

Le SSL VPN permet aux **utilisateurs nomades** d'accéder au réseau interne via un simple navigateur ou un client léger (SecoClient), en HTTPS (port 443) — ça traverse **tous les NAT et proxys** d'hôtel/aéroport.

Modes d'accès :
- **Web proxy** : accès web aux applis internes via le portail (pas de client).
- **Port forwarding** : redirection de ports TCP (ex. RDP, SSH) via applet.
- **Network extension** : client VPN complet (tunnel réseau, comme un IPSec client).
- **File sharing** : accès aux partages de fichiers via le portail.

Licences : le nombre d'utilisateurs SSL VPN simultanés dépend du modèle **et de la licence** (100 par défaut sur beaucoup de modèles, extensible — à vérifier).

## 64. SSL VPN : configuration complète (portail + network extension)

```huawei
system-view
# ===== 1. Activer le SSL VPN et créer la passerelle =====
[USG] sslvpn gateway sslgw
[USG-sslvpn-gateway-sslgw] ip address 203.0.113.1 port 443
[USG-sslvpn-gateway-sslgw] quit
# ===== 2. Domaine d'authentification (base locale pour commencer) =====
[USG] sslvpn domain domaine-vpn
[USG-sslvpn-domain-domaine-vpn] quit
# ===== 3. Politique d'accès : qui accède à quoi =====
[USG] sslvpn policy group politique-nomades
[USG-sslvpn-policy-group-politique-nomades] resource ip 192.168.10.0 24   # réseau accessible
[USG-sslvpn-policy-group-politique-nomades] quit
# ===== 4. Utilisateur local de test (fictif) =====
[USG] aaa
[USG-aaa] local-user vpn-user1 password
[USG-aaa] local-user vpn-user1 service-type sslvpn
[USG-aaa] local-user vpn-user1 privilege level 3
[USG-aaa] quit
# ===== 5. Lier la passerelle au domaine =====
[USG] sslvpn gateway sslgw
[USG-sslvpn-gateway-sslgw] domain domaine-vpn
[USG-sslvpn-gateway-sslgw] quit
# ===== 6. Politique de sécurité : autoriser le 443 vers local =====
[USG] security-policy
[USG-policy-security] rule name sslvpn-443-in
[USG-policy-security-rule-sslvpn-443-in] source-zone untrust
[USG-policy-security-rule-sslvpn-443-in] destination-zone local
[USG-policy-security-rule-sslvpn-443-in] service https
[USG-policy-security-rule-sslvpn-443-in] action permit
[USG-policy-security] quit
save
```

⚠️ Les commandes SSL VPN **varient selon les versions** (V500 vs V600, la syntaxe `sslvpn gateway` / `virtual-gateway` a évolué) — **valide chaque commande sur ta version** avec `?` avant de l'appliquer en prod. Le web (VPN > SSL VPN) est souvent plus fiable que le CLI pour cette fonction.

## 65. SSL VPN : vérifications côté utilisateur et admin

Côté utilisateur : navigateur → `https://203.0.113.1` → portail → login → accès aux ressources ou lancement du client Network Extension.

Côté admin :
```huawei
[USG] display sslvpn gateway sslgw              # état de la passerelle
[USG] display sslvpn online-user                 # utilisateurs connectés (selon version)
[USG] display sslvpn session                    # sessions actives
```

🔧 « La page du portail ne s'ouvre pas » → checklist section 150. « Connecté mais pas d'accès réseau » → politique de sécurité pour le **trafic issu du tunnel** (les utilisateurs VPN arrivent souvent dans une zone dédiée ou trust — autoriser explicitement).

## 66. SSL VPN : certificats (ne pas rester en auto-signé)

Par défaut, le portail utilise un **certificat auto-signé** → alerte navigateur à chaque connexion, les utilisateurs s'habituent à cliquer « accepter le risque » (mauvaise hygiène).

Procédure :
1. Générer une CSR ou importer un certificat d'une **CA reconnue** (Let's Encrypt via ton infra, ou CA d'entreprise).
2. Web : System > Certificate Management > Import.
3. Lier le certificat à la passerelle SSL VPN.
4. 📋 Renouveler **avant expiration** — un portail VPN avec certificat expiré un lundi 8h, c'est 50 appels au support.

## 67. SSL VPN : bonnes pratiques de sécurité

- **Double authentification** : mot de passe + certificat client ou OTP (selon version/licence).
- **Contrôle de conformité** du poste (antivirus à jour) si la fonction est disponible.
- **Split tunneling** : ne router dans le VPN que le trafic interne (perf + vie privée) — ou **full tunneling** si la politique impose l'inspection de tout le trafic (à décider, pas par défaut).
- **Timeout de session** agressif (15–30 min d'inactivité).
- **Journaliser** les connexions (qui, quand, depuis quelle IP).

## 68. GRE : tunnel simple entre deux USG

GRE = encapsulation sans chiffrement. Utile pour transporter des protocoles non-IP ou du multicast (ex. OSPF entre sites), **souvent protégé par IPSec par-dessus** (GRE over IPSec).

