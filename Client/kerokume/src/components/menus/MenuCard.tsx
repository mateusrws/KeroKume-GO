import { Link } from 'react-router-dom'
import type { Menu } from '../../types/menu'

type MenuCardProps = {
  menu: Menu
  onEdit: (menu: Menu) => void
  onDelete: (menuId: string) => void
  onActivate: (menuId: string) => void
}

export function MenuCard({ menu, onEdit, onDelete, onActivate }: MenuCardProps) {
  return (
    <article className="rounded-2xl border border-[var(--accent-100)] bg-[var(--card-bg)] p-4 shadow-md">
      <div className="mb-4 flex items-center justify-between gap-2">
        <h2 className="text-lg font-bold text-[var(--brand-900)]">{menu.name}</h2>
        <span className={`rounded-full px-2 py-1 text-xs font-bold ${menu.isActive ? 'bg-green-100 text-green-700' : 'bg-orange-100 text-orange-700'}`}>
          {menu.isActive ? 'Ativo' : 'Inativo'}
        </span>
      </div>

      <div className="grid grid-cols-2 gap-2 text-sm">
        <button className="rounded-lg border border-[var(--brand-300)] px-3 py-2 font-semibold text-[var(--brand-800)]" onClick={() => onEdit(menu)}>
          Editar
        </button>
        <button className="rounded-lg border border-red-200 px-3 py-2 font-semibold text-red-700" onClick={() => onDelete(menu.id)}>
          Excluir
        </button>
        <button className="rounded-lg border border-emerald-200 px-3 py-2 font-semibold text-emerald-700" onClick={() => onActivate(menu.id)}>
          Ativar
        </button>
        <Link className="rounded-lg border border-[var(--accent-200)] px-3 py-2 text-center font-semibold text-[var(--accent-900)]" to={`/my-menus/${menu.id}/foods`} state={{ menuName: menu.name }}>
          Gerenciar pratos
        </Link>
      </div>
    </article>
  )
}
