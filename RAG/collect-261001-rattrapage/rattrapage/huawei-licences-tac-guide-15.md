---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-15
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident", "license"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [2254, 2411]
sha256: dabb8833610558d73be45398f1dfabd37e6b6e985eb2a3d0ec39e1b38835ad21
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

- [ ] Sauvegarde **automatique hebdomadaire** de tous les équipements
      (TFTP/SFTP vers un serveur + copie hors site).
- [ ] Sauvegarde **manuelle avant/après chaque changement** de configuration.
- [ ] Nommage : `config_<équipement>_AAAAMMJJ-HHMM.cfg`.
- [ ] Test de restauration **une fois par an** sur un équipement de test ou
      un spare froid.
- [ ] Les sauvegardes contiennent des secrets (mots de passe, clés) :
      stockage **chiffré**, accès restreint.

## 131. Bulletins de sécurité : s'abonner et réagir

1. Sur le portail support, abonnez-vous aux **bulletins de sécurité**
   pour vos gammes (USG, S310, AR...).
2. À chaque bulletin critique : évaluer l'exposition de votre parc
   (versions concernées ? fonction vulnérable activée ?).
3. Vérifier **immédiatement** vos droits de téléchargement du patch
   (cas vécu n°105).
4. Planifier le patch : d'abord en maquette / sur un site pilote, puis
   généralisation avec fenêtre de maintenance.
