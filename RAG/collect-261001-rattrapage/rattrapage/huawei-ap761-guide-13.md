---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-13
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [1243, 1363]
sha256: 95c84f8a724c4d6a786ebe24706fa21f892a7429b374c7fe37ef22cfe4cc2ffc
---

# Guide ultra-complet — Huawei eKit AP761

**Configuration principe côté AP (Fat) :**
```
[AP761] radius-server template RAD-SRV
[AP761-radius-RAD-SRV] radius-server authentication 192.168.10.50 1812
[AP761-radius-RAD-SRV] radius-server accounting 192.168.10.50 1813
[AP761-radius-RAD-SRV] radius-server shared-key cipher CleRadiusForte_2026!
[AP761-radius-RAD-SRV] quit
[AP761-wlan-view] security-profile name SEC-ENT
[AP761-wlan-sec-prof-SEC-ENT] security wpa2 dot1x aes
# (attacher le template RADIUS au profil AAA — syntaxe à valider selon version)
```

**Choix EAP :**
| Méthode | Sécurité | Contrainte |
|---|---|---|
| PEAP-MSCHAPv2 | Bonne (tunnel TLS + login/mdp) | Mots de passe à gérer ; attaquable si le client ne vérifie pas le certificat serveur |
| EAP-TLS | Excellente (certificats mutuels) | PKI à déployer (mais tu as un guide CA interne — voir ton guide Debian) |
| EAP-TTLS | Bonne | Moins supportée côté clients |

**Point critique : le certificat du serveur RADIUS.** Les clients doivent **vérifier** le certificat du serveur (sinon : attaque par faux RADIUS / evil twin + credential harvesting). Déployer le certificat via GPO/MDM, et **former les utilisateurs à ne jamais accepter un certificat inconnu**.

## 49. Portail captif (Portal) pour invités

Le portail captif : le visiteur se connecte à `INVITE` (souvent en ouvert ou PSK simple), ouvre son navigateur, et tombe sur une **page d'authentification** (CGU à accepter, code reçu à l'accueil, identifiants temporaires).

**Ce que le portail fait :** de l'**authentification et du traçage** (qui était connecté quand — obligation légale en France pour un accès public : conservation des données de connexion).
**Ce que le portail ne fait PAS :** du **chiffrement**. Entre le client et l'AP, en portail ouvert, le trafic est en clair. D'où les mesures complémentaires obligatoires :

| Mesure | Pourquoi |
|---|---|
| **Isolation client** (chap. 53) | Un invité ne doit pas voir le PC de l'autre invité |
| **VLAN dédié** avec ACL « Internet seul » (chap. 52) | Un invité ne doit pas toucher le réseau interne |
| **HTTPS forcé côté services** | Les utilisateurs doivent privilégier les sites en HTTPS |
| **Bande passante limitée** par client | Un invité ne doit pas saturer le lien (QoS, chap. 54) |

**Problèmes classiques (cas 12) :** la page du portail ne s'affiche pas parce que le client a un **DNS personnalisé**, un **VPN actif**, ou un navigateur qui bloque la détection (CNA). Prévoir une **page d'aide** affichée à l'accueil avec l'URL directe du portail et la procédure.

**Obligation légale (France) :** un accès Wi-Fi public doit permettre l'identification des utilisateurs (LCEN) — le portail avec enregistrement (email, SMS, ticket) répond à ce besoin. **Conserver les journaux** selon les délais légaux — à valider avec ton service juridique.

## 50. Filtrage MAC : utile ou placebo ?

Le filtrage MAC (whitelist des adresses MAC autorisées) est supporté par l'AP761 [constructeur]. Verdict honnête : **c'est un placebo de sécurité**, pas une protection.

**Pourquoi ça ne protège pas :**
- Une adresse MAC se **capture en clair** dans l'air (même en WPA2) et se **spoofe en 10 secondes** sur n'importe quel OS.
- La gestion d'une whitelist à la main est un enfer administratif (chaque nouveau téléphone = un ticket).

