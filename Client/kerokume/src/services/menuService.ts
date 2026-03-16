import { api } from '../lib/api'
import type { ApiListResponse, ApiSingleResponse } from '../types/api'
import type { Menu, MenuPayload } from '../types/menu'

export async function getMenusByRestaurant(restaurantId: string) {
  const response = await api.get<ApiListResponse<Menu>>(`/menus/${restaurantId}`)
  return response.data.data ?? []
}

export async function createMenu(payload: MenuPayload) {
  const response = await api.post<ApiSingleResponse<Menu>>('/menu', payload)
  return response.data.data
}

export async function updateMenu(menuId: string, payload: MenuPayload) {
  await api.put(`/menu/${menuId}`, payload)
}

export async function deleteMenu(menuId: string) {
  await api.delete(`/menu/${menuId}`)
}

export async function activateMenu(menuId: string) {
  await api.patch(`/menu/${menuId}`)
}
