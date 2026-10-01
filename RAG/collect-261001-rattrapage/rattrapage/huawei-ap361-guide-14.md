---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-14
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [2197, 2364]
sha256: 114e7cc103ef2163c9faaca5a0819607b38ada4697c9147fa10156ee4df27cde
---

# Huawei eKit AP361 — Guide ultra-complet

- [ ] **Visuel** : chaque AP est-il toujours à sa place, bien clipsé, LED normale ?
- [ ] **Poussière** : soufflette sur les ouïes (AP éteint si démontage, sinon
      soufflage léger sans démonter).
- [ ] **Fixations** : test de traction sur 10 % des AP (tous si < 10).
- [ ] **Câbles** : jarretières en baie non écrasées, étiquettes lisibles.
- [ ] **Radio** : tour à l'analyseur — nouveaux SSID voisins ? canaux toujours
      pertinents ? niveaux conformes à la mesure initiale ?
- [ ] **Logs** : revue des 3 derniers mois (reboots inexpliqués ? échecs d'asso
      en hausse ?).
- [ ] **Clé invités** : rotation effectuée ?
- [ ] **Spare** : l'AP de secours est-il toujours là, sur la bonne version ?

## 140. Le stock de spare : votre assurance panne

- **1 AP361 de spare** minimum par site de plus de 10 AP (ou 1 pour 20 AP sur
  des petits sites groupés).
- Le spare est **pré-configuré** (même version firmware, onboarding déjà fait
  ou procédure écrite) : en cas de panne, échange en 15 minutes.
- L'AP remplacé part en **RMA** ; le spare est remplacé dès le retour.
- Stockez-le dans son carton, au sec, avec son kit de fixation.

## 141. Nettoyage et inspection : les détails qui comptent

- **Soufflette** (air sec) sur les ouïes d'aération : la poussière isole
  thermiquement. En atelier/boulangerie : trimestriel obligatoire.
