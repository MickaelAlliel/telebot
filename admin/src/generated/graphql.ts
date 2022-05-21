import { useQuery, UseQueryOptions } from 'react-query';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };

function fetcher<TData, TVariables>(query: string, variables?: TVariables) {
  return async (): Promise<TData> => {
    const res = await fetch(import.meta.env.VITE_GRAPHQL_ENDPOINT as string, {
    method: "POST",
    ...({"headers":{"Content-Type":"application/json"}}),
      body: JSON.stringify({ query, variables }),
    });

    const json = await res.json();

    if (json.errors) {
      const { message } = json.errors[0];

      throw new Error(message);
    }

    return json.data;
  }
}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  /**
   * A signed eight-byte integer. The upper big integer values are greater than the
   * max value for a JavaScript number. Therefore all big integers will be output as
   * strings and not numbers.
   */
  BigInt: any;
  /** A location in a connection that can be used for resuming pagination. */
  Cursor: any;
  /**
   * A point in time as described by the [ISO
   * 8601](https://en.wikipedia.org/wiki/ISO_8601) standard. May or may not include a timezone.
   */
  Datetime: string;
};

/** A filter to be used against BigInt fields. All fields are combined with a logical ‘and.’ */
export type BigIntFilter = {
  /** Not equal to the specified value, treating null like an ordinary value. */
  distinctFrom?: InputMaybe<Scalars['BigInt']>;
  /** Equal to the specified value. */
  equalTo?: InputMaybe<Scalars['BigInt']>;
  /** Greater than the specified value. */
  greaterThan?: InputMaybe<Scalars['BigInt']>;
  /** Greater than or equal to the specified value. */
  greaterThanOrEqualTo?: InputMaybe<Scalars['BigInt']>;
  /** Included in the specified list. */
  in?: InputMaybe<Array<Scalars['BigInt']>>;
  /** Is null (if `true` is specified) or is not null (if `false` is specified). */
  isNull?: InputMaybe<Scalars['Boolean']>;
  /** Less than the specified value. */
  lessThan?: InputMaybe<Scalars['BigInt']>;
  /** Less than or equal to the specified value. */
  lessThanOrEqualTo?: InputMaybe<Scalars['BigInt']>;
  /** Equal to the specified value, treating null like an ordinary value. */
  notDistinctFrom?: InputMaybe<Scalars['BigInt']>;
  /** Not equal to the specified value. */
  notEqualTo?: InputMaybe<Scalars['BigInt']>;
  /** Not included in the specified list. */
  notIn?: InputMaybe<Array<Scalars['BigInt']>>;
};

/** All input for the create `Expense` mutation. */
export type CreateExpenseInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `Expense` to be created by this mutation. */
  expense: ExpenseInput;
};

/** The output of our create `Expense` mutation. */
export type CreateExpensePayload = {
  __typename?: 'CreateExpensePayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `Expense` that was created by this mutation. */
  expense?: Maybe<Expense>;
  /** An edge for our `Expense`. May be used by Relay 1. */
  expenseEdge?: Maybe<ExpensesEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our create `Expense` mutation. */
export type CreateExpensePayloadExpenseEdgeArgs = {
  orderBy?: InputMaybe<Array<ExpensesOrderBy>>;
};

/** A filter to be used against Datetime fields. All fields are combined with a logical ‘and.’ */
export type DatetimeFilter = {
  /** Not equal to the specified value, treating null like an ordinary value. */
  distinctFrom?: InputMaybe<Scalars['Datetime']>;
  /** Equal to the specified value. */
  equalTo?: InputMaybe<Scalars['Datetime']>;
  /** Greater than the specified value. */
  greaterThan?: InputMaybe<Scalars['Datetime']>;
  /** Greater than or equal to the specified value. */
  greaterThanOrEqualTo?: InputMaybe<Scalars['Datetime']>;
  /** Included in the specified list. */
  in?: InputMaybe<Array<Scalars['Datetime']>>;
  /** Is null (if `true` is specified) or is not null (if `false` is specified). */
  isNull?: InputMaybe<Scalars['Boolean']>;
  /** Less than the specified value. */
  lessThan?: InputMaybe<Scalars['Datetime']>;
  /** Less than or equal to the specified value. */
  lessThanOrEqualTo?: InputMaybe<Scalars['Datetime']>;
  /** Equal to the specified value, treating null like an ordinary value. */
  notDistinctFrom?: InputMaybe<Scalars['Datetime']>;
  /** Not equal to the specified value. */
  notEqualTo?: InputMaybe<Scalars['Datetime']>;
  /** Not included in the specified list. */
  notIn?: InputMaybe<Array<Scalars['Datetime']>>;
};

/** All input for the `deleteExpenseById` mutation. */
export type DeleteExpenseByIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['BigInt'];
};

