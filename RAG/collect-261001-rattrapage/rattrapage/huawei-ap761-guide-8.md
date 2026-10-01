---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-8
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [647, 756]
sha256: c1dd5f131bfb2b7dc3ffae34062a647fcdadc7977fcdc521a70515dddc1cb4e5
---

# Guide ultra-complet — Huawei eKit AP761

**Test du câble avant de quitter le chantier :** testeur de continuité a minima, **testeur qualifiant** (longueur, diaphonie, PoE) idéalement. Un câble qui « marche » à 100 Mbps au lieu du Gbps à cause d'une paire abîmée = un AP bridé dont personne ne comprendra la cause pendant des mois. **Noter les résultats dans le dossier de site.**

## 25. PoE : budget, injecteurs, switches — calculs

**Rappel valeurs constructeur :** 802.3at/af, **17.7 W max**, fonctions restreintes en 802.3af.

**Dimensionner le budget — la méthode :**
```
Par AP761 :  17.7 W (AP)
           +  2.5 W (pertes câble 100 m en PoE+)
           + 20 % de marge
           ≈ 25 W à réserver par AP sur le budget du switch
```

**Tableau de dimensionnement :**
| Switch (budget PoE total) | AP761 alimentables (règle des 25 W) |
|---|---|
| 8 ports, 130 W | 5 |
| 8 ports, 240 W | 9 (limité par le nb de ports) |
| 24 ports, 370 W | 14 |
| 24 ports, 740 W | 24 (limité par le nb de ports) |