5. Tracer : bulletin → décision → patch appliqué (dossier d'audit, section 138).

## 132. Fin de vie (EoS/EoL) : anticiper, pas subir

- **EoS (End of Sale)** : on ne peut plus acheter l'équipement → les
  extensions de parc doivent changer de modèle.
- **EoL (End of Life)** : fin du support logiciel, puis matériel → tout
  RMA devient impossible à terme.
- Surveillez les annonces de fin de vie sur le portail pour vos modèles.
- Règle de pilotage : **ne jamais démarrer un nouveau site** sur un modèle
  annoncé en EoS ; planifier le renouvellement des modèles en EoL à 3 ans.

## 133. Qui paie quoi : imputation interne des coûts

Pour éviter les débats budgétaires en pleine urgence, faites valider en
amont :

| Coût | Imputation suggérée |
|---|---|
| Renouvellements de licences UTM | Budget récurrent du service (OPEX) |
| Contrats Hi-Care | Budget récurrent du service (OPEX) |
| RMA sous contrat | Coût nul (couvert) — tracer le temps passé |
| Achat suite à dommage exclu | Budget exceptionnel + plan d'action correctif |
| Spares froids | Investissement du service (amorti sur les incidents évités) |
| Formation de l'équipe | Budget formation |

## 134. Retour d'expérience : le débrief post-incident (template)

À faire sous 5 jours ouvrés après tout incident S1/S2 :

```text
REX INCIDENT — [titre] — [date]
1. Chronologie factuelle (heures + actions) :
2. Cause racine (technique) :
3. Cause racine (organisationnelle) : [ex. alerte non vue, doc manquante]
4. Ce qui a bien fonctionné :
5. Ce qui a mal fonctionné :
6. Actions correctives (qui / quoi / échéance) :
   - [ ] ...
   - [ ] ...
7. Faut-il mettre à jour ce guide / les checklists ? (oui/non + quoi)
Participants : [...]
```

Un REX sans action corrective datée et nommée est une réunion pour rien.

## 135. Former l'équipe : plan de montée en compétence administrative

| Public | Contenu | Format |
|---|---|---|
| Tous les techniciens | Fiche réflexe 2h du matin (127), collecte de logs (76–78), ouverture d'un ticket (70) | 1h d'atelier + exercice « à blanc » |
| Référent licences | Parties B–C, registre, script de contrôle (124) | Tutorat + ce guide |
| Chef de service / adjoint | Parties D–G, négociation contrats, REX (134), budget | Ce guide + revue annuelle |
| Nouvel arrivant | Lecture des parties A–B + visite du dossier parc (129) | Kit d'accueil, semaine 1 |

## 136. Lexique anglais-français : phrases types pour le TAC

| Situation | Phrase type (anglais) |
|---|---|
| Ouverture | "We are experiencing [symptom] on [model] since [time]." |
| Sévérité | "Business impact: [X users] affected, no workaround in place. Requesting S[S1/S2/S3/S4]." |
| Envoi de logs | "Please find attached the diagnostic-information output collected during the issue." |
| Demande de contournement | "Could you please provide a workaround while root cause analysis is ongoing?" |
| Escalade | "No progress since [date]. Requesting senior engineer review." (section 75) |
| RMA | "Hardware failure confirmed. Please proceed with RMA under contract [number]." |
| Licence temporaire | "Is it possible to issue a temporary license to cover the gap until [date]?" |
| Clôture | "Issue resolved since [date/time], stable for [X] hours. Ticket can be closed." |

## 137. Les 10 erreurs les plus fréquentes dans les tickets

1. Pas de version VRP exacte (« la dernière » ne veut rien dire).
2. Pas d'ESN/S/N → le TAC ne peut pas vérifier les droits.
3. Logs collectés **après** le reboot qui a tout effacé.
4. Captures d'écran au lieu de texte copié-collé (non cherchable).
5. Sévérité gonflée « pour aller plus vite » (section 108).
6. Ticket ouvert par email sans n° de contrat → 24h de perdues.
7. Deux tickets en parallèle pour le même problème (canaux différents).
8. Description = interprétation (« le DHCP est cassé ») au lieu du symptôme.
9. Contact injoignable (téléphone du bureau un dimanche).
10. Ticket fermé dès le contournement, sans cause racine → récidive.

## 138. Le dossier d'audit : ce que l'auditeur demandera

Un auditeur (sécurité, qualité, financier) vérifiera typiquement :

- [ ] L'inventaire des équipements (S/N, modèles, versions).
- [ ] Les licences : droits valides pour les fonctions utilisées ?
- [ ] Les contrats de support : couverture réelle vs criticité ?
- [ ] Les tickets S1/S2 de l'année : délais respectés ? REX faits ?
- [ ] Les RMA : traçabilité complète (n° RMA, retours) ?
- [ ] Les sauvegardes de configuration : existent-elles ? testées ?
- [ ] Les patchs de sécurité : appliqués dans quels délais ?
- [ ] Les accès au portail : nominatifs ? revus après les départs ?

Si chaque section de ce guide est appliquée, ce dossier **existe déjà**.

## 139. Plan de continuité : le classeur « si je suis absent »

Le chef de service part en congés : son adjoint doit pouvoir gérer une
urgence administrative. Le classeur (physique ou partagé) contient :

1. Ce guide (imprimé ou PDF à jour).
2. Les contacts (section 117) **vérifiés il y a moins de 3 mois**.
3. Les n° de contrats Hi-Care et les échéances (calendrier section 34).
4. Les accès : où sont les comptes portail (pas les mots de passe : la
   procédure de réinitialisation), qui a les droits.
5. Le registre des licences + l'emplacement des `.dat`.
6. La fiche réflexe 2h du matin (section 127).
7. La procédure RMA résumée en une page (partie G).

Testez-le : une fois par an, l'adjoint gère **seul** un exercice.

## 140. Index des commandes CLI citées dans le guide

| Commande | Usage | Section |
|---|---|---|
| `display esn` | Récupérer l'ESN pour les licences | 4, 16 |
| `display version` | Version VRP pour les tickets | 69 |
| `display license` | Vérifier fonctions et dates d'expiration | 26–27 |
| `display device` | S/N et état matériel | 69, 92 |
| `display diagnostic-information` | Pièce maîtresse pour le TAC | 76 |
| `display logbuffer` | Logs en mémoire (volatils) | 77 |
| `display logfile` | Logs persistants | 77 |
| `display current-configuration` | Config active (sauvegarde, diagnostic) | 76, 130 |
| `display environment` | Température, alimentations, ventilateurs | 92 |
| `license active <fichier>` | Activer un fichier de licence | 24 |
| `dir` | Vérifier la présence du `.dat` en flash | 23 |
| `tftp <srv> get <fichier>` | Transférer le `.dat` | 23 |

*(La syntaxe exacte peut varier selon les modèles et versions VRP —
toujours confirmer avec la documentation de votre version.)*

## 141. FAQ express — les questions qu'on pose tout le temps

**Q1. J'ai perdu le fichier `.dat`, que faire ?**
R : S'il a déjà été activé, l'équipement continue de fonctionner (la licence
survit au reboot). Régénérez un fichier depuis le portail avec le même ESN
si vous en avez besoin (réinstallation, RMA). Archivez-le cette fois (section 22).

