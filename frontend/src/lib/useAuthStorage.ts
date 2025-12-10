import { useCallback, useEffect, useState } from 'react';

export default function useAuthStorage() {
  const getStorage = useCallback(() => {
    if (typeof window === 'undefined') return null;

    try {
      const stored = sessionStorage.getItem('token');
      return stored;
    } catch {
      return null;
    }
  }, []);

  const [token, setToken] = useState(getStorage);

  useEffect(() => {
    if (typeof window === 'undefined') return;

    const nowToken = sessionStorage.getItem('token');
    if (token == nowToken) return;

    if (token) {
      sessionStorage.setItem('token', token);
    } else {
      sessionStorage.removeItem('token');
    }
  }, [token]);

  return { token, setToken };
}
