import { QueryClient, QueryClientProvider } from 'react-query';
import { ReactQueryDevtools } from 'react-query/devtools';
import Dashboard from './components/Dashboard';

const queryClient = new QueryClient();

export const App = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <Dashboard></Dashboard>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
};
