---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-1
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["distribution", "arr"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1, 85]
sha256: d376fe64de240710338eb0a03fe32ee91effffafae282ebdd1d7676b50ef1a82
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

> **Public :** Zelef, chef de service systèmes & énergies — technicien de terrain, responsable d'équipe.
> **Ton :** concret, direct, sans jargon inutile. Ce guide est fait pour être utilisé sur site, pas pour être beau sur une étagère.
> **Périmètre vérifié par recherche web (septembre 2026) :** la plateforme Huawei eKit, lancée en 2023 pour le marché de la distribution mondiale, couvre les PME via les gammes **eKitEngine** (réseau : passerelles AR, points d'accès Wi-Fi, switchs, pare-feux USG), **eKitOptix** (MiniFTTO fibre), **eKitStor** (stockage) et **IdeaHub** (collaboration). Ce guide se concentre sur le **datacom** : AR, AP, switchs, USG6000F-S, app eKit et cloud SNC.
> **Règle d'honnêteté de ce guide :** tout ce qui vient des fiches techniques constructeur est sourcé. Quand une info dépend de la version logicielle ou n'a pas pu être vérifiée, c'est écrit noir sur blanc : **« à vérifier sur la documentation officielle »**. Aucune procédure n'est inventée.

---

## SOMMAIRE

- **Sections 1 à 15** — L'écosystème eKit : positionnement, gammes, modèles réels
- **Sections 16 à 30** — L'app eKit : installation, compte, onboarding, organisation par sites
- **Sections 31 à 45** — Le cloud eKit (SME Network Center) : comptes, rôles, multi-sites, limites honnêtes
- **Sections 46 à 60** — Scénarios de déploiement types complets (bureau 20 postes, hôtel, boutique, entrepôt)
- **Sections 61 à 75** — Interopérabilité, migration, bonnes pratiques (adressage, VLAN, SSID, PoE, ordre de mise en service)
- **Sections 76 à 90** — Supervision, limites vs vrai NMS, quand passer au Datacom entreprise
- **Sections 91 à 100** — Sauvegarde, mises à jour, dépannage (15+ cas terrain)
- **Sections 101 à 115** — Maintenance, sécurité, pense-bête de poche, glossaire, quiz, pour aller plus loin

---

## 1. eKit, c'est quoi exactement ?

Huawei eKit est la marque **distribution** de Huawei pour les PME, lancée en 2023. Traduction terrain : ce n'est pas la gamme « opérateur » ni la gamme « grande entreprise » (celle-là, c'est Datacom avec iMaster NCE, les switchs S600 et compagnie). eKit, c'est la gamme que ton **distributeur local** te vend, avec :

- des produits **pré-packagés par scénario** (bureau, hôtel, école, commerce) ;
- une gestion pensée pour des gens qui ne sont **pas des experts réseau** : l'app mobile eKit et le cloud SME Network Center (SNC) ;
- une promesse constructeur : **management app/cloud gratuit, sans licence** (c'est ce qu'annoncent Huawei et ses distributeurs — voir section 33 pour les nuances).

En chiffres (annoncés par Huawei au MWC 2025) : plus de **70 pays**, plus de **400 partenaires Gold distribution**, croissance de l'activité de l'ordre de 30 %. L'Afrique (Afrique du Sud en tête) fait partie des marchés les plus dynamiques. Moralité : tu trouveras du stock et du support, ce n'est pas une gamme fantôme.

## 2. Positionnement : pour qui, pour quoi — et surtout PAS pour quoi

**C'est fait pour :**
- PME de 5 à ~200 utilisateurs ;
- sites uniques ou multi-sites simples (boutiques, agences, hôtels, écoles, cliniques) ;
- équipes IT réduites ou prestataires qui déploient vite ;
- budgets serrés mais besoin de Wi-Fi 6/7 propre et de PoE.

**Ce n'est PAS fait pour :**
- les campus de plusieurs milliers d'utilisateurs avec roaming fin et 802.1X complexe ;
- les datacenters (pas de 25/40/100G sérieux dans la gamme eKit — le S620 monte en 25G/100G mais reste positionné PME) ;
- les environnements qui exigent du routage dynamique avancé (OSPF/BGP multi-zones), du MPLS, du VXLAN ;
- les sites où la **redondance matérielle** (double alim, stacking, VRRP/HSRP) est contractuelle.

