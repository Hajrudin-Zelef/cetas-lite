---
id: collect-261001-general-networking/general-networking/double-authentification-mfa-13-etapes-90-min-2026-1
title: "deploy-mfa-hardening.ps1"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["cyber"]
source: docs/RAG/collect-261001-general-networking/double-authentification-mfa-13-etapes-90-min-2026.md
source_anchor: ""
source_lines: [1, 32]
sha256: 8326e6aeb845e79bccf36f04592d931631f3975dc786dc7aeeca017986ae450e
---

# deploy-mfa-hardening.ps1

Seules 27 % des entreprises de moins de 25 salariés ont activé une forme de double authentification sur leurs comptes professionnels, contre 87 % des groupes de plus de 10 000 employés, selon les données 2025 de JumpCloud. Cet écart explique en grande partie pourquoi les petites structures restent la cible favorite des campagnes de phishing qui visent Microsoft 365 et les messageries professionnelles. Ce tutoriel détaille, étape par étape, comment activer une double authentification qui résiste vraiment aux techniques de contournement actuelles, avec les commandes PowerShell nécessaires, les pièges classiques à éviter et un script complet prêt à déployer sur votre tenant Microsoft Entra ID.

## Pourquoi la double authentification est votre première ligne de défense en 2026

Le mot de passe seul a cessé de protéger grand-chose depuis longtemps. Les campagnes de phishing automatisées, souvent vendues comme service clé en main sur des forums spécialisés, permettent à n’importe quel attaquant de récupérer des identifiants en quelques minutes à peine. La double authentification (MFA, pour multi-factor authentication) ajoute une couche de vérification qui bloque la majorité de ces tentatives, même quand le mot de passe a déjà fuité dans une base de données compromise.

En France, l’ANSSI a mis à jour sa stratégie nationale de cybersécurité en mars 2026, avec un accent marqué sur la résilience des identités numériques et la réduction de l’impact des compromissions de comptes. Le CERT-FR, dans son bilan des menaces 2025, continue de recommander un contact direct ([email protected] ou le 3218) dès qu’un compte semble compromis, signe que le vol d’identifiants reste une source majeure de sollicitations. Au niveau européen, l’ENISA a lancé l’exercice Cyber Europe 2026 pour tester la réponse collective des États membres face à des scénarios de compromission à grande échelle.

Les chiffres mondiaux racontent la même histoire. Le taux d’adoption du MFA au sein des effectifs professionnels a atteint 70 % en janvier 2025, contre 66 % un an plus tôt, selon le rapport Okta Secure Sign-in Trends 2025. Cette moyenne cache toutefois un déséquilibre net : les grandes entreprises dépassent largement les 80 %, quand les TPE et PME plafonnent souvent sous les 35 %. Ce sont précisément ces structures plus petites qui manquent de ressources pour surveiller manuellement chaque tentative de connexion suspecte, ce qui rend l’automatisation décrite plus loin d’autant plus utile.

Ce guide s’adresse aux administrateurs systèmes, aux responsables IT de PME et aux développeurs qui gèrent un tenant Microsoft 365 ou Entra ID et veulent passer d’un MFA basique à une configuration qui tient face aux techniques de contournement actuelles. Il s’inscrit dans notre couverture plus large de la cybersécurité, où nous suivons chaque semaine les failles, ransomwares et nouvelles obligations réglementaires qui touchent les entreprises européennes.

Un compte Microsoft 365 compromis ne se limite presque jamais à une boîte mail lue en douce. Une fois à l’intérieur, un attaquant peut créer des règles de transfert automatique invisibles, usurper l’identité d’un dirigeant pour valider un virement (la fraude dite au président), ou utiliser le compte comme point d’entrée pour déployer un ransomware sur l’ensemble du réseau via OneDrive et SharePoint. La double authentification n’élimine pas ce risque à elle seule, mais elle ferme la porte la plus utilisée pour y entrer, ce qui justifie de la traiter comme un projet prioritaire plutôt que comme une case à cocher dans un audit de conformité.

