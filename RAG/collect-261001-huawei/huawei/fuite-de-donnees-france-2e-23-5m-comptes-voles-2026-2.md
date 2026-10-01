---
id: collect-261001-huawei/huawei/fuite-de-donnees-france-2e-23-5m-comptes-voles-2026-2
title: "Aucune verification des droits d'acces : il suffit d'incrementer le numero"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "cost", "data breach", "incident"]
source: docs/RAG/collect-261001-huawei/fuite-de-donnees-france-2e-23-5m-comptes-voles-2026.md
source_anchor: ""
source_lines: [36, 75]
sha256: 91839101e6fcd14cea14284185d482d17a538d6c404d0652f51dc9e18638fa5c
---

# Aucune verification des droits d'acces : il suffit d'incrementer le numero

L’ANTS n’est pas un cas isolé. Vingt-quatre heures plus tôt, le **14 avril 2026**, le ministère de l’Éducation nationale confirmait une cyberattaque survenue dès la fin décembre 2025 via la plateforme **ÉduConnect**, le portail d’authentification des familles et des élèves. Près de **3,5 millions d’élèves mineurs** ont vu leurs données exposées. Là encore, le mécanisme repose sur une faille IDOR : l’attaquant, rattaché au groupe « DumpSec », aurait usurpé l’identité d’un agent administratif puis modifié un identifiant dans l’URL pour consulter les dossiers d’autres établissements.

Le secteur éducatif paie un lourd tribut en 2026. Le 24 mars, l’outil de gestion **Compas** exposait les données de **243 000 enseignants**. Au mois d’avril, c’est le **CROUS** qui voyait **774 000 dossiers étudiants** extraits de ses systèmes. Trois incidents en quelques semaines, tous concentrés sur l’écosystème scolaire et universitaire, qui manipule des données personnelles de mineurs particulièrement protégées par le RGPD.

Cette répétition n’est pas fortuite. Les systèmes d’information de l’éducation cumulent plusieurs vulnérabilités : ils sont anciens, fortement interconnectés, ouverts à des millions d’utilisateurs aux compétences techniques hétérogènes, et reposent souvent sur des prestataires multiples. Chaque maillon devient une porte d’entrée potentielle vers des bases nationales centralisées.

## Tchap : la messagerie chiffrée de l’État compromise en juin 2026

La séquence la plus récente concerne **Tchap**, la messagerie instantanée chiffrée développée par l’État pour ses agents. Le **7 juin 2026**, un compte utilisateur a été compromis à la suite d’une usurpation d’identité, un incident analysé en coordination avec l’ANSSI. Selon la communication de la DINUM, l’auteur revendique l’exposition de **73 467 agents** de l’État sur plus de 825 000 inscrits, soit moins de 9 % des utilisateurs – un volume que la direction interministérielle du numérique n’a pas validé.

Point crucial, et rassurant : il ne s’agit pas d’une faille technique du chiffrement. L’**historique des conversations privées chiffrées de bout en bout n’a pas été accessible**. Seuls les échanges des salons publics, couvrant une période allant de juin 2023 à juin 2026, ont pu être consultés. L’accès initial proviendrait d’un compte rattaché à l’environnement de l’Éducation nationale – un nouvel exemple de l’effet domino entre administrations interconnectées. Le Monde Informatique relève que des agents de ministères régaliens – Intérieur, Armées, Justice, Affaires étrangères, Économie et Finances – figureraient parmi les comptes potentiellement concernés.

L’incident Tchap illustre une vérité dérangeante : même les outils conçus pour la sécurité ne valent que par la robustesse de leurs identités. Une messagerie peut chiffrer parfaitement ses messages ; si un seul compte d’agent est usurpé, une partie de son contenu public s’expose. La gestion des accès et des identités redevient le nerf de la guerre, au même titre que la protection des mots de passe via un gestionnaire de mots de passe robuste.

## Chronologie 2026 : six mois de fuites administratives en cascade

