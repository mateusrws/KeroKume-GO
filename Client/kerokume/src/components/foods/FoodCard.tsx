import type { Food } from '../../types/food'

type FoodCardProps = {
  food: Food
  onEdit: (food: Food) => void
  onDelete: (foodId: string) => void
}

export function FoodCard({ food, onEdit, onDelete }: FoodCardProps) {
  return (
    <article className="rounded-2xl border border-[var(--accent-100)] bg-white p-4 shadow-sm">
      <div className="flex items-start justify-between gap-3">
        <div>
          <h3 className="font-bold text-[var(--brand-900)]">{food.name}</h3>
          <p className="text-sm text-[var(--text-muted)]">{food.description}</p>
          <p className="mt-2 text-sm font-semibold text-[var(--accent-900)]">
            R$ {food.price.toFixed(2)} • {food.foodCategory}
          </p>
        </div>
        <span className={`rounded-full px-2 py-1 text-xs font-semibold ${food.isAvailable ? 'bg-green-100 text-green-700' : 'bg-zinc-100 text-zinc-600'}`}>
          {food.isAvailable ? 'Disponível' : 'Indisponível'}
        </span>
      </div>

      <div className="mt-4 flex gap-2">
        <button className="rounded-lg border border-[var(--brand-300)] px-3 py-2 text-sm font-semibold text-[var(--brand-800)]" onClick={() => onEdit(food)}>
          Editar
        </button>
        <button className="rounded-lg border border-red-200 px-3 py-2 text-sm font-semibold text-red-700" onClick={() => onDelete(food.id)}>
          Excluir
        </button>
      </div>
    </article>
  )
}
