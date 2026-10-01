---
id: collect-261001-general-networking/general-networking/double-authentification-mfa-13-etapes-90-min-2026-4
title: "deploy-mfa-hardening.ps1"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-general-networking/double-authentification-mfa-13-etapes-90-min-2026.md
source_anchor: ""
source_lines: [220, 284]
sha256: 7f9d4816afbe406463fc6d367542c49bbce270d8e17d89611d23a3ff96282d67
---

# deploy-mfa-hardening.ps1

```
Étape 1/3 : audit des méthodes MFA enregistrées...
Rapport enregistré dans audit-mfa-2026.csv
Étape 2/3 : création de la politique d'accès conditionnel (mode audit)...
Politique créée en mode rapport uniquement.
Étape 3/3 : activation de FIDO2 comme méthode résistante au phishing...
Déploiement terminé. Vérifiez le portail Entra ID avant de passer en mode enforced.
```
Trois vérifications confirment ensuite que le déploiement s’est bien déroulé.

- La politique apparaît dans Entra ID, sous Protection puis Accès conditionnel, avec l’état attendu.
- Le rapport d’inscription MFA de l’étape 2 montre une progression du pourcentage d’utilisateurs inscrits chaque semaine.
- Les journaux de connexion affichent `authenticationRequirement: multiFactorAuthentication` pour les applications ciblées.

Si l’un de ces trois points ne correspond pas à ce qui est attendu, consultez la section dépannage plus bas avant de basculer la politique en mode `enabled`.

## Adoption du MFA en entreprise : les chiffres 2025-2026

Les chiffres suivants aident à situer une organisation par rapport au reste du marché, et à justifier si besoin un budget dédié auprès de la direction.

| Taille d’entreprise | Taux d’adoption MFA | Source | 
|---|---|---|
| Jusqu’à 25 salariés | 27 % | JumpCloud, 2025 | 
| 26 à 100 salariés | 34 % | JumpCloud, 2025 | 
| 1 001 à 10 000 salariés | 78 % | JumpCloud, 2025 | 
| Plus de 10 000 salariés | 87 % | JumpCloud, 2025 | 
| Secteur technologique (tous effectifs) | 87 % | JumpCloud, 2025 | 
| Moyenne mondiale, tous secteurs | 70 % (+4 points sur un an) | Okta Secure Sign-in Trends 2025 | 
| Authentification sans mot de passe résistante au phishing | 14,0 % (contre 8,6 % un an plus tôt) | Okta Secure Sign-in Trends 2025 | 

L’écart entre grandes et petites structures ne s’explique pas uniquement par le budget. Selon une enquête JumpCloud menée en 2024 auprès de plus de 1 000 professionnels IT de PME, 83 % des répondants imposaient déjà le MFA pour l’ensemble des ressources de l’entreprise, ce qui montre que la volonté existe mais que la mise en œuvre technique reste le principal frein. Côté applications grand public, la bascule progresse aussi : Descope indique que 94 % des organisations proposent une forme de MFA à leurs clients, mais seulement 10 % l’appliquent de façon uniforme sur l’ensemble de leurs applications. La fragmentation reste donc la norme, y compris chez les éditeurs de logiciels eux-mêmes.

## 5 erreurs courantes qui sabotent votre déploiement MFA

Ces pièges reviennent le plus souvent dans les retours d’expérience des équipes IT qui déploient le MFA à grande échelle.

- **Ne pas créer de compte break-glass avant d’activer une politique large.** Sans ce filet de sécurité, une politique mal configurée peut bloquer tous les administrateurs en même temps, y compris de leurs propres comptes.
- **Laisser le SMS comme méthode de secours par défaut.** Cela réintroduit exactement la faille que la migration vers FIDO2 ou Authenticator cherche à combler.
- **Déployer sans campagne de communication interne.** Une bascule surprise génère un pic de tickets et pousse certains utilisateurs à contourner les règles, par exemple en partageant leur téléphone avec un collègue.
- **Oublier les comptes de service et les applications héritées.** Ces comptes cassent silencieusement en production s’ils sont inclus dans une politique MFA sans exception explicite.
- **Ignorer le rapport d’inscription avant l’enforcement.** Basculer une politique en mode obligatoire sans vérifier qui n’est pas encore inscrit bloque les retardataires du jour au lendemain.
- **Ne pas activer le number matching.** Sans cette option, la porte reste ouverte aux attaques par lassitude MFA décrites plus haut.