## Double authentification (MFA) : définition et fonctionnement technique

La double authentification combine au moins deux éléments de nature différente parmi trois catégories reconnues : un secret que vous connaissez (mot de passe, code PIN), un objet que vous possédez (téléphone, clé de sécurité, carte à puce) et une caractéristique qui vous est propre (empreinte digitale, reconnaissance faciale). Le NIST détaille ces niveaux de garantie d’authentification (AAL1 à AAL3) dans son guide SP 800-63B, qui sert de référence à la plupart des standards utilisés par les fournisseurs cloud.

Dans la pratique, trois familles de méthodes dominent le marché actuel. Les codes à usage unique (OTP) envoyés par SMS ou générés par une application comme Microsoft Authenticator reposent sur l’algorithme TOTP normalisé (RFC 6238), qui produit un code à six chiffres renouvelé toutes les 30 secondes à partir d’une clé secrète partagée. Les notifications push demandent simplement à l’utilisateur d’approuver la connexion depuis son téléphone, un mécanisme confortable qui a toutefois ouvert la porte aux attaques par lassitude MFA, où l’attaquant spamme les demandes d’approbation jusqu’à ce qu’un utilisateur fatigué clique sur « Approuver » par réflexe.

La troisième famille, bâtie sur le standard FIDO2/WebAuthn, fonctionne différemment : elle génère une paire de clés cryptographiques liée au domaine exact du service. Une clé de sécurité FIDO2 ou un passkey ne peut tout simplement pas être utilisé sur un faux site, même parfaitement imité, puisque la signature cryptographique échoue si le domaine ne correspond pas exactement. C’est cette propriété qui rend cette famille de méthodes réellement résistante au phishing, une distinction que Microsoft et Google mettent aujourd’hui en avant dans leur documentation.

Microsoft Entra ID (anciennement Azure Active Directory) centralise la gestion de ces méthodes via une politique unique, l’« Authentication methods policy » décrite dans la documentation Microsoft Learn, accessible aussi bien depuis le portail d’administration que depuis l’API Microsoft Graph. C’est ce qui permet d’automatiser entièrement son déploiement, comme dans les étapes suivantes.

## MFA classique vs MFA résistante au phishing : ce que les attaquants exploitent

Un code OTP reçu par SMS ou généré par une application TOTP protège contre le vol de mot de passe seul, mais pas contre l’interception en temps réel. Les kits de phishing-as-a-service comme Tycoon2FA ou EvilProxy, documentés par plusieurs éditeurs de sécurité, fonctionnent sur le principe de l’adversary-in-the-middle (AiTM) : ils placent un serveur proxy entre la victime et le vrai portail de connexion Microsoft. La victime saisit son mot de passe, reçoit et valide son code MFA normalement, mais le proxy capture au passage le jeton de session final. L’attaquant rejoue ensuite ce jeton pour accéder au compte, sans jamais avoir eu besoin de connaître le mot de passe ni le code à l’avance.

Ce type d’attaque explique pourquoi le SMS et même les applications TOTP classiques ne suffisent plus pour protéger des comptes à privilèges : administrateurs, direction financière, accès aux données RH. Le SIM swapping, où un attaquant convainc un opérateur téléphonique de transférer un numéro vers une carte SIM qu’il contrôle, reste par ailleurs une méthode active pour intercepter des codes envoyés par SMS.

Face à cela, les méthodes FIDO2, les passkeys et Windows Hello for Business ferment la porte à ce type de contournement, parce que la preuve cryptographique est liée au domaine réel du service et ne peut pas être rejouée ailleurs. Microsoft recommande désormais explicitement de migrer les comptes sensibles vers ces méthodes en priorité, tout en conservant le MFA classique pour les comptes à moindre risque le temps de la transition. Le rapport Okta Secure Sign-in Trends 2025 confirme cette bascule : l’adoption de l’authentification sans mot de passe résistante au phishing a progressé de 63 % en un an, passant de 8,6 % à 14,0 % des utilisateurs professionnels suivis par la plateforme.

