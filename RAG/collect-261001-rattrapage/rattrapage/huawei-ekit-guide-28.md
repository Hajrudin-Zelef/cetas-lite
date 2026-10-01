---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-28
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [2468, 2504]
sha256: cab45db60bfe9a4ef8c34bc9db084c5a89a3beae0ca2aa4fd139c7f74e33ac39
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

1. « C'est bizarre, ça marchait chez moi. » → « Je reproduis le problème et je reviens vers vous avec un diagnostic. »
2. « C'est la faute de l'opérateur / du matériel. » → « J'ai identifié la cause, voici ce qu'on fait. » (même si c'est l'opérateur, c'est toi le responsable du dossier)
3. « Il faut tout changer. » → « Voici ce qu'on garde, voici ce qu'on remplace, et pourquoi. »
4. « Je ne sais pas. » (sec) → « Je ne sais pas encore, je vérifie et je vous réponds aujourd'hui. » (puis le faire vraiment)
5. « C'est compliqué à expliquer. » → explique simplement (si tu ne peux pas l'expliquer simplement, c'est que tu ne l'as pas compris — retourne à la maquette)
La confiance se gagne en disant la vérité vite, pas en ayant toujours raison.

## 228. Modèle de message post-installation (J+7)

```
Bonjour [Nom],
Une semaine après la mise en service de votre réseau, petit point :
- Supervision : tout est au vert depuis le [date], aucune alerte critique.
- Ce que nous avons ajusté : [ex. : puissance radio de l'AP accueil].
- Rappel : votre interlocuteur reste [nom] au [téléphone] (astreinte : [plages]).
- En pièce jointe : la fiche de votre installation + le PV signé.
- Prochaine étape : [audit trimestriel / proposition de contrat de maintenance].
Bonne journée,
[Ton nom] – [Ton service]
```
Un message comme ça, 7 jours après, vaut 10 campagnes de pub : le client se sent suivi, pas abandonné.

## 229. Signaler un bug ou demander une fonction : les bons canaux

1. **Distributeur Gold** : premier niveau — avec logs, captures, SN, versions, procédure de reproduction (un bug sans reproduction est un bug qui ne sera pas corrigé).
2. **Support Huawei** (via le distributeur ou support.huawei.com) : pour les bugs confirmés et les questions de compatibilité.
3. **Forums e.huawei.com** : pour les retours d'expérience et les cas tordus déjà rencontrés par d'autres intégrateurs.
4. **En interne** : note le bug dans le dossier du site + une entrée « contournement » dans ce guide — en attendant le correctif, c'est ton contournement qui fait tourner le site.
Ne laisse jamais un bug « en l'air » : soit il est signalé avec preuves, soit il est contourné et documenté.

## 230. Crédits, avertissement final et mot de la fin (bis)

- **Sources** : fiches techniques constructeur (AR180/AR280, AP361/AP761/AP266/AP673H, S220/S310/S620, USG6000F-S), communiqués Huawei (MWC 2025, Intelligent Office 2.0), portail ekit.huawei.com — vérifiées par recherche web en septembre 2026.
- **Avertissement** : la gamme eKit évolue vite ; toute section marquée « à vérifier sur la documentation officielle » est une action à faire avant chiffrage/déploiement, pas une clause de style.
- **Exemples** : adresses IP, noms de SSID, mots de passe et prix de ce guide sont **fictifs ou indicatifs** — ne les déploie jamais tels quels.
- **Responsabilité** : ce guide est un aide-mémoire de terrain, pas un substitut aux documentations constructeur ni aux normes locales. En cas de contradiction, la doc officielle et la réglementation gagnent toujours.
- Voilà, chef : 230 sections, de la visite technique à la fin de contrat, en passant par la panne de 3 h du matin. Le reste, c'est le terrain qui l'écrira — avec tes annotations en marge.