Mises bout à bout, les fuites de 2026 dessinent une trajectoire d’escalade continue. Le tableau ci-dessous récapitule les incidents majeurs recensés en France depuis février, du secteur bancaire à la santé en passant par l’éducation et l’administration centrale.

| Date | Organisation | Personnes touchées | Type d’attaque | Auteur présumé | 
|---|---|---|---|---|
| 18 fév. 2026 | FICOBA / DGFiP | 1,2 million de comptes bancaires | Usurpation d’identifiants | Non identifié | 
| 26 fév. 2026 | Cegedim Santé | 15 millions de patients (164 000 données sensibles) | Exfiltration | Non identifié | 
| 24 mars 2026 | Compas (Éducation) | 243 000 enseignants | Exfiltration | Non identifié | 
| 14 avr. 2026 | ÉduConnect | 3,5 millions d’élèves mineurs | Faille IDOR | DumpSec | 
| 15-20 avr. 2026 | ANTS / France Titres | 11,7 M confirmés (18-19 M revendiqués) | Faille IDOR | breach3d | 
| Avril 2026 | CROUS | 774 000 étudiants | Extraction de données | Non identifié | 
| 7 juin 2026 | Tchap (messagerie de l’État) | 73 467 agents (revendiqués) | Compromission de compte | Non identifié | 

Deux constats s’imposent. D’abord, la **convergence des modes opératoires** : la faille IDOR frappe coup sur coup ÉduConnect et l’ANTS, prouvant que la même négligence de conception se réplique d’un portail public à l’autre. Ensuite, l’**ampleur des volumes** : en additionnant les seuls incidents administratifs, on dépasse aisément les 15 millions de citoyens directement concernés, soit plus d’un Français sur cinq.

## Pourquoi la France est devenue la cible numéro un

Comment expliquer que la France ait grimpé si haut dans le triste palmarès des **cyberattaques** ? Plusieurs facteurs se cumulent. Le premier est **géopolitique** : les positions internationales de Paris attirent régulièrement des groupes hacktivistes, notamment pro-russes, qui voient dans les administrations françaises une cible politiquement rentable.

Le deuxième facteur est paradoxal : c’est l’**avance de la France dans la numérisation de ses services publics**. France Connect, l’ANTS, ÉduConnect, Tchap, France Travail, l’Assurance maladie… le pays a concentré une part exceptionnelle de la vie administrative de ses citoyens dans des bases de données nationales. Cette centralisation, efficace pour l’usager, crée des « jackpots » de données : une seule faille suffit pour exposer des dizaines de millions de dossiers.

Le troisième facteur est la **fragilité du tissu économique**. Les PME, les ETI et les collectivités territoriales – souvent dépourvues d’équipes de sécurité dédiées – constituent une surface d’attaque immense et mal défendue. Les secteurs critiques (énergie, santé, transport, finance) y sont nombreux et fortement interdépendants. Enfin, la **chaîne de sous-traitance** multiplie les points d’entrée : un prestataire compromis ouvre la porte à des dizaines de clients, comme l’a montré l’attaque sur Cegedim Santé et ses 15 millions de patients. Le secteur de la santé n’est pas épargné outre-Atlantique non plus : selon le tracker d’Axis Intelligence mis à jour en août 2026, confirmé par le régulateur américain **HHS OCR** et par **HIPAA Journal**, la fuite touchant le prestataire **Conduent Business Services**, déclarée le **4 juin 2026**, a exposé **62 224 658 personnes**, ce qui en fait la troisième plus grande fuite de données de santé jamais recensée aux États-Unis. Cette exposition a un coût très concret : selon le rapport IBM *Cost of a Data Breach 2026*, une fuite de données coûte désormais en moyenne **11,5 millions de dollars** aux organisations américaines, en hausse de **11 %** sur un an, après un exercice 2025 déjà établi à 10,22 millions de dollars – soit environ 2,3 fois le coût moyen mondial de l’époque, alors mesuré à 4,44 millions de dollars.

## Le ransomware accélère en Europe : +55 % d’attaques en 2026

