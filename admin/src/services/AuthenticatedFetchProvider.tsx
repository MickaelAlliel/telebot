import { useAuthenticator } from '@aws-amplify/ui-react';

export const authedFetcher = <TData, TVariables>(
  query: string,
  options?: RequestInit['headers']
): ((variables?: TVariables) => Promise<TData>) => {
  const { user } = useAuthenticator();
  return async (variables?: TVariables): Promise<TData> => {
    const res = await fetch(import.meta.env.VITE_GRAPHQL_ENDPOINT as string, {
      method: 'POST',
      ...{
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${user
            .getSignInUserSession()
            ?.getIdToken()
            .getJwtToken()}`,
        },
      },
      ...options,
      body: JSON.stringify({ query, variables }),
    });

    const json = await res.json();

    if (json.errors) {
      const { message } = json.errors[0];
      throw new Error(message);
    }
    return json.data;
  };
};
