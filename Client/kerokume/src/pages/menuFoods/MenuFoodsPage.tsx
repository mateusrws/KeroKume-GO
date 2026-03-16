import { useEffect, useMemo, useState } from 'react'
import { Link, useLocation, useParams, useSearchParams } from 'react-router-dom'
import toast, { Toaster } from 'react-hot-toast'
import { FoodCard } from '../../components/foods/FoodCard'
import { FoodFormModal } from '../../components/foods/FoodFormModal'
import { useFoods } from '../../hooks/useFoods'
import type { Food } from '../../types/food'

const SELECTED_MENU_ID_KEY = 'selected-menu-id'

export function MenuFoodsPage() {
  const { menuId: menuIdFromParams } = useParams<{ menuId: string }>()
  const [searchParams] = useSearchParams()
  const menuIdFromQuery = searchParams.get('menuId')

  const selectedMenuId = useMemo(() => {
    return menuIdFromParams ?? menuIdFromQuery ?? localStorage.getItem(SELECTED_MENU_ID_KEY)
  }, [menuIdFromParams, menuIdFromQuery])

  useEffect(() => {
    if (selectedMenuId) {
      localStorage.setItem(SELECTED_MENU_ID_KEY, selectedMenuId)
    }
  }, [selectedMenuId])

  const location = useLocation()
  const menuName = useMemo(() => (location.state as { menuName?: string } | null)?.menuName ?? 'Cardápio', [location.state])

  const { foods, loading, error, handleCreate, handleDelete, handleUpdate } = useFoods(selectedMenuId)

  const [isCreateOpen, setIsCreateOpen] = useState(false)
  const [editingFood, setEditingFood] = useState<Food | null>(null)

  async function createFood(values: { name: string; description: string; price: number; foodCategory: 'COMIDA' | 'BEBIDA' }) {
    try {
      await handleCreate(values)
      toast.success('Prato criado com sucesso!')
    } catch {
      toast.error('Erro ao criar prato. Confirme se um menu foi selecionado.')
    }
  }

  async function updateFood(values: { name: string; description: string; price: number; foodCategory: 'COMIDA' | 'BEBIDA' }) {
    if (!editingFood) return
    try {
      await handleUpdate(editingFood.id, values)
      toast.success('Prato atualizado com sucesso!')
    } catch {
      toast.error('Erro ao atualizar prato. Confirme se um menu foi selecionado.')
    }
  }

  async function removeFood(foodId: string) {
    try {
      await handleDelete(foodId)
      toast.success('Prato removido com sucesso!')
    } catch {
      toast.error('Erro ao remover prato. Confirme se um menu foi selecionado.')
    }
  }

  return (
    <main className="min-h-screen bg-gradient-to-b from-[var(--bg-from)] via-[var(--bg-via)] to-[var(--bg-to)] px-6 py-8">
      <Toaster position="top-center" />
      <section className="mx-auto max-w-5xl">
        <div className="mb-4 flex flex-wrap items-center justify-between gap-3">
          <div>
            <Link className="text-sm font-semibold text-[var(--brand-800)]" to="/my-menus">← Voltar para cardápios</Link>
            <h1 className="mt-2 text-3xl font-black text-[var(--brand-950)]">{menuName}</h1>
            <p className="text-[var(--text-muted)]">Gerencie as comidas e bebidas deste cardápio.</p>
          </div>

          <button className="rounded-xl bg-[var(--brand-700)] px-4 py-2 font-semibold text-white" onClick={() => setIsCreateOpen(true)}>
            + Novo prato
          </button>
        </div>

        {!selectedMenuId && (
          <p className="mb-3 rounded-lg bg-red-100 px-3 py-2 text-sm text-red-700">
            Nenhum menu selecionado. Volte para cardápios e abra um menu específico.
          </p>
        )}

        {error && <p className="mb-3 rounded-lg bg-red-100 px-3 py-2 text-sm text-red-700">{error}</p>}

        <section className="grid gap-3 md:grid-cols-2">
          {loading && <p className="text-sm text-[var(--text-muted)]">Carregando pratos...</p>}
          {!loading && foods.map((food) => (
            <FoodCard key={food.id} food={food} onDelete={removeFood} onEdit={setEditingFood} />
          ))}
        </section>
      </section>

      {isCreateOpen && <FoodFormModal onClose={() => setIsCreateOpen(false)} onSubmit={createFood} />}
      {editingFood && <FoodFormModal food={editingFood} onClose={() => setEditingFood(null)} onSubmit={updateFood} />}
    </main>
  )
}
