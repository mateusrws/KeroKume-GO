import { useParams } from 'react-router-dom'
import { useFoods } from '../../hooks/useFoods'

export function PublicMenuPage() {
  const { menuId } = useParams<{ menuId: string }>()
  const { foods, loading, error } = useFoods(menuId ?? null)

  return (
    <main className="min-h-screen bg-gradient-to-b from-[var(--bg-from)] via-[var(--bg-via)] to-[var(--bg-to)] px-6 py-8">
      <section className="mx-auto max-w-3xl rounded-2xl border border-[var(--accent-100)] bg-white/80 p-6 shadow-xl">
        <p className="inline-block rounded-full bg-[var(--brand-100)] px-3 py-1 text-xs font-bold text-[var(--brand-800)]">Menu público</p>
        <h1 className="mt-2 text-3xl font-black text-[var(--brand-950)]">Cardápio do restaurante</h1>

        {error && <p className="mt-4 rounded-lg bg-red-100 px-3 py-2 text-sm text-red-700">{error}</p>}
        {loading && <p className="mt-4 text-sm text-[var(--text-muted)]">Carregando itens...</p>}

        <div className="mt-5 space-y-3">
          {foods.map((food) => (
            <article key={food.id} className="rounded-xl border border-[var(--accent-100)] bg-white p-4">
              <div className="flex items-start justify-between gap-2">
                <div>
                  <h2 className="font-bold text-[var(--brand-900)]">{food.name}</h2>
                  <p className="text-sm text-[var(--text-muted)]">{food.description}</p>
                </div>
                <p className="font-semibold text-[var(--accent-900)]">R$ {food.price.toFixed(2)}</p>
              </div>
            </article>
          ))}
        </div>
      </section>
    </main>
  )
}
