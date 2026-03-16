import './NotFound.css'

export const NotFound = () => {
  return (
    <main className="not-found-page">
      <section className="not-found-card">
        <p className="not-found-badge">Erro de navegação</p>
        <p className="not-found-code">404</p>
        <h1>Página não encontrada</h1>
        <p className="not-found-description">
          A página que você tentou acessar não existe ou foi movida. Volte para continuar gerenciando seus
          cardápios.
        </p>

        <div className="not-found-actions">
          <a href="#" className="primary-action">
            Ir para início
          </a>
          <a href="#" className="secondary-action">
            Falar com suporte
          </a>
        </div>
      </section>
    </main>
  )
}