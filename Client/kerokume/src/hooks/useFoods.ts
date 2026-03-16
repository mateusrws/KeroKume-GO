import { useCallback, useEffect, useState } from 'react'
import { createFood, deleteFood, getFoodsByMenu, updateFood } from '../services/foodService'
import type { Food, FoodPayload } from '../types/food'

export function useFoods(menuId: string | null) {
  const [foods, setFoods] = useState<Food[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const loadFoods = useCallback(async () => {
    if (!menuId) return
    setLoading(true)
    setError(null)

    try {
      const data = await getFoodsByMenu(menuId)
      setFoods(data)
    } catch {
      setError('Não foi possível carregar os pratos deste cardápio.')
    } finally {
      setLoading(false)
    }
  }, [menuId])

  useEffect(() => {
    void loadFoods()
  }, [loadFoods])

  async function handleCreate(payload: Omit<FoodPayload, 'menuId'>) {
    if (!menuId) return
    await createFood({ ...payload, menuId })
    await loadFoods()
  }

  async function handleUpdate(foodId: string, payload: Omit<FoodPayload, 'menuId'>) {
    if (!menuId) return
    await updateFood(foodId, { ...payload, menuId })
    await loadFoods()
  }

  async function handleDelete(foodId: string) {
    await deleteFood(foodId)
    await loadFoods()
  }

  return {
    foods,
    loading,
    error,
    handleCreate,
    handleUpdate,
    handleDelete,
  }
}
