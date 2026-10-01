---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-1
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [1, 158]
sha256: 0b0e9ed942bf831880df164979d705dc8ff2e40e625ec20c0f011acf6e138157
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

> **Guide de terrain, pas une brochure.** Ce document couvre le Huawei NetEngine AR720 de A à Z :
> hardware, démarrage, WAN, LAN, NAT, routage, VPN, sécurité, QoS, supervision, dépannage terrain.
> Rédigé en français pour un usage opérationnel (chef de service systèmes & énergies et son équipe).
>
> **Contexte de vérification :** les valeurs constructeur citées viennent de la fiche datasheet
> « Huawei NetEngine AR700 Series Enterprise Router Data Sheet » (AR720 / AR730).
> La version VRP exacte du firmware varie selon le lot et le fournisseur — **à vérifier sur la fiche
> du modèle exact** livré (commande `display version` une fois le routeur en main).
> Quand une valeur est incertaine ou dépend de la version, c'est écrit explicitement.

---

## Sommaire

- [Partie A — Prise en main](#partie-a--prise-en-main)
- [Partie B — WAN : accès Internet](#partie-b--wan--accès-internet)
- [Partie C — LAN et services locaux](#partie-c--lan-et-services-locaux)
- [Partie D — NAT](#partie-d--nat)
- [Partie E — Routage](#partie-e--routage)
- [Partie F — VPN](#partie-f--vpn)
- [Partie G — Sécurité](#partie-g--sécurité)
- [Partie H — QoS](#partie-h--qos)
- [Partie I — Supervision](#partie-i--supervision)
- [Partie J — Exploitation : sauvegarde, firmware, licences](#partie-j--exploitation--sauvegarde-firmware-licences)
- [Partie K — Dépannage terrain (15+ cas)](#partie-k--dépannage-terrain-15-cas-détaillés)
- [Partie L — Maintenance préventive](#partie-l--maintenance-préventive)
- [Partie M — Durcissement](#partie-m--durcissement)
- [Partie N — Cas pratiques complets commentés](#partie-n--cas-pratiques-complets-commentés)
- [Partie O — Pense-bête de poche](#partie-o--pense-bête-de-poche)
- [Partie P — Glossaire](#partie-p--glossaire)
- [Partie Q — Quiz](#partie-q--quiz-10-questions)
- [Partie R — Pour aller plus loin](#partie-r--pour-aller-plus-loin)

---

# Partie A — Prise en main

## 1. Positionnement du produit : c'est quoi, l'AR720 ?

Le Huawei NetEngine AR720 est un **routeur de branche d'entreprise** (segment PME / agences /
boutiques / établissements scolaires), dans la série NetEngine AR700.
Il intègre sur une seule plateforme 1U : routage, commutation, firewall, VPN, NAT, contrôleur
Wi-Fi (gestion d'AP Huawei via CAPWAP), et options 4G/5G/voix via cartes d'extension.

Concrètement, sur le terrain : c'est la boîte qu'on pose dans une agence pour sortir sur
Internet, monter des tunnels VPN vers le siège, faire du NAT, du DHCP, du filtrage, et
éventuellement basculer sur une 4G en secours. Le tout piloté par **VRP** (Versatile
Routing Platform), le même OS que les gros routeurs Huawei — donc la syntaxe CLI qu'on
apprend ici se réutilise sur du NetEngine 8000.

## 2. Caractéristiques hardware officielles (datasheet AR700)

| Élément | Valeur constructeur |
|---|---|
| Processeur | ARM64 4 cœurs, 1,4 GHz |
| Performance de commutation | 9 à 25 Mpps |
| Débit sortant (forwarding) | 4 Gbit/s |
| Nombre de terminaux connectés (capacité) | 700 |
| Ports WAN fixes | 2 x GE combo (RJ45 cuivre / SFP optique) |
| Ports LAN fixes | 8 x GE électriques RJ45 (configurables en ports WAN) |
| Slots SIC (Smart Interface Card) | 2 |
| Slot WSIC | 0 par défaut / 1 maximum |
| Slot MIC / XSIC | 0 / 0 (non supportés sur l'AR720) |
| Port console | 1 x RJ45 console série |
| Ports USB | 2 x USB 2.0 |
| Mémoire RAM | 4 Go |
| Flash | 1 Go (512 Mo utilisables client, note datasheet) |
| Hot swapping | Supporté |
| Châssis | 1U, 44,4 x 442,0 x 220,4 mm, 3,2 kg |
| Alimentation | AC intégrée, 100–240 V, 50/60 Hz (plage max 90–264 V, 47–63 Hz) |
| Consommation | 33 W max (28 W typique) |
| Ventilation | Module ventilateur intégré non extractible (built-in unpluggable), flux gauche → droite |
| Bruit | 52,3 dB(A) |
| Température de fonctionnement | 0 °C à 45 °C |
| Humidité | 5 % à 95 % sans condensation |
| Alimentation redondante | Non supportée |
| PoE | Non supporté |

Note terrain : 4 Gbit/s de débit sortant, c'est **la capacité de la plateforme**, pas un débit
garanti avec tous les services activés. Dès qu'on active firewall + IPS + VPN + QoS en même
temps, le débit utile réel chute. Dimensionner avec de la marge (viser 50–60 % de la capacité
en charge nominale).

## 3. Ce que le « combo » WAN veut dire en pratique

Un port « combo » = une paire RJ45 cuivre + cage SFP qui partagent **la même interface logique**.
On ne peut utiliser qu'un seul des deux médias à la fois pour un port donné.

- `GE0/0/0` : WAN1 combo (RJ45 ou SFP).
- `GE0/0/1` : WAN2 combo (RJ45 ou SFP).
- `GE0/0/2` à `GE0/0/9` : les 8 ports LAN RJ45.

La commutation du média sur un port combo se fait en général en CLI (`combo enable` ou
équivalent selon la version VRP — **à vérifier sur la fiche du modèle exact**, la syntaxe
varie selon les versions V300).

Piège classique : brancher un RJ45 sur le port WAN1 alors que le SFP est déjà inséré et actif —
ou l'inverse. Si le lien ne monte pas, vérifier d'abord quel média est actif :

```
display interface GigabitEthernet 0/0/0
```

## 4. Les slots d'extension SIC : à quoi ça sert

Les 2 slots SIC (Smart Interface Card) acceptent des cartes filles. Typiquement sur la série
AR (à valider avec le catalogue du fournisseur pour l'AR720, **à vérifier sur la fiche du
modèle exact**) :

- cartes 4G LTE / 5G (modem cellulaire avec slot SIM) — très utile en backup WAN ;
- cartes E1/T1 (liaisons louées classiques) ;
- cartes voix (FXS/FXO) ;
- cartes GE supplémentaires ou modules fibre.

Le hot swapping est supporté : on peut insérer/retirer une carte à chaud. Règle de prudence
terrain : on ne fait jamais ça en pleine production sans fenêtre de maintenance, même si le
constructeur l'autorise. Un retrait brutal peut faire tomber la carte et perturber le slot.

Le slot WSIC (0 par défaut, 1 max) correspond aux cartes Wi-Fi intégrées (l'AR720 peut
embarquer un point d'accès Wi-Fi via carte WSIC selon la référence commerciale — **à vérifier
sur la fiche du modèle exact**).

## 5. Face avant / face arrière : ce qu'on y trouve

**Face avant (panneau de ports) :**

- 2 ports combo WAN (étiquetés WAN1 / WAN2, chacun avec RJ45 + cage SFP).
- 8 ports LAN RJ45 (étiquetés LAN1…LAN8 ou GE0/0/2…GE0/0/9).
- 1 port console RJ45.
- 2 ports USB (USB 2.0).
- LED d'état : SYS (système), PWR (alimentation), et LED par port (link/act).

**Face arrière :**

- Prise d'alimentation AC (cordons standard).
- Sortie d'air (le flux va de gauche à droite — ne pas coller le routeur contre un mur à gauche,
  laisser de l'espace pour l'entrée d'air).

**Lecture des LED (convention Huawei classique) :**

| LED | État | Signification |
|---|---|---|
| PWR | Vert fixe | Alimentation OK |
| PWR | Éteinte | Pas d'alimentation |
| SYS | Vert clignotant lentement | Démarrage en cours / fonctionnement normal |
| SYS | Rouge | Panne système ou boot échoué |
| Port | Vert fixe | Lien établi |
| Port | Vert clignotant | Lien + trafic |
| Port | Éteinte | Pas de lien |

Règle d'or du premier diagnostic : **SYS rouge + aucun port allumé = problème de boot**
(firmware corrompu, clé USB de secours à préparer). SYS vert + port éteint = problème de
câble / média combo / négociation auto.

## 6. Alimentation et environnement : les contraintes à respecter

