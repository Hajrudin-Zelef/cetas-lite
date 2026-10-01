---
id: collect-261001-general-networking/general-networking/attaque-supply-chain-axios-npm-2026-rat-nord-coreen-3
title: "attaque-supply-chain-axios-npm-2026-rat-nord-coreen"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Google", "Microsoft"]
dates: []
keywords: ["agent", "aws", "containment", "cyber", "open source"]
source: docs/RAG/collect-261001-general-networking/attaque-supply-chain-axios-npm-2026-rat-nord-coreen.md
source_anchor: ""
source_lines: [74, 126]
sha256: 8949849bfcd26ec0f3b80214fc8a153bb4d1050a29d8b5d3d75daf18daf3df59
---

# attaque-supply-chain-axios-npm-2026-rat-nord-coreen

Pour les équipes de sécurité et les développeurs potentiellement affectés par l’attaque supply chain Axios npm, voici les mesures de remédiation recommandées par Huntress, Microsoft et le SANS Institute :

**Étape 1 – Détection :** Vérifier si les versions [email protected], [email protected] ou [email protected] apparaissent dans vos fichiers package-lock.json, yarn.lock ou pnpm-lock.yaml entre le 31 mars 00h21 UTC et 03h29 UTC. Examiner les logs npm pour toute installation durant cette fenêtre.

**Étape 2 – Containment :** Bloquer immédiatement tout trafic sortant vers le domaine sfrclak[.]com, l’adresse IP 142.11.206.73 et le port 8000 au niveau du pare-feu périmétrique. Surveiller les logs réseau pour les connexions sortantes suspectes, les comportements de beaconing et les requêtes HTTP POST anomales.

**Étape 3 – Rotation des secrets :** Tout système ayant exécuté le payload malveillant doit être considéré comme **entièrement compromis**. Rotation immédiate de tous les credentials accessibles : tokens npm, clés AWS, clés SSH, secrets CI/CD, valeurs .env, credentials cloud, tokens OAuth et clés API.

**Étape 4 – Reconstruction :** Les systèmes infectés doivent être reconstruits à partir de zéro. Sur Windows, vérifier spécifiquement les mécanismes de persistance au redémarrage mis en place par le RAT.

**Étape 5 – Prévention future :** Désactiver les scripts postinstall npm (`--ignore-scripts`), utiliser des lockfiles épinglés (versions exactes plutôt que caret ^), activer les « cooldowns » npm et implémenter des outils de Software Composition Analysis (SCA) comme Socket, Snyk ou ArmorCode dans vos pipelines CI/CD.

## Tableau des indicateurs de compromission (IOCs) de l’attaque Axios

| Type d’IOC | Valeur | Description | 
|---|---|---|
| Domaine C2 | sfrclak[.]com | Serveur de commande et contrôle principal | 
| Adresse IP C2 | 142.11.206.73 | Hébergé chez Hostwinds LLC (AS54290) | 
| Port C2 | 8000 | Port de communication du RAT | 
| Chemin de campagne | /6202033 | Inversé = date de l’attaque (3-30-2026) | 
| Paquet npm malveillant | [email protected] | Typosquat de crypto-js avec payload RAT | 
| Version Axios compromise | [email protected] | Branche latest, publiée 31/03 00:21 UTC | 
| Version Axios compromise | [email protected] | Branche legacy, publiée 31/03 01:00 UTC | 
| User-Agent anomale | IE8/Windows XP | Identique sur Windows, macOS et Linux | 
| Registrar du domaine C2 | Namecheap Inc | Domaine enregistré le 30/03 16:03 UTC | 
| VPN de l’attaquant | Nœud AstrillVPN | Lié à UNC1069 (Google TIG) | 

## Le contexte géopolitique : la Corée du Nord intensifie ses attaques sur l’open source

L’attaque supply chain Axios npm s’inscrit dans une stratégie nord-coréenne de plus en plus agressive ciblant les chaînes d’approvisionnement logicielles pour financer le régime de Pyongyang. Selon les estimations des Nations Unies, les cyberattaques nord-coréennes ont généré plus de **3 milliards de dollars** entre 2017 et 2024, principalement via le vol de cryptomonnaies et d’actifs numériques. En 2025-2026, cette stratégie s’est élargie pour inclure le ciblage systématique de l’écosystème open source.

