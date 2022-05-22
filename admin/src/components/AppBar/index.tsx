import { useAuthenticator } from '@aws-amplify/ui-react';
import { Paid } from '@mui/icons-material';
import { AppBar, Box, Button, Toolbar, Typography } from '@mui/material';
import { useNavigate } from 'react-router-dom';

export const CustomAppBar = () => {
  const navigate = useNavigate();
  const { user } = useAuthenticator();
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position={'static'} color={'secondary'}>
        <Toolbar>
          <Paid color={'inherit'} fontSize={'large'} />
          <Typography flexGrow={1} variant="h6">
            PeachExpense
          </Typography>
          <Typography flexGrow={1}>Hello, {user.attributes?.email}!</Typography>
          <Button onClick={() => navigate('/logout')} color={'inherit'}>
            Sign out
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
};