/** All input for the `deleteExpense` mutation. */
export type DeleteExpenseInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Expense` to be deleted. */
  nodeId: Scalars['ID'];
};

/** The output of our delete `Expense` mutation. */
export type DeleteExpensePayload = {
  __typename?: 'DeleteExpensePayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedExpenseId?: Maybe<Scalars['ID']>;
  /** The `Expense` that was deleted by this mutation. */
  expense?: Maybe<Expense>;
  /** An edge for our `Expense`. May be used by Relay 1. */
  expenseEdge?: Maybe<ExpensesEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our delete `Expense` mutation. */
export type DeleteExpensePayloadExpenseEdgeArgs = {
  orderBy?: InputMaybe<Array<ExpensesOrderBy>>;
};

export type Expense = Node & {
  __typename?: 'Expense';
  amount: Scalars['Float'];
  category: Scalars['String'];
  createdAt: Scalars['Datetime'];
  id: Scalars['BigInt'];
  method: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  ownerName: Scalars['String'];
  timestamp: Scalars['Datetime'];
};

/** A condition to be used against `Expense` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type ExpenseCondition = {
  /** Checks for equality with the object’s `amount` field. */
  amount?: InputMaybe<Scalars['Float']>;
  /** Checks for equality with the object’s `category` field. */
  category?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['BigInt']>;
  /** Checks for equality with the object’s `method` field. */
  method?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `ownerName` field. */
  ownerName?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `timestamp` field. */
  timestamp?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against `Expense` object types. All fields are combined with a logical ‘and.’ */
export type ExpenseFilter = {
  /** Filter by the object’s `amount` field. */
  amount?: InputMaybe<FloatFilter>;
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<ExpenseFilter>>;
  /** Filter by the object’s `category` field. */
  category?: InputMaybe<StringFilter>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<BigIntFilter>;
  /** Filter by the object’s `method` field. */
  method?: InputMaybe<StringFilter>;
  /** Negates the expression. */
  not?: InputMaybe<ExpenseFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<ExpenseFilter>>;
  /** Filter by the object’s `ownerName` field. */
  ownerName?: InputMaybe<StringFilter>;
  /** Filter by the object’s `timestamp` field. */
  timestamp?: InputMaybe<DatetimeFilter>;
};

/** An input for mutations affecting `Expense` */
export type ExpenseInput = {
  amount: Scalars['Float'];
  category: Scalars['String'];
  createdAt: Scalars['Datetime'];
  id?: InputMaybe<Scalars['BigInt']>;
  method: Scalars['String'];
  ownerName: Scalars['String'];
  timestamp: Scalars['Datetime'];
};

/** Represents an update to a `Expense`. Fields that are set will be updated. */
export type ExpensePatch = {
  amount?: InputMaybe<Scalars['Float']>;
  category?: InputMaybe<Scalars['String']>;
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['BigInt']>;
  method?: InputMaybe<Scalars['String']>;
  ownerName?: InputMaybe<Scalars['String']>;
  timestamp?: InputMaybe<Scalars['Datetime']>;
};

/** A connection to a list of `Expense` values. */
export type ExpensesConnection = {
  __typename?: 'ExpensesConnection';
  /** A list of edges which contains the `Expense` and cursor to aid in pagination. */
  edges: Array<ExpensesEdge>;
  /** A list of `Expense` objects. */
  nodes: Array<Expense>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Expense` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Expense` edge in the connection. */
export type ExpensesEdge = {
  __typename?: 'ExpensesEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Expense` at the end of the edge. */
  node: Expense;
};

