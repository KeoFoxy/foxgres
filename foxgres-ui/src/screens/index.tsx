import React, { FC } from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { styled } from '@mui/material/styles';

const LoginCard = styled(Card)({
  width: 500,
  margin: 'auto',
  marginTop: '10%',
  padding: '20px',
  background: '#FFF',
});

const LoginButton = styled(Button)({
  marginTop: '20px',
});

const StyledTextField = styled(TextField)({
  marginBottom: '10px',
});

const Login: FC = () => {
  return (
    <Container component="main" maxWidth="md">
      <LoginCard>
        <CardContent>
          <Box
            sx={{
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
            }}
          >
            <Typography component="h1" variant="h5">
              Вход в систему
            </Typography>
            <StyledTextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              id="email"
              label="Логин"
              name="email"
              autoComplete="email"
              autoFocus
            />
            <StyledTextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              name="password"
              label="Пароль"
              type="password"
              id="password"
              autoComplete="current-password"
            />
            <LoginButton
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
            >
              Войти
            </LoginButton>
          </Box>
        </CardContent>
      </LoginCard>
    </Container>
  );
}

export default Login;