**Où ça reste utile :**
- **SSID IoT** : limiter les associations à la liste des caméras/capteurs connus — pas comme sécurité, mais comme **garde-fou** (éviter qu'un visiteur ne se connecte au SSID IoT par erreur) et comme **inventaire**.
- **Détection** : une alerte « MAC inconnue sur SSID IoT » est un signal d'investigation, pas une barrière.

**Ne jamais** présenter le filtrage MAC comme une mesure de sécurité dans un dossier ou un audit : un auditeur sérieux le recalera.

## 51. 802.11w (PMF) : protection des trames de management

**PMF (Protected Management Frames, 802.11w)** [constructeur] : chiffre/signe les trames de management (désauthentification, désassociation). Sans PMF, un attaquant peut **éjecter n'importe quel client** en forgeant une trame de désauthentification — c'est l'attaque la plus simple du Wi-Fi.

| Mode PMF | Effet |
|---|---|
| Désactivé | Vulnérable aux attaques de désauthentification |
| **Optionnel** (recommandé en transition) | Les clients compatibles en profitent, les autres se connectent quand même |
| Obligatoire | Sécurité max, mais **exclut les clients non-PMF** (beaucoup d'IoT) |

**Recommandation :**
- SSID salariés (parc moderne) : **PMF obligatoire** si le parc le supporte, sinon optionnel.
- SSID IoT : **désactivé ou optionnel** (les caméras/capteurs sont souvent non-PMF).
- **WPA3 impose PMF** : en WPA3-only ou transition, le PMF est de fait requis pour les clients SAE.

**En CLI (principe) :**
```
[AP761-wlan-sec-prof-SEC-CORP] pmf optional
# ou : pmf mandatory / pmf disable
```

## 52. VLAN : segmentation par SSID

**Supporté** : 802.1Q, **VLAN par SSID** [constructeur]. C'est la base de la segmentation : chaque SSID = un VLAN = un sous-réseau = une politique.

**Schéma type :**
```
SSID ENTREPRISE ──► VLAN 20 ──► 192.168.20.0/24 ──► accès LAN + Internet
SSID INVITE     ──► VLAN 30 ──► 192.168.30.0/24 ──► Internet seul (ACL)
SSID IOT        ──► VLAN 40 ──► 192.168.40.0/24 ──► serveurs IoT seuls (ACL)
Management      ──► VLAN 10 ──► 192.168.10.0/24 ──► AP, switch, supervision
```

**Configuration :** le VAP attache le VLAN (`service-vlan vlan-id 20` — chap. 34), et le port de l'AP côté switch est en **trunk** avec le VLAN 10 en natif (PVID) :

```
# Côté switch (exemple générique) :
# port de l'AP : trunk, PVID 10, allowed 10, 20, 30, 40
```

**Règles de firewall inter-VLAN (à mettre sur le routeur/pare-feu) :**
| Flux | Règle |
|---|---|
| INVITE → LAN | ❌ Refusé |
| INVITE → Internet | ✅ Autorisé |
| INVITE → INVITE (inter-clients) | ❌ Refusé (isolation, chap. 53) |
| IOT → Internet | ❌ Refusé (sauf NTP/DNS si besoin) |
| IOT → serveur vidéosurveillance | ✅ Autorisé (ports caméra uniquement) |
| ENTREPRISE → tout | ✅ (selon politique interne) |
| Tout → Management (VLAN 10) | ❌ Refusé sauf postes d'admin |

**Erreur classique :** mettre le management de l'AP dans le même VLAN que les invités « pour simplifier ». Non : le management est la porte d'entrée vers toute l'infrastructure — **VLAN dédié, accès restreint**.

## 53. Isolation client (STA isolation) dans un VLAN

**STA isolation** (isolation intra-VLAN) [constructeur] : les clients d'un même SSID/VLAN **ne peuvent pas communiquer entre eux** — tout passe par la passerelle. Supporté par l'AP761.

**Où l'activer :**
- ✅ **SSID INVITE : toujours.** Deux visiteurs ne doivent jamais se voir (attaques MITM, partage de fichiers malveillant).
- ✅ **SSID IOT : souvent.** Une caméra n'a aucune raison de parler à une autre caméra.
- ❌ **SSID ENTREPRISE : généralement non.** Le partage d'imprimante, le Chromecast de la salle de réunion, les applis pair-à-pair en ont besoin. À activer seulement si la politique l'exige (et prévoir les exceptions).

**Limite :** l'isolation se fait **au niveau de l'AP**. Si deux clients sont sur deux AP différents du même VLAN sans isolation amont (sur le switch), ils se voient quand même. Pour une isolation totale, combiner avec du **Private VLAN** ou des ACL sur le switch — à vérifier sur la fiche du modèle exact.

## 54. QoS / WMM : prioriser la voix et la vidéo

**WMM (Wi-Fi Multimedia)** [constructeur] : le Wi-Fi reprend les 4 classes de priorité 802.11e — voix, vidéo, best effort, background. L'AP761 fait le **mapping et l'ordonnancement par priorité** [constructeur].

