import { createContext } from 'react'

type AuthUser = {
  name?: string
  email?: string
}

export type AuthContextValue = {
  user: AuthUser | null
  login: (userData: AuthUser) => void
  logout: () => void
  loading: boolean
}

export const AuthContext = createContext<AuthContextValue | null>(null)
