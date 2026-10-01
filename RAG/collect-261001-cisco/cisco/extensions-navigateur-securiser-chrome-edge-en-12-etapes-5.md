---
id: collect-261001-cisco/cisco/extensions-navigateur-securiser-chrome-edge-en-12-etapes-5
title: "Dans chrome://extensions, pour chaque extension :"
domain: cisco
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/extensions-navigateur-securiser-chrome-edge-en-12-etapes.md
source_anchor: ""
source_lines: [259, 305]
sha256: 72b05bba163d6249e1b71c0cc0808de0ef3e6c4a0f77d466ee91ec97a6f0c996
---

# Dans chrome://extensions, pour chaque extension :

Ce cas concret illustre pourquoi la vérification de l’identité de l’éditeur ne peut pas être un contrôle réalisé une seule fois à l’installation. C’est un point de contrôle qui doit être répété à chaque audit périodique, en comparant le nom de l’éditeur actuel avec celui noté lors de l’installation initiale. Les outils de gestion centralisée comme LayerX automatisent cette comparaison à l’échelle d’un parc entier de postes, en envoyant une alerte dès qu’un changement d’éditeur est détecté sur une extension déjà déployée, ce qui réduit la fenêtre d’exposition de plusieurs semaines à quelques heures.

## Extensions de protection à considérer pour renforcer son navigateur

Paradoxalement, certaines extensions permettent justement de renforcer la sécurité du navigateur plutôt que de l’affaiblir, à condition de les choisir avec le même niveau d’exigence que celui appliqué au reste de l’inventaire. Les bloqueurs de traceurs et de scripts malveillants filtrent une partie du trafic publicitaire avant même qu’il n’atteigne la page affichée, ce qui réduit la surface d’attaque exposée aux campagnes de publicité malveillante (malvertising). Ces extensions doivent elles-mêmes provenir d’éditeurs identifiables et faire l’objet de mises à jour régulières, le même paradoxe de confiance s’appliquant à un outil de sécurité qu’à n’importe quel autre module.

Les gestionnaires de mots de passe intégrés à des solutions reconnues constituent une autre catégorie d’extension à privilégier plutôt qu’à bannir : ils réduisent le risque de réutilisation de mots de passe faibles, un facteur de compromission bien plus fréquent que celui d’une extension malveillante isolée. La condition reste la même : vérifier que l’éditeur est établi, que le code fait l’objet d’audits de sécurité publiés, et que l’historique de mises à jour est régulier et documenté.

Enfin, des extensions dédiées à la vérification de la réputation des sites visités, qui croisent l’URL affichée avec des bases de données de sites frauduleux connus, complètent utilement les protections déjà intégrées aux navigateurs modernes. Elles ne remplacent pas une vigilance de base face au phishing, mais ajoutent une couche de vérification automatisée qui a fait ses preuves dans les entreprises appliquant une politique de sécurité en profondeur, où plusieurs mécanismes de protection se chevauchent délibérément.

## Sources et références officielles

Pour suivre l’actualité des vulnérabilités touchant les navigateurs, consultez régulièrement les avis de sécurité du CERT-FR, qui publie chaque semaine les failles critiques identifiées sur les logiciels les plus répandus. Le blog sécurité de Malwarebytes détaille régulièrement les correctifs Chrome à appliquer en priorité. Pour la gestion des extensions, les pages d’aide officielles restent la référence la plus fiable : Google Chrome, Microsoft Edge et Mozilla Firefox. Enfin, pour le contexte européen sur la cybersécurité liée à l’IA, l’ENISA publie des lignes directrices régulièrement mises à jour.

## Foire aux questions

**Faut-il désinstaller toutes ses extensions par précaution ?**

Non, ce n’est ni nécessaire ni réaliste. L’objectif est de conserver les extensions utiles en limitant leurs permissions au strict nécessaire, et de supprimer celles qui présentent un risque disproportionné par rapport à leur usage réel.

**Les extensions du Chrome Web Store sont-elles automatiquement sûres ?**

Non. Le passage par le magasin officiel réduit le risque par rapport au sideloading, mais n’élimine pas la possibilité qu’une extension légitime soit rachetée plus tard et modifiée pour intégrer du code malveillant lors d’une mise à jour.

**Quelle est la différence entre “Au clic” et “Sur des sites spécifiques” ?**

« Au clic » exige une action explicite de l’utilisateur à chaque utilisation de l’extension sur une page. « Sur des sites spécifiques » donne un accès permanent, mais limité à une liste de domaines définie manuellement. Le premier réglage est plus restrictif et donc plus sûr par défaut.

**Les extensions IA sont-elles plus dangereuses que les extensions classiques ?**

Selon les données disponibles en 2026, les extensions dopées à l’IA présentent un taux de vulnérabilité supérieur de 60 % et un risque d’accès à des données sensibles trois fois plus élevé que les extensions classiques, principalement à cause des permissions étendues nécessaires à leur fonctionnement.

**À quelle fréquence faut-il refaire cet audit ?**

Tous les trois mois pour un usage personnel suffit dans la plupart des cas. Pour une organisation soumise à des exigences de conformité comme NIS2, l’audit doit être intégré au cycle de gestion des vulnérabilités existant, généralement mensuel ou trimestriel selon la criticité des postes.

**Faut-il désactiver complètement le mode développeur même pour un usage occasionnel ?**

Si vous devez tester une extension non publiée sur le magasin officiel, activez le mode développeur temporairement, effectuez le test, puis désactivez-le immédiatement après. Le laisser actif en permanence multiplie inutilement la fenêtre d’exposition au sideloading.

**Un profil de navigateur sans extension est-il vraiment plus sûr pour les opérations bancaires ?**

Oui, car il élimine mécaniquement le risque qu’une extension compromise intercepte les données de session sur ce profil précis, les cookies et le stockage local n’étant pas partagés entre profils.

**Les outils comme ExtensionPedia ou CRXcavator sont-ils gratuits ?**

ExtensionPedia propose une consultation gratuite pour vérifier une extension individuelle. CRXcavator et LayerX ciblent davantage un usage professionnel à l’échelle d’un parc de postes, avec des offres adaptées à la gestion centralisée en entreprise.
