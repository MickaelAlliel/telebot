import { useAuthenticator } from '@aws-amplify/ui-react';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

export const Logout = () => {
  const { signOut } = useAuthenticator();
  const navigate = useNavigate();
  useEffect(() => {
    signOut();
    navigate('/');
  }, []);
  return <div>Signing out...</div>;
};
