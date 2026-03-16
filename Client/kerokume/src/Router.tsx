import { createBrowserRouter } from "react-router-dom"
import App from "./App"
import { Login } from "./pages/Login/login"
import { Register } from "./pages/Register/register"
import { NotFound } from "./pages/NotFoundPage/NotFound"
import { ProtectedRoute } from "../auth/ProviderRoute"
import { RestaurantMenus } from "./pages/menus/Menus"

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/register",
    element: <Register />,
  },
  {
    path: "*",
    element: <NotFound />,
  },
  {
    element: <ProtectedRoute />,
    children: [
      {
        path: "/my-menus",
        element: <RestaurantMenus/>
      },
    ]
  }
])