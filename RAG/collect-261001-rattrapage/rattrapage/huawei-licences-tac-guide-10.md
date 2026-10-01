---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-10
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [1504, 1660]
sha256: 52c7bb4c1c0c2277f7b7e6154f79a759f6a52ee182fa7286610bd732e858d49d
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

| Terme | Signification |
|---|---|
| **Numéro RMA** | Référence unique de l'échange, générée par Huawei après validation du diagnostic |
| **AR** (Advance Replacement) | La pièce neuve est expédiée avant le retour de la défectueuse |
| **RFR** (Return For Repair) | Vous renvoyez d'abord ; réparation sous ~30 jours ouvrés |
| **DOA** (Dead On Arrival) | Matériel défectueux dès la réception → procédure accélérée |
| **Faulty part** | La pièce défectueuse à retourner |
| **Replacement part** | La pièce de remplacement reçue |

## 92. Étape 1 — Constater et qualifier la panne

Avant tout appel :

1. [ ] Décrire le symptôme (méthode section 79).
2. [ ] Vérifier les causes triviales : alimentation, câble, SFP, température
      (`display environment`, `display device`), config récente.
3. [ ] Collecter `display diagnostic-information` (section 76).
4. [ ] Prendre en photo les LED et l'étiquette S/N.
5. [ ] Déterminer la sévérité (S1–S4, sections 63–68) selon l'impact réel.

**Ne déclarez pas « matériel HS » trop vite** : ~30 % des RMA demandés
s'avèrent être des problèmes de configuration ou d'environnement
(alimentation, chaleur). Le TAC le vérifiera de toute façon à l'étape 3.

## 93. Étape 2 — Ouvrir le ticket TAC

- Ouvrez le ticket selon la partie E (template section 70).
- Précisez que vous suspectez une **panne matérielle** et demandez un
  diagnostic en vue d'un RMA.
- Joignez : diagnostic-information, photos LED, logbuffer autour de la panne.

## 94. Étape 3 — Le diagnostic à distance avec le TAC

L'ingénieur TAC va :

1. Analyser vos logs et la sortie diagnostic.
2. Vous faire exécuter des tests complémentaires (boucles, changement de
   port/SFP, boot sur une autre partition...).
3. Écarter les causes logicielles et environnementales.
4. **Statuer** : panne matérielle confirmée → passage à l'étape 4 ;
   cause non matérielle → poursuite du diagnostic logiciel.

**Votre rôle :** exécuter précisément ce qui est demandé, noter les résultats,
rester joignable. Chaque test non fait = 24h de retard.

## 95. Étape 4 — Validation du remplacement et numéro RMA

Quand la panne matérielle est confirmée :

1. Le TAC (ou le partenaire selon le circuit) génère le **numéro RMA**.
2. Vous recevez : le n° RMA, la référence de la pièce expédiée, le délai
   annoncé selon votre palier (section 85), les instructions de retour.
3. **Notez le n° RMA dans votre registre** et communiquez-le à l'équipe.

> Le délai contractuel (NBD, 4h...) court à partir de la **génération du
> n° RMA**, pas de l'ouverture du ticket. D'où l'intérêt d'un diagnostic
> rapide et bien documenté.

## 96. Étape 5 — Expédition de la pièce de remplacement

- **Standard 9x5xNBD** : RMA généré avant 15h00 → pièce reçue le jour ouvré
  suivant ; après 15h00 → enregistrement le jour ouvré suivant, réception
  le surlendemain.
- **Premier 24x7x4 (P1/P2)** : pièce reçue sous 4h, 24x7.
- **Basic (NBD-S)** : pièce expédiée le jour ouvré suivant (réception J+2/J+3).
- Suivez le tracking communiqué ; prévenez le site de réception.

## 97. Étape 6 — Réception et vérification du matériel de remplacement

À la réception, **avant** de renvoyer le défectueux :

- [ ] Contrôler le colis (choc, humidité) et prendre des photos.
- [ ] Vérifier la référence reçue = la référence attendue.
- [ ] Relever le **nouvel ESN / S/N** (`display esn`).
- [ ] Tester sommairement (boot, quelques ports) avant mise en production.
- [ ] Si **DOA** (défectueux à réception) : photo + ticket immédiat, ne pas
      renvoyer l'ancien avant d'avoir une solution (cas vécu n°106).

## 98. Étape 7 — Retour du matériel défectueux

