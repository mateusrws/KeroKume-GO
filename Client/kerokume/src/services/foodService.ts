import { api } from '../lib/api'
import type { ApiListResponse } from '../types/api'
import type { Food, FoodPayload } from '../types/food'

export async function getFoodsByMenu(menuId: string) {
  const response = await api.get<ApiListResponse<Food>>(`/foods/${menuId}`)
  return response.data.data ?? []
}

export async function createFood(payload: FoodPayload) {
  await api.post('/foods', payload)
}

export async function updateFood(foodId: string, payload: FoodPayload) {
  await api.put(`/foods/${foodId}`, payload)
}

export async function deleteFood(foodId: string) {
  await api.delete(`/foods/${foodId}`)
}
