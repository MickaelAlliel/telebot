import { Chance } from 'chance';
import { groupBy } from 'lodash';
import { Expense } from '../../generated/graphql';
import { StackedBarChartDataType } from './StackedBarChart';

export const groupedExpenses = (
  data: Omit<Expense, 'createdAt'>[]
): StackedBarChartDataType => {
  if (!data) return { labels: [], datasets: [] };
  const expensesPerMethod = groupBy(data, 'method');
  const expensesPerCategory = groupBy(data, 'category');
  const expensesPerMethodPerCategory = Object.entries(expensesPerMethod)
    .map(([key, group]) => {
      return { [key]: groupBy(group, 'category') };
    })
    .reduce((obj, item) => Object.assign(obj, item), {});

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
};
