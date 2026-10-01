---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-7
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [908, 1068]
sha256: 3a84f5bf8c07dd245df6e1b7ecd7aadbb9ac452200cc262c7ec1bbc4bb1925f9
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

```huawei
[USG] nat-policy
[USG-policy-nat] rule name hairpin-web
[USG-policy-nat-rule-hairpin-web] source-zone trust
[USG-policy-nat-rule-hairpin-web] destination-zone dmz
[USG-policy-nat-rule-hairpin-web] destination-address 203.0.113.20 32
[USG-policy-nat-rule-hairpin-web] service http
[USG-policy-nat-rule-hairpin-web] action nat static
[USG-policy-nat-rule-hairpin-web] static-to 172.16.1.10
[USG-policy-nat-rule-hairpin-web] quit
[USG-policy-nat] quit
# + politique trust→dmz correspondante (même logique que la règle untrust→dmz, section 48)
save
```

## 52. NAT pour le VPN : exemption NAT (nat no-nat)

⚠️ **Le trafic VPN ne doit PAS être NATé.** Quand une règle NAT easy-ip couvre aussi le trafic vers le tunnel, le VPN casse. Solution : règle d'exemption **avant** la règle NAT :

```huawei
[USG] nat-policy
[USG-policy-nat] rule name no-nat-vpn
[USG-policy-nat-rule-no-nat-vpn] source-zone trust
[USG-policy-nat-rule-no-nat-vpn] destination-zone untrust
[USG-policy-nat-rule-no-nat-vpn] source-address 192.168.10.0 24
[USG-policy-nat-rule-no-nat-vpn] destination-address 192.168.20.0 24   # réseau distant du VPN
[USG-policy-nat-rule-no-nat-vpn] action no-nat
[USG-policy-nat-rule-no-nat-vpn] quit
[USG-policy-nat] move rule no-nat-vpn before nat-lan-internet
[USG-policy-nat] quit
save
```

**Ordre : exemption d'abord, NAT ensuite.** C'est l'erreur n°1 des VPN qui « montent mais ne passent pas de trafic » (section 147).

## 53. ALG : quand le NAT casse les protocoles (FTP, SIP, PPTP)

FTP/SIP/H.323 embarquent des IP **dans la charge utile** — le NAT d'en-tête ne suffit pas. L'USG a des **ALG** (Application Level Gateways) :

```huawei
system-view
[USG] firewall alg ftp enable
[USG] firewall alg sip enable
[USG] firewall alg pptp enable
# Voir l'état :
[USG] display firewall alg
save
```

🔧 FTP qui liste mais ne télécharge pas en mode actif/passif → 90 % ALG ou politique sur les ports de données. SIP qui s'enregistre mais pas d'audio → ALG SIP + ouvrir RTP (UDP 10000-20000 typique, à vérifier).

## 54. Diagnostiquer le NAT : les commandes qui tranchent

```huawei
[USG] display nat-policy rule all                    # règles NAT dans l'ordre
[USG] display nat-policy statistics rule nat-lan-internet   # compteurs par règle
[USG] display firewall session table source inside 192.168.10.25  # voir la traduction live
[USG] display nat address-group pool-fai             # état du pool (adresses utilisées)
[USG] display nat server-map                         # mappings statiques (DNAT)
[USG] display firewall session table verbose         # détail : NAT avant/après
```

📋 Checklist « le NAT ne marche pas » :
- [ ] La règle NAT est-elle **avant** toute règle qui matcherait aussi ?
- [ ] Les **zones** source/destination de la règle NAT sont-elles correctes ?
- [ ] La **politique de sécurité** autorise-t-elle le flux (avec les bonnes adresses post-NAT) ?
- [ ] L'interface de sortie a-t-elle une **route par défaut** ?
- [ ] Pas de **no-nat** qui court-circuite par erreur ?

## 55. NAT64 / IPv6 : le mot rapide

L'USG6000 supporte le NAT IPv6/IPv4 (NAT64, IVI) pour la transition. En pratique PME/industrie en 2026 : **rarement utilisé** — la plupart des déploiements restent en IPv4 + NAT44. Si ton FAI te pousse de l'IPv6 natif, configure le firewall en **dual-stack** avec des politiques IPv6 dédiées plutôt qu'un NAT64 complexe, sauf besoin d'interconnexion spécifique.

## 56. Web : configurer le NAT en interface graphique

Chemin (New Web UI) : **Policy > NAT Policy**. Assistant de création : choisir le type (Source NAT / Destination NAT / Bidirectional), les zones, les objets d'adresse, le mode (Easy IP / Address Pool / Static).
💡 Le web est pratique pour **visualiser l'ordre des règles** (glisser-déposer pour réordonner) — en CLI, `move rule` fait pareil.

