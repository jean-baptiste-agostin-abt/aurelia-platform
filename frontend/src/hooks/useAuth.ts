import { useState } from 'react';

export default function useAuth() {
  const [token, setToken] = useState<string | null>(localStorage.getItem('token'));

  const saveToken = (t: string) => {
    localStorage.setItem('token', t);
    setToken(t);
  };

  return { token, saveToken };
}
