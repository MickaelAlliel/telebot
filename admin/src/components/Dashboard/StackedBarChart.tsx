import {
  BarElement,
  CategoryScale,
  Chart as ChartJS,
  ChartData,
  ChartOptions,
  Legend,
  LinearScale,
  Title,
  Tooltip,
} from 'chart.js';
import { Bar } from 'react-chartjs-2';

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

export const StackedBarChart: React.FC<{
  title: string;
  data: ChartData<'bar', { x: string; y: number }[]>;
  height?: string;
}> = ({ title, data, height }) => {
  const options: ChartOptions<'bar'> = {
    plugins: {
      title: {
        display: true,
        text: title,
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
    interaction: {
      mode: 'x',
    },
  };

  return (
    <Bar
      options={options}
      data={data}
      width={'100%'}
      height={height ?? '500px'}
    />
  );
};
