import { useState } from 'react'
import type { PropsWithChildren } from 'react'
import { AuthContext } from './context'

type AuthUser = {
  name?: string
  email?: string
}

const initialUser = (() => {
  const storedUser = localStorage.getItem('user')
  return storedUser ? (JSON.parse(storedUser) as AuthUser) : null
})()

export const AuthProvider = ({ children }: PropsWithChildren) => {
  const [user, setUser] = useState<AuthUser | null>(initialUser)

  const login = (userData: AuthUser) => {
    setUser(userData)
    localStorage.setItem('user', JSON.stringify(userData))
  }

  const logout = () => {
    setUser(null)
    localStorage.removeItem('user')
  }

  return <AuthContext.Provider value={{ user, login, logout, loading: false }}>{children}</AuthContext.Provider>
}
