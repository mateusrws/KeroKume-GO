import axios from "axios"
import "./register.css"


export const Register = () => {
  
  const BASE_URL = import.meta.env.VITE_BASE_URL

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault()

    const formData = new FormData(event.currentTarget)

    const data = {
      name: formData.get("name") as string,
      email: formData.get("email") as string,
      password: formData.get("password") as string,
      description: formData.get("description") as string,
    }

    console.log(data)

    try {
      const response = await axios.post(BASE_URL+"/restaurant", data)

      console.log("Usuário criado:", response.data)

    } catch (error) {
      console.error("Erro ao registrar:", error)
    }
  }

  return (
    <main className="login-page">
      <section className="login-card">
        <div className="login-brand">
          <p className="login-badge">Bem-vindo ao KeroKume</p>
          <h1>Criar conta</h1>
          <p className="login-subtitle">
            Gerencie cardápios, pedidos e restaurantes em um só lugar.
          </p>
        </div>

        <form className="login-form" onSubmit={handleSubmit}>
          
          <label htmlFor="name">Nome</label>
          <input
            id="name"
            name="name"
            type="text"
            placeholder="KeroKume Restaurante"
            required
          />

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

          <label htmlFor="description">Descrição</label>
          <textarea
            id="description"
            name="description"
            placeholder="Digite uma descrição"
          />

          <button type="submit">Criar conta</button>
        </form>

        <p className="login-footer">
          Já tem conta? <a href="/login">Entrar</a>
        </p>
      </section>
    </main>
  )
}