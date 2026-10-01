---
id: collect-261001-ia-llm/ia-llm/plan-cybersecurite-ia-ue-2026-80-de-phishing-par-ia-2
title: "plan-cybersecurite-ia-ue-2026-80-de-phishing-par-ia"
domain: ia-llm
role: reference
task: reference
actors: ["CISA", "Google", "Microsoft"]
dates: []
keywords: ["cyber", "apache", "cybersecurity", "exploit", "mai"]
source: docs/RAG/collect-261001-ia-llm/plan-cybersecurite-ia-ue-2026-80-de-phishing-par-ia.md
source_anchor: ""
source_lines: [54, 107]
sha256: 39d6eabd1165eb2fe572f2bab9fd04fc2cf73d73b88a9cf04d84ad242579f165
---

# plan-cybersecurite-ia-ue-2026-80-de-phishing-par-ia

La directive NIS2 continue de fixer les obligations de cybersécurité et de signalement d’incidents pour les secteurs essentiels et importants. Notre guide technique sur la directive NIS2 détaille les étapes de mise en conformité pour les entreprises concernées. Le plan de Bruxelles ne modifie aucune de ces obligations. Il vise plutôt à accélérer la détection des vulnérabilités et la réponse aux incidents grâce à l’IA, ce qui devrait faciliter le respect des délais de signalement imposés par NIS2.

Le règlement DORA, propre au secteur financier, impose déjà une gestion structurée du risque numérique aux banques, assureurs et infrastructures de marché. Le Cyber Resilience Act, lui, encadre la sécurité des produits connectés commercialisés dans l’Union. Le plan d’action ne touche à aucun de ces deux textes. Il agit en complément, sur un terrain que ni NIS2, ni DORA, ni le CRA ne couvrent directement : l’usage de l’IA elle-même comme arme et comme bouclier.

## Ce qui ne change pas : aucune nouvelle obligation contraignante

Un point mérite d’être répété clairement : ce plan ne crée aucune nouvelle obligation légale. Aucune entreprise ne recevra d’amende pour ne pas avoir respecté une échéance du plan cybersécurité-IA, pour la simple raison qu’il ne s’agit pas d’un texte contraignant.

La vice-présidente exécutive de la Commission européenne pour la Souveraineté technologique, la Sécurité et la Démocratie, Henna Virkkunen, a résumé l’ambition politique du texte en une phrase.

« L’IA transforme la signification même de la cybersécurité. Nous devons suivre le rythme. » (« AI is transforming the meaning of cybersecurity. And we must keep pace. »)

Henna Virkkunen, vice-présidente exécutive de la Commission européenne pour la Souveraineté technologique, la Sécurité et la Démocratie (source)

Reste que l’écart entre cette ambition affichée et les moyens juridiques réellement mobilisés alimente déjà les critiques, détaillées plus loin dans cet article. Toutes les obligations contraignantes continuent de provenir des textes existants : le règlement sur l’IA, NIS2, DORA et le Cyber Resilience Act. Le plan agit en couche de coordination au-dessus de cet édifice, pas en remplacement. Pour les responsables conformité, suivre ce dossier relève de la veille stratégique, pas d’une nouvelle échéance juridique.

## La France en première ligne : ANSSI, CERT-FR et l’accord du 3 juillet

La France illustre bien comment ce plan européen s’articule avec l’action nationale.

Le 3 juillet 2026, quatre jours avant la publication du plan de Bruxelles, l’ANSSI, l’Autorité de contrôle prudentiel et de résolution (ACPR) et la Banque de France ont signé un accord de coordination cyber pour le secteur financier, en lien avec DORA et NIS2. L’accord porte les signatures de Vincent Strubel, directeur général de l’ANSSI, d’Emmanuelle Assouan, secrétaire générale de l’ACPR, et de Denis Beau, premier sous-gouverneur de la Banque de France.

Six jours plus tard, le 9 juillet, le CERT-FR a publié sept avis de sécurité en une seule journée, référencés CERTFR-2026-AVI-0849 à 0855, couvrant Wireshark, GitLab, Traefik, Juniper Networks, Palo Alto Networks, Microsoft Edge et Microsoft Azure Linux. L’avis sur Juniper Networks signale à lui seul plusieurs failles permettant exécution de code à distance, déni de service et fuite de données, dont les CVE-2026-33801 et CVE-2026-33799.

