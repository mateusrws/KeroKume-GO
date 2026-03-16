type QrCodeCardProps = {
  publicMenuUrl: string
}

export function QrCodeCard({ publicMenuUrl }: QrCodeCardProps) {
  const qrCodeUrl = `https://api.qrserver.com/v1/create-qr-code/?size=260x260&data=${encodeURIComponent(publicMenuUrl)}`

  return (
    <section className="rounded-2xl border border-[var(--accent-100)] bg-white p-5 text-center shadow-sm">
      <h2 className="text-lg font-bold text-[var(--brand-900)]">QR Code do menu ativo</h2>
      <p className="mt-1 text-sm text-[var(--text-muted)]">Escaneie para abrir o cardápio público.</p>
      <img className="mx-auto mt-4 rounded-xl border border-[var(--accent-100)]" src={qrCodeUrl} alt="QR Code do menu" />
      <p className="mt-3 break-all text-xs text-[var(--text-muted)]">{publicMenuUrl}</p>
    </section>
  )
}
