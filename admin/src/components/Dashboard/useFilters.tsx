import { uniq } from 'lodash';
import { useMemo, useState } from 'react';
import { ExpensesQuery } from '../../generated/graphql';

type UseFiltersFn = (
  timeframe: (string | null)[],
  data?: ExpensesQuery
) => {
  filterCategory: string | null;
  setFilterCategory: (newValue: string | null) => void;
  filterPaymentMethod: string | null;
  setFilterPaymentMethod: (newValue: string | null) => void;
  filterOwner: string | null;
  setFilterOwner: (newValue: string | null) => void;
  categoriesOptions: string[];
  paymentMethodsOptions: string[];
  ownersOptions: string[];

  startDate: string | null;
  setStartDate: (newValue: string | null) => void;
  endDate: string | null;
  setEndDate: (newValue: string | null) => void;
};

export const useFilters: UseFiltersFn = (timeframe, data) => {
  const [filterCategory, setFilterCategory] = useState<string | null>(null);
  const [filterPaymentMethod, setFilterPaymentMethod] = useState<string | null>(
    null
  );
  const [filterOwner, setFilterOwner] = useState<string | null>(null);
  const [startDate, setStartDate] = useState<string | null>(timeframe[0]);
  const [endDate, setEndDate] = useState<string | null>(timeframe[1]);

  const categoriesOptions = useMemo(() => {
    return uniq(data?.expenses?.nodes.map((exp) => exp.category));
  }, [data]);

  const paymentMethodsOptions = useMemo(() => {
    return uniq(data?.expenses?.nodes.map((exp) => exp.method));
  }, [data]);

  const ownersOptions = useMemo(() => {
    return uniq(data?.expenses?.nodes.map((exp) => exp.ownerName));
  }, [data]);

  return {
    filterCategory,
    setFilterCategory,
    categoriesOptions,
    filterPaymentMethod,
    setFilterPaymentMethod,
    paymentMethodsOptions,
    filterOwner,
    setFilterOwner,
    ownersOptions,
    startDate,
    setStartDate,
    endDate,
    setEndDate,
  };
};
