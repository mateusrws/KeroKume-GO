import { Navigate, Outlet } from 'react-router-dom';

export const ProtectedRoute = () => {
  const token = localStorage.getItem("token-login");

  // Se não houver token, redireciona para o login
  if (!token) {
    return <Navigate to="/login" replace />;
  }

  // Se houver token, renderiza a página que o usuário tentou acessar
  return <Outlet />;
};
