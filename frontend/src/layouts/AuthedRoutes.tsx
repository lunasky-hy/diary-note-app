import useAuthStorage from '../lib/useAuthStorage';
import { Navigate, Outlet } from 'react-router';

export default function AuthedRoutes() {
  const { token } = useAuthStorage();
  return token ? <Outlet /> : <Navigate to="/auth/signin" replace />;
}
