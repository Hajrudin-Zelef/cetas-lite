---
id: collect-261001-general-networking/general-networking/attaque-supply-chain-axios-npm-2026-rat-nord-coreen-2
title: "attaque-supply-chain-axios-npm-2026-rat-nord-coreen"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["attribution", "open source"]
source: docs/RAG/collect-261001-general-networking/attaque-supply-chain-axios-npm-2026-rat-nord-coreen.md
source_anchor: ""
source_lines: [33, 73]
sha256: 90fb032e88c078137e548affa1b3af8407140c1771e469a11b25782fca62f2ff
---

# attaque-supply-chain-axios-npm-2026-rat-nord-coreen

« Il ne s’agit pas simplement d’une bibliothèque compromise – c’est un vecteur d’attaque qui touche l’ensemble de la chaîne de développement logiciel moderne », a expliqué **Johannes Ullrich**, doyen de la recherche au SANS Technology Institute, lors du briefing d’urgence organisé le 1er avril 2026. « Un seul compte mainteneur compromis, et un paquet qui touche des millions de systèmes devient une arme. » L’analyse du SANS a souligné que l’attaque était « exactement le type d’attaque supply chain discuté lors de la conférence RSAC la semaine précédente », illustrant le décalage entre la sensibilisation théorique et la capacité de défense réelle de l’écosystème.

## Tableau comparatif : les grandes attaques supply chain npm de l’histoire

| Attaque | Date | Téléchargements hebdomadaires | Durée d’exposition | Type de payload | Attribution | 
|---|---|---|---|---|---|
| **Axios (2026)** | 31 mars 2026 | 100+ millions | ~3 heures | RAT multiplateforme | Corée du Nord (Sapphire Sleet) | 
| event-stream (2018) | Novembre 2018 | ~2 millions | ~2 mois | Vol de Bitcoin | Non attribué | 
| ua-parser-js (2021) | Octobre 2021 | ~8 millions | ~4 heures | Cryptominer + vol credentials | Non attribué | 
| colors.js (2022) | Janvier 2022 | ~23 millions | ~3 jours | Sabotage (boucle infinie) | Mainteneur (protestation) | 
| SolarWinds Orion (2020) | Décembre 2020 | N/A (logiciel entreprise) | ~9 mois | Backdoor (SUNBURST) | Russie (APT29) | 
| Codecov (2021) | Janvier-avril 2021 | N/A (outil CI) | ~2 mois | Exfiltration credentials | Non attribué | 

L’attaque supply chain Axios npm se distingue par trois facteurs : l’**échelle sans précédent** (100+ millions de téléchargements hebdomadaires, soit 50 fois plus que event-stream), la **rapidité d’exécution** (de la publication à la première infection en 89 secondes), et le ciblage simultané de **deux branches majeures** (latest et legacy). C’est aussi la première attaque supply chain npm de cette ampleur formellement attribuée à un acteur étatique nord-coréen.

## La réponse de l’industrie : npm, Microsoft et Palo Alto mobilisés

La réponse à l’attaque supply chain Axios npm a été rapide mais a mis en lumière les limites structurelles de la sécurité de l’écosystème npm. L’équipe npm/GitHub a retiré les versions malveillantes en environ 3 heures, un temps de réaction considérablement meilleur que lors d’incidents précédents, mais suffisant pour que le payload atteigne des centaines de systèmes. Le scanner automatisé de **Socket** a détecté [email protected] comme malware en seulement 6 minutes après sa publication – une prouesse technique qui n’a malheureusement pas empêché la publication subséquente des versions Axios compromises.

Microsoft a publié un guide de remédiation détaillé le 1er avril 2026, recommandant le blocage immédiat de tout trafic sortant vers sfrclak[.]com et l’adresse IP 142.11.206.73. « Tout projet avec des versions Axios supérieures à `axios@^1.14.0` ou `axios@^0.30.0` se connectait automatiquement au C2 de Sapphire Sleet lors de l’installation et téléchargeait un malware de deuxième étape », a précisé l’équipe Microsoft Threat Intelligence.

