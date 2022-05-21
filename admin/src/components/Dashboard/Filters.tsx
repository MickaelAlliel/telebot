import { Autocomplete, Button, TextField } from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers';
import { addDays, formatISO, subSeconds } from 'date-fns';
import { useCallback, useEffect, useMemo } from 'react';
import styled from 'styled-components';
import { Expense, ExpensesQuery } from '../../generated/graphql';
import { useFilters } from './useFilters';

const FiltersContainer = styled.div`
  display: flex;
  flex-direction: row;
  margin-bottom: 50px;
`;

const FilterAutocomplete = styled(Autocomplete)`
  width: 14rem;
  margin: 0 1rem;
`;

export const Filters: React.FC<{
  data?: ExpensesQuery;
  setFilteredData: (value: Omit<Expense, 'createdAt'>[]) => void;
  timeframe: (string | null)[];
  setTimeframe: (value: (string | null)[]) => void;
}> = ({ data, setFilteredData, timeframe, setTimeframe }) => {
  const {
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
  } = useFilters(timeframe, data);

  const onClearFilters = useCallback(() => {
    setFilterCategory(null);
    setFilterOwner(null);
    setFilterPaymentMethod(null);
  }, [setFilterCategory, setFilterOwner, setFilterPaymentMethod]);

  const filteredData: Omit<Expense, 'createdAt'>[] = useMemo(() => {
    const expenses = data?.allExpenses?.nodes;
    if (!expenses) return [];
    let filtered = expenses;
    if (filterCategory !== null) {
      filtered = filtered.filter((exp) => filterCategory === exp.category);
      console.table(filtered);
    }
    if (filterPaymentMethod !== null) {
      filtered = filtered.filter((exp) => filterPaymentMethod === exp.method);
    }
    if (filterOwner !== null) {
      filtered = filtered.filter((exp) => filterOwner === exp.ownerName);
    }
    return filtered;
  }, [
    data?.allExpenses?.nodes,
    filterCategory,
    filterPaymentMethod,
    filterOwner,
  ]);

  useEffect(
    () => setFilteredData(filteredData),
    [filteredData, setFilteredData]
  );

  useEffect(() => {
    setTimeframe([startDate, endDate]);
  }, [startDate, endDate, setTimeframe]);

  return (
    <FiltersContainer>
      <FilterAutocomplete
        options={categoriesOptions}
        value={filterCategory}
        onChange={(event, newValue) =>
          setFilterCategory(newValue as string | null)
        }
        inputValue={filterCategory ?? ''}
        onInputChange={(event, newInputValue) =>
          setFilterCategory(newInputValue)
        }
        renderInput={(params) => <TextField {...params} label="Category" />}
      />
      <FilterAutocomplete
        options={paymentMethodsOptions}
        value={filterPaymentMethod}
        onChange={(event, newValue) =>
          setFilterPaymentMethod(newValue as string | null)
        }
        inputValue={filterPaymentMethod ?? ''}
        onInputChange={(event, newInputValue) =>
          setFilterPaymentMethod(newInputValue)
        }
        renderInput={(params) => (
          <TextField {...params} label="Payment Method" />
        )}
      />
      <FilterAutocomplete
        options={ownersOptions}
        value={filterOwner}
        onChange={(event, newValue) =>
          setFilterOwner(newValue as string | null)
        }
        inputValue={filterOwner ?? ''}
        onInputChange={(event, newInputValue) => setFilterOwner(newInputValue)}
        renderInput={(params) => <TextField {...params} label="Owner" />}
      />
      <DatePicker
        onChange={(value) => {
          const typedValue = value as Date | null;
          if (typedValue === null) {
            setStartDate(null);
            return;
          }
          setStartDate(formatISO(typedValue));
        }}
        value={startDate}
        renderInput={(params) => <TextField {...params} label="From" />}
      />
      <DatePicker
        onChange={(value) => {
          const typedValue = value as Date | null;
          if (typedValue === null) {
            setEndDate(null);
            return;
          }
          let d = addDays(typedValue, 1);
          d = subSeconds(typedValue, 1);
          setEndDate(formatISO(d));
        }}
        value={endDate}
        renderInput={(params) => <TextField {...params} label="Until" />}
      />
      <Button onClick={onClearFilters}>CLEAR</Button>
    </FiltersContainer>
  );
};
