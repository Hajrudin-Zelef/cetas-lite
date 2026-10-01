---
id: collect-261001-general-networking/general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026-5
title: "Installe une autorité de certification locale de confiance"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google", "Microsoft"]
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026.md
source_anchor: ""
source_lines: [348, 394]
sha256: 0f136d385b9551965d142a789c8ca8a53693794fd147bc5ab09aa5bbda8e6860
---

# Installe une autorité de certification locale de confiance

```
app.delete('/devices/:credentialID', requireAuth, (req, res) => {
  const user = getUser(req.session.username);
  const before = user.devices.length;
  user.devices = user.devices.filter(
    d => d.credentialID !== req.params.credentialID
  );
  if (user.devices.length === before) {
    return res.status(404).json({ error: 'Passkey introuvable' });
  }
  res.json({ removed: true, remaining: user.devices.length });
});
// Middleware minimal : exige une session déjà authentifiée
function requireAuth(req, res, next) {
  if (!req.session.username) {
    return res.status(401).json({ error: 'Non authentifié' });
  }
  next();
}
```
Affichez cette liste d’appareils dans un écran de paramètres, avec la date d’ajout et, si possible, un nom lisible par l’utilisateur (“iPhone de Denis”, “PC bureau”). C’est ce détail d’interface, plus que le code cryptographique lui-même, qui détermine si vos utilisateurs feront réellement confiance au système le jour où ils perdent un appareil.

Votre projet de démonstration est désormais complet : inscription, stockage, connexion, gestion du compteur anti-clonage, révocation et autofill conditionnel. Il reste minimal volontairement, pour que chaque brique reste lisible, mais il couvre exactement le flux qu’utilisent en production Google, Microsoft ou GitHub.

## Intégrer les passkeys avec Keycloak et Authelia en entreprise

Coder son propre backend WebAuthn a un intérêt pédagogique réel, mais une entreprise gère rarement l’authentification application par application. Elle centralise ce travail dans un fournisseur d’identité. Si vous avez suivi notre tutoriel Keycloak, vous savez déjà que cette plateforme open source gère nativement WebAuthn comme méthode de connexion, avec une politique configurable par royaume (realm) : passkey obligatoire, optionnel, ou réservé à certains groupes d’utilisateurs.

Authelia propose une approche complémentaire pour qui protège déjà un reverse proxy avec du 2FA auto-hébergé : elle accepte WebAuthn comme second facteur ou, selon la configuration, comme facteur unique suffisant. Dans les deux cas, le principe reste identique à celui que vous venez de coder à la main, sauf que le fournisseur d’identité absorbe toute la complexité de génération des challenges et de vérification des signatures. Le choix entre coder sa propre couche WebAuthn ou s’appuyer sur un IAM dépend surtout du nombre d’applications à couvrir. En dessous de trois ou quatre services internes, une implémentation directe comme celle de ce tutoriel reste raisonnable. Au-delà, un fournisseur d’identité centralisé évite de dupliquer la logique de sécurité partout.

Concrètement, la décision se résume souvent à trois questions. Combien d’applications internes doivent partager la même politique d’authentification ? Votre équipe a-t-elle déjà une brique SSO en place, ou faut-il la construire en même temps que le support passkey ? Et surtout, qui doit pouvoir forcer ou désactiver le passkey pour un groupe d’utilisateurs donné, un administrateur IAM ou chaque équipe applicative séparément ? Une startup avec une poignée de services internes gagne du temps avec l’implémentation directe présentée dans ce tutoriel. Une organisation soumise à NIS2, avec des dizaines d’applications et des exigences d’audit, tire davantage de valeur d’un IAM centralisé qui journalise chaque authentification à un seul endroit.

## 5 erreurs courantes à éviter avec les passkeys

La plupart des implémentations WebAuthn ratées ne viennent pas d’une faille cryptographique, le protocole est solide par construction, mais de décisions de conception prises trop vite. Voici les erreurs qui reviennent le plus souvent dans les retours d’expérience d’équipes ayant déployé les passkeys en production.

- **Ne prévoir aucune méthode de récupération.** Un utilisateur qui perd son seul appareil et n’a synchronisé aucun passkey ailleurs se retrouve bloqué. Prévoyez toujours un second facteur alternatif ou des codes de secours à usage unique, générés et affichés une seule fois au moment de l’inscription.
- **Confondre passkey synchronisable et passkey lié à l’appareil dans l’interface.** Les utilisateurs doivent comprendre si leur passkey survit au remplacement de leur téléphone ou non, sans quoi la confusion génère des tickets de support inutiles. Un simple badge “synchronisé” ou “lié à cet appareil” dans l’écran de gestion des passkeys suffit à lever l’ambiguïté.
- **Fixer attestationType sur ‘direct’ sans raison précise.** Cette option complexifie l’implémentation, ralentit l’inscription et soulève des questions de confidentialité, pour un bénéfice réel que seuls certains secteurs réglementés exigent, comme la finance ou l’accès à des systèmes gouvernementaux.
- **Changer de domaine sans anticiper l’impact.** Le rpID est gravé dans chaque passkey créé. Migrer de app.exemple.fr vers exemple.fr casse tous les passkeys existants et force une réinscription générale. Si un changement de domaine est prévu, planifiez une période de transition avec les deux méthodes actives en parallèle.
- **Retirer le mot de passe trop tôt.** Tant qu’une partie des utilisateurs n’a pas adopté les passkeys, coupez la corde de rappel progressivement, service par service, plutôt que d’un coup pour tout le monde.
- **Sous-dimensionner les tests multi-navigateurs.** Un flux validé uniquement sur Chrome desktop peut se comporter différemment sur Safari mobile, notamment sur la gestion de l’autofill conditionnel. Testez systématiquement sur au moins un navigateur mobile et un navigateur desktop avant chaque mise en production.

## Dépannage : 8 problèmes fréquents et leurs solutions

Voici les erreurs les plus signalées lors de l’implémentation de WebAuthn, avec la cause réelle derrière chaque message. Gardez cette liste ouverte dans un onglet pendant vos premiers tests, la plupart des blocages viennent d’une des huit causes ci-dessous.