Le 8 juillet, cinq acteurs français de la santé numérique, Doctolib, Alan, Implicity, Lifen et Resilience Care, ont alerté le gouvernement sur le risque de voir la qualification SecNumCloud devenir obligatoire pour leur secteur, une contrainte technique et financière qu’ils jugent disproportionnée. Ce dossier reste, à l’heure où ces lignes sont écrites, en discussion à Bercy.

## Les alertes CERT-FR qui illustrent l’urgence

Le tableau ci-dessous reprend les avis de sécurité les plus significatifs publiés par le CERT-FR et par le bulletin de cyberveille du secteur santé début juillet 2026. Il donne une photographie concrète de ce que « répondre plus vite grâce à l’IA », l’un des objectifs du plan européen, signifierait en pratique pour les équipes de sécurité françaises.

| Référence | Produit concerné | Score CVSS | Faille exploitée | 
|---|---|---|---|
| CERTFR-2026-AVI-0855 | Microsoft Azure Linux | Non communiqué | Non | 
| CERTFR-2026-AVI-0854 | Microsoft Edge | Non communiqué | Non | 
| CERTFR-2026-AVI-0853 | Palo Alto Networks | Non communiqué | Non | 
| CERTFR-2026-AVI-0852 | Juniper Networks (Junos OS, cRPD) | Non communiqué | Non | 
| CERTFR-2026-AVI-0851 | Traefik (CVE-2026-54763) | 9,1 | Non | 
| CERTFR-2026-AVI-0850 | GitLab | Non communiqué | Non | 
| CERTFR-2026-AVI-0849 | Wireshark | Non communiqué | Non | 
| Bulletin santé numérique | Adobe (CVE-2026-48282) | 10,0 | Oui | 
| Bulletin santé numérique | Apache Camel Keycloak (CVE-2026-53913) | 9,8 | Non | 
| Bulletin santé numérique | WordPress (CVE-2026-5524) | 9,8 | Non | 
| Bulletin santé numérique | Google (CVE-2026-15129) | 9,6 | Non | 

Sur les onze avis recensés ici, un seul concerne une faille activement exploitée : celle touchant Adobe, notée 10 sur 10 à l’échelle CVSS, le score maximal possible. Les autres restent préventifs, mais leur volume, sept avis CERT-FR en une seule journée, donne une idée du rythme auquel les équipes de sécurité françaises doivent absorber l’information.

## Comparaison internationale : États-Unis, Royaume-Uni, Union européenne

L’Union européenne n’est pas seule à légiférer, ou à éviter de légiférer, sur l’IA appliquée à la cybersécurité. Washington et Londres ont emprunté des chemins différents.

Aux États-Unis, un décret présidentiel intitulé « Promoting Advanced Artificial Intelligence Innovation and Security », signé en juin 2026, charge la Cybersecurity and Infrastructure Security Agency (CISA) d’émettre des directives opérationnelles contraignantes et prévoit la création d’une plateforme fédérale de mutualisation sur la cybersécurité de l’IA. Le texte prévoit aussi un canal d’échange volontaire avec les développeurs de modèles de pointe et un processus classifié de test de leurs capacités offensives. Ce virage réglementaire s’accompagne d’un basculement déjà bien engagé du capital-risque local : selon J.P. Morgan (juin 2026), 72 % des opérations de capital-risque en cybersécurité aux États-Unis concernaient dès mai 2026 des start-up également classées IA, pour 4,6 milliards de dollars investis depuis le début de l’année. Il n’existe cependant toujours pas de loi fédérale unique sur l’IA aux États-Unis : la Californie, le Colorado, l’Illinois ou la ville de New York appliquent chacun leurs propres règles.

Le Royaume-Uni mise sur les codes de bonnes pratiques plutôt que sur la loi. Le National Cyber Security Centre (NCSC) pilote depuis janvier 2025 un Code of Practice for the Cyber Security of AI, fixant des standards minimaux pour les développeurs et opérateurs de systèmes d’IA. Le gouvernement britannique a par ailleurs mené une consultation publique sur la cybersécurité de l’IA, et depuis le 12 mai 2026, un nouveau règlement oblige l’Information Commissioner’s Office (ICO) à élaborer un code de pratique formel sur l’IA et les décisions automatisées.

