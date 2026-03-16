export const menuCard = () => {
  <article key={menu.title} className="menu-card">
    <div className="menu-card-top">
      <p className="menu-category">{menu.category}</p>
      <span className={`menu-status ${menu.status === 'Publicado' ? 'is-published' : 'is-draft'}`}>
        {menu.status}
      </span>
    </div>
 
    <h2>{menu.title}</h2>
 
    <div className="menu-meta">
      <p>{menu.items} itens</p>
      <p>Atualizado: {menu.updatedAt}</p>
    </div>
 
    <button className="menu-open-btn">Abrir menu</button>
  </article>
}