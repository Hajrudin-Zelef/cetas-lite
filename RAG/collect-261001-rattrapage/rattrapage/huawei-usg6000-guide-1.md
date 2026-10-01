---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-1
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [1, 115]
sha256: 0324240dabc185b8b86021a68dcf06621b50557eed4980c3b3a78b9062cf7da3
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

> **Public :** Zelef, chef de service systèmes & énergies — ton terrain, concret, direct.
> **Périmètre :** série classique **USG6000** (USG6305 → USG6680), logiciel **V500R005C20 / V600R007** (à vérifier sur la fiche du modèle exact).
> **Sources vérifiées (sept. 2026) :** fiche technique Huawei USG6000 Series NGFW Overview, specs matérielles par modèle (switch-router-olt.com reprenant les datasheets Huawei), guide « Quick Configuration Guide (With New Web UI) » basé sur USG6000E V600R007C00, white paper technique HUAWEI Secospace USG6000 Series.
> **Règle d'or :** toute valeur marquée « à vérifier sur la fiche du modèle exact » doit être confirmée avant dimensionnement ou achat. Les mots de passe des exemples sont fictifs.
> **Il existe déjà** un guide général `huawei_usg_guide.md` (684 lignes) : celui-ci va beaucoup plus loin — configs complètes, dépannage terrain approfondi, UTM avec impact perf honnête, HA, durcissement.

---

## Sommaire rapide

| Bloc | Sections | Contenu |
|------|----------|---------|
| A | 1–10 | Hardware : modèles, ports, LED, alimentation, console, capacités |
| B | 11–18 | Initialisation : premier démarrage, mot de passe, management, upgrade |
| C | 19–28 | Zones de sécurité |
| D | 29–44 | Politiques de sécurité (interzone) |
| E | 45–56 | NAT |
| F | 57–72 | VPN : IPSec, SSL VPN, GRE, L2TP |
| G | 73–88 | UTM : AV, IPS, URL, applicatif, anti-spam + impact perf |
| H | 89–96 | Authentification : locale, AD/LDAP, RADIUS |
| I | 97–106 | Haute disponibilité (VGMP/HRP) |
| J | 107–114 | Logs, rapports, supervision |
| K | 115–124 | Sauvegarde, firmware, licences |
| L | 125–132 | Maintenance préventive |
| M | 133–142 | Durcissement |
| N | 143–160 | Dépannage terrain (18 cas) |
| O | 161–168 | Cas pratiques complets commentés |
| P | 169–176 | Erreurs classiques à ne pas commettre |
| Q | 177–185 | Pense-bête de poche, glossaire, quiz, pour aller plus loin |