- **Ne jamais** : eau, solvant, nettoyeur haute pression (oui, on l'a déjà vu).
- **Dôme** : chiffon microfibre légèrement humide si traces.
- **Vérifiez les nids** : dans les entrepôts, les oiseaux et les toiles
  d'araignées adorent les AP chauds perchés.
- **Après travaux** : ronde systématique (peinture sur le dôme = atténuation ;
  AP déplacé par les peintres = repositionner + re-mesurer).

## 142. Revue documentaire annuelle

- [ ] Plan de masse à jour (déménagements, nouveaux AP) ?
- [ ] Tableau d'adressage à jour ?
- [ ] Plan radio toujours valable (nouveaux voisins, nouvelles cloisons) ?
- [ ] Mots de passe du coffre à jour ? Comptes des partants révoqués ?
- [ ] Release notes du firmware lues (failles corrigées depuis votre version ?)
- [ ] Contrats de garantie / RMA : dates de fin connues ?

## 143. Indicateurs de vieillissement : quand remplacer

Un AP361 n'est pas éternel. Signes de fin de vie :

- Reboots spontanés de plus en plus fréquents (hors cause réseau/PoE).
- Dégradation du débit radio sans cause externe (composants RF qui dérivent).
- Firmware **plus supporté** (plus de correctifs de sécurité) → remplacement
  planifié, pas en urgence.
- Besoin fonctionnel dépassé (densité qui explose, besoin 6 GHz) → montée en gamme.

**Durée de vie typique** : 5 à 7 ans en intérieur propre. Prévoyez le
renouvellement au **budget N+5** dès l'achat initial.

---
---

# P. DURCISSEMENT

## 144. Comptes et mots de passe : la base

- [ ] **Un compte par personne** (traçabilité). Pas de « admin » partagé.
- [ ] Mot de passe initial **changé jour 1** (ne jamais laisser le défaut).
- [ ] Robustesse : 20+ caractères, générés aléatoirement, au **coffre d'équipe**.
      Exemple **fictif** : `Grue-Torche-9917-Management!Bleu`.
- [ ] **Révocation immédiate** au départ d'un technicien/prestataire.
- [ ] Revue des comptes **semestrielle** : qui a accès ? pourquoi encore ?
- [ ] Verrouillage après N tentatives (si la fonction existe — 🔎 selon version)
      pour contrer le brute-force.

## 145. Management VLAN + restriction d'accès

- AP administrés sur le **VLAN 10 MGMT**, injoignable depuis les VLANs
  utilisateurs (règle firewall).
- Si l'AP le permet (🔎 selon version) : **ACL d'administration** n'autorisant
  que les IP des postes admin (ex. `192.168.10.100-110`).
- SSH : **jamais de Telnet** en production. SSH restreint au VLAN management.
- SNMP : v3 uniquement, communauté/credentials au coffre, restreint au superviseur.

## 146. HTTPS, certificats, chiffrement

- Interface web en **HTTPS uniquement** (désactiver HTTP si possible — 🔎).
- Certificat : le certificat auto-signé d'usine déclenche des alertes navigateur.
  Envisagez un certificat d'une **CA interne** (voir votre guide PKI) pour les
  équipements gérés en Fat.
- Vérifiez que les protocoles obsolètes sont désactivés (SSLv3, TLS 1.0/1.1 —
  🔎 selon ce que le firmware expose).

## 147. Désactiver les services inutiles

Principe : **tout ce qui n'est pas utilisé est désactivé**.

- [ ] Telnet : **désactivé** (SSH à la place, ou rien si cloud-only).
- [ ] HTTP : désactivé si HTTPS suffit.
- [ ] SNMP v1/v2c : désactivés si v3 en place (ou SNMP totalement si non utilisé).
- [ ] Découverte (LLDP/CDP) : utile en exploitation — à garder **sauf** exigence
      de discrétion, mais ne jamais diffuser d'infos sensibles dedans.
- [ ] Portail captif local : désactivé si non utilisé.

## 148. 802.1X sur le port filaire (si supporté)

Si le switch et l'AP le permettent : authentifier l'AP lui-même en 802.1X sur le
port du switch (le port ne s'ouvre que pour un AP connu). C'est le niveau
« paranoïaque sain » : quelqu'un qui débranche un AP pour brancher son PC se
retrouve avec un port mort.

🔎 Vérifiez le support côté AP361 et côté switch avant de promettre ça à la direction.

## 149. Journalisation des accès admin

- Toute connexion admin doit laisser une trace (syslog, §107).
- Revue **mensuelle** : connexions inhabituelles (nuit, week-end, IP inconnue) ?
- En cas d'incident : ces logs sont vos preuves. Sans logs, pas d'enquête.

## 150. Checklist de durcissement (audit semestriel)

- [ ] Comptes nominatifs, pas de défaut, pas de partagé
- [ ] Mots de passe robustes, au coffre, rotation annuelle minimum
- [ ] VLAN management isolé, ACL admin en place
- [ ] Telnet/HTTP/SNMP v1-v2c désactivés (ou justification écrite)
- [ ] HTTPS forcé, TLS récents
- [ ] Firmware à jour (failles connues corrigées)
- [ ] Logs d'accès admin centralisés et revus
- [ ] Comptes des prestataires/partants révoqués

## 151. En cas d'incident de sécurité : conduite à tenir

1. **Isoler** : déconnectez l'AP suspect du réseau (port du switch en shutdown)
   — ne l'éteignez pas tout de suite si vous voulez les logs volatils.
2. **Préserver** : exportez les logs (syslog, eKit) avant toute manipulation.
3. **Qualifier** : rogue AP ? Evil twin ? Compromission d'un compte admin ?
4. **Éradiquer** : reset + reconfiguration saine, rotation de **toutes** les clés,
   révocation des comptes suspects.
5. **Retour d'expérience** : fiche d'incident + mise à jour de ce guide.

---
---

# Q. PENSE-BÊTE DE POCHE

## 152. Les 10 commandes / actions qui sauvent

| # | Action | Où / comment |
|---|---|---|
| 1 | Voir l'état des AP | App eKit > Site > Équipements |
| 2 | Redémarrer un AP à distance | eKit > AP > Redémarrer (ou cycle PoE du port switch) |
| 3 | Voir la vitesse du port | eKit (état AP) ou `display interface` sur le switch |
| 4 | Changer un canal | eKit > Radio > canal manuel |
| 5 | Voir les clients d'un AP | eKit > AP > Clients |
| 6 | Sauvegarder la config (Fat) | `save` (+ copie vers serveur) |
| 7 | Voir les logs | Syslog central / eKit |
| 8 | Tester un câble | Testeur wiremap + longueur |
| 9 | Tester le PoE | Testeur PoE (tension 44-57 V) |
| 10 | Reset d'usine | Trombone > 5 s, attendre 2-4 min |

## 153. Valeurs à connaître par cœur

- AP361 : **2x2 ax, 575 + 1200 Mbit/s, 1 GE, 802.3af, 8,8 W**
- 2,4 GHz : canaux **1/6/11**, **20 MHz** toujours
- 5 GHz : **36-48** sans DFS, **40 MHz** par défaut
- Puissance de départ : **14 dBm**, cible **-60 dBm** en zone de travail
- Voix : **-67 dBm minimum**, **802.11r** activé
- PoE : **100 m max**, cuivre 100 %, budget = Σ + 20 %
- SSID : **max 4**, diffusés, clés 20+ caractères
- Un AP361 ≈ **40-50 clients** confortables

## 154. Numéros et contacts du site (à compléter)

