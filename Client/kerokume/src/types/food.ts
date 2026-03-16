export type FoodCategory = 'COMIDA' | 'BEBIDA'

export type Food = {
  id: string
  name: string
  description: string
  price: number
  foodCategory: FoodCategory
  isAvailable: boolean
  pathImg?: string
}

export type FoodPayload = {
  name: string
  description: string
  price: number
  foodCategory: FoodCategory
  menuId: string
}