/** Methods to use when ordering `Expense`. */
export enum ExpensesOrderBy {
  AmountAsc = 'AMOUNT_ASC',
  AmountDesc = 'AMOUNT_DESC',
  CategoryAsc = 'CATEGORY_ASC',
  CategoryDesc = 'CATEGORY_DESC',
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  MethodAsc = 'METHOD_ASC',
  MethodDesc = 'METHOD_DESC',
  Natural = 'NATURAL',
  OwnerNameAsc = 'OWNER_NAME_ASC',
  OwnerNameDesc = 'OWNER_NAME_DESC',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  TimestampAsc = 'TIMESTAMP_ASC',
  TimestampDesc = 'TIMESTAMP_DESC'
}

/** A filter to be used against Float fields. All fields are combined with a logical ‘and.’ */
export type FloatFilter = {
  /** Not equal to the specified value, treating null like an ordinary value. */
  distinctFrom?: InputMaybe<Scalars['Float']>;
  /** Equal to the specified value. */
  equalTo?: InputMaybe<Scalars['Float']>;
  /** Greater than the specified value. */
  greaterThan?: InputMaybe<Scalars['Float']>;
  /** Greater than or equal to the specified value. */
  greaterThanOrEqualTo?: InputMaybe<Scalars['Float']>;
  /** Included in the specified list. */
  in?: InputMaybe<Array<Scalars['Float']>>;
  /** Is null (if `true` is specified) or is not null (if `false` is specified). */
  isNull?: InputMaybe<Scalars['Boolean']>;
  /** Less than the specified value. */
  lessThan?: InputMaybe<Scalars['Float']>;
  /** Less than or equal to the specified value. */
  lessThanOrEqualTo?: InputMaybe<Scalars['Float']>;
  /** Equal to the specified value, treating null like an ordinary value. */
  notDistinctFrom?: InputMaybe<Scalars['Float']>;
  /** Not equal to the specified value. */
  notEqualTo?: InputMaybe<Scalars['Float']>;
  /** Not included in the specified list. */
  notIn?: InputMaybe<Array<Scalars['Float']>>;
};