---
---

# BLOC F — VPN

## 57. Vue d'ensemble : quel VPN pour quel besoin

| Besoin | Technologie | Section |
|--------|-------------|---------|
| Relier deux sites fixes | **IPSec site-à-site** | 58–62 |
| Nomades / télétravail | **SSL VPN** | 63–67 |
| Tunnel simple sans chiffrement (ou sous IPSec) | **GRE** | 68–69 |
| Vieux clients Windows natifs | **L2TP over IPSec** | 70 |
| Tout le trafic chiffré + routage dynamique | **GRE over IPSec / IPSec + OSPF** | 71 |

Règle : **IPSec** pour le site-à-site (robuste, standard), **SSL VPN** pour les utilisateurs (simple, traverse les NAT/proxys).

## 58. IPSec site-à-site : architecture et prérequis

```
Site A (USG-A)                                    Site B (USG-B)
LAN 192.168.10.0/24                               LAN 192.168.20.0/24
  │                                                 │
  │ GE1/0/1 : 203.0.113.1 (publique)                │ GE1/0/1 : 198.51.100.1 (publique)
  └────────────── tunnel IPSec ─────────────────────┘
```

📋 Prérequis AVANT de configurer :
- [ ] Les deux USG se **pingent** sur leurs IP publiques (politique local : autoriser ICMP/UDP 500/4500).
- [ ] **NTP synchronisé** des deux côtés (les certificats/PSK y sont sensibles).
- [ ] Les plages LAN des deux sites **ne se chevauchent pas**.
- [ ] **Exemption NAT** prévue (section 52).
- [ ] Ports : **UDP 500** (IKE), **UDP 4500** (NAT-T), **protocole 50** (ESP) ouverts vers la zone local.

## 59. IPSec site-à-site : configuration complète (côté A)

```huawei
system-view
# ===== 1. Proposition IKE (phase 1) =====
[USG-A] ike proposal 10
[USG-A-ike-proposal-10] encryption-algorithm aes-256
[USG-A-ike-proposal-10] authentication-algorithm sha2-256
[USG-A-ike-proposal-10] authentication-method pre-share
[USG-A-ike-proposal-10] dh group14
[USG-A-ike-proposal-10] quit
# ===== 2. Peer IKE : le correspondant =====
[USG-A] ike peer peer-site-b
[USG-A-ike-peer-peer-site-b] exchange-mode main
[USG-A-ike-peer-peer-site-b] pre-shared-key TestOnly123!    # MOT DE PASSE FICTIF - à changer
[USG-A-ike-peer-peer-site-b] ike-proposal 10
[USG-A-ike-peer-peer-site-b] remote-address 198.51.100.1
[USG-A-ike-peer-peer-site-b] local-address 203.0.113.1
[USG-A-ike-peer-peer-site-b] quit
# ===== 3. Proposition IPSec (phase 2) : transform-set =====
[USG-A] ipsec proposal trans-site-b
[USG-A-ipsec-proposal-trans-site-b] encapsulation-mode tunnel
[USG-A-ipsec-proposal-trans-site-b] transform esp
[USG-A-ipsec-proposal-trans-site-b] esp encryption-algorithm aes-256
[USG-A-ipsec-proposal-trans-site-b] esp authentication-algorithm sha2-256
[USG-A-ipsec-proposal-trans-site-b] quit
# ===== 4. ACL : quel trafic va dans le tunnel =====
[USG-A] acl 3001
[USG-A-acl-adv-3001] rule permit ip source 192.168.10.0 0.0.0.255 destination 192.168.20.0 0.0.0.255
[USG-A-acl-adv-3001] quit
# ===== 5. Politique IPSec : on assemble =====
[USG-A] ipsec policy policy-site-b 10 isakmp
[USG-A-ipsec-policy-isakmp-policy-site-b-10] security acl 3001
[USG-A-ipsec-policy-isakmp-policy-site-b-10] ike-peer peer-site-b
[USG-A-ipsec-policy-isakmp-policy-site-b-10] proposal trans-site-b
[USG-A-ipsec-policy-isakmp-policy-site-b-10] quit
# ===== 6. Application sur l'interface WAN =====
[USG-A] interface GigabitEthernet 1/0/1
[USG-A-GigabitEthernet1/0/1] ipsec policy policy-site-b
[USG-A-GigabitEthernet1/0/1] quit
save
```

Côté B : **miroir** (inverser les IP et les réseaux dans l'ACL). La PSK doit être **identique** des deux côtés.

## 60. IPSec : politiques de sécurité pour le tunnel

Le trafic VPN traverse les zones — il faut l'autoriser :

