export type Menu = {
  id: string
  name: string
  isActive?: boolean
  restaurantId?: string
}

export type MenuPayload = {
  name: string
  restaurantId: string
}
