---
id: collect-261001-general-networking/general-networking/bug-bounty-trouver-sa-premiere-faille-en-13-etapes-4
title: "Installer Go si nécessaire"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["apache", "mai", "valuation"]
source: docs/RAG/collect-261001-general-networking/bug-bounty-trouver-sa-premiere-faille-en-13-etapes.md
source_anchor: ""
source_lines: [159, 219]
sha256: 2be0197be72aeb00afd1872b1729e377c24c538a344b32fe35d19db23587e47e
---

# Installer Go si nécessaire

Il est recommandé de tester la reproduction au moins deux fois, à quelques minutes d’intervalle, en partant d’un état propre (nouvelle session, cache navigateur vidé) pour éliminer les faux positifs liés à un état de session résiduel. C’est également le moment d’évaluer honnêtement l’impact réel : une vulnérabilité technique sans impact métier démontrable (accès à une information déjà publique, par exemple) sera généralement reclassée en sévérité faible ou refusée, même si elle est techniquement valide.

## Étape 11 : Rédiger un rapport qui sera accepté

La qualité de la rédaction pèse autant que la qualité technique de la découverte dans la décision finale d’une équipe de triage. Un rapport efficace suit une structure constante : un titre précis indiquant le type de vulnérabilité et l’endpoint concerné, un résumé de l’impact métier en une ou deux phrases compréhensibles par un non-technicien, les étapes de reproduction numérotées avec les requêtes exactes, une preuve de concept (capture d’écran ou vidéo courte), et une évaluation de sévérité justifiée, idéalement selon le référentiel CVSS.

Il faut éviter deux écueils fréquents chez les débutants : exagérer l’impact pour obtenir une prime plus élevée (les équipes de triage expérimentées détectent rapidement cette pratique et elle nuit durablement à la réputation du compte), et au contraire noyer le rapport dans des détails techniques non pertinents qui masquent l’essentiel. Un rapport clair et concis, même sur une vulnérabilité de sévérité moyenne, obtient généralement une réponse plus rapide et un traitement plus favorable qu’un rapport confus sur une faille critique.

## Étape 12 : Gérer les retours et les délais de traitement

Après soumission, plusieurs issues sont possibles : validation directe, demande d’informations complémentaires, reclassement de sévérité, marquage comme doublon d’un rapport existant, ou fermeture pour absence d’impact suffisant. Les statistiques par programme publiées sur les plateformes indiquent généralement un délai moyen de première réponse et un délai moyen de résolution, deux indicateurs à consulter avant même de choisir son programme initial. Un programme dont le délai de première réponse dépasse plusieurs semaines décourage généralement les chercheurs les plus actifs, ce qui peut paradoxalement réduire la concurrence sur ce programme précis.

En cas de désaccord sur la sévérité ou le montant de la prime proposée, la plupart des plateformes prévoient un mécanisme de médiation. Il est préférable de l’utiliser avec mesure et arguments techniques précis plutôt que de multiplier les relances, qui peuvent être perçues négativement par l’équipe de triage lors de l’examen des futurs rapports du même chercheur.

## Étape 13 : Passer à l’échelle et se spécialiser

Une fois les premiers rapports validés, la progression naturelle consiste à se spécialiser sur une classe de vulnérabilités ou un type de technologie (applications mobiles, API GraphQL, infrastructures cloud) plutôt qu’à multiplier les programmes généralistes. Cette spécialisation permet d’accéder plus rapidement aux programmes privés sur invitation, généralement mieux rémunérés et moins saturés de chercheurs. Elle justifie aussi l’investissement dans Burp Suite Professional, dont le scanner actif automatisé et l’absence de limitation de débit sur Intruder deviennent rentables dès que le volume de tests augmente.

Construire un historique public de découvertes (avec l’accord des entreprises concernées, via les pages de remerciements ou “Hall of Fame”) reste également le levier le plus efficace pour obtenir des invitations à des programmes privés, bien plus qu’un nombre élevé de certifications.

## Table des récompenses moyennes par sévérité en 2026

Les montants varient fortement d’un programme à l’autre, mais les données indicatives compilées à partir de plusieurs plateformes en 2026 donnent un ordre de grandeur utile pour calibrer ses attentes.

| Sévérité | Récompense moyenne indicative | Exemple de vulnérabilité | 
|---|---|---|
| Critique | ≈ 4 200 $ | Exécution de code à distance, prise de contrôle de compte admin | 
| Élevée | ≈ 2 150 $ | IDOR avec accès à des données sensibles d’autres utilisateurs | 
| Moyenne | ≈ 750 $ | XSS stocké, contournement partiel d’autorisation | 
| Faible | ≈ 250 $ | Absence d’en-tête de sécurité avec impact limité | 

Ces chiffres doivent être pris comme des repères, pas comme des garanties. Certains programmes majeurs révisent régulièrement leurs barèmes à la baisse, comme l’a documenté The Register en mai 2026 concernant l’Internet Bug Bounty de HackerOne, dont la récompense critique est passée de 9 250 à 2 257 dollars. Toujours vérifier la table de récompenses spécifique au programme ciblé avant de soumettre un rapport.

## Erreurs courantes des débutants en bug bounty

Certaines erreurs reviennent systématiquement et expliquent une large part des rapports jugés non valides sur les plateformes majeures.

- **Ignorer le scope** : tester des sous-domaines ou des fonctionnalités explicitement exclus, ce qui entraîne un rejet immédiat et parfois une sanction sur le compte.
- **Soumettre des doublons** : ne pas vérifier si une vulnérabilité similaire a déjà été signalée, alors que la plupart des programmes publient un historique des rapports déjà traités.
- **Scanner trop agressivement** : dépasser la limite de requêtes par seconde tolérée, ce qui peut être interprété comme une attaque par déni de service.
- **Exagérer la sévérité** : présenter une faille mineure comme critique sans preuve d’impact métier réel, ce qui nuit à la crédibilité du compte sur le long terme.
- **Négliger la reproduction** : soumettre un rapport sans avoir vérifié que la faille est reproductible de façon fiable et constante.

## Exemple de sortie attendue lors d’une reconnaissance réussie

Voici à quoi ressemble une sortie typique de reconnaissance une fois les commandes de l’étape 4 exécutées sur un domaine réel, avec les identifiants remplacés par un exemple générique.

```
$ cat live_hosts.txt
https://app.exemple-cible.com [200] [Tableau de bord] [React,Nginx]
https://staging.exemple-cible.com [401] [Non autorisé] [Apache,PHP]
https://api.exemple-cible.com [200] [API Gateway] [Kong]
https://old-portal.exemple-cible.com [200] [Ancien portail 2022] [jQuery 1.9,IIS]
$ nuclei -l live_hosts.txt -severity high,critical -rate-limit 10
[2026-09-22 10:14:02] [exposed-config] [high] https://old-portal.exemple-cible.com/config.old.php
[2026-09-22 10:14:19] [outdated-jquery] [medium] https://old-portal.exemple-cible.com
```
Dans cet exemple type, le sous-domaine `old-portal`, visiblement un ancien portail non maintenu, concentre les résultats les plus intéressants, ce qui illustre la règle générale : les actifs oubliés valent presque toujours mieux que les applications principales, plus souvent auditées et corrigées en priorité par les équipes internes.

## Guide de dépannage : problèmes courants et solutions

