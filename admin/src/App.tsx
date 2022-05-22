import { Authenticator } from '@aws-amplify/ui-react';
import '@aws-amplify/ui-react/styles.css'; // default theme
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { Amplify } from 'aws-amplify';
import { QueryClient, QueryClientProvider } from 'react-query';
import { ReactQueryDevtools } from 'react-query/devtools';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Logout } from './components/Auth/Logout';
import Dashboard from './components/Dashboard';

import styled from 'styled-components';
import awsExports from './aws-exports';
import { CustomAppBar } from './components/AppBar';
Amplify.configure(awsExports);

const queryClient = new QueryClient();

const StyledAuthenticator = styled(Authenticator)`
  height: 100%;
`;

export const ProtectedAppRouter = () => {
  return (
    <StyledAuthenticator hideSignUp={true}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<App />}></Route>
          <Route path="/logout" element={<Logout />} />
        </Routes>
      </BrowserRouter>
    </StyledAuthenticator>
  );
};

export const App = () => {
  return (
    <LocalizationProvider dateAdapter={AdapterDateFns}>
      <QueryClientProvider client={queryClient}>
        <CustomAppBar />
        <Dashboard />
        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider>
    </LocalizationProvider>
  );
};
