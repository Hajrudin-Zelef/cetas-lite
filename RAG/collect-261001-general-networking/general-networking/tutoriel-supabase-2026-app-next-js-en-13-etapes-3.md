---
id: collect-261001-general-networking/general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes-3
title: "macOS (Homebrew)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Apple", "Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes.md
source_anchor: ""
source_lines: [280, 440]
sha256: f74e8df55f181ea1357fba29067891525d3eadbccbb3f8eed60fc870b034c587
---

# macOS (Homebrew)

Supabase Auth, basé sur GoTrue (un fork open-source de Netlify Identity, dont la documentation d’auto-hébergement de février 2026 pin désormais l’image de production à la version 2.186.0), supporte plus de 20 fournisseurs OAuth – Google, GitHub, Apple, Discord, Azure AD, Facebook, Spotify, LinkedIn, GitLab, Bitbucket, Twitter/X, Notion, Slack, Twitch, Zoom, Figma, Keycloak, Kakao, Workos, et tout fournisseur OIDC ou SAML 2.0 personnalisé. Pour ce tutoriel, nous combinons l’authentification par **Magic Link** (mot de passe sans mot de passe via email) et OAuth GitHub.

```
// src/app/login/page.tsx
'use client'
import { useState } from 'react'
import { createClient } from '@/lib/supabase/client'
export default function LoginPage() {
  const [email, setEmail] = useState('')
  const [sent, setSent] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const supabase = createClient()
  async function handleMagicLink(e: React.FormEvent) {
    e.preventDefault()
    setError(null)
    const { error } = await supabase.auth.signInWithOtp({
      email,
      options: {
        emailRedirectTo: `${location.origin}/auth/callback`,
      },
    })
    if (error) setError(error.message)
    else setSent(true)
  }
  async function handleGitHub() {
    await supabase.auth.signInWithOAuth({
      provider: 'github',
      options: { redirectTo: `${location.origin}/auth/callback` },
    })
  }
  if (sent) return <p>Vérifiez votre boîte mail !</p>
  return (
    <div className="max-w-sm mx-auto p-6">
      <h1 className="text-2xl font-bold mb-4">Connexion</h1>
      <form onSubmit={handleMagicLink} className="space-y-3">
        <input
          type="email"
          required
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="[email protected]"
          className="w-full px-3 py-2 border rounded"
        />
        <button type="submit" className="w-full bg-blue-600 text-white py-2 rounded">
          Envoyer le lien magique
        </button>
      </form>
      <button onClick={handleGitHub} className="w-full mt-3 bg-gray-900 text-white py-2 rounded">
        Connexion avec GitHub
      </button>
      {error && <p className="text-red-600 mt-2">{error}</p>}
    </div>
  )
}
```
Créez le route handler de callback OAuth qui échange le code contre une session :

```
// src/app/auth/callback/route.ts
import { createClient } from '@/lib/supabase/server'
import { NextResponse } from 'next/server'
export async function GET(request: Request) {
  const { searchParams, origin } = new URL(request.url)
  const code = searchParams.get('code')
  const next = searchParams.get('next') ?? '/app'
  if (code) {
    const supabase = await createClient()
    const { error } = await supabase.auth.exchangeCodeForSession(code)
    if (!error) return NextResponse.redirect(`${origin}${next}`)
  }
  return NextResponse.redirect(`${origin}/login?error=auth_callback`)
}
```
En local, les emails de Magic Link arrivent dans **Inbucket** sur http://127.0.0.1:54324 – aucun envoi réel. Pour la production, configurez un fournisseur SMTP (Resend, Postmark, SendGrid) dans `config.toml` ou via le dashboard Supabase, sous peine de quota strict de 4 emails par heure sur le service par défaut.

## Étape 7 : Lister, créer et modifier des tickets via Server Actions

Les Server Actions de Next.js 16 simplifient drastiquement les mutations : plus besoin d’écrire un endpoint API REST séparé. La fonction tourne sur le serveur, peut accéder directement au client Supabase server-side, et invalide automatiquement le cache du Router via `revalidatePath()`.

```
// src/app/app/actions.ts
'use server'
import { createClient } from '@/lib/supabase/server'
import { revalidatePath } from 'next/cache'
import { redirect } from 'next/navigation'
export async function createTicket(formData: FormData) {
  const supabase = await createClient()
  const { data: { user } } = await supabase.auth.getUser()
  if (!user) redirect('/login')
  const title = String(formData.get('title') ?? '').trim()
  const description = String(formData.get('description') ?? '').trim()
  const priority = String(formData.get('priority') ?? 'medium')
  if (title.length < 3) return { error: 'Titre trop court (min 3 caractères)' }
  const { error } = await supabase.from('tickets').insert({
    user_id: user.id,
    title,
    description,
    priority,
  })
  if (error) return { error: error.message }
  revalidatePath('/app')
  return { success: true }
}
export async function updateTicketStatus(id: string, status: string) {
  const supabase = await createClient()
  const { error } = await supabase
    .from('tickets')
    .update({ status })
    .eq('id', id)
  if (error) return { error: error.message }
  revalidatePath('/app')
  revalidatePath(`/app/${id}`)
  return { success: true }
}
```
La page liste utilise un Server Component qui charge les tickets directement depuis PostgreSQL – Row Level Security garantit que seuls les tickets de l’utilisateur connecté sont retournés, sans aucune logique de filtrage côté application :

```
// src/app/app/page.tsx
import { createClient } from '@/lib/supabase/server'
import Link from 'next/link'
import { createTicket } from './actions'
export default async function TicketsPage() {
  const supabase = await createClient()
  const { data: tickets } = await supabase
    .from('tickets')
    .select('id, title, status, priority, created_at')
    .order('created_at', { ascending: false })
    .limit(50)
  return (
    <main className="max-w-3xl mx-auto p-6">
      <h1 className="text-3xl font-bold">Mes tickets</h1>
      <form action={createTicket} className="mt-6 space-y-2 border p-4 rounded">
        <input name="title" placeholder="Titre" required className="w-full border p-2 rounded" />
        <textarea name="description" placeholder="Description" className="w-full border p-2 rounded" />
        <select name="priority" className="w-full border p-2 rounded">
          <option value="low">Faible</option>
          <option value="medium" selected>Moyenne</option>
          <option value="high">Haute</option>
          <option value="urgent">Urgente</option>
        </select>
        <button className="bg-blue-600 text-white px-4 py-2 rounded">Créer</button>
      </form>
      <ul className="mt-6 space-y-2">
        {tickets?.map(t => (
          <li key={t.id} className="border p-3 rounded">
            <Link href={`/app/${t.id}`} className="font-semibold">{t.title}</Link>
            <span className="ml-2 text-sm text-gray-500">{t.status} · {t.priority}</span>
          </li>
        ))}
      </ul>
    </main>
  )
}
```
## Étape 8 : Uploader des pièces jointes vers Supabase Storage

Supabase Storage est compatible avec l’API S3 d’Amazon, avec un bucket dédié par usage et des politiques d’accès basées sur les mêmes règles RLS que la base. Le tier Free offre 1 GB de stockage et 2 GB de bande passante mensuelle ; le Pro monte à 100 GB de stockage et 250 GB de bande passante incluse.

Créez un bucket privé via une migration SQL :

