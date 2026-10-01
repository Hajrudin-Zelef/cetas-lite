---
id: collect-261001-general-networking/general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes-4
title: "macOS (Homebrew)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes.md
source_anchor: ""
source_lines: [441, 643]
sha256: f290e59966762a76f0e38dd49c3b822d39b380628dfe6c69b2f95259ad84def5
---

# macOS (Homebrew)

```
-- supabase/migrations/20260406130000_storage_attachments.sql
insert into storage.buckets (id, name, public, file_size_limit, allowed_mime_types)
values (
  'ticket-attachments',
  'ticket-attachments',
  false,
  10485760,  -- 10 MB max par fichier
  array['image/png','image/jpeg','image/webp','application/pdf']
)
on conflict (id) do nothing;
-- Policy : un utilisateur ne lit que ses propres uploads
create policy "Users read own attachments" on storage.objects
  for select using (
    bucket_id = 'ticket-attachments' and
    (storage.foldername(name))[1] = auth.uid()::text
  );
create policy "Users upload to own folder" on storage.objects
  for insert with check (
    bucket_id = 'ticket-attachments' and
    (storage.foldername(name))[1] = auth.uid()::text
  );
```
Côté client, créez un composant d’upload qui place chaque fichier dans le dossier `{user_id}/{ticket_id}/` :

```
// src/components/UploadAttachment.tsx
'use client'
import { useState } from 'react'
import { createClient } from '@/lib/supabase/client'
export function UploadAttachment({ ticketId }: { ticketId: string }) {
  const [uploading, setUploading] = useState(false)
  const supabase = createClient()
  async function handleUpload(e: React.ChangeEvent<HTMLInputElement>) {
    const file = e.target.files?.[0]
    if (!file) return
    setUploading(true)
    const { data: { user } } = await supabase.auth.getUser()
    if (!user) return
    const path = `${user.id}/${ticketId}/${Date.now()}-${file.name}`
    const { error } = await supabase.storage
      .from('ticket-attachments')
      .upload(path, file, { cacheControl: '3600', upsert: false })
    if (error) alert(error.message)
    else {
      await supabase.from('tickets').update({ attachment_path: path }).eq('id', ticketId)
      location.reload()
    }
    setUploading(false)
  }
  return (
    <input
      type="file"
      accept="image/*,application/pdf"
      onChange={handleUpload}
      disabled={uploading}
    />
  )
}
```
Pour afficher la pièce jointe, générez une URL signée temporaire – jamais d’URL publique sur un bucket privé :

```
const { data } = await supabase.storage
  .from('ticket-attachments')
  .createSignedUrl(path, 3600)  // valide 1 heure
```
## Étape 9 : Activer la mise à jour temps réel via Realtime

Supabase Realtime, écrit en Elixir et basé sur Phoenix Channels, propage les changements PostgreSQL via WebSockets – utile pour afficher de nouveaux messages instantanément. Le tier Free supporte 200 connexions simultanées et 2 millions de messages par mois ; le tier Pro monte à 500 connexions et 5 millions de messages inclus.

Activez la publication Realtime sur les tables concernées :

```
-- supabase/migrations/20260406140000_enable_realtime.sql
alter publication supabase_realtime add table public.messages;
alter publication supabase_realtime add table public.tickets;
```
Côté React, abonnez-vous aux insertions sur la table `messages` :

