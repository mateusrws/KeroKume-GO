import './login.css'

export const Login = () => {
  return (
    <main className="login-page">
      <section className="login-card">
        <div className="login-brand">
          <p className="login-badge">Bem-vindo ao KeroKume</p>
          <h1>Entrar na sua conta</h1>
          <p className="login-subtitle">Gerencie cardápios, pedidos e restaurantes em um só lugar.</p>
        </div>

        <form className="login-form">
          <label htmlFor="email">E-mail</label>
          <input id="email" type="email" placeholder="voce@restaurante.com" />

          <label htmlFor="password">Senha</label>
          <input id="password" type="password" placeholder="••••••••" />

          <button type="submit">Entrar</button>
        </form>

        <p className="login-footer">Ainda não tem conta? <a href="#">Criar agora</a></p>
      </section>
    </main>
  )
}