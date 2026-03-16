import { useState } from 'react'
import type { Menu } from '../../types/menu'

type MenuFormModalProps = {
  menu?: Menu | null
  onClose: () => void
  onSubmit: (values: { name: string }) => Promise<void>
}

export function MenuFormModal({ menu, onClose, onSubmit }: MenuFormModalProps) {
  const [name, setName] = useState(menu?.name ?? '')
  const [loading, setLoading] = useState(false)

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()
    setLoading(true)
    try {
      await onSubmit({ name })
      onClose()
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/40 p-4">
      <form className="w-full max-w-md rounded-2xl bg-white p-6 shadow-xl" onSubmit={handleSubmit}>
        <h3 className="text-xl font-bold text-[var(--brand-900)]">{menu ? 'Editar cardápio' : 'Novo cardápio'}</h3>

        <label className="mt-4 block text-sm font-semibold text-[var(--brand-900)]" htmlFor="menuName">Nome</label>
        <input id="menuName" className="mt-2 w-full rounded-xl border border-[var(--accent-200)] px-3 py-2" value={name} onChange={(event) => setName(event.target.value)} required />

        <div className="mt-5 flex gap-3">
          <button type="button" className="w-full rounded-xl border border-[var(--accent-200)] px-4 py-2" onClick={onClose}>Cancelar</button>
          <button disabled={loading} className="w-full rounded-xl bg-[var(--brand-700)] px-4 py-2 font-semibold text-white">
            {loading ? 'Salvando...' : 'Salvar'}
          </button>
        </div>
      </form>
    </div>
  )
}
