import { useState } from 'react'
import type { Food, FoodCategory } from '../../types/food'

type FoodFormModalProps = {
  food?: Food | null
  onClose: () => void
  onSubmit: (values: { name: string; description: string; price: number; foodCategory: FoodCategory }) => Promise<void>
}

export function FoodFormModal({ food, onClose, onSubmit }: FoodFormModalProps) {
  const [name, setName] = useState(food?.name ?? '')
  const [description, setDescription] = useState(food?.description ?? '')
  const [price, setPrice] = useState(food?.price.toString() ?? '')
  const [foodCategory, setFoodCategory] = useState<FoodCategory>(food?.foodCategory ?? 'COMIDA')

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()
    await onSubmit({
      name,
      description,
      price: Number(price),
      foodCategory,
    })
    onClose()
  }

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/40 p-4">
      <form className="w-full max-w-lg rounded-2xl bg-white p-6 shadow-xl" onSubmit={handleSubmit}>
        <h3 className="text-xl font-bold text-[var(--brand-900)]">{food ? 'Editar prato' : 'Novo prato'}</h3>

        <div className="mt-4 grid gap-3">
          <input className="rounded-xl border border-[var(--accent-200)] px-3 py-2" placeholder="Nome" value={name} onChange={(event) => setName(event.target.value)} required />
          <textarea className="rounded-xl border border-[var(--accent-200)] px-3 py-2" placeholder="Descrição" value={description} onChange={(event) => setDescription(event.target.value)} required />
          <input className="rounded-xl border border-[var(--accent-200)] px-3 py-2" type="number" step="0.01" placeholder="Preço" value={price} onChange={(event) => setPrice(event.target.value)} required />

          <select className="rounded-xl border border-[var(--accent-200)] px-3 py-2" value={foodCategory} onChange={(event) => setFoodCategory(event.target.value as FoodCategory)}>
            <option value="COMIDA">Comida</option>
            <option value="BEBIDA">Bebida</option>
          </select>
        </div>

        <div className="mt-5 flex gap-3">
          <button type="button" className="w-full rounded-xl border border-[var(--accent-200)] px-4 py-2" onClick={onClose}>Cancelar</button>
          <button className="w-full rounded-xl bg-[var(--brand-700)] px-4 py-2 font-semibold text-white">Salvar</button>
        </div>
      </form>
    </div>
  )
}
