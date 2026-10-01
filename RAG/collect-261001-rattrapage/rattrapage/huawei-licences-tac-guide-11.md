---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-11
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [1661, 1808]
sha256: d03e7738db5fbf9862b048c1ac44939cc38c10398979fc6e33fd2bf705f32dc3
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

**Situation.** Ticket S3 sur des déconnexions Wi-Fi intermittentes (AP761).
Le TAC demande le `display diagnostic-information` **au moment du problème**
et les logs du contrôleur sur la plage horaire. L'équipe n'a que des logs
d'il y a 3 jours, le logbuffer a tourné depuis.

**Gestion.** Mise en place d'une collecte proactive : export syslog permanent
vers le serveur de logs, script qui fige le diagnostic-information dès
qu'une alerte remonte. Deuxième occurrence capturée « en flagrant délit » :
le TAC identifie un bug connu, patch fourni.

**Leçon.** 1) Sans logs **horodatés au moment des faits**, le TAC ne peut
rien faire : c'est parole contre silence. 2) La **centralisation syslog**
n'est pas une option, c'est un prérequis du support. 3) Documenter la
procédure de collecte (sections 76–78) et la tester **avant** la panne.

## 104. Cas n°4 — ESN illisible sur l'étiquette

**Situation.** Un S310 à remplacer : l'étiquette au dos est effacée (chaleur
+ poussière du local technique). Le `display esn` est inaccessible : le
switch ne boote plus. Impossible de générer le fichier de licence du
remplacement sans ESN... qui est celui de l'**ancien** équipement, pas du
nouveau — rappel utile : pour le nouvel équipement, c'est **son** ESN qu'il
faut (il boote, lui).

**Gestion.** En réalité le blocage venait d'ailleurs : le technicien cherchait
l'ESN de l'ancien pour « transférer » la licence, alors qu'il fallait
simplement relever l'ESN du **nouvel** équipement (`display esn` OK) et
demander la réémission du droit. Pour l'ancien ESN (traçabilité du retour),
le S/N partiel + la photo de l'étiquette + le bon de livraison d'origine
ont suffi au partenaire.

**Leçon.** 1) À la réception : **photographier toutes les étiquettes** et
archiver (checklist section 114). 2) Enregistrer les S/N sur le portail
dès la réception (section 49) : le portail devient la source de vérité quand
l'étiquette meurt. 3) Comprendre le sens du transfert (section 32) évite de
chercher le mauvais ESN.

## 105. Cas n°5 — Firmware non téléchargeable sans contrat

**Situation.** Faille critique publiée sur la version VRP des AR720. L'équipe
veut patcher. Sur le portail : le firmware cible est visible mais le
téléchargement est **refusé** — le contrat logiciel a expiré il y a 4 mois.

**Gestion.** Deux options : 1) renouveler le contrat logiciel (devis
partenaire, 48h), 2) contournement temporaire (ACL, désactivation du service
vulnérable) en attendant. L'équipe choisit le contournement documenté +
renouvellement en urgence. Le patch est appliqué à J+3.

**Leçon.** 1) Le droit de télécharger les firmwares est **contractuel**
(section 51), pas technique. 2) Les bulletins de sécurité doivent déclencher
une vérification **immédiate** des droits de téléchargement. 3) Maintenir les
contrats logiciels comme les licences : calendrier partagé (section 34).

## 106. Cas n°6 — Pièce DOA à la réception

**Situation.** RMA d'un S310 : la pièce de remplacement arrive, le technicien
la monte... et elle ne boote pas non plus (emballage éventré, choc visible
sur un coin).

**Gestion.** Photos immédiates du colis et de l'équipement, ticket rouvert
avec mention **DOA**, l'ancien équipement (toujours sur site) n'est **pas**
renvoyé. Le TAC valide le DOA en 2h, nouvelle pièce expédiée en NBD. Leçon
apprise : la vérification de la section 97 (tester avant de renvoyer
l'ancien) a évité de se retrouver sans aucun switch.

**Leçon.** 1) Ne renvoyez **jamais** le défectueux avant d'avoir testé le
remplacement. 2) Photographiez systématiquement les colis à réception.
3) Un DOA se traite en priorité : signalez-le dans l'heure.

## 107. Cas n°7 — RMA refusé : hors garantie et dommage exclu

