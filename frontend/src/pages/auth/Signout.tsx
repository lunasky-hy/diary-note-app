import { useEffect } from 'react';
import useAuthStorage from '../../lib/useAuthStorage';
import { Navigate } from 'react-router';

export default function Signout() {
  const { setToken } = useAuthStorage();

  useEffect(() => {
    setToken(null);
  }, [setToken]);

  return <Navigate to="/auth/signin" />;
}
