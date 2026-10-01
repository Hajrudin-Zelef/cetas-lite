---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-6
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [737, 907]
sha256: ca71d5c28868f394481a83501faf914a8b2ec137c03ce71f89cd9a8cadd68c7e
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

Complément de la section 23 — le socle « le firewall parle au SI » :
```huawei
[USG] security-policy
# Le firewall envoie ses logs (local→trust) :
[USG-policy-security] rule name local-syslog-out
[USG-policy-security-rule-local-syslog-out] source-zone local
[USG-policy-security-rule-local-syslog-out] destination-zone trust
[USG-policy-security-rule-local-syslog-out] destination-address 192.168.10.100 32
[USG-policy-security-rule-local-syslog-out] service syslog
[USG-policy-security-rule-local-syslog-out] action permit
[USG-policy-security-rule-local-syslog-out] quit
# Le firewall interroge le NTP (local→untrust) :
[USG-policy-security] rule name local-ntp-out
[USG-policy-security-rule-local-ntp-out] source-zone local
[USG-policy-security-rule-local-ntp-out] destination-zone untrust
[USG-policy-security-rule-local-ntp-out] service ntp
[USG-policy-security-rule-local-ntp-out] action permit
[USG-policy-security-rule-local-ntp-out] quit
[USG-policy-security] quit
save
```

⚠️ Avec `default action deny`, **ces règles sont obligatoires** : sans elles, pas de NTP (→ horloge fausse → logs et VPN en vrac) et pas de syslog (→ aveugle).

## 44. Politique par défaut : le filet, pas la stratégie

```huawei
[USG] security-policy
[USG-policy-security] default action deny   # base globale restrictive
[USG-policy-security] quit
save
```

💡 La politique par défaut n'est qu'un **filet** : des règles explicites pour chaque flux légitime, `deny` pour tout le reste. Certaines versions permettent d'affiner l'action par défaut **par couple de zones** (syntaxe à vérifier avec `?`) — utile pour être strict sur untrust→* tout en restant souple sur local→trust (management).

---

---

# BLOC E — NAT : TRADUCTION D'ADRESSES

## 45. Les 3 types de NAT sur USG (et quand les utiliser)

| Type | Sens | Cas d'usage |
|------|------|-------------|
| **NAT de source** (SNAT) | privé → public | LAN vers Internet (le cas n°1) |
| **NAT de destination** (DNAT / server-map) | public → privé | Publier un serveur (web, mail) |
| **NAT bidirectionnel** | les deux | DMZ avec IP publiques dédiées |

Principe VRP : le NAT se configure dans la vue **`nat-policy`** (pas dans security-policy). Le NAT s'applique **avant** la politique de sécurité pour le trafic entrant, **après** pour le sortant — l'ordre compte pour le debug (section 143).

## 46. Easy IP : le NAT de source simple (90 % des cas)

« Easy IP » = le firewall utilise **l'IP de l'interface de sortie** comme adresse source NATée. Idéal quand le FAI donne une seule IP (ou un petit bloc DHCP/PPPoE).

```huawei
system-view
[USG] nat-policy
[USG-policy-nat] rule name nat-lan-internet
[USG-policy-nat-rule-nat-lan-internet] source-zone trust
[USG-policy-nat-rule-nat-lan-internet] destination-zone untrust
[USG-policy-nat-rule-nat-lan-internet] source-address 192.168.10.0 24
[USG-policy-nat-rule-nat-lan-internet] action nat easy-ip
[USG-policy-nat-rule-nat-lan-internet] quit
[USG-policy-nat] quit
save
```

🔧 Vérifier : `display nat-policy rule all`, `display firewall session table` (voir les adresses traduites), `display nat session`.

## 47. NAT avec pool d'adresses (plusieurs IP publiques)

Quand le FAI fournit un **bloc d'IP publiques** (ex. /29 = 6 IP utilisables) :

```huawei
system-view
[USG] nat address-group pool-fai
[USG-address-group-pool-fai] section 0 203.0.113.10 203.0.113.14   # IP fictives
[USG-address-group-pool-fai] quit
[USG] nat-policy
[USG-policy-nat] rule name nat-lan-pool
[USG-policy-nat-rule-nat-lan-pool] source-zone trust
[USG-policy-nat-rule-nat-lan-pool] destination-zone untrust
[USG-policy-nat-rule-nat-lan-pool] source-address 192.168.10.0 24
[USG-policy-nat-rule-nat-lan-pool] action nat address-group pool-fai
[USG-policy-nat-rule-nat-lan-pool] quit
[USG-policy-nat] quit
save
```

