---
id: collect-261001-general-networking/general-networking/extensions-jetbrains-malveillantes-70-000-installs-2026-2
title: "Indicateur de compromission confirmé - campagne JetBrains Marketplace (juin 2026)"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Microsoft"]
dates: []
keywords: ["arr", "deepseek", "exploit", "incident", "valuation"]
source: docs/RAG/collect-261001-general-networking/extensions-jetbrains-malveillantes-70-000-installs-2026.md
source_anchor: ""
source_lines: [47, 103]
sha256: 9085617706128b8c7a0e148a0d6c6017aa9c25ba1a7c47a9dcd0b11eeae1664c
---

# Indicateur de compromission confirmé - campagne JetBrains Marketplace (juin 2026)

Face à ce type de campagne, la question naturelle est de savoir si une marketplace protège mieux ses utilisateurs qu’une autre. La réponse, sur la base des informations rendues publiques, reste nuancée : aucune des trois grandes plateformes d’extensions pour IDE ne documente publiquement un contrôle manuel systématique avant publication.

| Critère | JetBrains Marketplace | VS Code Marketplace | Open VSX Registry | 
|---|---|---|---|
| Modèle de contrôle | Signalement puis retrait a posteriori | Signalement (« Report a concern ») puis retrait | Non détaillé publiquement | 
| Désactivation à distance des plugins déjà installés | Oui, kill-switch confirmé en juin 2026 | Retrait du Marketplace, désactivation non systématique | Non détaillé publiquement | 
| Extensions malveillantes retirées (connu) | 15 (une campagne) | 110+ (année 2025) | Non communiqué publiquement | 
| Installations cumulées touchées (connu) | ~70 000 | Plusieurs centaines de millions cumulées, selon plusieurs études | Non communiqué publiquement | 

Microsoft a par exemple retiré plus de 110 extensions malveillantes du VS Code Marketplace au cours de la seule année 2025, un chiffre qui illustre le volume que la modération après coup doit absorber. JetBrains, de son côté, a démontré avec l’incident de juin 2026 qu’elle dispose d’un mécanisme de désactivation à distance capable de neutraliser des plugins déjà installés, une capacité que tous les gestionnaires d’extensions n’annoncent pas aussi clairement. Pour Open VSX, le registre ouvert utilisé notamment par des forks de VS Code, les politiques de vérification détaillées ne sont pas documentées publiquement de façon comparable, ce qui rend un comparatif chiffré difficile à établir avec rigueur.

## Pourquoi les assistants de code IA sont devenus une cible prioritaire

Les assistants IA pour le développement sont passés, en l’espace de deux ou trois ans, du statut de gadget testé par curiosité à celui d’outil que beaucoup de développeurs ouvrent plusieurs fois par jour. Cette adoption a un revers direct : chaque assistant IA repose sur une clé API, souvent reliée à un compte facturé à l’usage, parfois connectée à des dépôts de code privés ou à des environnements d’entreprise.

Pour un attaquant, imiter un outil IA plutôt qu’une extension généraliste présente un double avantage. La demande est forte, car les développeurs cherchent activement de nouveaux plugins IA et acceptent d’en tester plusieurs. La configuration elle-même y aide aussi : entrer une clé API dans un champ de paramètres est un geste routinier, presque automatique, que la campagne JetBrains a exploité sans avoir besoin de le modifier. Le fait que les extensions incriminées reprenaient des noms proches de marques réelles comme DeepSeek plutôt que d’inventer une identité de toutes pièces renforçait la confiance au moment de l’installation.

## L’impact pour les développeurs et les entreprises en France et en Europe

JetBrains n’a pas communiqué de répartition géographique des 70 000 installations concernées, et aucune entreprise française ou européenne n’a été identifiée publiquement à ce stade comme victime confirmée. L’exposition potentielle reste néanmoins réelle : les IDE JetBrains (IntelliJ IDEA, PyCharm, WebStorm, GoLand, Rider) comptent parmi les environnements de développement les plus utilisés dans les entreprises technologiques européennes, du secteur bancaire aux éditeurs de logiciels.

