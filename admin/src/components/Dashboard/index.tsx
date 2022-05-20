import { useExpensesQuery } from '../../generated/graphql';

const Dashboard = () => {
  const { data, error, isFetching } = useExpensesQuery();

  if (error) return <div>err</div>;
  if (isFetching) return <div>loading</div>;
  return <div>{data?.allExpenses?.nodes?.map((exp) => exp.amount)}</div>;
};

export default Dashboard;
