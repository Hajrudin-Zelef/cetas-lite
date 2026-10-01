---
id: collect-261001-general-networking/general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026-4
title: "Installe une autorité de certification locale de confiance"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026.md
source_anchor: ""
source_lines: [185, 347]
sha256: ed30fa477e1596b9260f58ae511e85c9a005fa156f3790b7d56a4d059f7fc4d4
---

# Installe une autorité de certification locale de confiance

```
// public/register.js
import { startRegistration } from '@simplewebauthn/browser';
async function inscrirePasskey(username) {
  const optionsResponse = await fetch('/register/options', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username }),
  });
  const options = await optionsResponse.json();
  let attestation;
  try {
    attestation = await startRegistration({ optionsJSON: options });
  } catch (error) {
    console.error('Enregistrement annulé ou échoué :', error);
    return;
  }
  const verifyResponse = await fetch('/register/verify', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username, response: attestation }),
  });
  const { verified } = await verifyResponse.json();
  console.log(verified ? 'Passkey enregistré avec succès' : 'Échec de l’enregistrement');
}
```
Le serveur reçoit ensuite cette réponse et doit la vérifier avant de faire confiance au nouveau passkey. C’est l’étape où **@simplewebauthn/server** valide la signature, le challenge et l’origine.

```
app.post('/register/verify', async (req, res) => {
  const { username, response } = req.body;
  const user = getUser(username);
  try {
    const verification = await verifyRegistrationResponse({
      response,
      expectedChallenge: user.currentChallenge,
      expectedOrigin: origin,
      expectedRPID: rpID,
    });
    const { verified, registrationInfo } = verification;
    if (verified && registrationInfo) {
      user.devices.push({
        credentialID: registrationInfo.credential.id,
        credentialPublicKey: registrationInfo.credential.publicKey,
        counter: registrationInfo.credential.counter,
        transports: response.response.transports,
      });
    }
    res.json({ verified });
  } catch (error) {
    console.error(error);
    res.status(400).json({ error: error.message });
  }
});
```
## Étapes 11 et 12 : la connexion sans mot de passe

L’inscription est terminée. Reste à permettre à l’utilisateur de revenir se connecter sans jamais retaper de mot de passe. Le principe est symétrique : le serveur génère des options d’authentification, le navigateur les transmet à l’authentificateur via `navigator.credentials.get()`, et le serveur vérifie la signature retournée.

```
app.post('/login/options', async (req, res) => {
  const { username } = req.body;
  const user = getUser(username);
  const options = await generateAuthenticationOptions({
    rpID,
    allowCredentials: user.devices.map(d => ({
      id: d.credentialID,
      transports: d.transports,
    })),
    userVerification: 'preferred',
  });
  user.currentChallenge = options.challenge;
  res.json(options);
});
app.post('/login/verify', async (req, res) => {
  const { username, response } = req.body;
  const user = getUser(username);
  const device = user.devices.find(d => d.credentialID === response.id);
  if (!device) {
    return res.status(400).json({ error: 'Passkey inconnu pour cet utilisateur' });
  }
  try {
    const verification = await verifyAuthenticationResponse({
      response,
      expectedChallenge: user.currentChallenge,
      expectedOrigin: origin,
      expectedRPID: rpID,
      credential: {
        id: device.credentialID,
        publicKey: device.credentialPublicKey,
        counter: device.counter,
      },
    });
    if (verification.verified) {
      device.counter = verification.authenticationInfo.newCounter;
      req.session.username = username;
    }
    res.json({ verified: verification.verified });
  } catch (error) {
    console.error(error);
    res.status(400).json({ error: error.message });
  }
});
app.listen(3000, () => console.log('Serveur démarré sur https://localhost:3000'));
```
Le compteur **counter** mérite une explication. Chaque authentificateur incrémente en théorie un compteur à chaque signature. Si le serveur reçoit un jour un compteur inférieur à celui déjà enregistré, cela peut signaler un clonage de l’authentificateur, et la connexion doit être rejetée. Certains authentificateurs de plateforme renvoient systématiquement zéro : la bibliothèque SimpleWebAuthn gère ce cas particulier automatiquement, mais il faut le savoir avant de déboguer un faux problème.

Côté client, la fonction équivalente à l’inscription s’appelle **startAuthentication**.

```
// public/login.js
import { startAuthentication } from '@simplewebauthn/browser';
async function connexionPasskey(username) {
  const optionsResponse = await fetch('/login/options', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username }),
  });
  const options = await optionsResponse.json();
  const assertion = await startAuthentication({ optionsJSON: options });
  const verifyResponse = await fetch('/login/verify', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username, response: assertion }),
  });
  const { verified } = await verifyResponse.json();
  if (verified) {
    window.location.href = '/tableau-de-bord';
  }
}
```
## Étape 13 : tester en HTTPS local, déployer et activer l’autofill conditionnel

WebAuthn refuse de fonctionner sur une origine HTTP non sécurisée, à l’exception stricte de localhost. Dès que vous testez sur un autre nom d’hôte ou que vous préparez un déploiement, générez un certificat local avec mkcert.

```
# Installe une autorité de certification locale de confiance
mkcert -install
# Génère un certificat valable pour localhost et ses variantes
mkcert localhost 127.0.0.1 ::1
# Démarrez ensuite votre serveur en HTTPS avec les fichiers générés
# (adaptez server.js pour charger localhost-key.pem et localhost.pem)
node server-https.js
```
Une fois le certificat en place, un test manuel complet ressemble à ceci. Ouvrez la page d’inscription, saisissez un nom d’utilisateur, cliquez sur le bouton de création de passkey. Le navigateur affiche sa fenêtre native, vous demande votre empreinte ou votre code PIN, puis la console affiche la confirmation suivante :

```
$ curl -X POST https://localhost:3000/register/verify \
  -H "Content-Type: application/json" \
  -d '{"username":"denis","response":{"id":"...","rawId":"...","type":"public-key"}}'
{"verified":true}
```
Dernier raffinement, particulièrement apprécié des utilisateurs : l’autofill conditionnel. Il permet au navigateur de suggérer directement un passkey disponible dès que l’utilisateur clique dans le champ de connexion, sans bouton dédié. Deux conditions techniques suffisent : ajouter l’attribut `autocomplete="username webauthn"` sur le champ de saisie, et déclencher l’appel avec l’option **useBrowserAutofill**.

```
import { startAuthentication, browserSupportsWebAuthnAutofill } from '@simplewebauthn/browser';
if (await browserSupportsWebAuthnAutofill()) {
  startAuthentication({ optionsJSON: options, useBrowserAutofill: true })
    .then(handleAuthenticationResponse);
}
```
Une dernière brique manque pour que le projet soit vraiment utilisable : la possibilité, pour l’utilisateur, de révoquer un passkey perdu ou compromis sans devoir contacter le support. Ajoutez une route dédiée qui retire le credential de la liste des appareils enregistrés.

