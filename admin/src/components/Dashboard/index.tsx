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
import { groupBy } from 'lodash';
import { useMemo } from 'react';
import { Bar } from 'react-chartjs-2';
import styled from 'styled-components';
import { useExpensesQuery } from '../../generated/graphql';

const Container = styled.div``;

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
  scales: {
    x: {
      stacked: true,
    },
    y: {
      stacked: true,
    },
  },
  maintainAspectRatio: false,
};

const Dashboard = () => {
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

    return { labels: Object.keys(expensesPerCategory), datasets };
  }, [data]);

  if (error) return <div>err</div>;
  if (isFetching) return <div>loading</div>;

  return (
    <Container>
      <Bar options={options} data={chartData} width={'100%'} height={'500px'} />
    </Container>
  );
};

export default Dashboard;