Pour une équipe technique, le risque ne se limite pas au coût direct d’une clé API réutilisée par un tiers. Une clé exposée peut ouvrir l’accès à des historiques de conversation avec l’assistant IA, susceptibles de contenir des extraits de code propriétaire collés pour obtenir de l’aide, ou servir de point d’entrée si l’organisation partage les mêmes schémas d’authentification entre plusieurs services. Les équipes de sécurité interne doivent désormais intégrer un nouvel élément à leurs audits : la provenance et la légitimité des extensions IA installées par les développeurs, souvent sans validation centralisée préalable.

## Ce qu’une clé API compromise permet réellement de faire

Le préjudice d’une fuite de clé API dépend largement de ce à quoi cette clé donne accès. Dans le cas de la campagne JetBrains, la clé volée correspond à un compte de fournisseur IA facturé à l’usage : un tiers peut donc consommer des crédits payants au nom de la victime, sans qu’aucune alerte ne se déclenche avant la réception de la facture ou l’épuisement du quota.

Le risque ne s’arrête pas toujours là. Selon le fournisseur IA concerné, une clé compromise peut aussi exposer un historique de requêtes, donc potentiellement des fragments de code envoyés pour analyse ou complétion. C’est précisément cette incertitude, difficile à quantifier sans audit complet, qui pousse JetBrains à recommander une rotation systématique des clés plutôt qu’une évaluation au cas par cas pour les 70 000 installations concernées.

## Comment savoir si vous êtes concerné

JetBrains a publié la liste des 15 extensions retirées ainsi que l’adresse IP du serveur de collecte dans son bulletin de sécurité du 16 juin 2026. Un développeur ou une équipe IT peut vérifier son exposition de plusieurs façons : consulter l’historique des extensions installées dans chaque IDE JetBrains, rechercher spécifiquement les plugins portant les noms DeepSeek AI Assist, DeepSeek Junit Test et DeepSeek Git Commit, et examiner les journaux réseau sortants à la recherche de connexions vers l’adresse 39.107.60.51.

Même en l’absence de trace d’installation active, JetBrains rappelle que ses systèmes ont désactivé les plugins à distance dès le 16 juin 2026. Une extension qui aurait disparu de l’interface sans action de l’utilisateur constitue donc, en elle-même, un indice qu’elle faisait partie de la liste incriminée.

## Les mesures à prendre immédiatement

Au-delà de la vérification, JetBrains et StepSecurity recommandent une série d’actions concrètes pour toute personne ou organisation ayant pu être exposée aux extensions JetBrains malveillantes :

- Régénérer immédiatement toute clé API configurée dans une extension IA tierce installée avant le 16 juin 2026
- Bloquer l’adresse IP 39.107.60.51 au niveau du pare-feu ou du DNS d’entreprise
- Vérifier les journaux de facturation des fournisseurs IA concernés à la recherche d’une activité inhabituelle
- Auditer la liste complète des extensions installées sur les postes de développement, IDE par IDE
- Limiter, quand c’est possible, les clés API à des périmètres et quotas restreints plutôt qu’à des accès larges

Voici l’indicateur de compromission principal publié par les chercheurs, utile pour un blocage rapide côté pare-feu ou proxy :

```
# Indicateur de compromission confirmé - campagne JetBrains Marketplace (juin 2026)
IP du serveur de collecte : 39.107.60.51
Hébergeur : Alibaba Cloud (Pékin, Chine)
Protocole d'exfiltration observé : HTTP non chiffré
Extensions confirmées : DeepSeek AI Assist, DeepSeek Junit Test, DeepSeek Git Commit (+12 autres, retirées)
Action recommandée : bloquer l'IP en sortie, régénérer toute clé API exposée
```
## Ce que cette affaire révèle sur la sécurité des outils IA en 2026