**Situation.** Un AP761 ne s'allume plus après un orage. Ticket ouvert,
diagnostic TAC : traces de **surtension** sur l'alimentation (pas de
parafoudre sur le site), équipement hors garantie standard de toute façon.
RMA refusé : dommage exclu + pas de contrat.

**Gestion.** Devis de réparation proposé par Huawei (montant dissuasif) ;
décision : achat d'un AP neuf via le partenaire (délai 1 semaine) +
installation d'un **parafoudre** et vérification de la terre sur le site.
Le coût total (AP + parafoudre + déplacement) est imputé en « leçon ».

**Leçon.** 1) La garantie ne couvre pas les dommages externes (section 87) :
protégez les sites (parafoudre, terre < 5 ohms — voir le guide onduleurs).
2) Un équipement hors contrat = RMA impossible : c'est un achat, pas un échange.
3) Après un orage, inspectez **tout** le site, pas seulement l'équipement HS.

## 108. Cas n°8 — Sévérité S1 requalifiée en S3

**Situation.** Panne du lien principal d'une agence (AR720). Le technicien
d'astreinte ouvre un **S1** « pour aller plus vite ». Le TAC constate : le
lien 4G de secours fonctionne, l'agence travaille (dégradé). Requalification
en **S3**, réponse en 2h au lieu de 30 min. Le technicien s'énerve, le ton
monte.

**Gestion.** Le responsable (prévenu par l'astreinte) rappelle le cadre :
le S1 exige un **impact critique réel** (section 64). Le lien est rétabli
via le TAC en 4h en S3. En débrief : on ne « gonfle » pas la sévérité.

**Leçon.** 1) Une sévérité abusive **retarde** le traitement (requalification
+ friction) au lieu de l'accélérer. 2) Décrivez l'impact **factuellement**,
laissez le TAC qualifier. 3) Vos S1 doivent rester crédibles (section 68) :
c'est un capital.

## 109. Cas n°9 — Renouvellement Hi-Care oublié

**Situation.** Le contrat Hi-Care des USG expire un 31 décembre. L'oubli est
découvert le 15 janvier, lors d'une panne mineure. Le TAC ouvre quand même le
ticket (bonne volonté) mais : pas de SLA contractuel, pas de téléchargement
de patch, et le renouvellement « après expiration » exige une procédure de
remise en conformité avec inspection.

**Gestion.** Renouvellement en urgence (surcoût + délai), panne traitée en
« best effort ». Mise en place du calendrier partagé (section 34) avec
alertes J-90/J-60/J-30 **et** responsable nommé.

**Leçon.** 1) Un contrat expiré = des droits perdus immédiatement, même si
le TAC fait un geste commercial ponctuel. 2) Le renouvellement après
expiration coûte plus cher et prend plus de temps. 3) Nommez un
**responsable** des échéances, pas « l'équipe ».

## 110. Cas n°10 — Transfert de licence oublié après RMA

**Situation.** RMA d'un USG6000 parfaitement exécuté : pièce reçue en NBD,
config restaurée, tout fonctionne... sauf l'IPS, inactif. Personne n'a
transféré la licence vers le nouvel ESN (section 32). Découvert 3 semaines
plus tard lors d'un contrôle.

**Gestion.** Demande de transfert en urgence au partenaire (24h), activation
du nouveau `.dat`, vérification `display license`. Le site est resté 3
semaines sans inspection UTM sans que personne ne s'en aperçoive.

**Leçon.** 1) Ajoutez le transfert de licence à la **checklist RMA**
(étape 8, section 99) : c'est une case à cocher, pas une option. 2) Après tout
RMA : `display license` systématique. 3) La supervision doit alerter sur
l'état des licences, pas seulement sur le ping.

## 111. Cas n°11 — eKit : qui appeler quand ça ne répond plus ?

**Situation.** Un AP361 eKit ne diffuse plus. Le technicien appelle la
hotline Huawei : « pour la gamme eKit, merci de contacter votre partenaire
distributeur ». Il appelle le partenaire : « envoyez-nous le diagnostic par
email ». 48h de flottement, chacun renvoyant vers l'autre.

**Gestion.** Le chef de service clarifie le circuit (section 118) : pour
eKit, le **partenaire est le point d'entrée unique** (niveau 1), qui escalade
vers Huawei si nécessaire. Il fait inscrire ce circuit dans le contrat de
maintenance avec des **délais de réponse du partenaire** écrits noir sur blanc.

