import { createBrowserRouter } from "react-router-dom"
import App from "./App"
import { Login } from "./pages/Login/login"
import { Register } from "./pages/Register/register"

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
  }
])