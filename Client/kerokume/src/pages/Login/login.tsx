import toast, { Toaster } from 'react-hot-toast'
import './login.css'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'
import { useState } from 'react'

type LoginResponse = {
  data: {
    Token: string
    Message: string
  }
}

export const Login = () => {

  const nav = useNavigate()
  const [loading, setLoading] = useState(false)

  // Submit
  async function handlerSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()

    const BASE_URL = import.meta.env.VITE_BASE_URL
    const formData = new FormData(event.currentTarget)

    const email = formData.get("email") as string
    const password = formData.get("password") as string

    setLoading(true)

    try {

      const response = await axios.post<LoginResponse>(
        BASE_URL + "/login",
        { email, password }
      )

      const token = response.data.data.Token

      localStorage.setItem("token-login", token)

      toast.success("Login realizado com sucesso 👏", {
        duration: 3000,
        position: "top-center",
      })

      setTimeout(() => {
        nav('/my-menus')
      }, 1200)

    } catch (error: any) {

      console.error(error)

      toast.error(
        error?.response?.data?.message || "Erro ao fazer login"
      )

    } finally {
      setLoading(false)
    }
  }

  return (
    <main className="login-page">

      <Toaster position="top-center" />

      <section className="login-card">
        <div className="login-brand">
          <p className="login-badge">Bem-vindo ao KeroKume</p>
          <h1>Entrar na sua conta</h1>
          <p className="login-subtitle">
            Gerencie cardápios, pedidos e restaurantes em um só lugar.
          </p>
        </div>

        <form className="login-form" onSubmit={handlerSubmit}>

          <label htmlFor="email">E-mail</label>
          <input
            id="email"
            name="email"
            type="email"
            placeholder="voce@restaurante.com"
            required
          />

          <label htmlFor="password">Senha</label>
          <input
            id="password"
            name="password"
            type="password"
            placeholder="••••••••"
            required
          />

          <button type="submit" disabled={loading}>
            {loading ? "Entrando..." : "Entrar"}
          </button>

        </form>

        <p className="login-footer">
          Ainda não tem conta? <a href="/register">Criar agora</a>
        </p>
      </section>

    </main>
  )
}