La campagne TeamPCP, dont l’attaque Axios représente le point culminant, a compromis au moins cinq projets open source majeurs en mars 2026 : Trivy (outil de sécurité conteneurs d’Aqua Security, 19 mars), KICS (outil d’analyse Infrastructure as Code de Checkmarx, 23 mars), LiteLLM (bibliothèque d’abstraction LLM, 24 mars), Telnyx (SDK de communication, 27 mars) et finalement Axios (31 mars). Cette escalade progressive – ciblant d’abord des outils de niche avant de s’attaquer à une bibliothèque universelle – témoigne d’une approche méthodique de montée en puissance.

« Ce que nous observons est un changement de doctrine. Les attaquants nord-coréens ne se contentent plus de cibler les échanges crypto – ils ciblent les développeurs eux-mêmes comme vecteur d’accès initial aux entreprises », a analysé **Adam Meyers**, vice-président senior de l’intelligence chez CrowdStrike, dans un briefing du 2 avril 2026. Cette analyse est corroborée par le rapport de cybersécurité 2026 qui identifie les attaques supply chain comme l’une des menaces les plus critiques de l’année.

## Implications pour la sécurité de la chaîne d’approvisionnement logicielle en Europe

L’attaque supply chain Axios npm relance le débat sur la sécurité des dépendances open source en Europe, alors que le **Cyber Resilience Act (CRA)** de l’Union européenne entre progressivement en vigueur. Ce règlement, adopté en 2024, impose aux fabricants de produits contenant des éléments numériques de garantir la sécurité de leurs chaînes d’approvisionnement logicielles – y compris les dépendances open source. Les entreprises ont jusqu’à 2027 pour se conformer pleinement aux exigences du CRA, mais l’attaque Axios met en évidence l’urgence d’une mise en conformité accélérée.

En France, l’**ANSSI** (Agence nationale de la sécurité des systèmes d’information) a renforcé ses recommandations sur la gestion des dépendances open source dans sa mise à jour de mars 2026, préconisant l’utilisation systématique de SBOM (Software Bill of Materials) et l’audit régulier des composants tiers. Le secteur des ESN (Entreprises de Services du Numérique), qui représente **34,8 milliards d’euros** en France en 2026, est particulièrement concerné car ces entreprises développent et maintiennent des applications critiques pour des clients dans les secteurs de la banque, de la santé et des infrastructures.

L’attaque intervient également dans un contexte de montée en puissance des initiatives de souveraineté numérique européenne. La compromission de la Commission européenne par ShinyHunters en mars 2026 (350 Go de données volées sur AWS) et l’attaque Axios npm démontrent que la dépendance de l’Europe aux infrastructures et bibliothèques développées en dehors de ses frontières constitue un risque systémique. Le plan de l’UE de 200 milliards d’euros pour les gigafactories IA devra nécessairement inclure un volet robuste de sécurité de la supply chain logicielle.

## 5 prédictions pour l’avenir de la sécurité npm après l’attaque Axios

**1. npm imposera l’authentification multi-facteurs obligatoire pour tous les mainteneurs de paquets à fort impact d’ici fin 2026.** GitHub a déjà rendu le 2FA obligatoire pour les contributeurs actifs en 2024. L’attaque Axios devrait accélérer l’extension de cette exigence à tous les mainteneurs de paquets dépassant un seuil de téléchargements hebdomadaires (probablement 1 million).

**2. Les « cooldowns » npm deviendront la norme par défaut d’ici Q3 2026.** Le mécanisme de cooldown, qui impose un délai entre la publication d’une nouvelle version et sa disponibilité pour les installations automatiques, aurait pu limiter significativement l’impact de l’attaque Axios. npm devrait l’activer par défaut pour tous les paquets à fort impact.

**3. Le marché des outils de Software Composition Analysis (SCA) dépassera 5 milliards de dollars en 2027.** L’attaque Axios npm va accélérer l’adoption d’outils comme Socket, Snyk, ArmorCode et Sonatype dans les entreprises européennes, particulièrement dans le contexte de la conformité au Cyber Resilience Act.

**4. Au moins deux autres attaques supply chain npm d’ampleur comparable se produiront en 2026.** La campagne TeamPCP montre une escalade méthodique. Les acteurs étatiques ont démontré leur capacité à compromettre les comptes mainteneurs et publier des versions malveillantes – d’autres bibliothèques populaires (Express, Lodash, React) pourraient être ciblées.