**Comment utiliser ce guide :**
- Cherche un sujet avec `Ctrl+F` sur le numéro de section.
- Tous les blocs ` ``` ` marqués `huawei` sont des commandes CLI VRP à coller en vue système (`system-view`).
- Les chemins web correspondent à la « New Web UI » (V600R007) ; sur V500, les menus peuvent différer légèrement.
- ⚠️ = point qui casse en prod si on l'oublie. 🔧 = commande de diagnostic. 📋 = checklist.

---

# BLOC A — HARDWARE : LA SÉRIE USG6000

## 1. La série USG6000 en une page

La série **Huawei USG6000** est la gamme de firewalls **Next-Generation (NGFW/UTM)** historique de Huawei : pare-feu stateful + VPN + UTM (antivirus, IPS, filtrage URL, contrôle applicatif, anti-spam, Anti-DDoS) dans un seul boîtier, piloté par le système **VRP** (le même que les routeurs/switchs Huawei — si tu connais `system-view`, tu es chez toi).

Trois générations coexistent :
- **USG6000 classique** (celle de ce guide) : V500R005C20 et versions antérieures.
- **USG6000E** : successeur logiciel/matériel, V600R006/V600R007, « New Web UI ».
- **USG6000F / HiSecEngine** : génération « AI firewall », la plus récente.

En pratique terrain, 90 % des commandes CLI de ce guide fonctionnent à l'identique sur USG6000 et USG6000E. Quand ça diverge, c'est signalé.

## 2. Les modèles de la série (débits constructeur vérifiés)

Débits firewall constructeur (paquets UDP 1518 octets, valeurs marketing — **divise par 2 à 3 pour estimer le débit UTM réel**, voir section 73) :

| Modèle | Format | Débit firewall | Ports en standard | Usage typique |
|--------|--------|----------------|-------------------|---------------|
| USG6305 / 6305-W | Desktop | 500 Mbit/s | 4 × GE | Très petite agence, Wi-Fi intégré (-W) |
| USG6310S / -W / -WL-OVS | Desktop | 1 Gbit/s | 8 × GE | Petite agence |
| USG6320 | Desktop | 2 Gbit/s | 8 × GE + slot extension | Agence moyenne |
| USG6330 | 1U | 1 Gbit/s | 4 × GE + 2 × combo | Petite PME (rack) |
| USG6350 | 1U | 2 Gbit/s | 4 × GE + 2 × combo | PME |
| USG6360 | 1U | 3 Gbit/s | 4 × GE + 2 × combo | PME |
| USG6370 | 1U | 4 Gbit/s | 8 × GE + 4 × SFP | PME / petite industrie |
| USG6380 | 1U | 6 Gbit/s | 8 × GE + 4 × SFP | PME exigeante |
| USG6390 | 1U | 8 Gbit/s | 8 × GE + 4 × SFP | ETI, petit datacenter |
| USG6620 | 1U | 12 Gbit/s | 8 × GE + 4 × SFP | ETI, campus |
| USG6630 | 1U | 16 Gbit/s | 8 × GE + 4 × SFP | Campus, pré-datacenter |
| USG6650 | 3U | 20 Gbit/s | 2 × 10GE + 8 × GE + 8 × SFP | Datacenter |
| USG6660 | 3U | 25 Gbit/s | 2 × 10GE + 8 × GE + 8 × SFP | Datacenter |
| USG6670 | 3U | 35 Gbit/s | 4 × 10GE + 16 × GE + 8 × SFP | Gros datacenter |
| USG6680 | 3U | 40 Gbit/s | 4 × 10GE + 16 × GE + 8 × SFP | Très gros datacenter |

Notes terrain :
- La gamme couvre **500 Mbit/s → 40 Gbit/s** en identification applicative ; la capacité **IPS + AV monte jusqu'à ~15 Gbit/s** sur les gros modèles (valeur constructeur, conditions de test idéales).
- Les modèles 6650→6680 acceptent des **cartes d'extension WSIC** (8 × GE RJ45, 2 × 10GE SFP+ + 8 × GE, 8 × GE SFP, bypass 4 × GE) et des **cartes XSIC** sur 6660.
- ⚠️ Les débits ci-dessus sont des **débits firewall purs**. Avec UTM activé (IPS+AV+URL), compte **30 à 50 %** de ces valeurs selon le mix de trafic. Dimensionne TOUJOURS sur le débit UTM, pas sur le débit firewall (section 73).

## 3. Ports physiques : ce que tu branches où

Sur un USG6000 typique (exemple USG6680, à adapter au modèle exact) :
- **Port console** : 1 × RJ45 (+ 1 × Mini-USB sur les gros modèles — un seul utilisable à la fois). Réglages : **9600-8-N-1**, câble console RJ45/DB9 ou USB-série.
- **Port de management hors-bande** : 1 × RJ45 (souvent `GE0/0/0` en interface de management dédiée).
- **Ports de service** : GE électriques RJ45 10/100/1000, SFP (fibre 1G), SFP+ (10G) selon modèle.
- **USB 2.0** : 1 à 2 ports — servent à l'**upgrade firmware**, l'import/export de config, les mises à jour de signatures UTM via clé USB.
- **Slots WSIC/XSIC** : cartes d'extension (modèles 3U).
- **Disques durs SAS 2,5"** (300/600/1200 Go, optionnels) : **RAID1** possible avec 2 disques identiques, **hot-swap** — servent au **stockage local des logs et rapports**. Sans disque, les logs partent en syslog ou se perdent au reboot (buffer limité).

## 4. LED et boutons : lecture rapide en baie

| LED | État | Signification |
|-----|------|---------------|
| PWR | Vert fixe | Alimentation OK |
| PWR | Éteint | Pas d'alimentation / module HS |
| SYS | Vert clignotant lent | Démarrage en cours |
| SYS | Vert fixe | Système démarré, OK |
| SYS | Rouge | Alarme système (voir logs) |
| ACT (par port) | Vert clignotant | Trafic sur le port |
| LINK (par port) | Vert fixe | Lien établi |
| HA | Vert | État HA normal (selon modèle) |

🔧 En cas de doute : `display device` et `display alarm` donnent l'état matériel vu par le système (section 144).

## 5. Alimentation : AC, DC, redondance

- **Modèles desktop/1U** : alimentation **interne ou externe** selon modèle, **non redondante** en général — prévois un onduleur (voir ton guide onduleurs) et, pour les sites critiques, un second USG en HA plutôt qu'une alim redondante.
- **Modèles 3U (6650→6680)** : **1+1 redondance**, modules **hot-swap**, versions **AC** (100–240 V) et **DC** (-48 V/–60 V télécom) disponibles.
- USG6680-AC : jusqu'à 700 W par module ; USG6680-DC : 350 W (valeurs à vérifier sur la fiche du modèle exact).
- ⚠️ **Ne mélange jamais** un module AC et un module DC dans le même châssis. En redondance 1+1, les deux modules doivent être **identiques**.
- Sens du flux d'air : **aspiration avant/côté gauche, extraction côté droit** (vu de l'arrière) — ne plaque pas le flanc droit contre un mur en baie.

## 6. Console et premier accès : la méthode qui marche toujours