/** The root mutation type which contains root level fields which mutate data. */
export type Mutation = {
  __typename?: 'Mutation';
  /** Creates a single `Expense`. */
  createExpense?: Maybe<CreateExpensePayload>;
  /** Deletes a single `Expense` using its globally unique id. */
  deleteExpense?: Maybe<DeleteExpensePayload>;
  /** Deletes a single `Expense` using a unique key. */
  deleteExpenseById?: Maybe<DeleteExpensePayload>;
  /** Updates a single `Expense` using its globally unique id and a patch. */
  updateExpense?: Maybe<UpdateExpensePayload>;
  /** Updates a single `Expense` using a unique key and a patch. */
  updateExpenseById?: Maybe<UpdateExpensePayload>;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateExpenseArgs = {
  input: CreateExpenseInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteExpenseArgs = {
  input: DeleteExpenseInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteExpenseByIdArgs = {
  input: DeleteExpenseByIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateExpenseArgs = {
  input: UpdateExpenseInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateExpenseByIdArgs = {
  input: UpdateExpenseByIdInput;
};

/** An object with a globally unique `ID`. */
export type Node = {
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
};

/** Information about pagination in a connection. */
export type PageInfo = {
  __typename?: 'PageInfo';
  /** When paginating forwards, the cursor to continue. */
  endCursor?: Maybe<Scalars['Cursor']>;
  /** When paginating forwards, are there more items? */
  hasNextPage: Scalars['Boolean'];
  /** When paginating backwards, are there more items? */
  hasPreviousPage: Scalars['Boolean'];
  /** When paginating backwards, the cursor to continue. */
  startCursor?: Maybe<Scalars['Cursor']>;
};

/** The root query type which gives access points into the data universe. */
export type Query = Node & {
  __typename?: 'Query';
  /** Reads and enables pagination through a set of `Expense`. */
  allExpenses?: Maybe<ExpensesConnection>;
  /** Reads a single `Expense` using its globally unique `ID`. */
  expense?: Maybe<Expense>;
  expenseById?: Maybe<Expense>;
  /** Fetches an object given its globally unique `ID`. */
  node?: Maybe<Node>;
  /** The root query type must be a `Node` to work well with Relay 1 mutations. This just resolves to `query`. */
  nodeId: Scalars['ID'];
  /**
   * Exposes the root query type nested one level down. This is helpful for Relay 1
   * which can only query top level fields if they are in a particular form.
   */
  query: Query;
};


/** The root query type which gives access points into the data universe. */
export type QueryAllExpensesArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<ExpenseCondition>;
  filter?: InputMaybe<ExpenseFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<ExpensesOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryExpenseArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryExpenseByIdArgs = {
  id: Scalars['BigInt'];
};


/** The root query type which gives access points into the data universe. */
export type QueryNodeArgs = {
  nodeId: Scalars['ID'];
};

/** A filter to be used against String fields. All fields are combined with a logical ‘and.’ */
export type StringFilter = {
  /** Not equal to the specified value, treating null like an ordinary value. */
  distinctFrom?: InputMaybe<Scalars['String']>;
  /** Not equal to the specified value, treating null like an ordinary value (case-insensitive). */
  distinctFromInsensitive?: InputMaybe<Scalars['String']>;
  /** Ends with the specified string (case-sensitive). */
  endsWith?: InputMaybe<Scalars['String']>;
  /** Ends with the specified string (case-insensitive). */
  endsWithInsensitive?: InputMaybe<Scalars['String']>;
  /** Equal to the specified value. */
  equalTo?: InputMaybe<Scalars['String']>;
  /** Equal to the specified value (case-insensitive). */
  equalToInsensitive?: InputMaybe<Scalars['String']>;
  /** Greater than the specified value. */
  greaterThan?: InputMaybe<Scalars['String']>;
  /** Greater than the specified value (case-insensitive). */
  greaterThanInsensitive?: InputMaybe<Scalars['String']>;
  /** Greater than or equal to the specified value. */
  greaterThanOrEqualTo?: InputMaybe<Scalars['String']>;
  /** Greater than or equal to the specified value (case-insensitive). */
  greaterThanOrEqualToInsensitive?: InputMaybe<Scalars['String']>;
  /** Included in the specified list. */
  in?: InputMaybe<Array<Scalars['String']>>;
  /** Included in the specified list (case-insensitive). */
  inInsensitive?: InputMaybe<Array<Scalars['String']>>;
  /** Contains the specified string (case-sensitive). */
  includes?: InputMaybe<Scalars['String']>;
  /** Contains the specified string (case-insensitive). */
  includesInsensitive?: InputMaybe<Scalars['String']>;
  /** Is null (if `true` is specified) or is not null (if `false` is specified). */
  isNull?: InputMaybe<Scalars['Boolean']>;
  /** Less than the specified value. */
  lessThan?: InputMaybe<Scalars['String']>;
  /** Less than the specified value (case-insensitive). */
  lessThanInsensitive?: InputMaybe<Scalars['String']>;
  /** Less than or equal to the specified value. */
  lessThanOrEqualTo?: InputMaybe<Scalars['String']>;
  /** Less than or equal to the specified value (case-insensitive). */
  lessThanOrEqualToInsensitive?: InputMaybe<Scalars['String']>;
  /** Matches the specified pattern (case-sensitive). An underscore (_) matches any single character; a percent sign (%) matches any sequence of zero or more characters. */
  like?: InputMaybe<Scalars['String']>;
  /** Matches the specified pattern (case-insensitive). An underscore (_) matches any single character; a percent sign (%) matches any sequence of zero or more characters. */
  likeInsensitive?: InputMaybe<Scalars['String']>;
  /** Equal to the specified value, treating null like an ordinary value. */
  notDistinctFrom?: InputMaybe<Scalars['String']>;
  /** Equal to the specified value, treating null like an ordinary value (case-insensitive). */
  notDistinctFromInsensitive?: InputMaybe<Scalars['String']>;
  /** Does not end with the specified string (case-sensitive). */
  notEndsWith?: InputMaybe<Scalars['String']>;
  /** Does not end with the specified string (case-insensitive). */
  notEndsWithInsensitive?: InputMaybe<Scalars['String']>;
  /** Not equal to the specified value. */
  notEqualTo?: InputMaybe<Scalars['String']>;
  /** Not equal to the specified value (case-insensitive). */
  notEqualToInsensitive?: InputMaybe<Scalars['String']>;
  /** Not included in the specified list. */
  notIn?: InputMaybe<Array<Scalars['String']>>;
  /** Not included in the specified list (case-insensitive). */
  notInInsensitive?: InputMaybe<Array<Scalars['String']>>;
  /** Does not contain the specified string (case-sensitive). */
  notIncludes?: InputMaybe<Scalars['String']>;
  /** Does not contain the specified string (case-insensitive). */
  notIncludesInsensitive?: InputMaybe<Scalars['String']>;
  /** Does not match the specified pattern (case-sensitive). An underscore (_) matches any single character; a percent sign (%) matches any sequence of zero or more characters. */
  notLike?: InputMaybe<Scalars['String']>;
  /** Does not match the specified pattern (case-insensitive). An underscore (_) matches any single character; a percent sign (%) matches any sequence of zero or more characters. */
  notLikeInsensitive?: InputMaybe<Scalars['String']>;
  /** Does not start with the specified string (case-sensitive). */
  notStartsWith?: InputMaybe<Scalars['String']>;
  /** Does not start with the specified string (case-insensitive). */
  notStartsWithInsensitive?: InputMaybe<Scalars['String']>;
  /** Starts with the specified string (case-sensitive). */
  startsWith?: InputMaybe<Scalars['String']>;
  /** Starts with the specified string (case-insensitive). */
  startsWithInsensitive?: InputMaybe<Scalars['String']>;
};

/** All input for the `updateExpenseById` mutation. */
export type UpdateExpenseByIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** An object where the defined keys will be set on the `Expense` being updated. */
  expensePatch: ExpensePatch;
  id: Scalars['BigInt'];
};

/** All input for the `updateExpense` mutation. */
export type UpdateExpenseInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** An object where the defined keys will be set on the `Expense` being updated. */
  expensePatch: ExpensePatch;
  /** The globally unique `ID` which will identify a single `Expense` to be updated. */
  nodeId: Scalars['ID'];
};

/** The output of our update `Expense` mutation. */
export type UpdateExpensePayload = {
  __typename?: 'UpdateExpensePayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `Expense` that was updated by this mutation. */
  expense?: Maybe<Expense>;
  /** An edge for our `Expense`. May be used by Relay 1. */
  expenseEdge?: Maybe<ExpensesEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our update `Expense` mutation. */
export type UpdateExpensePayloadExpenseEdgeArgs = {
  orderBy?: InputMaybe<Array<ExpensesOrderBy>>;
};

export type ExpensesQueryVariables = Exact<{ [key: string]: never; }>;


export type ExpensesQuery = { __typename?: 'Query', allExpenses?: { __typename?: 'ExpensesConnection', totalCount: number, nodes: Array<{ __typename?: 'Expense', id: any, amount: number, category: string, method: string, ownerName: string, timestamp: string, nodeId: string }>, pageInfo: { __typename?: 'PageInfo', endCursor?: any | null, hasNextPage: boolean, hasPreviousPage: boolean, startCursor?: any | null } } | null };


export const ExpensesDocument = `
    query Expenses {
  allExpenses {
    nodes {
      id
      amount
      category
      method
      ownerName
      timestamp
      nodeId
    }
    pageInfo {
      endCursor
      hasNextPage
      hasPreviousPage
      startCursor
    }
    totalCount
  }
}
    `;
export const useExpensesQuery = <
      TData = ExpensesQuery,
      TError = unknown
    >(
      variables?: ExpensesQueryVariables,
      options?: UseQueryOptions<ExpensesQuery, TError, TData>
    ) =>
    useQuery<ExpensesQuery, TError, TData>(
      variables === undefined ? ['Expenses'] : ['Expenses', variables],
      fetcher<ExpensesQuery, ExpensesQueryVariables>(ExpensesDocument, variables),
      options
    );