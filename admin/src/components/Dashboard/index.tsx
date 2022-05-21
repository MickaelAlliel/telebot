import { Chance } from 'chance';
import {
  BarElement,
  CategoryScale,
  Chart as ChartJS,
  Legend,
  LinearScale,
  Title,
  Tooltip,
} from 'chart.js';
import ca from 'date-fns/esm/locale/ca/index.js';
import { filter, flatMap, groupBy, map, method } from 'lodash';
import { useMemo } from 'react';
import { Bar } from 'react-chartjs-2';
import { useExpensesQuery } from '../../generated/graphql';

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

export const options = {
  plugins: {
    title: {
      display: true,
      text: 'Expenses',
    },
  },
  responsive: true,
  interaction: {
    mode: 'index' as const,
    intersect: false,
  },
  scales: {
    x: {
      stacked: true,
    },
    y: {
      stacked: true,
    },
  },
};

const Dashboard = () => {
  const { data, error, isFetching } = useExpensesQuery();

  const chartData = useMemo(() => {
    if (!data?.allExpenses?.nodes) return { labels: [], datasets: [] };
    const expensesPerCategory = groupBy(data.allExpenses.nodes, 'category');
    const expensesPerCategoryPerMethod = Object.entries(expensesPerCategory)
      .map(([key, group]) => {
        return { [key]: groupBy(group, 'method') };
      })
      .reduce((obj, item) => Object.assign(obj, item));

    console.table(expensesPerCategoryPerMethod);

    const datasets = Object.entries(expensesPerCategoryPerMethod).map(
      ([label, group]) => {
        return {
          label,
          stack: label,
          data: Object.entries(group).map(([key, value]) => ({
            x: label,
            y: value.reduce(
              (previousValue, currentValue) =>
                previousValue + currentValue.amount,
              0
            ),
          })),
        };
      }
    );

    console.table(datasets);

    // const datasets = map(expensesPerMethodPerCategory, (methodGroup) => {
    //   console.info('AA');
    //   console.table(methodGroup);
    //   return {
    //     label,
    //     data: expenses.map((e) => ({
    //       x: label,
    //       y: e.amount,
    //     })),
    //     backgroundColor: new Chance().color(),
    //   };
    // });
    return { labels: Object.keys(expensesPerCategory), datasets };
  }, [data]);

  if (error) return <div>err</div>;
  if (isFetching) return <div>loading</div>;

  return <Bar options={options} data={chartData} />;
};

export default Dashboard;
