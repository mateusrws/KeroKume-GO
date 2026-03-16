import { api } from '../lib/api'
import type { ApiListResponse, ApiSingleResponse } from '../types/api'
import type { Menu, MenuPayload } from '../types/menu'

type MenuApi = {
  id?: string
  Id?: string
  ID?: string
  name?: string
  Name?: string
  isActive?: boolean
  IsActive?: boolean
  restaurantId?: string
  RestaurantId?: string
}

function normalizeMenu(menu: MenuApi): Menu {
  return {
    id: menu.id ?? menu.Id ?? menu.ID ?? '',
    name: menu.name ?? menu.Name ?? 'Sem nome',
    isActive: menu.isActive ?? menu.IsActive ?? false,
    restaurantId: menu.restaurantId ?? menu.RestaurantId,
  }
}

export async function getMenusByRestaurant(restaurantId: string) {
  const response = await api.get<ApiListResponse<MenuApi>>(`/menus/${restaurantId}`)
  const rawMenus = response.data.data ?? []
  return rawMenus.map(normalizeMenu).filter((menu) => menu.id)
}

export async function createMenu(payload: MenuPayload) {
  const response = await api.post<ApiSingleResponse<MenuApi>>('/menu', payload)
  return response.data.data ? normalizeMenu(response.data.data) : undefined
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
