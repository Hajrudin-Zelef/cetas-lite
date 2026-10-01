---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-9
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "zero-day"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [1224, 1403]
sha256: 9ea3b335708f5b51867e7b52f92a641876a955e16d905920f2dda0a27104514e
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

```huawei
system-view
# Côté A :
[USG-A] interface Tunnel 0
[USG-A-Tunnel0] ip address 10.255.0.1 30
[USG-A-Tunnel0] tunnel-protocol gre
[USG-A-Tunnel0] source 203.0.113.1
[USG-A-Tunnel0] destination 198.51.100.1
[USG-A-Tunnel0] quit
[USG-A] firewall zone untrust
[USG-A-zone-untrust] add interface Tunnel 0
[USG-A-zone-untrust] quit
# Route vers le LAN distant via le tunnel :
[USG-A] ip route-static 192.168.20.0 24 Tunnel 0
save
```

⚠️ GRE seul = **trafic en clair** sur Internet. En prod, toujours GRE **over IPSec** (appliquer une `ipsec policy` sur l'interface Tunnel ou utiliser un profil IPSec).

## 69. GRE over IPSec : le combo routage + chiffrement

```huawei
# Sur l'interface Tunnel existante, ajouter la protection IPSec :
[USG-A] interface Tunnel 0
[USG-A-Tunnel0] ipsec policy policy-gre
[USG-A-Tunnel0] quit
# Avec une ACL qui matche le trafic GRE (protocole 47) :
[USG-A] acl 3002
[USG-A-acl-adv-3002] rule permit gre source 203.0.113.1 0 destination 198.51.100.1 0
[USG-A-acl-adv-3002] quit
```

Avantage : tu peux faire passer de l'**OSPF** dans le tunnel → routage dynamique entre sites, bascule automatique. Voir section 71.

## 70. L2TP over IPSec : pour les clients Windows natifs

```huawei
system-view
# Serveur L2TP :
[USG] l2tp enable
[USG] l2tp-group 1
[USG-l2tp-1] tunnel password Fictif123!          # mot de passe tunnel (fictif)
[USG-l2tp-1] allow l2tp virtual-template 1 remote client1
[USG-l2tp-1] quit
[USG] interface Virtual-Template 1
[USG-Virtual-Template1] ip address 192.168.99.1 24
[USG-Virtual-Template1] ppp authentication-mode chap
[USG-Virtual-Template1] quit
# Pool d'adresses pour les clients :
[USG] ip pool pool-l2tp 192.168.99.10 192.168.99.100
# Utilisateur :
[USG] aaa
[USG-aaa] local-user client1 password
[USG-aaa] local-user client1 service-type ppp
[USG-aaa] quit
save
```

Ensuite, protéger avec IPSec (politique IPSec sur l'interface WAN avec ACL UDP 1701). Côté Windows : créer une connexion VPN L2TP/IPSec avec clé pré-partagée. ⚠️ L2TP seul (sans IPSec) = **mots de passe en clair** — jamais en prod.

## 71. Routage dynamique dans le VPN (OSPF sur GRE/IPSec)

```huawei
# Côté A (après avoir monté le Tunnel 0, section 68) :
[USG-A] ospf 1
[USG-A-ospf-1] area 0
[USG-A-ospf-1-area-0.0.0.0] network 192.168.10.0 0.0.0.255
[USG-A-ospf-1-area-0.0.0.0] network 10.255.0.0 0.0.0.3
[USG-A-ospf-1-area-0.0.0.0] quit
[USG-A-ospf-1] quit
# Sécuriser OSPF (authentification MD5) :
[USG-A] interface Tunnel 0
[USG-A-Tunnel0] ospf authentication-mode md5 1 FictifOspf123
[USG-A-Tunnel0] quit
save
```

Vérif : `display ospf peer`, `display ip routing-table`. Si le tunnel tombe, OSPF reroute via la route de secours (si tu en as une).

## 72. VPN : tableau de diagnostic express

| Symptôme | `display ike sa` | Cause probable | Section |
|----------|------------------|----------------|---------|
| Rien ne monte | vide | Réseau / politique local / NAT | 147 |
| Phase 1 absente | vide | PSK / proposal IKE / ports | 147 |
| Phase 1 OK, pas de phase 2 | RD mais pas d'IPSec SA | ACL / transform-set | 147 |
| Tunnel OK, pas de trafic | SA actives | Politique / NAT / routage | 147 |
| Tunnel instable | SA qui flappent | DPD / lifetimes / double NAT | 148 |

---
---

# BLOC G — UTM : INSPECTION DE CONTENU (AV, IPS, URL, APPLICATIF, ANTI-SPAM)

## 73. L'UTM et la performance : la vérité honnête

**C'est LA section à lire avant d'activer quoi que ce soit.**

- Le **débit firewall** (section 2) mesure du simple filtrage stateful sur des paquets UDP calibrés.
- Dès que tu actives **AV + IPS + URL filtering**, le firewall doit **reassembler les flux TCP, décoder les protocoles (HTTP, SMTP, FTP...), matcher des dizaines de milliers de signatures** — c'est 3 à 10× plus coûteux en CPU.
- Ordres de grandeur constatés (à vérifier sur la fiche du modèle exact, conditions réelles) :
  - **Firewall seul** : 100 % du débit nominal.
  - **+ IPS** : ~60–70 % du débit nominal.
  - **+ IPS + AV** : ~40–60 % (le constructeur annonce jusqu'à ~15 Gbit/s IPS+AV sur les gros modèles, en conditions de test).
  - **+ IPS + AV + URL + SSL inspection** : ~25–40 %.

**Règles de décision :**
1. **Dimensionne sur le débit UTM**, pas sur le débit firewall. Si ton lien fait 1 Gbit/s et que tu veux IPS+AV partout, prends un modèle dont le débit IPS+AV ≥ 1 Gbit/s avec de la marge.
2. **N'active pas tout, partout** : applique les profils UTM uniquement sur les règles qui en ont besoin (LAN→Internet oui ; backup→NAS non).
3. **L'inspection SSL/TLS** est la plus coûteuse (débit dédié souvent 2–5× inférieur au débit IPS). Ne l'active que sur les flux à risque, avec consentement/cadre légal (inspection = lecture du contenu chiffré).
4. 🔧 Surveille : `display cpu`, `display memory`, et les compteurs de **sessions dropped par manque de ressources**.

## 74. Architecture UTM sur USG : profils + politiques

Logique en 2 temps :
1. **Créer un profil** (ex. `av-profile-lan`) : quoi détecter, quelle action.
2. **Attacher le profil à une règle de politique** (section 35) : sur quel flux l'appliquer.

```
Politique LAN→Internet (permit)
   ├── profile AV      → scanne les fichiers téléchargés
   ├── profile IPS     → bloque les exploits connus
   ├── profile URL     → filtre les catégories de sites
   └── profile APP     → contrôle les applications (P2P, réseaux sociaux...)
```

Avantage : **un même profil réutilisable** sur plusieurs règles ; modification centralisée.

## 75. Antivirus : configuration du profil

```huawei
system-view
[USG] profile type av name av-profile-lan
[USG-profile-av-av-profile-lan] exception application any   # exceptions éventuelles
# Action en cas de virus détecté : alert / block (recommandé : block)
[USG-profile-av-av-profile-lan] quit
# Attacher à la règle (déjà vu section 35) :
[USG] security-policy
[USG-policy-security] rule name LAN-web-out
[USG-policy-security-rule-LAN-web-out] profile av-profile-lan
[USG-policy-security-rule-LAN-web-out] quit
[USG-policy-security] quit
save
```

Protocoles scannés typiques : HTTP, FTP, SMTP, POP3, IMAP (à vérifier selon version/licence).
⚠️ **Limites** : l'AV réseau ne remplace pas l'antivirus des postes (il ne voit pas le chiffré sans inspection SSL, ni les menaces zero-day). C'est une **couche**, pas une solution.

## 76. Mise à jour des signatures AV/IPS (vital)

Des signatures vieilles de 6 mois = une passoire qui a l'air de protéger.

```huawei
[USG] display av update info          # version des signatures AV (selon version)
[USG] display ips signature version   # version des signatures IPS
```

- **Automatique** : planifier la mise à jour quotidienne (web : System > Update Center, ou CLI selon version). Nécessite un accès Internet (ou un serveur de mise à jour local).
- **Manuelle** : téléchargement sur support.huawei.com → import via web ou clé USB.
- 📋 Checklist : [ ] MAJ auto activée [ ] dernière MAJ < 7 jours [ ] alerte si échec de MAJ (syslog).
- 🔧 Après une MAJ de signatures, **surveille les faux positifs** (une signature trop agressive peut bloquer une appli métier — voir section 153).

## 77. IPS : choisir le bon niveau de protection

L'IPS détecte les **exploits** (tentatives d'intrusion, scans, shellcodes), pas les virus.

```huawei
system-view
[USG] profile type ips name ips-profile-lan
# Assigner un niveau : les profils prédéfinis vont de "léger" à "strict"
# (noms exacts à vérifier selon version : p.ex. default, strict)
[USG-profile-ips-ips-profile-lan] quit
[USG] security-policy
[USG-policy-security] rule name LAN-web-out
[USG-policy-security-rule-LAN-web-out] profile ips-profile-lan
[USG-policy-security-rule-LAN-web-out] quit
[USG-policy-security] quit
save
```

