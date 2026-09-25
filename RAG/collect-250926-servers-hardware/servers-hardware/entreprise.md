---
id: collect-250926-servers-hardware/servers-hardware/entreprise
title: "Entreprise"
domain: servers-hardware
role: reference
task: reference
actors: ["JFrog"]
dates: []
keywords: ["open source"]
source: docs/RAG/clean4/entreprise.md
source_anchor: ""
source_lines: [1, 76]
sha256: 46a60372f85bad29951e9772e8152a800e2750be0f22e5246194e8c95bf9d8ec
---

# Entreprise

Utiliser OpenCode en toute sÃ©curitÃ© dans votre organisation.

OpenCode Enterprise est destinÃ© aux organisations qui souhaitent sâassurer que leur code et leurs donnÃ©es ne quittent jamais leur infrastructure. Cela est possible grÃ¢ce Ã une configuration centralisÃ©e qui sâintÃ¨gre Ã votre SSO et Ã votre passerelle IA interne.

Pour dÃ©marrer avec OpenCode Enterprise :

1. Faites un essai en interne avec votre Ã©quipe.
2. **Contactez-nous** pour discuter des options de tarification et de mise en Åuvre.

OpenCode est open source et ne stocke aucune de vos donnÃ©es de code ou de contexte, vos dÃ©veloppeurs peuvent donc simplement commencer et effectuer un essai.

**OpenCode ne stocke pas votre code ni vos donnÃ©es contextuelles.** Tous les traitements sâeffectuent localement ou via des appels API directs Ã  votre fournisseur dâIA.

Cela signifie que tant que vous faites appel Ã un fournisseur de confiance ou Ã une passerelle IA, vous pouvez utiliser OpenCode en toute sÃ©curitÃ©.

La seule mise en garde ici concerne la fonctionnalitÃ© facultative `/share`.

Si un utilisateur active la fonctionnalitÃ© `/share`, la conversation et les donnÃ©es qui y sont associÃ©es sont envoyÃ©es au service que nous utilisons pour hÃ©berger ces pages de partage sur opencode.ai.

Les donnÃ©es sont actuellement servies via le rÃ©seau pÃ©riphÃ©rique de notre CDN et sont mises en cache en pÃ©riphÃ©rie, Ã proximitÃ© de vos utilisateurs.

Nous vous recommandons de dÃ©sactiver cette option pour votre essai.

**Vous possÃ©dez tout le code produit par OpenCode.** Il nây a aucune restriction de licence ni revendication de propriÃ©tÃ©.

Nous utilisons un modÃ¨le par siÃ¨ge pour OpenCode Enterprise. Si vous disposez de votre propre passerelle LLM, nous ne facturons pas les jetons utilisÃ©s. Pour plus de dÃ©tails sur les options de tarification et de mise en Åuvre, **contactez-nous**.

Une fois que vous avez terminÃ© votre essai et que vous Ãªtes prÃªt Ã  utiliser OpenCode Ã  votre organisation, vous pouvez **nous contacter** pour discuter des options de tarification et de mise en Åuvre.

Nous pouvons configurer OpenCode pour utiliser une seule configuration centrale pour lâensemble de votre organisation.

Cette configuration centralisÃ©e peut sâintÃ©grer Ã votre fournisseur SSO et garantit que tous les utilisateurs accÃ¨dent uniquement Ã votre passerelle IA interne.

GrÃ¢ce Ã la configuration centrale, OpenCode peut sâintÃ©grer au fournisseur SSO de votre organisation pour lâauthentification.

Cela permet Ã OpenCode dâobtenir les informations dâidentification de votre passerelle IA interne via votre systÃ¨me de gestion des identitÃ©s existant.

Avec la configuration centrale, OpenCode peut Ã©galement Ãªtre configurÃ© pour utiliser uniquement votre passerelle IA interne.

Vous pouvez Ã©galement dÃ©sactiver tous les autres fournisseurs dâIA, en vous assurant que toutes les demandes transitent par lâinfrastructure approuvÃ©e de votre organisation.

Bien que nous vous recommandons de dÃ©sactiver les pages de partage pour garantir que vos donnÃ©es ne quittent jamais votre organisation, nous pouvons Ã©galement vous aider Ã les auto-hÃ©berger sur votre infrastructure.

Ceci est actuellement sur notre feuille de route. Si vous Ãªtes intÃ©ressÃ©, **faites-le-nous savoir**.

## Quâest-ce que OpenCode Entreprise ?

OpenCode Enterprise est destinÃ© aux organisations qui souhaitent sâassurer que leur code et leurs donnÃ©es ne quittent jamais leur infrastructure. Il peut le faire en utilisant une configuration centralisÃ©e qui sâintÃ¨gre Ã votre SSO et Ã votre passerelle IA interne.

## Comment dÃ©marrer avec OpenCode Enterprise ?

Commencez simplement par un essai interne avec votre Ã©quipe. OpenCode par dÃ©faut ne stocke pas votre code ni vos donnÃ©es contextuelles, ce qui facilite le dÃ©marrage.

Ensuite, **contactez-nous** pour discuter des options de tarification et de mise en Åuvre.

## Comment fonctionne la tarification dâentreprise ?

Nous proposons des tarifs dâentreprise par siÃ¨ge. Si vous disposez de votre propre passerelle LLM, nous ne facturons pas les jetons utilisÃ©s. Pour plus de dÃ©tails, **contactez-nous** pour un devis personnalisÃ© basÃ© sur les besoins de votre organisation.

## Mes donnÃ©es sont-elles sÃ©curisÃ©es avec OpenCode Enterprise ?

Oui. OpenCode ne stocke pas votre code ni vos donnÃ©es contextuelles. Tout le traitement sâeffectue localement ou via des appels API directs Ã votre fournisseur dâIA. GrÃ¢ce Ã la configuration centrale et Ã lâintÃ©gration SSO, vos donnÃ©es restent sÃ©curisÃ©es au sein de lâinfrastructure de votre organisation.

## Pouvons-nous utiliser notre propre registre privÃ© NPM ?

OpenCode prend en charge les registres npm privÃ©s via la prise en charge native des fichiers `.npmrc` de Bun. Si votre organisation utilise un registre privÃ©, tel que JFrog Artifactory, Nexus ou similaire, assurez-vous que les dÃ©veloppeurs sont authentifiÃ©s avant dâexÃ©cuter OpenCode.

Pour configurer lâauthentification avec votre registre privÃ© :

Cela crÃ©e `~/.npmrc` avec les dÃ©tails dâauthentification. OpenCode le dÃ©tectera automatiquement.

Alternativement, vous pouvez configurer manuellement un fichier `.npmrc` :

Les dÃ©veloppeurs doivent Ãªtre connectÃ©s au registre privÃ© avant dâexÃ©cuter OpenCode pour garantir que les packages peuvent Ãªtre installÃ©s Ã partir du registre de votre entreprise.