Modes du pool (à vérifier selon version) :
- **PAT** (par défaut) : plusieurs privés partagent les IP du pool avec des ports différents — le plus courant.
- **1:1** : une IP privée = une IP publique (réserver pour les serveurs).
- **No-PAT** : traduction d'adresse sans traduction de port (consomme une IP publique par session).

## 48. NAT de destination : publier un serveur web (port mapping)

Scénario : serveur web interne `172.16.1.10:80`, IP publique `203.0.113.20`, port 80.

```huawei
system-view
[USG] nat-policy
[USG-policy-nat] rule name dnat-web
[USG-policy-nat-rule-dnat-web] source-zone untrust
[USG-policy-nat-rule-dnat-web] destination-zone dmz
[USG-policy-nat-rule-dnat-web] destination-address 203.0.113.20 32
[USG-policy-nat-rule-dnat-web] service http
[USG-policy-nat-rule-dnat-web] action nat static
[USG-policy-nat-rule-dnat-web] static-to 172.16.1.10
[USG-policy-nat-rule-dnat-web] quit
[USG-policy-nat] quit
# Politique de sécurité associée (OBLIGATOIRE) :
[USG] security-policy
[USG-policy-security] rule name untrust-web-dmz
[USG-policy-security-rule-untrust-web-dmz] source-zone untrust
[USG-policy-security-rule-untrust-web-dmz] destination-zone dmz
[USG-policy-security-rule-untrust-web-dmz] destination-address 172.16.1.10 32
[USG-policy-security-rule-untrust-web-dmz] service http https
[USG-policy-security-rule-untrust-web-dmz] action permit
[USG-policy-security-rule-untrust-web-dmz] quit
[USG-policy-security] quit
save
```

⚠️ **Erreur classique** : configurer le NAT mais oublier la **politique de sécurité** (ou la mettre avec la mauvaise adresse destination — après DNAT, la politique voit **l'IP privée**). Si « le NAT est configuré mais ça ne passe pas », c'est 90 % du temps la politique.

## 49. Port mapping avancé : un port public vers un port privé différent

```huawei
[USG] nat-policy
[USG-policy-nat] rule name dnat-rdp
[USG-policy-nat-rule-dnat-rdp] source-zone untrust
[USG-policy-nat-rule-dnat-rdp] destination-zone trust
[USG-policy-nat-rule-dnat-rdp] destination-address 203.0.113.20 32
[USG-policy-nat-rule-dnat-rdp] service 0 protocol tcp destination-port 13389
[USG-policy-nat-rule-dnat-rdp] action nat static
[USG-policy-nat-rule-dnat-rdp] static-to 192.168.10.50 port 3389
[USG-policy-nat-rule-dnat-rdp] quit
[USG-policy-nat] quit
```

⚠️ **Sécurité** : exposer du RDP/SSH sur Internet = cible n°1 des bots. Si vraiment nécessaire : port non standard + **restriction par IP source** + authentification forte + envisager le **SSL VPN** à la place (section 63).

## 50. NAT bidirectionnel (Two-way NAT)

Pour une DMZ dont les serveurs doivent **sortir avec leur propre IP publique** ET être **joints depuis Internet sur cette même IP** :

```huawei
system-view
[USG] nat-policy
[USG-policy-nat] rule name nat-bidir-serveur1
[USG-policy-nat-rule-nat-bidir-serveur1] source-zone dmz
[USG-policy-nat-rule-nat-bidir-serveur1] destination-zone untrust
[USG-policy-nat-rule-nat-bidir-serveur1] source-address 172.16.1.10 32
[USG-policy-nat-rule-nat-bidir-serveur1] action nat static bidirectional
[USG-policy-nat-rule-nat-bidir-serveur1] static-to 203.0.113.21
[USG-policy-nat-rule-nat-bidir-serveur1] quit
[USG-policy-nat] quit
save
```

Le `bidirectional` crée automatiquement la traduction inverse. Vérifie avec `display nat server-map`.

## 51. NAT et hairpin : accéder au serveur par son IP publique depuis le LAN

Problème : depuis le LAN, `http://203.0.113.20` ne répond pas (le paquet sort en NAT puis revient — l'USG le droppe souvent).

Solutions :
1. **Split-DNS** : en interne, le nom DNS résout vers `172.16.1.10` (recommandé, le plus propre).
2. **NAT hairpin** (si pas de DNS interne) : règle NAT trust→dmz qui traduit la destination :

