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
    const json = atob(normalized)
    const data = JSON.parse(json) as { Sum?: string; sum?: string }
    return data.Sum ?? data.sum ?? null
  } catch {
    return null
  }
}
