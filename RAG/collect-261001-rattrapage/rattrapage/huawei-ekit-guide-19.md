---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-19
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Huawei"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1630, 1705]
sha256: 589e3c82b5af9f6fa3c1d4bfea760bfa3b02122e116af0ed061e8abf0f9fbc8d
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

| Critère | Huawei eKit | Concurrent cloud avec licence/an | Solution « box grand public » |
|---|---|---|---|
| Coût gestion cloud | Annoncé **gratuit/sans licence** | Licence annuelle par équipement | Gratuit mais pas de gestion |
| Déploiement | App mobile, scan QR, rapide | App/cloud aussi, parfois plus guidé | Manuel, un par un |
| Multi-sites | Oui (SNC) | Oui | Non |
| Fonctions avancées | Limitées (pas de NAC, routage simple) | Variables, souvent + complètes | Quasi inexistantes |
| Wi-Fi 7 | Oui (AP572/673/772) | Oui chez la plupart | Rare |
| Verrouillage | Écosystème Huawei | Écosystème du concurrent | Aucun (mais aucun support) |
| Support local | Distributeur Gold | Variable | Aucun |

**Lecture :** eKit gagne sur le **rapport fonctions/prix** en PME simple. Il perd dès que le besoin devient « entreprise » (section 83). Ne dénigre jamais un concurrent devant un client : compare des **faits** (prix, fonctions, support), pas des slogans.

## 150. Procédure — reset usine et ré-onboarding propre

1. **Sauvegarde** la config du site (section 89) — même si l'équipement est en vrac, la config du *site* dans le cloud reste la référence.
2. Trouve le **bouton reset** (trou ou bouton selon modèle — voir le guide d'installation du modèle).
3. **Équipement sous tension**, appuie 5-10 s (durée exacte **à vérifier sur la documentation officielle** du modèle) jusqu'au changement de LED.
4. **Attention :** le reset interrompt le service (l'AP/la passerelle redémarre d'usine). Jamais en journée sur un site en production.
5. Après redémarrage : l'équipement repasse en mode setup → **ré-onboard** via l'app (sections 19-20) sur le bon site → il récupère la configuration du site.
6. Vérifie : SSID diffusés, VLAN, PoE, état vert dans le cloud.
7. Note l'opération dans le registre (106) : un reset n'est jamais « pour rien », il y avait une cause — écris-la.

## 151. Procédure — remplacement RMA sans coupure longue

1. Dès la panne : **équipement de secours** (ton stock, section 105) onboardé sur le site — il prend la config automatiquement. Coupure : 10-15 min.
2. Déclare la panne au **distributeur** : SN, facture, symptômes, photos, ce que tu as déjà testé (un RMA bien documenté va 3× plus vite).
3. À réception du remplacement : onboard sur le site, **vérifie la config**, remets l'équipement de secours en stock.
4. Mets à jour la **fiche site** (nouveau SN) et le registre.
5. Analyse : pourquoi est-il tombé en panne ? (surtension → revoir la protection ; surchauffe → revoir la ventilation). Un RMA sans analyse = un 2e RMA dans 6 mois.

## 152. Procédure — changer le PSK invités sur N sites (multi-sites)

1. Génère le **nouveau PSK** (20+ caractères, gestionnaire de mots de passe) — jamais un mot simple, jamais l'ancien + « 1 ».
2. Dans le cloud, modifie le SSID invités du **modèle de configuration** (pas site par site).
3. Applique au **site pilote** d'abord, teste avec un téléphone (connexion + portail).
4. Déploie **site par site en heures creuses** (pas les 12 boutiques à 12 h un samedi).
5. Communique le nouveau PSK aux responsables (canal sûr, pas d'affichage public en clair si portail à code).
6. Mets à jour le **coffre** (102) et note la date de rotation (prochaine : +12 mois max, ou au départ d'un employé sensible).

## 153. Procédure — ajouter un VLAN sur un site existant (sans tout casser)

1. **Planifie** : ID VLAN libre, sous-réseau (jamais en conflit), usage, quels ports/AP concernés.
2. **Sauvegarde** la config actuelle (89).
3. Crée le VLAN + le scope DHCP dans le cloud, **sans l'appliquer** aux ports d'abord.
4. Applique sur **1 port test** (ou 1 AP test) : branche un PC/téléphone, vérifie IP + connectivité + isolation.
5. Déploie port par port (ou groupe d'AP par groupe), en vérifiant à chaque étape.
6. Mets à jour : fiche site (65), plan de brassage, documentation.
7. **Heures creuses** obligatoires si le VLAN touche des ports en production (une erreur de VLAN = des utilisateurs coupés).

## 154. Procédure — portail captif : mise en place pas à pas (logique)

1. Choisir le **mode** : portail de l'AP (simple, sans serveur) ou portail de l'USG (plus d'options : base locale, RADIUS, AD).
2. Créer le **VLAN invités** (20/30) avec DHCP, DNS fonctionnel, **isolation** activée.
3. Créer le **SSID invités** rattaché au VLAN, activer le portail captif.
4. Personnaliser la **page** : logo client, CGU, mentions légales, langue(s).
5. Définir les **quotas** : débit/client (ex. fictif : 10 Mbit/s), durée de session (ex. : 4 h commerce, 24 h hôtel), volume si besoin.
6. Tester sur **Android + iPhone** (les deux), avec et sans VPN, avec date/heure correctes et incorrectes (cas n°7).
7. Préparer le **PSK de secours** et l'affichette d'aide (section 97).
8. Former l'accueil : « voici ce que vous dites au client qui n'y arrive pas » (2 phrases, pas un manuel).

## 155. Procédure — audit de sécurité annuel d'un site eKit

- [ ] Firmwares : tous à N ou N-1 ? (sinon : planifier MAJ)
- [ ] Mots de passe : PSK tournés dans l'année ? Admin par défaut changés partout ?
- [ ] Accès cloud : revue des comptes (qui a quoi ? ex-employés retirés ?)
- [ ] VLAN : invités toujours isolés ? Aucun équipement « oublié » sur le mauvais VLAN ?
- [ ] SSID : toujours 3 max ? Aucun SSID « TEST » ou « TEMP » qui traîne ?
- [ ] SNMP : v3 uniquement ? Strings par défaut changées ?
- [ ] Ports inutilisés du switch : **désactivés** (un port libre = une prise d'entrée physique)
- [ ] Sauvegardes : existent, récentes, testées ?
- [ ] Onduleur : test batterie OK ? Autonomie conforme ?
- [ ] Photos/baie : étiquetage toujours lisible ?
Chaque point non conforme = un ticket avec échéance. L'audit tient en 2 h et évite 90 % des incidents « bêtes ».

## 156. Erreurs classiques n°17 à 25 (la suite du dépannage)

