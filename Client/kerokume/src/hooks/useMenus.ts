import { useCallback, useEffect, useState } from 'react'
import {
  activateMenu,
  createMenu,
  deleteMenu,
  getMenusByRestaurant,
  updateMenu,
} from '../services/menuService'
import type { Menu, MenuPayload } from '../types/menu'

export function useMenus(restaurantId: string | null) {
  const [menus, setMenus] = useState<Menu[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const loadMenus = useCallback(async () => {
    if (!restaurantId) return

    setLoading(true)
    setError(null)

    try {
      const data = await getMenusByRestaurant(restaurantId)
      setMenus(data)
    } catch {
      setError('Não foi possível carregar os cardápios.')
    } finally {
      setLoading(false)
    }
  }, [restaurantId])

  useEffect(() => {
    void loadMenus()
  }, [loadMenus])

  async function handleCreate(payload: Omit<MenuPayload, 'restaurantId'>) {
    if (!restaurantId) return
    await createMenu({ ...payload, restaurantId })
    await loadMenus()
  }

  async function handleUpdate(menuId: string, payload: Omit<MenuPayload, 'restaurantId'>) {
    if (!restaurantId) return
    await updateMenu(menuId, { ...payload, restaurantId })
    await loadMenus()
  }

  async function handleDelete(menuId: string) {
    await deleteMenu(menuId)
    await loadMenus()
  }

  async function handleActivate(menuId: string) {
    await activateMenu(menuId)
    await loadMenus()
  }

  return {
    menus,
    loading,
    error,
    loadMenus,
    handleCreate,
    handleUpdate,
    handleDelete,
    handleActivate,
  }
}
