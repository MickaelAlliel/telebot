import { format, subDays } from 'date-fns';
import { useEffect, useState } from 'react';
import styled from 'styled-components';
import { Expense, useExpensesQuery } from '../../generated/graphql';
import { Loading } from '../Loading';
import { Filters } from './Filters';
import { groupedExpenses } from './groupedExpenses';
import { StackedBarChart, StackedBarChartDataType } from './StackedBarChart';

const Container = styled.div`
  margin: 50px;
  display: flex;
  flex-direction: column;
`;

const ChartContainer = styled.div`
  height: 100%;
  width: 100%;
`;

const Dashboard = () => {
  const [chartData, setChartData] = useState<StackedBarChartDataType | null>(
    null
  );
  const [timeframe, setTimeframe] = useState<(string | null)[]>([
    format(subDays(new Date(), 7), 'yyyy-MM-dd'),
  ]);
  const { data, error, isFetching } = useExpensesQuery({
    from: timeframe[0],
    to: timeframe[1] ?? undefined,
  });
  const [filteredData, setFilteredData] = useState<
    Omit<Expense, 'createdAt'>[]
  >([]);

  useEffect(() => {
    setChartData(groupedExpenses(filteredData));
  }, [filteredData]);

  if (error) return <div>Error...</div>;
  if (isFetching) return <Loading />;

  return (
    <Container>
      <Filters
        data={data}
        setFilteredData={setFilteredData}
        timeframe={timeframe}
        setTimeframe={setTimeframe}
      />
      <ChartContainer>
        {chartData && <StackedBarChart title={'Expense'} data={chartData} />}
      </ChartContainer>
    </Container>
  );
};

export default Dashboard;