Retiens la phrase qui te sauvera en réunion : *« eKit, c'est du solide pour du simple. Dès que le besoin devient complexe, on bascule sur du Datacom entreprise. »* (Critères de bascule détaillés en sections 83-90.)

## 3. Les quatre piliers de l'écosystème eKit

| Pilier | Nom commercial | Contenu | Ce guide couvre |
|---|---|---|---|
| Datacom | **eKitEngine** | Passerelles AR, AP Wi-Fi, switchs, pare-feux USG | Oui, en détail |
| Fibre bureau | **eKitOptix** (MiniFTTO) | F700D (AP Wi-Fi 7 optique 3-en-1), FG736 | Oui, survol + cas hôtel |
| Stockage | **eKitStor** | SSD Xtreme 200E, Shield 210 portable | Non (hors périmètre réseau) |
| Collaboration | **IdeaHub** | S3, B3 (écrans interactifs 4K) | Non (mention uniquement) |

Le datacom eKitEngine est le cœur de ton métier de déploiement. Les trois autres piliers existent, tu dois savoir qu'ils existent, mais ton réseau s'arrête au datacom.

## 4. Les passerelles eKitEngine AR : AR180, AR180 Plus, AR180 Pro, AR280, AR281

La passerelle AR eKit, c'est la « box pro » : **routage + switch + Wi-Fi + gestion Internet** dans un seul boîtier compact. D'après les fiches techniques constructeur (datasheets AR180/AR280, 2025) :

| Modèle | Ports | Wi-Fi | Capacité annoncée | PoE | Refroidissement |
|---|---|---|---|---|---|
| **AR180** | 4× GE LAN + 1× 2.5GE WAN | Wi-Fi 7 double bande (2,4/5 GHz, 2×2) | ~100 terminaux, 8 tunnels IPsec | Non | Passif (sans ventilateur) |
| **AR180 Plus** | 4× GE LAN + 1× 2.5GE WAN | Wi-Fi 7 double bande | ~100 terminaux, 8 tunnels IPsec | Non | Passif |
| **AR180 Pro** | 5× GE (configurables) + 1× 2.5GE | Wi-Fi 7 double bande | ~150 terminaux, 16 tunnels IPsec, **8 AP max, 32 équipements gérés**, 256 VLAN | Non | Passif (boîtier métal) |
| **AR280** | 4× GE (1 LAN dédié + 3 LAN/WAN) + 1× 2.5GE (WAN/LAN) | — (pas de Wi-Fi intégré d'après la fiche) | ~150 terminaux, 16 tunnels IPsec, **16 AP max, 32 équipements gérés** | Non (conso max 49 W dont 41 W PoE — à vérifier : la fiche mentionne une capacité PoE, probablement sur un modèle dérivé) | Passif |
| **AR281** | à vérifier sur la documentation officielle | — | Annoncé en 2026 comme « 1 appareil = 5 fonctions » : routage, **NVR**, **WAC**, accès ligne privée, gestion du comportement en ligne | à vérifier | à vérifier |

**Points terrain à retenir :**
- Débit VPN IPsec annoncé : **200 Mbit/s** (AR180/AR280). C'est une valeur labo. Si ton client fait du VPN site-à-site intensif, dimensionne large.
- Débit de contrôle de bande passante (egress) : **2 Gbit/s** annoncé.
- Les AR180 ont **4 antennes Wi-Fi externes intelligentes** (non démontables, livrées montées), gain max 3 dBi en 2,4 GHz / 4 dBi en 5 GHz, puissance d'émission max 23 dBm (réglementation locale applicable).
- **Pas de port console** sur AR180/AR280 d'après la fiche technique. Traduction : pas de plan B série si tu perds l'accès IP. Le reset physique existe (bouton reset), mais il coupe le service — voir section dépannage.
- Alimentation : adaptateur secteur externe, 100-240 V. Pas de redondance d'alimentation. Pour un site critique, prévois un onduleur (lien direct avec ton guide onduleurs).
- Température de fonctionnement AR180 : fiche à vérifier selon modèle (la série accepte typiquement 0 à 40/45 °C — **à vérifier sur la documentation officielle** du modèle exact que tu déploies).

## 5. Les points d'accès eKitEngine : la gamme réelle

Voici les modèles **réellement existants** d'après les fiches constructeur et les catalogues distributeurs (pas de modèles inventés) :

### 5.1. Wi-Fi 6 d'entrée de gamme (le gros des déploiements PME)

