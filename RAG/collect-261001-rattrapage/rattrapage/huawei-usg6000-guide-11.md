---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-11
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [1558, 1735]
sha256: bfba497009b24adda2a5e77cbe6451165c7f80849a9844b248d25baf4a9fa72b
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

- **EICAR** : télécharge le fichier test `eicar.com` (inoffensif, détecté par tous les AV) depuis le LAN → l'USG doit le **bloquer** et logger. Si ça passe, ton profil AV n'est pas actif sur la règle.
- **URL test** : des sites de test constructeur/standards (ex. pages de test de filtrage) pour valider le blocage par catégorie.
- **IPS** : déclencheurs de test (scan de ports depuis l'extérieur vers la DMZ) → vérifier les logs IPS.
- 📋 Après chaque changement de profil UTM : **re-teste EICAR**. 2 minutes, zéro risque, 100 % utile.

## 86. UTM et faux positifs : la procédure

1. L'utilisateur signale un blocage légitime → récupérer **l'heure, l'URL/IP, l'utilisateur**.
2. Logs UTM → identifier le **profil et la signature** en cause.
3. Décision : **exception ciblée** (whitelist d'URL, exclusion de signature) — jamais « désactiver le profil ».
4. Documenter l'exception (qui, quoi, pourquoi, date de revue).
5. 📋 **Revue trimestrielle** des exceptions : une whitelist oubliée, c'est une porte ouverte.

## 87. Signatures : mise à jour hors-ligne (site isolé)

Si l'USG n'a pas d'accès Internet :
1. Télécharger les packs de signatures sur support.huawei.com (depuis un poste connecté).
2. Les copier sur **clé USB**.
3. Web : System > Update > Local Update, ou CLI selon version.
4. Vérifier les versions (`display av update info`).
5. 📋 Planifier : quelqu'un doit le faire **tous les mois** — mets-le dans le planning de maintenance (section 126).

## 88. Web : superviser l'UTM en un coup d'œil

Chemin (New Web UI) : **Monitor > Security Protection** (ou Threat/UTM selon version) : top des menaces bloquées, top des URL bloquées, état des signatures, utilisation CPU par moteur.
💡 En routine hebdo, 5 minutes sur cet écran valent mieux qu'un rapport mensuel que personne ne lit.

---
---

# BLOC H — AUTHENTIFICATION DES UTILISATEURS

## 89. Pourquoi authentifier : la politique par utilisateur

Sans authentification, tes règles voient des **IP**. Avec, elles voient des **personnes** : « le service compta accède à l'appli métier, les stagiaires non » — même en DHCP, même en Wi-Fi.

Modes supportés : **base locale**, **RADIUS**, **LDAP/AD**, **certificats**, **SAML/OTP** (selon version/licence).

## 90. Base locale : pour les petits sites et le secours

```huawei
system-view
[USG] aaa
[USG-aaa] local-user alice password
[USG-aaa] local-user alice service-type sslvpn web
[USG-aaa] local-user alice privilege level 3
[USG-aaa] local-user alice access-limit 2      # 2 sessions simultanées max
[USG-aaa] quit
# Groupe local :
[USG] user-group comptabilite
[USG-user-group-comptabilite] quit
[USG] aaa
[USG-aaa] local-user alice user-group comptabilite
[USG-aaa] quit
save
```

⚠️ La base locale, c'est pour **< 20 utilisateurs** ou le **compte de secours**. Au-delà : AD/LDAP/RADIUS (sinon tu passes ta vie à gérer les mots de passe).

## 91. AD/LDAP : brancher l'annuaire d'entreprise

```huawei
system-view
[USG] ldap-server template template-ad
[USG-ldap-template-ad] ldap-server authentication 192.168.10.5 389   # IP fictive : ton AD
[USG-ldap-template-ad] ldap-server bind-dn cn=svc-firewall,ou=services,dc=entreprise,dc=local
[USG-ldap-template-ad] ldap-server bind-password Fictif123!
[USG-ldap-template-ad] ldap-server search-dn ou=utilisateurs,dc=entreprise,dc=local
[USG-ldap-template-ad] quit
[USG] aaa
[USG-aaa] authentication-scheme auth-ad
[USG-aaa-authentication-scheme-auth-ad] authentication-mode ldap
[USG-aaa-authentication-scheme-auth-ad] quit
[USG-aaa] domain domaine-entreprise
[USG-aaa-domain-domaine-entreprise] authentication-scheme auth-ad
[USG-aaa-domain-domaine-entreprise] ldap-server template-ad
[USG-aaa-domain-domaine-entreprise] quit
[USG-aaa] quit
save
```

🔧 Tester : `test-aaa ldap-template template-ad user alice password` (syntaxe à vérifier selon version) ou via le bouton « Test » du web. **Toujours tester AVANT** de basculer les utilisateurs dessus.

## 92. RADIUS : l'alternative standard

```huawei
system-view
[USG] radius-server template template-radius
[USG-radius-template-radius] radius-server authentication 192.168.10.6 1812
[USG-radius-template-radius] radius-server shared-key FictifRadius123!
[USG-radius-template-radius] quit
[USG] aaa
[USG-aaa] authentication-scheme auth-radius
[USG-aaa-authentication-scheme-auth-radius] authentication-mode radius
[USG-aaa-authentication-scheme-auth-radius] quit
[USG-aaa] domain domaine-entreprise
[USG-aaa-domain-domaine-entreprise] authentication-scheme auth-radius
[USG-aaa-domain-domaine-entreprise] radius-server template-radius
[USG-aaa-domain-domaine-entreprise] quit
[USG-aaa] quit
save
```

💡 RADIUS = le choix naturel si tu as déjà un **NPS/FreeRADIUS** (souvent couplé à l'authentification Wi-Fi/802.1X — même annuaire, même logique).

## 93. Politique par utilisateur/groupe : l'exemple qui parle

```huawei
[USG] security-policy
[USG-policy-security] rule name compta-appli-metier
[USG-policy-security-rule-compta-appli-metier] source-zone trust
[USG-policy-security-rule-compta-appli-metier] destination-zone dmz
[USG-policy-security-rule-compta-appli-metier] user-group comptabilite
[USG-policy-security-rule-compta-appli-metier] destination-address 172.16.1.30 32
[USG-policy-security-rule-compta-appli-metier] service https
[USG-policy-security-rule-compta-appli-metier] action permit
[USG-policy-security-rule-compta-appli-metier] quit
# Les autres utilisateurs : règle générale SANS critère user, placée APRÈS :
[USG-policy-security] rule name lan-dmz-general
[USG-policy-security-rule-lan-dmz-general] source-zone trust
[USG-policy-security-rule-lan-dmz-general] destination-zone dmz
[USG-policy-security-rule-lan-dmz-general] service ping
[USG-policy-security-rule-lan-dmz-general] action permit
[USG-policy-security-rule-lan-dmz-general] quit
[USG-policy-security] quit
save
```

⚠️ Pour que le critère `user` fonctionne, l'utilisateur doit être **authentifié** (portail, SSL VPN, 802.1X...) — sinon la règle ne matche jamais. Prévois une **politique d'authentification** (qui doit s'authentifier, quand).

## 94. Portail captif : authentifier les utilisateurs du LAN

```huawei
system-view
[USG] web-auth-server template portail-interne port 8080
[USG] authentication-policy
[USG-policy-auth] rule name auth-lan
[USG-policy-auth-rule-auth-lan] source-zone trust
[USG-policy-auth-rule-auth-lan] destination-zone untrust
[USG-policy-auth-rule-auth-lan] source-address 192.168.10.0 24
[USG-policy-auth-rule-auth-lan] action auth
[USG-policy-auth-rule-auth-lan] quit
[USG-policy-auth] quit
save
```

L'utilisateur ouvre son navigateur → portail → login → ses flux sont ensuite associés à son identité. (Syntaxe exacte à vérifier selon version — le web **Policy > Authentication Policy** est plus direct.)

## 95. Verrouillage et politique de mot de passe (comptes locaux)

```huawei
system-view
[USG] aaa
[USG-aaa] local-user password-policy
# Complexité, longueur min, historique, durée de vie (paramètres selon version)
[USG-aaa] quit
save
```

Voir aussi section 133 (durcissement). Règle : **les comptes locaux d'admin** suivent la même politique que les utilisateurs — pas de `Admin123` qui traîne.

## 96. Diagnostic d'authentification

```huawei
[USG] display aaa online-fail-record       # échecs de connexion récents
[USG] display local-user                   # comptes locaux
[USG] display ldap-server template template-ad   # état du template LDAP
```

🔧 « Personne ne s'authentifie depuis ce matin » → vérifier dans l'ordre : **connectivité vers l'AD** (ping/LDAP port 389), **compte de bind** (mot de passe expiré ?), **politique** (quelqu'un a touché aux règles ?), **heure** (désynchro NTP = échec Kerberos/LDAP).

---
---

# BLOC I — HAUTE DISPONIBILITÉ (HOT STANDBY)

## 97. Principe : VGMP + HRP, en clair

