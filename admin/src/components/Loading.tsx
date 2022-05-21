import { CircularProgress, Typography } from '@mui/material';
import styled from 'styled-components';

const Container = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
`;

export const Loading = () => {
  return (
    <Container>
      <CircularProgress size={'10rem'} />
      <Typography variant={'h5'} mt={4}>
        Loading...
      </Typography>
    </Container>
  );
};