## Dépannage : 9 problèmes fréquents et leurs solutions

Voici les incidents les plus signalés lors d’un déploiement MFA sur Microsoft Entra ID, avec la cause probable et la correction à appliquer.

- **Erreur « Insufficient privileges to complete the operation » lors de Connect-MgGraph.** Les scopes demandés à la connexion ne correspondent pas aux permissions consenties par un administrateur. Revérifiez la liste des scopes et redonnez le consentement admin si nécessaire.
- **Code AADSTS50076 ou AADSTS50079 à la connexion.** Ce n’est pas un bug : l’utilisateur doit simplement terminer son inscription MFA obligatoire avant de continuer.
- **Notification push jamais reçue.** Vérifiez les autorisations de notification en arrière-plan de l’application, la connectivité réseau du téléphone, et réenregistrez l’appareil si le problème persiste.
- **Clé FIDO2 non reconnue par le navigateur.** Le navigateur est probablement obsolète et ne supporte pas WebAuthn correctement, ou le pilote USB de la clé manque sur le poste.
- **Boucle de vérification MFA infinie.** Un cookie de session corrompu ou une horloge système désynchronisée en est souvent la cause, puisque le TOTP dépend d’une horloge précise à quelques secondes près.
- **Comptes de service bloqués par la politique.** Excluez-les explicitement de la politique d’accès conditionnel et migrez-les vers des identités managées quand c’est techniquement possible.
- **Rapport d’inscription vide ou incomplet.** Il s’agit généralement d’un délai de réplication côté Microsoft Graph. Réessayez après quelques heures ou vérifiez que le scope Reports.Read.All a bien été accordé.
- **Conflit entre plusieurs politiques d’accès conditionnel.** Utilisez l’outil de simulation « What If » d’Entra ID pour repérer les chevauchements avant qu’ils ne bloquent un utilisateur légitime.
- **Synchronisation Microsoft Entra Connect en retard.** Les nouveaux groupes de sécurité utilisés dans une politique n’apparaissent pas immédiatement : forcez un cycle de synchronisation delta pour accélérer la propagation.

## Conseils avancés : accès conditionnel basé sur le risque et sans mot de passe

Une fois le socle des 13 étapes en place, plusieurs options permettent d’aller plus loin. La fonctionnalité Authentication Strengths d’Entra ID permet d’exiger spécifiquement une méthode résistante au phishing, FIDO2 ou certificat, pour les applications les plus sensibles, tout en laissant Authenticator classique pour le reste du parc. Cette granularité évite d’imposer du matériel physique à l’ensemble de l’organisation alors que seuls quelques comptes en ont réellement besoin.

Identity Protection, disponible avec Entra ID P2, va plus loin que le simple accès conditionnel classique en calculant un score de risque par utilisateur et par connexion, basé sur des signaux comme les adresses IP anonymisées, les schémas de connexion inhabituels ou la présence de l’identifiant dans une fuite de données connue. Coupler ce score à une politique qui force automatiquement une réinscription MFA en cas de risque élevé referme rapidement la fenêtre d’exploitation d’un identifiant volé.

Enfin, pensez à faire remonter la télémétrie MFA (échecs répétés, changements de méthode, nouvelles inscriptions) vers votre SIEM plutôt que de la laisser dormir dans les journaux Entra ID. Une règle de détection simple, comme celle utilisée à l’étape 13, suffit souvent à repérer une tentative de contournement avant qu’elle ne se transforme en incident. À moyen terme, l’objectif réaliste pour la plupart des organisations reste une bascule complète vers le sans mot de passe : moins de tickets de réinitialisation, moins de surface d’attaque, et une expérience de connexion plus rapide pour les utilisateurs.

