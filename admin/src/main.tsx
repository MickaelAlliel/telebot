import { Authenticator } from '@aws-amplify/ui-react';
import React from 'react';
import ReactDOM from 'react-dom/client';
import { ProtectedAppRouter } from './App';
import './index.css';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Authenticator.Provider>
      <ProtectedAppRouter />
    </Authenticator.Provider>
  </React.StrictMode>
);
