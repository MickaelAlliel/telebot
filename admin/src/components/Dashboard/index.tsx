import styled from 'styled-components';
import { StackedBarChart } from './StackedBarChart';
import { useGroupedExpenses } from './useGroupedExpenses';

const Container = styled.div`
  display: flex;
`;

const Dashboard = () => {
  const { chartData, error, isFetching } = useGroupedExpenses();

  if (error) return <div>err</div>;
  if (isFetching) return <div>loading</div>;

  return (
    <Container>
      <StackedBarChart title={'Expense'} data={chartData} />
    </Container>
  );
};

export default Dashboard;