**Injecteur PoE (quand le switch n'est pas PoE) :**
- Prendre un injecteur **802.3at, 30 W, gigabit** (pas 100 Mbps).
- L'injecteur se place **côté local technique**, pas en extérieur (sauf modèle extérieur — cher et rare).
- Un injecteur par AP : ça fait du câble et des prises 230 V en plus — à 4 AP et plus, un switch PoE devient plus propre et souvent moins cher.

**Chaînage interdit :** pas de PoE « pass-through » en cascade sur l'AP761 (il n'a pas de port PoE-Out — à vérifier sur la fiche du modèle exact, mais le datasheet ne mentionne qu'un PoE-In). Chaque AP = un câble PoE home-run jusqu'au switch/injecteur.

**Onduleur et PoE — ton métier, Zelef :**
- Un switch PoE 24 ports plein peut tirer **400–800 W** : c'est une charge à intégrer dans le dimensionnement de l'onduleur du local technique (voir ton guide onduleurs : `~/workspace/user/files/onduleurs_ups_guide.md`).
- En cas de coupure, les AP sont souvent les premiers qu'on accepte de perdre — ou les derniers, si le Wi-Fi porte la téléphonie ou le contrôle d'accès. **Décider à l'avance** et le noter dans le plan de délestage.
- **Astuce :** programmer l'extinction des SSID non critiques sur ordre (via le cloud eKit ou un script) pour allonger l'autonomie sur batterie.

## 26. Protection foudre et mise à la terre

L'AP761 a une protection **6 kV sur les ports Ethernet** [constructeur] : c'est une protection contre les **surtensions induites**, pas contre la foudre directe. Un AP sur un mât en extérieur est une cible : la protection se conçoit en couches.

**Les 3 couches :**
1. **Paratonnerre / zone de protection :** si le site est en zone à risque (plateau exposé, pylône), l'AP doit être dans le volume de protection d'un paratonnerre (règle de la sphère fictive — voir un spécialiste, norme NF EN 62305).
2. **Parafoudre Ethernet (SPD) :** un parafoudre **RJ45 PoE-compatible** (il doit laisser passer le PoE !) monté **à chaque extrémité** du câble : un près de l'AP, un à l'entrée du bâtiment. Raccordé à la terre avec un conducteur **court et rectiligne** (16 mm² cuivre typique — selon norme locale).
3. **Mise à la terre de l'AP et du mât :** le boîtier métal et le mât sont reliés à la terre. Le blindage du câble est mis à la terre **d'un seul côté** (côté bâtiment) pour éviter les boucles de masse — sauf consigne contraire du fabricant du parafoudre.

**Ce qu'il ne faut PAS faire :**
- ❌ Compter sur les 6 kV internes comme seule protection sur un site exposé.
- ❌ Mettre un parafoudre Ethernet non-PoE : il bloque ou dégrade l'alimentation.
- ❌ Raccorder la terre du parafoudre sur une « terre » flottante ou sur la tuyauterie sans continuité.
- ❌ Oublier la **vérification annuelle** : un parafoudre qui a encaissé une surtension est peut-être mort sans le montrer (voyant d'état à contrôler, chap. 101).

**Après un orage :** si l'AP ne répond plus, tester dans l'ordre : injecteur/switch PoE → parafoudre (le shunter temporairement pour tester) → câble → AP. Les parafoudres sacrifient souvent leur vie pour sauver l'AP : c'est leur travail.

## 27. Checklist pré-installation chantier

À cocher **avant** de monter sur l'échelle. Un chantier préparé = une demi-journée ; un chantier improvisé = deux jours et trois allers-retours.

**Site :**
- [ ] Emplacement validé sur plan avec le secteur de 65° tracé (chap. 22).
- [ ] Support vérifié (porteur, vertical, rigide).
- [ ] Accès maintenance future réfléchi (échelle, nacelle, chemin).
- [ ] Autorisations obtenues (copropriété, mairie si façade visible/protégée).

**Électrique / réseau :**
- [ ] Chemin de câble défini et mesuré (< 100 m ou fibre prévue).
- [ ] Switch PoE 802.3at avec **budget suffisant** (chap. 25) ou injecteur 30 W gigabit.
- [ ] Parafoudres prévus si site exposé (chap. 26).
- [ ] Prise 230 V pour l'injecteur si besoin, sur circuit secouru si le Wi-Fi est critique.

**Logique :**
- [ ] Plan d'adressage : VLAN management, SSID, sous-réseaux.
- [ ] SSID et clés **générés** (pas « on verra sur place ») — clés fictives notées dans le dossier de site, jamais « admin/admin ».
- [ ] Compte cloud eKit prêt si mode cloud ; AC joignable si mode Fit.

**Matériel :**
- [ ] AP761 + kit de montage + presse-étoupes vérifiés (chap. 5).
- [ ] Câble extérieur blindé 100 % cuivre en quantité suffisante (+ 10 % de mou).
- [ ] Connecteurs RJ45 blindés Cat6 + capuchons.
- [ ] Outillage : perceuse, chevilles adaptées, niveau, testeur câble, smartphone avec app eKit, étiqueteuse.

**Le jour J :**
- [ ] Photo de l'étiquette AP (S/N, MAC, QR) avant montage.
- [ ] Test du câble **avant** raccordement final.
- [ ] Mise sous tension, vérification LED, adoption cloud/Fat (chap. 28–35).
- [ ] Test de couverture : tour de la zone avec un smartphone (relevés RSSI, chap. 60).
- [ ] Photos « après » + mise à jour du dossier de site.

## 28. Première mise en route : les 3 chemins

Selon le mode choisi (chap. 11), la première mise en route suit l'un de ces trois chemins. **Ne mélange pas les méthodes** : un AP à moitié configuré en Fat puis adopté en cloud, c'est une heure de perdue.

| Mode | Chemin | Chapitres |
|---|---|---|
| **Cloud** (recommandé PME multi-sites) | App eKit → scan QR → créer/adopter dans un site → config poussée par le cloud | 29, 30 |
| **Fat** (autonome) | Web local ou CLI → config manuelle complète | 31, 34 |
| **Fit** (contrôleur) | L'AP cherche l'AC (DHCP option 43 ou DNS) → l'AC pousse la config | 32 |

**Dans tous les cas, l'ordre est :**
1. Alimenter l'AP (PoE) — attendre le démarrage complet (LED verte fixe, ~2–3 min).
2. L'AP obtient une IP (DHCP du réseau — prévoir un serveur DHCP sur le VLAN de management).
3. Adopter/configurer selon le mode.
4. Vérifier : SSID diffusés, client test connecté, débit testé, supervision OK.
5. **Sauvegarder** (chap. 35).

**Piège n° 1 : le VLAN de management.** Si ton switch envoie le port en trunk avec un VLAN natif différent de celui où est le DHCP, l'AP n'aura jamais d'IP. Pour la mise en route, le plus simple : port en **access** sur le VLAN de management, DHCP actif. On passe en trunk après (chap. 52).

**Piège n° 2 : l'heure.** Un AP dont l'horloge est fausse (pas de NTP) génère des certificats/logs incohérents et peut refuser de se connecter au cloud (TLS). Configurer le NTP dès le début (chap. 70).

## 29. Onboarding via l'app HUAWEI eKit (QR code)

L'app **HUAWEI eKit** (Android/iOS, gratuite) est la télécommande de l'écosystème : onboarding, supervision, alertes. C'est la méthode la plus rapide pour un AP761 en mode cloud.

