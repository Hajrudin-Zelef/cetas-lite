---
id: collect-261001-general-networking/general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026-6
title: "Installe une autorité de certification locale de confiance"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026.md
source_anchor: ""
source_lines: [395, 435]
sha256: 1145c4af3f6f16fb2f546120ea4b2ff311ed0faebc448826b0be971f2d6d30d9
---

# Installe une autorité de certification locale de confiance

- **“SecurityError: The relying party ID is not a registrable domain suffix”** : le rpID déclaré côté serveur ne correspond pas exactement au domaine affiché dans la barre d’adresse. Vérifiez qu’il n’inclut ni protocole, ni port, ni sous-domaine incorrect.
- **“Origin does not match” lors de la vérification** : expectedOrigin doit reprendre le protocole complet (https://) et le port exact si celui-ci n’est pas standard. Une simple différence entre 3000 et 443 suffit à faire échouer la vérification.
- **WebAuthn refuse de se déclencher en HTTP** : en dehors de localhost strict, une origine sécurisée HTTPS est obligatoire. Utilisez mkcert en développement.
- **Un passkey créé sur un appareil n’apparaît pas sur un autre** : vérifiez que le compte cloud (iCloud, Google, Microsoft) est bien connecté et synchronisé sur les deux appareils, et que la synchronisation du trousseau est activée dans les réglages système.
- **“NotAllowedError: The operation either timed out or was not allowed”** : l’utilisateur a annulé la demande biométrique, le délai par défaut a expiré, ou l’appel n’a pas été déclenché directement par une interaction utilisateur (clic, tap).
- **L’autofill conditionnel n’apparaît jamais dans le champ de connexion** : vérifiez la présence de l’attribut autocomplete=”username webauthn” sur le champ, et confirmez via browserSupportsWebAuthnAutofill() que le navigateur teste supporte réellement cette fonctionnalité.
- **Erreur CORS entre le frontend et le backend** : si les deux tournent sur des ports ou domaines différents en développement, configurez credentials: ‘include’ côté fetch et les en-têtes CORS appropriés côté serveur.
- **“Unexpected registration response challenge”** : le challenge stocké en session a expiré ou a été écrasé par une requête concurrente. Vérifiez la durée de vie de votre session et évitez de partager le même challenge entre deux onglets ouverts simultanément.

## Passkeys vs clés de sécurité physiques vs TOTP : que choisir ?

Ces trois méthodes ne s’excluent pas forcément. Beaucoup d’organisations les combinent selon le profil de l’utilisateur et la criticité du compte protégé, plutôt que d’en imposer une seule à l’ensemble des effectifs.

La confusion la plus fréquente concerne la résistance au phishing d’une application TOTP. Le code à six chiffres généré par Google Authenticator ou une application équivalente reste, sur le papier, un facteur “que vous possédez”. Le problème, c’est qu’il n’est lié à aucune origine web. Un attaquant qui présente un faux formulaire de connexion peut relayer ce code en temps réel vers le vrai site, en profitant de la fenêtre de validité de quelques dizaines de secondes. C’est exactement le type d’attaque que documentent régulièrement les rapports d’incidents sur le phishing avancé. Un passkey échappe structurellement à ce scénario, puisque la signature produite par l’authentificateur est mathématiquement liée au domaine réel et ne peut pas être rejouée ailleurs, même par un relais en temps réel.

| Critère | Passkey (plateforme) | Clé de sécurité physique | Application TOTP | 
|---|---|---|---|
| Résistance au phishing | Élevée | Élevée | Faible à moyenne | 
| Coût | Gratuit (intégré à l’appareil) | Achat matériel requis | Gratuit | 
| Portabilité entre appareils | Selon synchronisation | Totale, tant que la clé est sur soi | Manuelle (export/QR code) | 
| Résilience si appareil perdu | Bonne si synchronisé dans le cloud | Nécessite une clé de secours | Nécessite les codes de récupération | 
| Cas d’usage recommandé | Grand public, usage quotidien multi-appareils | Comptes à très haut privilège, administrateurs systèmes | Services ne supportant pas encore WebAuthn | 

Pour un compte administrateur ou un accès à de l’infrastructure critique, une architecture Zero Trust combine souvent un passkey lié à l’appareil avec une clé de sécurité physique dédiée aux opérations les plus sensibles, en plus d’une vérification contextuelle (localisation, appareil connu, horaires habituels).

## Conseils avancés pour un déploiement en production

Une fois le flux de base validé, quelques ajustements séparent une démonstration d’un système prêt pour de vrais utilisateurs.

Stockez le challenge côté serveur, jamais dans un cookie visible côté client, et donnez-lui une durée de vie courte, de l’ordre de deux à cinq minutes. Journalisez chaque tentative d’enregistrement et de connexion avec l’horodatage, l’adresse IP et le résultat, ces journaux deviennent la première source d’investigation en cas d’incident. Proposez systématiquement un second passkey ou une clé de secours dès l’inscription plutôt que d’attendre que l’utilisateur en fasse la demande après avoir déjà perdu l’accès à son compte.

Ajoutez une limite de fréquence (rate limiting) sur les routes d’options d’enregistrement et de connexion. Rien n’empêche techniquement un script d’appeler `/login/options` en boucle pour un nom d’utilisateur donné, même si l’attaque ne peut pas aboutir sans l’authentificateur physique de la victime. Limiter ces routes à quelques appels par minute et par adresse IP évite malgré tout de transformer votre serveur en cible facile pour du bruit inutile ou une tentative d’énumération de comptes existants.

Surveillez également le taux d’échec de vérification dans le temps. Un pic soudain d’échecs sur `/register/verify` ou `/login/verify` signale souvent un problème de configuration après un déploiement, comme un rpID ou un origin mal mis à jour, plutôt qu’une attaque. Un tableau de bord simple qui trace ce taux d’échec heure par heure suffit à détecter ce genre de régression avant qu’elle n’atteigne tous vos utilisateurs.

Pensez aussi à la migration progressive plutôt qu’au basculement brutal. Affichez le passkey comme option recommandée à côté du mot de passe existant pendant plusieurs semaines, mesurez le taux d’adoption réel, puis seulement ensuite programmez la suppression du mot de passe pour les comptes qui ont confirmé leur passkey. Cette approche progressive réduit considérablement le volume de tickets de support au moment du changement. Enfin, documentez clairement pour votre équipe support les trois ou quatre scénarios de perte d’accès les plus probables (appareil perdu, changement d’écosystème, désinstallation accidentelle du gestionnaire de mots de passe) avec, pour chacun, la procédure exacte à suivre. Un support qui improvise face à un utilisateur bloqué crée plus de friction que le mot de passe qu’il était censé remplacer.

## Passkeys, RGPD et conformité en France

Du point de vue du RGPD, un passkey traite en réalité moins de données personnelles qu’un système de mot de passe classique. Aucun secret partagé n’est stocké côté serveur, seule une clé publique inutilisable en dehors de son contexte cryptographique. En choisissant **attestationType: ‘none’** comme recommandé plus haut, vous évitez même de collecter des informations sur le matériel exact de l’utilisateur, ce qui simplifie l’analyse d’impact relative à la protection des données pour ce traitement spécifique.

