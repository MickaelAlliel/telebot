import { Chance } from 'chance';
import { ChartData } from 'chart.js';
import { groupBy } from 'lodash';
import { useMemo } from 'react';
import { useExpensesQuery } from '../../generated/graphql';

export const useGroupedExpenses = (): {
  chartData: ChartData<'bar', { x: string; y: number }[]>;
  error: unknown;
  isFetching: boolean;
} => {
  const { data, error, isFetching } = useExpensesQuery();

  const chartData = useMemo(() => {
    if (!data?.allExpenses?.nodes) return { labels: [], datasets: [] };
    const expensesPerMethod = groupBy(data.allExpenses.nodes, 'method');
    const expensesPerCategory = groupBy(data.allExpenses.nodes, 'category');
    const expensesPerMethodPerCategory = Object.entries(expensesPerMethod)
      .map(([key, group]) => {
        return { [key]: groupBy(group, 'category') };
      })
      .reduce((obj, item) => Object.assign(obj, item));

    const datasets = Object.entries(expensesPerMethodPerCategory).map(
      ([label, group]) => {
        return {
          label,
          stack: 'bar',
          data: Object.entries(group).map(([key, value]) => ({
            x: key,
            y: value.reduce(
              (previousValue, currentValue) =>
                previousValue + currentValue.amount,
              0
            ),
          })),
          backgroundColor: new Chance().color({
            format: 'hex',
            casing: 'upper',
          }),
        };
      }
    );

    return {
      labels: Object.keys(expensesPerCategory),
      datasets,
    };
  }, [data]);

  return { chartData, error, isFetching };
};
