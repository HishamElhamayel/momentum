import { StatusBar } from 'expo-status-bar';
import { Text } from 'react-native';

import { Container } from 'app/components/Container';
import Register from 'app/screen/Register';
import './global.css';

export default function App() {
  return (
    <>
      <StatusBar style="auto" />
      <Container>
        <Register />
      </Container>
    </>
  );
}
