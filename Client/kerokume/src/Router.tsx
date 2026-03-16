import { createBrowserRouter } from 'react-router-dom'
import App from './App'
import { Login } from './pages/Login/login'
import { Register } from './pages/Register/register'
import { NotFound } from './pages/NotFoundPage/NotFound'
import { ProtectedRoute } from '../auth/ProviderRoute'
import { RestaurantMenus } from './pages/menus/Menus'
import { MenuFoodsPage } from './pages/menuFoods/MenuFoodsPage'
import { PublicMenuPage } from './pages/publicMenu/PublicMenuPage'

export const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
  },
  {
    path: '/login',
    element: <Login />,
  },
  {
    path: '/register',
    element: <Register />,
  },
  {
    path: '/public-menu/:menuId',
    element: <PublicMenuPage />,
  },
  {
    element: <ProtectedRoute />,
    children: [
      {
        path: '/my-menus',
        element: <RestaurantMenus />,
      },
      {
        path: '/my-menus/:menuId/foods',
        element: <MenuFoodsPage />,
      },
    ],
  },
  {
    path: '*',
    element: <NotFound />,
  },
])
