function App() {
  const categories = [
    'Comida Brasileira',
    'Massas Artesanais',
    'Hambúrgueres',
    'Sushi e Poke',
    'Marmitas'
  ]

  const highlights = [
    {
      title: 'Cardápio digital em minutos',
      description:
        'Crie, edite e publique seu cardápio sem complicação. Perfeito para restaurantes, bares e cafés.',
    },
    {
      title: 'Pedidos sem confusão',
      description:
        'Organize os pedidos por mesa, retirada ou delivery com atualizações em tempo real para sua equipe.',
    },
    {
      title: 'Visual que abre o apetite',
      description:
        'Fotos, descrições e destaques de pratos com uma paleta quente que remete ao clima de restaurante.',
    },
  ]

  return (
    <div className="flex min-h-screen items-center bg-gradient-to-b from-[var(--bg-from)] via-[var(--bg-via)] to-[var(--bg-to)] text-[var(--text-base)]">
      <div className="mx-auto w-full max-w-6xl px-6 py-12">
        <header className="mb-8 flex items-center justify-between">
          <div className="text-2xl font-black tracking-tight text-[var(--brand-900)]">KeroKume</div>
          <button className="rounded-full bg-[var(--brand-700)] px-5 py-2 text-sm font-semibold text-white transition hover:bg-[var(--brand-800)] cursor-pointer">
            Entrar
          </button>
        </header>

        <main className="grid items-center gap-10 lg:grid-cols-2">
          <section className="text-center lg:text-left">
            <p className="mb-3 inline-flex rounded-full bg-[var(--brand-100)] px-3 py-1 text-sm font-semibold text-[var(--brand-800)]">
              Cardápios online para restaurantes
            </p>
            <h1 className="text-4xl font-extrabold leading-tight text-[var(--brand-950)] sm:text-5xl">
              Seu menu digital com cara de casa cheia.
            </h1>
            <p className="mx-auto mt-5 max-w-xl text-lg text-[var(--text-muted)] lg:mx-0">
              Encante seus clientes com uma página inicial elegante, prática e pronta para converter visitas em pedidos.
            </p>

            <div className="mt-8 flex flex-wrap justify-center gap-3 lg:justify-start">
              {categories.map((category) => (
                <span
                  key={category}
                  className="rounded-full border border-[var(--accent-200)] bg-[var(--chip-bg)] px-4 py-2 text-sm font-medium text-[var(--accent-900)]"
                >
                  {category}
                </span>
              ))}
            </div>

            <div className="mt-10 flex flex-wrap justify-center gap-4 lg:justify-start">
              <button className="rounded-xl bg-[var(--brand-700)] px-6 py-3 font-semibold text-white shadow-lg shadow-red-900/20 transition hover:-translate-y-0.5 hover:bg-[var(--brand-800)] cursor-pointer">
                Criar meu cardápio
              </button>
              <button className="rounded-xl border border-[var(--brand-300)] bg-white px-6 py-3 font-semibold text-[var(--brand-800)] transition hover:bg-red-50 cursor-pointer">
                Ver demonstração
              </button>
            </div>
          </section>

          <section className="rounded-3xl border border-[var(--accent-100)] bg-[var(--card-bg)] p-6 shadow-2xl shadow-orange-900/10 backdrop-blur">
            <div className="mb-5 rounded-2xl bg-gradient-to-r from-[var(--brand-700)] to-[var(--accent-600)] p-5 text-white">
              <p className="text-sm uppercase tracking-widest text-orange-100">Restaurante destaque</p>
              <h2 className="mt-2 text-2xl font-bold">Cantina do Chef</h2>
              <p className="mt-1 text-orange-100">Lasanha da casa • Ravioli artesanal • Tiramisù</p>
            </div>

            <div className="space-y-4">
              {highlights.map((item) => (
                <article key={item.title} className="rounded-xl border border-[var(--accent-100)] bg-amber-50/70 p-4">
                  <h3 className="font-bold text-[var(--brand-900)]">{item.title}</h3>
                  <p className="mt-1 text-sm text-[var(--text-muted)]">{item.description}</p>
                </article>
              ))}
            </div>
          </section>
        </main>
      </div>
    </div>
  )
}

export default App