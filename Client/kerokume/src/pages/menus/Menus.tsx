import './Menus.css'

type Menu = {
  title: string
  category: string
  items: number
  updatedAt: string
  status: 'Publicado' | 'Rascunho'
}

const menus: Menu[] = [
  { title: 'Almoço Executivo', category: 'Comida Brasileira', items: 14, updatedAt: 'Hoje, 11:30', status: 'Publicado' },
  { title: 'Jantar Premium', category: 'Massas e Carnes', items: 22, updatedAt: 'Ontem, 20:10', status: 'Publicado' },
  { title: 'Happy Hour', category: 'Bebidas & Petiscos', items: 18, updatedAt: 'Hoje, 17:00', status: 'Publicado' },
  { title: 'Sobremesas da Casa', category: 'Doces Artesanais', items: 9, updatedAt: 'Seg, 14:40', status: 'Rascunho' },
  { title: 'Delivery da Noite', category: 'Combos e Lanches', items: 16, updatedAt: 'Hoje, 18:15', status: 'Publicado' },
  { title: 'Menu de Fim de Semana', category: 'Especiais do Chef', items: 12, updatedAt: 'Dom, 09:55', status: 'Rascunho' },
]

export const RestaurantMenus = () => {
  return (
    <main className="menus-page">
      <section className="menus-shell">
        <header className="menus-header">
          <div>
            <p className="menus-badge">Painel do restaurante</p>
            <h1>Todos os cardápios</h1>
            <p className="menus-subtitle">Gerencie seus menus e publique atualizações em segundos.</p>
          </div>
          <button className="new-menu-btn">+ Novo cardápio</button>
        </header>

        <section className="menus-grid">
          {menus.map((menu) => (
           
          ))}
        </section>
      </section>
    </main>
  )
}