```
// src/components/MessagesList.tsx
'use client'
import { useEffect, useState } from 'react'
import { createClient } from '@/lib/supabase/client'
type Message = { id: string; body: string; created_at: string; user_id: string }
export function MessagesList({ ticketId, initial }: { ticketId: string; initial: Message[] }) {
  const [messages, setMessages] = useState(initial)
  const supabase = createClient()
  useEffect(() => {
    const channel = supabase
      .channel(`messages:${ticketId}`)
      .on(
        'postgres_changes',
        {
          event: 'INSERT',
          schema: 'public',
          table: 'messages',
          filter: `ticket_id=eq.${ticketId}`,
        },
        (payload) => {
          setMessages((prev) => [...prev, payload.new as Message])
        }
      )
      .subscribe()
    return () => { supabase.removeChannel(channel) }
  }, [ticketId])
  return (
    <ul className="space-y-2">
      {messages.map((m) => (
        <li key={m.id} className="border p-2 rounded">
          <p>{m.body}</p>
          <time className="text-xs text-gray-500">{new Date(m.created_at).toLocaleString('fr-FR')}</time>
        </li>
      ))}
    </ul>
  )
}
```
Important : les politiques RLS s’appliquent aussi aux événements Realtime depuis la version 2024 – un utilisateur ne reçoit que les changements qu’il aurait pu lire via une requête classique. Vérifiez que vos policies SELECT sont correctes, sous peine de manquer des notifications légitimes.

## Étape 10 : Écrire une Edge Function Deno qui envoie un email

Les Edge Functions Supabase tournent sur Deno 2.4.1, déployées dans plus de 30 régions edge à travers le monde – latence moyenne sous 50 ms en Europe. Le tier Free inclut 500 000 invocations par mois ; le Pro 2 millions, avec un démarrage à froid sous 100 ms grâce à V8 Isolates.

`supabase functions new notify-new-ticket`
Éditez la fonction générée pour appeler l’API Resend :

```
// supabase/functions/notify-new-ticket/index.ts
import "jsr:@supabase/functions-js/edge-runtime.d.ts"
interface TicketWebhook {
  type: 'INSERT'
  table: 'tickets'
  record: {
    id: string
    title: string
    user_id: string
    priority: string
  }
}
Deno.serve(async (req: Request) => {
  if (req.method !== 'POST') {
    return new Response('Method not allowed', { status: 405 })
  }
  const payload: TicketWebhook = await req.json()
  const { title, priority, user_id } = payload.record
  const resendKey = Deno.env.get('RESEND_API_KEY')
  if (!resendKey) {
    return new Response('Missing RESEND_API_KEY', { status: 500 })
  }
  const res = await fetch('https://api.resend.com/emails', {
    method: 'POST',
    headers: {
      Authorization: `Bearer ${resendKey}`,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      from: '[email protected]',
      to: '[email protected]',
      subject: `[${priority.toUpperCase()}] Nouveau ticket : ${title}`,
      text: `Un nouveau ticket a été créé par ${user_id}.\nTitre : ${title}`,
    }),
  })
  if (!res.ok) {
    return new Response('Resend API error', { status: 502 })
  }
  return new Response(JSON.stringify({ ok: true }), {
    headers: { 'Content-Type': 'application/json' },
  })
})
```
Déployez la fonction et configurez le secret :

```
# Servir en local
supabase functions serve notify-new-ticket --env-file ./supabase/functions/.env
# Déployer en production
supabase secrets set RESEND_API_KEY=re_votre_cle_ici
supabase functions deploy notify-new-ticket --no-verify-jwt
```
Branchez ensuite un Database Webhook dans le dashboard Supabase : *Database → Webhooks → Create*, table `tickets`, événement `INSERT`, target l’URL de votre fonction. À chaque création de ticket, l’email partira automatiquement.

## Étape 11 : Lier le projet local au cloud et déployer en production

Connectez votre projet local à votre projet cloud Supabase et appliquez toutes les migrations en une commande :

```
# Récupérer la liste des projets
supabase projects list
# Lier le projet (remplacez REF par votre Project Ref)
supabase link --project-ref abcdefghijklmnopqrst
# Pousser toutes les migrations locales vers la prod
supabase db push
# Sortie attendue
Connecting to remote database...
Applying migration 20260406120000_create_tickets_schema.sql...
Applying migration 20260406130000_storage_attachments.sql...
Applying migration 20260406140000_enable_realtime.sql...
Finished supabase db push.
```
Mettez à jour `.env.local` avec les clés de production récupérées dans *Project Settings → API*, puis déployez le front Next.js sur Vercel :

