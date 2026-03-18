const TOKEN_KEY = 'token-login'

export function getAuthToken() {
  return localStorage.getItem(TOKEN_KEY)
}

export function getRestaurantIdFromToken() {
  const token = getAuthToken()
  if (!token) return null

  const payload = token.split('.')[1]
  if (!payload) return null

  try {
    const normalized = payload.replace(/-/g, '+').replace(/_/g, '/')
    const padded = normalized.padEnd(Math.ceil(normalized.length / 4) * 4, '=')
    const json = atob(padded)
    const data = JSON.parse(json) as {
      Sum?: string
      sum?: string
      sub?: string
      id?: string
      restaurantId?: string
      restaurant_id?: string
    }

    return data.Sum ?? data.sum ?? data.restaurantId ?? data.restaurant_id ?? data.sub ?? data.id ?? null
  } catch {
    return null
  }
}