- Utilisez l'**emballage de la pièce de remplacement** (prévu pour).
- Joignez : n° RMA **lisible** sur le colis et dans le colis, fiche de
  description de la panne.
- Expédiez via le transporteur / l'adresse indiqués (ne pas improviser).
- Conservez la **preuve de dépôt** jusqu'à confirmation de réception par Huawei.
- **Délai de retour** : en général sous 10–15 jours (**à vérifier sur vos
  instructions RMA**). Au-delà : facturation de la pièce non retournée possible.

## 99. Étape 8 — Transfert de licence et restauration de configuration

Sur l'équipement de remplacement :

1. [ ] Restaurer la **configuration** (sauvegarde — vous en avez une, n'est-ce pas ?).
2. [ ] Transférer les **licences** vers le nouvel ESN (section 32).
3. [ ] Vérifier `display license` : fonctions + dates OK.
4. [ ] Mettre à jour le registre (ancien ESN → « retourné », nouvel ESN → actif).
5. [ ] Mettre à jour l'enregistrement sur le portail (nouveau S/N sur le site).
6. [ ] Supervision : vérifier que l'équipement remonte correctement.

## 100. RMA et sécurité des données

- Les équipements réseau contiennent votre **configuration** (mots de passe,
  clés VPN, certificats). Avant retour : faites un **reset usine** si
  l'équipement démarre encore (`reset saved-configuration`, puis reboot —
  commande exacte **à vérifier sur le portail officiel** pour votre modèle).
- Si l'équipement ne démarre plus : signalez-le au TAC ; demandez la
  procédure de **non-retour pour raison de sécurité** si votre politique
  l'exige (certains contrats prévoient la conservation/destruction locale
  du support — **à vérifier dans votre contrat**).
- Ne laissez jamais un disque/SSD contenant des données partir sans traçabilité.

---

# H. CAS VÉCUS COMMENTÉS

> Les cas ci-dessous sont des **scénarios types reconstitués** à partir de
> situations classiques du support réseau d'entreprise. Les noms, dates et
> valeurs sont fictifs. Chacun se termine par la leçon opérationnelle.

## 101. Cas n°1 — Licence UTM expirée un vendredi soir

**Situation.** Vendredi 18h42 : l'USG6000 du siège ne fait plus d'inspection
IPS. Les logs indiquent `The license file has expired`. La licence 1 an,
achetée avec l'équipement, expirait ce jour. Personne n'avait vu l'alerte
J-30 (partie dans les spams de la boîte générique).

**Gestion.** S2 ouvert au TAC le vendredi soir ; le partenaire, injoignable
avant lundi, ne peut générer le `.dat` de renouvellement que le lundi matin.
Le TAC confirme qu'aucune licence temporaire ne peut être délivrée sans
l'aval commercial. Le site passe le week-end sans inspection UTM (firewall
de base maintenu, navigation restreinte par consigne).

**Leçon.** 1) Les alertes d'expiration doivent partir vers une **boîte
surveillée** + SMS d'astreinte, pas une boîte générique. 2) Le script de
contrôle nocturne (`display license`, section 27) aurait alerté à J-90.
3) Envisager le renouvellement **3 ans** pour lisser le risque.

## 102. Cas n°2 — RMA urgent avant un audit

**Situation.** À J-5 d'un audit de sécurité, le second USG du cluster HA
rend l'âme (panne d'alimentation). Le site n'est plus redondé. L'auditeur
va forcément le relever.

**Gestion.** Ticket S2 (le site fonctionne sur le membre restant : pas de S1,
mais l'impact potentiel justifie S2). Diagnostic TAC en 3h, panne matérielle
confirmée, n° RMA généré à 14h10 → contrat Standard 9x5xNBD → pièce reçue
le lendemain 10h. Remplacement, restauration de config, transfert de licence
vers le nouvel ESN (section 32) : HA rétabli à J-3. L'audit se passe bien,
avec le rapport RMA en pièce justificative.

**Leçon.** 1) La règle des **15h00** (section 85) : un RMA généré à 14h10
change tout vs 15h10. 2) Garder le **dossier RMA complet** : c'est une preuve
de maîtrise pour l'auditeur. 3) Sur un cluster HA, chaque membre doit avoir
ses licences (section 41) — vérifié ici avant l'audit, ouf.

## 103. Cas n°3 — Le TAC demande des logs introuvables

