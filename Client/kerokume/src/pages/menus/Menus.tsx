import { useState } from 'react'
import toast, { Toaster } from 'react-hot-toast'
import { MenuCard } from '../../components/menus/MenuCard'
import { MenuFormModal } from '../../components/menus/MenuFormModal'
import { QrCodeCard } from '../../components/menus/QrCodeCard'
import { useMenus } from '../../hooks/useMenus'
import { getRestaurantIdFromToken } from '../../lib/auth'
import type { Menu } from '../../types/menu'

export const RestaurantMenus = () => {
  const restaurantId = getRestaurantIdFromToken()
  const { menus, loading, error, handleCreate, handleDelete, handleUpdate, handleActivate } = useMenus(restaurantId)

  const [editingMenu, setEditingMenu] = useState<Menu | null>(null)
  const [isCreateOpen, setIsCreateOpen] = useState(false)

  const activeMenu = menus.find((menu) => menu.isActive)
  const publicMenuUrl = `${window.location.origin}/public-menu/${activeMenu?.id ?? ''}`

  async function createMenu(values: { name: string }) {
    try {
      await handleCreate(values)
      toast.success('Cardápio criado com sucesso!')
    } catch {
      toast.error('Erro ao criar cardápio.')
    }
  }

  async function updateMenu(values: { name: string }) {
    if (!editingMenu) return
    try {
      await handleUpdate(editingMenu.id, values)
      toast.success('Cardápio atualizado com sucesso!')
    } catch {
      toast.error('Erro ao atualizar cardápio.')
    }
  }

  async function deleteCurrentMenu(menuId: string) {
    try {
      await handleDelete(menuId)
      toast.success('Cardápio removido com sucesso!')
    } catch {
      toast.error('Erro ao remover cardápio.')
    }
  }

  async function activateCurrentMenu(menuId: string) {
    try {
      await handleActivate(menuId)
      toast.success('Cardápio ativado com sucesso!')
    } catch {
      toast.error('Erro ao ativar cardápio.')
    }
  }

  return (
    <main className="min-h-screen bg-gradient-to-b from-[var(--bg-from)] via-[var(--bg-via)] to-[var(--bg-to)] px-6 py-8">
      <Toaster position="top-center" />
      <section className="mx-auto max-w-6xl">
        <header className="mb-4 flex flex-wrap items-end justify-between gap-3">
          <div>
            <p className="inline-block rounded-full bg-[var(--brand-100)] px-3 py-1 text-xs font-bold text-[var(--brand-800)]">Painel do restaurante</p>
            <h1 className="mt-2 text-3xl font-black text-[var(--brand-950)]">Todos os cardápios</h1>
            <p className="text-[var(--text-muted)]">Gerencie menus, pratos e publicação.</p>
          </div>

          <button onClick={() => setIsCreateOpen(true)} className="rounded-xl bg-[var(--brand-700)] px-4 py-2 font-semibold text-white">
            + Novo cardápio
          </button>
        </header>

        {error && <p className="mb-3 rounded-lg bg-red-100 px-3 py-2 text-sm text-red-700">{error}</p>}

        <section className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          {loading && <p className="text-sm text-[var(--text-muted)]">Carregando cardápios...</p>}
          {!loading && menus.map((menu) => (
            <MenuCard
              key={menu.id}
              menu={menu}
              onDelete={deleteCurrentMenu}
              onEdit={setEditingMenu}
              onActivate={activateCurrentMenu}
            />
          ))}
        </section>

        {activeMenu && (
          <section className="mt-6 max-w-lg">
            <QrCodeCard publicMenuUrl={publicMenuUrl} />
          </section>
        )}
      </section>

      {isCreateOpen && <MenuFormModal onClose={() => setIsCreateOpen(false)} onSubmit={createMenu} />}
      {editingMenu && <MenuFormModal menu={editingMenu} onClose={() => setEditingMenu(null)} onSubmit={updateMenu} />}
    </main>
  )
}