Unit 42 de Palo Alto Networks a confirmé un « impact répandu » du RAT affectant tous les systèmes d’exploitation majeurs, publiant une liste complète d’indicateurs de compromission (IOCs) et recommandant une surveillance des logs réseau pour les connexions sortantes suspectes sur le port 8000, les comportements de beaconing et les requêtes HTTP POST anomales. **Lior Cohen**, directeur de la recherche sur les menaces chez Palo Alto Networks, a souligné : « La sophistication de cette attaque – du typosquatting à la persistance multiplateforme en passant par le C2 mimant le trafic npm – montre que les acteurs étatiques investissent massivement dans le ciblage des chaînes d’approvisionnement logicielles. »

## Le piratage du compte mainteneur : comment jasonsaayman a été compromis

L’attaque supply chain Axios npm a été rendue possible par le piratage du compte npm du mainteneur principal, identifié sous le pseudonyme **jasonsaayman**. Selon l’analyse de Malwarebytes et de Huntress, l’attaquant a probablement obtenu l’accès via un **token npm longue durée** combiné au mécanisme de « trusted publishing ». Les versions malveillantes ont été publiées manuellement après minuit UTC le 31 mars 2026, un horaire atypique qui coïncidait avec la nuit en Europe et le dimanche soir aux États-Unis – minimisant les chances de détection humaine rapide.

Malwarebytes a averti dans son analyse du 31 mars que l’attaquant pourrait avoir obtenu « un accès au dépôt, des clés de signature, des clés API ou d’autres secrets pouvant être utilisés pour backdoorer de futures releases ou attaquer votre infrastructure ». Les permissions de l’attaquant dépassaient celles des autres collaborateurs du projet, ce qui a retardé la révocation de l’accès. Cette situation illustre un problème fondamental de l’écosystème npm : la sécurité de bibliothèques utilisées par des millions de développeurs repose souvent sur un **seul compte individuel**, protégé par des mécanismes d’authentification parfois insuffisants.

« Le problème n’est pas qu’un mainteneur ait été compromis – c’est que l’architecture même de npm permet qu’un seul point de défaillance ait un impact sur 100 millions d’installations hebdomadaires », a commenté **Ax Sharma**, chercheur en sécurité chez Sonatype, spécialiste de la sécurité des chaînes d’approvisionnement logicielles. Cette analyse rejoint les préoccupations croissantes de la communauté open source concernant la « fatigue des mainteneurs » et le sous-investissement chronique dans la sécurité des projets critiques.

## Impact sur les entreprises européennes et les pipelines CI/CD

L’attaque supply chain Axios npm a des implications particulièrement préoccupantes pour les entreprises européennes. La fenêtre d’exposition de 3 heures (00h21-03h29 UTC) correspondait à la tranche horaire de 01h21 à 04h29 CET – en pleine nuit pour la plupart des équipes DevOps européennes. Cependant, de nombreux pipelines CI/CD s’exécutent en continu, et tout build déclenché durant cette période avec une dépendance vers Axios utilisant le versioning sémantique (caret ^) aurait automatiquement téléchargé et exécuté le malware.

Les entreprises européennes utilisant des environnements de développement cloud (GitHub Actions, GitLab CI, Jenkins, CircleCI) sont particulièrement vulnérables car les runners CI partagent souvent des secrets d’environnement – tokens de déploiement, clés API, credentials de bases de données – qui auraient été exfiltrés par le RAT. Le règlement européen **NIS2**, entré en vigueur en octobre 2024, impose aux entreprises de secteurs critiques de signaler les incidents de cybersécurité sous 24 heures, ce qui a poussé plusieurs organisations européennes à lancer des audits d’urgence de leurs dépendances npm dès le 31 mars.

Selon un rapport du cabinet Wavestone publié en février 2026, la surface d’attaque SaaS a été « multipliée par 3 en 2025 via les incidents de rebond chez les partenaires ». L’attaque Axios npm illustre parfaitement ce phénomène : une seule bibliothèque compromise dans la chaîne d’approvisionnement peut affecter des milliers d’applications en aval. Le marché numérique français, estimé à **74,3 milliards d’euros** en 2026 selon Numeum (avec une croissance de 4,3 %), est particulièrement exposé compte tenu de la dépendance massive de l’écosystème tech français à JavaScript et Node.js.

## Guide de remédiation : les étapes critiques pour les équipes de sécurité

