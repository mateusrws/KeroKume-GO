import { api } from '../lib/api'
import type { ApiListResponse } from '../types/api'
import type { Food, FoodPayload } from '../types/food'

type FoodApi = {
  id?: string
  Id?: string
  ID?: string
  name?: string
  Name?: string
  description?: string
  Description?: string
  price?: number
  Price?: number
  foodCategory?: Food['foodCategory']
  FoodCategory?: Food['foodCategory']
  isAvailable?: boolean
  IsAvailable?: boolean
  pathImg?: string
  PathImg?: string
}

function normalizeFood(food: FoodApi): Food {
  return {
    id: food.id ?? food.Id ?? food.ID ?? '',
    name: food.name ?? food.Name ?? 'Sem nome',
    description: food.description ?? food.Description ?? '',
    price: food.price ?? food.Price ?? 0,
    foodCategory: food.foodCategory ?? food.FoodCategory ?? 'COMIDA',
    isAvailable: food.isAvailable ?? food.IsAvailable ?? true,
    pathImg: food.pathImg ?? food.PathImg,
  }
}

export async function getFoodsByMenu(menuId: string) {
  const response = await api.get<ApiListResponse<FoodApi>>(`/foods/${menuId}`, {
    params: { menuId },
  })
  const rawFoods = response.data.data ?? []
  return rawFoods.map(normalizeFood).filter((food) => food.id)
}

export async function createFood(payload: FoodPayload) {
  await api.post('/foods', payload, {
    params: { menuId: payload.menuId },
  })
}

export async function updateFood(foodId: string, payload: FoodPayload) {
  await api.put(`/foods/${foodId}`, payload, {
    params: { menuId: payload.menuId },
  })
}

export async function deleteFood(foodId: string, menuId: string) {
  await api.delete(`/foods/${foodId}`, {
    params: { menuId },
    data: { menuId },
  })
}
